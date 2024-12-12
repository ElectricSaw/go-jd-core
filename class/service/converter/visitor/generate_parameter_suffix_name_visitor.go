package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewGenerateParameterSuffixNameVisitor() *GenerateParameterSuffixNameVisitor {
	return &GenerateParameterSuffixNameVisitor{}
}

type GenerateParameterSuffixNameVisitor struct {
	_type.AbstractNopTypeArgumentVisitor

	suffix string
}

func (v *GenerateParameterSuffixNameVisitor) Suffix() string {
	return v.suffix
}

func (v *GenerateParameterSuffixNameVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) {
	switch t.JavaPrimitiveFlags() {
	case intmod.FlagByte:
		v.suffix = "Byte"
	case intmod.FlagChar:
		v.suffix = "Char"
	case intmod.FlagDouble:
		v.suffix = "Double"
	case intmod.FlagFloat:
		v.suffix = "Float"
	case intmod.FlagInt:
		v.suffix = "Int"
	case intmod.FlagLong:
		v.suffix = "Long"
	case intmod.FlagShort:
		v.suffix = "Short"
	case intmod.FlagBoolean:
		v.suffix = "Boolean"
	default:
		v.suffix = "Unknown"
	}
}

func (v *GenerateParameterSuffixNameVisitor) VisitObjectType(t intmod.IObjectType) {
	v.suffix = t.Name()
}

func (v *GenerateParameterSuffixNameVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	v.suffix = t.Name()
}

func (v *GenerateParameterSuffixNameVisitor) VisitGenericType(t intmod.IGenericType) {
	v.suffix = t.Name()
}
