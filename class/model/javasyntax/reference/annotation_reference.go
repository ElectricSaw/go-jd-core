package reference

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

func NewAnnotationReference(typ _type.IObjectType) *AnnotationReference {
	return &AnnotationReference{
		typ: typ,
	}
}

func NewAnnotationReferenceWithEv(typ _type.IObjectType, elementValue IElementValue) *AnnotationReference {
	return &AnnotationReference{
		typ:          typ,
		elementValue: elementValue,
	}
}

func NewAnnotationReferenceWithEvp(typ _type.IObjectType, elementValuePairs IElementValuePair) *AnnotationReference {
	return &AnnotationReference{
		typ:               typ,
		elementValuePairs: elementValuePairs,
	}
}

func NewAnnotationReferenceWithAll(typ _type.IObjectType, elementValue IElementValue, elementValuePairs IElementValuePair) *AnnotationReference {
	return &AnnotationReference{
		typ:               typ,
		elementValue:      elementValue,
		elementValuePairs: elementValuePairs,
	}
}

type AnnotationReference struct {
	IAnnotationReference

	typ               _type.IObjectType
	elementValue      IElementValue
	elementValuePairs IElementValuePair
}

func (r *AnnotationReference) Type() _type.IObjectType {
	return r.typ
}

func (r *AnnotationReference) ElementValue() IElementValue {
	return r.elementValue
}

func (r *AnnotationReference) ElementValuePairs() IElementValuePair {
	return r.elementValuePairs
}

//func (r *AnnotationReference) Equals(other interface{}) bool {
//
//}

func (r *AnnotationReference) Accept(visitor ReferenceVisitor) {
	visitor.VisitAnnotationReference(r)
}
