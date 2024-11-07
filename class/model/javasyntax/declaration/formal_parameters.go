package declaration

type FormalParameters struct {
	FormalParameters []FormalParameter
}

func (d *FormalParameters) Accept(visitor DeclarationVisitor) {
	visitor.VisitFormalParameters(d)
}
