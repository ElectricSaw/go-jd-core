package localvariable

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"math"
)

func NewPrimitiveLocalVariable(index, offset int, typ model.PrimitiveType, name string) PrimitiveLocalVariable {
	return PrimitiveLocalVariable{
		declared:         offset == 0,
		index:            index,
		fromOffset:       offset,
		toOffset:         offset,
		name:             name,
		references:       util.NewDefaultList[ILocalVariable](),
		variablesOnRight: util.NewSet[ILocalVariable](),
		variablesOnLeft:  util.NewSet[ILocalVariable](),
		flags:            typ.Flags,
	}
}

func NewPrimitiveLocalVariableWithVar(index, offset int, primitiveLocalVariable PrimitiveLocalVariable) PrimitiveLocalVariable {
	v := PrimitiveLocalVariable{
		declared:         offset == 0,
		index:            index,
		fromOffset:       offset,
		toOffset:         offset,
		name:             "",
		references:       util.NewDefaultList[ILocalVariable](),
		variablesOnRight: util.NewSet[ILocalVariable](),
		variablesOnLeft:  util.NewSet[ILocalVariable](),
	}

	valueFlags := primitiveLocalVariable.Flags()

	if valueFlags&model.FlagInt != 0 {
		v.flags = valueFlags
	} else if valueFlags&model.FlagShort != 0 {
		v.flags = valueFlags | model.FlagInt
	} else if valueFlags&model.FlagChar != 0 {
		v.flags = valueFlags | model.FlagInt | model.FlagShort
	} else if valueFlags&model.FlagByte != 0 {
		v.flags = valueFlags | model.FlagInt | model.FlagShort
	} else {
		v.flags = valueFlags
	}

	return v
}

type PrimitiveLocalVariable struct {
	frame            IFrame
	next             ILocalVariable
	declared         bool
	index            int
	fromOffset       int
	toOffset         int
	name             string
	references       util.IList[ILocalVariable]
	variablesOnRight util.ISet[ILocalVariable]
	variablesOnLeft  util.ISet[ILocalVariable]
	flags            int
}

func (v *PrimitiveLocalVariable) Frame() IFrame {
	return v.frame
}

func (v *PrimitiveLocalVariable) Next() ILocalVariable {
	return v.next
}

func (v *PrimitiveLocalVariable) IsDeclared() bool {
	return v.declared
}

func (v *PrimitiveLocalVariable) Index() int {
	return v.index
}

func (v *PrimitiveLocalVariable) FromOffset() int {
	return v.fromOffset
}

func (v *PrimitiveLocalVariable) ToOffset() int {
	return v.toOffset
}

func (v *PrimitiveLocalVariable) Type() model.IType {
	switch v.flags {
	case model.FlagBoolean:
		return &model.PtTypeBoolean
	case model.FlagChar:
		return &model.PtTypeChar
	case model.FlagFloat:
		return &model.PtTypeFloat
	case model.FlagDouble:
		return &model.PtTypeDouble
	case model.FlagByte:
		return &model.PtTypeByte
	case model.FlagShort:
		return &model.PtTypeShort
	case model.FlagInt:
		return &model.PtTypeInt
	case model.FlagLong:
		return &model.PtTypeLong
	case model.FlagVoid:
		return &model.PtTypeVoid
	}

	if v.flags == (model.FlagChar | model.FlagInt) {
		return &model.PtMaybeCharType
	}
	if v.flags == (model.FlagChar | model.FlagShort | model.FlagInt) {
		return &model.PtMaybeShortType
	}
	if v.flags == (model.FlagChar | model.FlagChar | model.FlagShort | model.FlagInt) {
		return &model.PtMaybeByteType
	}
	if v.flags == (model.FlagChar | model.FlagByte | model.FlagChar | model.FlagShort | model.FlagInt) {
		return &model.PtMaybeBooleanType
	}
	if v.flags == (model.FlagChar | model.FlagShort | model.FlagInt) {
		return &model.PtMaybeNegativeByteType
	}
	if v.flags == (model.FlagChar | model.FlagInt) {
		return &model.PtMaybeNegativeShortType
	}
	if v.flags == (model.FlagChar | model.FlagByte | model.FlagShort | model.FlagInt) {
		return &model.PtMaybeNegativeBooleanType
	}

	return &model.PtTypeInt
}

func (v *PrimitiveLocalVariable) Name() string {
	return v.name
}

func (v *PrimitiveLocalVariable) Dimension() int {
	return 0
}

func (v *PrimitiveLocalVariable) Flags() int {
	return v.flags
}

func (v *PrimitiveLocalVariable) References() util.IList[ILocalVariable] {
	return v.references
}

func (v *PrimitiveLocalVariable) SetFrame(frame IFrame) {
	v.frame = frame
}

func (v *PrimitiveLocalVariable) SetNext(localVariable ILocalVariable) {
	v.next = localVariable
}

func (v *PrimitiveLocalVariable) SetDeclared(declared bool) {
	v.declared = declared

}

func (v *PrimitiveLocalVariable) SetIndex(index int) {
	v.index = index
}

