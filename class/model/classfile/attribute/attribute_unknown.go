package attribute

func NewUnknownAttribute() *UnknownAttribute {
	return &UnknownAttribute{}
}

type UnknownAttribute struct{}

func (a UnknownAttribute) attributeIgnoreFunc() {}
