package model

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"math"
)

/////////////////////////////////////////////////////////////////////////
//  New Functions
/////////////////////////////////////////////////////////////////////////

func NewArrayExpression(expression, index IExpression) ArrayExpression {
	return NewArrayExpressionWithAll(UnknownLineNumber, expression, index)
}

func NewArrayExpressionWithAll(lineNumber int, expression, index IExpression) ArrayExpression {
	e := ArrayExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		lineNumber:  lineNumber,
		typ:         CreateItemType(expression),
		priority:    1,
		expression:  expression,
		index:       index,
	}
	e.SetValue(&e)
	return e
}

func NewBinaryOperatorExpression(lineNumber int, typ IType, leftExpression IExpression,
	operator string, rightExpression IExpression, priority int) BinaryOperatorExpression {
	e := BinaryOperatorExpression{
		DefaultBase:     *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:      lineNumber,
		Type:            typ,
		LeftExpression:  leftExpression,
		Operator:        operator,
		RightExpression: rightExpression,
		Priority:        priority,
	}
	e.SetValue(&e)
	return e
}

var True = NewBooleanExpression(true)
var False = NewBooleanExpression(false)

func NewBooleanExpression(value bool) BooleanExpression {
	return NewBooleanExpressionWithLineNumber(0, value)
}

func NewBooleanExpressionWithLineNumber(lineNumber int, value bool) BooleanExpression {
	e := BooleanExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        nil,
		Priority:    0,
		Value:       value,
	}
	e.SetValue(&e)
	return e
}

func NewCastExpression(typ IType, expression IExpression) CastExpression {
	return NewCastExpressionWithAll(UnknownLineNumber, typ, expression, true)
}

func NewCastExpressionWithLineNumber(lineNumber int, typ IType, expression IExpression) CastExpression {
	return NewCastExpressionWithAll(lineNumber, typ, expression, true)
}

func NewCastExpressionWithAll(lineNumber int, typ IType, expression IExpression, explicit bool) CastExpression {
	e := CastExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        typ,
		Priority:    3,
		Expression:  expression,
		IsExplicit:  explicit,
	}
	e.SetValue(&e)
	return e
}

func NewCommentExpression(text string) CommentExpression {
	e := CommentExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		Text:        text,
	}
	e.SetValue(&e)
	return e
}

func NewConstructorInvocationExpression(objectType *ObjectType, descriptor string,
	parameters IExpression) ConstructorInvocationExpression {
	return NewConstructorInvocationExpressionWithAll(0, objectType, descriptor, parameters)
}

func NewConstructorInvocationExpressionWithAll(lineNumber int, objectType *ObjectType,
	descriptor string, parameters IExpression) ConstructorInvocationExpression {
	e := ConstructorInvocationExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        &PtTypeVoid,
		Priority:    1,
		ObjectType:  objectType,
		Descriptor:  descriptor,
		Parameters:  parameters,
	}
	e.SetValue(&e)
	return e
}

func NewConstructorReferenceExpression(typ IType, objectType *ObjectType,
	descriptor string) ConstructorReferenceExpression {
	return NewConstructorReferenceExpressionWithAll(UnknownLineNumber, typ, objectType, descriptor)
}

func NewConstructorReferenceExpressionWithAll(lineNumber int, typ IType,
	objectType *ObjectType, descriptor string) ConstructorReferenceExpression {
	e := ConstructorReferenceExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        typ,
		Priority:    0,
		ObjectType:  objectType,
		Descriptor:  descriptor,
	}
	e.SetValue(&e)
	return e
}

func NewDoubleConstantExpression(value float64) DoubleConstantExpression {
	return NewDoubleConstantExpressionWithAll(UnknownLineNumber, value)
}

func NewDoubleConstantExpressionWithAll(lineNumber int, value float64) DoubleConstantExpression {
	e := DoubleConstantExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        &PtTypeDouble,
		Priority:    0,
		Value:       value,
	}
	e.SetValue(&e)
	return e
}

func NewEnumConstantReferenceExpression(typ *ObjectType, name string) EnumConstantReferenceExpression {
	return NewEnumConstantReferenceExpressionWithAll(UnknownLineNumber, typ, name)
}

func NewEnumConstantReferenceExpressionWithAll(lineNumber int, typ *ObjectType, name string) EnumConstantReferenceExpression {
	e := EnumConstantReferenceExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        typ,
		Priority:    0,
		Name:        name,
	}
	e.SetValue(&e)
	return e
}

func NewExpressions() Expressions {
	return NewExpressionsWithCapacity(0)
}

func NewExpressionsWithCapacity(capacity int) Expressions {
	return Expressions{
		DefaultList: *util.NewDefaultListWithCapacity[IExpression](capacity).(*util.DefaultList[IExpression]),
	}
}

func NewExpressionsWithElements(elements ...IExpression) Expressions {
	return Expressions{
		DefaultList: *util.NewDefaultListWithElements[IExpression](elements...).(*util.DefaultList[IExpression]),
	}
}

func NewFieldReferenceExpression(typ IType, expression IExpression,
	internalTypeName, name, descriptor string) FieldReferenceExpression {
	return NewFieldReferenceExpressionWithAll(UnknownLineNumber, typ, expression, internalTypeName, name, descriptor)
}

func NewFieldReferenceExpressionWithAll(lineNumber int, typ IType, expression IExpression,
	internalTypeName, name, descriptor string) FieldReferenceExpression {
	e := FieldReferenceExpression{
		DefaultBase:      *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:       lineNumber,
		Type:             typ,
		Priority:         0,
		Expression:       expression,
		InternalTypeName: internalTypeName,
		Name:             name,
		Descriptor:       descriptor,
	}
	e.SetValue(&e)
	return e
}

func NewFloatConstantExpression(value float32) FloatConstantExpression {
	return NewFloatConstantExpressionWithAll(UnknownLineNumber, value)
}

func NewFloatConstantExpressionWithAll(lineNumber int, value float32) FloatConstantExpression {
	e := FloatConstantExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        &PtTypeFloat,
		Priority:    0,
		value:       value,
	}
	e.SetValue(&e)
	return e
}

func NewInstanceOfExpression(expression IExpression,
	instanceOfType *ObjectType) InstanceOfExpression {
	return NewInstanceOfExpressionWithAll(UnknownLineNumber, expression, instanceOfType)
}

func NewInstanceOfExpressionWithAll(lineNumber int, expression IExpression,
	instanceOfType *ObjectType) InstanceOfExpression {
	e := InstanceOfExpression{
		LineNumber:     lineNumber,
		Type:           &PtTypeBoolean,
		Priority:       8,
		expression:     expression,
		instanceOfType: instanceOfType,
	}
	e.SetValue(&e)
	return e
}

func NewIntegerConstantExpression(typ IType, value int) IntegerConstantExpression {
	return NewIntegerConstantExpressionWithAll(UnknownLineNumber, typ, value)
}

func NewIntegerConstantExpressionWithAll(lineNumber int, typ IType, value int) IntegerConstantExpression {
	e := IntegerConstantExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        typ,
		Priority:    0,
		Value:       value,
	}
	e.SetValue(&e)
	return e
}

func NewLambdaFormalParametersExpression(typ IType, formalParameters *FormalParameter,
	statements IStatement) LambdaFormalParametersExpression {
	return NewLambdaFormalParametersExpressionWithAll(0, typ, formalParameters, statements)
}

func NewLambdaFormalParametersExpressionWithAll(lineNumber int, typ IType,
	formalParameters *FormalParameter, statements IStatement) LambdaFormalParametersExpression {
	e := LambdaFormalParametersExpression{
		DefaultBase:      *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:       lineNumber,
		Type:             typ,
		Priority:         0,
		Statements:       statements,
		FormalParameters: formalParameters,
	}
	e.SetValue(&e)
	return e
}

func NewLambdaIdentifiersExpression(typ IType, returnedType IType,
	paramNames util.IList[string], statements IStatement) LambdaIdentifiersExpression {
	return NewLambdaIdentifiersExpressionWithAll(UnknownLineNumber, typ, returnedType, paramNames, statements)
}

func NewLambdaIdentifiersExpressionWithAll(lineNumber int, typ IType,
	returnedType IType, paramNames util.IList[string], statements IStatement) LambdaIdentifiersExpression {
	e := LambdaIdentifiersExpression{
		DefaultBase:    *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:     lineNumber,
		Type:           typ,
		Priority:       0,
		Statements:     statements,
		ReturnedType:   returnedType,
		ParameterNames: paramNames,
	}
	e.SetValue(&e)
	return e
}

func NewLengthExpression(expression IExpression) LengthExpression {
	return NewLengthExpressionWithAll(UnknownLineNumber, expression)
}

func NewLengthExpressionWithAll(lineNumber int, expression IExpression) LengthExpression {
	e := LengthExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        &PtTypeInt,
		Priority:    0,
		Expression:  expression,
	}
	e.SetValue(&e)
	return e
}

func NewLocalVariableReferenceExpression(typ IType, name string) LocalVariableReferenceExpression {
	return NewLocalVariableReferenceExpressionWithAll(UnknownLineNumber, typ, name)
}

func NewLocalVariableReferenceExpressionWithAll(lineNumber int, typ IType, name string) LocalVariableReferenceExpression {
	e := LocalVariableReferenceExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        typ,
		Priority:    0,
		Name:        name,
	}
	e.SetValue(&e)
	return e
}

func NewLongConstantExpression(value int64) LongConstantExpression {
	return NewLongConstantExpressionWithAll(UnknownLineNumber, value)
}

func NewLongConstantExpressionWithAll(lineNumber int, value int64) LongConstantExpression {
	e := LongConstantExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        &PtTypeLong,
		Priority:    0,
		Value:       value,
	}
	e.SetValue(&e)
	return e
}

func NewMethodInvocationExpression(typ IType, expression IExpression,
	internalTypeName, name, descriptor string) MethodInvocationExpression {
	return NewMethodInvocationExpressionWithAll(UnknownLineNumber, typ, expression, internalTypeName, name, descriptor, nil)
}

func NewMethodInvocationExpressionWithLineNumber(lineNumber int, typ IType,
	expression IExpression, internalTypeName, name, descriptor string) MethodInvocationExpression {
	return NewMethodInvocationExpressionWithAll(lineNumber, typ, expression, internalTypeName, name, descriptor, nil)
}

func NewMethodInvocationExpressionWithParam(typ IType, expression IExpression,
	internalTypeName, name, descriptor string, parameters IExpression) MethodInvocationExpression {
	return NewMethodInvocationExpressionWithAll(UnknownLineNumber, typ, expression, internalTypeName, name, descriptor, parameters)
}

func NewMethodInvocationExpressionWithAll(lineNumber int, typ IType, expression IExpression,
	internalTypeName, name, descriptor string, parameters IExpression) MethodInvocationExpression {
	e := MethodInvocationExpression{
		DefaultBase:      *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:       lineNumber,
		Type:             typ,
		Priority:         1,
		Expression:       expression,
		InternalTypeName: internalTypeName,
		Name:             name,
		Descriptor:       descriptor,
		Parameters:       parameters,
	}
	e.SetValue(&e)
	return e
}

func NewMethodReferenceExpression(typ IType, expression IExpression,
	internalTypeName, name, descriptor string) MethodReferenceExpression {
	return NewMethodReferenceExpressionWithAll(UnknownLineNumber, typ, expression, internalTypeName, name, descriptor)
}

func NewMethodReferenceExpressionWithAll(lineNumber int, typ IType, expression IExpression,
	internalTypeName, name, descriptor string) MethodReferenceExpression {
	e := MethodReferenceExpression{
		DefaultBase:      *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:       lineNumber,
		Type:             typ,
		Priority:         0,
		Expression:       expression,
		InternalTypeName: internalTypeName,
		Name:             name,
		Descriptor:       descriptor,
	}
	e.SetValue(&e)
	return e
}

func NewNewArray(lineNumber int, typ IType, dimensionExpressionList IExpression) NewArray {
	e := NewArray{
		DefaultBase:             *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:              lineNumber,
		Type:                    typ,
		Priority:                0,
		DimensionExpressionList: dimensionExpressionList,
	}
	e.SetValue(&e)
	return e
}

func NewNewExpression(lineNumber int, typ *ObjectType, descriptor string) NewExpression {
	return NewNewExpressionWithAll(lineNumber, typ, descriptor, nil)
}

func NewNewExpressionWithAll(lineNumber int, typ *ObjectType, descriptor string, bodyDeclaration *BodyDeclaration) NewExpression {
	e := NewExpression{
		DefaultBase:     *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:      lineNumber,
		Type:            typ,
		Priority:        0,
		Descriptor:      descriptor,
		BodyDeclaration: bodyDeclaration,
	}
	e.SetValue(&e)
	return e
}

func NewNewInitializedArray(typ IType, arrayInitializer ArrayVariableInitializer) NewInitializedArray {
	return NewNewInitializedArrayWithAll(UnknownLineNumber, typ, arrayInitializer)
}

func NewNewInitializedArrayWithAll(lineNumber int, typ IType, arrayInitializer ArrayVariableInitializer) NewInitializedArray {
	e := NewInitializedArray{
		DefaultBase:      *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:       lineNumber,
		Type:             typ,
		Priority:         0,
		ArrayInitializer: arrayInitializer,
	}
	e.SetValue(&e)
	return e
}

var NeNoExpression = NewNoExpression()

func NewNoExpression() NoExpression {
	e := NoExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  UnknownLineNumber,
		Type:        &PtTypeVoid,
		Priority:    0,
	}
	e.SetValue(&e)
	return e
}

func NewNullExpression(typ IType) NullExpression {
	return NewNullExpressionWithAll(UnknownLineNumber, typ)
}

func NewNullExpressionWithAll(lineNumber int, typ IType) NullExpression {
	e := NullExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        typ,
		Priority:    0,
	}
	e.SetValue(&e)
	return e
}

func NewObjectTypeReferenceExpression(typ *ObjectType) ObjectTypeReferenceExpression {
	return NewObjectTypeReferenceExpressionWithAll(UnknownLineNumber, typ, true)
}

func NewObjectTypeReferenceExpressionWithLineNumber(lineNumber int, typ *ObjectType) ObjectTypeReferenceExpression {
	return NewObjectTypeReferenceExpressionWithAll(lineNumber, typ, true)
}

func NewObjectTypeReferenceExpressionWithExplicit(typ *ObjectType, explicit bool) ObjectTypeReferenceExpression {
	return NewObjectTypeReferenceExpressionWithAll(UnknownLineNumber, typ, explicit)
}

func NewObjectTypeReferenceExpressionWithAll(lineNumber int, typ *ObjectType, explicit bool) ObjectTypeReferenceExpression {
	e := ObjectTypeReferenceExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        typ,
		Priority:    0,
		IsExplicit:  explicit,
	}
	e.SetValue(&e)
	return e
}

func NewParenthesesExpression(expression IExpression) ParenthesesExpression {
	e := ParenthesesExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  expression.GetLineNumber(),
		Priority:    0,
		Expression:  expression,
	}
	e.SetValue(&e)
	return e
}

func NewPostOperatorExpression(operator string, expression IExpression) PostOperatorExpression {
	return NewPostOperatorExpressionWithAll(UnknownLineNumber, operator, expression)
}

func NewPostOperatorExpressionWithAll(lineNumber int, operator string, expression IExpression) PostOperatorExpression {
	e := PostOperatorExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Priority:    1,
		Operator:    operator,
		Expression:  expression,
	}
	e.SetValue(&e)
	return e
}

func NewPreOperatorExpression(operator string, expression IExpression) PreOperatorExpression {
	return NewPreOperatorExpressionWithAll(UnknownLineNumber, operator, expression)
}

func NewPreOperatorExpressionWithAll(lineNumber int, operator string, expression IExpression) PreOperatorExpression {
	e := PreOperatorExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Priority:    2,
		Operator:    operator,
		Expression:  expression,
	}
	e.SetValue(&e)
	return e
}

var EmptyString = NewStringConstantExpression("")

func NewStringConstantExpression(text string) StringConstantExpression {
	return NewStringConstantExpressionWithAll(UnknownLineNumber, text)
}

func NewStringConstantExpressionWithAll(lineNumber int, text string) StringConstantExpression {
	e := StringConstantExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Priority:    0,
		Text:        text,
	}
	e.SetValue(&e)
	return e
}

func NewSuperConstructorInvocationExpression(objectType *ObjectType, descriptor string,
	parameters IExpression) SuperConstructorInvocationExpression {
	return NewSuperConstructorInvocationExpressionWithAll(UnknownLineNumber, objectType, descriptor, parameters)
}

func NewSuperConstructorInvocationExpressionWithAll(lineNumber int, objectType *ObjectType,
	descriptor string, parameters IExpression) SuperConstructorInvocationExpression {
	e := SuperConstructorInvocationExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        &PtTypeVoid,
		Priority:    1,
		ObjectType:  objectType,
		Descriptor:  descriptor,
		Parameters:  parameters,
	}
	e.SetValue(&e)
	return e
}

func NewSuperExpression(typ IType) SuperExpression {
	return NewSuperExpressionWithAll(UnknownLineNumber, typ)
}

func NewTernaryOperatorExpression(typ IType, condition IExpression,
	trueExpression IExpression, falseExpression IExpression) TernaryOperatorExpression {
	return NewTernaryOperatorExpressionWithAll(UnknownLineNumber, typ, condition, trueExpression, falseExpression)
}

func NewTernaryOperatorExpressionWithAll(lineNumber int, typ IType, condition IExpression,
	trueExpression IExpression, falseExpression IExpression) TernaryOperatorExpression {
	e := TernaryOperatorExpression{
		DefaultBase:     *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:      lineNumber,
		Type:            typ,
		Priority:        15,
		Condition:       condition,
		TrueExpression:  trueExpression,
		FalseExpression: falseExpression,
	}
	e.SetValue(&e)
	return e
}

func NewThisExpression(typ IType) ThisExpression {
	return NewThisExpressionWithAll(UnknownLineNumber, typ)
}

func NewThisExpressionWithAll(lineNumber int, typ IType) ThisExpression {
	e := ThisExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Priority:    0,
		Type:        typ,
		IsExplicit:  true,
	}
	e.SetValue(&e)
	return e
}

