package visitor

import (
	"bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	_ "bitbucket.org/coontec/go-jd-core/class/model/classfile"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	_ "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"strings"
	"unicode"
)

func NewInitInnerClassVisitor() *InitInnerClassVisitor {
	v := &InitInnerClassVisitor{}
	v.updateFieldDeclarationsAndReferencesVisitor.parent = v
	return v
}

type InitInnerClassVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	updateFieldDeclarationsAndReferencesVisitor UpdateFieldDeclarationsAndReferencesVisitor
	syntheticInnerFieldNames                    util.DefaultList[string]
	outerTypeFieldName                          string
}

func (v *InitInnerClassVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitInnerClassVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitInnerClassVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitInnerClassVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitInnerClassVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)

	// Init attributes
	v.outerTypeFieldName = ""
	v.syntheticInnerFieldNames.Clear()
	// Visit methods
	v.SafeAcceptListDeclaration(ConvertMethodDeclarations(bodyDeclaration.MethodDeclarations()))
	// Init values
	bodyDeclaration.SetOuterTypeFieldName(v.outerTypeFieldName)

	if !v.syntheticInnerFieldNames.IsEmpty() {
		bodyDeclaration.SetSyntheticInnerFieldNames(v.syntheticInnerFieldNames.ToSlice())
	}

	if v.outerTypeFieldName != "" || !v.syntheticInnerFieldNames.IsEmpty() {
		v.updateFieldDeclarationsAndReferencesVisitor.VisitBodyDeclaration(bodyDeclaration)
	}
}

func (v *InitInnerClassVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	cfcd := decl.(intsrv.IClassFileConstructorDeclaration)
	classFile := cfcd.ClassFile()
	outerClassFile := classFile.OuterClassFile()
	removeFirstParameter := false

	v.syntheticInnerFieldNames.Clear()

	// Search synthetic field initialization
	if cfcd.Statements().IsList() {
		iterator := cfcd.Statements().Iterator()

		for iterator.HasNext() {
			state := iterator.Next()

			if state.IsExpressionStatement() {
				expr := state.Expression()

				if expr.IsSuperConstructorInvocationExpression() {
					// 'super(...)'
					break
				}

				if expr.IsConstructorInvocationExpression() {
					// 'this(...)'
					if (outerClassFile != nil) && !classFile.IsStatic() {
						// Inner non-static class --> First parameter is the synthetic outer reference
						removeFirstParameter = true
					}
					break
				}

				if expr.IsBinaryOperatorExpression() {
					e := expr.LeftExpression()

					if e.IsFieldReferenceExpression() {
						name := e.Name()

						if strings.HasPrefix(name, "this$") {
							v.outerTypeFieldName = name
							removeFirstParameter = true
						} else if strings.HasPrefix(name, "val$") {
							v.syntheticInnerFieldNames.Add(name)
						}
					}
				}
			}

			_ = iterator.Remove()
		}
	}

	// Remove synthetic parameters
	parameters := cfcd.FormalParameters()

	if parameters != nil {
		if parameters.IsList() {
			list := util.NewDefaultList[intmod.IDeclaration]()

			if removeFirstParameter {
				// Remove outer this
				list.RemoveAt(0)
			}

			count := v.syntheticInnerFieldNames.Size()

			if count > 0 {
				// Remove outer local variable reference
				size := list.Size()
				list = list.SubList(size-count, size)
				list.Clear()
			}
		} else if removeFirstParameter || !v.syntheticInnerFieldNames.IsEmpty() {
			// Remove outer this and outer local variable reference
			cfcd.SetFormalParameters(nil)
		}
	}

	// Anonymous class constructor ?
	if outerClassFile != nil {
		outerTypeName := outerClassFile.InternalTypeName()
		internalTypeName := cfcd.ClassFile().InternalTypeName()
		var minmum int

		if strings.HasPrefix(internalTypeName, outerTypeName+"$") {
			minmum = len(outerTypeName) + 1
		} else {
			minmum = strings.LastIndex(internalTypeName, "$") + 1
		}

		if unicode.IsDigit(rune(internalTypeName[minmum])) {
			i := len(internalTypeName)
			anonymousFlag := true

			for i--; i > minmum; {
				if unicode.IsDigit(rune(internalTypeName[i])) {
					anonymousFlag = false
					break
				}
			}

			if anonymousFlag {
				// Mark anonymous class constructor
				cfcd.SetFlags(cfcd.Flags() | intmod.FlagAnonymous)
			}
		}
	}
}

