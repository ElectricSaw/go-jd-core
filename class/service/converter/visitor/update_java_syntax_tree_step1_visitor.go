package visitor

import (
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
)

func NewUpdateJavaSyntaxTreeStep1Visitor(typeMaker *utils.TypeMaker) *UpdateJavaSyntaxTreeStep1Visitor {
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
