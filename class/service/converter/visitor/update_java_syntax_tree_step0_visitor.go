package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
)

func NewUpdateJavaSyntaxTreeStep0Visitor(typeMaker intsrv.ITypeMaker) intsrv.IUpdateJavaSyntaxTreeStep0Visitor {
	return &UpdateJavaSyntaxTreeStep0Visitor{
		updateOuterFieldTypeVisitor:   NewUpdateOuterFieldTypeVisitor(typeMaker),
		updateBridgeMethodTypeVisitor: NewUpdateBridgeMethodTypeVisitor(typeMaker),
	}
}

type UpdateJavaSyntaxTreeStep0Visitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	updateOuterFieldTypeVisitor   intsrv.IUpdateOuterFieldTypeVisitor
	updateBridgeMethodTypeVisitor intsrv.IUpdateBridgeMethodTypeVisitor
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)
	genericTypesSupported := bodyDeclaration.ClassFile().MajorVersion() >= 49 // (majorVersion >= Java 5)

	if genericTypesSupported {
		v.updateOuterFieldTypeVisitor.SafeAcceptListDeclaration(
			ConvertTypeDeclarations(bodyDeclaration.InnerTypeDeclarations()))
		v.updateBridgeMethodTypeVisitor.VisitBodyDeclaration(decl)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitAnnotationDeclaration(_ intmod.IAnnotationDeclaration) {
}
func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitEnumDeclaration(_ intmod.IEnumDeclaration) {}
