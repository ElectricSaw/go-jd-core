package reference

import _type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"

func NewAnnotationReference(typ _type.ObjectType) *AnnotationReference {
	return &AnnotationReference{
		typ: typ,
	}
}

func NewAnnotationReferenceWithEv(typ _type.ObjectType, elementValue IElementValue) *AnnotationReference {
	return &AnnotationReference{
		typ:          typ,
		elementValue: elementValue,
	}
}

func NewAnnotationReferenceWithEvp(typ _type.ObjectType, elementValuePairs IElementValuePair) *AnnotationReference {
	return &AnnotationReference{
		typ:               typ,
		elementValuePairs: elementValuePairs,
	}
}

func NewAnnotationReferenceWithAll(typ _type.ObjectType, elementValue IElementValue, elementValuePairs IElementValuePair) *AnnotationReference {
	return &AnnotationReference{
		typ:               typ,
		elementValue:      elementValue,
		elementValuePairs: elementValuePairs,
	}
}

type AnnotationReference struct {
	typ               _type.ObjectType
	elementValue      IElementValue
	elementValuePairs IElementValuePair
}

func (r *AnnotationReference) GetType() _type.ObjectType {
	return r.typ
}

func (r *AnnotationReference) GetElementValue() IElementValue {
	return r.elementValue
}

func (r *AnnotationReference) GetElementValuePairs() IElementValuePair {
	return r.elementValuePairs
}

//func (r *AnnotationReference) Equals(other interface{}) bool {
//
//}

func (r *AnnotationReference) Accept(visitor ReferenceVisitor) {
	visitor.VisitAnnotationReference(r)
}
