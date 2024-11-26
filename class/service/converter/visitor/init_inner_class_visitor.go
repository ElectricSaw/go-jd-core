package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"strings"
)

func NewInitInnerClassVisitor() *InitInnerClassVisitor {
	return &InitInnerClassVisitor{}
}

type InitInnerClassVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	updateFieldDeclarationsAndReferencesVisitor UpdateFieldDeclarationsAndReferencesVisitor
	syntheticInnerFieldNames                    util.DefaultList[string]
	outerTypeFieldName                          string
}


func (v *InitInnerClassVisitor) VisitAnnotationDeclaration( decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration());
}


func (v *InitInnerClassVisitor) VisitClassDeclaration( decl intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration());
}


func (v *InitInnerClassVisitor) VisitEnumDeclaration( decl intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration());
}


func (v *InitInnerClassVisitor) VisitInterfaceDeclaration( decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration());
}


func (v *InitInnerClassVisitor) VisitBodyDeclaration( decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration);

	// Init attributes
	v.outerTypeFieldName = "";
	v.syntheticInnerFieldNames.Clear();
	// Visit methods
	v.SafeAcceptListDeclaration(ConvertMethodDeclarations(bodyDeclaration.MethodDeclarations()));
	// Init values
	bodyDeclaration.SetOuterTypeFieldName(v.outerTypeFieldName);

	if !v.syntheticInnerFieldNames.IsEmpty() {
		bodyDeclaration.SetSyntheticInnerFieldNames(v.syntheticInnerFieldNames.ToSlice());
	}

	if v.outerTypeFieldName != "" || !v.syntheticInnerFieldNames.IsEmpty() {
		v.updateFieldDeclarationsAndReferencesVisitor.VisitBodyDeclaration(bodyDeclaration);
	}
}



func (v *InitInnerClassVisitor) VisitConstructorDeclaration( decl intmod.IConstructorDeclaration) {
	cfcd := decl.(intsrv.IClassFileConstructorDeclaration);
	classFile := cfcd.ClassFile();
	outerClassFile := classFile.OuterClassFile();
	removeFirstParameter := false;

	v.syntheticInnerFieldNames.Clear();

	// Search synthetic field initialization
	if cfcd.Statements().IsList() {
		iterator := cfcd.Statements().Iterator();

		for iterator.HasNext() {
			statement := iterator.Next();

			if statement.IsExpressionStatement() {
				expression := statement.Expression();

				if expression.IsSuperConstructorInvocationExpression() {
					// 'super(...)'
					break;
				}

				if expression.IsConstructorInvocationExpression() {
					// 'this(...)'
					if (outerClassFile != nil) && !classFile.IsStatic() {
						// Inner non-static class --> First parameter is the synthetic outer reference
						removeFirstParameter = true;
					}
					break;
				}

				if expression.IsBinaryOperatorExpression() {
					e := expression.LeftExpression();

					if e.IsFieldReferenceExpression() {
						name := e.Name();

						if strings.HasPrefix(name, "this$") {
							v.outerTypeFieldName = name;
							removeFirstParameter = true;
						} else if strings.HasPrefix(name, "val$") {
							v.syntheticInnerFieldNames.Add(name);
						}
					}
				}
			}

			iterator.Remove();
		}
	}

	// Remove synthetic parameters
	parameters := cfcd.FormalParameters();

	if parameters != nil {
		if parameters.IsList() {
			list := util.NewDefaultList[intmod.IDeclaration]()

			if removeFirstParameter {
				// Remove outer this
				list.RemoveAt(0);
			}

			count := v.syntheticInnerFieldNames.Size();

			if count > 0 {
				// Remove outer local variable reference
				size := list.Size();
				list.SubList(size - count, size).clear();
			}
		} else if (removeFirstParameter || !syntheticInnerFieldNames.isEmpty()) {
			// Remove outer this and outer local variable reference
			cfcd.setFormalParameters(null);
		}
	}

	// Anonymous class constructor ?
	if (outerClassFile != null) {
		String outerTypeName = outerClassFile.getInternalTypeName();
		String internalTypeName = cfcd.getClassFile().getInternalTypeName();
		int min;

		if (internalTypeName.startsWith(outerTypeName + '$')) {
			min = outerTypeName.length() + 1;
		} else {
			min = internalTypeName.lastIndexOf('$') + 1;
		}

		if (Character.isDigit(internalTypeName.charAt(min))) {
			int i = internalTypeName.length();
			boolean anonymousFlag = true;

			while (--i > min) {
				if (!Character.isDigit(internalTypeName.charAt(i))) {
					anonymousFlag = false;
					break;
				}
			}

			if (anonymousFlag) {
				// Mark anonymous class constructor
				cfcd.setFlags(cfcd.getFlags() | Declaration.FLAG_ANONYMOUS);
			}
		}
	}
}

