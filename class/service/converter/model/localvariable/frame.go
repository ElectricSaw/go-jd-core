package localvariable

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	srvdecl "bitbucket.org/coontec/go-jd-core/class/service/converter/model/javasyntax/declaration"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/visitor"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
	"math"
	"sort"
	"strings"
	"unicode"
)

var GlobalAbstractLocalVariableComparator = NewAbstractLocalVariableComparator()
var CapitalizedJavaLanguageKeywords = []string{
	"Abstract", "Continue", "For", "New", "Switch", "Assert", "Default", "Goto", "Package", "Synchronized",
	"Boolean", "Do", "If", "Private", "This", "Break", "Double", "Implements", "Protected", "Throw", "Byte", "Else",
	"Import", "Public", "Throws", "Case", "Enum", "Instanceof", "Return", "Transient", "Catch", "Extends", "Int",
	"Short", "Try", "Char", "Final", "Interface", "Static", "Void", "Class", "Finally", "Long", "Strictfp",
	"Volatile", "Const", "Float", "Native", "Super", "While"}

func NewFrame(parent intsrv.IFrame, stat intmod.IStatements) intsrv.IFrame {
	return &Frame{
		localVariableArray: make([]intsrv.ILocalVariable, 0),
		newExpressions:     make(map[intmod.INewExpression]intsrv.ILocalVariable),
		children:           make([]intsrv.IFrame, 0),
		parent:             parent,
		stat:               stat,
	}
}

type Frame struct {
	localVariableArray     []intsrv.ILocalVariable
	newExpressions         map[intmod.INewExpression]intsrv.ILocalVariable
	children               []intsrv.IFrame
	parent                 intsrv.IFrame
	stat                   intmod.IStatements
	exceptionLocalVariable intsrv.ILocalVariable
}

func (f *Frame) Statements() intmod.IStatements {
	return f.stat
}

func (f *Frame) AddLocalVariable(lv intsrv.ILocalVariable) {
	// Java의 assert 대체 코드
	if lv.Next() != nil {
		fmt.Println("Frame.AddLocalVariable: add local variable failed")
		return
	}

	index := lv.Index()

	// 배열 크기 늘리기
	if index >= len(f.localVariableArray) {
		newArray := make([]intsrv.ILocalVariable, index*2)
		copy(newArray, f.localVariableArray)
		f.localVariableArray = newArray
	}

	next := f.localVariableArray[index]

	// 중복 추가 방지
	if next != lv {
		f.localVariableArray[index] = lv
		lv.SetNext(next)
		lv.SetFrame(f)
	}
}

func (f *Frame) LocalVariable(index int) intsrv.ILocalVariableReference {
	if index < len(f.localVariableArray) {
		lv := f.localVariableArray[index]
		if lv != nil {
			return lv
		}
	}
	return f.parent.LocalVariable(index)
}

func (f *Frame) Parent() intsrv.IFrame {
	return f.parent
}

func (f *Frame) SetExceptionLocalVariable(e intsrv.ILocalVariable) {
	f.exceptionLocalVariable = e
}