func NewTypeReferenceDotClassExpression(typeDotClass IType) TypeReferenceDotClassExpression {
	return NewTypeReferenceDotClassExpressionWithAll(UnknownLineNumber, typeDotClass)
}

func NewTypeReferenceDotClassExpressionWithAll(lineNumber int, typeDotClass IType) TypeReferenceDotClassExpression {
	e := TypeReferenceDotClassExpression{
		DefaultBase:  *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:   lineNumber,
		Type:         OtTypeClass.CreateTypeWithArgs(typeDotClass.(ITypeArgument)),
		TypeDotClass: typeDotClass,
		Priority:     0,
	}
	e.SetValue(&e)
	return e
}

/////////////////////////////////////////////////////////////////////////
//  Interfaces
/////////////////////////////////////////////////////////////////////////

type IExpression interface {
	Accept(visitor IExpressionVisitor)

	IsArrayExpression() bool
	IsBinaryOperatorExpression() bool
	IsBooleanExpression() bool
	IsCastExpression() bool
	IsConstructorInvocationExpression() bool
	IsDoubleConstantExpression() bool
	IsFieldReferenceExpression() bool
	IsFloatConstantExpression() bool
	IsIntegerConstantExpression() bool
	IsLengthExpression() bool
	IsLocalVariableReferenceExpression() bool
	IsLongConstantExpression() bool
	IsMethodInvocationExpression() bool
	IsNewArray() bool
	IsNewExpression() bool
	IsNewInitializedArray() bool
	IsNullExpression() bool
	IsObjectTypeReferenceExpression() bool
	IsPostOperatorExpression() bool
	IsPreOperatorExpression() bool
	IsStringConstantExpression() bool
	IsSuperConstructorInvocationExpression() bool
	IsSuperExpression() bool
	IsTernaryOperatorExpression() bool
	IsThisExpression() bool

	GetLineNumber() int
	GetType() IType
	GetPriority() int

	GetDimensionExpressionList() IExpression
	GetParameters() IExpression

	GetCondition() IExpression
	GetExpression() IExpression
	GetTrueExpression() IExpression
	GetFalseExpression() IExpression
	GetIndex() IExpression
	GetLeftExpression() IExpression
	GetRightExpression() IExpression

	GetDescriptor() string
	GetDoubleValue() float64
	GetFloatValue() float32
	GetIntegerValue() int
	GetInternalTypeName() string
	GetLongValue() int64
	GetName() string
	GetObjectType() *ObjectType
	GetOperator() string
	GetStringValue() string

	String() string
}

type IExpressionVisitor interface {
	VisitArrayExpression(expression *ArrayExpression)
	VisitBinaryOperatorExpression(expression *BinaryOperatorExpression)
	VisitBooleanExpression(expression *BooleanExpression)
	VisitCastExpression(expression *CastExpression)
	VisitCommentExpression(expression *CommentExpression)
	VisitConstructorInvocationExpression(expression *ConstructorInvocationExpression)
	VisitConstructorReferenceExpression(expression *ConstructorReferenceExpression)
	VisitDoubleConstantExpression(expression *DoubleConstantExpression)
	VisitEnumConstantReferenceExpression(expression *EnumConstantReferenceExpression)
	VisitExpressions(expression *Expressions)
	VisitFieldReferenceExpression(expression *FieldReferenceExpression)
	VisitFloatConstantExpression(expression *FloatConstantExpression)
	VisitIntegerConstantExpression(expression *IntegerConstantExpression)
	VisitInstanceOfExpression(expression *InstanceOfExpression)
	VisitLambdaFormalParametersExpression(expression *LambdaFormalParametersExpression)
	VisitLambdaIdentifiersExpression(expression *LambdaIdentifiersExpression)
	VisitLengthExpression(expression *LengthExpression)
	VisitLocalVariableReferenceExpression(expression *LocalVariableReferenceExpression)
	VisitLongConstantExpression(expression *LongConstantExpression)
	VisitMethodInvocationExpression(expression *MethodInvocationExpression)
	VisitMethodReferenceExpression(expression *MethodReferenceExpression)
	VisitNewArray(expression *NewArray)
	VisitNewExpression(expression *NewExpression)
	VisitNewInitializedArray(expression *NewInitializedArray)
	VisitNoExpression(expression *NoExpression)
	VisitNullExpression(expression *NullExpression)
	VisitObjectTypeReferenceExpression(expression *ObjectTypeReferenceExpression)
	VisitParenthesesExpression(expression *ParenthesesExpression)
	VisitPostOperatorExpression(expression *PostOperatorExpression)
	VisitPreOperatorExpression(expression *PreOperatorExpression)
	VisitStringConstantExpression(expression *StringConstantExpression)
	VisitSuperConstructorInvocationExpression(expression *SuperConstructorInvocationExpression)
	VisitSuperExpression(expression *SuperExpression)
	VisitTernaryOperatorExpression(expression *TernaryOperatorExpression)
	VisitThisExpression(expression *ThisExpression)
	VisitTypeReferenceDotClassExpression(expression *TypeReferenceDotClassExpression)
}

/////////////////////////////////////////////////////////////////////////
//  Structures
/////////////////////////////////////////////////////////////////////////

type ArrayExpression struct {
	util.DefaultBase[IExpression]

	lineNumber int
	typ        IType
	priority   int
	expression IExpression
	index      IExpression
}

func (e *ArrayExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitArrayExpression(e)
}

func (e *ArrayExpression) IsArrayExpression() bool                      { return true }
func (e *ArrayExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *ArrayExpression) IsBooleanExpression() bool                    { return false }
func (e *ArrayExpression) IsCastExpression() bool                       { return false }
func (e *ArrayExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *ArrayExpression) IsDoubleConstantExpression() bool             { return false }
func (e *ArrayExpression) IsFieldReferenceExpression() bool             { return false }
func (e *ArrayExpression) IsFloatConstantExpression() bool              { return false }
func (e *ArrayExpression) IsIntegerConstantExpression() bool            { return false }
func (e *ArrayExpression) IsLengthExpression() bool                     { return false }
func (e *ArrayExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *ArrayExpression) IsLongConstantExpression() bool               { return false }
func (e *ArrayExpression) IsMethodInvocationExpression() bool           { return false }
func (e *ArrayExpression) IsNewArray() bool                             { return false }
func (e *ArrayExpression) IsNewExpression() bool                        { return false }
func (e *ArrayExpression) IsNewInitializedArray() bool                  { return false }
func (e *ArrayExpression) IsNullExpression() bool                       { return false }
func (e *ArrayExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *ArrayExpression) IsPostOperatorExpression() bool               { return false }
func (e *ArrayExpression) IsPreOperatorExpression() bool                { return false }
func (e *ArrayExpression) IsStringConstantExpression() bool             { return false }
func (e *ArrayExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *ArrayExpression) IsSuperExpression() bool                      { return false }
func (e *ArrayExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *ArrayExpression) IsThisExpression() bool                       { return false }

func (e *ArrayExpression) GetLineNumber() int { return e.lineNumber }
func (e *ArrayExpression) GetType() IType     { return e.typ }
func (e *ArrayExpression) GetPriority() int   { return e.priority }

func (e *ArrayExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *ArrayExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *ArrayExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *ArrayExpression) GetExpression() IExpression      { return e.expression }
func (e *ArrayExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *ArrayExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *ArrayExpression) GetIndex() IExpression           { return e.index }
func (e *ArrayExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *ArrayExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *ArrayExpression) GetDescriptor() string       { return "" }
func (e *ArrayExpression) GetDoubleValue() float64     { return 0 }
func (e *ArrayExpression) GetFloatValue() float32      { return 0 }
func (e *ArrayExpression) GetIntegerValue() int        { return 0 }
func (e *ArrayExpression) GetInternalTypeName() string { return "" }
func (e *ArrayExpression) GetLongValue() int64         { return 0 }
func (e *ArrayExpression) GetName() string             { return "" }
func (e *ArrayExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *ArrayExpression) GetOperator() string         { return "" }
func (e *ArrayExpression) GetStringValue() string      { return "" }

func (e *ArrayExpression) String() string {
	return fmt.Sprintf("ArrayExpression{%v[%v]}", e.expression, e.index)
}

type BinaryOperatorExpression struct {
	util.DefaultBase[IExpression]

	LineNumber      int
	Type            IType
	Priority        int
	LeftExpression  IExpression
	Operator        string
	RightExpression IExpression
}

func (e *BinaryOperatorExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitBinaryOperatorExpression(e)
}

func (e *BinaryOperatorExpression) IsArrayExpression() bool                      { return false }
func (e *BinaryOperatorExpression) IsBinaryOperatorExpression() bool             { return true }
func (e *BinaryOperatorExpression) IsBooleanExpression() bool                    { return false }
func (e *BinaryOperatorExpression) IsCastExpression() bool                       { return false }
func (e *BinaryOperatorExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *BinaryOperatorExpression) IsDoubleConstantExpression() bool             { return false }
func (e *BinaryOperatorExpression) IsFieldReferenceExpression() bool             { return false }
func (e *BinaryOperatorExpression) IsFloatConstantExpression() bool              { return false }
func (e *BinaryOperatorExpression) IsIntegerConstantExpression() bool            { return false }
func (e *BinaryOperatorExpression) IsLengthExpression() bool                     { return false }
func (e *BinaryOperatorExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *BinaryOperatorExpression) IsLongConstantExpression() bool               { return false }
func (e *BinaryOperatorExpression) IsMethodInvocationExpression() bool           { return false }
func (e *BinaryOperatorExpression) IsNewArray() bool                             { return false }
func (e *BinaryOperatorExpression) IsNewExpression() bool                        { return false }
func (e *BinaryOperatorExpression) IsNewInitializedArray() bool                  { return false }
func (e *BinaryOperatorExpression) IsNullExpression() bool                       { return false }
func (e *BinaryOperatorExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *BinaryOperatorExpression) IsPostOperatorExpression() bool               { return false }
func (e *BinaryOperatorExpression) IsPreOperatorExpression() bool                { return false }
func (e *BinaryOperatorExpression) IsStringConstantExpression() bool             { return false }
func (e *BinaryOperatorExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *BinaryOperatorExpression) IsSuperExpression() bool                      { return false }
func (e *BinaryOperatorExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *BinaryOperatorExpression) IsThisExpression() bool                       { return false }

func (e *BinaryOperatorExpression) GetLineNumber() int { return e.LineNumber }
func (e *BinaryOperatorExpression) GetType() IType     { return e.Type }
func (e *BinaryOperatorExpression) GetPriority() int   { return e.Priority }

func (e *BinaryOperatorExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *BinaryOperatorExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *BinaryOperatorExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *BinaryOperatorExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *BinaryOperatorExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *BinaryOperatorExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *BinaryOperatorExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *BinaryOperatorExpression) GetLeftExpression() IExpression  { return e.LeftExpression }
func (e *BinaryOperatorExpression) GetRightExpression() IExpression { return e.RightExpression }

func (e *BinaryOperatorExpression) GetDescriptor() string       { return "" }
func (e *BinaryOperatorExpression) GetDoubleValue() float64     { return 0 }
func (e *BinaryOperatorExpression) GetFloatValue() float32      { return 0 }
func (e *BinaryOperatorExpression) GetIntegerValue() int        { return 0 }
func (e *BinaryOperatorExpression) GetInternalTypeName() string { return "" }
func (e *BinaryOperatorExpression) GetLongValue() int64         { return 0 }
func (e *BinaryOperatorExpression) GetName() string             { return "" }
func (e *BinaryOperatorExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *BinaryOperatorExpression) GetOperator() string         { return e.Operator }
func (e *BinaryOperatorExpression) GetStringValue() string      { return "" }

func (e *BinaryOperatorExpression) String() string {
	return fmt.Sprintf("BinaryOperatorExpression{%s %s %s}", e.LeftExpression, e.Operator, e.RightExpression)
}

type BooleanExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	Value      bool
}

func (e *BooleanExpression) IsTrue() bool {
	return e.Value
}

func (e *BooleanExpression) IsFalse() bool {
	return !e.Value
}

func (e *BooleanExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitBooleanExpression(e)
}

func (e *BooleanExpression) IsArrayExpression() bool                      { return false }
func (e *BooleanExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *BooleanExpression) IsBooleanExpression() bool                    { return true }
func (e *BooleanExpression) IsCastExpression() bool                       { return false }
func (e *BooleanExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *BooleanExpression) IsDoubleConstantExpression() bool             { return false }
func (e *BooleanExpression) IsFieldReferenceExpression() bool             { return false }
func (e *BooleanExpression) IsFloatConstantExpression() bool              { return false }
func (e *BooleanExpression) IsIntegerConstantExpression() bool            { return false }
func (e *BooleanExpression) IsLengthExpression() bool                     { return false }
func (e *BooleanExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *BooleanExpression) IsLongConstantExpression() bool               { return false }
func (e *BooleanExpression) IsMethodInvocationExpression() bool           { return false }
func (e *BooleanExpression) IsNewArray() bool                             { return false }
func (e *BooleanExpression) IsNewExpression() bool                        { return false }
func (e *BooleanExpression) IsNewInitializedArray() bool                  { return false }
func (e *BooleanExpression) IsNullExpression() bool                       { return false }
func (e *BooleanExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *BooleanExpression) IsPostOperatorExpression() bool               { return false }
func (e *BooleanExpression) IsPreOperatorExpression() bool                { return false }
func (e *BooleanExpression) IsStringConstantExpression() bool             { return false }
func (e *BooleanExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *BooleanExpression) IsSuperExpression() bool                      { return false }
func (e *BooleanExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *BooleanExpression) IsThisExpression() bool                       { return false }

func (e *BooleanExpression) GetLineNumber() int { return e.LineNumber }
func (e *BooleanExpression) GetType() IType     { return &PtTypeBoolean }
func (e *BooleanExpression) GetPriority() int   { return e.Priority }

func (e *BooleanExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *BooleanExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *BooleanExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *BooleanExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *BooleanExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *BooleanExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *BooleanExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *BooleanExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *BooleanExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *BooleanExpression) GetDescriptor() string       { return "" }
func (e *BooleanExpression) GetDoubleValue() float64     { return 0 }
func (e *BooleanExpression) GetFloatValue() float32      { return 0 }
func (e *BooleanExpression) GetIntegerValue() int        { return 0 }
func (e *BooleanExpression) GetInternalTypeName() string { return "" }
func (e *BooleanExpression) GetLongValue() int64         { return 0 }
func (e *BooleanExpression) GetName() string             { return "" }
func (e *BooleanExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *BooleanExpression) GetOperator() string         { return "" }
func (e *BooleanExpression) GetStringValue() string      { return "" }

func (e *BooleanExpression) String() string {
	value := "false"
	if e.Value {
		value = "true"
	}
	return fmt.Sprintf("BooleanExpression{%s}", value)
}

type CastExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	Expression IExpression
	IsExplicit bool
}

func (e *CastExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitCastExpression(e)
}

func (e *CastExpression) IsArrayExpression() bool                      { return false }
func (e *CastExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *CastExpression) IsBooleanExpression() bool                    { return false }
func (e *CastExpression) IsCastExpression() bool                       { return true }
func (e *CastExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *CastExpression) IsDoubleConstantExpression() bool             { return false }
func (e *CastExpression) IsFieldReferenceExpression() bool             { return false }
func (e *CastExpression) IsFloatConstantExpression() bool              { return false }
func (e *CastExpression) IsIntegerConstantExpression() bool            { return false }
func (e *CastExpression) IsLengthExpression() bool                     { return false }
func (e *CastExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *CastExpression) IsLongConstantExpression() bool               { return false }
func (e *CastExpression) IsMethodInvocationExpression() bool           { return false }
func (e *CastExpression) IsNewArray() bool                             { return false }
func (e *CastExpression) IsNewExpression() bool                        { return false }
func (e *CastExpression) IsNewInitializedArray() bool                  { return false }
func (e *CastExpression) IsNullExpression() bool                       { return false }
func (e *CastExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *CastExpression) IsPostOperatorExpression() bool               { return false }
func (e *CastExpression) IsPreOperatorExpression() bool                { return false }
func (e *CastExpression) IsStringConstantExpression() bool             { return false }
func (e *CastExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *CastExpression) IsSuperExpression() bool                      { return false }
func (e *CastExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *CastExpression) IsThisExpression() bool                       { return false }

func (e *CastExpression) GetLineNumber() int { return e.LineNumber }
func (e *CastExpression) GetType() IType     { return e.Type }
func (e *CastExpression) GetPriority() int   { return e.Priority }

func (e *CastExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *CastExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *CastExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *CastExpression) GetExpression() IExpression      { return e.Expression }
func (e *CastExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *CastExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *CastExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *CastExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *CastExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *CastExpression) GetDescriptor() string       { return "" }
func (e *CastExpression) GetDoubleValue() float64     { return 0 }
func (e *CastExpression) GetFloatValue() float32      { return 0 }
func (e *CastExpression) GetIntegerValue() int        { return 0 }
func (e *CastExpression) GetInternalTypeName() string { return "" }
func (e *CastExpression) GetLongValue() int64         { return 0 }
func (e *CastExpression) GetName() string             { return "" }
func (e *CastExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *CastExpression) GetOperator() string         { return "" }
func (e *CastExpression) GetStringValue() string      { return "" }

func (e *CastExpression) String() string {
	return fmt.Sprintf("CastExpression{cast (%s) %s }", e.Type, e.Expression)
}

type CommentExpression struct {
	util.DefaultBase[IExpression]

	Text string
}

func (e *CommentExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitCommentExpression(e)
}

