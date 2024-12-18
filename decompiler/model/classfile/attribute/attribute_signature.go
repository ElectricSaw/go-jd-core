package attribute

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

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
