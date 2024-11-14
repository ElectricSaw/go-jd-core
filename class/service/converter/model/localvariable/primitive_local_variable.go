package localvariable

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/class/service/converter/utils"
	"fmt"
)

func NewPrimitiveLocalVariable(index, offset int, typ *_type.PrimitiveType, name string) *PrimitiveLocalVariable {
	return &PrimitiveLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, name),
		flags:                 typ.Flags(),
	}
}

func NewPrimitiveLocalVariableWithVar(index, offset int, primitiveLocalVariable *PrimitiveLocalVariable) *PrimitiveLocalVariable {
	v := &PrimitiveLocalVariable{
		AbstractLocalVariable: *NewAbstractLocalVariable(index, offset, ""),
	}

	valueFlags := primitiveLocalVariable.flags

	if valueFlags&_type.FlagInt != 0 {
		v.flags = valueFlags
	} else if valueFlags&_type.FlagShort != 0 {
		v.flags = valueFlags | _type.FlagInt
	} else if valueFlags&_type.FlagChar != 0 {
		v.flags = valueFlags | _type.FlagInt | _type.FlagShort
	} else if valueFlags&_type.FlagByte != 0 {
		v.flags = valueFlags | _type.FlagInt | _type.FlagShort
	} else {
		v.flags = valueFlags
	}

	return v
}

type PrimitiveLocalVariable struct {
	AbstractLocalVariable

	flags int
}

func (v *PrimitiveLocalVariable) Type() _type.IType {
	switch v.flags {
	case _type.FlagBoolean:
		return _type.PtTypeBoolean
	case _type.FlagChar:
		return _type.PtTypeChar
	case _type.FlagFloat:
		return _type.PtTypeFloat
	case _type.FlagDouble:
		return _type.PtTypeDouble
	case _type.FlagByte:
		return _type.PtTypeByte
	case _type.FlagShort:
		return _type.PtTypeShort
	case _type.FlagInt:
		return _type.PtTypeInt
	case _type.FlagLong:
		return _type.PtTypeLong
	case _type.FlagVoid:
		return _type.PtTypeVoid
	}

	if v.flags == (_type.FlagChar | _type.FlagInt) {
		return _type.PtMaybeCharType
	}
	if v.flags == (_type.FlagChar | _type.FlagShort | _type.FlagInt) {
		return _type.PtMaybeShortType
	}
	if v.flags == (_type.FlagChar | _type.FlagChar | _type.FlagShort | _type.FlagInt) {
		return _type.PtMaybeByteType
	}
	if v.flags == (_type.FlagChar | _type.FlagByte | _type.FlagChar | _type.FlagShort | _type.FlagInt) {
		return _type.PtMaybeBooleanType
	}
	if v.flags == (_type.FlagChar | _type.FlagShort | _type.FlagInt) {
		return _type.PtMaybeNegativeByteType
	}
	if v.flags == (_type.FlagChar | _type.FlagInt) {
		return _type.PtMaybeNegativeShortType
	}
	if v.flags == (_type.FlagChar | _type.FlagByte | _type.FlagShort | _type.FlagInt) {
		return _type.PtMaybeNegativeBooleanType
	}

	return _type.PtTypeInt
}

func (v *PrimitiveLocalVariable) Dimension() int {
	return 0
}

func (v *PrimitiveLocalVariable) SetType(typ *_type.PrimitiveType) {
	v.flags = typ.Flags()
}

func (v *PrimitiveLocalVariable) Accept(visitor LocalVariableVisitor) {
	visitor.VisitPrimitiveLocalVariable(v)
}

func (v *PrimitiveLocalVariable) String() string {
	sb := "PrimitiveLocalVariable{"

	if v.flags&_type.FlagBoolean != 0 {
		sb += "boolean "
	}
	if v.flags&_type.FlagChar != 0 {
		sb += "char "
	}
	if v.flags&_type.FlagFloat != 0 {
		sb += "float "
	}
	if v.flags&_type.FlagDouble != 0 {
		sb += "double "
	}
	if v.flags&_type.FlagByte != 0 {
		sb += "byte "
	}
	if v.flags&_type.FlagShort != 0 {
		sb += "short "
	}
	if v.flags&_type.FlagInt != 0 {
		sb += "int "
	}
	if v.flags&_type.FlagLong != 0 {
		sb += "long "
	}
	if v.flags&_type.FlagVoid != 0 {
		sb += "void "
	}

	sb += fmt.Sprintf("%s, index=%d", v.name, v.index)

	if v.next != nil {
		sb += fmt.Sprintf(", next=%s", v.next)
	}

	sb += "}"

	return sb
}

func (v *PrimitiveLocalVariable) IsAssignableFrom(typeBounds map[string]_type.IType, typ _type.IType) bool {
	if typ.Dimension() == 0 && typ.IsPrimitiveType() {
		return (v.flags & (typ.(*_type.PrimitiveType).RightFlags())) != 0
	}
	return false
}

func (v *PrimitiveLocalVariable) TypeOnRight(typeBounds map[string]_type.IType, typ _type.IType) {
	if typ.IsPrimitiveType() {
		if typ.Dimension() == 0 {
			return
		}

		f := typ.(*_type.PrimitiveType).RightFlags()

		if v.flags&f != 0 {
			old := v.flags
			v.flags &= f

			if old != v.flags {
				v.FireChangeEvent(typeBounds)
			}
		}
	}
}

func (v *PrimitiveLocalVariable) TypeOnLeft(typeBounds map[string]_type.IType, typ _type.IType) {
	if typ.IsPrimitiveType() {
		if typ.Dimension() == 0 {
			return
		}

		f := typ.(*_type.PrimitiveType).LeftFlags()

		if v.flags&f != 0 {
			old := v.flags
			v.flags &= f

			if old != v.flags {
				v.FireChangeEvent(typeBounds)
			}
		}
	}
}

func (v *PrimitiveLocalVariable) IsAssignableFromWithVariable(typeBounds map[string]_type.IType, variable ILocalVariableReference) bool {
	if variable.IsPrimitiveLocalVariable() {
		variableFlags := variable.(*PrimitiveLocalVariable).flags
		typ := utils.GetPrimitiveTypeFromFlags(variableFlags)

		if typ != nil {
			variableFlags = typ.RightFlags()
		}

		return v.flags&variableFlags != 0
	}
	return false
}

func (v *PrimitiveLocalVariable) VariableOnRight(typeBounds map[string]_type.IType, variable ILocalVariableReference) {
	if variable.Dimension() == 0 {
		return
	}

	v.AddVariableOnRight(variable)

	old := v.flags
	variableFlags := variable.(*PrimitiveLocalVariable).flags
	typ := utils.GetPrimitiveTypeFromFlags(variableFlags)

	if typ != nil {
		variableFlags = typ.RightFlags()
	}

	v.flags &= variableFlags

	if old != v.flags {
		v.FireChangeEvent(typeBounds)
	}
}

func (v *PrimitiveLocalVariable) VariableOnLeft(typeBounds map[string]_type.IType, variable ILocalVariableReference) {
	if variable.Dimension() == 0 {
		return
	}

	v.AddVariableOnLeft(variable)

	old := v.flags
	variableFlags := variable.(*PrimitiveLocalVariable).flags
	typ := utils.GetPrimitiveTypeFromFlags(variableFlags)

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
