package visitor

import (
	"fmt"
	"strings"

	"github.com/ElectricSaw/go-jd-core/decompiler/api"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javafragment"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/token"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/fragmenter/visitor/fragutil"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewCompilationUnitVisitor(loader api.Loader, mainInternalTypeName string, majorVersion int,
	importsFragment intmod.IImportsFragment) *CompilationUnitVisitor {
	v := &CompilationUnitVisitor{
		StatementVisitor:           *NewStatementVisitor(loader, mainInternalTypeName, majorVersion, importsFragment),
		singleLineStatementVisitor: NewSingleLineStatementVisitor(),
		mainInternalName:           mainInternalTypeName,
	}
	v.annotationVisitor = NewAnnotationVisitor(v)
	return v
}

type CompilationUnitVisitor struct {
	StatementVisitor

	annotationVisitor          *AnnotationVisitor
	singleLineStatementVisitor *SingleLineStatementVisitor
	mainInternalName           string
}

func (v *CompilationUnitVisitor) VisitAnnotationDeclaration(declaration intmod.IAnnotationDeclaration) {
	if (declaration.Flags() & intmod.FlagSynthetic) == 0 {
		v.fragments.Add(javafragment.StartMovableTypeBlock)

		v.buildFragmentsForTypeDeclaration(declaration, declaration.Flags() & ^intmod.FlagAbstract, token.Annotation)

		v.fragments.AddTokensFragment(v.tokens)

		annotationDeclaratorList := declaration.AnnotationDeclarators()
		bodyDeclaration := declaration.BodyDeclaration()

		if (annotationDeclaratorList == nil) && (bodyDeclaration == nil) {
			v.tokens.Add(token.Space)
			v.tokens.Add(token.LeftRightCurlyBrackets)
		} else {
			fragmentCount1 := v.fragments.Size()
			start := fragutil.AddStartTypeBody(v.fragments)
			fragmentCount2 := v.fragments.Size()

			v.storeContext()
			v.currentInternalTypeName = declaration.InternalTypeName()
			v.currentTypeName = declaration.Name()

			if annotationDeclaratorList != nil {
				annotationDeclaratorList.AcceptDeclaration(v)

				if bodyDeclaration != nil {
					fragutil.AddSpacerBetweenMembers(v.fragments)
				}
			}

			v.SafeAcceptDeclaration(bodyDeclaration)

			v.restoreContext()

			if fragmentCount2 == v.fragments.Size() {
				v.fragments.SubList(fragmentCount1, fragmentCount2).Clear()
				v.tokens.Add(token.Space)
				v.tokens.Add(token.LeftRightCurlyBrackets)
			} else {
				fragutil.AddEndTypeBody(v.fragments, start)
			}
		}

		v.fragments.Add(javafragment.EndMovableBlock)
	}
}

func (v *CompilationUnitVisitor) VisitAnnotationElementValue(reference intmod.IAnnotationElementValue) {
	v.VisitAnnotationReference(reference)
}

func (v *CompilationUnitVisitor) VisitAnnotationReference(reference intmod.IAnnotationReference) {
	v.visitAnnotationReferenceAnnotationReference(reference)
}

func (v *CompilationUnitVisitor) visitAnnotationReferenceAnnotationReference(reference intmod.IAnnotationReference) {
	v.tokens.Add(token.At)

	typ := reference.Type()
	typ.AcceptTypeVisitor(v)

	elementValue := reference.ElementValue()

	if elementValue == nil {
		elementValuePairs := reference.ElementValuePairs()

		if elementValuePairs != nil {
			v.tokens.Add(token.StartParametersBlock)
			elementValuePairs.Accept(v)
			v.tokens.Add(token.EndParametersBlock)
		}
	} else {
		v.tokens.Add(token.StartParametersBlock)
		elementValue.Accept(v)
		v.tokens.Add(token.EndParametersBlock)
	}
}

func (v *CompilationUnitVisitor) VisitAnnotationReferences(list intmod.IAnnotationReferences) {
	size := list.Size()

	if size > 0 {
		iterator := list.Iterator()
		iterator.Next().Accept(v)

		for i := 1; i < size; i++ {
			v.tokens.Add(token.Space)
			iterator.Next().Accept(v)
		}
	}
}

func (v *CompilationUnitVisitor) VisitArrayVariableInitializer(declaration intmod.IArrayVariableInitializer) {
	size := declaration.Size()

	if size > 0 {
		v.fragments.AddTokensFragment(v.tokens)

		start := fragutil.AddStartArrayInitializerBlock(v.fragments)

		if size > 10 {
			fragutil.AddNewLineBetweenArrayInitializerBlock(v.fragments)
		}

		v.tokens = NewTokens(v)

		declaration.Get(0).AcceptDeclaration(v)

		for i := 1; i < size; i++ {
			if v.tokens.IsEmpty() {
				fragutil.AddSpacerBetweenArrayInitializerBlock(v.fragments)

				if (size > 10) && (i%10 == 0) {
					fragutil.AddNewLineBetweenArrayInitializerBlock(v.fragments)
				}
			} else if (size > 10) && (i%10 == 0) {
				v.fragments.AddTokensFragment(v.tokens)

				fragutil.AddSpacerBetweenArrayInitializerBlock(v.fragments)
				fragutil.AddNewLineBetweenArrayInitializerBlock(v.fragments)

				v.tokens = NewTokens(v)
			} else {
				v.tokens.Add(token.CommaSpace)
			}

			declaration.Get(i).AcceptDeclaration(v)
		}

		v.fragments.AddTokensFragment(v.tokens)

		if v.inExpressionFlag {
			fragutil.AddEndArrayInitializerInParameter(v.fragments, start)
		} else {
			fragutil.AddEndArrayInitializer(v.fragments, start)
		}

		v.tokens = NewTokens(v)
	} else {
		v.tokens.Add(token.LeftRightCurlyBrackets)
	}
}

func (v *CompilationUnitVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	v.SafeAcceptDeclaration(declaration.MemberDeclarations())
}

