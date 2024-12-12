package token

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewCharacterConstantToken(character string, ownerInternalName string) intmod.ICharacterConstantToken {
	return &CharacterConstantToken{character, ownerInternalName}
}

type CharacterConstantToken struct {
	character         string
	ownerInternalName string
}

func (t *CharacterConstantToken) Character() string {
	return t.character
}

func (t *CharacterConstantToken) SetCharacter(character string) {
	t.character = character
}

func (t *CharacterConstantToken) OwnerInternalName() string {
	return t.ownerInternalName
}

func (t *CharacterConstantToken) SetOwnerInternalName(ownerInternalName string) {
	t.ownerInternalName = ownerInternalName
}

func (t *CharacterConstantToken) Accept(visitor intmod.ITokenVisitor) {
	visitor.VisitCharacterConstantToken(t)
}

func (t *CharacterConstantToken) String() string {
	return fmt.Sprintf("CharacterConstantToken { '%s' }", t.character)
}
