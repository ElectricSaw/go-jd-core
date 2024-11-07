package _type

func NewTypes() *Types {
	return &Types{}
}

type Types struct {
	AbstractType

	Types []IType
}

func (t *Types) Add(ty IType) {
	t.Types = append(t.Types, ty)
}

func (t *Types) IsTypes() bool {
	return true
}

func (t *Types) AcceptTypeVisitor(visitor TypeVisitor) {
	visitor.VisitTypes(t)
}