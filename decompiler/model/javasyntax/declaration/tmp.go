package declaration

const (
	// Access flags for Class, Field, Method, Nested class, Module, Module Requires, Module Exports, Module Opens
	FlagPublic       = 0x0001 // C  F  M  N  .  .  .  .
	FlagPrivate      = 0x0002 // .  F  M  N  .  .  .  .
	FlagProtected    = 0x0004 // .  F  M  N  .  .  .  .
	FlagStatic       = 0x0008 // C  F  M  N  .  .  .  .
	FlagFinal        = 0x0010 // C  F  M  N  .  .  .  .
	FlagSynchronized = 0x0020 // .  .  M  .  .  .  .  .
	FlagSuper        = 0x0020 // C  .  .  .  .  .  .  .
	FlagOpen         = 0x0020 // .  .  .  .  Mo .  .  .
	FlagTransitive   = 0x0020 // .  .  .  .  .  MR .  .
	FlagVolatile     = 0x0040 // .  F  .  .  .  .  .  .
	FlagBridge       = 0x0040 // .  .  M  .  .  .  .  .
	FlagStaticPhase  = 0x0040 // .  .  .  .  .  MR .  .
	FlagTransient    = 0x0080 // .  F  .  .  .  .  .  .
	FlagVarArgs      = 0x0080 // .  .  M  .  .  .  .  .
	FlagNative       = 0x0100 // .  .  M  .  .  .  .  .
	FlagInterface    = 0x0200 // C  .  .  N  .  .  .  .
	FlagAnonymous    = 0x0200 // .  .  M  .  .  .  .  . // Custom flag
	FlagAbstract     = 0x0400 // C  .  M  N  .  .  .  .
	FlagStrict       = 0x0800 // .  .  M  .  .  .  .  .
	FlagSynthetic    = 0x1000 // C  F  M  N  Mo MR ME MO
	FlagAnnotation   = 0x2000 // C  .  .  N  .  .  .  .
	FlagEnum         = 0x4000 // C  F  .  N  .  .  .  .
	FlagModule       = 0x8000 // C  .  .  .  .  .  .  .
	FlagMandated     = 0x8000 // .  .  .  .  Mo MR ME MO

	// Extension
	FlagDefault = 0x10000 // .  .  M  .  .  .  .  .
)

type IDeclaration interface {
	AcceptDeclaration(visitor IDeclarationVisitor)
}

type IDeclarationVisitor interface {
	VisitAnnotationDeclaration(declaration *AnnotationDeclaration)
	VisitArrayVariableInitializer(declaration *ArrayVariableInitializer)
	VisitBodyDeclaration(declaration *BodyDeclaration)
	VisitClassDeclaration(declaration *ClassDeclaration)
	VisitConstructorDeclaration(declaration *ConstructorDeclaration)
	VisitEnumDeclaration(declaration *EnumDeclaration)
	VisitEnumDeclarationConstant(declaration *Constant)
	VisitExpressionVariableInitializer(declaration *ExpressionVariableInitializer)
	VisitFieldDeclaration(declaration *FieldDeclaration)
	VisitFieldDeclarator(declaration *FieldDeclarator)
	VisitFieldDeclarators(declarations *FieldDeclarators)
	VisitFormalParameter(declaration *FormalParameter)
	VisitFormalParameters(declarations *FormalParameters)
	VisitInstanceInitializerDeclaration(declaration *InstanceInitializerDeclaration)
	VisitInterfaceDeclaration(declaration *InterfaceDeclaration)
	VisitLocalVariableDeclaration(declaration *LocalVariableDeclaration)
	VisitLocalVariableDeclarator(declarator *LocalVariableDeclarator)
	VisitLocalVariableDeclarators(declarators *LocalVariableDeclarators)
	VisitMethodDeclaration(declaration *MethodDeclaration)
	VisitMemberDeclarations(declarations *MemberDeclarations)
	VisitModuleDeclaration(declarations *ModuleDeclaration)
	VisitStaticInitializerDeclaration(declaration *StaticInitializerDeclaration)
	VisitTypeDeclarations(declarations *TypeDeclarations)
}

type IMemberDeclaration interface {
	IsClassDeclaration() bool
	AcceptDeclaration(visitor IDeclarationVisitor)
}

type ILocalVariableDeclarator interface {
	LineNumber() int
	AcceptDeclaration(visitor IDeclarationVisitor)
}

type IVariableInitializer interface {
	LineNumber() int
	IsExpressionVariableInitializer() bool
	Expression() IExpression
	AcceptDeclaration(visitor IDeclarationVisitor)
}