func (v *CompilationUnitVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	if (declaration.Flags() & intmod.FlagSynthetic) == 0 {
		v.fragments.Add(javafragment.StartMovableTypeBlock)

		v.buildFragmentsForClassOrInterfaceDeclaration(declaration, declaration.Flags(), token.Class2)

		v.tokens.Add(token.StartDeclarationOrStatementBlock)

		// Build v.fragments for super type
		superType := declaration.SuperType()
		if superType != nil && superType != _type.OtTypeObject {
			v.fragments.AddTokensFragment(v.tokens)
			fragutil.AddSpacerBeforeExtends(v.fragments)

			v.tokens = NewTokens(v)
			v.tokens.Add(token.Exports)
			v.tokens.Add(token.Space)
			superType.AcceptTypeVisitor(v)
			v.fragments.AddTokensFragment(v.tokens)

			v.tokens = NewTokens(v)
		}

		// Build v.fragments for interfaces
		interfaces := declaration.Interfaces()
		if interfaces != nil {
			if !v.tokens.IsEmpty() {
				v.fragments.AddTokensFragment(v.tokens)
			}

			fragutil.AddSpacerBeforeImplements(v.fragments)

			v.tokens = NewTokens(v)
			v.tokens.Add(token.Implements)
			v.tokens.Add(token.Space)
			interfaces.AcceptTypeVisitor(v)
			v.fragments.AddTokensFragment(v.tokens)

			v.tokens = NewTokens(v)
		}

		v.tokens.Add(token.EndDeclarationOrStatementBlock)
		v.fragments.AddTokensFragment(v.tokens)

		bodyDeclaration := declaration.BodyDeclaration()

		if bodyDeclaration == nil {
			v.tokens.Add(token.Space)
			v.tokens.Add(token.LeftRightCurlyBrackets)
		} else {
			fragmentCount1 := v.fragments.Size()
			start := fragutil.AddStartTypeBody(v.fragments)
			fragmentCount2 := v.fragments.Size()

			v.storeContext()
			v.currentInternalTypeName = declaration.InternalTypeName()
			v.currentTypeName = declaration.Name()
			bodyDeclaration.AcceptDeclaration(v)
			v.restoreContext()

			if fragmentCount2 == v.fragments.Size() {
				v.fragments.SubList(fragmentCount1, fragmentCount2).Clear()
				v.tokens.Add(token.Space)
				v.tokens.Add(token.LeftRightCurlyBrackets)
			} else {
				fragutil.AddEndTypeBody(v.fragments, start)
			}
		}

		v.fragments.Add(javafragment.EndMovableBlock)
	}
}

func (v *CompilationUnitVisitor) VisitCompilationUnit(compilationUnit *javasyntax.CompilationUnit) {
	// Init
	v.fragments.Clear()
	v.contextStack.Clear()
	v.currentInternalTypeName = ""

	// Add fragment for package
	index := strings.LastIndex(v.mainInternalName, "/")
	if index != -1 {
		v.tokens = NewTokens(v)

		v.tokens.Add(token.Package)
		v.tokens.Add(token.Space)
		v.tokens.Add(token.NewTextToken(strings.ReplaceAll(v.mainInternalName[:index], "/", ".")))
		v.tokens.Add(token.Semicolon)

		v.fragments.AddTokensFragment(v.tokens)

		fragutil.AddSpacerAfterPackage(v.fragments)
	}

	// Add fragment for imports
	if !v.importsFragment.IsEmpty() {
		v.fragments.Add(v.importsFragment)

		if !v.fragments.IsEmpty() {
			fragutil.AddSpacerAfterImports(v.fragments)
		}
	}

	fragutil.AddSpacerBeforeMainDeclaration(v.fragments)

	// Visit all compilation unit
	v.AbstractJavaSyntaxVisitor.VisitCompilationUnit(compilationUnit)
}

func (v *CompilationUnitVisitor) VisitConstructorDeclaration(declaration intmod.IConstructorDeclaration) {
	if (declaration.Flags() & (intmod.FlagSynthetic | intmod.FlagBridge)) == 0 {
		statements := declaration.Statements()

		if (declaration.Flags() & intmod.FlagAnonymous) == 0 {
			v.fragments.Add(javafragment.StartMovableMethodBlock)

			v.tokens = NewTokens(v)

			// Build v.fragments for annotations
			annotationReferences := declaration.AnnotationReferences()

			if annotationReferences != nil {
				annotationReferences.Accept(v.annotationVisitor)
				v.fragments.AddTokensFragment(v.tokens)
				fragutil.AddSpacerAfterMemberAnnotations(v.fragments)
				v.tokens = NewTokens(v)
			}

			// Build v.tokens for access
			v.buildTokensForMethodAccessFlags(declaration.Flags())

			// Build v.tokens for type parameters
			typeParameters := declaration.TypeParameters()

			if typeParameters != nil {
				v.tokens.Add(token.LeftAngleBracket)
				typeParameters.AcceptTypeParameterVisitor(v)
				v.tokens.Add(token.RightAngleBracket)
				v.tokens.Add(token.Space)
			}

			// Build token for type declaration
			v.tokens.Add(token.NewDeclarationToken(intmod.ConstructorToken, v.currentInternalTypeName, v.currentTypeName, declaration.Descriptor()))

			v.storeContext()
			v.currentMethodParamNames.Clear()

			formalParameters := declaration.FormalParameters()

			if formalParameters == nil {
				v.tokens.Add(token.LeftRightRoundBrackets)
			} else {
				v.tokens.Add(token.StartParametersBlock)
				v.fragments.AddTokensFragment(v.tokens)

				formalParameters.AcceptDeclaration(v)

				v.tokens = NewTokens(v)
				v.tokens.Add(token.EndParametersBlock)
			}

			exceptionTypes := declaration.ExceptionTypes()

			if exceptionTypes != nil {
				v.tokens.Add(token.Space)
				v.tokens.Add(token.Throws)
				v.tokens.Add(token.Space)
				exceptionTypes.AcceptTypeVisitor(v)
			}

			if statements != nil {
				v.fragments.AddTokensFragment(v.tokens)
				v.singleLineStatementVisitor.Init()
				statements.AcceptStatement(v.singleLineStatementVisitor)

				singleLineStatement := v.singleLineStatementVisitor.IsSingleLineStatement()
				fragmentCount1 := v.fragments.Size()
				var start intmod.IStartBodyFragment

				if singleLineStatement {
					start = fragutil.AddStartSingleStatementMethodBody(v.fragments)
				} else {
					start = fragutil.AddStartMethodBody(v.fragments)
				}

				fragmentCount2 := v.fragments.Size()

				statements.AcceptStatement(v)

				if fragmentCount2 == v.fragments.Size() {
					v.fragments.SubList(fragmentCount1, fragmentCount2).Clear()
					v.tokens.Add(token.Space)
					v.tokens.Add(token.LeftRightCurlyBrackets)
				} else if singleLineStatement {
					fragutil.AddEndSingleStatementMethodBody(v.fragments, start)
				} else {
					fragutil.AddEndMethodBody(v.fragments, start)
				}
			}

			v.restoreContext()

			v.fragments.Add(javafragment.EndMovableBlock)
		} else if (statements != nil) && (statements.Size() > 0) {
			fragmentCount0 := v.fragments.Size()
			v.fragments.Add(javafragment.StartMovableMethodBlock)

			start := fragutil.AddStartInstanceInitializerBlock(v.fragments)
			fragmentCount2 := v.fragments.Size()

			v.tokens = NewTokens(v)

			statements.AcceptStatement(v)

			if fragmentCount2 == v.fragments.Size() {
				v.fragments.SubList(fragmentCount0, fragmentCount2).Clear()
			} else {
				fragutil.AddEndInstanceInitializerBlock(v.fragments, start)
			}

			v.fragments.Add(javafragment.EndMovableBlock)
		}
	}
}