func (v *PrimitiveLocalVariable) SetFromOffset(fromOffset int) {
	v.fromOffset = fromOffset
}

func (v *PrimitiveLocalVariable) SetToOffset(offset int) {
	if v.fromOffset > offset {
		v.fromOffset = offset
	}
	if v.toOffset < offset {
		v.toOffset = offset
	}
}

func (v *PrimitiveLocalVariable) SetToOffsetForce(offset int, force bool) {
	v.toOffset = offset
}

func (v *PrimitiveLocalVariable) SetType(t model.IType) {
	v.flags = t.(*model.PrimitiveType).Flags
}

func (v *PrimitiveLocalVariable) SetName(name string) {
	v.name = name
}

func (v *PrimitiveLocalVariable) SetDimension(dimension int) {
}

func (v *PrimitiveLocalVariable) SetFlags(flags int) {
	v.flags = flags
}

func (v *PrimitiveLocalVariable) Accept(visitor ILocalVariableVisitor) {
	visitor.VisitPrimitiveLocalVariable(v)
}

func (v *PrimitiveLocalVariable) AddReference(reference ILocalVariableReference) {
	v.references.Add(reference.(ILocalVariable))
}

func (v *PrimitiveLocalVariable) IsAssignableFrom(typeBounds map[string]model.IType, typ model.IType) bool {
	if typ.GetDimension() == 0 && typ.IsPrimitiveType() {
		return (v.flags & (typ.(*model.PrimitiveType).RightFlags)) != 0
	}
	return false
}

func (v *PrimitiveLocalVariable) TypeOnRight(typeBounds map[string]model.IType, typ model.IType) {
	if typ.IsPrimitiveType() {
		if typ.GetDimension() == 0 {
			return
		}

		f := typ.(*model.PrimitiveType).RightFlags

		if v.flags&f != 0 {
			old := v.flags
			v.flags &= f

			if old != v.flags {
				v.FireChangeEvent(typeBounds)
			}
		}
	}
}

func (v *PrimitiveLocalVariable) TypeOnLeft(typeBounds map[string]model.IType, typ model.IType) {
	if typ.IsPrimitiveType() {
		if typ.GetDimension() == 0 {
			return
		}

		f := typ.(*model.PrimitiveType).LeftFlags

		if v.flags&f != 0 {
			old := v.flags
			v.flags &= f

			if old != v.flags {
				v.FireChangeEvent(typeBounds)
			}
		}
	}
}

func (v *PrimitiveLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]model.IType, variable ILocalVariable) bool {
	if variable.IsPrimitiveLocalVariable() {
		variableFlags := variable.(*PrimitiveLocalVariable).flags
		typ := GetPrimitiveTypeFromFlags(variableFlags)

		if typ != nil {
			variableFlags = typ.RightFlags
		}

		return v.flags&variableFlags != 0
	}
	return false
}

func (v *PrimitiveLocalVariable) VariableOnRight(typeBounds map[string]model.IType, variable ILocalVariable) {
	if variable.Dimension() == 0 {
		return
	}

	v.AddVariableOnRight(variable)

	old := v.flags
	variableFlags := variable.(*PrimitiveLocalVariable).flags
	typ := GetPrimitiveTypeFromFlags(variableFlags)

	if typ != nil {
		variableFlags = typ.RightFlags
	}

	v.flags &= variableFlags

	if old != v.flags {
		v.FireChangeEvent(typeBounds)
	}
}

func (v *PrimitiveLocalVariable) VariableOnLeft(typeBounds map[string]model.IType, variable ILocalVariable) {
	if variable.Dimension() == 0 {
		return
	}

	v.AddVariableOnLeft(variable)

	old := v.flags
	variableFlags := variable.(*PrimitiveLocalVariable).flags
	typ := GetPrimitiveTypeFromFlags(variableFlags)

	if typ != nil {
		variableFlags = typ.LeftFlags
	}

	v.flags &= variableFlags

	if old != v.flags {
		v.FireChangeEvent(typeBounds)
	}
}

func (v *PrimitiveLocalVariable) FireChangeEvent(typeBounds map[string]model.IType) {
	if v.variablesOnLeft != nil {
		for _, variable := range v.variablesOnLeft.ToSlice() {
			v.VariableOnRight(typeBounds, variable)
		}
	}
	if v.variablesOnRight != nil {
		for _, variable := range v.variablesOnRight.ToSlice() {
			v.VariableOnLeft(typeBounds, variable)
		}
	}
}

func (v *PrimitiveLocalVariable) AddVariableOnLeft(variable ILocalVariable) {
	if v.variablesOnLeft == nil {
		v.variablesOnLeft = util.NewSet[ILocalVariable]()
		v.variablesOnLeft.Add(variable)
		variable.AddVariableOnRight(v)
	} else if !v.variablesOnLeft.Contains(variable) {
		v.variablesOnLeft.Add(variable)
		variable.AddVariableOnRight(v)
	}
}

