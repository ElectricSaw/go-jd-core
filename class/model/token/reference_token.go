package token

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewReferenceToken(typ int, internalTypeName, name, descriptor, ownerInternalName string) intmod.IReferenceToken {
	return &ReferenceToken{
		DeclarationToken:  *NewDeclarationToken(typ, internalTypeName, name, descriptor).(*DeclarationToken),
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

func (t *ReferenceToken) SetOwnerInternalName(ownerInternalName string) {
	t.ownerInternalName = ownerInternalName
}

func (t *ReferenceToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitReferenceToken(t)
}

func (t *ReferenceToken) String() string {
	return fmt.Sprintf("ReferenceToken { '%s' }", t.ownerInternalName)
}