func (f *Frame) MergeLocalVariable(typeBounds map[string]intmod.IType, localVariableMaker intsrv.ILocalVariableMaker, lv intsrv.ILocalVariable) {
	index := lv.Index()
	var alvToMerge intsrv.ILocalVariable

	if index < len(f.localVariableArray) {
		alvToMerge = f.localVariableArray[index]
	} else {
		alvToMerge = nil
	}

	if alvToMerge != nil {
		if !lv.IsAssignableFromWithVariable(typeBounds, alvToMerge) && !alvToMerge.IsAssignableFromWithVariable(typeBounds, lv) {
			alvToMerge = nil
		} else if (lv.Name() != "") && (alvToMerge.Name() != "") && !(lv.Name() == alvToMerge.Name()) {
			alvToMerge = nil
		}
	}

	if alvToMerge == nil {
		if f.children != nil {
			for _, child := range f.children {
				child.MergeLocalVariable(typeBounds, localVariableMaker, lv)
			}
		}
	} else if lv != alvToMerge {
		for _, reference := range alvToMerge.References() {
			reference.SetLocalVariable(lv)
			lv.AddReference(reference)
		}

		lv.SetFromOffset(alvToMerge.FromOffset())

		typ := lv.Type()
		alvToMerype := alvToMerge.Type()

		if lv.IsAssignableFromWithVariable(typeBounds, alvToMerge) || localVariableMaker.IsCompatible(lv, alvToMerge.Type()) {
			if typ.IsPrimitiveType() {
				plv := lv.(*PrimitiveLocalVariable)
				plvToMerype := alvToMerge.(*PrimitiveLocalVariable)
				t := GetCommonPrimitiveType(plv.Type().(intmod.IPrimitiveType), plvToMerype.Type().(intmod.IPrimitiveType))

				if t == nil {
					t = _type.PtTypeInt
				}

				plv.SetType(t.CreateType(typ.Dimension()).(intmod.IPrimitiveType))
			}
		} else {
			if typ.IsPrimitiveType() {
				plv := lv.(*PrimitiveLocalVariable)

				if alvToMerge.IsAssignableFromWithVariable(typeBounds, lv) || localVariableMaker.IsCompatible(alvToMerge, lv.Type()) {
					plv.SetType(alvToMerype.(intmod.IPrimitiveType))
				} else {
					plv.SetType(_type.PtTypeInt)
				}
			} else if typ.IsObjectType() {
				olv := lv.(*ObjectLocalVariable)

				if alvToMerge.IsAssignableFromWithVariable(typeBounds, lv) || localVariableMaker.IsCompatible(alvToMerge, lv.Type()) {
					olv.SetType(typeBounds, alvToMerype)
				} else {
					dimension := alvToMerge.Dimension()
					if lv.Dimension() >= alvToMerge.Dimension() {
						dimension = lv.Dimension()
					}
					olv.SetType(typeBounds, _type.OtTypeObject.CreateType(dimension))
				}
			}
		}

		f.localVariableArray[index] = alvToMerge.Next()
	}
}

func (f *Frame) RemoveLocalVariable(lv intsrv.ILocalVariable) {
	index := lv.Index()
	var alvToRemove intsrv.ILocalVariable

	if (index < len(f.localVariableArray)) && (f.localVariableArray[index] == lv) {
		alvToRemove = lv
	} else {
		alvToRemove = nil
	}

	if alvToRemove == nil {
		if f.children != nil {
			for _, child := range f.children {
				child.RemoveLocalVariable(lv)
			}
		}
	} else {
		f.localVariableArray[index] = alvToRemove.Next()
		alvToRemove.SetNext(nil)
	}
}

func (f *Frame) AddChild(child intsrv.IFrame) {
	if f.children == nil {
		f.children = make([]intsrv.IFrame, 0)
	}
	f.children = append(f.children, child)
}

func (f *Frame) Close() {
	// Update type for 'new' expression
	if f.newExpressions != nil {
		for key, value := range f.newExpressions {
			ot1 := key.ObjectType()
			ot2 := value.Type().(intmod.IObjectType)

			if (ot1.TypeArguments() == nil) && (ot2.TypeArguments() != nil) {
				key.SetObjectType(ot1.CreateTypeWithArgs(ot2.TypeArguments()))
			}
		}
	}
}

func (f *Frame) CreateNames(parentNames []string) {
	names := make([]string, 0, len(parentNames))
	copy(names, parentNames)
	types := make(map[intmod.IType]bool)
	length := len(f.localVariableArray)

	for i := 0; i < length; i++ {
		lv := f.localVariableArray[i]

		for lv != nil {
			if _, ok := types[lv.Type()]; ok {
				types[lv.Type()] = true // Non unique type
			} else {
				types[lv.Type()] = false // Unique type
			}

			if lv.Name() != "" {
				if contains(names, lv.Name()) {
					lv.SetName("")
				} else {
					names = append(names, lv.Name())
				}
			}

			lv = lv.Next()
		}
	}

	if f.exceptionLocalVariable != nil {
		if _, ok := types[f.exceptionLocalVariable.Type()]; ok {
			types[f.exceptionLocalVariable.Type()] = true // Non unique type
		} else {
			types[f.exceptionLocalVariable.Type()] = false // Unique type
		}
	}

	if len(types) != 0 {
		visitor := NewGenerateLocalVariableNameVisitor(names, types)

		for i := 0; i < length; i++ {
			lv := f.localVariableArray[i]
			for lv != nil {
				if lv.Name() == "" {
					lv.Type().(intmod.ITypeArgumentVisitable).AcceptTypeArgumentVisitor(visitor)
					lv.SetName(visitor.Name())
				}
			}
		}

		if f.exceptionLocalVariable != nil {
			f.exceptionLocalVariable.Type().(intmod.ITypeArgumentVisitable).AcceptTypeArgumentVisitor(visitor)
			f.exceptionLocalVariable.SetName(visitor.Name())
		}
	}

	// Recursive call
	if f.children != nil {
		for _, child := range f.children {
			child.CreateNames(names)
		}
	}
}

