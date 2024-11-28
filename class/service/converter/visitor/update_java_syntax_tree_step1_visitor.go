package visitor

import (
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
)

func NewUpdateJavaSyntaxTreeStep1Visitor(typeMaker intsrv.ITypeMaker) *UpdateJavaSyntaxTreeStep1Visitor {
	return &UpdateJavaSyntaxTreeStep1Visitor{
		updateOuterFieldTypeVisitor:   *NewUpdateOuterFieldTypeVisitor(typeMaker),
		updateBridgeMethodTypeVisitor: *NewUpdateBridgeMethodTypeVisitor(typeMaker),
	}
}

type UpdateJavaSyntaxTreeStep1Visitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	createInstructionsVisitor  *CreateInstructionsVisitor
	initInnerClassStep1Visitor *InitInnerClassVisitor
}