func (e *CommentExpression) IsArrayExpression() bool                      { return false }
func (e *CommentExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *CommentExpression) IsBooleanExpression() bool                    { return false }
func (e *CommentExpression) IsCastExpression() bool                       { return false }
func (e *CommentExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *CommentExpression) IsDoubleConstantExpression() bool             { return false }
func (e *CommentExpression) IsFieldReferenceExpression() bool             { return false }
func (e *CommentExpression) IsFloatConstantExpression() bool              { return false }
func (e *CommentExpression) IsIntegerConstantExpression() bool            { return false }
func (e *CommentExpression) IsLengthExpression() bool                     { return false }
func (e *CommentExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *CommentExpression) IsLongConstantExpression() bool               { return false }
func (e *CommentExpression) IsMethodInvocationExpression() bool           { return false }
func (e *CommentExpression) IsNewArray() bool                             { return false }
func (e *CommentExpression) IsNewExpression() bool                        { return false }
func (e *CommentExpression) IsNewInitializedArray() bool                  { return false }
func (e *CommentExpression) IsNullExpression() bool                       { return false }
func (e *CommentExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *CommentExpression) IsPostOperatorExpression() bool               { return false }
func (e *CommentExpression) IsPreOperatorExpression() bool                { return false }
func (e *CommentExpression) IsStringConstantExpression() bool             { return false }
func (e *CommentExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *CommentExpression) IsSuperExpression() bool                      { return false }
func (e *CommentExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *CommentExpression) IsThisExpression() bool                       { return false }

func (e *CommentExpression) GetLineNumber() int { return UnknownLineNumber }
func (e *CommentExpression) GetType() IType     { return &PtTypeVoid }
func (e *CommentExpression) GetPriority() int   { return 0 }

func (e *CommentExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *CommentExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *CommentExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *CommentExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *CommentExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *CommentExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *CommentExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *CommentExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *CommentExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *CommentExpression) GetDescriptor() string       { return "" }
func (e *CommentExpression) GetDoubleValue() float64     { return 0 }
func (e *CommentExpression) GetFloatValue() float32      { return 0 }
func (e *CommentExpression) GetIntegerValue() int        { return 0 }
func (e *CommentExpression) GetInternalTypeName() string { return "" }
func (e *CommentExpression) GetLongValue() int64         { return 0 }
func (e *CommentExpression) GetName() string             { return "" }
func (e *CommentExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *CommentExpression) GetOperator() string         { return "" }
func (e *CommentExpression) GetStringValue() string      { return "" }

func (e *CommentExpression) String() string {
	return fmt.Sprintf("CommentExpression{%s}", e.Text)
}

type ConstructorInvocationExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	ObjectType *ObjectType
	Descriptor string
	Parameters IExpression
}

func (e *ConstructorInvocationExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitConstructorInvocationExpression(e)
}

func (e *ConstructorInvocationExpression) IsArrayExpression() bool                      { return false }
func (e *ConstructorInvocationExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *ConstructorInvocationExpression) IsBooleanExpression() bool                    { return false }
func (e *ConstructorInvocationExpression) IsCastExpression() bool                       { return false }
func (e *ConstructorInvocationExpression) IsConstructorInvocationExpression() bool      { return true }
func (e *ConstructorInvocationExpression) IsDoubleConstantExpression() bool             { return false }
func (e *ConstructorInvocationExpression) IsFieldReferenceExpression() bool             { return false }
func (e *ConstructorInvocationExpression) IsFloatConstantExpression() bool              { return false }
func (e *ConstructorInvocationExpression) IsIntegerConstantExpression() bool            { return false }
func (e *ConstructorInvocationExpression) IsLengthExpression() bool                     { return false }
func (e *ConstructorInvocationExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *ConstructorInvocationExpression) IsLongConstantExpression() bool               { return false }
func (e *ConstructorInvocationExpression) IsMethodInvocationExpression() bool           { return false }
func (e *ConstructorInvocationExpression) IsNewArray() bool                             { return false }
func (e *ConstructorInvocationExpression) IsNewExpression() bool                        { return false }
func (e *ConstructorInvocationExpression) IsNewInitializedArray() bool                  { return false }
func (e *ConstructorInvocationExpression) IsNullExpression() bool                       { return false }
func (e *ConstructorInvocationExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *ConstructorInvocationExpression) IsPostOperatorExpression() bool               { return false }
func (e *ConstructorInvocationExpression) IsPreOperatorExpression() bool                { return false }
func (e *ConstructorInvocationExpression) IsStringConstantExpression() bool             { return false }
func (e *ConstructorInvocationExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *ConstructorInvocationExpression) IsSuperExpression() bool                      { return false }
func (e *ConstructorInvocationExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *ConstructorInvocationExpression) IsThisExpression() bool                       { return false }

func (e *ConstructorInvocationExpression) GetLineNumber() int { return e.LineNumber }
func (e *ConstructorInvocationExpression) GetType() IType     { return e.Type }
func (e *ConstructorInvocationExpression) GetPriority() int   { return e.Priority }

func (e *ConstructorInvocationExpression) GetDimensionExpressionList() IExpression {
	return &NeNoExpression
}
func (e *ConstructorInvocationExpression) GetParameters() IExpression { return e.Parameters }

func (e *ConstructorInvocationExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *ConstructorInvocationExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *ConstructorInvocationExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *ConstructorInvocationExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *ConstructorInvocationExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *ConstructorInvocationExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *ConstructorInvocationExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *ConstructorInvocationExpression) GetDescriptor() string       { return "" }
func (e *ConstructorInvocationExpression) GetDoubleValue() float64     { return 0 }
func (e *ConstructorInvocationExpression) GetFloatValue() float32      { return 0 }
func (e *ConstructorInvocationExpression) GetIntegerValue() int        { return 0 }
func (e *ConstructorInvocationExpression) GetInternalTypeName() string { return "" }
func (e *ConstructorInvocationExpression) GetLongValue() int64         { return 0 }
func (e *ConstructorInvocationExpression) GetName() string             { return "" }
func (e *ConstructorInvocationExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *ConstructorInvocationExpression) GetOperator() string         { return "" }
func (e *ConstructorInvocationExpression) GetStringValue() string      { return "" }

func (e *ConstructorInvocationExpression) String() string {
	return fmt.Sprintf("ConstructorInvocationExpression{call this(%s)}", e.Descriptor)
}

type ConstructorReferenceExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	ObjectType *ObjectType
	Descriptor string
}

func (e *ConstructorReferenceExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitConstructorReferenceExpression(e)
}

func (e *ConstructorReferenceExpression) IsArrayExpression() bool                      { return false }
func (e *ConstructorReferenceExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *ConstructorReferenceExpression) IsBooleanExpression() bool                    { return false }
func (e *ConstructorReferenceExpression) IsCastExpression() bool                       { return false }
func (e *ConstructorReferenceExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *ConstructorReferenceExpression) IsDoubleConstantExpression() bool             { return false }
func (e *ConstructorReferenceExpression) IsFieldReferenceExpression() bool             { return false }
func (e *ConstructorReferenceExpression) IsFloatConstantExpression() bool              { return false }
func (e *ConstructorReferenceExpression) IsIntegerConstantExpression() bool            { return false }
func (e *ConstructorReferenceExpression) IsLengthExpression() bool                     { return false }
func (e *ConstructorReferenceExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *ConstructorReferenceExpression) IsLongConstantExpression() bool               { return false }
func (e *ConstructorReferenceExpression) IsMethodInvocationExpression() bool           { return false }
func (e *ConstructorReferenceExpression) IsNewArray() bool                             { return false }
func (e *ConstructorReferenceExpression) IsNewExpression() bool                        { return false }
func (e *ConstructorReferenceExpression) IsNewInitializedArray() bool                  { return false }
func (e *ConstructorReferenceExpression) IsNullExpression() bool                       { return false }
func (e *ConstructorReferenceExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *ConstructorReferenceExpression) IsPostOperatorExpression() bool               { return false }
func (e *ConstructorReferenceExpression) IsPreOperatorExpression() bool                { return false }
func (e *ConstructorReferenceExpression) IsStringConstantExpression() bool             { return false }
func (e *ConstructorReferenceExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *ConstructorReferenceExpression) IsSuperExpression() bool                      { return false }
func (e *ConstructorReferenceExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *ConstructorReferenceExpression) IsThisExpression() bool                       { return false }

func (e *ConstructorReferenceExpression) GetLineNumber() int { return e.LineNumber }
func (e *ConstructorReferenceExpression) GetType() IType     { return e.Type }
func (e *ConstructorReferenceExpression) GetPriority() int   { return e.Priority }

func (e *ConstructorReferenceExpression) GetDimensionExpressionList() IExpression {
	return &NeNoExpression
}
func (e *ConstructorReferenceExpression) GetParameters() IExpression { return &NeNoExpression }

func (e *ConstructorReferenceExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *ConstructorReferenceExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *ConstructorReferenceExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *ConstructorReferenceExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *ConstructorReferenceExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *ConstructorReferenceExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *ConstructorReferenceExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *ConstructorReferenceExpression) GetDescriptor() string       { return e.Descriptor }
func (e *ConstructorReferenceExpression) GetDoubleValue() float64     { return 0 }
func (e *ConstructorReferenceExpression) GetFloatValue() float32      { return 0 }
func (e *ConstructorReferenceExpression) GetIntegerValue() int        { return 0 }
func (e *ConstructorReferenceExpression) GetInternalTypeName() string { return "" }
func (e *ConstructorReferenceExpression) GetLongValue() int64         { return 0 }
func (e *ConstructorReferenceExpression) GetName() string             { return "" }
func (e *ConstructorReferenceExpression) GetObjectType() *ObjectType  { return e.ObjectType }
func (e *ConstructorReferenceExpression) GetOperator() string         { return "" }
func (e *ConstructorReferenceExpression) GetStringValue() string      { return "" }

func (e *ConstructorReferenceExpression) String() string { return "ConstructorReferenceExpression{}" }

type DoubleConstantExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	Value      float64
}

func (e *DoubleConstantExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitDoubleConstantExpression(e)
}

func (e *DoubleConstantExpression) IsArrayExpression() bool                      { return false }
func (e *DoubleConstantExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *DoubleConstantExpression) IsBooleanExpression() bool                    { return false }
func (e *DoubleConstantExpression) IsCastExpression() bool                       { return false }
func (e *DoubleConstantExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *DoubleConstantExpression) IsDoubleConstantExpression() bool             { return true }
func (e *DoubleConstantExpression) IsFieldReferenceExpression() bool             { return false }
func (e *DoubleConstantExpression) IsFloatConstantExpression() bool              { return false }
func (e *DoubleConstantExpression) IsIntegerConstantExpression() bool            { return false }
func (e *DoubleConstantExpression) IsLengthExpression() bool                     { return false }
func (e *DoubleConstantExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *DoubleConstantExpression) IsLongConstantExpression() bool               { return false }
func (e *DoubleConstantExpression) IsMethodInvocationExpression() bool           { return false }
func (e *DoubleConstantExpression) IsNewArray() bool                             { return false }
func (e *DoubleConstantExpression) IsNewExpression() bool                        { return false }
func (e *DoubleConstantExpression) IsNewInitializedArray() bool                  { return false }
func (e *DoubleConstantExpression) IsNullExpression() bool                       { return false }
func (e *DoubleConstantExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *DoubleConstantExpression) IsPostOperatorExpression() bool               { return false }
func (e *DoubleConstantExpression) IsPreOperatorExpression() bool                { return false }
func (e *DoubleConstantExpression) IsStringConstantExpression() bool             { return false }
func (e *DoubleConstantExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *DoubleConstantExpression) IsSuperExpression() bool                      { return false }
func (e *DoubleConstantExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *DoubleConstantExpression) IsThisExpression() bool                       { return false }

func (e *DoubleConstantExpression) GetLineNumber() int { return e.LineNumber }
func (e *DoubleConstantExpression) GetType() IType     { return e.Type }
func (e *DoubleConstantExpression) GetPriority() int   { return e.Priority }

func (e *DoubleConstantExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *DoubleConstantExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *DoubleConstantExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *DoubleConstantExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *DoubleConstantExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *DoubleConstantExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *DoubleConstantExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *DoubleConstantExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *DoubleConstantExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *DoubleConstantExpression) GetDescriptor() string       { return "" }
func (e *DoubleConstantExpression) GetDoubleValue() float64     { return e.Value }
func (e *DoubleConstantExpression) GetFloatValue() float32      { return 0 }
func (e *DoubleConstantExpression) GetIntegerValue() int        { return 0 }
func (e *DoubleConstantExpression) GetInternalTypeName() string { return "" }
func (e *DoubleConstantExpression) GetLongValue() int64         { return 0 }
func (e *DoubleConstantExpression) GetName() string             { return "" }
func (e *DoubleConstantExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *DoubleConstantExpression) GetOperator() string         { return "" }
func (e *DoubleConstantExpression) GetStringValue() string      { return "" }

func (e *DoubleConstantExpression) String() string {
	return fmt.Sprintf("DoubleConstantExpression{%f}", e.Value)
}

type EnumConstantReferenceExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	Name       string
}

func (e *EnumConstantReferenceExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitEnumConstantReferenceExpression(e)
}

func (e *EnumConstantReferenceExpression) IsArrayExpression() bool                      { return false }
func (e *EnumConstantReferenceExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *EnumConstantReferenceExpression) IsBooleanExpression() bool                    { return false }
func (e *EnumConstantReferenceExpression) IsCastExpression() bool                       { return false }
func (e *EnumConstantReferenceExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *EnumConstantReferenceExpression) IsDoubleConstantExpression() bool             { return false }
func (e *EnumConstantReferenceExpression) IsFieldReferenceExpression() bool             { return false }
func (e *EnumConstantReferenceExpression) IsFloatConstantExpression() bool              { return false }
func (e *EnumConstantReferenceExpression) IsIntegerConstantExpression() bool            { return false }
func (e *EnumConstantReferenceExpression) IsLengthExpression() bool                     { return false }
func (e *EnumConstantReferenceExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *EnumConstantReferenceExpression) IsLongConstantExpression() bool               { return false }
func (e *EnumConstantReferenceExpression) IsMethodInvocationExpression() bool           { return false }
func (e *EnumConstantReferenceExpression) IsNewArray() bool                             { return false }
func (e *EnumConstantReferenceExpression) IsNewExpression() bool                        { return false }
func (e *EnumConstantReferenceExpression) IsNewInitializedArray() bool                  { return false }
func (e *EnumConstantReferenceExpression) IsNullExpression() bool                       { return false }
func (e *EnumConstantReferenceExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *EnumConstantReferenceExpression) IsPostOperatorExpression() bool               { return false }
func (e *EnumConstantReferenceExpression) IsPreOperatorExpression() bool                { return false }
func (e *EnumConstantReferenceExpression) IsStringConstantExpression() bool             { return false }
func (e *EnumConstantReferenceExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *EnumConstantReferenceExpression) IsSuperExpression() bool                      { return false }
func (e *EnumConstantReferenceExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *EnumConstantReferenceExpression) IsThisExpression() bool                       { return false }

func (e *EnumConstantReferenceExpression) GetLineNumber() int { return e.LineNumber }
func (e *EnumConstantReferenceExpression) GetType() IType     { return e.Type }
func (e *EnumConstantReferenceExpression) GetPriority() int   { return e.Priority }

func (e *EnumConstantReferenceExpression) GetDimensionExpressionList() IExpression {
	return &NeNoExpression
}
func (e *EnumConstantReferenceExpression) GetParameters() IExpression { return &NeNoExpression }

func (e *EnumConstantReferenceExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *EnumConstantReferenceExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *EnumConstantReferenceExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *EnumConstantReferenceExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *EnumConstantReferenceExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *EnumConstantReferenceExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *EnumConstantReferenceExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *EnumConstantReferenceExpression) GetDescriptor() string       { return "" }
func (e *EnumConstantReferenceExpression) GetDoubleValue() float64     { return 0 }
func (e *EnumConstantReferenceExpression) GetFloatValue() float32      { return 0 }
func (e *EnumConstantReferenceExpression) GetIntegerValue() int        { return 0 }
func (e *EnumConstantReferenceExpression) GetInternalTypeName() string { return "" }
func (e *EnumConstantReferenceExpression) GetLongValue() int64         { return 0 }
func (e *EnumConstantReferenceExpression) GetName() string             { return e.Name }
func (e *EnumConstantReferenceExpression) GetObjectType() *ObjectType  { return e.Type.(*ObjectType) }
func (e *EnumConstantReferenceExpression) GetOperator() string         { return "" }
func (e *EnumConstantReferenceExpression) GetStringValue() string      { return "" }

func (e *EnumConstantReferenceExpression) String() string {
	return fmt.Sprintf("EnumConstantReferenceExpression{type=%s, Name=%s}", e.Type.String(), e.Name)
}

type Expressions struct {
	util.DefaultList[IExpression]
}

func (e *Expressions) Accept(visitor IExpressionVisitor) {
	visitor.VisitExpressions(e)
}

func (e *Expressions) IsArrayExpression() bool                      { return false }
func (e *Expressions) IsBinaryOperatorExpression() bool             { return false }
func (e *Expressions) IsBooleanExpression() bool                    { return false }
func (e *Expressions) IsCastExpression() bool                       { return false }
func (e *Expressions) IsConstructorInvocationExpression() bool      { return false }
func (e *Expressions) IsDoubleConstantExpression() bool             { return false }
func (e *Expressions) IsFieldReferenceExpression() bool             { return false }
func (e *Expressions) IsFloatConstantExpression() bool              { return false }
func (e *Expressions) IsIntegerConstantExpression() bool            { return false }
func (e *Expressions) IsLengthExpression() bool                     { return false }
func (e *Expressions) IsLocalVariableReferenceExpression() bool     { return false }
func (e *Expressions) IsLongConstantExpression() bool               { return false }
func (e *Expressions) IsMethodInvocationExpression() bool           { return false }
func (e *Expressions) IsNewArray() bool                             { return false }
func (e *Expressions) IsNewExpression() bool                        { return false }
func (e *Expressions) IsNewInitializedArray() bool                  { return false }
func (e *Expressions) IsNullExpression() bool                       { return false }
func (e *Expressions) IsObjectTypeReferenceExpression() bool        { return false }
func (e *Expressions) IsPostOperatorExpression() bool               { return false }
func (e *Expressions) IsPreOperatorExpression() bool                { return false }
func (e *Expressions) IsStringConstantExpression() bool             { return false }
func (e *Expressions) IsSuperConstructorInvocationExpression() bool { return false }
func (e *Expressions) IsSuperExpression() bool                      { return false }
func (e *Expressions) IsTernaryOperatorExpression() bool            { return false }
func (e *Expressions) IsThisExpression() bool                       { return false }

func (e *Expressions) GetLineNumber() int { return UnknownLineNumber }
func (e *Expressions) GetType() IType     { return nil }
func (e *Expressions) GetPriority() int   { return -1 }