func (v *CompilationUnitVisitor) VisitElementValueArrayInitializerElementValue(reference intmod.IElementValueArrayInitializerElementValue) {
	v.tokens.Add(token.StartArrayInitializerBlock)
	v.SafeAcceptReference(reference.ElementValueArrayInitializer())
	v.tokens.Add(token.EndArrayInitializerBlock)
}

func (v *CompilationUnitVisitor) VisitElementValues(references intmod.IElementValues) {
	iterator := references.Iterator()

	iterator.Next().Accept(v)
	for iterator.HasNext() {
		v.tokens.Add(token.CommaSpace)
		iterator.Next().Accept(v)
	}
}

func (v *CompilationUnitVisitor) VisitExpressionElementValue(reference intmod.IExpressionElementValue) {
	reference.Expression().Accept(v)
}

func (v *CompilationUnitVisitor) VisitElementValuePairs(references intmod.IElementValuePairs) {
	iterator := references.Iterator()

	iterator.Next().Accept(v)

	for iterator.HasNext() {
		v.tokens.Add(token.CommaSpace)
		iterator.Next().Accept(v)
	}
}

func (v *CompilationUnitVisitor) VisitElementValuePair(reference intmod.IElementValuePair) {
	v.tokens.Add(token.NewTextToken(reference.Name()))
	v.tokens.Add(token.SpaceEqualSpace)
	reference.ElementValue().Accept(v)
}

func (v *CompilationUnitVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	if (declaration.Flags() & intmod.FlagSynthetic) == 0 {
		v.fragments.Add(javafragment.StartMovableTypeBlock)

		v.buildFragmentsForTypeDeclaration(declaration, declaration.Flags(), token.Enum)

		// Build v.fragments for interfaces
		interfaces := declaration.Interfaces()
		if interfaces != nil {
			v.tokens.Add(token.StartDeclarationOrStatementBlock)

			v.fragments.AddTokensFragment(v.tokens)

			fragutil.AddSpacerBeforeImplements(v.fragments)

			v.tokens = NewTokens(v)
			v.tokens.Add(token.Implements)
			v.tokens.Add(token.Space)
			interfaces.AcceptTypeVisitor(v)
			v.fragments.AddTokensFragment(v.tokens)

			v.tokens = NewTokens(v)

			v.tokens.Add(token.EndDeclarationOrStatementBlock)
		}

		v.fragments.AddTokensFragment(v.tokens)

		start := fragutil.AddStartTypeBody(v.fragments)

		v.storeContext()
		v.currentInternalTypeName = declaration.InternalTypeName()
		v.currentTypeName = declaration.Name()

		constants := util.NewDefaultListWithSlice(declaration.Constants())

		if (constants != nil) && (!constants.IsEmpty()) {
			preferredLineNumber := 0

			for _, constant := range constants.ToSlice() {
				if (constant.Arguments() != nil) || (constant.BodyDeclaration() != nil) {
					preferredLineNumber = 1
					break
				}
			}

			v.fragments.Add(javafragment.StartMovableFieldBlock)

			constants.Get(0).AcceptDeclaration(v)

			size := constants.Size()

			for i := 1; i < size; i++ {
				fragutil.AddSpacerBetweenEnumValues(v.fragments, preferredLineNumber)
				constants.Get(i).AcceptDeclaration(v)
			}

			v.fragments.Add(javafragment.Semicolon)
			v.fragments.Add(javafragment.EndMovableBlock)
		}

		bodyDeclaration := declaration.BodyDeclaration()

		if bodyDeclaration != nil {
			if (constants != nil) && (!constants.IsEmpty()) {
				f := v.fragments

				v.fragments = NewFragments()
				bodyDeclaration.AcceptDeclaration(v)

				if !v.fragments.IsEmpty() {
					fragutil.AddSpacerBetweenMembers(f)
					f.AddAll(v.fragments.ToSlice())
				}

				v.fragments = f
			} else {
				bodyDeclaration.AcceptDeclaration(v)
			}
		}

		v.restoreContext()

		fragutil.AddEndTypeBody(v.fragments, start)

		v.fragments.Add(javafragment.EndMovableBlock)
	}
}

