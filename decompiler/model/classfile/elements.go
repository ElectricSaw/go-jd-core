package classfile

import "fmt"

func NewElementValueAnnotationValue(annotationValue Annotation) ElementValueAnnotationValue {
	return ElementValueAnnotationValue{
		AnnotationValue: annotationValue,
	}
}

func NewElementValueArrayValue(values []IElementValue) ElementValueArrayValue {
	return ElementValueArrayValue{
		Values: values,
	}
}

func NewElementValueClassInfo(classInfo string) ElementValueClassInfo {
	return ElementValueClassInfo{
		ClassInfo: classInfo,
	}
}

func NewElementValueEnumConstValue(descriptor string, constName string) ElementValueEnumConstValue {
	return ElementValueEnumConstValue{
		Descriptor: descriptor,
		ConstName:  constName,
	}
}

func NewElementValuePair(elementName string, elementValue IElementValue) ElementValuePair {
	return ElementValuePair{
		ElementName:  elementName,
		ElementValue: elementValue,
	}
}

func NewElementValuePrimitiveType(type_ int, constValue IConstantValue) ElementValuePrimitiveType {
	return ElementValuePrimitiveType{
		Type:       type_,
		ConstValue: constValue,
	}
}

type IElementValue interface {
	Accept(attribute IElementValueVisitor)
	String() string
}

type IElementValueVisitor interface {
	VisitPrimitiveType(elementValue IElementValue)
	VisitClassInfo(elementValue IElementValue)
	VisitAnnotationValue(elementValue IElementValue)
	VisitEnumConstValue(elementValue IElementValue)
	VisitArrayValue(elementValue IElementValue)
}

type ElementValueAnnotationValue struct {
	AnnotationValue Annotation
}

func (e *ElementValueAnnotationValue) Accept(visitor IElementValueVisitor) {
	visitor.VisitAnnotationValue(e)
}

func (e *ElementValueAnnotationValue) String() string {
	return fmt.Sprintf("ElementValueAnnotationValue{ %s }", e.AnnotationValue)
}

type ElementValueArrayValue struct {
	Values []IElementValue
}

func (e *ElementValueArrayValue) Accept(visitor IElementValueVisitor) {
	visitor.VisitArrayValue(e)
}

func (e *ElementValueArrayValue) String() string {
	return fmt.Sprintf("ElementValueArrayValue{ values: %d }", len(e.Values))
}

type ElementValueClassInfo struct {
	ClassInfo string
}

func (e *ElementValueClassInfo) Accept(visitor IElementValueVisitor) {
	visitor.VisitClassInfo(e)
}

func (e *ElementValueClassInfo) String() string {
	return fmt.Sprintf("ElementValueClassInfo{ classinfo: %s }", e.ClassInfo)
}

type ElementValueEnumConstValue struct {
	Descriptor string
	ConstName  string
}

func (e *ElementValueEnumConstValue) Accept(visitor IElementValueVisitor) {
	visitor.VisitEnumConstValue(e)
}

func (e *ElementValueEnumConstValue) String() string {
	return fmt.Sprintf("ElementValueEnumConstValue{ descriptor: %s , constName: %s }", e.Descriptor, e.ConstName)
}

type ElementValuePair struct {
	ElementName  string
	ElementValue IElementValue
}

func (e *ElementValuePair) String() string {
	return fmt.Sprintf("ElementValuePair{ elementName: %s elementValue: %s }", e.ElementName, e.ElementValue)
}

type ElementValuePrimitiveType struct {
	/*
	 * type = {'B', 'D', 'F', 'I', 'J', 'S', 'Z', 'C', 's'}
	 */
	Type       int
	ConstValue IConstantValue
}

func (e *ElementValuePrimitiveType) Accept(visitor IElementValueVisitor) {
	visitor.VisitPrimitiveType(e)
}

func (e *ElementValuePrimitiveType) String() string {
	return fmt.Sprintf("ElementValuePrimitiveType{ type: %d constName: %s }", e.Type, e.ConstValue)
}
