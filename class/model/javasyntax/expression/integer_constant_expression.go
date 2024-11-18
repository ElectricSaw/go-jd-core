package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
	"math"
)

func NewIntegerConstantExpression(typ intsyn.IType, value int) intsyn.IIntegerConstantExpression {
	return &IntegerConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		value:                            value,
	}
}

func NewIntegerConstantExpressionWithAll(lineNumber int, typ intsyn.IType, value int) intsyn.IIntegerConstantExpression {
	return &IntegerConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		value:                            value,
	}
}

type IntegerConstantExpression struct {
	AbstractLineNumberTypeExpression

	value int
}

func (e *IntegerConstantExpression) IntegerValue() int {
	return e.value
}

func (e *IntegerConstantExpression) SetType(typ intsyn.IType) {
	e.checkType(typ)
	e.AbstractLineNumberTypeExpression.SetType(typ)
}

func (e *IntegerConstantExpression) IsIntegerConstantExpression() bool {
	return true
}

func (e *IntegerConstantExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitIntegerConstantExpression(e)
}

func (e *IntegerConstantExpression) String() string {
	return fmt.Sprintf("IntegerConstantExpression{type=%s, value=%d}", e.typ, e.value)
}

func (e *IntegerConstantExpression) checkType(typ intsyn.IType) bool {
	if typ.IsPrimitiveType() {
		valueType := GetPrimitiveTypeFromValue(e.value)
		pt, ok := e.typ.(*_type.PrimitiveType)
		if ok {
			return pt.Flags()&valueType.Flags() != 0
		}
	}
	return false
}

func GetPrimitiveTypeFromValue(value int) intsyn.IPrimitiveType {
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
