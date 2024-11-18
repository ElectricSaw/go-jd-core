package declaration

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"

type AbstractNopDeclarationVisitor struct {
}

func (v *AbstractNopDeclarationVisitor) VisitAnnotationDeclaration(declaration intsyn.IAnnotationDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitArrayVariableInitializer(declaration intsyn.IArrayVariableInitializer) {
}

func (v *AbstractNopDeclarationVisitor) VisitBodyDeclaration(declaration intsyn.IBodyDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitClassDeclaration(declaration intsyn.IClassDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitConstructorDeclaration(declaration intsyn.IConstructorDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitEnumDeclaration(declaration intsyn.IEnumDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitEnumDeclarationConstant(declaration intsyn.IConstant) {
}

func (v *AbstractNopDeclarationVisitor) VisitExpressionVariableInitializer(declaration intsyn.IExpressionVariableInitializer) {
}

func (v *AbstractNopDeclarationVisitor) VisitFieldDeclaration(declaration intsyn.IFieldDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitFieldDeclarator(declaration intsyn.IFieldDeclarator) {}

func (v *AbstractNopDeclarationVisitor) VisitFieldDeclarators(declarations intsyn.IFieldDeclarators) {
}

func (v *AbstractNopDeclarationVisitor) VisitFormalParameter(declaration intsyn.IFormalParameter) {}

func (v *AbstractNopDeclarationVisitor) VisitFormalParameters(declarations intsyn.IFormalParameters) {
}

func (v *AbstractNopDeclarationVisitor) VisitInstanceInitializerDeclaration(declaration intsyn.IInstanceInitializerDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitInterfaceDeclaration(declaration intsyn.IInterfaceDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitLocalVariableDeclaration(declaration intsyn.ILocalVariableDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitLocalVariableDeclarator(declarator intsyn.ILocalVariableDeclarator) {
}

func (v *AbstractNopDeclarationVisitor) VisitLocalVariableDeclarators(declarators intsyn.ILocalVariableDeclarators) {
}

func (v *AbstractNopDeclarationVisitor) VisitMethodDeclaration(declaration intsyn.IMethodDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitMemberDeclarations(declarations intsyn.IMemberDeclarations) {
}

func (v *AbstractNopDeclarationVisitor) VisitModuleDeclaration(declarations intsyn.IModuleDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitStaticInitializerDeclaration(declaration intsyn.IStaticInitializerDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitTypeDeclarations(declarations intsyn.ITypeDeclarations) {
}