func (v *InitInnerClassVisitor) VisitMethodDeclaration(_ intmod.IMethodDeclaration) {
}

func (v *InitInnerClassVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
}

func NewUpdateFieldDeclarationsAndReferencesVisitor(parent *InitInnerClassVisitor) *UpdateFieldDeclarationsAndReferencesVisitor {
	return &UpdateFieldDeclarationsAndReferencesVisitor{
		parent: parent,
	}
}

type UpdateFieldDeclarationsAndReferencesVisitor struct {
	AbstractUpdateExpressionVisitor

	parent          *InitInnerClassVisitor
	bodyDeclaration intsrv.IClassFileBodyDeclaration
	syntheticField  bool
}

func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)
	v.SafeAcceptListDeclaration(ConvertFieldDeclarations(bodyDeclaration.FieldDeclarations()))
	v.SafeAcceptListDeclaration(ConvertMethodDeclarations(bodyDeclaration.MethodDeclarations()))
}

func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	v.syntheticField = false
	decl.FieldDeclarators().AcceptDeclaration(v)

	if v.syntheticField {
		decl.SetFlags(decl.Flags() | intmod.FlagSynthetic)
	}
}

func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	name := decl.Name()

	if name == v.parent.outerTypeFieldName || v.parent.syntheticInnerFieldNames.Contains(name) {
		v.syntheticField = true
	}
}

func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitStaticInitializerDeclaration(decl intmod.IStaticInitializerDeclaration) {
}

func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitNewExpression(expr intmod.INewExpression) {
	if expr.Parameters() != nil {
		expr.SetParameters(v.UpdateBaseExpression(expr.Parameters()))
		expr.Parameters().Accept(v)
	}
	v.SafeAcceptDeclaration(expr.BodyDeclaration())
}

func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	if strings.HasPrefix(expr.Name(), "this$") {
		if expr.InternalTypeName() == v.bodyDeclaration.InternalTypeName() {
			if expr.Name() == v.parent.outerTypeFieldName {
				objectType := expr.Type().(intmod.IObjectType)
				var exp intmod.IExpression
				if expr.Expression() == nil {
					exp = expr
				} else {
					exp = expr.Expression()
				}

				expr.SetExpression(expression.NewObjectTypeReferenceExpressionWithLineNumber(
					exp.LineNumber(), objectType.CreateTypeWithArgs(nil)))
				expr.SetName("this")
			}
		} else {
			typeDeclaration := v.bodyDeclaration.InnerTypeDeclaration(expr.InternalTypeName())

			if typeDeclaration != nil && typeDeclaration.IsClassDeclaration() {
				cfbd := typeDeclaration.BodyDeclaration().(intsrv.IClassFileBodyDeclaration)
				outerInternalTypeName := cfbd.OuterBodyDeclaration().InternalTypeName()
				objectType := expr.Type().(intmod.IObjectType)

				if outerInternalTypeName == objectType.InternalName() {
					var exp intmod.IExpression

					if expr.Expression() == nil {
						exp = expr
					} else {
						exp = expr.Expression()
					}

					expr.SetExpression(expression.NewObjectTypeReferenceExpressionWithLineNumber(
						exp.LineNumber(), objectType.CreateTypeWithArgs(nil)))
					expr.SetName("this")
				}
			}
		}
	} else if strings.HasPrefix(expr.Name(), "val$") {
		expr.SetName(expr.Name()[4:])
		expr.SetExpression(nil)
	} else {
		v.AbstractUpdateExpressionVisitor.VisitFieldReferenceExpression(expr)
	}
}

func (v *UpdateFieldDeclarationsAndReferencesVisitor) updateExpression(expr intmod.IExpression) intmod.IExpression {
	if expr.IsLocalVariableReferenceExpression() {
		if expr.Name() != "" && expr.Name() == v.parent.outerTypeFieldName && expr.Type().IsObjectType() {
			objectType := expr.Type().(intmod.IObjectType)
			if v.bodyDeclaration.OuterBodyDeclaration().InternalTypeName() == objectType.InternalName() {
				return expression.NewFieldReferenceExpression(objectType.(intmod.IType),
					expression.NewObjectTypeReferenceExpressionWithLineNumber(
						expr.LineNumber(), objectType.CreateTypeWithArgs(nil)),
					objectType.InternalName(), "this", objectType.Descriptor())
			}
		}
	}

	return expr
}

