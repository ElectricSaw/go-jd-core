package expression

import (
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
	"math"
)

func NewIntegerConstantExpression(typ _type.IType, value int) *IntegerConstantExpression {
	return &IntegerConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		value:                            value,
	}
}

func NewIntegerConstantExpressionWithAll(lineNumber int, typ _type.IType, value int) *IntegerConstantExpression {
	return &IntegerConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		value:                            value,
	}
}

type IntegerConstantExpression struct {
	AbstractLineNumberTypeExpression

	value int
}

func (e *IntegerConstantExpression) GetIntegerValue() int {
	return e.value
}

func (e *IntegerConstantExpression) SetType(typ _type.IType) {
	e.checkType(typ)
	e.AbstractLineNumberTypeExpression.SetType(typ)
}

func (e *IntegerConstantExpression) checkType(typ _type.IType) bool {
	if typ.IsPrimitiveType() {
		valueType := GetPrimitiveTypeFromValue(e.value)
		pt, ok := e.typ.(*_type.PrimitiveType)
		if ok {
			return pt.GetFlags()&valueType.GetFlags() != 0
		}
	}
	return false
}

func (e *IntegerConstantExpression) IsIntegerConstantExpression() bool {
	return true
}

func (e *IntegerConstantExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitIntegerConstantExpression(e)
}

func (e *IntegerConstantExpression) String() string {
	return fmt.Sprintf("IntegerConstantExpression{type=%s, value=%d}", e.typ, e.value)
}

func GetPrimitiveTypeFromValue(value int) *_type.PrimitiveType {
	if value >= 0 {
		if value <= 1 {
			return _type.PtMaybeBooleanType
		}
		if value <= math.MaxInt8 {
			return _type.PtMaybeByteType
		}
		if value <= math.MaxInt16 {
			return _type.PtMaybeShortType
		}
		if value <= math.MaxUint16 {
			return _type.PtMaybeCharType
		}
	} else {
		if value >= math.MinInt8 {
			return _type.PtMaybeNegativeByteType
		}
		if value <= math.MinInt16 {
			return _type.PtMaybeNegativeShortType
		}
	}
	return _type.PtMaybeIntType
}
