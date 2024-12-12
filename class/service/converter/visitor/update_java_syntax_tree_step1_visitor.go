package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
)

func NewUpdateJavaSyntaxTreeStep1Visitor(typeMaker intsrv.ITypeMaker) intsrv.IUpdateJavaSyntaxTreeStep1Visitor {
	return &UpdateJavaSyntaxTreeStep1Visitor{
		createInstructionsVisitor:  NewCreateInstructionsVisitor(typeMaker),
		initInnerClassStep1Visitor: NewInitInnerClassVisitor(),
	}
}

type UpdateJavaSyntaxTreeStep1Visitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	createInstructionsVisitor  intsrv.ICreateInstructionsVisitor
	initInnerClassStep1Visitor intsrv.IInitInnerClassVisitor
}

func (v *UpdateJavaSyntaxTreeStep1Visitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)

	// Visit inner types
	if bodyDeclaration.InnerTypeDeclarations() != nil {
		tmp := make([]intmod.IDeclaration, 0)
		for _, item := range bodyDeclaration.InnerTypeDeclarations() {
			tmp = append(tmp, item)
		}
		v.AcceptListDeclaration(tmp)
	}

	// Visit declaration
	v.createInstructionsVisitor.VisitBodyDeclaration(decl)
	v.initInnerClassStep1Visitor.VisitBodyDeclaration(decl)
}
