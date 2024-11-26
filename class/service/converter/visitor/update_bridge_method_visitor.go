package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	srvexp "bitbucket.org/coontec/go-jd-core/class/service/converter/model/javasyntax/expression"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
	"fmt"
	"strings"
)

func NewUpdateBridgeMethodVisitor(typeMaker *utils.TypeMaker) *UpdateBridgeMethodVisitor {
	v := &UpdateBridgeMethodVisitor{
		bodyDeclarationsVisitor: &BodyDeclarationsVisitor{
			mapped: make(map[string]intsrv.IClassFileMethodDeclaration),
		},
		bridgeMethodDeclarations: make(map[string]map[string]intsrv.IClassFileMethodDeclaration),
		typeMaker:                typeMaker,
	}

	v.bodyDeclarationsVisitor.bridgeMethodDeclarationsLink = v.bridgeMethodDeclarations

	return v
}

type UpdateBridgeMethodVisitor struct {
	AbstractUpdateExpressionVisitor

	bodyDeclarationsVisitor  *BodyDeclarationsVisitor
	bridgeMethodDeclarations map[string]map[string]intsrv.IClassFileMethodDeclaration
	typeMaker                *utils.TypeMaker
}

func (v *UpdateBridgeMethodVisitor) Init(bodyDeclaration intsrv.IClassFileBodyDeclaration) bool {
	v.bridgeMethodDeclarations = make(map[string]map[string]intsrv.IClassFileMethodDeclaration)
	v.bodyDeclarationsVisitor.bridgeMethodDeclarationsLink = v.bridgeMethodDeclarations
	v.bodyDeclarationsVisitor.VisitBodyDeclaration(bodyDeclaration)
	return len(v.bridgeMethodDeclarations) != 0
}

func (v *UpdateBridgeMethodVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.SetExpression(v.updateExpression(expr.Expression()))

	if expr.Parameters() != nil {
		expr.SetParameters(v.UpdateBaseExpression(expr.Parameters()))
		expr.Parameters().Accept(v)
	}

	expr.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) updateExpression(expr intmod.IExpression) intmod.IExpression {
	if !expr.IsMethodInvocationExpression() {
		return expr
	}

	mie1 := expr.(intsrv.IClassFileMethodInvocationExpression)
	mapped := v.bridgeMethodDeclarations[mie1.Expression().Type().Descriptor()]

	if mapped == nil {
		return expr
	}

	bridgeMethodDeclaration := mapped[mie1.Name()+mie1.Descriptor()].(intsrv.IClassFileMethodDeclaration)

	if bridgeMethodDeclaration == nil {
		return expr
	}

	statement := bridgeMethodDeclaration.Statements().First()
	var exp intmod.IExpression

	if statement.IsReturnExpressionStatement() {
		exp = statement.Expression()
	} else if statement.IsExpressionStatement() {
		exp = statement.Expression()
	} else {
		return expr
	}

	parameterTypes := bridgeMethodDeclaration.ParameterTypes()
	parameterTypesCount := 0

	if parameterTypes != nil {
		parameterTypesCount = parameterTypes.Size()
	}

	if exp.IsFieldReferenceExpression() {
		fre := getFieldReferenceExpression(exp)

		if parameterTypesCount == 0 {
			expr = fre.Expression()
		} else {
			expr = mie1.Parameters().First()
		}

		return expression.NewFieldReferenceExpressionWithAll(mie1.LineNumber(), fre.Type(), expr, fre.InternalTypeName(), fre.Name(), fre.Descriptor())
	} else if exp.IsMethodInvocationExpression() {
		mie2 := exp.(intmod.IMethodInvocationExpression)
		methodTypes := v.typeMaker.MakeMethodTypes2(mie2.InternalTypeName(), mie2.Name(), mie2.Descriptor())

		if methodTypes != nil {
			if mie2.Expression().IsObjectTypeReferenceExpression() {
				// Static method invocation
				return srvexp.NewClassFileMethodInvocationExpression(mie1.LineNumber(), nil,
					methodTypes.ReturnedType, mie2.Expression(), mie2.InternalTypeName(),
					mie2.Name(), mie2.Descriptor(), methodTypes.ParameterTypes, mie1.Parameters())
			} else {
				mie1Parameters := mie1.Parameters()
				var newParameters intmod.IExpression

				switch mie1Parameters.Size() {
				case 0, 1:
				case 2:
					newParameters = mie1Parameters.ToSlice()[1]
				default:
					p := mie1Parameters.ToSlice()
					newParameters = expression.NewExpressions()
					newParameters.(intmod.IExpressions).AddAll(p[1 : 1+len(p)])
				}

				return srvexp.NewClassFileMethodInvocationExpression(mie1.LineNumber(), nil,
					methodTypes.ReturnedType, mie1Parameters.First(), mie2.InternalTypeName(),
					mie2.Name(), mie2.Descriptor(), methodTypes.ParameterTypes, newParameters)
			}
		}
	} else if exp.IsBinaryOperatorExpression() {
		fre := getFieldReferenceExpression(exp.LeftExpression())

		if parameterTypesCount == 1 {
			return expression.NewBinaryOperatorExpression(
				mie1.LineNumber(), mie1.Type(),
				expression.NewFieldReferenceExpression(fre.Type(), fre.Expression(), fre.InternalTypeName(),
					fre.Name(), fre.Descriptor()), exp.Operator(), mie1.Parameters().First(), exp.Priority())
		} else if parameterTypesCount == 2 {
			parameters := mie1.Parameters().ToSlice()

			return expression.NewBinaryOperatorExpression(
				mie1.LineNumber(), mie1.Type(),
				expression.NewFieldReferenceExpression(fre.Type(), parameters[0], fre.InternalTypeName(), fre.Name(), fre.Descriptor()),
				exp.Operator(), parameters[1], exp.Priority())
		}
	} else if exp.IsPostOperatorExpression() {
		fre := getFieldReferenceExpression(exp.Expression())

		if parameterTypesCount == 0 {
			expr = fre.Expression()
		} else {
			expr = mie1.Parameters().First()
		}

		return expression.NewPostOperatorExpressionWithAll(mie1.LineNumber(), exp.Operator(),
			expression.NewFieldReferenceExpression(fre.Type(), expr, fre.InternalTypeName(),
				fre.Name(), fre.Descriptor()))
	} else if exp.IsPreOperatorExpression() {
		fre := getFieldReferenceExpression(exp.Expression())

		if parameterTypesCount == 0 {
			expr = fre.Expression()
		} else {
			expr = mie1.Parameters().First()
		}

		return expression.NewPreOperatorExpressionWithAll(
			mie1.LineNumber(), exp.Operator(), expression.NewFieldReferenceExpression(
				fre.Type(), expr, fre.InternalTypeName(), fre.Name(), fre.Descriptor()))
	} else if exp.IsIntegerConstantExpression() {
		return exp
	}

	return expr
}

type BodyDeclarationsVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	bridgeMethodDeclarationsLink map[string]map[string]intsrv.IClassFileMethodDeclaration
	mapped                       map[string]intsrv.IClassFileMethodDeclaration
}

func (v *BodyDeclarationsVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *BodyDeclarationsVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *BodyDeclarationsVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
}

func (v *BodyDeclarationsVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
}

func (v *BodyDeclarationsVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)
	methodDeclarations := bodyDeclaration.MethodDeclarations()

	if methodDeclarations != nil && len(methodDeclarations) != 0 {
		backup := v.mapped
		v.mapped = make(map[string]intsrv.IClassFileMethodDeclaration)
		v.AcceptListDeclaration(ConvertMethodDeclarations(methodDeclarations))

		if len(v.mapped) != 0 {
			key := fmt.Sprintf("L%s;", bodyDeclaration.InternalTypeName())
			v.bridgeMethodDeclarationsLink[key] = v.mapped
		}

		v.mapped = backup
	}

	v.SafeAcceptListDeclaration(ConvertInnerTypeDeclarations(bodyDeclaration.InnerTypeDeclarations()))
}

func (v *BodyDeclarationsVisitor) VisitStaticInitializerDeclaration(decl intmod.IStaticInitializerDeclaration) {
}

func (v *BodyDeclarationsVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
}

func (v *BodyDeclarationsVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	if (decl.Flags() & intmod.FlagStatic) == 0 {
		return
	}

	statements := decl.Statements()

	if statements == nil || statements.Size() != 1 {
		return
	}

	name := decl.Name()
	if !strings.HasPrefix(name, "access$") {
		return
	}

	bridgeMethodDeclaration := decl.(intsrv.IClassFileMethodDeclaration)
	if !v.checkBridgeMethodDeclaration(bridgeMethodDeclaration) {
		return
	}

	v.mapped[name+decl.Descriptor()] = bridgeMethodDeclaration
}

