package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

type AbstractNopDeclarationVisitor struct {
}

func (v *AbstractNopDeclarationVisitor) VisitAnnotationDeclaration(declaration intmod.IAnnotationDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitArrayVariableInitializer(declaration intmod.IArrayVariableInitializer) {
}

func (v *AbstractNopDeclarationVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitConstructorDeclaration(declaration intmod.IConstructorDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitEnumDeclarationConstant(declaration intmod.IConstant) {
}

func (v *AbstractNopDeclarationVisitor) VisitExpressionVariableInitializer(declaration intmod.IExpressionVariableInitializer) {
}

func (v *AbstractNopDeclarationVisitor) VisitFieldDeclaration(declaration intmod.IFieldDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitFieldDeclarator(declaration intmod.IFieldDeclarator) {}

func (v *AbstractNopDeclarationVisitor) VisitFieldDeclarators(declarations intmod.IFieldDeclarators) {
}

func (v *AbstractNopDeclarationVisitor) VisitFormalParameter(declaration intmod.IFormalParameter) {}

func (v *AbstractNopDeclarationVisitor) VisitFormalParameters(declarations intmod.IFormalParameters) {
}

func (v *AbstractNopDeclarationVisitor) VisitInstanceInitializerDeclaration(declaration intmod.IInstanceInitializerDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitLocalVariableDeclaration(declaration intmod.ILocalVariableDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitLocalVariableDeclarator(declarator intmod.ILocalVariableDeclarator) {
}

func (v *AbstractNopDeclarationVisitor) VisitLocalVariableDeclarators(declarators intmod.ILocalVariableDeclarators) {
}

func (v *AbstractNopDeclarationVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitMemberDeclarations(declarations intmod.IMemberDeclarations) {
}

func (v *AbstractNopDeclarationVisitor) VisitModuleDeclaration(declarations intmod.IModuleDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitStaticInitializerDeclaration(declaration intmod.IStaticInitializerDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitTypeDeclarations(declarations intmod.ITypeDeclarations) {
}