func (f *Frame) UpdateLocalVariableInForStatements(typeMaker intsrv.ITypeMaker) {
	// Recursive call first
	if f.children != nil {
		for _, child := range f.children {
			child.UpdateLocalVariableInForStatements(typeMaker)
		}
	}

	// Split local variable ranges in init 'for' statements
	searchLocalVariableVisitor := visitor.NewSearchLocalVariableVisitor()
	undeclaredInExpressionStatements := make([]intsrv.ILocalVariable, 0)

	for _, stat := range f.stat.Statements().ToSlice() {
		if stat.IsForStatement() {
			if stat.Init() == nil {
				if stat.Condition() != nil {
					searchLocalVariableVisitor.Init()
					stat.Condition().Accept(searchLocalVariableVisitor)
					for _, variable := range searchLocalVariableVisitor.Variables() {
						undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable)
					}
				}
				if stat.Update() != nil {
					searchLocalVariableVisitor.Init()
					stat.Update().Accept(searchLocalVariableVisitor)
					for _, variable := range searchLocalVariableVisitor.Variables() {
						undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable)
					}
				}
				if stat.Statements() != nil {
					searchLocalVariableVisitor.Init()
					stat.Statements().Accept(searchLocalVariableVisitor)
					for _, variable := range searchLocalVariableVisitor.Variables() {
						undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable)
					}
				}
			}
		} else {
			searchLocalVariableVisitor.Init()
			stat.Accept(searchLocalVariableVisitor)
			for _, variable := range searchLocalVariableVisitor.Variables() {
				undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable)
			}
		}
	}

	searchUndeclaredLocalVariableVisitor := visitor.NewSearchUndeclaredLocalVariableVisitor()
	undeclaredInForStatements := make(map[intsrv.ILocalVariable][]intsrv.IClassFileForStatement)

	for _, stat := range f.stat.Statements().ToSlice() {
		if stat.IsForStatement() {
			fs := stat.(intsrv.IClassFileForStatement)

			if fs.Init() != nil {
				searchUndeclaredLocalVariableVisitor.Init()
				fs.Init().Accept(searchUndeclaredLocalVariableVisitor)
				searchUndeclaredLocalVariableVisitor.RemoveAll(undeclaredInExpressionStatements)

				for _, lv := range searchUndeclaredLocalVariableVisitor.Variables() {
					list := undeclaredInForStatements[lv]
					if list == nil {
						list = make([]intsrv.IClassFileForStatement, 0)
						undeclaredInForStatements[lv] = list
					}
					list = append(list, fs)
				}
			}
		}
	}

	if len(undeclaredInForStatements) != 0 {
		createLocalVariableVisitor := visitor.NewCreateLocalVariableVisitor(typeMaker)

		for lv, listFS := range undeclaredInForStatements {
			// Split local variable range
			firstFS := listFS[0]

			for i := 1; i < len(listFS); i++ {
				f.createNewLocalVariable(createLocalVariableVisitor, listFS[i], lv)
			}

			if lv.Frame() == f {
				lv.SetFromOffset(firstFS.FromOffset())
				lv.SetToOffsetWithForce(firstFS.ToOffset(), true)
			} else {
				f.createNewLocalVariable(createLocalVariableVisitor, firstFS, lv)

				if len(lv.References()) == 0 {
					lv.Frame().RemoveLocalVariable(lv)
				}
			}
		}
	}
}

func (f *Frame) createNewLocalVariable(createLocalVariableVisitor *visitor.CreateLocalVariableVisitor,
	fs intsrv.IClassFileForStatement, lv intsrv.ILocalVariable) {
	fromOffset := fs.FromOffset()
	toOffset := fs.ToOffset()
	createLocalVariableVisitor.Init(lv.Index(), fromOffset)
	lv.Accept(createLocalVariableVisitor)
	newLV := createLocalVariableVisitor.LocalVariable()

	newLV.SetToOffsetWithForce(toOffset, true)
	f.AddLocalVariable(newLV)
	iteratorLVR := util.NewIteratorWithSlice(lv.References())

	for iteratorLVR.HasNext() {
		lvr := iteratorLVR.Next()
		offset := lvr.(intsrv.IClassFileLocalVariableReferenceExpression).Offset()

		if fromOffset <= offset && offset <= toOffset {
			lvr.SetLocalVariable(newLV)
			newLV.AddReference(lvr)
			_ = iteratorLVR.Remove()
		}
	}
}

