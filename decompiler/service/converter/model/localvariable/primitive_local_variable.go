package localvariable

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
	"math"
)

func NewPrimitiveLocalVariable(index, offset int, typ intmod.IPrimitiveType, name string) intsrv.IPrimitiveLocalVariable {
	return &PrimitiveLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, name).(*AbstractLocalVariable),
		flags:                 typ.Flags(),
	}
}

func NewPrimitiveLocalVariableWithVar(index, offset int, primitiveLocalVariable intsrv.IPrimitiveLocalVariable) intsrv.IPrimitiveLocalVariable {
	v := &PrimitiveLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, "").(*AbstractLocalVariable),
	}

	valueFlags := primitiveLocalVariable.Flags()

	if valueFlags&intmod.FlagInt != 0 {
		v.flags = valueFlags
	} else if valueFlags&intmod.FlagShort != 0 {
		v.flags = valueFlags | intmod.FlagInt
	} else if valueFlags&intmod.FlagChar != 0 {
		v.flags = valueFlags | intmod.FlagInt | intmod.FlagShort
	} else if valueFlags&intmod.FlagByte != 0 {
		v.flags = valueFlags | intmod.FlagInt | intmod.FlagShort
	} else {
		v.flags = valueFlags
	}

	return v
}

type PrimitiveLocalVariable struct {
	AbstractLocalVariable

	flags int
}

func (v *PrimitiveLocalVariable) Flags() int {
	return v.flags
}

func (v *PrimitiveLocalVariable) SetFlags(flags int) {
	v.flags = flags
}

func (v *PrimitiveLocalVariable) Type() intmod.IType {
	switch v.flags {
	case intmod.FlagBoolean:
		return _type.PtTypeBoolean.(intmod.IType)
	case intmod.FlagChar:
		return _type.PtTypeChar.(intmod.IType)
	case intmod.FlagFloat:
		return _type.PtTypeFloat.(intmod.IType)
	case intmod.FlagDouble:
		return _type.PtTypeDouble.(intmod.IType)
	case intmod.FlagByte:
		return _type.PtTypeByte.(intmod.IType)
	case intmod.FlagShort:
		return _type.PtTypeShort.(intmod.IType)
	case intmod.FlagInt:
		return _type.PtTypeInt.(intmod.IType)
	case intmod.FlagLong:
		return _type.PtTypeLong.(intmod.IType)
	case intmod.FlagVoid:
		return _type.PtTypeVoid.(intmod.IType)
	}

	if v.flags == (intmod.FlagChar | intmod.FlagInt) {
		return _type.PtMaybeCharType.(intmod.IType)
	}
	if v.flags == (intmod.FlagChar | intmod.FlagShort | intmod.FlagInt) {
		return _type.PtMaybeShortType.(intmod.IType)
	}
	if v.flags == (intmod.FlagChar | intmod.FlagChar | intmod.FlagShort | intmod.FlagInt) {
		return _type.PtMaybeByteType.(intmod.IType)
	}
	if v.flags == (intmod.FlagChar | intmod.FlagByte | intmod.FlagChar | intmod.FlagShort | intmod.FlagInt) {
		return _type.PtMaybeBooleanType.(intmod.IType)
	}
	if v.flags == (intmod.FlagChar | intmod.FlagShort | intmod.FlagInt) {
		return _type.PtMaybeNegativeByteType.(intmod.IType)
	}
	if v.flags == (intmod.FlagChar | intmod.FlagInt) {
		return _type.PtMaybeNegativeShortType.(intmod.IType)
	}
	if v.flags == (intmod.FlagChar | intmod.FlagByte | intmod.FlagShort | intmod.FlagInt) {
		return _type.PtMaybeNegativeBooleanType.(intmod.IType)
	}

	return _type.PtTypeInt.(intmod.IType)
}

func (v *PrimitiveLocalVariable) Dimension() int {
	return 0
}

