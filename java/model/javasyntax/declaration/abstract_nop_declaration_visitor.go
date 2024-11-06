package declaration

type AbstractNopDeclarationVisitor struct {
}

func (v *AbstractNopDeclarationVisitor) VisitAnnotationDeclaration(declaration *AnnotationDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitArrayVariableInitializer(declaration *ArrayVariableInitializer) {
}

func (v *AbstractNopDeclarationVisitor) VisitBodyDeclaration(declaration *BodyDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitClassDeclaration(declaration *ClassDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitConstructorDeclaration(declaration *ConstructorDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitEnumDeclaration(declaration *EnumDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitEnumDeclarationConstant(declaration *Constant) {
}

func (v *AbstractNopDeclarationVisitor) VisitExpressionVariableInitializer(declaration *ExpressionVariableInitializer) {
}

func (v *AbstractNopDeclarationVisitor) VisitFieldDeclaration(declaration *FieldDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitFieldDeclarator(declaration *FieldDeclarator) {}

func (v *AbstractNopDeclarationVisitor) VisitFieldDeclarators(declarations *FieldDeclarators) {}

func (v *AbstractNopDeclarationVisitor) VisitFormalParameter(declaration *FormalParameter) {}

func (v *AbstractNopDeclarationVisitor) VisitFormalParameters(declarations *FormalParameters) {}

func (v *AbstractNopDeclarationVisitor) VisitInstanceInitializerDeclaration(declaration *InstanceInitializerDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitInterfaceDeclaration(declaration *InterfaceDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitLocalVariableDeclaration(declaration *LocalVariableDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitLocalVariableDeclarator(declarator *LocalVariableDeclarator) {
}

func (v *AbstractNopDeclarationVisitor) VisitLocalVariableDeclarators(declarators *LocalVariableDeclarators) {
}

func (v *AbstractNopDeclarationVisitor) VisitMethodDeclaration(declaration *MethodDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitMemberDeclarations(declarations *MemberDeclarations) {}

func (v *AbstractNopDeclarationVisitor) VisitModuleDeclaration(declarations *ModuleDeclaration) {}

func (v *AbstractNopDeclarationVisitor) VisitStaticInitializerDeclaration(declaration *StaticInitializerDeclaration) {
}

func (v *AbstractNopDeclarationVisitor) VisitTypeDeclarations(declarations *TypeDeclarations) {}