func (f *Frame) CreateDeclarations(containsLineNumber bool) {
	// Create inline declarations
	f.createInlineDeclarations()

	// Create start-block declarations
	f.createStartBlockDeclarations()

	// Merge declarations
	if containsLineNumber {
		f.mergeDeclarations()
	}

	// Recursive call
	if f.children != nil {
		for _, child := range f.children {
			child.CreateDeclarations(containsLineNumber)
		}
	}
}

func (f *Frame) createInlineDeclarations() {
	// FIXME: createInlineDeclarations() 메소드 재검토 필요 (정상동작 여부)
	mapped := f.createMapForInlineDeclarations()

	if len(mapped) != 0 {
		searchUndeclaredLocalVariableVisitor := visitor.NewSearchUndeclaredLocalVariableVisitor()

		for key, value := range mapped {
			statements := key.Statements()
			iterator := statements.ListIterator()
			undeclaredLocalVariables := util.NewSetWithSlice[intsrv.ILocalVariable](value)

			for iterator.HasNext() {
				state := iterator.Next()
				searchUndeclaredLocalVariableVisitor.Init()
				state.Accept(searchUndeclaredLocalVariableVisitor)

				//undeclaredLocalVariablesInStatement := searchUndeclaredLocalVariableVisitor.Variables()
				undeclaredLocalVariablesInStatement := util.NewSetWithSlice[intsrv.ILocalVariable](searchUndeclaredLocalVariableVisitor.Variables())
				undeclaredLocalVariablesInStatement.RetainAll(undeclaredLocalVariables.ToSlice())

				if !undeclaredLocalVariablesInStatement.IsEmpty() {
					index1 := iterator.NextIndex()

					if state.IsExpressionStatement() {
						f.createInlineDeclarations2(undeclaredLocalVariables,
							undeclaredLocalVariablesInStatement, iterator, state.(intmod.IExpressionStatement))
					} else if state.IsForStatement() {
						f.createInlineDeclarations3(undeclaredLocalVariables,
							undeclaredLocalVariablesInStatement, state.(intsrv.IClassFileForStatement))
					}

					if !undeclaredLocalVariablesInStatement.IsEmpty() {
						// Set the cursor before current state
						index2 := iterator.NextIndex() + undeclaredLocalVariablesInStatement.Size()

						for iterator.NextIndex() >= index1 {
							iterator.Previous()
						}

						sorted := make([]intsrv.ILocalVariable, 0)
						sorted = append(sorted, undeclaredLocalVariablesInStatement.ToSlice()...)
						sort.SliceIsSorted(sorted, func(i, j int) bool {
							return sorted[i].Index() > sorted[j].Index()
						})

						for _, lv := range sorted {
							_ = iterator.Add(statement.NewLocalVariableDeclarationStatement(lv.Type(),
								srvdecl.NewClassFileLocalVariableDeclarator(lv)))
							lv.SetDeclared(true)
							undeclaredLocalVariables.Remove(lv)
						}

						for iterator.NextIndex() < index2 {
							iterator.Next()
						}
					}
				}

				if undeclaredLocalVariables.IsEmpty() {
					break
				}
			}
		}
	}
}

func (f *Frame) createMapForInlineDeclarations() map[intsrv.IFrame][]intsrv.ILocalVariable {
	mapped := make(map[intsrv.IFrame][]intsrv.ILocalVariable)
	i := len(f.localVariableArray)

	for ; i > 0; i-- {
		lv := f.localVariableArray[i]
		for lv != nil {
			if lv.Frame() == f && !lv.IsDeclared() {
				variablesToDeclare := mapped[f]
				if variablesToDeclare == nil {
					variablesToDeclare = make([]intsrv.ILocalVariable, 0)
					mapped[f] = variablesToDeclare
				}
				variablesToDeclare = append(variablesToDeclare, lv)
			}
			lv = lv.Next()
		}
	}

	return mapped
}

func (f *Frame) createInlineDeclarations2(
	undeclaredLocalVariables util.ISet[intsrv.ILocalVariable],
	undeclaredLocalVariablesInStatement util.ISet[intsrv.ILocalVariable],
	iterator util.IListIterator[intmod.IStatement],
	es intmod.IExpressionStatement) {

	if es.Expression().IsBinaryOperatorExpression() {
		boe := es.Expression()

		if boe.Operator() == "=" {
			expressions := expression.NewExpressions()

			f.splitMultiAssignment(math.MaxInt, undeclaredLocalVariablesInStatement, expressions, boe)
			_ = iterator.Remove()

			for _, exp := range expressions.ToSlice() {
				_ = iterator.Add(f.newDeclarationStatement(undeclaredLocalVariables, undeclaredLocalVariablesInStatement, exp))
			}

			if expressions.IsEmpty() {
				_ = iterator.Add(es)
			}
		}
	}
}

