package utils

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"strings"
)

var EmptyArray = declaration.NewArrayVariableInitializer(_type.PtTypeVoid)

func MakeNewArrayMaker(statements intmod.IStatements, newArray intmod.IExpression) intmod.IExpression {
	if !statements.IsEmpty() {
		ae := statements.Last().Expression().LeftExpression()

		if ae.Expression() == newArray && ae.Index().IsIntegerConstantExpression() {
			return expression.NewNewInitializedArrayWithAll(newArray.LineNumber(), newArray.Type(),
				createVariableInitializer(statements.ListIterator(), newArray))
		}
	}

	return newArray
}

func createVariableInitializer(li util.IListIterator[intmod.IStatement],
	newArray intmod.IExpression) intmod.IArrayVariableInitializer {
	statement := li.Previous()
	_ = li.Remove()

	typ := newArray.Type()
	boe := statement.Expression()
	array := declaration.NewArrayVariableInitializer(typ.CreateType(typ.Dimension() - 1))
	index := boe.LeftExpression().Index().IntegerValue()

	array.Add(declaration.NewExpressionVariableInitializer(boe.RightExpression()))

	for li.HasPrevious() {
		boe = li.Previous().Expression()

		if boe.LeftExpression().IsArrayExpression() {
			ae := boe.LeftExpression()

			if ae.Index().IsIntegerConstantExpression() {
				if ae.Expression() == newArray {
					index = ae.Index().IntegerValue()
					array.Add(declaration.NewExpressionVariableInitializer(boe.RightExpression()))
					_ = li.Remove()
					continue
				} else if ae.Expression().IsNewArray() {
					lastE := array.Last().Expression()

					if lastE.IsNewArray() && (ae.Expression() == lastE) {
						array.RemoveLast()
						li.Next()
						array.Add(createVariableInitializer(li, lastE))
						continue
					}
				}
			}
		}

		li.Next()
		break
	}

	// Replace 'new XXX[0][]' with '{}'
	vii := array.ListIterator()

	for vii.HasNext() {
		expr := vii.Next().Expression()

		if expr.IsNewArray() {
			del := expr.DimensionExpressionList()

			if del.IsIntegerConstantExpression() && (del.IntegerValue() == 0) {
				t := expr.Type()

				if typ.Dimension() == (t.Dimension()+1) &&
					len(typ.Descriptor()) == (len(t.Descriptor())+1) &&
					strings.HasSuffix(typ.Descriptor(), t.Descriptor()) {
					_ = vii.Set(EmptyArray)
				}
			}
		}
	}

	// Padding
	if index > 0 {
		var evi intmod.IExpressionVariableInitializer

		typ = typ.CreateType(typ.Dimension() - 1)

		if (typ.Dimension() == 0) && typ.IsPrimitiveType() {
			evi = declaration.NewExpressionVariableInitializer(expression.NewIntegerConstantExpression(typ, 0))
		} else {
			evi = declaration.NewExpressionVariableInitializer(expression.NewNullExpression(typ))
		}

		for ; index > 0; index-- {
			array.Add(evi)
		}
	}

	array.Reverse()

	return array
}