func (e *Expressions) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *Expressions) GetParameters() IExpression              { return &NeNoExpression }

func (e *Expressions) GetCondition() IExpression       { return &NeNoExpression }
func (e *Expressions) GetExpression() IExpression      { return &NeNoExpression }
func (e *Expressions) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *Expressions) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *Expressions) GetIndex() IExpression           { return &NeNoExpression }
func (e *Expressions) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *Expressions) GetRightExpression() IExpression { return &NeNoExpression }

func (e *Expressions) GetDescriptor() string       { return "" }
func (e *Expressions) GetDoubleValue() float64     { return 0 }
func (e *Expressions) GetFloatValue() float32      { return 0 }
func (e *Expressions) GetIntegerValue() int        { return 0 }
func (e *Expressions) GetInternalTypeName() string { return "" }
func (e *Expressions) GetLongValue() int64         { return 0 }
func (e *Expressions) GetName() string             { return "" }
func (e *Expressions) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *Expressions) GetOperator() string         { return "" }
func (e *Expressions) GetStringValue() string      { return "" }

func (e *Expressions) String() string { return "" }

type FieldReferenceExpression struct {
	util.DefaultBase[IExpression]

	LineNumber       int
	Type             IType
	Priority         int
	Expression       IExpression
	InternalTypeName string
	Name             string
	Descriptor       string
}

func (e *FieldReferenceExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitFieldReferenceExpression(e)
}

func (e *FieldReferenceExpression) IsArrayExpression() bool                      { return false }
func (e *FieldReferenceExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *FieldReferenceExpression) IsBooleanExpression() bool                    { return false }
func (e *FieldReferenceExpression) IsCastExpression() bool                       { return false }
func (e *FieldReferenceExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *FieldReferenceExpression) IsDoubleConstantExpression() bool             { return false }
func (e *FieldReferenceExpression) IsFieldReferenceExpression() bool             { return true }
func (e *FieldReferenceExpression) IsFloatConstantExpression() bool              { return false }
func (e *FieldReferenceExpression) IsIntegerConstantExpression() bool            { return false }
func (e *FieldReferenceExpression) IsLengthExpression() bool                     { return false }
func (e *FieldReferenceExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *FieldReferenceExpression) IsLongConstantExpression() bool               { return false }
func (e *FieldReferenceExpression) IsMethodInvocationExpression() bool           { return false }
func (e *FieldReferenceExpression) IsNewArray() bool                             { return false }
func (e *FieldReferenceExpression) IsNewExpression() bool                        { return false }
func (e *FieldReferenceExpression) IsNewInitializedArray() bool                  { return false }
func (e *FieldReferenceExpression) IsNullExpression() bool                       { return false }
func (e *FieldReferenceExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *FieldReferenceExpression) IsPostOperatorExpression() bool               { return false }
func (e *FieldReferenceExpression) IsPreOperatorExpression() bool                { return false }
func (e *FieldReferenceExpression) IsStringConstantExpression() bool             { return false }
func (e *FieldReferenceExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *FieldReferenceExpression) IsSuperExpression() bool                      { return false }
func (e *FieldReferenceExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *FieldReferenceExpression) IsThisExpression() bool                       { return false }

func (e *FieldReferenceExpression) GetLineNumber() int { return e.LineNumber }
func (e *FieldReferenceExpression) GetType() IType     { return e.Type }
func (e *FieldReferenceExpression) GetPriority() int   { return e.Priority }

func (e *FieldReferenceExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *FieldReferenceExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *FieldReferenceExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *FieldReferenceExpression) GetExpression() IExpression      { return e.Expression }
func (e *FieldReferenceExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *FieldReferenceExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *FieldReferenceExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *FieldReferenceExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *FieldReferenceExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *FieldReferenceExpression) GetDescriptor() string       { return e.Descriptor }
func (e *FieldReferenceExpression) GetDoubleValue() float64     { return 0 }
func (e *FieldReferenceExpression) GetFloatValue() float32      { return 0 }
func (e *FieldReferenceExpression) GetIntegerValue() int        { return 0 }
func (e *FieldReferenceExpression) GetInternalTypeName() string { return e.InternalTypeName }
func (e *FieldReferenceExpression) GetLongValue() int64         { return 0 }
func (e *FieldReferenceExpression) GetName() string             { return e.Name }
func (e *FieldReferenceExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *FieldReferenceExpression) GetOperator() string         { return "" }
func (e *FieldReferenceExpression) GetStringValue() string      { return "" }

func (e *FieldReferenceExpression) String() string {
	return fmt.Sprintf("FieldReferenceExpression{type=%s, Expression=%s, Name=%s, Descriptor=%s }", e.Type, e.Expression, e.Name, e.Descriptor)
}

type FloatConstantExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	value      float32
}

func (e *FloatConstantExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitFloatConstantExpression(e)
}

func (e *FloatConstantExpression) IsArrayExpression() bool                      { return false }
func (e *FloatConstantExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *FloatConstantExpression) IsBooleanExpression() bool                    { return false }
func (e *FloatConstantExpression) IsCastExpression() bool                       { return false }
func (e *FloatConstantExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *FloatConstantExpression) IsDoubleConstantExpression() bool             { return false }
func (e *FloatConstantExpression) IsFieldReferenceExpression() bool             { return false }
func (e *FloatConstantExpression) IsFloatConstantExpression() bool              { return true }
func (e *FloatConstantExpression) IsIntegerConstantExpression() bool            { return false }
func (e *FloatConstantExpression) IsLengthExpression() bool                     { return false }
func (e *FloatConstantExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *FloatConstantExpression) IsLongConstantExpression() bool               { return false }
func (e *FloatConstantExpression) IsMethodInvocationExpression() bool           { return false }
func (e *FloatConstantExpression) IsNewArray() bool                             { return false }
func (e *FloatConstantExpression) IsNewExpression() bool                        { return false }
func (e *FloatConstantExpression) IsNewInitializedArray() bool                  { return false }
func (e *FloatConstantExpression) IsNullExpression() bool                       { return false }
func (e *FloatConstantExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *FloatConstantExpression) IsPostOperatorExpression() bool               { return false }
func (e *FloatConstantExpression) IsPreOperatorExpression() bool                { return false }
func (e *FloatConstantExpression) IsStringConstantExpression() bool             { return false }
func (e *FloatConstantExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *FloatConstantExpression) IsSuperExpression() bool                      { return false }
func (e *FloatConstantExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *FloatConstantExpression) IsThisExpression() bool                       { return false }

func (e *FloatConstantExpression) GetLineNumber() int { return e.LineNumber }
func (e *FloatConstantExpression) GetType() IType     { return e.Type }
func (e *FloatConstantExpression) GetPriority() int   { return e.Priority }

func (e *FloatConstantExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *FloatConstantExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *FloatConstantExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *FloatConstantExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *FloatConstantExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *FloatConstantExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *FloatConstantExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *FloatConstantExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *FloatConstantExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *FloatConstantExpression) GetDescriptor() string       { return "" }
func (e *FloatConstantExpression) GetDoubleValue() float64     { return 0 }
func (e *FloatConstantExpression) GetFloatValue() float32      { return e.value }
func (e *FloatConstantExpression) GetIntegerValue() int        { return 0 }
func (e *FloatConstantExpression) GetInternalTypeName() string { return "" }
func (e *FloatConstantExpression) GetLongValue() int64         { return 0 }
func (e *FloatConstantExpression) GetName() string             { return "" }
func (e *FloatConstantExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *FloatConstantExpression) GetOperator() string         { return "" }
func (e *FloatConstantExpression) GetStringValue() string      { return "" }

func (e *FloatConstantExpression) String() string {
	return fmt.Sprintf("FloatConstantExpression{ %.2f }", e.value)
}

type InstanceOfExpression struct {
	util.DefaultBase[IExpression]

	LineNumber     int
	Type           IType
	Priority       int
	expression     IExpression
	instanceOfType IType
}

func (e *InstanceOfExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitInstanceOfExpression(e)
}

func (e *InstanceOfExpression) IsArrayExpression() bool                      { return false }
func (e *InstanceOfExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *InstanceOfExpression) IsBooleanExpression() bool                    { return false }
func (e *InstanceOfExpression) IsCastExpression() bool                       { return false }
func (e *InstanceOfExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *InstanceOfExpression) IsDoubleConstantExpression() bool             { return false }
func (e *InstanceOfExpression) IsFieldReferenceExpression() bool             { return false }
func (e *InstanceOfExpression) IsFloatConstantExpression() bool              { return false }
func (e *InstanceOfExpression) IsIntegerConstantExpression() bool            { return false }
func (e *InstanceOfExpression) IsLengthExpression() bool                     { return false }
func (e *InstanceOfExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *InstanceOfExpression) IsLongConstantExpression() bool               { return false }
func (e *InstanceOfExpression) IsMethodInvocationExpression() bool           { return false }
func (e *InstanceOfExpression) IsNewArray() bool                             { return false }
func (e *InstanceOfExpression) IsNewExpression() bool                        { return false }
func (e *InstanceOfExpression) IsNewInitializedArray() bool                  { return false }
func (e *InstanceOfExpression) IsNullExpression() bool                       { return false }
func (e *InstanceOfExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *InstanceOfExpression) IsPostOperatorExpression() bool               { return false }
func (e *InstanceOfExpression) IsPreOperatorExpression() bool                { return false }
func (e *InstanceOfExpression) IsStringConstantExpression() bool             { return false }
func (e *InstanceOfExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *InstanceOfExpression) IsSuperExpression() bool                      { return false }
func (e *InstanceOfExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *InstanceOfExpression) IsThisExpression() bool                       { return false }

func (e *InstanceOfExpression) GetLineNumber() int { return e.LineNumber }
func (e *InstanceOfExpression) GetType() IType     { return e.Type }
func (e *InstanceOfExpression) GetPriority() int   { return e.Priority }

func (e *InstanceOfExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *InstanceOfExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *InstanceOfExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *InstanceOfExpression) GetExpression() IExpression      { return e.expression }
func (e *InstanceOfExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *InstanceOfExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *InstanceOfExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *InstanceOfExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *InstanceOfExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *InstanceOfExpression) GetDescriptor() string       { return "" }
func (e *InstanceOfExpression) GetDoubleValue() float64     { return 0 }
func (e *InstanceOfExpression) GetFloatValue() float32      { return 0 }
func (e *InstanceOfExpression) GetIntegerValue() int        { return 0 }
func (e *InstanceOfExpression) GetInternalTypeName() string { return "" }
func (e *InstanceOfExpression) GetLongValue() int64         { return 0 }
func (e *InstanceOfExpression) GetName() string             { return "" }
func (e *InstanceOfExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *InstanceOfExpression) GetOperator() string         { return "" }
func (e *InstanceOfExpression) GetStringValue() string      { return "" }

func (e *InstanceOfExpression) String() string {
	return fmt.Sprintf("InstanceOfExpression{ line-number=%d, type=%s, priority=%d }", e.LineNumber, e.Type.String(), e.Priority)
}

type IntegerConstantExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	Value      int
}

func (e *IntegerConstantExpression) checkType(typ IType) bool {
	if typ.IsPrimitiveType() {
		valueType := GetPrimitiveTypeFromValue(e.Value)
		pt, ok := e.Type.(*PrimitiveType)
		if ok {
			return pt.Flags&valueType.Flags != 0
		}
	}
	return false
}

func (e *IntegerConstantExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitIntegerConstantExpression(e)
}

func (e *IntegerConstantExpression) IsArrayExpression() bool                      { return false }
func (e *IntegerConstantExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *IntegerConstantExpression) IsBooleanExpression() bool                    { return false }
func (e *IntegerConstantExpression) IsCastExpression() bool                       { return false }
func (e *IntegerConstantExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *IntegerConstantExpression) IsDoubleConstantExpression() bool             { return false }
func (e *IntegerConstantExpression) IsFieldReferenceExpression() bool             { return false }
func (e *IntegerConstantExpression) IsFloatConstantExpression() bool              { return false }
func (e *IntegerConstantExpression) IsIntegerConstantExpression() bool            { return true }
func (e *IntegerConstantExpression) IsLengthExpression() bool                     { return false }
func (e *IntegerConstantExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *IntegerConstantExpression) IsLongConstantExpression() bool               { return false }
func (e *IntegerConstantExpression) IsMethodInvocationExpression() bool           { return false }
func (e *IntegerConstantExpression) IsNewArray() bool                             { return false }
func (e *IntegerConstantExpression) IsNewExpression() bool                        { return false }
func (e *IntegerConstantExpression) IsNewInitializedArray() bool                  { return false }
func (e *IntegerConstantExpression) IsNullExpression() bool                       { return false }
func (e *IntegerConstantExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *IntegerConstantExpression) IsPostOperatorExpression() bool               { return false }
func (e *IntegerConstantExpression) IsPreOperatorExpression() bool                { return false }
func (e *IntegerConstantExpression) IsStringConstantExpression() bool             { return false }
func (e *IntegerConstantExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *IntegerConstantExpression) IsSuperExpression() bool                      { return false }
func (e *IntegerConstantExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *IntegerConstantExpression) IsThisExpression() bool                       { return false }

func (e *IntegerConstantExpression) GetLineNumber() int { return e.LineNumber }
func (e *IntegerConstantExpression) GetType() IType     { return e.Type }
func (e *IntegerConstantExpression) GetPriority() int   { return e.Priority }

func (e *IntegerConstantExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *IntegerConstantExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *IntegerConstantExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *IntegerConstantExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *IntegerConstantExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *IntegerConstantExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *IntegerConstantExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *IntegerConstantExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *IntegerConstantExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *IntegerConstantExpression) GetDescriptor() string       { return "" }
func (e *IntegerConstantExpression) GetDoubleValue() float64     { return 0 }
func (e *IntegerConstantExpression) GetFloatValue() float32      { return 0 }
func (e *IntegerConstantExpression) GetIntegerValue() int        { return e.Value }
func (e *IntegerConstantExpression) GetInternalTypeName() string { return "" }
func (e *IntegerConstantExpression) GetLongValue() int64         { return 0 }
func (e *IntegerConstantExpression) GetName() string             { return "" }
func (e *IntegerConstantExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *IntegerConstantExpression) GetOperator() string         { return "" }
func (e *IntegerConstantExpression) GetStringValue() string      { return "" }

func (e *IntegerConstantExpression) String() string {
	return fmt.Sprintf("IntegerConstantExpression{type=%s, Value=%d}", e.Type, e.Value)
}

type LambdaFormalParametersExpression struct {
	util.DefaultBase[IExpression]

	LineNumber       int
	Type             IType
	Priority         int
	Statements       IStatement
	FormalParameters *FormalParameter
}

func (e *LambdaFormalParametersExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitLambdaFormalParametersExpression(e)
}

func (e *LambdaFormalParametersExpression) IsArrayExpression() bool                  { return false }
func (e *LambdaFormalParametersExpression) IsBinaryOperatorExpression() bool         { return false }
func (e *LambdaFormalParametersExpression) IsBooleanExpression() bool                { return false }
func (e *LambdaFormalParametersExpression) IsCastExpression() bool                   { return false }
func (e *LambdaFormalParametersExpression) IsConstructorInvocationExpression() bool  { return false }
func (e *LambdaFormalParametersExpression) IsDoubleConstantExpression() bool         { return false }
func (e *LambdaFormalParametersExpression) IsFieldReferenceExpression() bool         { return false }
func (e *LambdaFormalParametersExpression) IsFloatConstantExpression() bool          { return false }
func (e *LambdaFormalParametersExpression) IsIntegerConstantExpression() bool        { return false }
func (e *LambdaFormalParametersExpression) IsLengthExpression() bool                 { return false }
func (e *LambdaFormalParametersExpression) IsLocalVariableReferenceExpression() bool { return false }
func (e *LambdaFormalParametersExpression) IsLongConstantExpression() bool           { return false }
func (e *LambdaFormalParametersExpression) IsMethodInvocationExpression() bool       { return false }
func (e *LambdaFormalParametersExpression) IsNewArray() bool                         { return false }
func (e *LambdaFormalParametersExpression) IsNewExpression() bool                    { return false }
func (e *LambdaFormalParametersExpression) IsNewInitializedArray() bool              { return false }
func (e *LambdaFormalParametersExpression) IsNullExpression() bool                   { return false }
func (e *LambdaFormalParametersExpression) IsObjectTypeReferenceExpression() bool    { return false }
func (e *LambdaFormalParametersExpression) IsPostOperatorExpression() bool           { return false }
func (e *LambdaFormalParametersExpression) IsPreOperatorExpression() bool            { return false }
func (e *LambdaFormalParametersExpression) IsStringConstantExpression() bool         { return false }
func (e *LambdaFormalParametersExpression) IsSuperConstructorInvocationExpression() bool {
	return false
}
func (e *LambdaFormalParametersExpression) IsSuperExpression() bool           { return false }
func (e *LambdaFormalParametersExpression) IsTernaryOperatorExpression() bool { return false }
func (e *LambdaFormalParametersExpression) IsThisExpression() bool            { return false }

func (e *LambdaFormalParametersExpression) GetLineNumber() int        { return e.LineNumber }
func (e *LambdaFormalParametersExpression) GetType() IType            { return e.Type }
func (e *LambdaFormalParametersExpression) GetPriority() int          { return e.Priority }
func (e *LambdaFormalParametersExpression) GetStatements() IStatement { return e.Statements }

func (e *LambdaFormalParametersExpression) GetDimensionExpressionList() IExpression {
	return &NeNoExpression
}
func (e *LambdaFormalParametersExpression) GetParameters() IExpression { return &NeNoExpression }

func (e *LambdaFormalParametersExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *LambdaFormalParametersExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *LambdaFormalParametersExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *LambdaFormalParametersExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *LambdaFormalParametersExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *LambdaFormalParametersExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *LambdaFormalParametersExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *LambdaFormalParametersExpression) GetDescriptor() string       { return "" }
func (e *LambdaFormalParametersExpression) GetDoubleValue() float64     { return 0 }
func (e *LambdaFormalParametersExpression) GetFloatValue() float32      { return 0 }
func (e *LambdaFormalParametersExpression) GetIntegerValue() int        { return 0 }
func (e *LambdaFormalParametersExpression) GetInternalTypeName() string { return "" }
func (e *LambdaFormalParametersExpression) GetLongValue() int64         { return 0 }
func (e *LambdaFormalParametersExpression) GetName() string             { return "" }
func (e *LambdaFormalParametersExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *LambdaFormalParametersExpression) GetOperator() string         { return "" }
func (e *LambdaFormalParametersExpression) GetStringValue() string      { return "" }

