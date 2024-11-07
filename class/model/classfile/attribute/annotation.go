package attribute

func NewAnnotation(descriptor string, elementValuePairs []ElementValuePair) *Annotation {
	return &Annotation{
		descriptor:        descriptor,
		elementValuePairs: elementValuePairs,
	}
}

type Annotation struct {
	descriptor        string
	elementValuePairs []ElementValuePair
}

func (a Annotation) Descriptor() string {
	return a.descriptor
}

func (a Annotation) ElementValuePairs() []ElementValuePair {
	return a.elementValuePairs
}
