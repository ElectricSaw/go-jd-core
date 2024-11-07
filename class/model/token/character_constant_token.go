package token

import "fmt"

func NewCharacterConstantToken(character string, ownerInternalName string) *CharacterConstantToken {
	return &CharacterConstantToken{character, ownerInternalName}
}

type CharacterConstantToken struct {
	character         string
	ownerInternalName string
}

func (t *CharacterConstantToken) Character() string {
	return t.character
}

func (t *CharacterConstantToken) OwnerInternalName() string {
	return t.ownerInternalName
}

func (t *CharacterConstantToken) Accept(visitor TokenVisitor) {
	visitor.VisitCharacterConstantToken(t)
}

func (t *CharacterConstantToken) String() string {
	return fmt.Sprintf("CharacterConstantToken { '%s' }", t.character)
}
