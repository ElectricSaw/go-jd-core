package utils

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/expression"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	srvexp "github.com/ElectricSaw/go-jd-core/class/service/converter/model/javasyntax/expression"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func StringConcatenationUtilCreate1(expr intmod.IExpression, lineNumber int, typeName string) intmod.IExpression {
	if expr.IsMethodInvocationExpression() {
		mie := expr.(intmod.IMethodInvocationExpression)

		if (mie.Parameters() != nil) && !mie.Parameters().IsList() && "append" == mie.Name() {
			concatenatedStringExpression := mie.Parameters().First()
			exp := mie.Expression()
			firstParameterHaveGenericType := false

			for exp.IsMethodInvocationExpression() {
				mie = exp.(intmod.IMethodInvocationExpression)

				if mie.Parameters() == nil || mie.Parameters().IsList() || "append" != mie.Name() {
					break
				}

				firstParameterHaveGenericType = mie.Parameters().First().Type().IsGenericType()
				concatenatedStringExpression = expression.NewBinaryOperatorExpression(mie.LineNumber(),
					_type.OtTypeString, mie.Parameters(), "+", concatenatedStringExpression, 4)
				exp = mie.Expression()
			}

			if exp.IsNewExpression() {
				internalTypeName := exp.Type().Descriptor()

				if "Ljava/lang/StringBuilder;" == internalTypeName || "Ljava/lang/StringBuffer;" == internalTypeName {
					if exp.Parameters() == nil {
						if !firstParameterHaveGenericType {
							return concatenatedStringExpression
						}
					} else if !exp.Parameters().IsList() {
						exp = exp.Parameters().First()

						if _type.OtTypeString == exp.Type() {
							return expression.NewBinaryOperatorExpression(exp.LineNumber(), _type.OtTypeString,
								exp, "+", concatenatedStringExpression, 4)
						}
					}
				}
			}
		}
	}

	return srvexp.NewClassFileMethodInvocationExpression(lineNumber, nil, _type.OtTypeString,
		expr, typeName, "toString", "()Ljava/lang/String;", nil, nil)
}

func StringConcatenationUtilCreate2(recipe string, parameters intmod.IExpression) intmod.IExpression {
	st := util.NewStringTokenizer3(recipe, "\u0001", true)

	if st.HasMoreTokens() {
		token, _ := st.NextToken()
		var expr intmod.IExpression

		if token == "\u0001" {
			expr = createFirstStringConcatenationItem(parameters.First())
		} else {
			expr = expression.NewStringConstantExpression(token)
		}

		if parameters.IsList() {
			list := parameters.ToList()
			index := 0

			for st.HasMoreTokens() {
				token, _ = st.NextToken()
				var e intmod.IExpression

				if token == "\u0001" {
					expr = list.Get(index)
					index++
				} else {
					expr = expression.NewStringConstantExpression(token)
				}

				expr = expression.NewBinaryOperatorExpression(expr.LineNumber(), _type.OtTypeString, expr, "+", e, 6)
			}
		} else {
			for st.HasMoreTokens() {
				token, _ = st.NextToken()
				var e intmod.IExpression
				if token == "\u0001" {
					expr = parameters.First()
				} else {
					expr = expression.NewStringConstantExpression(token)
				}
				expr = expression.NewBinaryOperatorExpression(expr.LineNumber(), _type.OtTypeString, expr, "+", e, 6)
			}
		}

		return expr
	} else {
		return expression.EmptyString
	}
}

func StringConcatenationUtilCreate3(parameters intmod.IExpression) intmod.IExpression {
	switch parameters.Size() {
	case 0:
		return expression.EmptyString
	case 1:
		return createFirstStringConcatenationItem(parameters.First())
	default:
		iterator := parameters.Iterator()
		expr := createFirstStringConcatenationItem(iterator.Next())
		for iterator.HasNext() {
			expr = expression.NewBinaryOperatorExpression(expr.LineNumber(), _type.OtTypeString, expr, "+", iterator.Next(), 6)
		}
		return expr
	}
}

func createFirstStringConcatenationItem(expr intmod.IExpression) intmod.IExpression {
	if expr.Type() != _type.OtTypeString {
		expr = expression.NewBinaryOperatorExpression(expr.LineNumber(),
			_type.OtTypeString, expression.EmptyString, "+", expr, 6)
	}

	return expr
}