func NewUpdateNewExpressionVisitor(typeMaker intsrv.ITypeMaker) *UpdateNewExpressionVisitor {
	return &UpdateNewExpressionVisitor{
		typeMaker:                 typeMaker,
		finalLocalVariableNameMap: make(map[string]string),
		localClassDeclarations:    util.NewDefaultList[intsrv.IClassFileClassDeclaration](),
		newExpressions:            util.NewSet[intmod.INewExpression](),
	}

}

type UpdateNewExpressionVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	typeMaker                 intsrv.ITypeMaker
	bodyDeclaration           intsrv.IClassFileBodyDeclaration
	classFile                 classpath.IClassFile
	finalLocalVariableNameMap map[string]string
	localClassDeclarations    util.IList[intsrv.IClassFileClassDeclaration]
	newExpressions            util.ISet[intmod.INewExpression]
	lineNumber                int
}

func (v *UpdateNewExpressionVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)
	v.SafeAcceptListDeclaration(ConvertMethodDeclarations(bodyDeclaration.MethodDeclarations()))
}

func (v *UpdateNewExpressionVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.classFile = decl.(intsrv.IClassFileConstructorDeclaration).ClassFile()
	v.finalLocalVariableNameMap = make(map[string]string)
	v.localClassDeclarations.Clear()

	v.SafeAcceptStatement(decl.Statements())

	if len(v.finalLocalVariableNameMap) != 0 {
		visitor := NewUpdateParametersAndLocalVariablesVisitor(v)

		decl.Statements().AcceptStatement(visitor)

		if decl.FormalParameters() != nil {
			decl.FormalParameters().Accept(visitor)
		}
	}

	if !v.localClassDeclarations.IsEmpty() {
		v.localClassDeclarations.Sort(func(i, j int) bool {
			return v.localClassDeclarations.Get(i).FirstLineNumber() <
				v.localClassDeclarations.Get(j).FirstLineNumber()
		})
		decl.AcceptDeclaration(NewAddLocalClassDeclarationVisitor(v))
	}
}

func (v *UpdateNewExpressionVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	v.finalLocalVariableNameMap = make(map[string]string)
	v.localClassDeclarations.Clear()
	v.SafeAcceptStatement(decl.Statements())

	if len(v.finalLocalVariableNameMap) != 0 {
		visitor := NewUpdateParametersAndLocalVariablesVisitor(v)
		decl.Statements().AcceptStatement(visitor)

		if decl.FormalParameters() != nil {
			decl.FormalParameters().Accept(visitor)
		}
	}

	if !v.localClassDeclarations.IsEmpty() {
		v.localClassDeclarations.Sort(func(i, j int) bool {
			return v.localClassDeclarations.Get(i).FirstLineNumber() <
				v.localClassDeclarations.Get(j).FirstLineNumber()
		})
		decl.AcceptDeclaration(NewAddLocalClassDeclarationVisitor(v))
	}
}

func (v *UpdateNewExpressionVisitor) VisitStaticInitializerDeclaration(decl intmod.IStaticInitializerDeclaration) {
	v.finalLocalVariableNameMap = make(map[string]string)
	v.localClassDeclarations.Clear()
	v.SafeAcceptStatement(decl.Statements())

	if len(v.finalLocalVariableNameMap) != 0 {
		decl.Statements().AcceptStatement(NewUpdateParametersAndLocalVariablesVisitor(v))
	}

	if !v.localClassDeclarations.IsEmpty() {
		v.localClassDeclarations.Sort(func(i, j int) bool {
			return v.localClassDeclarations.Get(i).FirstLineNumber() <
				v.localClassDeclarations.Get(j).FirstLineNumber()
		})
		decl.Accept(NewAddLocalClassDeclarationVisitor(v))
	}
}

func (v *UpdateNewExpressionVisitor) VisitStatements(list intmod.IStatements) {
	if !list.IsEmpty() {
		iterator := list.ListIterator()

		for iterator.HasNext() {
			//iterator.next().accept(v);
			s := iterator.Next()
			s.AcceptStatement(v)

			if v.lineNumber == intmod.UnknownLineNumber && !v.localClassDeclarations.IsEmpty() {
				iterator.Previous()

				for _, typeDeclaration := range v.localClassDeclarations.ToSlice() {
					_ = iterator.Add(statement.NewTypeDeclarationStatement(typeDeclaration))
				}

				v.localClassDeclarations.Clear()
				iterator.Next()
			}
		}
	}
}