func (v *InitInnerClassVisitor) VisitMethodDeclaration( declaration intmod.IMethodDeclaration) {}
func (v *InitInnerClassVisitor) VisitStaticInitializerDeclaration( declaration intmod.IStaticInitializerDeclaration) {}

type UpdateFieldDeclarationsAndReferencesVisitor struct {
	AbstractUpdateExpressionVisitor

	bodyDeclaration intsrv.IClassFileBodyDeclaration
	syntheticField bool
}


func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitBodyDeclaration( decl intmod.IBodyDeclaration) {
	bodyDeclaration = (ClassFileBodyDeclaration)decl;
	safeAcceptListDeclaration(bodyDeclaration.getFieldDeclarations());
	safeAcceptListDeclaration(bodyDeclaration.getMethodDeclarations());
}


func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitFieldDeclaration( decl intmod.IFieldDeclaration) {
	syntheticField = false;
	decl.getFieldDeclarators().accept(this);

	if (syntheticField) {
		decl.setFlags(declaration.getFlags()|FLAG_SYNTHETIC);
	}
}


func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitFieldDeclarator( declarator intmod.IFieldDeclarator) {
	String name = declarator.getName();

	if (name.equals(outerTypeFieldName) || syntheticInnerFieldNames.contains(name)) {
		syntheticField = true;
	}
}

func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitStaticInitializerDeclaration( declaration intmod.IStaticInitializerDeclaration) {}


func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitMethodDeclaration( declaration intmod.IMethodDeclaration) {
	safeAccept(declaration.getStatements());
}


func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitNewExpression( expression intmod.INewExpression) {
	if (expression.getParameters() != null) {
		expression.setParameters(updateBaseExpression(expression.getParameters()));
		expression.getParameters().accept(this);
	}
	safeAccept(expression.getBodyDeclaration());
}


func (v *UpdateFieldDeclarationsAndReferencesVisitor) VisitFieldReferenceExpression( expression intmod.IFieldReferenceExpression) {
	if (expression.getName().startsWith("this$")) {
		if (expression.getInternalTypeName().equals(bodyDeclaration.getInternalTypeName())) {
			if (expression.getName().equals(outerTypeFieldName)) {
				ObjectType objectType = (ObjectType)expression.getType();
				Expression exp = (expression.getExpression() == null) ? expression : expression.getExpression();
				expression.setExpression(new ObjectTypeReferenceExpression(exp.getLineNumber(), objectType.createType(null)));
				expression.setName("this");
			}
		} else {
			ClassFileTypeDeclaration typeDeclaration = bodyDeclaration.getInnerTypeDeclaration(expression.getInternalTypeName());

			if ((typeDeclaration != null) && typeDeclaration.isClassDeclaration()) {
				ClassFileBodyDeclaration cfbd = (ClassFileBodyDeclaration) typeDeclaration.getBodyDeclaration();
				String outerInternalTypeName = cfbd.getOuterBodyDeclaration().getInternalTypeName();
				ObjectType objectType = (ObjectType)expression.getType();

				if (outerInternalTypeName.equals(objectType.getInternalName())) {
					Expression exp = (expression.getExpression() == null) ? expression : expression.getExpression();
					expression.setExpression(new ObjectTypeReferenceExpression(exp.getLineNumber(), objectType.createType(null)));
					expression.setName("this");
				}
			}
		}
	} else if (expression.getName().startsWith("val$")) {
		expression.setName(expression.getName().substring(4));
		expression.setExpression(null);
	} else {
		super.Visit(expression);
	}
}