func (e *LambdaFormalParametersExpression) String() string {
	return fmt.Sprintf("LambdaFormalParametersExpression{%s -> %d}", e.FormalParameters, e.Statements)
}

type LambdaIdentifiersExpression struct {
	util.DefaultBase[IExpression]

	LineNumber     int
	Type           IType
	Priority       int
	Statements     IStatement
	ReturnedType   IType
	ParameterNames util.IList[string]
}

func (e *LambdaIdentifiersExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitLambdaIdentifiersExpression(e)
}

func (e *LambdaIdentifiersExpression) IsArrayExpression() bool                      { return false }
func (e *LambdaIdentifiersExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *LambdaIdentifiersExpression) IsBooleanExpression() bool                    { return false }
func (e *LambdaIdentifiersExpression) IsCastExpression() bool                       { return false }
func (e *LambdaIdentifiersExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *LambdaIdentifiersExpression) IsDoubleConstantExpression() bool             { return false }
func (e *LambdaIdentifiersExpression) IsFieldReferenceExpression() bool             { return false }
func (e *LambdaIdentifiersExpression) IsFloatConstantExpression() bool              { return false }
func (e *LambdaIdentifiersExpression) IsIntegerConstantExpression() bool            { return false }
func (e *LambdaIdentifiersExpression) IsLengthExpression() bool                     { return false }
func (e *LambdaIdentifiersExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *LambdaIdentifiersExpression) IsLongConstantExpression() bool               { return false }
func (e *LambdaIdentifiersExpression) IsMethodInvocationExpression() bool           { return false }
func (e *LambdaIdentifiersExpression) IsNewArray() bool                             { return false }
func (e *LambdaIdentifiersExpression) IsNewExpression() bool                        { return false }
func (e *LambdaIdentifiersExpression) IsNewInitializedArray() bool                  { return false }
func (e *LambdaIdentifiersExpression) IsNullExpression() bool                       { return false }
func (e *LambdaIdentifiersExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *LambdaIdentifiersExpression) IsPostOperatorExpression() bool               { return false }
func (e *LambdaIdentifiersExpression) IsPreOperatorExpression() bool                { return false }
func (e *LambdaIdentifiersExpression) IsStringConstantExpression() bool             { return false }
func (e *LambdaIdentifiersExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *LambdaIdentifiersExpression) IsSuperExpression() bool                      { return false }
func (e *LambdaIdentifiersExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *LambdaIdentifiersExpression) IsThisExpression() bool                       { return false }

func (e *LambdaIdentifiersExpression) GetLineNumber() int        { return e.LineNumber }
func (e *LambdaIdentifiersExpression) GetType() IType            { return e.Type }
func (e *LambdaIdentifiersExpression) GetPriority() int          { return e.Priority }
func (e *LambdaIdentifiersExpression) GetStatements() IStatement { return e.Statements }

func (e *LambdaIdentifiersExpression) GetDimensionExpressionList() IExpression {
	return &NeNoExpression
}
func (e *LambdaIdentifiersExpression) GetParameters() IExpression { return &NeNoExpression }

func (e *LambdaIdentifiersExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *LambdaIdentifiersExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *LambdaIdentifiersExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *LambdaIdentifiersExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *LambdaIdentifiersExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *LambdaIdentifiersExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *LambdaIdentifiersExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *LambdaIdentifiersExpression) GetDescriptor() string       { return "" }
func (e *LambdaIdentifiersExpression) GetDoubleValue() float64     { return 0 }
func (e *LambdaIdentifiersExpression) GetFloatValue() float32      { return 0 }
func (e *LambdaIdentifiersExpression) GetIntegerValue() int        { return 0 }
func (e *LambdaIdentifiersExpression) GetInternalTypeName() string { return "" }
func (e *LambdaIdentifiersExpression) GetLongValue() int64         { return 0 }
func (e *LambdaIdentifiersExpression) GetName() string             { return "" }
func (e *LambdaIdentifiersExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *LambdaIdentifiersExpression) GetOperator() string         { return "" }
func (e *LambdaIdentifiersExpression) GetStringValue() string      { return "" }

func (e *LambdaIdentifiersExpression) String() string {
	return fmt.Sprintf("LambdaIdentifiersExpression{%s -> %d}", e.ParameterNames, e.Statements)
}

type LengthExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	Expression IExpression
}

func (e *LengthExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitLengthExpression(e)
}

func (e *LengthExpression) IsArrayExpression() bool                      { return false }
func (e *LengthExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *LengthExpression) IsBooleanExpression() bool                    { return false }
func (e *LengthExpression) IsCastExpression() bool                       { return false }
func (e *LengthExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *LengthExpression) IsDoubleConstantExpression() bool             { return false }
func (e *LengthExpression) IsFieldReferenceExpression() bool             { return false }
func (e *LengthExpression) IsFloatConstantExpression() bool              { return false }
func (e *LengthExpression) IsIntegerConstantExpression() bool            { return false }
func (e *LengthExpression) IsLengthExpression() bool                     { return true }
func (e *LengthExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *LengthExpression) IsLongConstantExpression() bool               { return false }
func (e *LengthExpression) IsMethodInvocationExpression() bool           { return false }
func (e *LengthExpression) IsNewArray() bool                             { return false }
func (e *LengthExpression) IsNewExpression() bool                        { return false }
func (e *LengthExpression) IsNewInitializedArray() bool                  { return false }
func (e *LengthExpression) IsNullExpression() bool                       { return false }
func (e *LengthExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *LengthExpression) IsPostOperatorExpression() bool               { return false }
func (e *LengthExpression) IsPreOperatorExpression() bool                { return false }
func (e *LengthExpression) IsStringConstantExpression() bool             { return false }
func (e *LengthExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *LengthExpression) IsSuperExpression() bool                      { return false }
func (e *LengthExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *LengthExpression) IsThisExpression() bool                       { return false }

func (e *LengthExpression) GetLineNumber() int { return e.LineNumber }
func (e *LengthExpression) GetType() IType     { return e.Type }
func (e *LengthExpression) GetPriority() int   { return e.Priority }

func (e *LengthExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *LengthExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *LengthExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *LengthExpression) GetExpression() IExpression      { return e.Expression }
func (e *LengthExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *LengthExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *LengthExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *LengthExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *LengthExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *LengthExpression) GetDescriptor() string       { return "" }
func (e *LengthExpression) GetDoubleValue() float64     { return 0 }
func (e *LengthExpression) GetFloatValue() float32      { return 0 }
func (e *LengthExpression) GetIntegerValue() int        { return 0 }
func (e *LengthExpression) GetInternalTypeName() string { return "" }
func (e *LengthExpression) GetLongValue() int64         { return 0 }
func (e *LengthExpression) GetName() string             { return "" }
func (e *LengthExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *LengthExpression) GetOperator() string         { return "" }
func (e *LengthExpression) GetStringValue() string      { return "" }

func (e *LengthExpression) String() string {
	return fmt.Sprintf("LengthExpression{%s}", e.Expression)
}

type LocalVariableReferenceExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	Name       string
}

func (e *LocalVariableReferenceExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitLocalVariableReferenceExpression(e)
}

func (e *LocalVariableReferenceExpression) IsArrayExpression() bool                  { return false }
func (e *LocalVariableReferenceExpression) IsBinaryOperatorExpression() bool         { return false }
func (e *LocalVariableReferenceExpression) IsBooleanExpression() bool                { return false }
func (e *LocalVariableReferenceExpression) IsCastExpression() bool                   { return false }
func (e *LocalVariableReferenceExpression) IsConstructorInvocationExpression() bool  { return false }
func (e *LocalVariableReferenceExpression) IsDoubleConstantExpression() bool         { return false }
func (e *LocalVariableReferenceExpression) IsFieldReferenceExpression() bool         { return false }
func (e *LocalVariableReferenceExpression) IsFloatConstantExpression() bool          { return false }
func (e *LocalVariableReferenceExpression) IsIntegerConstantExpression() bool        { return false }
func (e *LocalVariableReferenceExpression) IsLengthExpression() bool                 { return false }
func (e *LocalVariableReferenceExpression) IsLocalVariableReferenceExpression() bool { return true }
func (e *LocalVariableReferenceExpression) IsLongConstantExpression() bool           { return false }
func (e *LocalVariableReferenceExpression) IsMethodInvocationExpression() bool       { return false }
func (e *LocalVariableReferenceExpression) IsNewArray() bool                         { return false }
func (e *LocalVariableReferenceExpression) IsNewExpression() bool                    { return false }
func (e *LocalVariableReferenceExpression) IsNewInitializedArray() bool              { return false }
func (e *LocalVariableReferenceExpression) IsNullExpression() bool                   { return false }
func (e *LocalVariableReferenceExpression) IsObjectTypeReferenceExpression() bool    { return false }
func (e *LocalVariableReferenceExpression) IsPostOperatorExpression() bool           { return false }
func (e *LocalVariableReferenceExpression) IsPreOperatorExpression() bool            { return false }
func (e *LocalVariableReferenceExpression) IsStringConstantExpression() bool         { return false }
func (e *LocalVariableReferenceExpression) IsSuperConstructorInvocationExpression() bool {
	return false
}
func (e *LocalVariableReferenceExpression) IsSuperExpression() bool           { return false }
func (e *LocalVariableReferenceExpression) IsTernaryOperatorExpression() bool { return false }
func (e *LocalVariableReferenceExpression) IsThisExpression() bool            { return false }

func (e *LocalVariableReferenceExpression) GetLineNumber() int { return e.LineNumber }
func (e *LocalVariableReferenceExpression) GetType() IType     { return e.Type }
func (e *LocalVariableReferenceExpression) GetPriority() int   { return e.Priority }

func (e *LocalVariableReferenceExpression) GetDimensionExpressionList() IExpression {
	return &NeNoExpression
}
func (e *LocalVariableReferenceExpression) GetParameters() IExpression { return &NeNoExpression }

func (e *LocalVariableReferenceExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *LocalVariableReferenceExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *LocalVariableReferenceExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *LocalVariableReferenceExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *LocalVariableReferenceExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *LocalVariableReferenceExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *LocalVariableReferenceExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *LocalVariableReferenceExpression) GetDescriptor() string       { return "" }
func (e *LocalVariableReferenceExpression) GetDoubleValue() float64     { return 0 }
func (e *LocalVariableReferenceExpression) GetFloatValue() float32      { return 0 }
func (e *LocalVariableReferenceExpression) GetIntegerValue() int        { return 0 }
func (e *LocalVariableReferenceExpression) GetInternalTypeName() string { return "" }
func (e *LocalVariableReferenceExpression) GetLongValue() int64         { return 0 }
func (e *LocalVariableReferenceExpression) GetName() string             { return e.Name }
func (e *LocalVariableReferenceExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *LocalVariableReferenceExpression) GetOperator() string         { return "" }
func (e *LocalVariableReferenceExpression) GetStringValue() string      { return "" }

func (e *LocalVariableReferenceExpression) String() string {
	return fmt.Sprintf("LocalVariableReferenceExpression{type=%s, Name=%s}", e.Type, e.Name)
}

type LongConstantExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	Value      int64
}

func (e *LongConstantExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitLongConstantExpression(e)
}

func (e *LongConstantExpression) IsArrayExpression() bool                      { return false }
func (e *LongConstantExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *LongConstantExpression) IsBooleanExpression() bool                    { return false }
func (e *LongConstantExpression) IsCastExpression() bool                       { return false }
func (e *LongConstantExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *LongConstantExpression) IsDoubleConstantExpression() bool             { return false }
func (e *LongConstantExpression) IsFieldReferenceExpression() bool             { return false }
func (e *LongConstantExpression) IsFloatConstantExpression() bool              { return false }
func (e *LongConstantExpression) IsIntegerConstantExpression() bool            { return false }
func (e *LongConstantExpression) IsLengthExpression() bool                     { return false }
func (e *LongConstantExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *LongConstantExpression) IsLongConstantExpression() bool               { return true }
func (e *LongConstantExpression) IsMethodInvocationExpression() bool           { return false }
func (e *LongConstantExpression) IsNewArray() bool                             { return false }
func (e *LongConstantExpression) IsNewExpression() bool                        { return false }
func (e *LongConstantExpression) IsNewInitializedArray() bool                  { return false }
func (e *LongConstantExpression) IsNullExpression() bool                       { return false }
func (e *LongConstantExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *LongConstantExpression) IsPostOperatorExpression() bool               { return false }
func (e *LongConstantExpression) IsPreOperatorExpression() bool                { return false }
func (e *LongConstantExpression) IsStringConstantExpression() bool             { return false }
func (e *LongConstantExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *LongConstantExpression) IsSuperExpression() bool                      { return false }
func (e *LongConstantExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *LongConstantExpression) IsThisExpression() bool                       { return false }

func (e *LongConstantExpression) GetLineNumber() int { return e.LineNumber }
func (e *LongConstantExpression) GetType() IType     { return e.Type }
func (e *LongConstantExpression) GetPriority() int   { return e.Priority }

func (e *LongConstantExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *LongConstantExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *LongConstantExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *LongConstantExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *LongConstantExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *LongConstantExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *LongConstantExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *LongConstantExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *LongConstantExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *LongConstantExpression) GetDescriptor() string       { return "" }
func (e *LongConstantExpression) GetDoubleValue() float64     { return 0 }
func (e *LongConstantExpression) GetFloatValue() float32      { return 0 }
func (e *LongConstantExpression) GetIntegerValue() int        { return 0 }
func (e *LongConstantExpression) GetInternalTypeName() string { return "" }
func (e *LongConstantExpression) GetLongValue() int64         { return e.Value }
func (e *LongConstantExpression) GetName() string             { return "" }
func (e *LongConstantExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *LongConstantExpression) GetOperator() string         { return "" }
func (e *LongConstantExpression) GetStringValue() string      { return "" }

func (e *LongConstantExpression) String() string {
	return fmt.Sprintf("LongConstantExpression{%d}", e.Value)
}

type MethodInvocationExpression struct {
	util.DefaultBase[IExpression]

	LineNumber               int
	Type                     IType
	Priority                 int
	Expression               IExpression
	InternalTypeName         string
	Name                     string
	Descriptor               string
	NonWildcardTypeArguments ITypeArgument
	Parameters               IExpression
}

func (e *MethodInvocationExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitMethodInvocationExpression(e)
}

func (e *MethodInvocationExpression) IsArrayExpression() bool                      { return false }
func (e *MethodInvocationExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *MethodInvocationExpression) IsBooleanExpression() bool                    { return false }
func (e *MethodInvocationExpression) IsCastExpression() bool                       { return false }
func (e *MethodInvocationExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *MethodInvocationExpression) IsDoubleConstantExpression() bool             { return false }
func (e *MethodInvocationExpression) IsFieldReferenceExpression() bool             { return false }
func (e *MethodInvocationExpression) IsFloatConstantExpression() bool              { return false }
func (e *MethodInvocationExpression) IsIntegerConstantExpression() bool            { return false }
func (e *MethodInvocationExpression) IsLengthExpression() bool                     { return false }
func (e *MethodInvocationExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *MethodInvocationExpression) IsLongConstantExpression() bool               { return false }
func (e *MethodInvocationExpression) IsMethodInvocationExpression() bool           { return true }
func (e *MethodInvocationExpression) IsNewArray() bool                             { return false }
func (e *MethodInvocationExpression) IsNewExpression() bool                        { return false }
func (e *MethodInvocationExpression) IsNewInitializedArray() bool                  { return false }
func (e *MethodInvocationExpression) IsNullExpression() bool                       { return false }
func (e *MethodInvocationExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *MethodInvocationExpression) IsPostOperatorExpression() bool               { return false }
func (e *MethodInvocationExpression) IsPreOperatorExpression() bool                { return false }
func (e *MethodInvocationExpression) IsStringConstantExpression() bool             { return false }
func (e *MethodInvocationExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *MethodInvocationExpression) IsSuperExpression() bool                      { return false }
func (e *MethodInvocationExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *MethodInvocationExpression) IsThisExpression() bool                       { return false }

func (e *MethodInvocationExpression) GetLineNumber() int { return e.LineNumber }
func (e *MethodInvocationExpression) GetType() IType     { return e.Type }
func (e *MethodInvocationExpression) GetPriority() int   { return e.Priority }

func (e *MethodInvocationExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *MethodInvocationExpression) GetParameters() IExpression              { return e.Parameters }

func (e *MethodInvocationExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *MethodInvocationExpression) GetExpression() IExpression      { return e.Expression }
func (e *MethodInvocationExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *MethodInvocationExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *MethodInvocationExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *MethodInvocationExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *MethodInvocationExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *MethodInvocationExpression) GetDescriptor() string       { return e.Descriptor }
func (e *MethodInvocationExpression) GetDoubleValue() float64     { return 0 }
func (e *MethodInvocationExpression) GetFloatValue() float32      { return 0 }
func (e *MethodInvocationExpression) GetIntegerValue() int        { return 0 }
func (e *MethodInvocationExpression) GetInternalTypeName() string { return e.InternalTypeName }
func (e *MethodInvocationExpression) GetLongValue() int64         { return 0 }
func (e *MethodInvocationExpression) GetName() string             { return e.Name }
func (e *MethodInvocationExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *MethodInvocationExpression) GetOperator() string         { return "" }
func (e *MethodInvocationExpression) GetStringValue() string      { return "" }

func (e *MethodInvocationExpression) String() string {
	return fmt.Sprintf("MethodInvocationExpression{call %s . %s (%s)}", e.Expression, e.Name, e.Descriptor)
}

type MethodReferenceExpression struct {
	util.DefaultBase[IExpression]

	LineNumber       int
	Type             IType
	Priority         int
	Expression       IExpression
	InternalTypeName string
	Name             string
	Descriptor       string
}

