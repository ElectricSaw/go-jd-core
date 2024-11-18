package expression

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewClassFileNewExpression(lineNumber int, typ _type.IObjectType) *ClassFileNewExpression {
	return &ClassFileNewExpression{
		NewExpression: *expression.NewNewExpression(lineNumber, typ, ""),
		bound:         false,
	}
}

func NewClassFileNewExpression2(lineNumber int, typ _type.IObjectType, bodyDeclaration declaration.BodyDeclaration) *ClassFileNewExpression {
	return &ClassFileNewExpression{
		NewExpression: *expression.NewNewExpressionWithAll(lineNumber, typ, "", bodyDeclaration),
		bound:         false,
	}
}

func NewClassFileNewExpression3(lineNumber int, typ _type.IObjectType, bodyDeclaration declaration.BodyDeclaration, bound bool) *ClassFileNewExpression {
	return &ClassFileNewExpression{
		NewExpression: *expression.NewNewExpressionWithAll(lineNumber, typ, "", bodyDeclaration),
		bound:         bound,
	}
}

type ClassFileNewExpression struct {
	expression.NewExpression

	parameterTypes _type.IType
	bound          bool
}

func (e *ClassFileNewExpression) ParameterTypes() _type.IType {
	return e.parameterTypes
}

func (e *ClassFileNewExpression) SetParameterTypes(parameterTypes _type.IType) {
	e.parameterTypes = parameterTypes
}

func (e *ClassFileNewExpression) IsBound() bool {
	return e.bound
}

func (e *ClassFileNewExpression) SetBound(bound bool) {
	e.bound = bound
}

func (e *ClassFileNewExpression) Set(descriptor string, parameterTypes _type.IType, parameters expression.IExpression) {
	e.SetDescriptor(descriptor)
	e.SetParameterTypes(parameterTypes)
	e.SetParameters(parameters)
}

func (e *ClassFileNewExpression) String() string {
	return fmt.Sprintf("ClassFileNewExpression{new %s}", e.Type())
}
