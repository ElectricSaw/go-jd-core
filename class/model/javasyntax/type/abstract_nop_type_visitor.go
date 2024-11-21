package _type

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

type AbstractNopTypeVisitor struct {
}

func (v *AbstractNopTypeVisitor) VisitPrimitiveType(y intmod.IPrimitiveType)     {}
func (v *AbstractNopTypeVisitor) VisitObjectType(y intmod.IObjectType)           {}
func (v *AbstractNopTypeVisitor) VisitInnerObjectType(y intmod.IInnerObjectType) {}
func (v *AbstractNopTypeVisitor) VisitTypes(types intmod.ITypes)                 {}
func (v *AbstractNopTypeVisitor) VisitGenericType(y intmod.IGenericType)         {}
