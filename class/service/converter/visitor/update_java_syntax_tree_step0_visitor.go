package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
)

func NewUpdateJavaSyntaxTreeStep0Visitor(typeMaker *utils.TypeMaker) *UpdateJavaSyntaxTreeStep0Visitor {
	return &UpdateJavaSyntaxTreeStep0Visitor{
		updateOuterFieldTypeVisitor:   *NewUpdateOuterFieldTypeVisitor(typeMaker),
		updateBridgeMethodTypeVisitor: *NewUpdateBridgeMethodTypeVisitor(typeMaker),
	}
}

type UpdateJavaSyntaxTreeStep0Visitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	updateOuterFieldTypeVisitor   UpdateOuterFieldTypeVisitor
	updateBridgeMethodTypeVisitor UpdateBridgeMethodTypeVisitor
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)
	genericTypesSupported := bodyDeclaration.ClassFile().MajorVersion() >= 49 // (majorVersion >= Java 5)

	if genericTypesSupported {
		v.updateOuterFieldTypeVisitor.SafeAcceptListDeclaration(
			ConvertInnerTypeDeclarations(bodyDeclaration.InnerTypeDeclarations()))
		v.updateBridgeMethodTypeVisitor.VisitBodyDeclaration(decl)
	}
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
}
func (v *UpdateJavaSyntaxTreeStep0Visitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {}
