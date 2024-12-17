package visitor

import (
	"fmt"
	"strings"

	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/expression"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	srvexp "github.com/ElectricSaw/go-jd-core/class/service/converter/model/javasyntax/expression"
)

func NewUpdateBridgeMethodVisitor(typeMaker intsrv.ITypeMaker) intsrv.IUpdateBridgeMethodVisitor {
	v := &UpdateBridgeMethodVisitor{
		bodyDeclarationsVisitor:  NewBodyDeclarationsVisitor(),
		bridgeMethodDeclarations: make(map[string]map[string]intsrv.IClassFileMethodDeclaration),
		typeMaker:                typeMaker,
	}

	v.bodyDeclarationsVisitor.SetBridgeMethodDeclarationsLink(v.bridgeMethodDeclarations)

	return v
}

type UpdateBridgeMethodVisitor struct {
	AbstractUpdateExpressionVisitor

	bodyDeclarationsVisitor  intsrv.IBodyDeclarationsVisitor
	bridgeMethodDeclarations map[string]map[string]intsrv.IClassFileMethodDeclaration
	typeMaker                intsrv.ITypeMaker
}

func (v *UpdateBridgeMethodVisitor) Init(bodyDeclaration intsrv.IClassFileBodyDeclaration) bool {
	v.bridgeMethodDeclarations = make(map[string]map[string]intsrv.IClassFileMethodDeclaration)
	v.bodyDeclarationsVisitor.SetBridgeMethodDeclarationsLink(v.bridgeMethodDeclarations)
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
					methodTypes.ReturnedType(), mie2.Expression(), mie2.InternalTypeName(),
					mie2.Name(), mie2.Descriptor(), methodTypes.ParameterTypes(), mie1.Parameters())
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
					methodTypes.ReturnedType(), mie1Parameters.First(), mie2.InternalTypeName(),
					mie2.Name(), mie2.Descriptor(), methodTypes.ParameterTypes(), newParameters)
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

func NewBodyDeclarationsVisitor() intsrv.IBodyDeclarationsVisitor {
	return &BodyDeclarationsVisitor{
		bridgeMethodDeclarationsLink: make(map[string]map[string]intsrv.IClassFileMethodDeclaration),
		mapped:                       make(map[string]intsrv.IClassFileMethodDeclaration),
	}
}

type BodyDeclarationsVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	bridgeMethodDeclarationsLink map[string]map[string]intsrv.IClassFileMethodDeclaration
	mapped                       map[string]intsrv.IClassFileMethodDeclaration
}

func (v *BodyDeclarationsVisitor) BridgeMethodDeclarationsLink() map[string]map[string]intsrv.IClassFileMethodDeclaration {
	return v.bridgeMethodDeclarationsLink
}

func (v *BodyDeclarationsVisitor) SetBridgeMethodDeclarationsLink(bridgeMethodDeclarationsLink map[string]map[string]intsrv.IClassFileMethodDeclaration) {
	v.bridgeMethodDeclarationsLink = bridgeMethodDeclarationsLink
}

func (v *BodyDeclarationsVisitor) Mapped() map[string]intsrv.IClassFileMethodDeclaration {
	return v.mapped
}