func (e *MethodReferenceExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitMethodReferenceExpression(e)
}

func (e *MethodReferenceExpression) IsArrayExpression() bool                      { return false }
func (e *MethodReferenceExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *MethodReferenceExpression) IsBooleanExpression() bool                    { return false }
func (e *MethodReferenceExpression) IsCastExpression() bool                       { return false }
func (e *MethodReferenceExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *MethodReferenceExpression) IsDoubleConstantExpression() bool             { return false }
func (e *MethodReferenceExpression) IsFieldReferenceExpression() bool             { return false }
func (e *MethodReferenceExpression) IsFloatConstantExpression() bool              { return false }
func (e *MethodReferenceExpression) IsIntegerConstantExpression() bool            { return false }
func (e *MethodReferenceExpression) IsLengthExpression() bool                     { return false }
func (e *MethodReferenceExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *MethodReferenceExpression) IsLongConstantExpression() bool               { return false }
func (e *MethodReferenceExpression) IsMethodInvocationExpression() bool           { return false }
func (e *MethodReferenceExpression) IsNewArray() bool                             { return false }
func (e *MethodReferenceExpression) IsNewExpression() bool                        { return false }
func (e *MethodReferenceExpression) IsNewInitializedArray() bool                  { return false }
func (e *MethodReferenceExpression) IsNullExpression() bool                       { return false }
func (e *MethodReferenceExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *MethodReferenceExpression) IsPostOperatorExpression() bool               { return false }
func (e *MethodReferenceExpression) IsPreOperatorExpression() bool                { return false }
func (e *MethodReferenceExpression) IsStringConstantExpression() bool             { return false }
func (e *MethodReferenceExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *MethodReferenceExpression) IsSuperExpression() bool                      { return false }
func (e *MethodReferenceExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *MethodReferenceExpression) IsThisExpression() bool                       { return false }

func (e *MethodReferenceExpression) GetLineNumber() int { return e.LineNumber }
func (e *MethodReferenceExpression) GetType() IType     { return e.Type }
func (e *MethodReferenceExpression) GetPriority() int   { return e.Priority }

func (e *MethodReferenceExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *MethodReferenceExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *MethodReferenceExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *MethodReferenceExpression) GetExpression() IExpression      { return e.Expression }
func (e *MethodReferenceExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *MethodReferenceExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *MethodReferenceExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *MethodReferenceExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *MethodReferenceExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *MethodReferenceExpression) GetDescriptor() string       { return e.Descriptor }
func (e *MethodReferenceExpression) GetDoubleValue() float64     { return 0 }
func (e *MethodReferenceExpression) GetFloatValue() float32      { return 0 }
func (e *MethodReferenceExpression) GetIntegerValue() int        { return 0 }
func (e *MethodReferenceExpression) GetInternalTypeName() string { return e.InternalTypeName }
func (e *MethodReferenceExpression) GetLongValue() int64         { return 0 }
func (e *MethodReferenceExpression) GetName() string             { return e.Name }
func (e *MethodReferenceExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *MethodReferenceExpression) GetOperator() string         { return "" }
func (e *MethodReferenceExpression) GetStringValue() string      { return "" }

func (e *MethodReferenceExpression) String() string {
	return fmt.Sprintf("MethodReferenceExpression{call %s . %s (%s)}", e.Expression, e.Name, e.Descriptor)
}

type NewArray struct {
	util.DefaultBase[IExpression]

	LineNumber              int
	Type                    IType
	Priority                int
	DimensionExpressionList IExpression
}

func (e *NewArray) Accept(visitor IExpressionVisitor) {
	visitor.VisitNewArray(e)
}

func (e *NewArray) IsArrayExpression() bool                      { return false }
func (e *NewArray) IsBinaryOperatorExpression() bool             { return false }
func (e *NewArray) IsBooleanExpression() bool                    { return false }
func (e *NewArray) IsCastExpression() bool                       { return false }
func (e *NewArray) IsConstructorInvocationExpression() bool      { return false }
func (e *NewArray) IsDoubleConstantExpression() bool             { return false }
func (e *NewArray) IsFieldReferenceExpression() bool             { return false }
func (e *NewArray) IsFloatConstantExpression() bool              { return false }
func (e *NewArray) IsIntegerConstantExpression() bool            { return false }
func (e *NewArray) IsLengthExpression() bool                     { return false }
func (e *NewArray) IsLocalVariableReferenceExpression() bool     { return false }
func (e *NewArray) IsLongConstantExpression() bool               { return false }
func (e *NewArray) IsMethodInvocationExpression() bool           { return false }
func (e *NewArray) IsNewArray() bool                             { return true }
func (e *NewArray) IsNewExpression() bool                        { return false }
func (e *NewArray) IsNewInitializedArray() bool                  { return false }
func (e *NewArray) IsNullExpression() bool                       { return false }
func (e *NewArray) IsObjectTypeReferenceExpression() bool        { return false }
func (e *NewArray) IsPostOperatorExpression() bool               { return false }
func (e *NewArray) IsPreOperatorExpression() bool                { return false }
func (e *NewArray) IsStringConstantExpression() bool             { return false }
func (e *NewArray) IsSuperConstructorInvocationExpression() bool { return false }
func (e *NewArray) IsSuperExpression() bool                      { return false }
func (e *NewArray) IsTernaryOperatorExpression() bool            { return false }
func (e *NewArray) IsThisExpression() bool                       { return false }

func (e *NewArray) GetLineNumber() int { return e.LineNumber }
func (e *NewArray) GetType() IType     { return e.Type }
func (e *NewArray) GetPriority() int   { return e.Priority }

func (e *NewArray) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *NewArray) GetParameters() IExpression              { return &NeNoExpression }

func (e *NewArray) GetCondition() IExpression       { return &NeNoExpression }
func (e *NewArray) GetExpression() IExpression      { return &NeNoExpression }
func (e *NewArray) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *NewArray) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *NewArray) GetIndex() IExpression           { return &NeNoExpression }
func (e *NewArray) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *NewArray) GetRightExpression() IExpression { return &NeNoExpression }

func (e *NewArray) GetDescriptor() string       { return "" }
func (e *NewArray) GetDoubleValue() float64     { return 0 }
func (e *NewArray) GetFloatValue() float32      { return 0 }
func (e *NewArray) GetIntegerValue() int        { return 0 }
func (e *NewArray) GetInternalTypeName() string { return "" }
func (e *NewArray) GetLongValue() int64         { return 0 }
func (e *NewArray) GetName() string             { return "" }
func (e *NewArray) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *NewArray) GetOperator() string         { return "" }
func (e *NewArray) GetStringValue() string      { return "" }

func (e *NewArray) String() string {
	return fmt.Sprintf("NewArray{%s}", e.Type)
}

type NewExpression struct {
	util.DefaultBase[IExpression]

	LineNumber      int
	Priority        int
	Type            *ObjectType
	Descriptor      string
	Parameters      IExpression
	BodyDeclaration *BodyDeclaration
}

func (e *NewExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitNewExpression(e)
}

func (e *NewExpression) IsArrayExpression() bool                      { return false }
func (e *NewExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *NewExpression) IsBooleanExpression() bool                    { return false }
func (e *NewExpression) IsCastExpression() bool                       { return false }
func (e *NewExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *NewExpression) IsDoubleConstantExpression() bool             { return false }
func (e *NewExpression) IsFieldReferenceExpression() bool             { return false }
func (e *NewExpression) IsFloatConstantExpression() bool              { return false }
func (e *NewExpression) IsIntegerConstantExpression() bool            { return false }
func (e *NewExpression) IsLengthExpression() bool                     { return false }
func (e *NewExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *NewExpression) IsLongConstantExpression() bool               { return false }
func (e *NewExpression) IsMethodInvocationExpression() bool           { return false }
func (e *NewExpression) IsNewArray() bool                             { return false }
func (e *NewExpression) IsNewExpression() bool                        { return true }
func (e *NewExpression) IsNewInitializedArray() bool                  { return false }
func (e *NewExpression) IsNullExpression() bool                       { return false }
func (e *NewExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *NewExpression) IsPostOperatorExpression() bool               { return false }
func (e *NewExpression) IsPreOperatorExpression() bool                { return false }
func (e *NewExpression) IsStringConstantExpression() bool             { return false }
func (e *NewExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *NewExpression) IsSuperExpression() bool                      { return false }
func (e *NewExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *NewExpression) IsThisExpression() bool                       { return false }

func (e *NewExpression) GetLineNumber() int { return e.LineNumber }
func (e *NewExpression) GetType() IType     { return e.Type }
func (e *NewExpression) GetPriority() int   { return e.Priority }

func (e *NewExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *NewExpression) GetParameters() IExpression              { return e.Parameters }

func (e *NewExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *NewExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *NewExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *NewExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *NewExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *NewExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *NewExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *NewExpression) GetDescriptor() string       { return e.Descriptor }
func (e *NewExpression) GetDoubleValue() float64     { return 0 }
func (e *NewExpression) GetFloatValue() float32      { return 0 }
func (e *NewExpression) GetIntegerValue() int        { return 0 }
func (e *NewExpression) GetInternalTypeName() string { return "" }
func (e *NewExpression) GetLongValue() int64         { return 0 }
func (e *NewExpression) GetName() string             { return "" }
func (e *NewExpression) GetObjectType() *ObjectType  { return e.Type }
func (e *NewExpression) GetOperator() string         { return "" }
func (e *NewExpression) GetStringValue() string      { return "" }

func (e *NewExpression) String() string {
	return fmt.Sprintf("NewExpression{new %s}", e.Type)
}

type NewInitializedArray struct {
	util.DefaultBase[IExpression]

	LineNumber       int
	Type             IType
	Priority         int
	ArrayInitializer ArrayVariableInitializer
}

func (e *NewInitializedArray) Accept(visitor IExpressionVisitor) {
	visitor.VisitNewInitializedArray(e)
}

func (e *NewInitializedArray) IsArrayExpression() bool                      { return false }
func (e *NewInitializedArray) IsBinaryOperatorExpression() bool             { return false }
func (e *NewInitializedArray) IsBooleanExpression() bool                    { return false }
func (e *NewInitializedArray) IsCastExpression() bool                       { return false }
func (e *NewInitializedArray) IsConstructorInvocationExpression() bool      { return false }
func (e *NewInitializedArray) IsDoubleConstantExpression() bool             { return false }
func (e *NewInitializedArray) IsFieldReferenceExpression() bool             { return false }
func (e *NewInitializedArray) IsFloatConstantExpression() bool              { return false }
func (e *NewInitializedArray) IsIntegerConstantExpression() bool            { return false }
func (e *NewInitializedArray) IsLengthExpression() bool                     { return false }
func (e *NewInitializedArray) IsLocalVariableReferenceExpression() bool     { return false }
func (e *NewInitializedArray) IsLongConstantExpression() bool               { return false }
func (e *NewInitializedArray) IsMethodInvocationExpression() bool           { return false }
func (e *NewInitializedArray) IsNewArray() bool                             { return false }
func (e *NewInitializedArray) IsNewExpression() bool                        { return false }
func (e *NewInitializedArray) IsNewInitializedArray() bool                  { return true }
func (e *NewInitializedArray) IsNullExpression() bool                       { return false }
func (e *NewInitializedArray) IsObjectTypeReferenceExpression() bool        { return false }
func (e *NewInitializedArray) IsPostOperatorExpression() bool               { return false }
func (e *NewInitializedArray) IsPreOperatorExpression() bool                { return false }
func (e *NewInitializedArray) IsStringConstantExpression() bool             { return false }
func (e *NewInitializedArray) IsSuperConstructorInvocationExpression() bool { return false }
func (e *NewInitializedArray) IsSuperExpression() bool                      { return false }
func (e *NewInitializedArray) IsTernaryOperatorExpression() bool            { return false }
func (e *NewInitializedArray) IsThisExpression() bool                       { return false }

func (e *NewInitializedArray) GetLineNumber() int { return e.LineNumber }
func (e *NewInitializedArray) GetType() IType     { return e.Type }
func (e *NewInitializedArray) GetPriority() int   { return e.Priority }

func (e *NewInitializedArray) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *NewInitializedArray) GetParameters() IExpression              { return &NeNoExpression }

func (e *NewInitializedArray) GetCondition() IExpression       { return &NeNoExpression }
func (e *NewInitializedArray) GetExpression() IExpression      { return &NeNoExpression }
func (e *NewInitializedArray) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *NewInitializedArray) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *NewInitializedArray) GetIndex() IExpression           { return &NeNoExpression }
func (e *NewInitializedArray) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *NewInitializedArray) GetRightExpression() IExpression { return &NeNoExpression }

func (e *NewInitializedArray) GetDescriptor() string       { return "" }
func (e *NewInitializedArray) GetDoubleValue() float64     { return 0 }
func (e *NewInitializedArray) GetFloatValue() float32      { return 0 }
func (e *NewInitializedArray) GetIntegerValue() int        { return 0 }
func (e *NewInitializedArray) GetInternalTypeName() string { return "" }
func (e *NewInitializedArray) GetLongValue() int64         { return 0 }
func (e *NewInitializedArray) GetName() string             { return "" }
func (e *NewInitializedArray) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *NewInitializedArray) GetOperator() string         { return "" }
func (e *NewInitializedArray) GetStringValue() string      { return "" }

func (e *NewInitializedArray) String() string {
	return fmt.Sprintf("NewInitializedArray{new %s [%s]}", e.Type, e.ArrayInitializer)
}

type NoExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
}

func (e *NoExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitNoExpression(e)
}

func (e *NoExpression) IsArrayExpression() bool                      { return false }
func (e *NoExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *NoExpression) IsBooleanExpression() bool                    { return false }
func (e *NoExpression) IsCastExpression() bool                       { return false }
func (e *NoExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *NoExpression) IsDoubleConstantExpression() bool             { return false }
func (e *NoExpression) IsFieldReferenceExpression() bool             { return false }
func (e *NoExpression) IsFloatConstantExpression() bool              { return false }
func (e *NoExpression) IsIntegerConstantExpression() bool            { return false }
func (e *NoExpression) IsLengthExpression() bool                     { return false }
func (e *NoExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *NoExpression) IsLongConstantExpression() bool               { return false }
func (e *NoExpression) IsMethodInvocationExpression() bool           { return false }
func (e *NoExpression) IsNewArray() bool                             { return false }
func (e *NoExpression) IsNewExpression() bool                        { return false }
func (e *NoExpression) IsNewInitializedArray() bool                  { return false }
func (e *NoExpression) IsNullExpression() bool                       { return false }
func (e *NoExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *NoExpression) IsPostOperatorExpression() bool               { return false }
func (e *NoExpression) IsPreOperatorExpression() bool                { return false }
func (e *NoExpression) IsStringConstantExpression() bool             { return false }
func (e *NoExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *NoExpression) IsSuperExpression() bool                      { return false }
func (e *NoExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *NoExpression) IsThisExpression() bool                       { return false }

func (e *NoExpression) GetLineNumber() int { return e.LineNumber }
func (e *NoExpression) GetType() IType     { return e.Type }
func (e *NoExpression) GetPriority() int   { return e.Priority }

func (e *NoExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *NoExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *NoExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *NoExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *NoExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *NoExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *NoExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *NoExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *NoExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *NoExpression) GetDescriptor() string       { return "" }
func (e *NoExpression) GetDoubleValue() float64     { return 0 }
func (e *NoExpression) GetFloatValue() float32      { return 0 }
func (e *NoExpression) GetIntegerValue() int        { return 0 }
func (e *NoExpression) GetInternalTypeName() string { return "" }
func (e *NoExpression) GetLongValue() int64         { return 0 }
func (e *NoExpression) GetName() string             { return "" }
func (e *NoExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *NoExpression) GetOperator() string         { return "" }
func (e *NoExpression) GetStringValue() string      { return "" }

func (e *NoExpression) String() string {
	return "NoExpression{}"
}

type NullExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
}

func (e *NullExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitNullExpression(e)
}

func (e *NullExpression) IsArrayExpression() bool                      { return false }
func (e *NullExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *NullExpression) IsBooleanExpression() bool                    { return false }
func (e *NullExpression) IsCastExpression() bool                       { return false }
func (e *NullExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *NullExpression) IsDoubleConstantExpression() bool             { return false }
func (e *NullExpression) IsFieldReferenceExpression() bool             { return false }
func (e *NullExpression) IsFloatConstantExpression() bool              { return false }
func (e *NullExpression) IsIntegerConstantExpression() bool            { return false }
func (e *NullExpression) IsLengthExpression() bool                     { return false }
func (e *NullExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *NullExpression) IsLongConstantExpression() bool               { return false }
func (e *NullExpression) IsMethodInvocationExpression() bool           { return false }
func (e *NullExpression) IsNewArray() bool                             { return false }
func (e *NullExpression) IsNewExpression() bool                        { return false }
func (e *NullExpression) IsNewInitializedArray() bool                  { return false }
func (e *NullExpression) IsNullExpression() bool                       { return true }
func (e *NullExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *NullExpression) IsPostOperatorExpression() bool               { return false }
func (e *NullExpression) IsPreOperatorExpression() bool                { return false }
func (e *NullExpression) IsStringConstantExpression() bool             { return false }
func (e *NullExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *NullExpression) IsSuperExpression() bool                      { return false }
func (e *NullExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *NullExpression) IsThisExpression() bool                       { return false }

func (e *NullExpression) GetLineNumber() int { return e.LineNumber }
func (e *NullExpression) GetType() IType     { return e.Type }
func (e *NullExpression) GetPriority() int   { return e.Priority }

func (e *NullExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *NullExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *NullExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *NullExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *NullExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *NullExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *NullExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *NullExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *NullExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *NullExpression) GetDescriptor() string       { return "" }
func (e *NullExpression) GetDoubleValue() float64     { return 0 }
func (e *NullExpression) GetFloatValue() float32      { return 0 }
func (e *NullExpression) GetIntegerValue() int        { return 0 }
func (e *NullExpression) GetInternalTypeName() string { return "" }
func (e *NullExpression) GetLongValue() int64         { return 0 }
func (e *NullExpression) GetName() string             { return "" }
func (e *NullExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *NullExpression) GetOperator() string         { return "" }
func (e *NullExpression) GetStringValue() string      { return "" }