func (v *CompilationUnitVisitor) VisitEnumDeclarationConstant(declaration intmod.IConstant) {
	v.tokens = NewTokens(v)

	// Build v.fragments for annotations
	annotationReferences := declaration.AnnotationReferences()

	if annotationReferences != nil {
		annotationReferences.Accept(v.annotationVisitor)
		v.fragments.AddTokensFragment(v.tokens)
		fragutil.AddSpacerAfterMemberAnnotations(v.fragments)
		v.tokens = NewTokens(v)
	}

	// Build token for type declaration
	v.tokens.AddLineNumberTokenAt(declaration.LineNumber())
	v.tokens.Add(token.NewDeclarationToken(intmod.FieldToken, v.currentInternalTypeName,
		declaration.Name(), fmt.Sprintf("L%s;", v.currentInternalTypeName)))

	v.storeContext()
	v.currentMethodParamNames.Clear()

	arguments := declaration.Arguments()
	if arguments != nil {
		v.tokens.Add(token.StartParametersBlock)
		arguments.Accept(v)
		v.tokens.Add(token.EndParametersBlock)
	}

	v.fragments.AddTokensFragment(v.tokens)

	bodyDeclaration := declaration.BodyDeclaration()

	if bodyDeclaration != nil {
		start := fragutil.AddStartTypeBody(v.fragments)
		bodyDeclaration.AcceptDeclaration(v)
		fragutil.AddEndTypeBody(v.fragments, start)
	}

	v.restoreContext()
}

func (v *CompilationUnitVisitor) VisitFieldDeclaration(declaration intmod.IFieldDeclaration) {
	if (declaration.Flags() & intmod.FlagSynthetic) == 0 {
		v.fragments.Add(javafragment.StartMovableFieldBlock)

		v.tokens = NewTokens(v)

		// Build v.fragments for annotations
		annotationReferences := declaration.AnnotationReferences()

		if annotationReferences != nil {
			annotationReferences.Accept(v.annotationVisitor)
			v.fragments.AddTokensFragment(v.tokens)
			fragutil.AddSpacerAfterMemberAnnotations(v.fragments)
			v.tokens = NewTokens(v)
		}

		// Build v.tokens for access
		v.buildTokensForFieldAccessFlags(declaration.Flags())

		typ := declaration.Type()

		typ.AcceptTypeVisitor(v)

		v.tokens.Add(token.StartDeclarationOrStatementBlock)
		v.fragments.AddTokensFragment(v.tokens)

		declaration.FieldDeclarators().AcceptDeclaration(v)

		v.fragments.Add(javafragment.EndDeclarationOrStatementBlockSemicolon)

		v.fragments.Add(javafragment.EndMovableBlock)
	}
}

func (v *CompilationUnitVisitor) VisitFieldDeclarator(fieldDeclarator intmod.IFieldDeclarator) {
	fieldDeclaration := fieldDeclarator.FieldDeclaration()
	variableInitializer := fieldDeclarator.VariableInitializer()
	descriptor := fieldDeclaration.Type().Descriptor()

	v.tokens = NewTokens(v)
	v.tokens.Add(token.Space)

	switch fieldDeclarator.Dimension() {
	case 0:
		v.tokens.Add(token.NewDeclarationToken(intmod.FieldToken, v.currentInternalTypeName, fieldDeclarator.Name(), descriptor))
		break
	case 1:
		v.tokens.Add(token.NewDeclarationToken(intmod.FieldToken, v.currentInternalTypeName, fieldDeclarator.Name(), "["+descriptor))
		v.tokens.Add(token.Dimension1)
		break
	case 2:
		v.tokens.Add(token.NewDeclarationToken(intmod.FieldToken, v.currentInternalTypeName, fieldDeclarator.Name(), "[["+descriptor))
		v.tokens.Add(token.Dimension2)
		break
	default:
		prefix := ""
		for i := 0; i < fieldDeclarator.Dimension(); i++ {
			prefix += "["
		}
		descriptor = prefix + descriptor
		v.tokens.Add(token.NewDeclarationToken(intmod.FieldToken, v.currentInternalTypeName, fieldDeclarator.Name(), descriptor))
		prefix = ""
		for i := 0; i < fieldDeclarator.Dimension(); i++ {
			prefix += "[]"
		}
		v.tokens.Add(token.NewTextToken(prefix))
		break
	}

	if variableInitializer == nil {
		v.fragments.AddTokensFragment(v.tokens)
	} else {
		v.tokens.Add(token.SpaceEqualSpace)
		variableInitializer.AcceptDeclaration(v)
		v.fragments.AddTokensFragment(v.tokens)
	}
}

func (v *CompilationUnitVisitor) VisitFieldDeclarators(declarators intmod.IFieldDeclarators) {
	size := declarators.Size()

	if size > 0 {
		iterator := declarators.Iterator()
		iterator.Next().AcceptDeclaration(v)

		for i := 1; i < size; i++ {
			fragutil.AddSpacerBetweenFieldDeclarators(v.fragments)
			iterator.Next().AcceptDeclaration(v)
		}
	}
}

func (v *CompilationUnitVisitor) VisitFormalParameter(declaration intmod.IFormalParameter) {
	annotationReferences := declaration.AnnotationReferences()

	if annotationReferences != nil {
		annotationReferences.Accept(v)
		v.tokens.Add(token.Space)
	}

	if declaration.IsVarargs() {
		arrayType := declaration.Type()
		typ := arrayType.CreateType(arrayType.Dimension() - 1)
		typ.AcceptTypeVisitor(v)
		v.tokens.Add(token.VarArgs)
	} else {
		if declaration.IsFinal() {
			v.tokens.Add(token.Final)
			v.tokens.Add(token.Space)
		}

		typ := declaration.Type()
		typ.AcceptTypeVisitor(v)
		v.tokens.Add(token.Space)
	}

	name := declaration.Name()

	v.tokens.Add(token.NewTextToken(name))
	v.currentMethodParamNames.Add(name)
}

