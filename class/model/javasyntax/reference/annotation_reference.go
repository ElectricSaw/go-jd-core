package reference

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewAnnotationReference(typ intsyn.IObjectType) intsyn.IAnnotationReference {
	return &AnnotationReference{
		typ: typ,
	}
}

func NewAnnotationReferenceWithEv(typ intsyn.IObjectType, elementValue intsyn.IElementValue) intsyn.IAnnotationReference {
	return &AnnotationReference{
		typ:          typ,
		elementValue: elementValue,
	}
}

func NewAnnotationReferenceWithEvp(typ intsyn.IObjectType, elementValuePairs intsyn.IElementValuePair) intsyn.IAnnotationReference {
	return &AnnotationReference{
		typ:               typ,
		elementValuePairs: elementValuePairs,
	}
}

func NewAnnotationReferenceWithAll(typ intsyn.IObjectType, elementValue intsyn.IElementValue,
	elementValuePairs intsyn.IElementValuePair) intsyn.IAnnotationReference {
	return &AnnotationReference{
		typ:               typ,
		elementValue:      elementValue,
		elementValuePairs: elementValuePairs,
	}
}

type AnnotationReference struct {
	typ               intsyn.IObjectType
	elementValue      intsyn.IElementValue
	elementValuePairs intsyn.IElementValuePair
}

func (r *AnnotationReference) Type() intsyn.IObjectType {
	return r.typ
}

func (r *AnnotationReference) ElementValue() intsyn.IElementValue {
	return r.elementValue
}

func (r *AnnotationReference) ElementValuePairs() intsyn.IElementValuePair {
	return r.elementValuePairs
}

func (r *AnnotationReference) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitAnnotationReference(r)
}
