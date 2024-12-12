package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/service/converter/visitor/utils"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewAggregateFieldsVisitor() intsrv.IAggregateFieldsVisitor {
	return &AggregateFieldsVisitor{}
}

type AggregateFieldsVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor
}

func (v *AggregateFieldsVisitor) VisitAnnotationDeclaration(declaration intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AggregateFieldsVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration := declaration.(intsrv.IClassFileBodyDeclaration)
	// Aggregate fields
	utils.Aggregate(util.NewDefaultListWithSlice(bodyDeclaration.FieldDeclarations()))
}

func (v *AggregateFieldsVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AggregateFieldsVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AggregateFieldsVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}