func (f *Frame) splitMultiAssignment(toOffset int,
	undeclaredLocalVariablesInStatement util.ISet[intsrv.ILocalVariable],
	expressions intmod.IExpressions, expr intmod.IExpression) intmod.IExpression {

	if expr.IsBinaryOperatorExpression() && expr.Operator() == "=" {
		rightExpression := f.splitMultiAssignment(toOffset, undeclaredLocalVariablesInStatement, expressions, expr.RightExpression())

		if expr.LeftExpression().IsLocalVariableReferenceExpression() {
			lvre := expr.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
			localVariable := lvre.LocalVariable().(intsrv.ILocalVariable)

			if undeclaredLocalVariablesInStatement.Contains(localVariable) && (localVariable.ToOffset() <= toOffset) {
				// Split multi assignment
				if rightExpression == expr.RightExpression() {
					expressions.Add(expr)
				} else {
					expressions.Add(expression.NewBinaryOperatorExpression(
						expr.LineNumber(), expr.Type(), lvre, "=", rightExpression, expr.Priority()))
				}
				// Return local variable
				return lvre
			}
		}
	}

	return expr
}

func (f *Frame) newDeclarationStatement(undeclaredLocalVariables util.ISet[intsrv.ILocalVariable],
	undeclaredLocalVariablesInStatement util.ISet[intsrv.ILocalVariable], boe intmod.IExpression) intmod.ILocalVariableDeclarationStatement {
	reference := boe.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
	localVariable := reference.LocalVariable().(intsrv.ILocalVariable)

	undeclaredLocalVariables.Remove(localVariable)
	undeclaredLocalVariablesInStatement.Remove(localVariable)
	localVariable.SetDeclared(true)

	typ := localVariable.Type()
	var variableInitializer intmod.IVariableInitializer

	if boe.RightExpression().IsNewInitializedArray() {
		if typ.IsObjectType() && typ.(intmod.IObjectType).TypeArguments() != nil {
			variableInitializer = declaration.NewExpressionVariableInitializer(boe.RightExpression())
		} else {
			variableInitializer = boe.RightExpression().(intmod.INewInitializedArray).ArrayInitializer()
		}
	} else {
		variableInitializer = declaration.NewExpressionVariableInitializer(boe.RightExpression())
	}

	return statement.NewLocalVariableDeclarationStatement(typ,
		srvdecl.NewClassFileLocalVariableDeclarator2(boe.LineNumber(),
			reference.LocalVariable().(intsrv.ILocalVariable), variableInitializer))
}

func (f *Frame) createInlineDeclarations3(undeclaredLocalVariables util.ISet[intsrv.ILocalVariable],
	undeclaredLocalVariablesInStatement util.ISet[intsrv.ILocalVariable], fs intsrv.IClassFileForStatement) {
	init := fs.Init()

	if init != nil {
		expressions := expression.NewExpressions()
		toOffset := fs.ToOffset()

		if init.IsList() {
			for _, exp := range init.ToSlice() {
				f.splitMultiAssignment(toOffset, undeclaredLocalVariablesInStatement, expressions, exp)
				if expressions.IsEmpty() {
					expressions.Add(exp)
				}
			}
		} else {
			f.splitMultiAssignment(toOffset, undeclaredLocalVariablesInStatement, expressions, init.First())
			if expressions.IsEmpty() {
				expressions.Add(init.First())
			}
		}

		if expressions.Size() == 1 {
			f.updateForStatement(undeclaredLocalVariables, undeclaredLocalVariablesInStatement, fs, expressions.First())
		} else {
			f.updateForStatement(undeclaredLocalVariables, undeclaredLocalVariablesInStatement, fs, expressions)
		}
	}
}

