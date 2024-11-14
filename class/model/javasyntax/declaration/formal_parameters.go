package declaration

type FormalParameters struct {
	FormalParameters []FormalParameter
}

func (d *FormalParameters) List() []Declaration {
	ret := make([]Declaration, 0, len(d.FormalParameters))
	for _, param := range d.FormalParameters {
		ret = append(ret, &param)
	}
	return ret
}

func (d *FormalParameters) Accept(visitor DeclarationVisitor) {
	visitor.VisitFormalParameters(d)
}
