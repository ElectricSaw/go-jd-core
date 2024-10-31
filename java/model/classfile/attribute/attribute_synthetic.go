package attribute

func NewAttributeSynthetic() *AttributeSynthetic {
	return &AttributeSynthetic{}
}

type AttributeSynthetic struct{}

func (a AttributeSynthetic) attributeIgnoreFunc() {}
