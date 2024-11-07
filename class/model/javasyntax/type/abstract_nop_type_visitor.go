package _type

type AbstractNopTypeVisitor struct {
}

func (v *AbstractNopTypeVisitor) VisitPrimitiveType(y *PrimitiveType)     {}
func (v *AbstractNopTypeVisitor) VisitObjectType(y *ObjectType)           {}
func (v *AbstractNopTypeVisitor) VisitInnerObjectType(y *InnerObjectType) {}
func (v *AbstractNopTypeVisitor) VisitTypes(types *Types)                 {}
func (v *AbstractNopTypeVisitor) VisitGenericType(y *GenericType)         {}
