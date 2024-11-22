package visitor

import "bitbucket.org/coontec/go-jd-core/class/model/javasyntax"

type UpdateJavaSyntaxTreeStep1Visitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	createInstructionsVisitor  *CreateInstructionsVisitor
	initInnerClassStep1Visitor *InitInnerClassVisitor
}
