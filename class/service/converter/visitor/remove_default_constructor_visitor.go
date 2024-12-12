package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewRemoveDefaultConstructorVisitor() intsrv.IRemoveDefaultConstructorVisitor {
	return &RemoveDefaultConstructorVisitor{}
}

type RemoveDefaultConstructorVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	constructorCounter int
	constructor        intsrv.IClassFileMemberDeclaration
}

func (v *RemoveDefaultConstructorVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *RemoveDefaultConstructorVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)
	methods := util.NewDefaultListWithSlice(bodyDeclaration.MethodDeclarations())

	v.constructor = nil
	v.constructorCounter = 0
	tmp1 := make([]intmod.IDeclaration, 0)
	for _, item := range methods.ToSlice() {
		tmp1 = append(tmp1, item)
	}
	v.SafeAcceptListDeclaration(tmp1)

	if (v.constructorCounter == 1) && (v.constructor != nil) {
		// Remove empty default constructor
		methods.Remove(v.constructor.(intsrv.IClassFileConstructorOrMethodDeclaration))
	}
}

func (v *RemoveDefaultConstructorVisitor) VisitFieldDeclaration(_ intmod.IFieldDeclaration) {}

func (v *RemoveDefaultConstructorVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	if (decl.Flags() & intmod.FlagAbstract) == 0 {
		cfcd := decl.(intsrv.IClassFileConstructorDeclaration)

		if (cfcd.Statements() != nil) && cfcd.Statements().IsStatements() {
			statements := cfcd.Statements().(intmod.IStatements)

			// Remove no-parameter super constructor call and anonymous class super constructor call
			iterator := statements.Iterator()

			for iterator.HasNext() {
				es := iterator.Next().Expression()

				if es.IsSuperConstructorInvocationExpression() {
					if (decl.Flags() & intmod.FlagAnonymous) == 0 {
						parameters := es.Parameters()

						if (parameters == nil) || (parameters.Size() == 0) {
							// Remove 'super();'
							_ = iterator.Remove()
							break
						}
					} else {
						// Remove anonymous class super constructor call
						_ = iterator.Remove()
						break
					}
				}
			}

			// Store empty default constructor
			if statements.IsEmpty() {
				if (cfcd.FormalParameters() == nil) || (cfcd.FormalParameters().Size() == 0) {
					v.constructor = cfcd
				}
			}
		}

		// Inc constructor counter
		v.constructorCounter++
	}
}

func (v *RemoveDefaultConstructorVisitor) VisitMethodDeclaration(_ intmod.IMethodDeclaration) {}

func (v *RemoveDefaultConstructorVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
}

func (v *RemoveDefaultConstructorVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *RemoveDefaultConstructorVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *RemoveDefaultConstructorVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}
