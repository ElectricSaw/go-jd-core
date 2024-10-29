package attribute

func NewAttributeSignature(signature string) AttributeSignature {
	return AttributeSignature{signature: signature}
}

type AttributeSignature struct {
	signature string
}

func (a AttributeSignature) Signature() string {
	return a.signature
}

func (a AttributeSignature) attributeIgnoreFunc() {}
