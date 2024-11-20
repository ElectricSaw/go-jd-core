package localvariable

//import (
//	"bitbucket.org/coontec/go-jd-core/class/model/model/declaration"
//	"bitbucket.org/coontec/go-jd-core/class/model/model/expression"
//	"bitbucket.org/coontec/go-jd-core/class/model/model/statement"
//	_type "bitbucket.org/coontec/go-jd-core/class/model/model/type"
//	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
//	"bitbucket.org/coontec/go-jd-core/class/service/converter/visitor"
//	srvstat "bitbucket.org/coontec/go-jd-core/class/service/converter/model/model/statement"
//	"fmt"
//	"strings"
//	"unicode"
//)
//
//var GlobalAbstractLocalVariableComparator = NewAbstractLocalVariableComparator()
//var CapitalizedJavaLanguageKeywords = []string{
//	"Abstract", "Continue", "For", "New", "Switch", "Assert", "Default", "Goto", "Package", "Synchronized",
//	"Boolean", "Do", "If", "Private", "This", "Break", "Double", "Implements", "Protected", "Throw", "Byte", "Else",
//	"Import", "Public", "Throws", "Case", "Enum", "Instanceof", "Return", "Transient", "Catch", "Extends", "Int",
//	"Short", "Try", "Char", "Final", "Interface", "Static", "Void", "Class", "Finally", "Long", "Strictfp",
//	"Volatile", "Const", "Float", "Native", "Super", "While"}
//
//func NewFrame(parent *Frame, stat *statement.Statements) *Frame {
//	return &Frame{
//		localVariableArray: make([]ILocalVariable, 0),
//		newExpressions: make(map[expression.NewExpression]ILocalVariable),
//		children:           make([]Frame, 0),
//		parent:             parent,
//		stat:               stat,
//	}
//}
//
//type Frame struct {
//	localVariableArray     []ILocalVariable
//	newExpressions         map[expression.NewExpression]ILocalVariable
//	children               []Frame
//	parent                 *Frame
//	stat                   *statement.Statements
//	exceptionLocalVariable ILocalVariable
//}
//
//func (f *Frame) AddLocalVariable(lv ILocalVariable) {
//	// Java의 assert 대체 코드
//	if lv.Next() != nil {
//		fmt.Println("Frame.AddLocalVariable: add local variable failed")
//		return
//	}
//
//	index := lv.Index()
//
//	// 배열 크기 늘리기
//	if index >= len(f.localVariableArray) {
//		newArray := make([]ILocalVariable, index*2)
//		copy(newArray, f.localVariableArray)
//		f.localVariableArray = newArray
//	}
//
//	next := f.localVariableArray[index]
//
//	// 중복 추가 방지
//	if next != lv {
//		f.localVariableArray[index] = lv
//		lv.SetNext(next)
//		lv.SetFrame(f)
//	}
//}
//
//func (f *Frame) LocalVariable(index int) ILocalVariable {
//	if index < len(f.localVariableArray) {
//		lv := f.localVariableArray[index]
//		if lv != nil {
//			return lv
//		}
//	}
//	return f.parent.LocalVariable(index)
//}
//
//func (f *Frame) Parent() *Frame {
//	return f.parent
//}
//
//func (f *Frame) SetExceptionLocalVariable(e ILocalVariable) {
//	f.exceptionLocalVariable = e
//}
//
//func (f *Frame)  MergeLocalVariable(typeBounds map[string]_type.IType, localVariableMaker utils.LocalVariableMaker, lv ILocalVariable) {
//	index := lv.Index();
//	var  alvToMerge ILocalVariable;
//
//	if index < len(f.localVariableArray) {
//		alvToMerge = f.localVariableArray[index];
//	} else {
//		alvToMerge = nil;
//	}
//
//	if alvToMerge != nil {
//		if !lv.IsAssignableFromWithVariable(typeBounds, alvToMerge) && !alvToMerge.IsAssignableFromWithVariable(typeBounds, lv) {
//			alvToMerge = nil;
//		} else if (lv.Name() != "") && (alvToMerge.Name() != "") && !(lv.Name() == alvToMerge.Name()) {
//			alvToMerge = nil;
//		}
//	}
//
//	if alvToMerge == nil {
//		if f.children != nil {
//			for _, child := range f.children {
//				child.MergeLocalVariable(typeBounds, localVariableMaker, lv)
//			}
//		}
//	} else if lv != alvToMerge {
//		for _, reference  := range alvToMerge.References() {
//			reference.SetLocalVariable(lv);
//			lv.AddReference(reference)
//		}
//
//		lv.SetFromOffSet(alvToMerge.FromOffSet());
//
//		typ := lv.Type();
//		alvToMerype := alvToMerge.Type();
//
//		if lv.IsAssignableFromWithVariable(typeBounds, alvToMerge) || localVariableMaker.IsCompatible(lv, alvToMerge.Type()) {
//			if typ.IsPrimitiveType() {
//				plv := lv.(*PrimitiveLocalVariable);
//				plvToMerype := alvToMerge.(*PrimitiveLocalVariable)
//				t := utils.CommonPrimitiveType(plv.Type().(*_type.PrimitiveType), plvToMerype.Type().(*_type.PrimitiveType));
//
//				if t == nil {
//					t = _type.PtTypeInt;
//				}
//
//				plv.SetType(t.CreateType(typ.Dimension()).(*_type.PrimitiveType));
//			}
//		} else {
//			if typ.IsPrimitiveType() {
//				plv := lv.(*PrimitiveLocalVariable);
//
//				if alvToMerge.IsAssignableFromWithVariable(typeBounds, lv) || localVariableMaker.IsCompatible(alvToMerge, lv.Type()) {
//					plv.SetType(alvToMerype.(*_type.PrimitiveType));
//				} else {
//					plv.SetType(_type.PtTypeInt);
//				}
//			} else if typ.IsObjectType() {
//				olv := lv.(*ObjectLocalVariable);
//
//				if alvToMerge.IsAssignableFromWithVariable(typeBounds, lv) || localVariableMaker.IsCompatible(alvToMerge, lv.Type()) {
//					olv.SetType(typeBounds, alvToMerype);
//				} else {
//					dimension := alvToMerge.Dimension()
//					if lv.Dimension() >= alvToMerge.Dimension() {
//						dimension = lv.Dimension();
//					}
//					olv.SetType(typeBounds, _type.OtTypeObject.CreateType(dimension));
//				}
//			}
//		}
//
//f.		localVariableArray[index] = alvToMerge.Next();
//	}
//}
//
//func (f *Frame)  RemoveLocalVariable(lv ILocalVariable) {
//	index := lv.Index();
//	var  alvToRemove ILocalVariable;
//
//	if (index < len(f.localVariableArray)) && (f.localVariableArray[index] == lv) {
//		alvToRemove = lv;
//	} else {
//		alvToRemove = nil;
//	}
//
//	if alvToRemove == nil {
//		if f.children != nil {
//			for _, child := range f.children {
//				child.removeLocalVariable(lv);
//			}
//		}
//	} else {
//		f.localVariableArray[index] = alvToRemove.Next();
//		alvToRemove.SetNext(nil);
//	}
//}
//
//func (f *Frame)  AddChild( child Frame) {
//	if f.children == nil {
//		f.children = make([]Frame, 0)
//	}
//	f.children = append(f.children, child);
//}
//
//func (f *Frame)  Close() {
//	// Update type for 'new' expression
//	if f.newExpressions != nil {
//		for key, value := range f.newExpressions{
//			 ot1 := key.ObjectType();
//			 ot2 := value.Type().(_type.IObjectType)
//
//			if (ot1.TypeArguments() == nil) && (ot2.TypeArguments() != nil) {
//				key.SetObjectType(ot1.CreateTypeWithArgs(ot2.TypeArguments()));
//			}
//		}
//	}
//}
//
//func (f *Frame)  CreateNames(parentNames []string) {
//	names := make([]string, 0, len(parentNames))
//	copy(names, parentNames)
//	types := make(map[_type.IType]bool)
//	length := len(f.localVariableArray)
//
//	for i := 0 ; i < length; i++ {
//		lv := f.localVariableArray[i];
//
//		for lv != nil {
//			if _, ok := types[lv.Type()]; ok {
//				types[lv.Type()] = true // Non unique type
//			} else {
//				types[lv.Type()] = false // Unique type
//			}
//
//			if lv.Name() != "" {
//				if contains(names, lv.Name()) {
//					lv.SetName("")
//				} else {
//					names = append(names, lv.Name());
//				}
//			}
//
//			lv = lv.Next();
//		}
//	}
//
//	if f.exceptionLocalVariable != nil {
//		if _, ok := types[f.exceptionLocalVariable.Type()]; ok  {
//			types[f.exceptionLocalVariable.Type()] = true // Non unique type
//		} else {
//			types[f.exceptionLocalVariable.Type()] = false // Unique type
//		}
//	}
//
//	if len(types) != 0 {
//		visitor := NewGenerateLocalVariableNameVisitor(names, types);
//
//		for i:=0; i< length;i++ {
//			lv := f.localVariableArray[i];
//			for lv != nil {
//				if lv.Name() == "" {
//					lv.Type().(_type.TypeArgumentVisitable).AcceptTypeArgumentVisitor(visitor)
//					lv.SetName(visitor.Name())
//				}
//			}
//		}
//
//		if f.exceptionLocalVariable != nil {
//			f.exceptionLocalVariable.Type().(_type.TypeArgumentVisitable).AcceptTypeArgumentVisitor(visitor);
//			f.exceptionLocalVariable.SetName(visitor.Name());
//		}
//	}
//
//	// Recursive call
//	if f.children != nil {
//		for _, child := range f.children {
//			child.CreateNames(names);
//		}
//	}
//}
//
//func (f *Frame)  UpdateLocalVariableInForStatements(typeMaker *utils.TypeMaker) {
//	// Recursive call first
//	if f.children != nil {
//		for _, child := range f.children {
//			child.UpdateLocalVariableInForStatements(typeMaker);
//		}
//	}
//
//	// Split local variable ranges in init 'for' statements
//	searchLocalVariableVisitor := visitor.NewSearchLocalVariableVisitor();
//	undeclaredInExpressionStatements := make([]ILocalVariable, 0)
//
//	for _, stat := range f.stat.Statements {
//		if stat.IsForStatement() {
//			if stat.Init() == nil {
//				if stat.Condition() != nil {
//					searchLocalVariableVisitor.Init();
//					stat.Condition().Accept(searchLocalVariableVisitor);
//					for _, variable := range searchLocalVariableVisitor.Variables() {
//						undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable);
//					}
//				}
//				if stat.Update() != nil {
//					searchLocalVariableVisitor.Init();
//					stat.Update().Accept(searchLocalVariableVisitor);
//					for _, variable := range searchLocalVariableVisitor.Variables() {
//						undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable);
//					}
//				}
//				if stat.Statements() != nil {
//					searchLocalVariableVisitor.Init();
//					stat.Statements().Accept(searchLocalVariableVisitor);
//					for _, variable := range searchLocalVariableVisitor.Variables() {
//						undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable);
//					}
//				}
//			}
//		} else {
//			searchLocalVariableVisitor.Init();
//			stat.Accept(searchLocalVariableVisitor);
//			for _, variable := range searchLocalVariableVisitor.Variables() {
//				undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable);
//			}
//		}
//	}
//
//	searchUndeclaredLocalVariableVisitor := visitor.NewSearchUndeclaredLocalVariableVisitor();
//	undeclaredInForStatements = make(map[ILocalVariable][]srvstat.ClassFileForStatement)
//
//	for _, stat := range f.stat.Statements {
//		if stat.IsForStatement() {
//			fs := stat.(srvstat.ClassFileForStatement)
//
//			if (fs.Init() != nil) {
//				searchUndeclaredLocalVariableVisitor.Init();
//				fs.Init().Accept(searchUndeclaredLocalVariableVisitor);
//				searchUndeclaredLocalVariableVisitor.Variables().removeAll(undeclaredInExpressionStatements);
//
//				for (AbstractLocalVariable lv : searchUndeclaredLocalVariableVisitor.Variables()) {
//					List<ClassFileForStatement> list = undeclaredInForStatements.(lv);
//					if (list == nil) {
//						undeclaredInForStatements.put(lv, list = new ArrayList<>());
//					}
//					list.add(fs);
//				}
//			}
//		}
//	}
//
//	if (!undeclaredInForStatements.isEmpty()) {
//		CreateLocalVariableVisitor createLocalVariableVisitor = new CreateLocalVariableVisitor(typeMaker);
//
//		for (Map.Entry<AbstractLocalVariable, List<ClassFileForStatement>> entry : undeclaredInForStatements.entrySet()) {
//			List<ClassFileForStatement> listFS = entry.Value();
//
//			// Split local variable range
//			AbstractLocalVariable lv = entry.Key();
//			Iterator<ClassFileForStatement> iteratorFS = listFS.iterator();
//			ClassFileForStatement firstFS = iteratorFS.next();
//
//			while (iteratorFS.hasNext()) {
//				createNewLocalVariable(createLocalVariableVisitor, iteratorFS.next(), lv);
//			}
//
//			if (lv.Frame() == this) {
//				lv.SetFromOffSet(firstFS.FromOffSet());
//				lv.SetToOffSet(firstFS.ToOffSet(), true);
//			} else {
//				createNewLocalVariable(createLocalVariableVisitor, firstFS, lv);
//
//				if (lv.References().isEmpty()) {
//					lv.Frame().removeLocalVariable(lv);
//				}
//			}
//		}
//	}
//}
//
//func (f *Frame)  createNewLocalVariable(CreateLocalVariableVisitor createLocalVariableVisitor, ClassFileForStatement fs,  lv ILocalVariable) {
//	int fromOffSet = fs.FromOffSet(), toOffSet = fs.ToOffSet();
//	createLocalVariableVisitor.init(lv.Index(), fromOffSet);
//	lv.accept(createLocalVariableVisitor);
//	AbstractLocalVariable newLV = createLocalVariableVisitor.LocalVariable();
//
//	newLV.SetToOffSet(toOffSet, true);
//	addLocalVariable(newLV);
//	Iterator<LocalVariableReference> iteratorLVR = lv.References().iterator();
//
//	while (iteratorLVR.hasNext()) {
//		LocalVariableReference lvr = iteratorLVR.next();
//		int offSet = ((ClassFileLocalVariableReferenceExpression) lvr).OffSet();
//
//		if ((fromOffSet <= offSet) && (offSet <= toOffSet)) {
//			lvr.SetLocalVariable(newLV);
//			newLV.addReference(lvr);
//			iteratorLVR.remove();
//		}
//	}
//}
//
//func (f *Frame)  createDeclarations( containsLineNumber bool) {
//	// Create inline declarations
//	createInlineDeclarations();
//
//	// Create start-block declarations
//	createStartBlockDeclarations();
//
//	// Merge declarations
//	if (containsLineNumber) {
//		mergeDeclarations();
//	}
//
//	// Recursive call
//	if (children != nil) {
//		for (Frame child : children) {
//			child.createDeclarations(containsLineNumber);
//		}
//	}
//}
//
//func (f *Frame)  createInlineDeclarations() {
//	HashMap<Frame, HashSet<AbstractLocalVariable>> map = createMapForInlineDeclarations();
//
//	if (!map.isEmpty()) {
//		SearchUndeclaredLocalVariableVisitor searchUndeclaredLocalVariableVisitor = new SearchUndeclaredLocalVariableVisitor();
//
//		for (Map.Entry<Frame, HashSet<AbstractLocalVariable>> entry : map.entrySet()) {
//			Statements statements = entry.Key().statements;
//			ListIterator<Statement> iterator = statements.listIterator();
//			HashSet<AbstractLocalVariable> undeclaredLocalVariables = entry.Value();
//
//			while (iterator.hasNext()) {
//				Statement statement = iterator.next();
//
//				searchUndeclaredLocalVariableVisitor.init();
//				statement.accept(searchUndeclaredLocalVariableVisitor);
//
//				HashSet<AbstractLocalVariable> undeclaredLocalVariablesInStatement = searchUndeclaredLocalVariableVisitor.Variables();
//				undeclaredLocalVariablesInStatement.retainAll(undeclaredLocalVariables);
//
//				if (!undeclaredLocalVariablesInStatement.isEmpty()) {
//					int index1 = iterator.nextIndex();
//
//					if (statement.isExpressionStatement()) {
//						createInlineDeclarations(undeclaredLocalVariables, undeclaredLocalVariablesInStatement, iterator, (ExpressionStatement)statement);
//					} else if (statement.isForStatement()) {
//						createInlineDeclarations(undeclaredLocalVariables, undeclaredLocalVariablesInStatement, (ClassFileForStatement)statement);
//					}
//
//					if (!undeclaredLocalVariablesInStatement.isEmpty()) {
//						// Set the cursor before current statement
//						int index2 = iterator.nextIndex() + undeclaredLocalVariablesInStatement.size();
//
//						while (iterator.nextIndex() >= index1) {
//							iterator.previous();
//						}
//
//						DefaultList<AbstractLocalVariable> sorted = new DefaultList<>(undeclaredLocalVariablesInStatement);
//						sorted.sort(ABSTRACT_LOCAL_VARIABLE_COMPARATOR);
//
//						for (AbstractLocalVariable lv : sorted) {
//							// Add declaration before current statement
//							iterator.add(new LocalVariableDeclarationStatement(lv.Type(), new ClassFileLocalVariableDeclarator(lv)));
//							lv.SetDeclared(true);
//							undeclaredLocalVariables.remove(lv);
//						}
//
//						// ReSet the cursor after current statement
//						while (iterator.nextIndex() < index2) {
//							iterator.next();
//						}
//					}
//				}
//
//				if (undeclaredLocalVariables.isEmpty()) {
//					break;
//				}
//			}
//		}
//	}
//}
//
//func (f *Frame) <Frame, HashSet<AbstractLocalVariable>> createMapForInlineDeclarations() HashMap{
//	HashMap<Frame, HashSet<AbstractLocalVariable>> map = new HashMap<>();
//	int i = localVariableArray.length;
//
//	while (i-- > 0) {
//		AbstractLocalVariable lv = localVariableArray[i];
//
//		while (lv != nil) {
//			if ((this == lv.Frame()) && !lv.isDeclared()) {
//				HashSet<AbstractLocalVariable> variablesToDeclare = map.(this);
//				if (variablesToDeclare == nil) {
//					map.put(this, variablesToDeclare = new HashSet<>());
//				}
//				variablesToDeclare.add(lv);
//			}
//			lv = lv.Next();
//		}
//	}
//
//	return map;
//}
//
//func (f *Frame)  createInlineDeclarationsvoid(
//	HashSet<AbstractLocalVariable> undeclaredLocalVariables, HashSet<AbstractLocalVariable> undeclaredLocalVariablesInStatement,
//	ListIterator<Statement> iterator, ExpressionStatement es) {
//
//	if (es.Expression().isBinaryOperatorExpression()) {
//		Expression boe = es.Expression();
//
//		if (boe.Operator().equals("=")) {
//			Expressions expressions = new Expressions();
//
//			splitMultiAssignment(Integer.MAX_VALUE, undeclaredLocalVariablesInStatement, expressions, boe);
//			iterator.remove();
//
//			for (Expression exp : expressions) {
//				iterator.add(newDeclarationStatement(undeclaredLocalVariables, undeclaredLocalVariablesInStatement, exp));
//			}
//
//			if (expressions.isEmpty()) {
//				iterator.add(es);
//			}
//		}
//	}
//}
//
//func (f *Frame)  splitMultiAssignmentExpression(
//	int toOffSet, HashSet<AbstractLocalVariable> undeclaredLocalVariablesInStatement, List<Expression> expressions, Expression expression) {
//
//	if (expression.isBinaryOperatorExpression() && expression.Operator().equals("=")) {
//		Expression rightExpression = splitMultiAssignment(toOffSet, undeclaredLocalVariablesInStatement, expressions, expression.RightExpression());
//
//		if (expression.LeftExpression().isLocalVariableReferenceExpression()) {
//			ClassFileLocalVariableReferenceExpression lvre = (ClassFileLocalVariableReferenceExpression)expression.LeftExpression();
//			AbstractLocalVariable localVariable = lvre.LocalVariable();
//
//			if (undeclaredLocalVariablesInStatement.contains(localVariable) && (localVariable.ToOffSet() <= toOffSet)) {
//				// Split multi assignment
//				if (rightExpression == expression.RightExpression()) {
//					expressions.add(expression);
//				} else {
//					expressions.add(new BinaryOperatorExpression(expression.LineNumber(), expression.Type(), lvre, "=", rightExpression, expression.Priority()));
//				}
//				// Return local variable
//				return lvre;
//			}
//		}
//	}
//
//	return expression;
//}
//
//func (f *Frame)  newDeclarationStatementLocalVariableDeclarationStatement(
//	HashSet<AbstractLocalVariable> undeclaredLocalVariables, HashSet<AbstractLocalVariable> undeclaredLocalVariablesInStatement, Expression boe) {
//
//	ClassFileLocalVariableReferenceExpression reference = (ClassFileLocalVariableReferenceExpression)boe.LeftExpression();
//	AbstractLocalVariable localVariable = reference.LocalVariable();
//
//	undeclaredLocalVariables.remove(localVariable);
//	undeclaredLocalVariablesInStatement.remove(localVariable);
//	localVariable.SetDeclared(true);
//
//	Type type = localVariable.Type();
//	VariableInitializer variableInitializer;
//
//	if (boe.RightExpression().isNewInitializedArray()) {
//		if (type.isObjectType() && (((ObjectType)type).TypeArguments() != nil)) {
//		variableInitializer = new ExpressionVariableInitializer(boe.RightExpression());
//		} else {
//		variableInitializer = ((NewInitializedArray) boe.RightExpression()).ArrayInitializer();
//		}
//	} else {
//		variableInitializer = new ExpressionVariableInitializer(boe.RightExpression());
//	}
//
//	return new LocalVariableDeclarationStatement(type, new ClassFileLocalVariableDeclarator(boe.LineNumber(), reference.LocalVariable(), variableInitializer));
//}
//
//func (f *Frame)  createInlineDeclarationsvoid(
//	HashSet<AbstractLocalVariable> undeclaredLocalVariables, HashSet<AbstractLocalVariable> undeclaredLocalVariablesInStatement, ClassFileForStatement fs) {
//
//	BaseExpression init = fs.Init();
//
//	if (init != nil) {
//		Expressions expressions = new Expressions();
//		int toOffSet = fs.ToOffSet();
//
//		if (init.isList()) {
//			for (Expression exp : init) {
//				splitMultiAssignment(toOffSet, undeclaredLocalVariablesInStatement, expressions, exp);
//				if (expressions.isEmpty()) {
//					expressions.add(exp);
//				}
//			}
//		} else {
//			splitMultiAssignment(toOffSet, undeclaredLocalVariablesInStatement, expressions, init.First());
//			if (expressions.isEmpty()) {
//				expressions.add(init.First());
//			}
//		}
//
//		if (expressions.size() == 1) {
//			updateForStatement(undeclaredLocalVariables, undeclaredLocalVariablesInStatement, fs, expressions.First());
//		} else {
//			updateForStatement(undeclaredLocalVariables, undeclaredLocalVariablesInStatement, fs, expressions);
//		}
//	}
//}
//
//func (f *Frame)  updateForStatementvoid(
//	HashSet<AbstractLocalVariable> undeclaredLocalVariables, HashSet<AbstractLocalVariable> undeclaredLocalVariablesInStatement,
//	ClassFileForStatement forStatement, Expression init) {
//
//	if (!init.isBinaryOperatorExpression()) {
//		return;
//	}
//
//	if (!init.LeftExpression().isLocalVariableReferenceExpression()) {
//		return;
//	}
//
//	ClassFileLocalVariableReferenceExpression reference = (ClassFileLocalVariableReferenceExpression)init.LeftExpression();
//	AbstractLocalVariable localVariable = reference.LocalVariable();
//
//	if (localVariable.isDeclared() || (localVariable.ToOffSet() > forStatement.ToOffSet())) {
//		return;
//	}
//
//	undeclaredLocalVariables.remove(localVariable);
//	undeclaredLocalVariablesInStatement.remove(localVariable);
//	localVariable.SetDeclared(true);
//
//	VariableInitializer variableInitializer = init.RightExpression().isNewInitializedArray() ?
//	((NewInitializedArray)init.RightExpression()).ArrayInitializer() :
//	new ExpressionVariableInitializer(init.RightExpression());
//
//	forStatement.SetDeclaration(new LocalVariableDeclaration(localVariable.Type(), new ClassFileLocalVariableDeclarator(init.LineNumber(), reference.LocalVariable(), variableInitializer)));
//	forStatement.SetInit(nil);
//}
//
//func (f *Frame)  updateForStatementvoid(
//	HashSet<AbstractLocalVariable> variablesToDeclare, HashSet<AbstractLocalVariable> foundVariables,
//	ClassFileForStatement forStatement, Expressions init) {
//
//	DefaultList<Expression> boes = new DefaultList<>();
//	DefaultList<AbstractLocalVariable> localVariables = new DefaultList<>();
//	Type type0 = nil, type1 = nil;
//	int minDimension = 0, maxDimension = 0;
//
//	for (Expression expression : init) {
//		if (!expression.isBinaryOperatorExpression()) {
//			return;
//		}
//		if (!expression.LeftExpression().isLocalVariableReferenceExpression()) {
//			return;
//		}
//
//		AbstractLocalVariable localVariable = ((ClassFileLocalVariableReferenceExpression)expression.LeftExpression()).LocalVariable();
//
//		if (localVariable.isDeclared() || (localVariable.ToOffSet() > forStatement.ToOffSet())) {
//			return;
//		}
//
//		if (type1 == nil) {
//			type1 = localVariable.Type();
//			type0 = type1.createType(0);
//			minDimension = maxDimension = type1.Dimension();
//		} else {
//			Type type2 = localVariable.Type();
//
//			if (type1.isPrimitiveType() && type2.isPrimitiveType()) {
//				Type type = PrimitiveTypeUtil.CommonPrimitiveType((PrimitiveType)type1, (PrimitiveType)type2);
//
//				if (type == nil) {
//					return;
//				}
//
//				type0 = type;
//				type1 = type.createType(type1.Dimension());
//				type2 = type.createType(type2.Dimension());
//			} else if (!type1.equals(type2) && !type0.equals(type2.createType(0))) {
//				return;
//			}
//
//			int dimension = type2.Dimension();
//
//			if (minDimension > dimension) {
//				minDimension = dimension;
//			}
//			if (maxDimension < dimension) {
//				maxDimension = dimension;
//			}
//		}
//
//		localVariables.add(localVariable);
//		boes.add(expression);
//	}
//
//	for (AbstractLocalVariable lv : localVariables) {
//		variablesToDeclare.remove(lv);
//		foundVariables.remove(lv);
//		lv.SetDeclared(true);
//	}
//
//	if (minDimension == maxDimension) {
//		forStatement.SetDeclaration(new LocalVariableDeclaration(type1, createDeclarators1(boes, false)));
//	} else {
//		forStatement.SetDeclaration(new LocalVariableDeclaration(type0, createDeclarators1(boes, true)));
//	}
//
//	forStatement.SetInit(nil);
//}
//
//func (f *Frame)  createDeclarators1(DefaultList<Expression> boes, boolean SetDimension) LocalVariableDeclarators{
//	LocalVariableDeclarators declarators = new LocalVariableDeclarators(boes.size());
//
//	for (Expression boe : boes) {
//		ClassFileLocalVariableReferenceExpression reference = (ClassFileLocalVariableReferenceExpression) boe.LeftExpression();
//		VariableInitializer variableInitializer = boe.RightExpression().isNewInitializedArray() ?
//		((NewInitializedArray) boe.RightExpression()).ArrayInitializer() :
//		new ExpressionVariableInitializer(boe.RightExpression());
//		LocalVariableDeclarator declarator = new ClassFileLocalVariableDeclarator(boe.LineNumber(), reference.LocalVariable(), variableInitializer);
//
//		if (SetDimension) {
//			declarator.SetDimension(reference.LocalVariable().Dimension());
//		}
//
//		declarators.add(declarator);
//	}
//
//	return declarators;
//}
//
//func (f *Frame)  createStartBlockDeclarations() void{
//	int addIndex = -1;
//	int i = localVariableArray.length;
//
//	while (i-- > 0) {
//		AbstractLocalVariable lv = localVariableArray[i];
//
//		while (lv != nil) {
//			if (!lv.isDeclared()) {
//				if (addIndex == -1) {
//					addIndex = AddIndex();
//				}
//				statements.add(addIndex, new LocalVariableDeclarationStatement(lv.Type(), new ClassFileLocalVariableDeclarator(lv)));
//				lv.SetDeclared(true);
//			}
//
//			lv = lv.Next();
//		}
//	}
//}
//
//func (f *Frame)  AddIndex() int{
//	int addIndex = 0;
//
//	if (parent.parent == nil) {
//		// Insert declarations after 'super' call invocation => Search index of SuperConstructorInvocationExpression.
//		int len = statements.size();
//
//		while (addIndex < len) {
//			Statement statement = statements.(addIndex++);
//			if (statement.isExpressionStatement()) {
//				Expression expression = statement.Expression();
//				if (expression.isSuperConstructorInvocationExpression() || expression.isConstructorInvocationExpression()) {
//					break;
//				}
//			}
//		}
//
//		if (addIndex >= len) {
//			addIndex = 0;
//		}
//	}
//
//	return addIndex;
//}
//
//func (f *Frame)  mergeDeclarations() void{
//	int size = statements.size();
//
//	if (size > 1) {
//		DefaultList<LocalVariableDeclarationStatement> declarations = new DefaultList<>();
//		ListIterator<Statement> iterator = statements.listIterator();
//
//		while (iterator.hasNext()) {
//			Statement previous;
//
//			do {
//				previous = iterator.next();
//			} while (!previous.isLocalVariableDeclarationStatement() && iterator.hasNext());
//
//			if (previous.isLocalVariableDeclarationStatement()) {
//				LocalVariableDeclarationStatement lvds1 = (LocalVariableDeclarationStatement) previous;
//				Type type1 = lvds1.Type();
//				Type type0 = type1.createType(0);
//				int minDimension = type1.Dimension();
//				int maxDimension = minDimension;
//				int lineNumber1 = lvds1.LocalVariableDeclarators().LineNumber();
//
//				declarations.clear();
//				declarations.add(lvds1);
//
//				while (iterator.hasNext()) {
//					Statement statement = iterator.next();
//
//					if (!statement.isLocalVariableDeclarationStatement()) {
//						iterator.previous();
//						break;
//					}
//
//					LocalVariableDeclarationStatement lvds2 = (LocalVariableDeclarationStatement) statement;
//					int lineNumber2 = lvds2.LocalVariableDeclarators().LineNumber();
//
//					if (lineNumber1 != lineNumber2) {
//						iterator.previous();
//						break;
//					}
//
//					lineNumber1 = lineNumber2;
//
//					Type type2 = lvds2.Type();
//
//					if (type1.isPrimitiveType() && type2.isPrimitiveType()) {
//						Type type = PrimitiveTypeUtil.CommonPrimitiveType((PrimitiveType)type1, (PrimitiveType)type2);
//
//						if (type == nil) {
//							iterator.previous();
//							break;
//						}
//
//						type0 = type;
//						type1 = type.createType(type1.Dimension());
//						type2 = type.createType(type2.Dimension());
//					} else if (!type1.equals(type2) && !type0.equals(type2.createType(0))) {
//						iterator.previous();
//						break;
//					}
//
//					int dimension = type2.Dimension();
//
//					if (minDimension > dimension) {
//						minDimension = dimension;
//					}
//					if (maxDimension < dimension) {
//						maxDimension = dimension;
//					}
//
//					declarations.add(lvds2);
//				}
//
//				int declarationSize = declarations.size();
//
//				if (declarationSize > 1) {
//					while (--declarationSize > 0) {
//						iterator.previous();
//						iterator.remove();
//					}
//
//					iterator.previous();
//
//					if (minDimension == maxDimension) {
//						iterator.Set(new LocalVariableDeclarationStatement(type1, createDeclarators2(declarations, false)));
//					} else {
//						iterator.Set(new LocalVariableDeclarationStatement(type0, createDeclarators2(declarations, true)));
//					}
//
//					iterator.next();
//				}
//			}
//		}
//	}
//}
//
//func (f *Frame)  createDeclarators2(DefaultList<LocalVariableDeclarationStatement> declarations, boolean SetDimension) LocalVariableDeclarators{
//	LocalVariableDeclarators declarators = new LocalVariableDeclarators(declarations.size());
//
//	for (LocalVariableDeclarationStatement declaration : declarations) {
//		LocalVariableDeclarator declarator = (LocalVariableDeclarator)declaration.LocalVariableDeclarators();
//
//		if (SetDimension) {
//			declarator.SetDimension(declaration.Type().Dimension());
//		}
//
//		declarators.add(declarator);
//	}
//
//	return declarators;
//}
//
//func NewGenerateLocalVariableNameVisitor(blackListNames []string, types map[_type.IType]bool) *GenerateLocalVariableNameVisitor {
//	return &GenerateLocalVariableNameVisitor{
//		blackListNames: blackListNames,
//		types:          types,
//	}
//}
//
//var IntegerNames = []string{"i", "j", "k", "m", "n"}
//
//type GenerateLocalVariableNameVisitor struct {
//	sb             string
//	blackListNames []string
//	types          map[_type.IType]bool
//	name           string
//}
//
//func (c *GenerateLocalVariableNameVisitor) Name() string {
//	return c.name
//}
//
//func (c *GenerateLocalVariableNameVisitor) capitalize(str string) {
//	if str != "" {
//		length := len(str)
//		if length > 0 {
//			firstChar := str[0]
//
//			if unicode.IsUpper(rune(firstChar)) {
//				c.sb += str
//			} else {
//				c.sb += strings.ToUpper(string(firstChar))
//				if length > 1 {
//					c.sb += str[1:]
//				}
//			}
//		}
//	}
//}
//
//func (c *GenerateLocalVariableNameVisitor) uncapitalize(str string) {
//	if str != "" {
//		length := len(str)
//		if length > 0 {
//			firstChar := str[0]
//			if unicode.IsLower(rune(firstChar)) {
//				c.sb += str
//			} else {
//				c.sb += strings.ToLower(string(firstChar))
//				if length > 1 {
//					c.sb += str[1:]
//				}
//			}
//		}
//	}
//}
//
//func (c *GenerateLocalVariableNameVisitor) generate(typ _type.IType) {
//	length := len(c.sb)
//	counter := 1
//
//	if c.types[typ] {
//		c.sb += fmt.Sprintf("%d", counter)
//		counter++
//	}
//
//	c.name = c.sb
//
//	for contains(c.blackListNames, c.name) {
//		c.sb = c.sb[:length]
//		c.sb += fmt.Sprintf("%d", counter)
//		counter++
//		c.name = c.sb
//	}
//
//	c.blackListNames = append(c.blackListNames, c.name)
//}
//
//func (c *GenerateLocalVariableNameVisitor) VisitPrimitiveType(t *_type.PrimitiveType) {
//	c.sb = ""
//
//	switch t.Dimension() {
//	case _type.FlagByte:
//		c.sb += "b"
//	case _type.FlagChar:
//		c.sb += "c"
//	case _type.FlagDouble:
//		c.sb += "d"
//	case _type.FlagFloat:
//		c.sb += "f"
//	case _type.FlagInt:
//		for _, in := range IntegerNames {
//			if !contains(c.blackListNames, in) {
//				c.blackListNames = append(c.blackListNames, in)
//				return
//			}
//		}
//		c.sb += "i"
//	case _type.FlagLong:
//		c.sb += "l"
//	case _type.FlagShort:
//		c.sb += "s"
//	case _type.FlagBoolean:
//		c.sb += "bool"
//	}
//
//	c.generate(t)
//}
//
//func (c *GenerateLocalVariableNameVisitor) Visit(t _type.IType, str string) {
//	c.sb = ""
//
//	switch t.Dimension() {
//	case 0:
//		if str == "Class" {
//			c.sb += "clazz"
//		} else if str == "String" {
//			c.sb += "str"
//		} else if str == "Boolean" {
//			c.sb += "bool"
//		} else {
//			c.uncapitalize(str)
//			if contains(CapitalizedJavaLanguageKeywords, str) {
//				c.sb += "_"
//			}
//		}
//	default:
//		c.sb += "arrayOf"
//		c.capitalize(str)
//	}
//
//	c.generate(t)
//}
//
//func (c *GenerateLocalVariableNameVisitor) VisitObjectType(t *_type.ObjectType) {
//	c.Visit(t, t.Name())
//}
//
//func (c *GenerateLocalVariableNameVisitor) VisitInnerObjectType(t *_type.InnerObjectType) {
//	c.Visit(t, t.Name())
//}
//
//func (c *GenerateLocalVariableNameVisitor) VisitGenericType(t *_type.GenericType) {
//	c.Visit(t, t.Name())
//}
//
//func (c *GenerateLocalVariableNameVisitor) VisitTypeArguments(arguments *_type.TypeArguments) {}
//
//func (c *GenerateLocalVariableNameVisitor) VisitDiamondTypeArgument(argument *_type.DiamondTypeArgument) {
//}
//
//func (c *GenerateLocalVariableNameVisitor) VisitWildcardExtendsTypeArgument(argument *_type.WildcardExtendsTypeArgument) {
//}
//
//func (c *GenerateLocalVariableNameVisitor) VisitWildcardSuperTypeArgument(argument *_type.WildcardSuperTypeArgument) {
//}
//
//func (c *GenerateLocalVariableNameVisitor) VisitWildcardTypeArgument(argument *_type.WildcardTypeArgument) {
//}
//
//func NewAbstractLocalVariableComparator() *AbstractLocalVariableComparator {
//	return &AbstractLocalVariableComparator{}
//}
//
//type AbstractLocalVariableComparator struct {
//}
//
//func (c *AbstractLocalVariableComparator) Compare(alv1, alv2 ILocalVariable) int {
//	return alv1.Index() - alv2.Index()
//}
//
//func contains(list []string, value string) bool {
//	for _, v := range list {
//		if v == value {
//			return true
//		}
//	}
//	return false
//}