func (v *UpdateFieldDeclarationsAndReferencesVisitor) updateExpression(expression intmod.IExpression) intmod.IExpression {
	if (expression.isLocalVariableReferenceExpression()) {
		if ((expression.getName() != null) && expression.getName().equals(outerTypeFieldName) && expression.getType().isObjectType()) {
			ObjectType objectType = (ObjectType)expression.getType();
			if (bodyDeclaration.getOuterBodyDeclaration().getInternalTypeName().equals(objectType.getInternalName())) {
				return new FieldReferenceExpression(objectType, new ObjectTypeReferenceExpression(expression.getLineNumber(), objectType.createType(null)), objectType.getInternalName(), "this", objectType.getDescriptor());
			}
		}
	}

	return expression;
}
}

func NewUpdateNewExpressionVisitor(typeMaker *utils.TypeMaker) *UpdateNewExpressionVisitor {
	return &UpdateNewExpressionVisitor {
		typeMaker: typeMaker,
		finalLocalVariableNameMap: make(map[string]string),
		localClassDeclarations: util.NewDefaultList[intsrv.IClassFileClassDeclaration](),
		newExpressions: util.NewSet[intmod.INewExpression],
	}

}

type UpdateNewExpressionVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	typeMaker *utils.TypeMaker
	bodyDeclaration intsrv.IClassFileBodyDeclaration
	classFile *classfile.ClassFile
	finalLocalVariableNameMap map[string]string
	localClassDeclarations *util.DefaultList[intsrv.IClassFileClassDeclaration]
	newExpressions util.ISet[intmod.INewExpression]
	lineNumber int
}

func (v *UpdateNewExpressionVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration = (ClassFileBodyDeclaration)declaration;
	safeAcceptListDeclaration(bodyDeclaration.getMethodDeclarations());
}


func (v *UpdateNewExpressionVisitor) VisitConstructorDeclaration( declaration intmod.IConstructorDeclaration) {
	classFile = ((ClassFileConstructorDeclaration)declaration).getClassFile();
	finalLocalVariableNameMap.clear();
	localClassDeclarations.clear();

	safeAccept(declaration.getStatements());

	if (! finalLocalVariableNameMap.isEmpty()) {
		UpdateParametersAndLocalVariablesVisitor Visitor = new UpdateParametersAndLocalVariablesVisitor();

		declaration.getStatements().accept(Visitor);

		if (declaration.getFormalParameters() != null) {
			declaration.getFormalParameters().accept(Visitor);
		}
	}

	if (! localClassDeclarations.isEmpty()) {
		localClassDeclarations.sort(new MemberDeclarationComparator());
		declaration.accept(new AddLocalClassDeclarationVisitor());
	}
}


func (v *UpdateNewExpressionVisitor) Visit(MethodDeclaration declaration) {
	finalLocalVariableNameMap.clear();
	localClassDeclarations.clear();
	safeAccept(declaration.getStatements());

	if (! finalLocalVariableNameMap.isEmpty()) {
		UpdateParametersAndLocalVariablesVisitor Visitor = new UpdateParametersAndLocalVariablesVisitor();

		declaration.getStatements().accept(Visitor);

		if (declaration.getFormalParameters() != null) {
			declaration.getFormalParameters().accept(Visitor);
		}
	}

	if (! localClassDeclarations.isEmpty()) {
		localClassDeclarations.sort(new MemberDeclarationComparator());
		declaration.accept(new AddLocalClassDeclarationVisitor());
	}
}


func (v *UpdateNewExpressionVisitor) Visit(StaticInitializerDeclaration declaration) {
	finalLocalVariableNameMap.clear();
	localClassDeclarations.clear();
	safeAccept(declaration.getStatements());

	if (! finalLocalVariableNameMap.isEmpty()) {
		declaration.getStatements().accept(new UpdateParametersAndLocalVariablesVisitor());
	}

	if (! localClassDeclarations.isEmpty()) {
		localClassDeclarations.sort(new MemberDeclarationComparator());
		declaration.accept(new AddLocalClassDeclarationVisitor());
	}
}



func (v *UpdateNewExpressionVisitor) Visit(Statements list) {
	if (!list.isEmpty()) {
		ListIterator<Statement> iterator = list.listIterator();

		while (iterator.hasNext()) {
			//iterator.next().accept(this);
			Statement s = iterator.next();
			s.accept(this);

			if ((lineNumber == Expression.UNKNOWN_LINE_NUMBER) && !localClassDeclarations.isEmpty()) {
				iterator.previous();

				for (TypeDeclaration typeDeclaration : localClassDeclarations) {
					iterator.add(new TypeDeclarationStatement(typeDeclaration));
				}

				localClassDeclarations.clear();
				iterator.next();
			}
		}
	}
}