func (v *CompilationUnitVisitor) VisitFormalParameters(declarations intmod.IFormalParameters) {
	size := declarations.Size()

	if size > 0 {
		iterator := declarations.Iterator()
		iterator.Next().AcceptDeclaration(v)

		for i := 1; i < size; i++ {
			v.tokens.Add(token.CommaSpace)
			iterator.Next().AcceptDeclaration(v)
		}
	}
}

func (v *CompilationUnitVisitor) VisitInstanceInitializerDeclaration(declaration intmod.IInstanceInitializerDeclaration) {
	statements := declaration.Statements()

	if statements != nil {
		v.fragments.Add(javafragment.StartMovableMethodBlock)

		v.storeContext()
		v.currentMethodParamNames.Clear()

		start := fragutil.AddStartMethodBody(v.fragments)
		statements.AcceptStatement(v)
		fragutil.AddEndMethodBody(v.fragments, start)

		v.fragments.Add(javafragment.EndMovableBlock)

		v.restoreContext()
	}
}

func (v *CompilationUnitVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	if (declaration.Flags() & intmod.FlagSynthetic) == 0 {
		v.fragments.Add(javafragment.StartMovableTypeBlock)

		v.buildFragmentsForClassOrInterfaceDeclaration(declaration, declaration.Flags() & ^intmod.FlagAbstract, token.Interface)

		v.tokens.Add(token.StartDeclarationOrStatementBlock)

		// Build v.fragments for interfaces
		interfaces := declaration.Interfaces()
		if interfaces != nil {
			v.fragments.AddTokensFragment(v.tokens)

			fragutil.AddSpacerBeforeImplements(v.fragments)

			v.tokens = NewTokens(v)
			v.tokens.Add(token.Extends)
			v.tokens.Add(token.Space)
			interfaces.AcceptTypeVisitor(v)
			v.fragments.AddTokensFragment(v.tokens)

			v.tokens = NewTokens(v)
		}

		v.tokens.Add(token.EndDeclarationOrStatementBlock)
		v.fragments.AddTokensFragment(v.tokens)

		bodyDeclaration := declaration.BodyDeclaration()
		if bodyDeclaration == nil {
			v.tokens.Add(token.Space)
			v.tokens.Add(token.LeftRightCurlyBrackets)
		} else {
			fragmentCount1 := v.fragments.Size()
			start := fragutil.AddStartTypeBody(v.fragments)
			fragmentCount2 := v.fragments.Size()

			v.storeContext()
			v.currentInternalTypeName = declaration.InternalTypeName()
			v.currentTypeName = declaration.Name()
			bodyDeclaration.AcceptDeclaration(v)
			v.restoreContext()

			if fragmentCount2 == v.fragments.Size() {
				v.fragments.SubList(fragmentCount1, fragmentCount2).Clear()
				v.tokens.Add(token.Space)
				v.tokens.Add(token.LeftRightCurlyBrackets)
			} else {
				fragutil.AddEndTypeBody(v.fragments, start)
			}
		}

		v.fragments.Add(javafragment.EndMovableBlock)
	}
}

func (v *CompilationUnitVisitor) VisitModuleDeclaration(declaration intmod.IModuleDeclaration) {
	needNewLine := false

	v.fragments.Clear()
	v.fragments.Add(javafragment.StartMovableTypeBlock)

	v.tokens = NewTokens(v)

	if (declaration.Flags() & intmod.FlagOpen) != 0 {
		v.tokens.Add(token.Open)
		v.tokens.Add(token.Space)
	}

	v.tokens.Add(token.Module)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.NewDeclarationToken(intmod.ModuleToken, declaration.InternalTypeName(), declaration.Name(), ""))
	v.fragments.AddTokensFragment(v.tokens)

	start := fragutil.AddStartTypeBody(v.fragments)

	v.tokens = NewTokens(v)

	if (declaration.Requires() != nil) && !declaration.Requires().IsEmpty() {
		iterator := declaration.Requires().Iterator()
		v.visitModuleDeclaration1(iterator.Next())
		for iterator.HasNext() {
			v.tokens.Add(token.NewLine1)
			v.visitModuleDeclaration1(iterator.Next())
		}
		needNewLine = true
	}

	if (declaration.Exports() != nil) && !declaration.Exports().IsEmpty() {
		if needNewLine {
			v.tokens.Add(token.NewLine2)
		}
		iterator := declaration.Exports().Iterator()
		v.visitModuleDeclaration2(iterator.Next(), token.Exports)
		for iterator.HasNext() {
			v.tokens.Add(token.NewLine1)
			v.visitModuleDeclaration2(iterator.Next(), token.Exports)
		}
		needNewLine = true
	}

	if (declaration.Opens() != nil) && !declaration.Opens().IsEmpty() {
		if needNewLine {
			v.tokens.Add(token.NewLine2)
		}
		iterator := declaration.Opens().Iterator()
		v.visitModuleDeclaration2(iterator.Next(), token.Opens)
		for iterator.HasNext() {
			v.tokens.Add(token.NewLine1)
			v.visitModuleDeclaration2(iterator.Next(), token.Opens)
		}
		needNewLine = true
	}

	if (declaration.Uses() != nil) && !declaration.Uses().IsEmpty() {
		if needNewLine {
			v.tokens.Add(token.NewLine2)
		}
		iterator := declaration.Uses().Iterator()
		v.visitModuleDeclaration3(iterator.Next())
		for iterator.HasNext() {
			v.tokens.Add(token.NewLine1)
			v.visitModuleDeclaration3(iterator.Next())
		}
		needNewLine = true
	}

	if (declaration.Provides() != nil) && !declaration.Provides().IsEmpty() {
		if needNewLine {
			v.tokens.Add(token.NewLine2)
		}
		iterator := declaration.Provides().Iterator()
		v.visitModuleDeclaration4(iterator.Next())
		for iterator.HasNext() {
			v.tokens.Add(token.NewLine1)
			v.visitModuleDeclaration4(iterator.Next())
		}
	}

	v.fragments.AddTokensFragment(v.tokens)

	fragutil.AddEndTypeBody(v.fragments, start)

	v.fragments.Add(javafragment.EndMovableBlock)
}