func (v *PrimitiveLocalVariable) SetType(typ intmod.IPrimitiveType) {
	v.flags = typ.Flags()
}

func (v *PrimitiveLocalVariable) Accept(visitor intsrv.ILocalVariableVisitor) {
	visitor.VisitPrimitiveLocalVariable(v)
}

func (v *PrimitiveLocalVariable) String() string {
	sb := "PrimitiveLocalVariable{"

	if v.flags&intmod.FlagBoolean != 0 {
		sb += "boolean "
	}
	if v.flags&intmod.FlagChar != 0 {
		sb += "char "
	}
	if v.flags&intmod.FlagFloat != 0 {
		sb += "float "
	}
	if v.flags&intmod.FlagDouble != 0 {
		sb += "double "
	}
	if v.flags&intmod.FlagByte != 0 {
		sb += "byte "
	}
	if v.flags&intmod.FlagShort != 0 {
		sb += "short "
	}
	if v.flags&intmod.FlagInt != 0 {
		sb += "int "
	}
	if v.flags&intmod.FlagLong != 0 {
		sb += "long "
	}
	if v.flags&intmod.FlagVoid != 0 {
		sb += "void "
	}

	sb += fmt.Sprintf("%s, index=%d", v.name, v.index)

	if v.next != nil {
		sb += fmt.Sprintf(", next=%s", v.next)
	}

	sb += "}"

	return sb
}

func (v *PrimitiveLocalVariable) IsAssignableFrom(typeBounds map[string]intmod.IType, typ intmod.IType) bool {
	if typ.Dimension() == 0 && typ.IsPrimitiveType() {
		return (v.flags & (typ.(intmod.IPrimitiveType).RightFlags())) != 0
	}
	return false
}

func (v *PrimitiveLocalVariable) TypeOnRight(typeBounds map[string]intmod.IType, typ intmod.IType) {
	if typ.IsPrimitiveType() {
		if typ.Dimension() == 0 {
			return
		}

		f := typ.(intmod.IPrimitiveType).RightFlags()

		if v.flags&f != 0 {
			old := v.flags
			v.flags &= f

			if old != v.flags {
				v.FireChangeEvent(typeBounds)
			}
		}
	}
}

func (v *PrimitiveLocalVariable) TypeOnLeft(typeBounds map[string]intmod.IType, typ intmod.IType) {
	if typ.IsPrimitiveType() {
		if typ.Dimension() == 0 {
			return
		}

		f := typ.(intmod.IPrimitiveType).LeftFlags()

		if v.flags&f != 0 {
			old := v.flags
			v.flags &= f

			if old != v.flags {
				v.FireChangeEvent(typeBounds)
			}
		}
	}
}

func (v *PrimitiveLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]intmod.IType, variable intsrv.ILocalVariable) bool {
	if variable.IsPrimitiveLocalVariable() {
		variableFlags := variable.(*PrimitiveLocalVariable).flags
		typ := GetPrimitiveTypeFromFlags(variableFlags)

		if typ != nil {
			variableFlags = typ.RightFlags()
		}

		return v.flags&variableFlags != 0
	}
	return false
}

func (v *PrimitiveLocalVariable) VariableOnRight(typeBounds map[string]intmod.IType, variable intsrv.ILocalVariable) {
	if variable.Dimension() == 0 {
		return
	}

	v.AddVariableOnRight(variable)

	old := v.flags
	variableFlags := variable.(*PrimitiveLocalVariable).flags
	typ := GetPrimitiveTypeFromFlags(variableFlags)

	if typ != nil {
		variableFlags = typ.RightFlags()
	}

	v.flags &= variableFlags

	if old != v.flags {
		v.FireChangeEvent(typeBounds)
	}
}

