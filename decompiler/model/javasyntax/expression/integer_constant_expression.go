package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
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
		pt, ok := e.typ.(*model.PrimitiveType)
		if ok {
			return pt.Flags()&valueType.Flags() != 0
		}
	}
	return false
}

func GetPrimitiveTypeFromValue(value int) intmod.IPrimitiveType {
	if value >= 0 {
		if value <= 1 {
			return model.PtMaybeBooleanType
		}
		if value <= math.MaxInt8 {
			return model.PtMaybeByteType
		}
		if value <= math.MaxInt16 {
			return model.PtMaybeShortType
		}
		if value <= math.MaxUint16 {
			return model.PtMaybeCharType
		}
	} else {
		if value >= math.MinInt8 {
			return model.PtMaybeNegativeByteType
		}
		if value <= math.MinInt16 {
			return model.PtMaybeNegativeShortType
		}
	}
	return model.PtMaybeIntType
}