func (v *UpdateNewExpressionVisitor) Visit(NewExpression expression) {
	if (!newExpressions.contains(expression)) {
		newExpressions.add(expression);

		ClassFileNewExpression ne = (ClassFileNewExpression)expression;
		ClassFileBodyDeclaration cfbd = null;

		if (ne.getBodyDeclaration() == null) {
			ObjectType type = ne.getObjectType();
			String internalName = type.getInternalName();
			ClassFileTypeDeclaration typeDeclaration = bodyDeclaration.getInnerTypeDeclaration(internalName);

			if (typeDeclaration == null) {
				for (ClassFileBodyDeclaration bd = bodyDeclaration; bd != null; bd = bd.getOuterBodyDeclaration()) {
					if (bd.getInternalTypeName().equals(internalName)) {
						cfbd = bd;
						break;
					}
				}
			} else if (typeDeclaration.isClassDeclaration()) {
				ClassFileClassDeclaration cfcd = (ClassFileClassDeclaration) typeDeclaration;
				cfbd = (ClassFileBodyDeclaration) cfcd.getBodyDeclaration();

				if ((type.getQualifiedName() == null) && (type.getName() != null)) {
				// Local class
				cfcd.setFlags(cfcd.getFlags() & (~FLAG_SYNTHETIC));
				localClassDeclarations.add(cfcd);
				bodyDeclaration.removeInnerType(internalName);
				lineNumber = ne.getLineNumber();
				}
			}
		} else {
			// Anonymous class
			cfbd = (ClassFileBodyDeclaration) ne.getBodyDeclaration();
		}

		if (cfbd != null) {
			BaseExpression parameters = ne.getParameters();
			BaseType parameterTypes = ne.getParameterTypes();

			if (parameters != null) {
				// Remove synthetic parameters
				DefaultList<String> syntheticInnerFieldNames = cfbd.getSyntheticInnerFieldNames();

				if (parameters.isList()) {
					DefaultList<Expression> list = parameters.getList();
					DefaultList<Type> types = parameterTypes.getList();

					if (cfbd.getOuterTypeFieldName() != null) {
						// Remove outer this
						list.removeFirst();
						types.removeFirst();
					}

					if (syntheticInnerFieldNames != null) {
						// Remove outer local variable reference
						int size = list.size();
						int count = syntheticInnerFieldNames.size();
						List<Expression> lastParameters = list.subList(size - count, size);
						Iterator<Expression> parameterIterator = lastParameters.iterator();
						Iterator<String> syntheticInnerFieldNameIterator = syntheticInnerFieldNames.iterator();

						while (parameterIterator.hasNext()) {
							Expression param = parameterIterator.next();
							String syntheticInnerFieldName = syntheticInnerFieldNameIterator.next();

							if (param.isCastExpression()) {
								param = param.getExpression();
							}

							if (param.isLocalVariableReferenceExpression()) {
								AbstractLocalVariable lv = ((ClassFileLocalVariableReferenceExpression) param).getLocalVariable();
								String localVariableName = syntheticInnerFieldName.substring(4);
								finalLocalVariableNameMap.put(lv.getName(), localVariableName);
							}
						}

						lastParameters.clear();
						types.subList(size - count, size).clear();
					}
				} else if (cfbd.getOuterTypeFieldName() != null) {
					// Remove outer this
					ne.setParameters(null);
					ne.setParameterTypes(null);
				} else if (syntheticInnerFieldNames != null) {
					// Remove outer local variable reference
					Expression param = parameters.getFirst();

					if (param.isCastExpression()) {
						param = param.getExpression();
					}

					if (param.isLocalVariableReferenceExpression()) {
						AbstractLocalVariable lv = ((ClassFileLocalVariableReferenceExpression) param).getLocalVariable();
						String localVariableName = syntheticInnerFieldNames.getFirst().substring(4);
						finalLocalVariableNameMap.put(lv.getName(), localVariableName);
						ne.setParameters(null);
						ne.setParameterTypes(null);
					}
				}

				// Is the last parameter synthetic ?
				parameters = ne.getParameters();

				if ((parameters != null) && (parameters.size() > 0) && parameters.getLast().isNullExpression()) {
					parameterTypes = ne.getParameterTypes();

					if (parameterTypes.getLast().getName() == null) {
						// Yes. Remove it.
						if (parameters.isList()) {
							parameters.getList().removeLast();
							parameterTypes.getList().removeLast();
						} else {
							ne.setParameters(null);
							ne.setParameterTypes(null);
						}
					}
				}
			}
		}
	}

	safeAccept(expression.getParameters());
}


