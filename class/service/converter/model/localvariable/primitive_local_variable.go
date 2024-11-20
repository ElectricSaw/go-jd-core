package localvariable

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
	"fmt"
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
		typ := utils.GetPrimitiveTypeFromFlags(variableFlags)

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
	typ := utils.GetPrimitiveTypeFromFlags(variableFlags)

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