func (v *UpdateNewExpressionVisitor) VisitNewExpression(expr intmod.INewExpression) {
	if !v.newExpressions.Contains(expr) {
		v.newExpressions.Add(expr)

		ne := expr.(intsrv.IClassFileNewExpression)
		var cfbd intsrv.IClassFileBodyDeclaration

		if ne.BodyDeclaration() == nil {
			typ := ne.ObjectType()
			internalName := typ.InternalName()
			typeDeclaration := v.bodyDeclaration.InnerTypeDeclaration(internalName).(intsrv.IClassFileTypeDeclaration)

			if typeDeclaration == nil {
				for bd := v.bodyDeclaration; bd != nil; bd = bd.OuterBodyDeclaration() {
					if bd.InternalTypeName() == internalName {
						cfbd = bd
						break
					}
				}
			} else if typeDeclaration.IsClassDeclaration() {
				cfcd := typeDeclaration.(intsrv.IClassFileClassDeclaration)
				cfbd = cfcd.BodyDeclaration().(intsrv.IClassFileBodyDeclaration)

				if typ.QualifiedName() == "" && typ.Name() != "" {
					// Local class
					cfcd.SetFlags(cfcd.Flags() & ^intmod.FlagSynthetic)
					v.localClassDeclarations.Add(cfcd)
					v.bodyDeclaration.RemoveInnerTypeDeclaration(internalName)
					v.lineNumber = ne.LineNumber()
				}
			}
		} else {
			// Anonymous class
			cfbd = ne.BodyDeclaration().(intsrv.IClassFileBodyDeclaration)
		}

		if cfbd != nil {
			parameters := ne.Parameters()
			parameterTypes := ne.ParameterTypes()

			if parameters != nil {
				// Remove synthetic parameters
				syntheticInnerFieldNames := util.NewDefaultListWithSlice[string](cfbd.SyntheticInnerFieldNames())

				if parameters.IsList() {
					list := util.NewDefaultListWithSlice[intmod.IExpression](parameters.ToSlice())
					types := util.NewDefaultListWithSlice[intmod.IType](parameterTypes.ToSlice())

					if cfbd.OuterTypeFieldName() != "" {
						// Remove outer this
						list.RemoveFirst()
						types.RemoveFirst()
					}

					if syntheticInnerFieldNames != nil {
						// Remove outer local variable reference
						size := list.Size()
						count := syntheticInnerFieldNames.Size()
						lastParameters := list.SubList(size-count, size)
						parameterIterator := lastParameters.Iterator()
						syntheticInnerFieldNameIterator := syntheticInnerFieldNames.Iterator()

						for parameterIterator.HasNext() {
							param := parameterIterator.Next()
							syntheticInnerFieldName := syntheticInnerFieldNameIterator.Next()

							if param.IsCastExpression() {
								param = param.Expression()
							}

							if param.IsLocalVariableReferenceExpression() {
								lv := param.(intsrv.IClassFileLocalVariableReferenceExpression).
									LocalVariable().(intsrv.ILocalVariable)
								localVariableName := syntheticInnerFieldName[4:]
								v.finalLocalVariableNameMap[lv.Name()] = localVariableName
							}
						}

						lastParameters.Clear()
						removal := types.SubList(size-count, size)
						types.RemoveAll(removal.ToSlice())
					}
				} else if cfbd.OuterTypeFieldName() != "" {
					// Remove outer this
					ne.SetParameters(nil)
					ne.SetParameterTypes(nil)
				} else if syntheticInnerFieldNames != nil {
					// Remove outer local variable reference
					param := parameters.First()

					if param.IsCastExpression() {
						param = param.Expression()
					}

					if param.IsLocalVariableReferenceExpression() {
						lv := param.(intsrv.IClassFileLocalVariableReferenceExpression).
							LocalVariable().(intsrv.ILocalVariable)
						localVariableName := syntheticInnerFieldNames.First()[4:]
						v.finalLocalVariableNameMap[lv.Name()] = localVariableName
						ne.SetParameters(nil)
						ne.SetParameterTypes(nil)
					}
				}

				// Is the last parameter synthetic ?
				parameters = ne.Parameters()

				if (parameters != nil) && (parameters.Size() > 0) && parameters.Last().IsNullExpression() {
					parameterTypes = ne.ParameterTypes()

					if parameterTypes.Last().Name() == "" {
						// Yes. Remove it.
						if parameters.IsList() {
							parameters.ToList().RemoveLast()
							parameterTypes.ToList().RemoveLast()
						} else {
							ne.SetParameters(nil)
							ne.SetParameterTypes(nil)
						}
					}
				}
			}
		}
	}

	v.SafeAcceptExpression(expr.Parameters())
}

