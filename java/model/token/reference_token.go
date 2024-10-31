package token

import "fmt"

func NewReferenceToken(typ int, internalTypeName string, name string, descriptor string, ownerInternalName string) *ReferenceToken {
	return &ReferenceToken{
		DeclarationToken:  NewDeclarationToken(typ, internalTypeName, name, descriptor),
		ownerInternalName: ownerInternalName,
	}
}

type ReferenceToken struct {
	DeclarationToken

	ownerInternalName string
}

func (t *ReferenceToken) OwnerInternalName() string {
	return t.ownerInternalName
}

func (t *ReferenceToken) Accept(visitor TokenVisitor) {
	visitor.VisitReferenceToken(t)
}

func (t *ReferenceToken) String() string {
	return fmt.Sprintf("ReferenceToken { '%s' }", t.ownerInternalName)
}
