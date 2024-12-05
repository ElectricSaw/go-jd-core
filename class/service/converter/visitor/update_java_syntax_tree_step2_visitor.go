package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
)

var aggregateFieldsVisitor = NewAggregateFieldsVisitor()
var sortMembersVisitor = NewSortMembersVisitor()
var autoboxingVisitor = NewAutoboxingVisitor()

func NewUpdateJavaSyntaxTreeStep2Visitor(typeMaker intsrv.ITypeMaker) *UpdateJavaSyntaxTreeStep2Visitor {
	return &UpdateJavaSyntaxTreeStep2Visitor{
		initStaticFieldVisitor:          NewInitStaticFieldVisitor(),
		initInstanceFieldVisitor:        NewInitInstanceFieldVisitor(),
		initEnumVisitor:                 NewInitEnumVisitor(),
		removeDefaultConstructorVisitor: NewRemoveDefaultConstructorVisitor(),
		replaceBridgeMethodVisitor:      NewUpdateBridgeMethodVisitor(typeMaker),
		initInnerClassStep2Visitor:      NewUpdateNewExpressionVisitor(typeMaker),
		addCastExpressionVisitor:        NewAddCastExpressionVisitor(typeMaker),
	}
}

type UpdateJavaSyntaxTreeStep2Visitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	initStaticFieldVisitor          *InitStaticFieldVisitor
	initInstanceFieldVisitor        *InitInstanceFieldVisitor
	initEnumVisitor                 *InitEnumVisitor
	removeDefaultConstructorVisitor *RemoveDefaultConstructorVisitor
	replaceBridgeMethodVisitor      *UpdateBridgeMethodVisitor
	initInnerClassStep2Visitor      *UpdateNewExpressionVisitor
	addCastExpressionVisitor        *AddCastExpressionVisitor

	typeDeclaration intmod.ITypeDeclaration
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration := declaration.(intsrv.IClassFileBodyDeclaration)

	// Visit inner types
	if bodyDeclaration.InnerTypeDeclarations() != nil {
		td := v.typeDeclaration
		tmp := make([]intmod.IDeclaration, 0)
		for _, item := range bodyDeclaration.InnerTypeDeclarations() {
			tmp = append(tmp, item)
		}
		v.AcceptListDeclaration(tmp)
		v.typeDeclaration = td
	}

	// Init bindTypeArgumentVisitor
	v.initStaticFieldVisitor.SetInternalTypeName(v.typeDeclaration.InternalTypeName())

	// Visit declaration
	v.initInnerClassStep2Visitor.VisitBodyDeclaration(declaration)
	v.initStaticFieldVisitor.VisitBodyDeclaration(declaration)
	v.initInstanceFieldVisitor.VisitBodyDeclaration(declaration)
	v.removeDefaultConstructorVisitor.VisitBodyDeclaration(declaration)
	aggregateFieldsVisitor.VisitBodyDeclaration(declaration)
	sortMembersVisitor.VisitBodyDeclaration(declaration)

	if bodyDeclaration.OuterBodyDeclaration() == nil {
		// Main body declaration

		if (bodyDeclaration.InnerTypeDeclarations() != nil) && v.replaceBridgeMethodVisitor.Init(bodyDeclaration) {
			// Replace bridge method invocation
			v.replaceBridgeMethodVisitor.VisitBodyDeclaration(bodyDeclaration)
		}

		// Add cast expressions
		v.addCastExpressionVisitor.VisitBodyDeclaration(declaration)

		// Autoboxing
		autoboxingVisitor.VisitBodyDeclaration(declaration)
	}
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitAnnotationDeclaration(declaration intmod.IAnnotationDeclaration) {
	v.typeDeclaration = declaration
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.typeDeclaration = declaration
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.typeDeclaration = declaration
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateJavaSyntaxTreeStep2Visitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.typeDeclaration = declaration

	// Remove 'static', 'final' and 'abstract' flags
	cfed := declaration.(intsrv.IClassFileEnumDeclaration)

	cfed.SetFlags(cfed.Flags() & ^(intmod.FlagStatic | intmod.FlagFinal | intmod.FlagAbstract))
	cfed.BodyDeclaration().Accept(v)
	v.initEnumVisitor.VisitBodyDeclaration(cfed.BodyDeclaration())
	cfed.SetConstants(v.initEnumVisitor.Constants().ToSlice())
}