func (v *UpdateNewExpressionVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	scie := expr.(intsrv.IClassFileSuperConstructorInvocationExpression)
	parameters := scie.Parameters()

	if (parameters != nil) && (parameters.Size() > 0) {
		// Remove outer 'this' reference parameter
		firstParameterType := parameters.First().Type()

		if firstParameterType.IsObjectType() && !v.classFile.IsStatic() && (v.bodyDeclaration.OuterTypeFieldName() != "") {
			superTypeTypes := v.typeMaker.MakeTypeTypes(v.classFile.SuperTypeName())

			if (superTypeTypes != nil) && superTypeTypes.ThisType().IsInnerObjectType() {
				if v.typeMaker.IsRawTypeAssignable(superTypeTypes.ThisType().OuterType(), firstParameterType.(intmod.IObjectType)) {
					scie.SetParameters(v.removeFirstItemExpression(parameters))
					scie.SetParameterTypes(v.removeFirstItemType(scie.ParameterTypes()))
				}
			}
		}

		// Remove last synthetic parameter
		expr.SetParameters(v.removeLastSyntheticParameter(scie.Parameters(), scie.ParameterTypes()))
	}
}

func (v *UpdateNewExpressionVisitor) VisitConstructorInvocationExpression(expr intmod.IConstructorInvocationExpression) {
	cie := expr.(intsrv.IClassFileConstructorInvocationExpression)
	parameters := cie.Parameters()

	if (parameters != nil) && (parameters.Size() > 0) {
		// Remove outer this reference parameter
		if v.bodyDeclaration.OuterTypeFieldName() != "" {
			cie.SetParameters(v.removeFirstItemExpression(parameters))
			cie.SetParameterTypes(v.removeFirstItemType(cie.ParameterTypes()))
		}

		// Remove last synthetic parameter
		cie.SetParameters(v.removeLastSyntheticParameter(cie.Parameters(), cie.ParameterTypes()))
	}
}

func (v *UpdateNewExpressionVisitor) removeFirstItemExpression(parameters intmod.IExpression) intmod.IExpression {
	if parameters.IsList() {
		parameters.ToList().RemoveFirst()
	} else {
		parameters = nil
	}

	return parameters
}

func (v *UpdateNewExpressionVisitor) removeFirstItemType(types intmod.IType) intmod.IType {
	if types.IsList() {
		types.ToList().RemoveFirst()
	} else {
		types = nil
	}

	return types
}

func (v *UpdateNewExpressionVisitor) removeLastSyntheticParameter(parameters intmod.IExpression, parameterTypes intmod.IType) intmod.IExpression {
	// Is the last parameter synthetic ?
	if (parameters != nil) && (parameters.Size() > 0) && parameters.Last().IsNullExpression() {
		if parameterTypes.Last().Name() == "" {
			// Yes. Remove it.
			if parameters.IsList() {
				parameters.ToList().RemoveFirst()
			} else {
				parameters = nil
			}
		}
	}

	return parameters
}

func NewUpdateParametersAndLocalVariablesVisitor(parent *UpdateNewExpressionVisitor) *UpdateParametersAndLocalVariablesVisitor {
	return &UpdateParametersAndLocalVariablesVisitor{
		parent: parent,
	}
}

type UpdateParametersAndLocalVariablesVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	parent *UpdateNewExpressionVisitor
	final  bool
}

func (v *UpdateParametersAndLocalVariablesVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	if value, ok := v.parent.finalLocalVariableNameMap[decl.Name()]; ok {
		decl.SetFinal(true)
		decl.SetName(value)
	}
}

func (v *UpdateParametersAndLocalVariablesVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.final = false
	stat.LocalVariableDeclarators().Accept(v)
	stat.SetFinal(v.final)
}

func (v *UpdateParametersAndLocalVariablesVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.final = false
	decl.LocalVariableDeclarators().Accept(v)
	decl.SetFinal(v.final)
}

