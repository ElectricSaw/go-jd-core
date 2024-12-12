package utils

import (
	intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/expression"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/reference"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewAnnotationConverter(typeMaker intsrv.ITypeMaker) intsrv.IAnnotationConverter {
	return &AnnotationConverter{
		TypeMaker: typeMaker,
	}
}

type AnnotationConverter struct {
	TypeMaker    intsrv.ITypeMaker
	ElementValue intmod.IElementValue
}

func (c *AnnotationConverter) ConvertWithAnnotations2(visibles, invisibles intcls.IAnnotations) intmod.IAnnotationReference {
	if visibles == nil {
		if invisibles == nil {
			return nil
		} else {
			return c.ConvertWithAnnotations(invisibles)
		}
	} else {
		if invisibles == nil {
			return c.ConvertWithAnnotations(visibles)
		} else {
			aral := reference.NewAnnotationReferences()

			for _, a := range visibles.Annotations() {
				aral.Add(c.ConvertWithAnnotation(a))
			}

			for _, a := range invisibles.Annotations() {
				aral.Add(c.ConvertWithAnnotation(a))
			}

			return aral
		}
	}
}

func (c *AnnotationConverter) ConvertWithAnnotations(annotations intcls.IAnnotations) intmod.IAnnotationReference {
	as := annotations.Annotations()

	if len(as) == 1 {
		return c.ConvertWithAnnotation(as[0])
	} else {
		aral := reference.NewAnnotationReferences()
		for _, a := range as {
			aral.Add(c.ConvertWithAnnotation(a))
		}

		return aral
	}
}

func (c *AnnotationConverter) ConvertWithAnnotation(annotation intcls.IAnnotation) intmod.IAnnotationReference {
	descriptor := annotation.Descriptor()
	ot := c.TypeMaker.MakeFromDescriptor(descriptor)
	elementValuePairs := annotation.ElementValuePairs()

	if elementValuePairs == nil {
		return reference.NewAnnotationReference(ot)
	} else if len(elementValuePairs) == 1 {
		elementValuePair := elementValuePairs[0]
		elementName := elementValuePair.Name()
		elementValue := elementValuePair.Value()

		if elementName == "name" {
			return reference.NewAnnotationReferenceWithEv(ot, c.ConvertWithElementValue(elementValue))
		} else {
			return reference.NewAnnotationReferenceWithEv(ot, reference.NewElementValuePair(elementName, c.ConvertWithElementValue(elementValue)))
		}
	} else {
		list := reference.NewElementValuePairs()

		for _, elementValuePair := range elementValuePairs {
			elementName := elementValuePair.Name()
			elementValue := elementValuePair.Value()
			list.Add(reference.NewElementValuePair(elementName, c.ConvertWithElementValue(elementValue)))
		}

		return reference.NewAnnotationReferenceWithEv(ot, list)
	}
}

func (c *AnnotationConverter) ConvertWithElementValue(ev intcls.IElementValue) intmod.IElementValue {
	ev.Accept(c)
	return c.ElementValue
}

func (c *AnnotationConverter) VisitPrimitiveType(elementValue intcls.IElementValuePrimitiveType) {
	switch elementValue.Type() {
	case 'B':
		c.ElementValue = reference.NewExpressionElementValue(
			expression.NewIntegerConstantExpression(_type.PtTypeByte,
				elementValue.Value().(intcls.IConstantInteger).Value()))
	case 'D':
		c.ElementValue = reference.NewExpressionElementValue(
			expression.NewDoubleConstantExpression(
				elementValue.Value().(intcls.IConstantDouble).Value()))
	case 'F':
		c.ElementValue = reference.NewExpressionElementValue(
			expression.NewFloatConstantExpression(
				elementValue.Value().(intcls.IConstantFloat).Value()))
	case 'I':
		c.ElementValue = reference.NewExpressionElementValue(
			expression.NewIntegerConstantExpression(_type.PtTypeInt,
				elementValue.Value().(intcls.IConstantInteger).Value()))
	case 'J':
		c.ElementValue = reference.NewExpressionElementValue(
			expression.NewLongConstantExpression(
				elementValue.Value().(intcls.IConstantLong).Value()))
	case 'S':
		c.ElementValue = reference.NewExpressionElementValue(
			expression.NewIntegerConstantExpression(_type.PtTypeShort,
				elementValue.Value().(intcls.IConstantInteger).Value()))
	case 'Z':
		c.ElementValue = reference.NewExpressionElementValue(
			expression.NewIntegerConstantExpression(_type.PtTypeBoolean,
				elementValue.Value().(intcls.IConstantInteger).Value()))
	case 'C':
		c.ElementValue = reference.NewExpressionElementValue(
			expression.NewIntegerConstantExpression(_type.PtTypeChar,
				elementValue.Value().(intcls.IConstantInteger).Value()))
	case 's':
		c.ElementValue = reference.NewExpressionElementValue(
			expression.NewStringConstantExpression(
				elementValue.Value().(intcls.IConstantUtf8).Value()))
	}
}

func (c *AnnotationConverter) VisitClassInfo(elementValue intcls.IElementValueClassInfo) {
	classInfo := elementValue.ClassInfo()
	ot := c.TypeMaker.MakeFromDescriptor(classInfo)
	c.ElementValue = reference.NewExpressionElementValue(expression.NewTypeReferenceDotClassExpression(ot))
}

func (c *AnnotationConverter) VisitAnnotationValue(elementValue intcls.IElementValueAnnotationValue) {
	annotationValue := elementValue.AnnotationValue()
	annotationReference := c.ConvertWithAnnotation(annotationValue)
	c.ElementValue = reference.NewAnnotationElementValue(annotationReference)
}

func (c *AnnotationConverter) VisitEnumConstValue(elementValue intcls.IElementValueEnumConstValue) {
	descriptor := elementValue.Descriptor()
	ot := c.TypeMaker.MakeFromDescriptor(descriptor)
	constName := elementValue.ConstName()
	internalTypeName := descriptor[1 : len(descriptor)-1]
	c.ElementValue = reference.NewExpressionElementValue(
		expression.NewFieldReferenceExpression(ot,
			expression.NewObjectTypeReferenceExpression(ot),
			internalTypeName, constName, descriptor,
		),
	)
}

func (c *AnnotationConverter) VisitArrayValue(elementValue intcls.IElementValueArrayValue) {
	values := elementValue.Values()

	if values == nil {
		c.ElementValue = reference.NewElementValueArrayInitializerElementValueEmpty()
	} else if len(values) == 1 {
		values[0].Accept(c)
		c.ElementValue = reference.NewElementValueArrayInitializerElementValue(c.ElementValue)
	} else {
		list := reference.NewElementValues()
		for _, value := range values {
			value.Accept(c)
			list.Add(c.ElementValue)
		}

		c.ElementValue = reference.NewElementValueArrayInitializerElementValues(list)
	}
}
