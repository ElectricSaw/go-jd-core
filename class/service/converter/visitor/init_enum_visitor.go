package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	srvdecl "bitbucket.org/coontec/go-jd-core/class/service/converter/model/javasyntax/declaration"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewInitEnumVisitor() *InitEnumVisitor {
	return &InitEnumVisitor{
		constants: util.NewDefaultList[intsrv.IClassFileConstant](),
	}
}

type InitEnumVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	bodyDeclaration         intsrv.IClassFileBodyDeclaration
	constantBodyDeclaration intmod.IBodyDeclaration
	constants               util.IList[intsrv.IClassFileConstant]
	lineNumber              int
	index                   int
	arguments               intmod.IExpression
}

func (v *InitEnumVisitor) Constants() util.IList[intmod.IConstant] {
	if !v.constants.IsEmpty() {
		v.constants.Sort(func(i, j int) bool {
			return v.constants.Get(i).Index() < v.constants.Get(j).Index()
		})
	}

	constants := make([]intmod.IConstant, 0)
	for _, constant := range v.constants.ToSlice() {
		constants = append(constants, constant)
	}

	return util.NewDefaultListWithSlice(constants)
}

func (v *InitEnumVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bd := v.bodyDeclaration

	v.bodyDeclaration = declaration.(intsrv.IClassFileBodyDeclaration)
	v.constants.Clear()
	tmp1 := make([]intmod.IDeclaration, 0)
	for _, item := range v.bodyDeclaration.FieldDeclarations() {
		tmp1 = append(tmp1, item)
	}
	v.SafeAcceptListDeclaration(tmp1)
	tmp2 := make([]intmod.IDeclaration, 0)
	for _, item := range v.bodyDeclaration.MethodDeclarations() {
		tmp2 = append(tmp2, item)
	}
	v.SafeAcceptListDeclaration(tmp2)
	v.bodyDeclaration = bd
}

func (v *InitEnumVisitor) VisitConstructorDeclaration(declaration intmod.IConstructorDeclaration) {
	if (declaration.Flags() & intmod.FlagAnonymous) != 0 {
		declaration.SetFlags(intmod.FlagSynthetic)
	} else if declaration.Statements().Size() <= 1 {
		declaration.SetFlags(intmod.FlagSynthetic)
	} else {
		parameters := declaration.FormalParameters().(intmod.IFormalParameters)
		// Remove name & index parameterTypes
		parameters.SubList(0, 2).Clear()
		// Remove super constructor call
		declaration.Statements().ToList().RemoveAt(0)
		// Fix flags
		declaration.SetFlags(0)
	}
}

func (v *InitEnumVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
}

func (v *InitEnumVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	if (declaration.Flags() & (intmod.FlagStatic | intmod.FlagPublic)) != 0 {
		if declaration.Name() == "values" || declaration.Name() == "valueOf" {
			cfmd := declaration.(intsrv.IClassFileMethodDeclaration)
			cfmd.SetFlags(cfmd.Flags() | intmod.FlagSynthetic)
		}
	}
}

func (v *InitEnumVisitor) VisitFieldDeclaration(declaration intmod.IFieldDeclaration) {
	if (declaration.Flags() & intmod.FlagEnum) != 0 {
		cffd := declaration.(intsrv.IClassFileFieldDeclaration)
		cffd.FieldDeclarators().Accept(v)
		cffd.SetFlags(cffd.Flags() | intmod.FlagSynthetic)
	}
}

func (v *InitEnumVisitor) VisitFieldDeclarator(declaration intmod.IFieldDeclarator) {
	v.constantBodyDeclaration = nil
	v.SafeAcceptDeclaration(declaration.VariableInitializer())
	v.constants.Add(srvdecl.NewClassFileConstant(v.lineNumber,
		declaration.Name(), v.index, v.arguments, v.constantBodyDeclaration))
}

func (v *InitEnumVisitor) VisitNewExpression(expression intmod.INewExpression) {
	parameters := expression.Parameters().(intmod.IExpressions)
	exp := parameters.Get(1)

	if exp.IsCastExpression() {
		exp = exp.Expression()
	}

	ice := exp.(intmod.IIntegerConstantExpression)
	v.lineNumber = expression.LineNumber()
	v.index = ice.IntegerValue()

	// Remove name & index
	if parameters.Size() == 2 {
		v.arguments = nil
	} else {
		parameters.SubList(0, 2).Clear()
		v.arguments = parameters
	}

	enumInternalTypeName := expression.ObjectType().InternalName()
	if enumInternalTypeName != v.bodyDeclaration.InternalTypeName() {
		typeDeclaration := v.bodyDeclaration.InnerTypeDeclaration(enumInternalTypeName).(intsrv.IClassFileTypeDeclaration)
		if typeDeclaration != nil {
			v.constantBodyDeclaration = typeDeclaration.BodyDeclaration()
		}
	}
}