func (v *BodyDeclarationsVisitor) checkBridgeMethodDeclaration(bridgeMethodDeclaration intsrv.IClassFileMethodDeclaration) bool {
	state := bridgeMethodDeclaration.Statements().First()
	var expr intmod.IExpression

	if state.IsReturnExpressionStatement() {
		expr = state.Expression()
	} else if state.IsExpressionStatement() {
		expr = state.Expression()
	} else {
		return false
	}

	parameterTypes := bridgeMethodDeclaration.ParameterTypes()
	parameterTypesCount := 0
	if parameterTypes != nil {
		parameterTypesCount = parameterTypes.Size()
	}

	if expr.IsFieldReferenceExpression() {
		fre := expr.(intmod.IFieldReferenceExpression)

		if parameterTypesCount == 0 {
			return (fre.Expression() != nil) && fre.Expression().IsObjectTypeReferenceExpression()
		} else if parameterTypesCount == 1 {
			return (fre.Expression() == nil) || v.checkLocalVariableReference(fre.Expression(), 0)
		}
	} else if expr.IsMethodInvocationExpression() {
		mie2 := expr.(intmod.IMethodInvocationExpression)

		if mie2.Expression().IsObjectTypeReferenceExpression() {
			mie2Parameters := mie2.Parameters()

			if (mie2Parameters == nil) || (mie2Parameters.Size() == 0) {
				return true
			}

			if mie2Parameters.IsList() {
				i := 0
				for _, parameter := range mie2Parameters.Iterator().List() {
					if v.checkLocalVariableReference(parameter, i) {
						i++
						return false
					}

					t := parameter.Type()
					if t == _type.PtTypeLong || t == _type.PtTypeDouble {
						i++
					}
				}
				return true
			}

			return v.checkLocalVariableReference(mie2Parameters, 0)
		} else if parameterTypesCount > 0 && v.checkLocalVariableReference(mie2.Expression(), 0) {
			mie2Parameters := mie2.Parameters()

			if mie2Parameters == nil || mie2Parameters.Size() == 0 {
				return true
			}

			if mie2Parameters.IsList() {
				i := 0
				for _, parameter := range mie2Parameters.Iterator().List() {
					if !v.checkLocalVariableReference(parameter, i) {
						i++
						return false
					}
					i++
					t := parameter.Type()
					if t == _type.PtTypeLong || t == _type.PtTypeDouble {
						i++
					}
				}
				return true
			}
			return v.checkLocalVariableReference(mie2Parameters, 1)
		}
	} else if expr.IsBinaryOperatorExpression() {
		if parameterTypesCount == 1 {
			if expr.LeftExpression().IsFieldReferenceExpression() && v.checkLocalVariableReference(expr.RightExpression(), 0) {
				fre := expr.LeftExpression().(intmod.IFieldReferenceExpression)
				return fre.Expression().IsObjectTypeReferenceExpression()
			}
		} else if parameterTypesCount == 2 {
			if expr.LeftExpression().IsFieldReferenceExpression() && v.checkLocalVariableReference(expr.RightExpression(), 1) {
				fre := expr.LeftExpression().(intmod.IFieldReferenceExpression)
				return v.checkLocalVariableReference(fre.Expression(), 0)
			}
		}
	} else if expr.IsPostOperatorExpression() {
		expr = expr.Expression()

		if expr.IsFieldReferenceExpression() {
			if (parameterTypesCount == 0) && expr.Expression().IsObjectTypeReferenceExpression() {
				return true
			} else if (parameterTypesCount == 1) && (expr.Expression() != nil) && v.checkLocalVariableReference(expr.Expression(), 0) {
				return true
			}
		}
	} else if (parameterTypesCount == 1) && expr.IsPreOperatorExpression() {
		expr = expr.Expression()

		if expr.IsFieldReferenceExpression() {
			if parameterTypesCount == 0 && expr.Expression().IsObjectTypeReferenceExpression() {
				return true
			} else if parameterTypesCount == 1 && expr.Expression() != nil && v.checkLocalVariableReference(expr.Expression(), 0) {
				return true
			}
		}
	} else if parameterTypesCount == 0 && expr.IsIntegerConstantExpression() {
		return true
	}

	return false
}

func (v *BodyDeclarationsVisitor) checkLocalVariableReference(expr intmod.IExpression, index int) bool {
	if expr.IsLocalVariableReferenceExpression() {
		variable := expr.(intsrv.IClassFileLocalVariableReferenceExpression)
		return variable.LocalVariable().(intsrv.ILocalVariable).Index() == index
	}

	return false
}

func getFieldReferenceExpression(expression intmod.IExpression) intmod.IFieldReferenceExpression {
	fre := expression.(intmod.IFieldReferenceExpression)
	freExpression := fre.Expression()

	if freExpression != nil && freExpression.IsObjectTypeReferenceExpression() {
		freExpression.(intmod.IObjectTypeReferenceExpression).SetExplicit(true)
	}

	return fre
}
