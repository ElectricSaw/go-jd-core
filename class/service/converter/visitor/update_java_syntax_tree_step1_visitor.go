package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
)

func NewUpdateJavaSyntaxTreeStep1Visitor(typeMaker intsrv.ITypeMaker) *UpdateJavaSyntaxTreeStep1Visitor {
	return &UpdateJavaSyntaxTreeStep1Visitor{
		createInstructionsVisitor:  NewCreateInstructionsVisitor(typeMaker),
		initInnerClassStep1Visitor: NewInitInnerClassVisitor(),
	}
}

type UpdateJavaSyntaxTreeStep1Visitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	createInstructionsVisitor  *CreateInstructionsVisitor
	initInnerClassStep1Visitor *InitInnerClassVisitor
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