func (f *Frame) updateForStatement(
	undeclaredLocalVariables util.ISet[intsrv.ILocalVariable],
	undeclaredLocalVariablesInStatement util.ISet[intsrv.ILocalVariable],
	forStatement intsrv.IClassFileForStatement, init intmod.IExpression) {

	if !init.IsBinaryOperatorExpression() {
		return
	}

	if !init.LeftExpression().IsLocalVariableReferenceExpression() {
		return
	}

	reference := init.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
	localVariable := reference.LocalVariable().(intsrv.ILocalVariable)

	if localVariable.IsDeclared() || (localVariable.ToOffset() > forStatement.ToOffset()) {
		return
	}

	undeclaredLocalVariables.Remove(localVariable)
	undeclaredLocalVariablesInStatement.Remove(localVariable)
	localVariable.SetDeclared(true)

	var variableInitializer intmod.IVariableInitializer

	if init.RightExpression().IsNewInitializedArray() {
		variableInitializer = init.RightExpression().(intmod.INewInitializedArray).ArrayInitializer()
	} else {
		variableInitializer = declaration.NewExpressionVariableInitializer(init.RightExpression())
	}

	forStatement.SetDeclaration(declaration.NewLocalVariableDeclaration(localVariable.Type(),
		srvdecl.NewClassFileLocalVariableDeclarator2(init.LineNumber(),
			reference.LocalVariable().(intsrv.ILocalVariable), variableInitializer)))
	forStatement.SetInit(nil)
}

func (f *Frame) updateForStatement2(
	variablesToDeclare util.ISet[intsrv.ILocalVariable], foundVariables util.ISet[intsrv.ILocalVariable],
	forStatement intsrv.IClassFileForStatement, init intmod.IExpressions) {
	boes := util.NewDefaultList[intmod.IExpression]()
	localVariables := util.NewDefaultList[intsrv.ILocalVariable]()
	var type0 intmod.IType
	var type1 intmod.IType
	minDimension := 0
	maxDimension := 0

	for _, expr := range init.ToSlice() {
		if !expr.IsBinaryOperatorExpression() {
			return
		}
		if !expr.LeftExpression().IsLocalVariableReferenceExpression() {
			return
		}

		localVariable := expr.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression).
			LocalVariable().(intsrv.ILocalVariable)

		if localVariable.IsDeclared() || (localVariable.ToOffset() > forStatement.ToOffset()) {
			return
		}

		if type1 == nil {
			type1 = localVariable.Type()
			type0 = type1.CreateType(0)
			minDimension = type1.Dimension()
			maxDimension = type1.Dimension()
		} else {
			type2 := localVariable.Type()

			if type1.IsPrimitiveType() && type2.IsPrimitiveType() {
				typ := GetCommonPrimitiveType(type1.(intmod.IPrimitiveType), type2.(intmod.IPrimitiveType))

				if typ == nil {
					return
				}

				type0 = typ
				type1 = typ.CreateType(type1.Dimension())
				type2 = typ.CreateType(type2.Dimension())
			} else if !(type1 == type2) && !(type0 == type2.CreateType(0)) {
				return
			}

			dimension := type2.Dimension()

			if minDimension > dimension {
				minDimension = dimension
			}
			if maxDimension < dimension {
				maxDimension = dimension
			}
		}

		localVariables.Add(localVariable)
		boes.Add(expr)
	}

	for _, lv := range localVariables.ToSlice() {
		variablesToDeclare.Remove(lv)
		foundVariables.Remove(lv)
		lv.SetDeclared(true)
	}

	if minDimension == maxDimension {
		forStatement.SetDeclaration(declaration.NewLocalVariableDeclaration(type1, f.createDeclarators1(boes, false)))
	} else {
		forStatement.SetDeclaration(declaration.NewLocalVariableDeclaration(type0, f.createDeclarators1(boes, true)))
	}

	forStatement.SetInit(nil)
}

func (f *Frame) createDeclarators1(boes util.IList[intmod.IExpression], setDimension bool) intmod.ILocalVariableDeclarators {
	declarators := declaration.NewLocalVariableDeclarators()

	for _, boe := range boes.ToSlice() {
		reference := boe.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
		var variableInitializer intmod.IVariableInitializer
		if boe.RightExpression().IsNewInitializedArray() {
			variableInitializer = boe.RightExpression().(intmod.INewInitializedArray).ArrayInitializer()
		} else {
			variableInitializer = declaration.NewExpressionVariableInitializer(boe.RightExpression())
		}
		declarator := srvdecl.NewClassFileLocalVariableDeclarator2(boe.LineNumber(),
			reference.LocalVariable().(intsrv.ILocalVariable), variableInitializer)

		if setDimension {
			declarator.SetDimension(reference.LocalVariable().(intsrv.ILocalVariable).Dimension())
		}

		declarators.Add(declarator)
	}

	return declarators
}