func (v *UpdateNewExpressionVisitor) Visit(SuperConstructorInvocationExpression expression) {
	ClassFileSuperConstructorInvocationExpression scie = (ClassFileSuperConstructorInvocationExpression)expression;
	BaseExpression parameters = scie.getParameters();

	if ((parameters != null) && (parameters.size() > 0)) {
		// Remove outer 'this' reference parameter
		Type firstParameterType = parameters.getFirst().getType();

		if (firstParameterType.isObjectType() && !classFile.isStatic() && (bodyDeclaration.getOuterTypeFieldName() != null)) {
			TypeMaker.TypeTypes superTypeTypes = typeMaker.makeTypeTypes(classFile.getSuperTypeName());

			if ((superTypeTypes != null) && superTypeTypes.thisType.isInnerObjectType()) {
				if (typeMaker.isRawTypeAssignable(superTypeTypes.thisType.getOuterType(), (ObjectType)firstParameterType)) {
				scie.setParameters(removeFirstItem(parameters));
				scie.setParameterTypes(removeFirstItem(scie.getParameterTypes()));
				}
			}
		}

		// Remove last synthetic parameter
		expression.setParameters(removeLastSyntheticParameter(scie.getParameters(), scie.getParameterTypes()));
	}
}


func (v *UpdateNewExpressionVisitor) Visit(ConstructorInvocationExpression expression) {
	ClassFileConstructorInvocationExpression cie = (ClassFileConstructorInvocationExpression)expression;
	BaseExpression parameters = cie.getParameters();

	if ((parameters != null) && (parameters.size() > 0)) {
		// Remove outer this reference parameter
		if (bodyDeclaration.getOuterTypeFieldName() != null) {
			cie.setParameters(removeFirstItem(parameters));
			cie.setParameterTypes(removeFirstItem(cie.getParameterTypes()));
		}

		// Remove last synthetic parameter
		cie.setParameters(removeLastSyntheticParameter(cie.getParameters(), cie.getParameterTypes()));
	}
}

func (v *UpdateNewExpressionVisitor) removeFirstItem(parameters intmod.IExpression) intmod.IExpression {
	if (parameters.isList()) {
		parameters.getList().removeFirst();
	} else {
		parameters = null;
	}

	return parameters;
}

func (v *UpdateNewExpressionVisitor) removeFirstItem(types intmod.IType)intmod.IType {
	if (types.isList()) {
		types.getList().removeFirst();
	} else {
		types = null;
	}

	return types;
}

func (v *UpdateNewExpressionVisitor)  removeLastSyntheticParameter(parameters intmod.IExpression, parameterTypes intmod.IType) intmod.IExpression {
	// Is the last parameter synthetic ?
	if ((parameters != null) && (parameters.size() > 0) && parameters.getLast().isNullExpression()) {
		if (parameterTypes.getLast().getName() == null) {
			// Yes. Remove it.
			if (parameters.isList()) {
				parameters.getList().removeLast();
			} else {
				parameters = null;
			}
		}
	}

	return parameters;
}

type UpdateParametersAndLocalVariablesVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	final bool
}


func (v *InitInnerClassVisitor) Visit(FormalParameter declaration) {
	if (finalLocalVariableNameMap.containsKey(declaration.getName())) {
		declaration.setFinal(true);
		declaration.setName(finalLocalVariableNameMap.get(declaration.getName()));
	}
}


func (v *InitInnerClassVisitor) Visit(LocalVariableDeclarationStatement statement) {
	fina1 = false;
	statement.getLocalVariableDeclarators().accept(this);
	statement.setFinal(fina1);
}


func (v *InitInnerClassVisitor) Visit(LocalVariableDeclaration declaration) {
	fina1 = false;
	declaration.getLocalVariableDeclarators().accept(this);
	declaration.setFinal(fina1);
}


