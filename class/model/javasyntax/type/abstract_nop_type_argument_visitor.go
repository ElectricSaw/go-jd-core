package _type

type AbstractTypeArgumentVisitor struct {
}

func (v *AbstractTypeArgumentVisitor) VisitTypeArguments(arguments *TypeArguments) {
	arguments.AcceptTypeArgumentVisitor(v)
}

func (v *AbstractTypeArgumentVisitor) VisitDiamondTypeArgument(argument *DiamondTypeArgument) {
}

func (v *AbstractTypeArgumentVisitor) VisitWildcardExtendsTypeArgument(argument *WildcardExtendsTypeArgument) {
	argument.AcceptTypeArgumentVisitor(v)
}

func (v *AbstractTypeArgumentVisitor) VisitWildcardSuperTypeArgument(argument *WildcardSuperTypeArgument) {
}

func (v *AbstractTypeArgumentVisitor) VisitWildcardTypeArgument(argument *WildcardTypeArgument) {
}

func (v *AbstractTypeArgumentVisitor) VisitPrimitiveType(t *PrimitiveType) {
}

func (v *AbstractTypeArgumentVisitor) VisitObjectType(t *ObjectType) {
	if visitable, ok := t.TypeArguments().(TypeArgumentVisitable); ok {
		v.SafeAcceptTypeArgumentVisitable(visitable)
	}
}

func (v *AbstractTypeArgumentVisitor) VisitInnerObjectType(t *InnerObjectType) {
	t.outerType.AcceptTypeArgumentVisitor(v)
	if visitable, ok := t.TypeArguments().(TypeArgumentVisitable); ok {
		v.SafeAcceptTypeArgumentVisitable(visitable)
	}
}

func (v *AbstractTypeArgumentVisitor) VisitGenericType(t *GenericType) {
}

func (v *AbstractTypeArgumentVisitor) SafeAcceptTypeArgumentVisitable(visitable TypeArgumentVisitable) {
	if visitable != nil {
		visitable.AcceptTypeArgumentVisitor(v)
	}
}

type AbstractNopTypeArgumentVisitor struct {
}

func (v *AbstractNopTypeArgumentVisitor) VisitTypeArguments(arguments *TypeArguments)            {}
func (v *AbstractNopTypeArgumentVisitor) VisitDiamondTypeArgument(argument *DiamondTypeArgument) {}
func (v *AbstractNopTypeArgumentVisitor) VisitWildcardExtendsTypeArgument(argument *WildcardExtendsTypeArgument) {
}
func (v *AbstractNopTypeArgumentVisitor) VisitWildcardSuperTypeArgument(argument *WildcardSuperTypeArgument) {
}
func (v *AbstractNopTypeArgumentVisitor) VisitWildcardTypeArgument(argument *WildcardTypeArgument) {}
func (v *AbstractNopTypeArgumentVisitor) VisitPrimitiveType(t *PrimitiveType)                      {}
func (v *AbstractNopTypeArgumentVisitor) VisitObjectType(t *ObjectType)                            {}
func (v *AbstractNopTypeArgumentVisitor) VisitInnerObjectType(t *InnerObjectType)                  {}
func (v *AbstractNopTypeArgumentVisitor) VisitGenericType(t *GenericType)                          {}
