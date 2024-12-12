package reference

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewAnnotationReference(typ intmod.IObjectType) intmod.IAnnotationReference {
	return &AnnotationReference{
		typ: typ,
	}
}

func NewAnnotationReferenceWithEv(typ intmod.IObjectType, elementValue intmod.IElementValue) intmod.IAnnotationReference {
	return &AnnotationReference{
		typ:          typ,
		elementValue: elementValue,
	}
}

func NewAnnotationReferenceWithEvp(typ intmod.IObjectType, elementValuePairs intmod.IElementValuePair) intmod.IAnnotationReference {
	return &AnnotationReference{
		typ:               typ,
		elementValuePairs: elementValuePairs,
	}
}

func NewAnnotationReferenceWithAll(typ intmod.IObjectType, elementValue intmod.IElementValue,
	elementValuePairs intmod.IElementValuePair) intmod.IAnnotationReference {
	return &AnnotationReference{
		typ:               typ,
		elementValue:      elementValue,
		elementValuePairs: elementValuePairs,
	}
}

type AnnotationReference struct {
	typ               intmod.IObjectType
	elementValue      intmod.IElementValue
	elementValuePairs intmod.IElementValuePair
}

func (r *AnnotationReference) Type() intmod.IObjectType {
	return r.typ
}

func (r *AnnotationReference) ElementValue() intmod.IElementValue {
	return r.elementValue
}

func (r *AnnotationReference) ElementValuePairs() intmod.IElementValuePair {
	return r.elementValuePairs
}

func (r *AnnotationReference) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitAnnotationReference(r)
}
