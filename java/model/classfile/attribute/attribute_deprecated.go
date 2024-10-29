package attribute

func NewAttributeDeprecated() AttributeDeprecated {
	return AttributeDeprecated{}
}

type AttributeDeprecated struct {
}

func (a AttributeDeprecated) attributeIgnoreFunc() {}
