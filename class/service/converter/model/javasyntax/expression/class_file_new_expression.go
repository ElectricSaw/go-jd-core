package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	"fmt"
)

func NewClassFileNewExpression(lineNumber int, typ intmod.IObjectType) intsrv.IClassFileNewExpression {
	e := &ClassFileNewExpression{
		NewExpression: *expression.NewNewExpression(lineNumber, typ, "").(*expression.NewExpression),
		bound:         false,
	}
	e.SetValue(e)
	return e
}

func NewClassFileNewExpression2(lineNumber int, typ intmod.IObjectType, bodyDeclaration intmod.IBodyDeclaration) intsrv.IClassFileNewExpression {
	return &ClassFileNewExpression{
		NewExpression: *expression.NewNewExpressionWithAll(lineNumber, typ, "", bodyDeclaration).(*expression.NewExpression),
		bound:         false,
	}
}

func NewClassFileNewExpression3(lineNumber int, typ intmod.IObjectType, bodyDeclaration intmod.IBodyDeclaration, bound bool) intsrv.IClassFileNewExpression {
	return &ClassFileNewExpression{
		NewExpression: *expression.NewNewExpressionWithAll(lineNumber, typ, "", bodyDeclaration).(*expression.NewExpression),
		bound:         bound,
	}
}

type ClassFileNewExpression struct {
	expression.NewExpression

	parameterTypes intmod.IType
	bound          bool
}

func (e *ClassFileNewExpression) ParameterTypes() intmod.IType {
	return e.parameterTypes
}

func (e *ClassFileNewExpression) SetParameterTypes(parameterTypes intmod.IType) {
	e.parameterTypes = parameterTypes
}

func (e *ClassFileNewExpression) IsBound() bool {
	return e.bound
}

func (e *ClassFileNewExpression) SetBound(bound bool) {
	e.bound = bound
}

func (e *ClassFileNewExpression) Set(descriptor string, parameterTypes intmod.IType, parameters intmod.IExpression) {
	e.SetDescriptor(descriptor)
	e.SetParameterTypes(parameterTypes)
	e.SetParameters(parameters)
}

func (e *ClassFileNewExpression) String() string {
	return fmt.Sprintf("ClassFileNewExpression{new %s}", e.Type())
}
