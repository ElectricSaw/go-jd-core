package token

import "fmt"

const (
	TypeToken        = 1
	FieldToken       = 2
	MethodToken      = 3
	ConstructorToken = 4
	PackageToken     = 5
	ModuleToken      = 6
)

func NewDeclarationToken(typ int, internalTypeName string, name string, descriptor string) *DeclarationToken {
	return &DeclarationToken{typ, internalTypeName, name, descriptor}
}

type DeclarationToken struct {
	typ              int
	internalTypeName string
	name             string
	descriptor       string
}

func (t *DeclarationToken) Type() int {
	return t.typ
}

func (t *DeclarationToken) InternalTypeName() string {
	return t.internalTypeName
}

func (t *DeclarationToken) Name() string {
	return t.name
}

func (t *DeclarationToken) Descriptor() string {
	return t.descriptor
}

func (t *DeclarationToken) Accept(visitor TokenVisitor) {
	visitor.VisitDeclarationToken(t)
}

func (t *DeclarationToken) String() string {
	return fmt.Sprintf("DeclarationToken { declaration='%s' }", t.name)
}