func (f *Frame) createStartBlockDeclarations() {
	addIndex := -1
	i := len(f.localVariableArray)

	for ; i > 0; i-- {
		lv := f.localVariableArray[i]
		for lv != nil {
			if lv.IsDeclared() {
				if addIndex == -1 {
					addIndex = f.AddIndex()
				}

				f.stat.AddAt(addIndex, statement.NewLocalVariableDeclarationStatement(
					lv.Type(), srvdecl.NewClassFileLocalVariableDeclarator(lv)))
				lv.SetDeclared(true)
			}

			lv = lv.Next()
		}
	}
}

func (f *Frame) AddIndex() int {
	addIndex := 0

	if f.parent.Parent() == nil {
		// Insert declarations after 'super' call invocation => Search index of SuperConstructorInvocationExpression.
		length := f.stat.Size()

		for addIndex < length {
			state := f.stat.Get(addIndex)
			addIndex++
			if state.IsExpressionStatement() {
				expr := state.Expression()
				if expr.IsSuperConstructorInvocationExpression() || expr.IsConstructorInvocationExpression() {
					break
				}
			}
		}

		if addIndex >= length {
			addIndex = 0
		}
	}

	return addIndex
}

func (f *Frame) mergeDeclarations() {
	size := f.stat.Size()

	if size > 1 {
		declarations := util.NewDefaultList[intmod.ILocalVariableDeclarationStatement]()
		iterator := f.stat.ListIterator()

		for iterator.HasNext() {
			previous := iterator.Next()

			for !previous.IsLocalVariableDeclarationStatement() && iterator.HasNext() {
				previous = iterator.Next()
			}

			if previous.IsLocalVariableDeclarationStatement() {
				lvds1 := previous.(intmod.ILocalVariableDeclarationStatement)
				type1 := lvds1.Type()
				type0 := type1.CreateType(0)
				minDimension := type1.Dimension()
				maxDimension := minDimension
				lineNumber1 := lvds1.LocalVariableDeclarators().LineNumber()

				declarations.Clear()
				declarations.Add(lvds1)

				for iterator.HasNext() {
					stat := iterator.Next()

					if !stat.IsLocalVariableDeclarationStatement() {
						iterator.Previous()
						break
					}

					lvds2 := stat.(intmod.ILocalVariableDeclarationStatement)
					lineNumber2 := lvds2.LocalVariableDeclarators().LineNumber()

					if lineNumber1 != lineNumber2 {
						iterator.Previous()
						break
					}

					lineNumber1 = lineNumber2
					type2 := lvds2.Type()

					if type1.IsPrimitiveType() && type2.IsPrimitiveType() {
						t := GetCommonPrimitiveType(type1.(intmod.IPrimitiveType), type2.(intmod.IPrimitiveType))

						if t == nil {
							iterator.Previous()
							break
						}

						type0 = t
						type1 = t.CreateType(type1.Dimension())
						type2 = t.CreateType(type2.Dimension())
					} else if type1 != type2 && type0 != type2.CreateType(0) {
						iterator.Previous()
						break
					}

					dimension := type2.Dimension()

					if minDimension > dimension {
						minDimension = dimension
					}
					if maxDimension < dimension {
						maxDimension = dimension
					}

					declarations.Add(lvds2)
				}

				declarationSize := declarations.Size()

				if declarationSize > 1 {
					for declarationSize--; declarationSize > 0; {
						iterator.Previous()
						_ = iterator.Remove()
					}

					iterator.Previous()

					if minDimension == maxDimension {
						_ = iterator.Set(statement.NewLocalVariableDeclarationStatement(type1, f.createDeclarators2(declarations, false)))
					} else {
						_ = iterator.Set(statement.NewLocalVariableDeclarationStatement(type0, f.createDeclarators2(declarations, true)))
					}

					iterator.Next()
				}
			}
		}
	}
}

func (f *Frame) createDeclarators2(declarations util.IList[intmod.ILocalVariableDeclarationStatement],
	setDimension bool) intmod.ILocalVariableDeclarators {
	declarators := declaration.NewLocalVariableDeclarators()

	for _, decl := range declarations.ToSlice() {
		declarator := decl.LocalVariableDeclarators().(intmod.ILocalVariableDeclarator)

		if setDimension {
			declarator.SetDimension(decl.Type().Dimension())
		}

		declarators.Add(declarator)
	}

	return declarators
}

func NewGenerateLocalVariableNameVisitor(blackListNames []string, types map[intmod.IType]bool) *GenerateLocalVariableNameVisitor {
	return &GenerateLocalVariableNameVisitor{
		blackListNames: blackListNames,
		types:          types,
	}
}

var IntegerNames = []string{"i", "j", "k", "m", "n"}

type GenerateLocalVariableNameVisitor struct {
	sb             string
	blackListNames []string
	types          map[intmod.IType]bool
	name           string
}

