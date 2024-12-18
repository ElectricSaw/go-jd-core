package token

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewDeclarationToken(typ int, internalTypeName, name, descriptor string) intmod.IDeclarationToken {
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

func (t *DeclarationToken) SetType(typ int) {
	t.typ = typ
}

func (t *DeclarationToken) InternalTypeName() string {
	return t.internalTypeName
}

func (t *DeclarationToken) SetInternalTypeName(internalName string) {
	t.internalTypeName = internalName
}

func (t *DeclarationToken) Name() string {
	return t.name
}

func (t *DeclarationToken) SetName(name string) {
	t.name = name
}

func (t *DeclarationToken) Descriptor() string {
	return t.descriptor
}

func (t *DeclarationToken) SetDescriptor(descriptor string) {
	t.descriptor = descriptor
}

func (t *DeclarationToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitDeclarationToken(t)
}

func (t *DeclarationToken) String() string {
	return fmt.Sprintf("DeclarationToken { declaration='%s' }", t.name)
}
