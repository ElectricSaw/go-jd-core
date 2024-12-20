package attribute

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewAnnotation(descriptor string, elementValuePairs []intcls.IElementValuePair) intcls.IAnnotation {
	return &Annotation{
		descriptor:        descriptor,
		elementValuePairs: elementValuePairs,
	}
}

type Annotation struct {
	descriptor        string
	elementValuePairs []intcls.IElementValuePair
}

func (a Annotation) Descriptor() string {
	return a.descriptor
}

func (a Annotation) ElementValuePairs() []intcls.IElementValuePair {
	return a.elementValuePairs
}