func (e *NullExpression) String() string {
	return fmt.Sprintf("NullExpression{type=%s}", e.Type)
}

type ObjectTypeReferenceExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       *ObjectType
	Priority   int
	IsExplicit bool
}

func (e *ObjectTypeReferenceExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitObjectTypeReferenceExpression(e)
}

func (e *ObjectTypeReferenceExpression) IsArrayExpression() bool                      { return false }
func (e *ObjectTypeReferenceExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *ObjectTypeReferenceExpression) IsBooleanExpression() bool                    { return false }
func (e *ObjectTypeReferenceExpression) IsCastExpression() bool                       { return false }
func (e *ObjectTypeReferenceExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *ObjectTypeReferenceExpression) IsDoubleConstantExpression() bool             { return false }
func (e *ObjectTypeReferenceExpression) IsFieldReferenceExpression() bool             { return false }
func (e *ObjectTypeReferenceExpression) IsFloatConstantExpression() bool              { return false }
func (e *ObjectTypeReferenceExpression) IsIntegerConstantExpression() bool            { return false }
func (e *ObjectTypeReferenceExpression) IsLengthExpression() bool                     { return false }
func (e *ObjectTypeReferenceExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *ObjectTypeReferenceExpression) IsLongConstantExpression() bool               { return false }
func (e *ObjectTypeReferenceExpression) IsMethodInvocationExpression() bool           { return false }
func (e *ObjectTypeReferenceExpression) IsNewArray() bool                             { return false }
func (e *ObjectTypeReferenceExpression) IsNewExpression() bool                        { return false }
func (e *ObjectTypeReferenceExpression) IsNewInitializedArray() bool                  { return false }
func (e *ObjectTypeReferenceExpression) IsNullExpression() bool                       { return false }
func (e *ObjectTypeReferenceExpression) IsObjectTypeReferenceExpression() bool        { return true }
func (e *ObjectTypeReferenceExpression) IsPostOperatorExpression() bool               { return false }
func (e *ObjectTypeReferenceExpression) IsPreOperatorExpression() bool                { return false }
func (e *ObjectTypeReferenceExpression) IsStringConstantExpression() bool             { return false }
func (e *ObjectTypeReferenceExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *ObjectTypeReferenceExpression) IsSuperExpression() bool                      { return false }
func (e *ObjectTypeReferenceExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *ObjectTypeReferenceExpression) IsThisExpression() bool                       { return false }

func (e *ObjectTypeReferenceExpression) GetLineNumber() int { return e.LineNumber }
func (e *ObjectTypeReferenceExpression) GetType() IType     { return e.Type }
func (e *ObjectTypeReferenceExpression) GetPriority() int   { return e.Priority }

func (e *ObjectTypeReferenceExpression) GetDimensionExpressionList() IExpression {
	return &NeNoExpression
}
func (e *ObjectTypeReferenceExpression) GetParameters() IExpression { return &NeNoExpression }

func (e *ObjectTypeReferenceExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *ObjectTypeReferenceExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *ObjectTypeReferenceExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *ObjectTypeReferenceExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *ObjectTypeReferenceExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *ObjectTypeReferenceExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *ObjectTypeReferenceExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *ObjectTypeReferenceExpression) GetDescriptor() string       { return "" }
func (e *ObjectTypeReferenceExpression) GetDoubleValue() float64     { return 0 }
func (e *ObjectTypeReferenceExpression) GetFloatValue() float32      { return 0 }
func (e *ObjectTypeReferenceExpression) GetIntegerValue() int        { return 0 }
func (e *ObjectTypeReferenceExpression) GetInternalTypeName() string { return "" }
func (e *ObjectTypeReferenceExpression) GetLongValue() int64         { return 0 }
func (e *ObjectTypeReferenceExpression) GetName() string             { return "" }
func (e *ObjectTypeReferenceExpression) GetObjectType() *ObjectType  { return e.Type }
func (e *ObjectTypeReferenceExpression) GetOperator() string         { return "" }
func (e *ObjectTypeReferenceExpression) GetStringValue() string      { return "" }

func (e *ObjectTypeReferenceExpression) String() string {
	return fmt.Sprintf("ObjectTypeReferenceExpression{%s}", e.Type)
}

type ParenthesesExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Priority   int
	Expression IExpression
}

func (e *ParenthesesExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitParenthesesExpression(e)
}

func (e *ParenthesesExpression) IsArrayExpression() bool                      { return false }
func (e *ParenthesesExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *ParenthesesExpression) IsBooleanExpression() bool                    { return false }
func (e *ParenthesesExpression) IsCastExpression() bool                       { return false }
func (e *ParenthesesExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *ParenthesesExpression) IsDoubleConstantExpression() bool             { return false }
func (e *ParenthesesExpression) IsFieldReferenceExpression() bool             { return false }
func (e *ParenthesesExpression) IsFloatConstantExpression() bool              { return false }
func (e *ParenthesesExpression) IsIntegerConstantExpression() bool            { return false }
func (e *ParenthesesExpression) IsLengthExpression() bool                     { return false }
func (e *ParenthesesExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *ParenthesesExpression) IsLongConstantExpression() bool               { return false }
func (e *ParenthesesExpression) IsMethodInvocationExpression() bool           { return false }
func (e *ParenthesesExpression) IsNewArray() bool                             { return false }
func (e *ParenthesesExpression) IsNewExpression() bool                        { return false }
func (e *ParenthesesExpression) IsNewInitializedArray() bool                  { return false }
func (e *ParenthesesExpression) IsNullExpression() bool                       { return false }
func (e *ParenthesesExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *ParenthesesExpression) IsPostOperatorExpression() bool               { return false }
func (e *ParenthesesExpression) IsPreOperatorExpression() bool                { return false }
func (e *ParenthesesExpression) IsStringConstantExpression() bool             { return false }
func (e *ParenthesesExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *ParenthesesExpression) IsSuperExpression() bool                      { return false }
func (e *ParenthesesExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *ParenthesesExpression) IsThisExpression() bool                       { return false }

func (e *ParenthesesExpression) GetLineNumber() int { return e.LineNumber }
func (e *ParenthesesExpression) GetType() IType     { return e.Expression.GetType() }
func (e *ParenthesesExpression) GetPriority() int   { return e.Priority }

func (e *ParenthesesExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *ParenthesesExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *ParenthesesExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *ParenthesesExpression) GetExpression() IExpression      { return e.Expression }
func (e *ParenthesesExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *ParenthesesExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *ParenthesesExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *ParenthesesExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *ParenthesesExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *ParenthesesExpression) GetDescriptor() string       { return "" }
func (e *ParenthesesExpression) GetDoubleValue() float64     { return 0 }
func (e *ParenthesesExpression) GetFloatValue() float32      { return 0 }
func (e *ParenthesesExpression) GetIntegerValue() int        { return 0 }
func (e *ParenthesesExpression) GetInternalTypeName() string { return "" }
func (e *ParenthesesExpression) GetLongValue() int64         { return 0 }
func (e *ParenthesesExpression) GetName() string             { return "" }
func (e *ParenthesesExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *ParenthesesExpression) GetOperator() string         { return "" }
func (e *ParenthesesExpression) GetStringValue() string      { return "" }

func (e *ParenthesesExpression) String() string { return "" }

type PostOperatorExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Priority   int
	Operator   string
	Expression IExpression
}

func (e *PostOperatorExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitPostOperatorExpression(e)
}

func (e *PostOperatorExpression) IsArrayExpression() bool                      { return false }
func (e *PostOperatorExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *PostOperatorExpression) IsBooleanExpression() bool                    { return false }
func (e *PostOperatorExpression) IsCastExpression() bool                       { return false }
func (e *PostOperatorExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *PostOperatorExpression) IsDoubleConstantExpression() bool             { return false }
func (e *PostOperatorExpression) IsFieldReferenceExpression() bool             { return false }
func (e *PostOperatorExpression) IsFloatConstantExpression() bool              { return false }
func (e *PostOperatorExpression) IsIntegerConstantExpression() bool            { return false }
func (e *PostOperatorExpression) IsLengthExpression() bool                     { return false }
func (e *PostOperatorExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *PostOperatorExpression) IsLongConstantExpression() bool               { return false }
func (e *PostOperatorExpression) IsMethodInvocationExpression() bool           { return false }
func (e *PostOperatorExpression) IsNewArray() bool                             { return false }
func (e *PostOperatorExpression) IsNewExpression() bool                        { return false }
func (e *PostOperatorExpression) IsNewInitializedArray() bool                  { return false }
func (e *PostOperatorExpression) IsNullExpression() bool                       { return false }
func (e *PostOperatorExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *PostOperatorExpression) IsPostOperatorExpression() bool               { return true }
func (e *PostOperatorExpression) IsPreOperatorExpression() bool                { return false }
func (e *PostOperatorExpression) IsStringConstantExpression() bool             { return false }
func (e *PostOperatorExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *PostOperatorExpression) IsSuperExpression() bool                      { return false }
func (e *PostOperatorExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *PostOperatorExpression) IsThisExpression() bool                       { return false }

func (e *PostOperatorExpression) GetLineNumber() int { return e.LineNumber }
func (e *PostOperatorExpression) GetType() IType     { return e.Expression.GetType() }
func (e *PostOperatorExpression) GetPriority() int   { return e.Priority }

func (e *PostOperatorExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *PostOperatorExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *PostOperatorExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *PostOperatorExpression) GetExpression() IExpression      { return e.Expression }
func (e *PostOperatorExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *PostOperatorExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *PostOperatorExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *PostOperatorExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *PostOperatorExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *PostOperatorExpression) GetDescriptor() string       { return "" }
func (e *PostOperatorExpression) GetDoubleValue() float64     { return 0 }
func (e *PostOperatorExpression) GetFloatValue() float32      { return 0 }
func (e *PostOperatorExpression) GetIntegerValue() int        { return 0 }
func (e *PostOperatorExpression) GetInternalTypeName() string { return "" }
func (e *PostOperatorExpression) GetLongValue() int64         { return 0 }
func (e *PostOperatorExpression) GetName() string             { return "" }
func (e *PostOperatorExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *PostOperatorExpression) GetOperator() string         { return e.Operator }
func (e *PostOperatorExpression) GetStringValue() string      { return "" }

func (e *PostOperatorExpression) String() string {
	return fmt.Sprintf("PostOperatorExpression{%s %s}", e.Expression, e.Operator)
}

type PreOperatorExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Priority   int
	Operator   string
	Expression IExpression
}

func (e *PreOperatorExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitPreOperatorExpression(e)
}

func (e *PreOperatorExpression) IsArrayExpression() bool                      { return false }
func (e *PreOperatorExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *PreOperatorExpression) IsBooleanExpression() bool                    { return false }
func (e *PreOperatorExpression) IsCastExpression() bool                       { return false }
func (e *PreOperatorExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *PreOperatorExpression) IsDoubleConstantExpression() bool             { return false }
func (e *PreOperatorExpression) IsFieldReferenceExpression() bool             { return false }
func (e *PreOperatorExpression) IsFloatConstantExpression() bool              { return false }
func (e *PreOperatorExpression) IsIntegerConstantExpression() bool            { return false }
func (e *PreOperatorExpression) IsLengthExpression() bool                     { return false }
func (e *PreOperatorExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *PreOperatorExpression) IsLongConstantExpression() bool               { return false }
func (e *PreOperatorExpression) IsMethodInvocationExpression() bool           { return false }
func (e *PreOperatorExpression) IsNewArray() bool                             { return false }
func (e *PreOperatorExpression) IsNewExpression() bool                        { return false }
func (e *PreOperatorExpression) IsNewInitializedArray() bool                  { return false }
func (e *PreOperatorExpression) IsNullExpression() bool                       { return false }
func (e *PreOperatorExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *PreOperatorExpression) IsPostOperatorExpression() bool               { return false }
func (e *PreOperatorExpression) IsPreOperatorExpression() bool                { return true }
func (e *PreOperatorExpression) IsStringConstantExpression() bool             { return false }
func (e *PreOperatorExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *PreOperatorExpression) IsSuperExpression() bool                      { return false }
func (e *PreOperatorExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *PreOperatorExpression) IsThisExpression() bool                       { return false }

func (e *PreOperatorExpression) GetLineNumber() int { return e.LineNumber }
func (e *PreOperatorExpression) GetType() IType     { return e.Expression.GetType() }
func (e *PreOperatorExpression) GetPriority() int   { return e.Priority }

func (e *PreOperatorExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *PreOperatorExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *PreOperatorExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *PreOperatorExpression) GetExpression() IExpression      { return e.Expression }
func (e *PreOperatorExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *PreOperatorExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *PreOperatorExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *PreOperatorExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *PreOperatorExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *PreOperatorExpression) GetDescriptor() string       { return "" }
func (e *PreOperatorExpression) GetDoubleValue() float64     { return 0 }
func (e *PreOperatorExpression) GetFloatValue() float32      { return 0 }
func (e *PreOperatorExpression) GetIntegerValue() int        { return 0 }
func (e *PreOperatorExpression) GetInternalTypeName() string { return "" }
func (e *PreOperatorExpression) GetLongValue() int64         { return 0 }
func (e *PreOperatorExpression) GetName() string             { return "" }
func (e *PreOperatorExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *PreOperatorExpression) GetOperator() string         { return e.Operator }
func (e *PreOperatorExpression) GetStringValue() string      { return "" }

func (e *PreOperatorExpression) String() string {
	return fmt.Sprintf("PreOperatorExpression{%s %s}", e.Operator, e.Expression)
}

type StringConstantExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Priority   int
	Text       string
}

func (e *StringConstantExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitStringConstantExpression(e)
}

func (e *StringConstantExpression) IsArrayExpression() bool                      { return false }
func (e *StringConstantExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *StringConstantExpression) IsBooleanExpression() bool                    { return false }
func (e *StringConstantExpression) IsCastExpression() bool                       { return false }
func (e *StringConstantExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *StringConstantExpression) IsDoubleConstantExpression() bool             { return false }
func (e *StringConstantExpression) IsFieldReferenceExpression() bool             { return false }
func (e *StringConstantExpression) IsFloatConstantExpression() bool              { return false }
func (e *StringConstantExpression) IsIntegerConstantExpression() bool            { return false }
func (e *StringConstantExpression) IsLengthExpression() bool                     { return false }
func (e *StringConstantExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *StringConstantExpression) IsLongConstantExpression() bool               { return false }
func (e *StringConstantExpression) IsMethodInvocationExpression() bool           { return false }
func (e *StringConstantExpression) IsNewArray() bool                             { return false }
func (e *StringConstantExpression) IsNewExpression() bool                        { return false }
func (e *StringConstantExpression) IsNewInitializedArray() bool                  { return false }
func (e *StringConstantExpression) IsNullExpression() bool                       { return false }
func (e *StringConstantExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *StringConstantExpression) IsPostOperatorExpression() bool               { return false }
func (e *StringConstantExpression) IsPreOperatorExpression() bool                { return false }
func (e *StringConstantExpression) IsStringConstantExpression() bool             { return true }
func (e *StringConstantExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *StringConstantExpression) IsSuperExpression() bool                      { return false }
func (e *StringConstantExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *StringConstantExpression) IsThisExpression() bool                       { return false }

func (e *StringConstantExpression) GetLineNumber() int { return e.LineNumber }
func (e *StringConstantExpression) GetType() IType     { return &OtTypeString }
func (e *StringConstantExpression) GetPriority() int   { return e.Priority }

func (e *StringConstantExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *StringConstantExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *StringConstantExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *StringConstantExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *StringConstantExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *StringConstantExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *StringConstantExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *StringConstantExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *StringConstantExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *StringConstantExpression) GetDescriptor() string       { return "" }
func (e *StringConstantExpression) GetDoubleValue() float64     { return 0 }
func (e *StringConstantExpression) GetFloatValue() float32      { return 0 }
func (e *StringConstantExpression) GetIntegerValue() int        { return 0 }
func (e *StringConstantExpression) GetInternalTypeName() string { return "" }
func (e *StringConstantExpression) GetLongValue() int64         { return 0 }
func (e *StringConstantExpression) GetName() string             { return "" }
func (e *StringConstantExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *StringConstantExpression) GetOperator() string         { return "" }
func (e *StringConstantExpression) GetStringValue() string      { return e.Text }

func (e *StringConstantExpression) String() string {
	return fmt.Sprintf("StringConstantExpression{\"%s\"}", e.Text)
}

type SuperConstructorInvocationExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       IType
	Priority   int
	ObjectType *ObjectType
	Descriptor string
	Parameters IExpression
}

func (e *SuperConstructorInvocationExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitSuperConstructorInvocationExpression(e)
}

func (e *SuperConstructorInvocationExpression) IsArrayExpression() bool                 { return false }
func (e *SuperConstructorInvocationExpression) IsBinaryOperatorExpression() bool        { return false }
func (e *SuperConstructorInvocationExpression) IsBooleanExpression() bool               { return false }
func (e *SuperConstructorInvocationExpression) IsCastExpression() bool                  { return false }
func (e *SuperConstructorInvocationExpression) IsConstructorInvocationExpression() bool { return false }
func (e *SuperConstructorInvocationExpression) IsDoubleConstantExpression() bool        { return false }
func (e *SuperConstructorInvocationExpression) IsFieldReferenceExpression() bool        { return false }
func (e *SuperConstructorInvocationExpression) IsFloatConstantExpression() bool         { return false }
func (e *SuperConstructorInvocationExpression) IsIntegerConstantExpression() bool       { return false }
func (e *SuperConstructorInvocationExpression) IsLengthExpression() bool                { return false }
func (e *SuperConstructorInvocationExpression) IsLocalVariableReferenceExpression() bool {
	return false
}
func (e *SuperConstructorInvocationExpression) IsLongConstantExpression() bool        { return false }
func (e *SuperConstructorInvocationExpression) IsMethodInvocationExpression() bool    { return false }
func (e *SuperConstructorInvocationExpression) IsNewArray() bool                      { return false }
func (e *SuperConstructorInvocationExpression) IsNewExpression() bool                 { return false }
func (e *SuperConstructorInvocationExpression) IsNewInitializedArray() bool           { return false }
func (e *SuperConstructorInvocationExpression) IsNullExpression() bool                { return false }
func (e *SuperConstructorInvocationExpression) IsObjectTypeReferenceExpression() bool { return false }
func (e *SuperConstructorInvocationExpression) IsPostOperatorExpression() bool        { return false }
func (e *SuperConstructorInvocationExpression) IsPreOperatorExpression() bool         { return false }
func (e *SuperConstructorInvocationExpression) IsStringConstantExpression() bool      { return false }
func (e *SuperConstructorInvocationExpression) IsSuperConstructorInvocationExpression() bool {
	return true
}
func (e *SuperConstructorInvocationExpression) IsSuperExpression() bool           { return false }
func (e *SuperConstructorInvocationExpression) IsTernaryOperatorExpression() bool { return false }
func (e *SuperConstructorInvocationExpression) IsThisExpression() bool            { return false }