func (v *BodyDeclarationsVisitor) SetMapped(mapped map[string]intsrv.IClassFileMethodDeclaration) {
	v.mapped = mapped
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

	v.SafeAcceptListDeclaration(ConvertTypeDeclarations(bodyDeclaration.InnerTypeDeclarations()))
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
				for _, parameter := range mie2Parameters.ToSlice() {
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
				for _, parameter := range mie2Parameters.ToSlice() {
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

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                                 $$$
// $$$ AbstractUpdateExpressionVisitor $$$
// $$$                                 $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *UpdateBridgeMethodVisitor) UpdateExpression(_ intmod.IExpression) intmod.IExpression {
	return nil
}

func (v *UpdateBridgeMethodVisitor) UpdateBaseExpression(baseExpression intmod.IExpression) intmod.IExpression {
	if baseExpression == nil {
		return nil
	}

	if baseExpression.IsList() {
		iterator := baseExpression.ToList().ListIterator()

		for iterator.HasNext() {
			value := iterator.Next()
			_ = iterator.Set(v.UpdateExpression(value))
		}

		return baseExpression
	}

	return v.UpdateExpression(baseExpression.First())
}

func (v *UpdateBridgeMethodVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateBridgeMethodVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.VisitInterfaceDeclaration(declaration)
}

func (v *UpdateBridgeMethodVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
}

func (v *UpdateBridgeMethodVisitor) VisitConstructorDeclaration(declaration intmod.IConstructorDeclaration) {
	v.SafeAcceptStatement(declaration.Statements())
}

func (v *UpdateBridgeMethodVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateBridgeMethodVisitor) VisitEnumDeclarationConstant(declaration intmod.IConstant) {
	if declaration.Arguments() != nil {
		declaration.SetArguments(v.UpdateBaseExpression(declaration.Arguments()))
		declaration.Arguments().Accept(v)
	}
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateBridgeMethodVisitor) VisitExpressionVariableInitializer(declaration intmod.IExpressionVariableInitializer) {
	if declaration.Expression() != nil {
		declaration.SetExpression(v.UpdateExpression(declaration.Expression()))
		declaration.Expression().Accept(v)
	}
}

func (v *UpdateBridgeMethodVisitor) VisitFieldDeclaration(declaration intmod.IFieldDeclaration) {
	declaration.FieldDeclarators().AcceptDeclaration(v)
}

func (v *UpdateBridgeMethodVisitor) VisitFieldDeclarator(declaration intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(declaration.VariableInitializer())
}

func (v *UpdateBridgeMethodVisitor) VisitFormalParameter(_ intmod.IFormalParameter) {
}

func (v *UpdateBridgeMethodVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateBridgeMethodVisitor) VisitLocalVariableDeclaration(declaration intmod.ILocalVariableDeclaration) {
	declaration.LocalVariableDeclarators().AcceptDeclaration(v)
}

func (v *UpdateBridgeMethodVisitor) VisitLocalVariableDeclarator(declarator intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(declarator.VariableInitializer())
}

func (v *UpdateBridgeMethodVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	v.SafeAcceptReference(declaration.AnnotationReferences())
	v.SafeAcceptStatement(declaration.Statements())
}

func (v *UpdateBridgeMethodVisitor) VisitArrayExpression(expression intmod.IArrayExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.SetIndex(v.UpdateExpression(expression.Index()))
	expression.Expression().Accept(v)
	expression.Index().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitBinaryOperatorExpression(expression intmod.IBinaryOperatorExpression) {
	expression.SetLeftExpression(v.UpdateExpression(expression.LeftExpression()))
	expression.SetRightExpression(v.UpdateExpression(expression.RightExpression()))
	expression.LeftExpression().Accept(v)
	expression.RightExpression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitCastExpression(expression intmod.ICastExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitFieldReferenceExpression(expression intmod.IFieldReferenceExpression) {
	if expression.Expression() != nil {
		expression.SetExpression(v.UpdateExpression(expression.Expression()))
		expression.Expression().Accept(v)
	}
}

func (v *UpdateBridgeMethodVisitor) VisitInstanceOfExpression(expression intmod.IInstanceOfExpression) {
	if expression.Expression() != nil {
		expression.SetExpression(v.UpdateExpression(expression.Expression()))
		expression.Expression().Accept(v)
	}
}

func (v *UpdateBridgeMethodVisitor) VisitLambdaFormalParametersExpression(expression intmod.ILambdaFormalParametersExpression) {
	expression.Statements().AcceptStatement(v)
}

func (v *UpdateBridgeMethodVisitor) VisitLengthExpression(expression intmod.ILengthExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitMethodReferenceExpression(expression intmod.IMethodReferenceExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitNewArray(expression intmod.INewArray) {
	if expression.DimensionExpressionList() != nil {
		expression.SetDimensionExpressionList(v.UpdateBaseExpression(expression.DimensionExpressionList()))
		expression.DimensionExpressionList().Accept(v)
	}
}

func (v *UpdateBridgeMethodVisitor) VisitNewExpression(expression intmod.INewExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
	// v.SafeAccept(expression.BodyDeclaration());
}

func (v *UpdateBridgeMethodVisitor) VisitNewInitializedArray(expression intmod.INewInitializedArray) {
	v.SafeAcceptDeclaration(expression.ArrayInitializer())
}

func (v *UpdateBridgeMethodVisitor) VisitParenthesesExpression(expression intmod.IParenthesesExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitPostOperatorExpression(expression intmod.IPostOperatorExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitPreOperatorExpression(expression intmod.IPreOperatorExpression) {
	expression.SetExpression(v.UpdateExpression(expression.Expression()))
	expression.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitSuperConstructorInvocationExpression(expression intmod.ISuperConstructorInvocationExpression) {
	if expression.Parameters() != nil {
		expression.SetParameters(v.UpdateBaseExpression(expression.Parameters()))
		expression.Parameters().Accept(v)
	}
}

func (v *UpdateBridgeMethodVisitor) VisitTernaryOperatorExpression(expression intmod.ITernaryOperatorExpression) {
	expression.SetCondition(v.UpdateExpression(expression.Condition()))
	expression.SetTrueExpression(v.UpdateExpression(expression.TrueExpression()))
	expression.SetFalseExpression(v.UpdateExpression(expression.FalseExpression()))
	expression.Condition().Accept(v)
	expression.TrueExpression().Accept(v)
	expression.FalseExpression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitExpressionElementValue(reference intmod.IExpressionElementValue) {
	reference.SetExpression(v.UpdateExpression(reference.Expression()))
	reference.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitAssertStatement(statement intmod.IAssertStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptExpression(statement.Message())
}

func (v *UpdateBridgeMethodVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	v.SafeAcceptExpression(statement.Condition())
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateBridgeMethodVisitor) VisitExpressionStatement(statement intmod.IExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateBridgeMethodVisitor) VisitForStatement(statement intmod.IForStatement) {
	v.SafeAcceptDeclaration(statement.Declaration())
	if statement.Init() != nil {
		statement.SetInit(v.UpdateBaseExpression(statement.Init()))
		statement.Init().Accept(v)
	}
	if statement.Condition() != nil {
		statement.SetCondition(v.UpdateExpression(statement.Condition()))
		statement.Condition().Accept(v)
	}
	if statement.Update() != nil {
		statement.SetUpdate(v.UpdateBaseExpression(statement.Update()))
		statement.Update().Accept(v)
	}
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateBridgeMethodVisitor) VisitIfStatement(statement intmod.IIfStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateBridgeMethodVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
	statement.ElseStatements().AcceptStatement(v)
}

func (v *UpdateBridgeMethodVisitor) VisitLambdaExpressionStatement(statement intmod.ILambdaExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitReturnExpressionStatement(statement intmod.IReturnExpressionStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.AcceptListStatement(statement.List())
}

func (v *UpdateBridgeMethodVisitor) VisitSwitchStatementExpressionLabel(statement intmod.IExpressionLabel) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
	statement.SetMonitor(v.UpdateExpression(statement.Monitor()))
	statement.Monitor().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateBridgeMethodVisitor) VisitThrowStatement(statement intmod.IThrowStatement) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitTryStatementCatchClause(statement intmod.ICatchClause) {
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateBridgeMethodVisitor) VisitTryStatementResource(statement intmod.IResource) {
	statement.SetExpression(v.UpdateExpression(statement.Expression()))
	statement.Expression().Accept(v)
}

func (v *UpdateBridgeMethodVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {
	statement.SetCondition(v.UpdateExpression(statement.Condition()))
	statement.Condition().Accept(v)
	v.SafeAcceptStatement(statement.Statements())
}

func (v *UpdateBridgeMethodVisitor) VisitConstructorReferenceExpression(_ intmod.IConstructorReferenceExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitDoubleConstantExpression(_ intmod.IDoubleConstantExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitEnumConstantReferenceExpression(_ intmod.IEnumConstantReferenceExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitFloatConstantExpression(_ intmod.IFloatConstantExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitLocalVariableReferenceExpression(_ intmod.ILocalVariableReferenceExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitLongConstantExpression(_ intmod.ILongConstantExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitNullExpression(_ intmod.INullExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitTypeReferenceDotClassExpression(_ intmod.ITypeReferenceDotClassExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitObjectTypeReferenceExpression(_ intmod.IObjectTypeReferenceExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitSuperExpression(_ intmod.ISuperExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitThisExpression(_ intmod.IThisExpression) {
}
func (v *UpdateBridgeMethodVisitor) VisitAnnotationReference(_ intmod.IAnnotationReference) {
}
func (v *UpdateBridgeMethodVisitor) VisitElementValueArrayInitializerElementValue(_ intmod.IElementValueArrayInitializerElementValue) {
}
func (v *UpdateBridgeMethodVisitor) VisitAnnotationElementValue(_ intmod.IAnnotationElementValue) {
}
func (v *UpdateBridgeMethodVisitor) VisitObjectReference(_ intmod.IObjectReference) {
}
func (v *UpdateBridgeMethodVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {}
func (v *UpdateBridgeMethodVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
}
func (v *UpdateBridgeMethodVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
}
func (v *UpdateBridgeMethodVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
}
func (v *UpdateBridgeMethodVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
}

func (v *UpdateBridgeMethodVisitor) VisitInnerObjectReference(_ intmod.IInnerObjectReference) {
}
func (v *UpdateBridgeMethodVisitor) VisitTypeArguments(_ intmod.ITypeArguments) {}
func (v *UpdateBridgeMethodVisitor) VisitWildcardExtendsTypeArgument(_ intmod.IWildcardExtendsTypeArgument) {
}
func (v *UpdateBridgeMethodVisitor) VisitObjectType(_ intmod.IObjectType)           {}
func (v *UpdateBridgeMethodVisitor) VisitInnerObjectType(_ intmod.IInnerObjectType) {}
func (v *UpdateBridgeMethodVisitor) VisitWildcardSuperTypeArgument(_ intmod.IWildcardSuperTypeArgument) {
}
func (v *UpdateBridgeMethodVisitor) VisitTypes(_ intmod.ITypes) {}
func (v *UpdateBridgeMethodVisitor) VisitTypeParameterWithTypeBounds(_ intmod.ITypeParameterWithTypeBounds) {
}

func (v *UpdateBridgeMethodVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *UpdateBridgeMethodVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateBridgeMethodVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateBridgeMethodVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *UpdateBridgeMethodVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *UpdateBridgeMethodVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *UpdateBridgeMethodVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *UpdateBridgeMethodVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *UpdateBridgeMethodVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *UpdateBridgeMethodVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *UpdateBridgeMethodVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateBridgeMethodVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateBridgeMethodVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