func (c *GenerateLocalVariableNameVisitor) Name() string {
	return c.name
}

func (c *GenerateLocalVariableNameVisitor) capitalize(str string) {
	if str != "" {
		length := len(str)
		if length > 0 {
			firstChar := str[0]

			if unicode.IsUpper(rune(firstChar)) {
				c.sb += str
			} else {
				c.sb += strings.ToUpper(string(firstChar))
				if length > 1 {
					c.sb += str[1:]
				}
			}
		}
	}
}

func (c *GenerateLocalVariableNameVisitor) uncapitalize(str string) {
	if str != "" {
		length := len(str)
		if length > 0 {
			firstChar := str[0]
			if unicode.IsLower(rune(firstChar)) {
				c.sb += str
			} else {
				c.sb += strings.ToLower(string(firstChar))
				if length > 1 {
					c.sb += str[1:]
				}
			}
		}
	}
}

func (c *GenerateLocalVariableNameVisitor) generate(typ intmod.IType) {
	length := len(c.sb)
	counter := 1

	if c.types[typ] {
		c.sb += fmt.Sprintf("%d", counter)
		counter++
	}

	c.name = c.sb

	for contains(c.blackListNames, c.name) {
		c.sb = c.sb[:length]
		c.sb += fmt.Sprintf("%d", counter)
		counter++
		c.name = c.sb
	}

	c.blackListNames = append(c.blackListNames, c.name)
}

func (c *GenerateLocalVariableNameVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) {
	c.sb = ""

	switch t.Dimension() {
	case intmod.FlagByte:
		c.sb += "b"
	case intmod.FlagChar:
		c.sb += "c"
	case intmod.FlagDouble:
		c.sb += "d"
	case intmod.FlagFloat:
		c.sb += "f"
	case intmod.FlagInt:
		for _, in := range IntegerNames {
			if !contains(c.blackListNames, in) {
				c.blackListNames = append(c.blackListNames, in)
				return
			}
		}
		c.sb += "i"
	case intmod.FlagLong:
		c.sb += "l"
	case intmod.FlagShort:
		c.sb += "s"
	case intmod.FlagBoolean:
		c.sb += "bool"
	}

	c.generate(t)
}

func (c *GenerateLocalVariableNameVisitor) Visit(t intmod.IType, str string) {
	c.sb = ""

	switch t.Dimension() {
	case 0:
		if str == "Class" {
			c.sb += "clazz"
		} else if str == "String" {
			c.sb += "str"
		} else if str == "Boolean" {
			c.sb += "bool"
		} else {
			c.uncapitalize(str)
			if contains(CapitalizedJavaLanguageKeywords, str) {
				c.sb += "_"
			}
		}
	default:
		c.sb += "arrayOf"
		c.capitalize(str)
	}

	c.generate(t)
}

func (c *GenerateLocalVariableNameVisitor) VisitObjectType(t intmod.IObjectType) {
	c.Visit(t.(intmod.IType), t.Name())
}

func (c *GenerateLocalVariableNameVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	c.Visit(t.(intmod.IType), t.Name())
}

func (c *GenerateLocalVariableNameVisitor) VisitGenericType(t intmod.IGenericType) {
	c.Visit(t.(intmod.IType), t.Name())
}

func (c *GenerateLocalVariableNameVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {}

func (c *GenerateLocalVariableNameVisitor) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {
}

func (c *GenerateLocalVariableNameVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
}

func (c *GenerateLocalVariableNameVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
}

func (c *GenerateLocalVariableNameVisitor) VisitWildcardTypeArgument(argument intmod.IWildcardTypeArgument) {
}

func NewAbstractLocalVariableComparator() *AbstractLocalVariableComparator {
	return &AbstractLocalVariableComparator{}
}

type AbstractLocalVariableComparator struct {
}

func (c *AbstractLocalVariableComparator) Compare(alv1, alv2 intsrv.ILocalVariable) int {
	return alv1.Index() - alv2.Index()
}

func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func retainAll(src, target []intsrv.ILocalVariable) []intsrv.ILocalVariable {
	// target의 값을 Set으로 저장
	targetSet := make(map[intsrv.ILocalVariable]struct{})
	for _, v := range target {
		targetSet[v] = struct{}{}
	}

	// src 슬라이스에서 target에 포함된 값만 남기기
	result := make([]intsrv.ILocalVariable, 0)
	for _, v := range src {
		if _, exists := targetSet[v]; exists {
			result = append(result, v)
		}
	}

	return result
}