func (v *UpdateParametersAndLocalVariablesVisitor) VisitLocalVariableDeclarator(declarator intmod.ILocalVariableDeclarator) {
	if value, ok := v.parent.finalLocalVariableNameMap[declarator.Name()]; ok {
		v.final = true
		declarator.SetName(value)
	}
}

func NewAddLocalClassDeclarationVisitor(parent *UpdateNewExpressionVisitor) *AddLocalClassDeclarationVisitor {
	return &AddLocalClassDeclarationVisitor{
		parent:                       parent,
		searchFirstLineNumberVisitor: NewSearchFirstLineNumberVisitor(),
		lineNumber:                   intmod.UnknownLineNumber,
	}
}

type AddLocalClassDeclarationVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	parent                       *UpdateNewExpressionVisitor
	searchFirstLineNumberVisitor *SearchFirstLineNumberVisitor
	lineNumber                   int
}

func (v *AddLocalClassDeclarationVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	cfcd := decl.(intsrv.IClassFileConstructorDeclaration)
	cfcd.SetStatements(v.addLocalClassDeclarations(cfcd.Statements()))
}

func (v *AddLocalClassDeclarationVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	cfmd := decl.(intsrv.IClassFileMethodDeclaration)
	cfmd.SetStatements(v.addLocalClassDeclarations(cfmd.Statements()))
}

func (v *AddLocalClassDeclarationVisitor) VisitStaticInitializerDeclaration(decl intmod.IStaticInitializerDeclaration) {
	cfsid := decl.(intsrv.IClassFileStaticInitializerDeclaration)
	cfsid.SetStatements(v.addLocalClassDeclarations(cfsid.Statements()))
}

func (v *AddLocalClassDeclarationVisitor) addLocalClassDeclarations(stat intmod.IStatement) intmod.IStatement {
	if !v.parent.localClassDeclarations.IsEmpty() {
		if stat.IsStatements() {
			stat.AcceptStatement(v)
		} else {
			decl := v.parent.localClassDeclarations.Get(0)

			v.searchFirstLineNumberVisitor.Init()
			stat.AcceptStatement(v.searchFirstLineNumberVisitor)

			if v.searchFirstLineNumberVisitor.LineNumber() != -1 {
				v.lineNumber = v.searchFirstLineNumberVisitor.LineNumber()
			}

			if decl.FirstLineNumber() <= v.lineNumber {
				list := statement.NewStatements()
				declarationIterator := v.parent.localClassDeclarations.Iterator()

				list.Add(statement.NewTypeDeclarationStatement(decl))
				declarationIterator.Next()
				_ = declarationIterator.Remove()

				for declarationIterator.HasNext() {
					decl = declarationIterator.Next()
					if decl.FirstLineNumber() <= v.lineNumber {
						list.Add(statement.NewTypeDeclarationStatement(decl))
						_ = declarationIterator.Remove()
					}
				}

				if stat.IsList() {
					list.AddAll(stat.ToSlice())
				} else {
					list.Add(stat.First())
				}
				stat = list
			} else {
				stat.AcceptStatement(v)
			}
		}
	}

	return stat
}

func (v *AddLocalClassDeclarationVisitor) VisitStatements(list intmod.IStatements) {
	if !v.parent.localClassDeclarations.IsEmpty() && !list.IsEmpty() {
		statementIterator := list.ListIterator()
		declarationIterator := v.parent.localClassDeclarations.Iterator()
		decl := declarationIterator.Next()

		for statementIterator.HasNext() {
			state := statementIterator.Next()

			v.searchFirstLineNumberVisitor.Init()
			state.AcceptStatement(v.searchFirstLineNumberVisitor)

			if v.searchFirstLineNumberVisitor.LineNumber() != -1 {
				v.lineNumber = v.searchFirstLineNumberVisitor.LineNumber()
			}

			for decl.FirstLineNumber() <= v.lineNumber {
				statementIterator.Previous()
				_ = statementIterator.Add(statement.NewTypeDeclarationStatement(decl))
				statementIterator.Next()
				_ = declarationIterator.Remove()

				if !declarationIterator.HasNext() {
					return
				}

				decl = declarationIterator.Next()
			}
		}
	}
}

type MemberDeclarationComparator struct {
}

func (c *MemberDeclarationComparator) Compare(md1, md2 intsrv.IClassFileMemberDeclaration) int {
	return md1.FirstLineNumber() - md2.FirstLineNumber()
}
