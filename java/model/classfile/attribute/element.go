package attribute

type ElementValue interface {
	Accept(attribute ElementValueVisitor)
}

type ElementValueVisitor interface {
	VisitPrimitiveType(elementValue *ElementValuePrimitiveType)
	VisitClassInfo(elementValue *ElementValueClassInfo)
	VisitAnnotationValue(elementValue *ElementValueAnnotationValue)
	VisitEnumConstValue(elementValue *ElementValueEnumConstValue)
	VisitArrayValue(elementValue *ElementValueArrayValue)
}
