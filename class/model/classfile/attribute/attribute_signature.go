package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewAttributeSignature(signature string) intcls.IAttributeSignature {
	return &AttributeSignature{signature: signature}
}

type AttributeSignature struct {
	signature string
}

func (a AttributeSignature) Signature() string {
	return a.signature
}

func (a AttributeSignature) IsAttribute() bool {
	return true
}
