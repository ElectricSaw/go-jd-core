package reference

type IReference interface {
	Accept(visitor ReferenceVisitor)
}

type IAnnotationReference interface {
	IReference
}

type ReferenceVisitor interface {
	VisitAnnotationElementValue(reference *AnnotationElementValue)
	VisitAnnotationReference(reference *AnnotationReference)
	VisitAnnotationReferences(references *AnnotationReferences)
	VisitElementValueArrayInitializerElementValue(reference *ElementValueArrayInitializerElementValue)
	VisitElementValues(references *ElementValues)
	VisitElementValuePair(reference *ElementValuePair)
	VisitElementValuePairs(references *ElementValuePairs)
	VisitExpressionElementValue(reference *ExpressionElementValue)
	VisitInnerObjectReference(reference *InnerObjectReference)
	VisitObjectReference(reference *ObjectReference)
}