func (v *PrimitiveLocalVariable) VariableOnLeft(typeBounds map[string]intmod.IType, variable intsrv.ILocalVariable) {
	if variable.Dimension() == 0 {
		return
	}

	v.AddVariableOnLeft(variable)

	old := v.flags
	variableFlags := variable.(*PrimitiveLocalVariable).flags
	typ := GetPrimitiveTypeFromFlags(variableFlags)

	if typ != nil {
		variableFlags = typ.LeftFlags()
	}

	v.flags &= variableFlags

	if old != v.flags {
		v.FireChangeEvent(typeBounds)
	}
}

func (v *PrimitiveLocalVariable) IsPrimitiveLocalVariable() bool {
	return true
}

func GetPrimitiveTypeFromDescriptor(descriptor string) intmod.IType {
	dimension := 0

	for descriptor[dimension] == '[' {
		dimension++
	}

	if dimension == 0 {
		return _type.GetPrimitiveType(int(descriptor[dimension])).(intmod.IType)
	} else {
		return _type.NewObjectTypeWithDescAndDim(descriptor[dimension:], dimension).(intmod.IType)
	}
}

func GetPrimitiveTypeFromValue(value int) intmod.IPrimitiveType {
	if value >= 0 {
		if value <= 1 {
			return _type.PtMaybeBooleanType.(intmod.IPrimitiveType)
		}
		if value <= math.MaxInt8 {
			return _type.PtMaybeByteType.(intmod.IPrimitiveType)
		}
		if value <= math.MaxInt16 {
			return _type.PtMaybeShortType.(intmod.IPrimitiveType)
		}
		if value <= '\uFFFF' {
			return _type.PtMaybeCharType.(intmod.IPrimitiveType)
		}
	} else {
		if value >= math.MinInt8 {
			return _type.PtMaybeNegativeByteType.(intmod.IPrimitiveType)
		}
		if value >= math.MinInt16 {
			return _type.PtMaybeNegativeShortType.(intmod.IPrimitiveType)
		}
	}
	return _type.PtMaybeIntType
}

func GetCommonPrimitiveType(pt1, pt2 intmod.IPrimitiveType) intmod.IPrimitiveType {
	return GetPrimitiveTypeFromFlags(pt1.Flags() & pt2.Flags())
}

func GetPrimitiveTypeFromFlags(flags int) intmod.IPrimitiveType {
	switch flags {
	case intmod.FlagBoolean:
		return _type.PtTypeBoolean
	case intmod.FlagChar:
		return _type.PtTypeChar
	case intmod.FlagFloat:
		return _type.PtTypeFloat
	case intmod.FlagDouble:
		return _type.PtTypeDouble
	case intmod.FlagByte:
		return _type.PtTypeByte
	case intmod.FlagShort:
		return _type.PtTypeShort
	case intmod.FlagInt:
		return _type.PtTypeInt
	case intmod.FlagLong:
		return _type.PtTypeLong
	case intmod.FlagVoid:
		return _type.PtTypeVoid
	default:
		if flags == intmod.FlagChar|intmod.FlagInt {
			return _type.PtMaybeCharType
		}
		if flags == intmod.FlagChar|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeShortType
		}
		if flags == intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeByteType
		}
		if flags == intmod.FlagBoolean|intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeBooleanType
		}
		if flags == intmod.FlagByte|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeNegativeByteType
		}
		if flags == intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeNegativeShortType
		}
		if flags == intmod.FlagBoolean|intmod.FlagByte|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeNegativeBooleanType
		}
	}

	return nil
}

func GetPrimitiveTypeFromTag(tag int) intmod.IType {
	switch tag {
	case 4:
		return _type.PtTypeBoolean.(intmod.IType)
	case 5:
		return _type.PtTypeChar.(intmod.IType)
	case 6:
		return _type.PtTypeFloat.(intmod.IType)
	case 7:
		return _type.PtTypeDouble.(intmod.IType)
	case 8:
		return _type.PtTypeByte.(intmod.IType)
	case 9:
		return _type.PtTypeShort.(intmod.IType)
	case 10:
		return _type.PtTypeInt.(intmod.IType)
	case 11:
		return _type.PtTypeLong.(intmod.IType)
	default:
		return nil
	}
}