func (v *InitInnerClassVisitor) Visit(LocalVariableDeclarator declarator) {
	if (finalLocalVariableNameMap.containsKey(declarator.getName())) {
		fina1 = true;
		declarator.setName(finalLocalVariableNameMap.get(declarator.getName()));
	}
}
}

func NewAddLocalClassDeclarationVisitor() *AddLocalClassDeclarationVisitor {
	return &AddLocalClassDeclarationVisitor{
		searchFirstLineNumberVisitor: *NewSearchFirstLineNumberVisitor(),
		lineNumber: intmod.UnknownLineNumber
	}
}

type AddLocalClassDeclarationVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	searchFirstLineNumberVisitor SearchFirstLineNumberVisitor
	lineNumber int
}


func (v *AddLocalClassDeclarationVisitor) VisitConstructorDeclaration( declaration intmod.IConstructorDeclaration) {
	ClassFileConstructorDeclaration cfcd = (ClassFileConstructorDeclaration)declaration;
	cfcd.setStatements(addLocalClassDeclarations(cfcd.getStatements()));
}


func (v *AddLocalClassDeclarationVisitor) VisitMethodDeclaration( declaration intmod.IMethodDeclaration) {
	ClassFileMethodDeclaration cfmd = (ClassFileMethodDeclaration)declaration;
	cfmd.setStatements(addLocalClassDeclarations(cfmd.getStatements()));
}

func (v *AddLocalClassDeclarationVisitor) VisitStaticInitializerDeclaration( declaration intmod.IStaticInitializerDeclaration) {
	ClassFileStaticInitializerDeclaration cfsid = (ClassFileStaticInitializerDeclaration)declaration;
	cfsid.setStatements(addLocalClassDeclarations(cfsid.getStatements()));
}

func (v *AddLocalClassDeclarationVisitor) addLocalClassDeclarations(statements intmod.IStatement) intmod.IStatement {
	if (!localClassDeclarations.isEmpty()) {
		if (statements.isStatements()) {
			statements.accept(this);
		} else {
			ClassFileClassDeclaration declaration = localClassDeclarations.get(0);

			searchFirstLineNumberVisitor.init();
			statements.accept(searchFirstLineNumberVisitor);

			if (searchFirstLineNumberVisitor.getLineNumber() != -1) {
				lineNumber = searchFirstLineNumberVisitor.getLineNumber();
			}

			if (declaration.getFirstLineNumber() <= lineNumber) {
				Statements list = new Statements();
				Iterator<ClassFileClassDeclaration> declarationIterator = localClassDeclarations.iterator();

				list.add(new TypeDeclarationStatement(declaration));
				declarationIterator.next();
				declarationIterator.remove();

				while (declarationIterator.hasNext() && ((declaration = declarationIterator.next()).getFirstLineNumber() <= lineNumber)) {
				list.add(new TypeDeclarationStatement(declaration));
				declarationIterator.remove();
				}

				if (statements.isList()) {
					list.addAll(statements.getList());
				} else {
					list.add(statements.getFirst());
				}
				statements = list;
			} else {
				statements.accept(this);
			}
		}
	}

	return statements;
}

func (v *AddLocalClassDeclarationVisitor)  VisitStatements( list intmod.IStatements) {
	if (!localClassDeclarations.isEmpty() && !list.isEmpty()) {
		ListIterator<Statement> statementIterator = list.listIterator();
		Iterator<ClassFileClassDeclaration> declarationIterator = localClassDeclarations.iterator();
		ClassFileClassDeclaration declaration = declarationIterator.next();

		while (statementIterator.hasNext()) {
			Statement statement = statementIterator.next();

			searchFirstLineNumberVisitor.init();
			statement.accept(searchFirstLineNumberVisitor);

			if (searchFirstLineNumberVisitor.getLineNumber() != -1) {
				lineNumber = searchFirstLineNumberVisitor.getLineNumber();
			}

			while (declaration.getFirstLineNumber() <= lineNumber) {
				statementIterator.previous();
				statementIterator.add(new TypeDeclarationStatement(declaration));
				statementIterator.next();
				declarationIterator.remove();

				if (!declarationIterator.hasNext()) {
					return;
				}

				declaration = declarationIterator.next();
			}

			type MemberDeclarationComparator struct{
			}

			func (c *MemberDeclarationComparator) Compare(md1, md2 intsrv.IClassFileMemberDeclaration) int {
				return md1.FirstLineNumber() - md2.FirstLineNumber();
			}
