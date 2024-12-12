package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/class/service/converter/visitor/utils"
)

func NewSortMembersVisitor() intsrv.ISortMembersVisitor {
	return &SortMembersVisitor{}
}

type SortMembersVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor
}

func (v *SortMembersVisitor) VisitAnnotationDeclaration(declaration intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *SortMembersVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration := declaration.(intsrv.IClassFileBodyDeclaration)
	innerTypes := bodyDeclaration.InnerTypeDeclarations()
	// Merge fields, getters & inner types
	members := utils.Merge(utils.ConvertField(bodyDeclaration.FieldDeclarations()),
		utils.ConvertMethod(bodyDeclaration.MethodDeclarations()), utils.ConvertTypes(innerTypes))
	bodyDeclaration.SetMemberDeclarations(members)
}

func (v *SortMembersVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *SortMembersVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *SortMembersVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func Merge(fields, methods, innerTypes []intsrv.IClassFileMemberDeclaration) intmod.IMemberDeclarations {
	var size int

	if fields != nil {
		size = len(fields)
	} else {
		size = 0
	}

	if methods != nil {
		size += len(methods)
	}

	if innerTypes != nil {
		size += len(innerTypes)
	}

	result := declaration.NewMemberDeclarationsWithCapacity(size)

	tmp := make([]intmod.IMemberDeclaration, 0)
	for _, item := range fields {
		tmp = append(tmp, item.(intmod.IMemberDeclaration))
	}
	for _, item := range methods {
		tmp = append(tmp, item.(intmod.IMemberDeclaration))
	}
	for _, item := range innerTypes {
		tmp = append(tmp, item.(intmod.IMemberDeclaration))
	}
	result.AddAll(tmp)

	//result.AddAll(ConvertMemberDeclaration(fields))
	//result.AddAll(ConvertMemberDeclaration(methods))
	//result.AddAll(ConvertMemberDeclaration(innerTypes))

	return result
}