func (v *CompilationUnitVisitor) visitModuleDeclaration1(moduleInfo intmod.IModuleInfo) {
	v.tokens.Add(token.Requires)

	if (moduleInfo.Flags() & intmod.FlagStatic) != 0 {
		v.tokens.Add(token.Space)
		v.tokens.Add(token.Static)
	}
	if (moduleInfo.Flags() & intmod.FlagTransitive) != 0 {
		v.tokens.Add(token.Space)
		v.tokens.Add(token.Transient)
	}

	v.tokens.Add(token.Space)
	v.tokens.Add(token.NewReferenceToken(intmod.ModuleToken, "module-info",
		moduleInfo.Name(), "", ""))
	v.tokens.Add(token.Semicolon)
}

func (v *CompilationUnitVisitor) visitModuleDeclaration2(packageInfo intmod.IPackageInfo, keywordToken intmod.IKeywordToken) {
	v.tokens.Add(keywordToken)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.NewReferenceToken(intmod.PackageToken, packageInfo.InternalName(),
		strings.ReplaceAll(packageInfo.InternalName(), "/", "."), "", ""))

	if (packageInfo.ModuleInfoNames() != nil) && len(packageInfo.ModuleInfoNames()) != 0 {
		moduleInfoNames := util.NewDefaultListWithSlice(packageInfo.ModuleInfoNames())
		v.tokens.Add(token.Space)
		v.tokens.Add(token.To)

		if moduleInfoNames.Size() == 1 {
			v.tokens.Add(token.Space)
			moduleInfoName := moduleInfoNames.Get(0)
			v.tokens.Add(token.NewReferenceToken(intmod.ModuleToken, "module-info",
				moduleInfoName, "", ""))
		} else {
			v.tokens.Add(token.StartDeclarationOrStatementBlock)
			v.tokens.Add(token.NewLine1)

			iterator := moduleInfoNames.Iterator()

			moduleInfoName := iterator.Next()
			v.tokens.Add(token.NewReferenceToken(intmod.ModuleToken, "module-info",
				moduleInfoName, "", ""))

			for iterator.HasNext() {
				v.tokens.Add(token.Comma)
				v.tokens.Add(token.NewLine1)
				moduleInfoName = iterator.Next()
				v.tokens.Add(token.NewReferenceToken(intmod.ModuleToken, "module-info",
					moduleInfoName, "", ""))
			}

			v.tokens.Add(token.EndDeclarationOrStatementBlock)
		}
	}

	v.tokens.Add(token.Semicolon)
}

func (v *CompilationUnitVisitor) visitModuleDeclaration3(internalTypeName string) {
	v.tokens.Add(token.Uses)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.NewReferenceToken(intmod.TypeToken, internalTypeName, strings.ReplaceAll(internalTypeName, "/", "."), "", ""))
	v.tokens.Add(token.Semicolon)
}

func (v *CompilationUnitVisitor) visitModuleDeclaration4(serviceInfo intmod.IServiceInfo) {
	v.tokens.Add(token.Provides)
	v.tokens.Add(token.Space)
	internalTypeName := serviceInfo.InternalTypeName()
	v.tokens.Add(token.NewReferenceToken(intmod.TypeToken, internalTypeName,
		strings.ReplaceAll(internalTypeName, "/", "."), "", ""))
	v.tokens.Add(token.Space)
	v.tokens.Add(token.With)

	if len(serviceInfo.ImplementationTypeNames()) == 1 {
		v.tokens.Add(token.Space)
		internalTypeName = serviceInfo.ImplementationTypeNames()[0]
		v.tokens.Add(token.NewReferenceToken(intmod.TypeToken, internalTypeName,
			strings.ReplaceAll(internalTypeName, "/", "."), "", ""))
	} else {
		v.tokens.Add(token.StartDeclarationOrStatementBlock)
		v.tokens.Add(token.NewLine1)

		iterator := util.NewIteratorWithSlice(serviceInfo.ImplementationTypeNames())

		internalTypeName = iterator.Next()
		v.tokens.Add(token.NewReferenceToken(intmod.TypeToken, internalTypeName,
			strings.ReplaceAll(internalTypeName, "/", "."), "", ""))

		for iterator.HasNext() {
			v.tokens.Add(token.Comma)
			v.tokens.Add(token.NewLine1)
			internalTypeName = iterator.Next()
			v.tokens.Add(token.NewReferenceToken(intmod.TypeToken, internalTypeName,
				strings.ReplaceAll(internalTypeName, "/", "."), "", ""))
		}

		v.tokens.Add(token.EndDeclarationOrStatementBlock)
	}

	v.tokens.Add(token.Semicolon)
}

func (v *CompilationUnitVisitor) VisitLocalVariableDeclaration(declaration intmod.ILocalVariableDeclaration) {
	if declaration.IsFinal() {
		v.tokens.Add(token.Final)
		v.tokens.Add(token.Space)
	}
	typ := declaration.Type()
	typ.AcceptTypeVisitor(v)
	v.tokens.Add(token.Space)
	declaration.LocalVariableDeclarators().AcceptDeclaration(v)
}

func (v *CompilationUnitVisitor) VisitLocalVariableDeclarator(declarator intmod.ILocalVariableDeclarator) {
	if declarator.VariableInitializer() == nil {
		v.tokens.AddLineNumberTokenAt(declarator.LineNumber())
		v.tokens.Add(token.NewTextToken(declarator.Name()))

		v.visitDimension(declarator.Dimension())
	} else {
		v.tokens.Add(token.NewTextToken(declarator.Name()))

		v.visitDimension(declarator.Dimension())

		v.tokens.Add(token.SpaceEqualSpace)
		declarator.VariableInitializer().AcceptDeclaration(v)
	}
}