func (e *SuperConstructorInvocationExpression) GetLineNumber() int { return e.LineNumber }
func (e *SuperConstructorInvocationExpression) GetType() IType     { return e.Type }
func (e *SuperConstructorInvocationExpression) GetPriority() int   { return e.Priority }

func (e *SuperConstructorInvocationExpression) GetDimensionExpressionList() IExpression {
	return &NeNoExpression
}
func (e *SuperConstructorInvocationExpression) GetParameters() IExpression { return e.Parameters }

func (e *SuperConstructorInvocationExpression) GetCondition() IExpression  { return &NeNoExpression }
func (e *SuperConstructorInvocationExpression) GetExpression() IExpression { return &NeNoExpression }
func (e *SuperConstructorInvocationExpression) GetTrueExpression() IExpression {
	return &NeNoExpression
}
func (e *SuperConstructorInvocationExpression) GetFalseExpression() IExpression {
	return &NeNoExpression
}
func (e *SuperConstructorInvocationExpression) GetIndex() IExpression { return &NeNoExpression }
func (e *SuperConstructorInvocationExpression) GetLeftExpression() IExpression {
	return &NeNoExpression
}
func (e *SuperConstructorInvocationExpression) GetRightExpression() IExpression {
	return &NeNoExpression
}

func (e *SuperConstructorInvocationExpression) GetDescriptor() string       { return e.Descriptor }
func (e *SuperConstructorInvocationExpression) GetDoubleValue() float64     { return 0 }
func (e *SuperConstructorInvocationExpression) GetFloatValue() float32      { return 0 }
func (e *SuperConstructorInvocationExpression) GetIntegerValue() int        { return 0 }
func (e *SuperConstructorInvocationExpression) GetInternalTypeName() string { return "" }
func (e *SuperConstructorInvocationExpression) GetLongValue() int64         { return 0 }
func (e *SuperConstructorInvocationExpression) GetName() string             { return "" }
func (e *SuperConstructorInvocationExpression) GetObjectType() *ObjectType  { return e.ObjectType }
func (e *SuperConstructorInvocationExpression) GetOperator() string         { return "" }
func (e *SuperConstructorInvocationExpression) GetStringValue() string      { return "" }

func (e *SuperConstructorInvocationExpression) String() string {
	return fmt.Sprintf("SuperConstructorInvocationExpression{call super(%s)}", e.Descriptor)
}

func NewSuperExpressionWithAll(lineNumber int, typ IType) SuperExpression {
	e := SuperExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Priority:    0,
		Type:        typ,
	}
	e.SetValue(&e)
	return e
}

type SuperExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Priority   int
	Type       IType
}

func (e *SuperExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitSuperExpression(e)
}

func (e *SuperExpression) IsArrayExpression() bool                      { return false }
func (e *SuperExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *SuperExpression) IsBooleanExpression() bool                    { return false }
func (e *SuperExpression) IsCastExpression() bool                       { return false }
func (e *SuperExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *SuperExpression) IsDoubleConstantExpression() bool             { return false }
func (e *SuperExpression) IsFieldReferenceExpression() bool             { return false }
func (e *SuperExpression) IsFloatConstantExpression() bool              { return false }
func (e *SuperExpression) IsIntegerConstantExpression() bool            { return false }
func (e *SuperExpression) IsLengthExpression() bool                     { return false }
func (e *SuperExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *SuperExpression) IsLongConstantExpression() bool               { return false }
func (e *SuperExpression) IsMethodInvocationExpression() bool           { return false }
func (e *SuperExpression) IsNewArray() bool                             { return false }
func (e *SuperExpression) IsNewExpression() bool                        { return false }
func (e *SuperExpression) IsNewInitializedArray() bool                  { return false }
func (e *SuperExpression) IsNullExpression() bool                       { return false }
func (e *SuperExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *SuperExpression) IsPostOperatorExpression() bool               { return false }
func (e *SuperExpression) IsPreOperatorExpression() bool                { return false }
func (e *SuperExpression) IsStringConstantExpression() bool             { return false }
func (e *SuperExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *SuperExpression) IsSuperExpression() bool                      { return true }
func (e *SuperExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *SuperExpression) IsThisExpression() bool                       { return false }

func (e *SuperExpression) GetLineNumber() int { return e.LineNumber }
func (e *SuperExpression) GetType() IType     { return e.Type }
func (e *SuperExpression) GetPriority() int   { return e.Priority }

func (e *SuperExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *SuperExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *SuperExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *SuperExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *SuperExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *SuperExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *SuperExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *SuperExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *SuperExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *SuperExpression) GetDescriptor() string       { return "" }
func (e *SuperExpression) GetDoubleValue() float64     { return 0 }
func (e *SuperExpression) GetFloatValue() float32      { return 0 }
func (e *SuperExpression) GetIntegerValue() int        { return 0 }
func (e *SuperExpression) GetInternalTypeName() string { return "" }
func (e *SuperExpression) GetLongValue() int64         { return 0 }
func (e *SuperExpression) GetName() string             { return "" }
func (e *SuperExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *SuperExpression) GetOperator() string         { return "" }
func (e *SuperExpression) GetStringValue() string      { return "" }

func (e *SuperExpression) String() string {
	return fmt.Sprintf("SuperExpression{%s}", e.Type)
}

type TernaryOperatorExpression struct {
	util.DefaultBase[IExpression]

	LineNumber      int
	Type            IType
	Priority        int
	Condition       IExpression
	TrueExpression  IExpression
	FalseExpression IExpression
}

func (e *TernaryOperatorExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitTernaryOperatorExpression(e)
}

func (e *TernaryOperatorExpression) IsArrayExpression() bool                      { return false }
func (e *TernaryOperatorExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *TernaryOperatorExpression) IsBooleanExpression() bool                    { return false }
func (e *TernaryOperatorExpression) IsCastExpression() bool                       { return false }
func (e *TernaryOperatorExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *TernaryOperatorExpression) IsDoubleConstantExpression() bool             { return false }
func (e *TernaryOperatorExpression) IsFieldReferenceExpression() bool             { return false }
func (e *TernaryOperatorExpression) IsFloatConstantExpression() bool              { return false }
func (e *TernaryOperatorExpression) IsIntegerConstantExpression() bool            { return false }
func (e *TernaryOperatorExpression) IsLengthExpression() bool                     { return false }
func (e *TernaryOperatorExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *TernaryOperatorExpression) IsLongConstantExpression() bool               { return false }
func (e *TernaryOperatorExpression) IsMethodInvocationExpression() bool           { return false }
func (e *TernaryOperatorExpression) IsNewArray() bool                             { return false }
func (e *TernaryOperatorExpression) IsNewExpression() bool                        { return false }
func (e *TernaryOperatorExpression) IsNewInitializedArray() bool                  { return false }
func (e *TernaryOperatorExpression) IsNullExpression() bool                       { return false }
func (e *TernaryOperatorExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *TernaryOperatorExpression) IsPostOperatorExpression() bool               { return false }
func (e *TernaryOperatorExpression) IsPreOperatorExpression() bool                { return false }
func (e *TernaryOperatorExpression) IsStringConstantExpression() bool             { return false }
func (e *TernaryOperatorExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *TernaryOperatorExpression) IsSuperExpression() bool                      { return false }
func (e *TernaryOperatorExpression) IsTernaryOperatorExpression() bool            { return true }
func (e *TernaryOperatorExpression) IsThisExpression() bool                       { return false }

func (e *TernaryOperatorExpression) GetLineNumber() int { return e.LineNumber }
func (e *TernaryOperatorExpression) GetType() IType     { return e.Type }
func (e *TernaryOperatorExpression) GetPriority() int   { return e.Priority }

func (e *TernaryOperatorExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *TernaryOperatorExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *TernaryOperatorExpression) GetCondition() IExpression       { return e.Condition }
func (e *TernaryOperatorExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *TernaryOperatorExpression) GetTrueExpression() IExpression  { return e.TrueExpression }
func (e *TernaryOperatorExpression) GetFalseExpression() IExpression { return e.FalseExpression }
func (e *TernaryOperatorExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *TernaryOperatorExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *TernaryOperatorExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *TernaryOperatorExpression) GetDescriptor() string       { return "" }
func (e *TernaryOperatorExpression) GetDoubleValue() float64     { return 0 }
func (e *TernaryOperatorExpression) GetFloatValue() float32      { return 0 }
func (e *TernaryOperatorExpression) GetIntegerValue() int        { return 0 }
func (e *TernaryOperatorExpression) GetInternalTypeName() string { return "" }
func (e *TernaryOperatorExpression) GetLongValue() int64         { return 0 }
func (e *TernaryOperatorExpression) GetName() string             { return "" }
func (e *TernaryOperatorExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *TernaryOperatorExpression) GetOperator() string         { return "" }
func (e *TernaryOperatorExpression) GetStringValue() string      { return "" }

func (e *TernaryOperatorExpression) String() string {
	return fmt.Sprintf("TernaryOperatorExpression{%s ? %s : %s}", e.Condition, e.TrueExpression, e.FalseExpression)
}

type ThisExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Priority   int
	Type       IType
	IsExplicit bool
}

func (e *ThisExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitThisExpression(e)
}

func (e *ThisExpression) IsArrayExpression() bool                      { return false }
func (e *ThisExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *ThisExpression) IsBooleanExpression() bool                    { return false }
func (e *ThisExpression) IsCastExpression() bool                       { return false }
func (e *ThisExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *ThisExpression) IsDoubleConstantExpression() bool             { return false }
func (e *ThisExpression) IsFieldReferenceExpression() bool             { return false }
func (e *ThisExpression) IsFloatConstantExpression() bool              { return false }
func (e *ThisExpression) IsIntegerConstantExpression() bool            { return false }
func (e *ThisExpression) IsLengthExpression() bool                     { return false }
func (e *ThisExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *ThisExpression) IsLongConstantExpression() bool               { return false }
func (e *ThisExpression) IsMethodInvocationExpression() bool           { return false }
func (e *ThisExpression) IsNewArray() bool                             { return false }
func (e *ThisExpression) IsNewExpression() bool                        { return false }
func (e *ThisExpression) IsNewInitializedArray() bool                  { return false }
func (e *ThisExpression) IsNullExpression() bool                       { return false }
func (e *ThisExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *ThisExpression) IsPostOperatorExpression() bool               { return false }
func (e *ThisExpression) IsPreOperatorExpression() bool                { return false }
func (e *ThisExpression) IsStringConstantExpression() bool             { return false }
func (e *ThisExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *ThisExpression) IsSuperExpression() bool                      { return false }
func (e *ThisExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *ThisExpression) IsThisExpression() bool                       { return true }

func (e *ThisExpression) GetLineNumber() int { return e.LineNumber }
func (e *ThisExpression) GetType() IType     { return e.Type }
func (e *ThisExpression) GetPriority() int   { return e.Priority }

func (e *ThisExpression) GetDimensionExpressionList() IExpression { return &NeNoExpression }
func (e *ThisExpression) GetParameters() IExpression              { return &NeNoExpression }

func (e *ThisExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *ThisExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *ThisExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *ThisExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *ThisExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *ThisExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *ThisExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *ThisExpression) GetDescriptor() string       { return "" }
func (e *ThisExpression) GetDoubleValue() float64     { return 0 }
func (e *ThisExpression) GetFloatValue() float32      { return 0 }
func (e *ThisExpression) GetIntegerValue() int        { return 0 }
func (e *ThisExpression) GetInternalTypeName() string { return "" }
func (e *ThisExpression) GetLongValue() int64         { return 0 }
func (e *ThisExpression) GetName() string             { return "" }
func (e *ThisExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *ThisExpression) GetOperator() string         { return "" }
func (e *ThisExpression) GetStringValue() string      { return "" }

func (e *ThisExpression) String() string {
	return fmt.Sprintf("ThisExpression{%s}", e.Type)
}

type TypeReferenceDotClassExpression struct {
	util.DefaultBase[IExpression]
	LineNumber   int
	Type         IType
	Priority     int
	TypeDotClass IType
}

func (e *TypeReferenceDotClassExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitTypeReferenceDotClassExpression(e)
}

func (e *TypeReferenceDotClassExpression) IsArrayExpression() bool                      { return false }
func (e *TypeReferenceDotClassExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *TypeReferenceDotClassExpression) IsBooleanExpression() bool                    { return false }
func (e *TypeReferenceDotClassExpression) IsCastExpression() bool                       { return false }
func (e *TypeReferenceDotClassExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *TypeReferenceDotClassExpression) IsDoubleConstantExpression() bool             { return false }
func (e *TypeReferenceDotClassExpression) IsFieldReferenceExpression() bool             { return false }
func (e *TypeReferenceDotClassExpression) IsFloatConstantExpression() bool              { return false }
func (e *TypeReferenceDotClassExpression) IsIntegerConstantExpression() bool            { return false }
func (e *TypeReferenceDotClassExpression) IsLengthExpression() bool                     { return false }
func (e *TypeReferenceDotClassExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *TypeReferenceDotClassExpression) IsLongConstantExpression() bool               { return false }
func (e *TypeReferenceDotClassExpression) IsMethodInvocationExpression() bool           { return false }
func (e *TypeReferenceDotClassExpression) IsNewArray() bool                             { return false }
func (e *TypeReferenceDotClassExpression) IsNewExpression() bool                        { return false }
func (e *TypeReferenceDotClassExpression) IsNewInitializedArray() bool                  { return false }
func (e *TypeReferenceDotClassExpression) IsNullExpression() bool                       { return false }
func (e *TypeReferenceDotClassExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *TypeReferenceDotClassExpression) IsPostOperatorExpression() bool               { return false }
func (e *TypeReferenceDotClassExpression) IsPreOperatorExpression() bool                { return false }
func (e *TypeReferenceDotClassExpression) IsStringConstantExpression() bool             { return false }
func (e *TypeReferenceDotClassExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *TypeReferenceDotClassExpression) IsSuperExpression() bool                      { return false }
func (e *TypeReferenceDotClassExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *TypeReferenceDotClassExpression) IsThisExpression() bool                       { return false }

func (e *TypeReferenceDotClassExpression) GetLineNumber() int { return e.LineNumber }
func (e *TypeReferenceDotClassExpression) GetType() IType     { return e.Type }
func (e *TypeReferenceDotClassExpression) GetPriority() int   { return e.Priority }

func (e *TypeReferenceDotClassExpression) GetDimensionExpressionList() IExpression {
	return &NeNoExpression
}
func (e *TypeReferenceDotClassExpression) GetParameters() IExpression { return &NeNoExpression }

func (e *TypeReferenceDotClassExpression) GetCondition() IExpression       { return &NeNoExpression }
func (e *TypeReferenceDotClassExpression) GetExpression() IExpression      { return &NeNoExpression }
func (e *TypeReferenceDotClassExpression) GetTrueExpression() IExpression  { return &NeNoExpression }
func (e *TypeReferenceDotClassExpression) GetFalseExpression() IExpression { return &NeNoExpression }
func (e *TypeReferenceDotClassExpression) GetIndex() IExpression           { return &NeNoExpression }
func (e *TypeReferenceDotClassExpression) GetLeftExpression() IExpression  { return &NeNoExpression }
func (e *TypeReferenceDotClassExpression) GetRightExpression() IExpression { return &NeNoExpression }

func (e *TypeReferenceDotClassExpression) GetDescriptor() string       { return "" }
func (e *TypeReferenceDotClassExpression) GetDoubleValue() float64     { return 0 }
func (e *TypeReferenceDotClassExpression) GetFloatValue() float32      { return 0 }
func (e *TypeReferenceDotClassExpression) GetIntegerValue() int        { return 0 }
func (e *TypeReferenceDotClassExpression) GetInternalTypeName() string { return "" }
func (e *TypeReferenceDotClassExpression) GetLongValue() int64         { return 0 }
func (e *TypeReferenceDotClassExpression) GetName() string             { return "" }
func (e *TypeReferenceDotClassExpression) GetObjectType() *ObjectType  { return &OtTypeUndefinedObject }
func (e *TypeReferenceDotClassExpression) GetOperator() string         { return "" }
func (e *TypeReferenceDotClassExpression) GetStringValue() string      { return "" }

func (e *TypeReferenceDotClassExpression) String() string {
	return fmt.Sprintf("TypeReferenceDotClassExpression{%s}", e.TypeDotClass)
}

/////////////////////////////////////////////////////////////////////////
//  Functions
/////////////////////////////////////////////////////////////////////////

func CreateItemType(expression IExpression) IType {
	typ := expression.GetType()
	dimension := typ.GetDimension()

	if dimension > 0 {
		return typ.CreateType(dimension - 1)
	}

	return typ.CreateType(0)
}

func GetPrimitiveTypeFromValue(value int) PrimitiveType {
	if value >= 0 {
		if value <= 1 {
			return PtMaybeBooleanType
		}
		if value <= math.MaxInt8 {
			return PtMaybeByteType
		}
		if value <= math.MaxInt16 {
			return PtMaybeShortType
		}
		if value <= math.MaxUint16 {
			return PtMaybeCharType
		}
	} else {
		if value >= math.MinInt8 {
			return PtMaybeNegativeByteType
		}
		if value <= math.MinInt16 {
			return PtMaybeNegativeShortType
		}
	}
	return PtMaybeIntType
}