func (v *PrimitiveLocalVariable) AddVariableOnRight(variable ILocalVariable) {
	if v.variablesOnRight == nil {
		v.variablesOnRight = util.NewSet[ILocalVariable]()
		v.variablesOnRight.Add(variable)
		variable.AddVariableOnLeft(v)
	} else if !v.variablesOnRight.Contains(variable) {
		v.variablesOnRight.Add(variable)
		variable.AddVariableOnLeft(v)
	}
}

func (v *PrimitiveLocalVariable) IsPrimitiveLocalVariable() bool {
	return true
}

func (v *PrimitiveLocalVariable) GetLocalVariable() ILocalVariableReference {
	return nil
}

func (v *PrimitiveLocalVariable) SetLocalVariable(_ ILocalVariableReference) {

}

func (v *PrimitiveLocalVariable) String() string {
	sb := "PrimitiveLocalVariable{"

	if v.flags&model.FlagBoolean != 0 {
		sb += "boolean "
	}
	if v.flags&model.FlagChar != 0 {
		sb += "char "
	}
	if v.flags&model.FlagFloat != 0 {
		sb += "float "
	}
	if v.flags&model.FlagDouble != 0 {
		sb += "double "
	}
	if v.flags&model.FlagByte != 0 {
		sb += "byte "
	}
	if v.flags&model.FlagShort != 0 {
		sb += "short "
	}
	if v.flags&model.FlagInt != 0 {
		sb += "int "
	}
	if v.flags&model.FlagLong != 0 {
		sb += "long "
	}
	if v.flags&model.FlagVoid != 0 {
		sb += "void "
	}

	sb += fmt.Sprintf("%s, index=%d", v.name, v.index)

	if v.next != nil {
		sb += fmt.Sprintf(", next=%s", v.next)
	}

	sb += "}"

	return sb
}

func GetPrimitiveTypeFromDescriptor(descriptor string) model.IType {
	dimension := 0

	for descriptor[dimension] == '[' {
		dimension++
	}

	if dimension == 0 {
		return model.GetPrimitiveType(int(descriptor[dimension]))
	} else {
		return model.NewObjectTypeWithDescAndDim(descriptor[dimension:], dimension).(model.IType)
	}
}

func GetPrimitiveTypeFromValue(value int) *model.PrimitiveType {
	if value >= 0 {
		if value <= 1 {
			return &model.PtMaybeBooleanType
		}
		if value <= math.MaxInt8 {
			return &model.PtMaybeByteType
		}
		if value <= math.MaxInt16 {
			return &model.PtMaybeShortType
		}
		if value <= '\uFFFF' {
			return &model.PtMaybeCharType
		}
	} else {
		if value >= math.MinInt8 {
			return &model.PtMaybeNegativeByteType
		}
		if value >= math.MinInt16 {
			return &model.PtMaybeNegativeShortType
		}
	}
	return &model.PtMaybeIntType
}

func GetCommonPrimitiveType(pt1, pt2 *model.PrimitiveType) *model.PrimitiveType {
	return GetPrimitiveTypeFromFlags(pt1.Flags & pt2.Flags)
}

func GetPrimitiveTypeFromFlags(flags int) *model.PrimitiveType {
	switch flags {
	case model.FlagBoolean:
		return &model.PtTypeBoolean
	case model.FlagChar:
		return &model.PtTypeChar
	case model.FlagFloat:
		return &model.PtTypeFloat
	case model.FlagDouble:
		return &model.PtTypeDouble
	case model.FlagByte:
		return &model.PtTypeByte
	case model.FlagShort:
		return &model.PtTypeShort
	case model.FlagInt:
		return &model.PtTypeInt
	case model.FlagLong:
		return &model.PtTypeLong
	case model.FlagVoid:
		return &model.PtTypeVoid
	default:
		if flags == model.FlagChar|model.FlagInt {
			return &model.PtMaybeCharType
		}
		if flags == model.FlagChar|model.FlagShort|model.FlagInt {
			return &model.PtMaybeShortType
		}
		if flags == model.FlagByte|model.FlagChar|model.FlagShort|model.FlagInt {
			return &model.PtMaybeByteType
		}
		if flags == model.FlagBoolean|model.FlagByte|model.FlagChar|model.FlagShort|model.FlagInt {
			return &model.PtMaybeBooleanType
		}
		if flags == model.FlagByte|model.FlagShort|model.FlagInt {
			return &model.PtMaybeNegativeByteType
		}
		if flags == model.FlagShort|model.FlagInt {
			return &model.PtMaybeNegativeShortType
		}
		if flags == model.FlagBoolean|model.FlagByte|model.FlagShort|model.FlagInt {
			return &model.PtMaybeNegativeBooleanType
		}
	}

	return nil
}

func GetPrimitiveTypeFromTag(tag int) model.IType {
	switch tag {
	case 4:
		return &model.PtTypeBoolean
	case 5:
		return &model.PtTypeChar
	case 6:
		return &model.PtTypeFloat
	case 7:
		return &model.PtTypeDouble
	case 8:
		return &model.PtTypeByte
	case 9:
		return &model.PtTypeShort
	case 10:
		return &model.PtTypeInt
	case 11:
		return &model.PtTypeLong
	default:
		return nil
	}
}
