package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
)

func NewSearchLocalVariableReferenceVisitor() *SearchLocalVariableReferenceVisitor {
	return &SearchLocalVariableReferenceVisitor{}
}

type SearchLocalVariableReferenceVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	index int
	found bool
}

func (v *SearchLocalVariableReferenceVisitor) Init(index int) {
	v.index = index
	v.found = false
}

func (v *SearchLocalVariableReferenceVisitor) ContainsReference() bool {
	return v.found
}

func (v *SearchLocalVariableReferenceVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	if v.index < 0 {
		v.found = true
	} else {
		refExpr := expr.(intsrv.IClassFileLocalVariableReferenceExpression)
		v.found = v.found || refExpr.LocalVariable().(intsrv.ILocalVariable).Index() == v.index
	}
}
