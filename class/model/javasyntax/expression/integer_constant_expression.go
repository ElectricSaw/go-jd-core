package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	"fmt"
	"math"
)

func NewIntegerConstantExpression(typ intmod.IType, value int) intmod.IIntegerConstantExpression {
	return NewIntegerConstantExpressionWithAll(0, typ, value)
}

func NewIntegerConstantExpressionWithAll(lineNumber int, typ intmod.IType, value int) intmod.IIntegerConstantExpression {
	e := &IntegerConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		value:                            value,
	}
	e.SetValue(e)
	return e
}

type IntegerConstantExpression struct {
	AbstractLineNumberTypeExpression

	value int
}

func (e *IntegerConstantExpression) IntegerValue() int {
	return e.value
}

func (e *IntegerConstantExpression) SetType(typ intmod.IType) {
	e.checkType(typ)
	e.AbstractLineNumberTypeExpression.SetType(typ)
}

func (e *IntegerConstantExpression) IsIntegerConstantExpression() bool {
	return true
}

func (e *IntegerConstantExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitIntegerConstantExpression(e)
}

func (e *IntegerConstantExpression) String() string {
	return fmt.Sprintf("IntegerConstantExpression{type=%s, value=%d}", e.typ, e.value)
}

func (e *IntegerConstantExpression) checkType(typ intmod.IType) bool {
	if typ.IsPrimitiveType() {
		valueType := GetPrimitiveTypeFromValue(e.value)
		pt, ok := e.typ.(*_type.PrimitiveType)
		if ok {
			return pt.Flags()&valueType.Flags() != 0
		}
	}
	return false
}

func GetPrimitiveTypeFromValue(value int) intmod.IPrimitiveType {
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