func (v *CompilationUnitVisitor) VisitLocalVariableDeclarators(declarators intmod.ILocalVariableDeclarators) {
	size := declarators.Size()

	if size > 0 {
		iterator := declarators.Iterator()
		iterator.Next().AcceptDeclaration(v)

		for i := 1; i < size; i++ {
			v.tokens.Add(token.CommaSpace)
			iterator.Next().AcceptDeclaration(v)
		}
	}
}

func (v *CompilationUnitVisitor) VisitMemberDeclarations(list intmod.IMemberDeclarations) {
	size := list.Size()

	if size > 0 {
		fragmentCount2 := v.fragments.Size()
		iterator := list.Iterator()

		iterator.Next().AcceptDeclaration(v)

		if size > 1 {
			fragmentCount1 := -1

			for i := 1; i < size; i++ {
				if fragmentCount2 < v.fragments.Size() {
					fragmentCount1 = v.fragments.Size()
					fragutil.AddSpacerBetweenMembers(v.fragments)
					fragmentCount2 = v.fragments.Size()
				}
				iterator.Next().AcceptDeclaration(v)
			}

			if (fragmentCount1 != -1) && (fragmentCount2 == v.fragments.Size()) {
				v.fragments.SubList(fragmentCount1, v.fragments.Size()).Clear()
			}
		}
	}
}

func (v *CompilationUnitVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	if (declaration.Flags() & (intmod.FlagSynthetic | intmod.FlagBridge)) == 0 {
		v.fragments.Add(javafragment.StartMovableMethodBlock)

		v.tokens = NewTokens(v)

		// Build v.fragments for annotations
		annotationReferences := declaration.AnnotationReferences()

		if annotationReferences != nil {
			annotationReferences.Accept(v.annotationVisitor)
			v.fragments.AddTokensFragment(v.tokens)
			fragutil.AddSpacerAfterMemberAnnotations(v.fragments)
			v.tokens = NewTokens(v)
		}

		// Build v.tokens for access
		v.buildTokensForMethodAccessFlags(declaration.Flags())

		// Build v.tokens for type parameters
		typeParameters := declaration.TypeParameters()

		if typeParameters != nil {
			v.tokens.Add(token.LeftAngleBracket)
			typeParameters.AcceptTypeParameterVisitor(v)
			v.tokens.Add(token.RightAngleBracket)
			v.tokens.Add(token.Space)
		}

		returnedType := declaration.ReturnedType()
		returnedType.AcceptTypeVisitor(v)
		v.tokens.Add(token.Space)

		// Build token for type declaration
		v.tokens.Add(token.NewDeclarationToken(intmod.MethodToken, v.currentInternalTypeName, declaration.Name(), declaration.Descriptor()))

		v.storeContext()
		v.currentMethodParamNames.Clear()

		formalParameters := declaration.FormalParameters()

		if formalParameters == nil {
			v.tokens.Add(token.LeftRightRoundBrackets)
		} else {
			v.tokens.Add(token.StartParametersBlock)
			v.fragments.AddTokensFragment(v.tokens)

			formalParameters.AcceptDeclaration(v)

			v.tokens = NewTokens(v)
			v.tokens.Add(token.EndParametersBlock)
		}

		exceptions := declaration.ExceptionTypes()

		if exceptions != nil {
			v.tokens.Add(token.Space)
			v.tokens.Add(token.Throws)
			v.tokens.Add(token.Space)
			exceptions.AcceptTypeVisitor(v)
		}

		statements := declaration.Statements()

		if statements == nil {
			elementValue := declaration.DefaultAnnotationValue()

			if elementValue == nil {
				v.tokens.Add(token.Semicolon)
				v.fragments.AddTokensFragment(v.tokens)
			} else {
				v.tokens.Add(token.Space)
				v.tokens.Add(token.Default2)
				v.tokens.Add(token.Space)
				v.fragments.AddTokensFragment(v.tokens)

				elementValue.Accept(v)

				v.tokens = NewTokens(v)
				v.tokens.Add(token.Semicolon)
				v.fragments.AddTokensFragment(v.tokens)
			}
		} else {
			v.fragments.AddTokensFragment(v.tokens)
			v.singleLineStatementVisitor.Init()
			statements.AcceptStatement(v.singleLineStatementVisitor)

			singleLineStatement := v.singleLineStatementVisitor.IsSingleLineStatement()
			fragmentCount1 := v.fragments.Size()
			var start intmod.IStartBodyFragment

			if singleLineStatement {
				start = fragutil.AddStartSingleStatementMethodBody(v.fragments)
			} else {
				start = fragutil.AddStartMethodBody(v.fragments)
			}

			fragmentCount2 := v.fragments.Size()

			statements.AcceptStatement(v)

			if fragmentCount2 == v.fragments.Size() {
				v.fragments.SubList(fragmentCount1, fragmentCount2).Clear()
				v.tokens.Add(token.Space)
				v.tokens.Add(token.LeftRightCurlyBrackets)
			} else if singleLineStatement {
				fragutil.AddEndSingleStatementMethodBody(v.fragments, start)
			} else {
				fragutil.AddEndMethodBody(v.fragments, start)
			}
		}

		v.restoreContext()

		v.fragments.Add(javafragment.EndMovableBlock)
	}
}

func (v *CompilationUnitVisitor) VisitObjectReference(reference intmod.IObjectReference) {
	v.VisitObjectType(reference.(intmod.IObjectType))
}

func (v *CompilationUnitVisitor) VisitInnerObjectReference(reference intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(reference.(intmod.IInnerObjectType))
}

func (v *CompilationUnitVisitor) VisitStaticInitializerDeclaration(declaration intmod.IStaticInitializerDeclaration) {
	statements := declaration.Statements()

	if statements != nil {
		v.fragments.Add(javafragment.StartMovableMethodBlock)

		v.storeContext()
		v.currentMethodParamNames.Clear()

		v.tokens = NewTokens(v)
		v.tokens.Add(token.Static)
		v.fragments.AddTokensFragment(v.tokens)

		start := fragutil.AddStartMethodBody(v.fragments)
		statements.AcceptStatement(v)
		fragutil.AddEndMethodBody(v.fragments, start)

		v.fragments.Add(javafragment.EndMovableBlock)

		v.restoreContext()
	}
}

func (v *CompilationUnitVisitor) VisitTypeDeclarations(declaration intmod.ITypeDeclarations) {
	if declaration.Size() > 0 {
		iterator := declaration.Iterator()

		iterator.Next().AcceptDeclaration(v)

		for iterator.HasNext() {
			fragutil.AddSpacerBetweenMembers(v.fragments)
			iterator.Next().AcceptDeclaration(v)
		}
	}
}

func (v *CompilationUnitVisitor) buildFragmentsForTypeDeclaration(declaration intmod.ITypeDeclaration,
	flags int, keyword intmod.IKeywordToken) {
	v.tokens = NewTokens(v)

	// Build v.fragments for annotations
	annotationReferences := declaration.AnnotationReferences()
	if annotationReferences != nil {
		annotationReferences.Accept(v.annotationVisitor)
		v.fragments.AddTokensFragment(v.tokens)
		fragutil.AddSpacerAfterMemberAnnotations(v.fragments)
		v.tokens = NewTokens(v)
	}

	// Build v.tokens for access
	v.buildTokensForTypeAccessFlags(flags)
	v.tokens.Add(keyword)
	v.tokens.Add(token.Space)

	// Build token for type declaration
	v.tokens.Add(token.NewDeclarationToken(intmod.TypeToken, declaration.InternalTypeName(), declaration.Name(), ""))
}

func (v *CompilationUnitVisitor) buildFragmentsForClassOrInterfaceDeclaration(declaration intmod.IInterfaceDeclaration, flags int, keyword intmod.IKeywordToken) {
	v.buildFragmentsForTypeDeclaration(declaration, flags, keyword)

	// Build v.tokens for type parameterTypes
	typeParameters := declaration.TypeParameters()

	if typeParameters != nil {
		v.tokens.Add(token.LeftAngleBracket)
		typeParameters.AcceptTypeParameterVisitor(v)
		v.tokens.Add(token.RightAngleBracket)
	}
}

func (v *CompilationUnitVisitor) buildTokensForTypeAccessFlags(flags int) {
	if (flags & intmod.FlagPublic) != 0 {
		v.tokens.Add(token.Public)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagProtected) != 0 {
		v.tokens.Add(token.Protected)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagPrivate) != 0 {
		v.tokens.Add(token.Private)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagStatic) != 0 {
		v.tokens.Add(token.Static)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagFinal) != 0 {
		v.tokens.Add(token.Final)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagAbstract) != 0 {
		v.tokens.Add(token.Abstract)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagSynthetic) != 0 {
		v.tokens.Add(token.StartComment)
		v.tokens.Add(token.CommentSynthetic)
		v.tokens.Add(token.EndComment)
		v.tokens.Add(token.Space)
	}
}

func (v *CompilationUnitVisitor) buildTokensForFieldAccessFlags(flags int) {
	if (flags & intmod.FlagPublic) != 0 {
		v.tokens.Add(token.Public)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagProtected) != 0 {
		v.tokens.Add(token.Protected)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagPrivate) != 0 {
		v.tokens.Add(token.Private)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagStatic) != 0 {
		v.tokens.Add(token.Static)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagFinal) != 0 {
		v.tokens.Add(token.Final)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagVolatile) != 0 {
		v.tokens.Add(token.Volatile)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagTransient) != 0 {
		v.tokens.Add(token.Transient)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagSynthetic) != 0 {
		v.tokens.Add(token.StartComment)
		v.tokens.Add(token.CommentSynthetic)
		v.tokens.Add(token.EndComment)
		v.tokens.Add(token.Space)
	}
}

func (v *CompilationUnitVisitor) buildTokensForMethodAccessFlags(flags int) {
	if (flags & intmod.FlagPublic) != 0 {
		v.tokens.Add(token.Public)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagProtected) != 0 {
		v.tokens.Add(token.Protected)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagPrivate) != 0 {
		v.tokens.Add(token.Private)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagStatic) != 0 {
		v.tokens.Add(token.Static)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagFinal) != 0 {
		v.tokens.Add(token.Final)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagSynchronized) != 0 {
		v.tokens.Add(token.Synchronized)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagBridge) != 0 {
		v.tokens.Add(token.StartComment)
		v.tokens.Add(token.CommentBridge)
		v.tokens.Add(token.EndComment)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagNative) != 0 {
		v.tokens.Add(token.Native)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagAbstract) != 0 {
		v.tokens.Add(token.Abstract)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagStrict) != 0 {
		v.tokens.Add(token.Strict)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagSynthetic) != 0 {
		v.tokens.Add(token.StartComment)
		v.tokens.Add(token.CommentSynthetic)
		v.tokens.Add(token.EndComment)
		v.tokens.Add(token.Space)
	}
	if (flags & intmod.FlagDefault) != 0 {
		v.tokens.Add(token.Default2)
		v.tokens.Add(token.Space)
	}
}

func NewAnnotationVisitor(parent *CompilationUnitVisitor) *AnnotationVisitor {
	return &AnnotationVisitor{
		parent: parent,
	}
}

type AnnotationVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	parent *CompilationUnitVisitor
}

func (v *AnnotationVisitor) VisitAnnotationReferences(list intmod.IAnnotationReferences) {
	if list.Size() > 0 {
		iterator := list.Iterator()
		iterator.Next().Accept(v)

		for iterator.HasNext() {
			v.parent.fragments.AddTokensFragment(v.parent.tokens)

			fragutil.AddSpacerBetweenMemberAnnotations(v.parent.fragments)

			v.parent.tokens = NewTokens(v.parent)

			iterator.Next().Accept(v)
		}
	}
}

func (v *AnnotationVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.parent.visitAnnotationReferenceAnnotationReference(ref)
}

func (v *AnnotationVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.parent.visitAnnotationReferenceAnnotationReference(ref)
}
