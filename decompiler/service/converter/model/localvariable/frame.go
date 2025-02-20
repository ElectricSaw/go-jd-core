package localvariable

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	srvdecl "github.com/ElectricSaw/go-jd-core/decompiler/service/converter/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/converter/visitor"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"math"
	"sort"
	"strings"
	"unicode"
)

var CapitalizedJavaLanguageKeywords = []string{
	"Abstract", "Continue", "For", "New", "Switch", "Assert", "Default", "Goto", "Package", "Synchronized",
	"Boolean", "Do", "If", "Private", "This", "Break", "Double", "Implements", "Protected", "Throw", "Byte", "Else",
	"Import", "Public", "Throws", "Case", "Enum", "Instanceof", "Return", "Transient", "Catch", "Extends", "Int",
	"Short", "Try", "Char", "Final", "Interface", "Static", "Void", "Class", "Finally", "Long", "Strictfp",
	"Volatile", "Const", "Float", "Native", "Super", "While"}

func NewFrame(parent IFrame, stat model.Statements) Frame {
	return Frame{
		LocalVariableArray: make([]ILocalVariable, 10),
		NewExpressions:     make(map[*model.NewExpression]ILocalVariable),
		Children:           util.NewDefaultList[IFrame](),
		Parent:             parent,
		Statements:         stat,
	}
}

type Frame struct {
	LocalVariableArray     []ILocalVariable
	NewExpressions         map[*model.NewExpression]ILocalVariable
	Children               util.IList[IFrame]
	Parent                 IFrame
	Statements             model.Statements
	ExceptionLocalVariable ILocalVariable
}

func (f *Frame) GetStatements() model.Statements {
	return f.Statements
}

func (f *Frame) GetLocalVariable(index int) ILocalVariableReference {
	if index < len(f.LocalVariableArray) {
		lv := f.LocalVariableArray[index]
		if lv != nil {
			return lv
		}
	}
	return f.Parent.GetLocalVariable(index)
}

func (f *Frame) GetParent() IFrame {
	return f.Parent
}

func (f *Frame) AddLocalVariable(lv ILocalVariable) {
	// Java의 assert 대체 코드
	if lv.Next() != nil {
		fmt.Println("Frame.AddLocalVariable: add local variable failed")
		return
	}

	index := lv.Index()

	// 배열 크기 늘리기
	if index >= len(f.LocalVariableArray) {
		newArray := make([]ILocalVariable, index*2)
		copy(newArray, f.LocalVariableArray)
		f.LocalVariableArray = newArray
	}

	next := f.LocalVariableArray[index]

	// 중복 추가 방지
	if next != lv {
		f.LocalVariableArray[index] = lv
		lv.SetNext(next)
		lv.SetFrame(f)
	}
}

func (f *Frame) MergeLocalVariable(typeBounds map[string]model.IType, localVariableMaker *visitor.LocalVariableMaker, lv ILocalVariable) {
	index := lv.Index()
	var alvToMerge ILocalVariable

	if index < len(f.LocalVariableArray) {
		alvToMerge = f.LocalVariableArray[index]
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
		if f.Children != nil {
			for _, child := range f.Children.ToSlice() {
				child.MergeLocalVariable(typeBounds, localVariableMaker, lv)
			}
		}
	} else if lv != alvToMerge {
		for _, reference := range alvToMerge.References().ToSlice() {
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
				t := GetCommonPrimitiveType(plv.Type().(*model.PrimitiveType), plvToMerype.Type().(*model.PrimitiveType))

				if t == nil {
					t = &model.PtTypeInt
				}

				plv.SetType(t.CreateType(typ.GetDimension()).(*model.PrimitiveType))
			}
		} else {
			if typ.IsPrimitiveType() {
				plv := lv.(*PrimitiveLocalVariable)

				if alvToMerge.IsAssignableFromWithVariable(typeBounds, lv) || localVariableMaker.IsCompatible(alvToMerge, lv.Type()) {
					plv.SetType(alvToMerype.(*model.PrimitiveType))
				} else {
					plv.SetType(&model.PtTypeInt)
				}
			} else if typ.IsObjectType() {
				olv := lv.(*ObjectLocalVariable)

				if alvToMerge.IsAssignableFromWithVariable(typeBounds, lv) || localVariableMaker.IsCompatible(alvToMerge, lv.Type()) {
					olv.SetTypeWithTypeBounds(typeBounds, alvToMerype)
				} else {
					dimension := alvToMerge.Dimension()
					if lv.Dimension() >= alvToMerge.Dimension() {
						dimension = lv.Dimension()
					}
					olv.SetTypeWithTypeBounds(typeBounds, model.OtTypeObject.CreateType(dimension))
				}
			}
		}

		f.LocalVariableArray[index] = alvToMerge.Next()
	}
}

func (f *Frame) RemoveLocalVariable(lv ILocalVariable) {
	index := lv.Index()
	var alvToRemove ILocalVariable

	if (index < len(f.LocalVariableArray)) && (f.LocalVariableArray[index] == lv) {
		alvToRemove = lv
	} else {
		alvToRemove = nil
	}

	if alvToRemove == nil {
		if f.Children != nil {
			for _, child := range f.Children.ToSlice() {
				child.RemoveLocalVariable(lv)
			}
		}
	} else {
		f.LocalVariableArray[index] = alvToRemove.Next()
		alvToRemove.SetNext(nil)
	}
}

func (f *Frame) AddChild(child IFrame) {
	if f.Children == nil {
		f.Children = util.NewDefaultList[IFrame]()
	}
	f.Children.Add(child)
}

func (f *Frame) Close() {
	// Update type for 'new' expression
	if f.NewExpressions != nil {
		for key, value := range f.NewExpressions {
			ot1 := key.GetObjectType()
			ot2 := value.Type().(*model.ObjectType)

			if (ot1.TypeArguments == nil) && (ot2.TypeArguments != nil) {
				key.Type = ot1.CreateTypeWithArgs(ot2.TypeArguments).(*model.ObjectType)
			}
		}
	}
}

func (f *Frame) CreateNames(parentNames []string) {
	names := make([]string, 0, len(parentNames))
	copy(names, parentNames)
	types := make(map[model.IType]bool)
	length := len(f.LocalVariableArray)

	for i := 0; i < length; i++ {
		lv := f.LocalVariableArray[i]

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

	if f.ExceptionLocalVariable != nil {
		if _, ok := types[f.ExceptionLocalVariable.Type()]; ok {
			types[f.ExceptionLocalVariable.Type()] = true // Non unique type
		} else {
			types[f.ExceptionLocalVariable.Type()] = false // Unique type
		}
	}

	if len(types) != 0 {
		visit0r := NewGenerateLocalVariableNameVisitor(names, types)

		for i := 0; i < length; i++ {
			lv := f.LocalVariableArray[i]
			for lv != nil {
				if lv.Name() == "" {
					lv.Type().(model.ITypeArgumentVisitable).AcceptTypeArgumentVisitor(visit0r)
					lv.SetName(visit0r.Name())
				}
				lv = lv.Next()
			}
		}

		if f.ExceptionLocalVariable != nil {
			f.ExceptionLocalVariable.Type().(model.ITypeArgumentVisitable).AcceptTypeArgumentVisitor(visit0r)
			f.ExceptionLocalVariable.SetName(visit0r.Name())
		}
	}

	// Recursive call
	if f.Children != nil {
		for _, child := range f.Children.ToSlice() {
			child.CreateNames(names)
		}
	}
}

func (f *Frame) UpdateLocalVariableInForStatements(typeMaker *visitor.TypeMaker) {
	// Recursive call first
	if f.Children != nil {
		for _, child := range f.Children.ToSlice() {
			child.UpdateLocalVariableInForStatements(typeMaker)
		}
	}

	// Split local variable ranges in init 'for' statements
	searchLocalVariableVisitor := visitor.NewSearchLocalVariableVisitor()
	undeclaredInExpressionStatements := make([]ILocalVariable, 0)

	for _, stat := range f.Statements.ToSlice() {
		if stat.IsForStatement() {
			if stat.GetInit() == nil {
				if stat.GetCondition() != nil {
					searchLocalVariableVisitor.Init()
					stat.GetCondition().Accept(searchLocalVariableVisitor)
					for _, variable := range searchLocalVariableVisitor.Variables() {
						undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable)
					}
				}
				if stat.GetUpdate() != nil {
					searchLocalVariableVisitor.Init()
					stat.GetUpdate().Accept(searchLocalVariableVisitor)
					for _, variable := range searchLocalVariableVisitor.Variables() {
						undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable)
					}
				}
				if stat.GetStatements() != nil {
					searchLocalVariableVisitor.Init()
					stat.GetStatements().AcceptStatement(searchLocalVariableVisitor)
					for _, variable := range searchLocalVariableVisitor.Variables() {
						undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable)
					}
				}
			}
		} else {
			searchLocalVariableVisitor.Init()
			stat.AcceptStatement(searchLocalVariableVisitor)
			for _, variable := range searchLocalVariableVisitor.Variables() {
				undeclaredInExpressionStatements = append(undeclaredInExpressionStatements, variable)
			}
		}
	}

	searchUndeclaredLocalVariableVisitor := visitor.NewSearchUndeclaredLocalVariableVisitor()
	undeclaredInForStatements := make(map[ILocalVariable][]*ClassFileForStatement)

	for _, stat := range f.Statements.ToSlice() {
		if stat.IsForStatement() {
			fs := stat.(*ClassFileForStatement)

			if fs.GetInit() != nil {
				searchUndeclaredLocalVariableVisitor.Init()
				fs.GetInit().Accept(searchUndeclaredLocalVariableVisitor)
				searchUndeclaredLocalVariableVisitor.RemoveAll(undeclaredInExpressionStatements)

				for _, lv := range searchUndeclaredLocalVariableVisitor.Variables() {
					list := undeclaredInForStatements[lv]
					if list == nil {
						list = make([]*ClassFileForStatement, 0)
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

				if lv.References().Size() == 0 {
					lv.Frame().RemoveLocalVariable(lv)
				}
			}
		}
	}
}

func (f *Frame) createNewLocalVariable(createLocalVariableVisitor *CreateLocalVariableVisitor,
	fs *ClassFileForStatement, lv ILocalVariable) {
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
	if f.Children != nil {
		for _, child := range f.Children.ToSlice() {
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
			undeclaredLocalVariables := util.NewSetWithSlice[ILocalVariable](value)

			for iterator.HasNext() {
				state := iterator.Next()
				searchUndeclaredLocalVariableVisitor.Init()
				state.AcceptStatement(searchUndeclaredLocalVariableVisitor)

				//undeclaredLocalVariablesInStatement := searchUndeclaredLocalVariableVisitor.Variables()
				undeclaredLocalVariablesInStatement := util.NewSetWithSlice[ILocalVariable](searchUndeclaredLocalVariableVisitor.Variables())
				undeclaredLocalVariablesInStatement.RetainAll(undeclaredLocalVariables.ToSlice())

				if !undeclaredLocalVariablesInStatement.IsEmpty() {
					index1 := iterator.NextIndex()

					if state.IsExpressionStatement() {
						f.createInlineDeclarations2(undeclaredLocalVariables,
							undeclaredLocalVariablesInStatement, iterator, state.(model.IExpressionStatement))
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

						sorted := make([]ILocalVariable, 0)
						sorted = append(sorted, undeclaredLocalVariablesInStatement.ToSlice()...)
						sort.SliceIsSorted(sorted, func(i, j int) bool {
							return sorted[i].Index() > sorted[j].Index()
						})

						for _, lv := range sorted {
							_ = iterator.Add(model.NewLocalVariableDeclarationStatement(lv.Type(),
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

func (f *Frame) createMapForInlineDeclarations() map[IFrame][]ILocalVariable {
	mapped := make(map[IFrame][]ILocalVariable)
	i := len(f.LocalVariableArray)

	for i > 0 {
		i--
		lv := f.LocalVariableArray[i]
		for lv != nil {
			if lv.Frame() == f && !lv.IsDeclared() {
				variablesToDeclare := mapped[f]
				if variablesToDeclare == nil {
					variablesToDeclare = make([]ILocalVariable, 0)
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
	undeclaredLocalVariables util.ISet[ILocalVariable],
	undeclaredLocalVariablesInStatement util.ISet[ILocalVariable],
	iterator util.IListIterator[model.IStatement],
	es model.ExpressionStatement) {

	if es.Expression.IsBinaryOperatorExpression() {
		boe := es.Expression

		if boe.GetOperator() == "=" {
			expressions := model.NewExpressions()

			f.splitMultiAssignment(math.MaxInt, undeclaredLocalVariablesInStatement, expressions, boe)
			_ = iterator.Remove()

			for _, exp := range expressions.ToSlice() {
				_ = iterator.Add(f.newDeclarationStatement(undeclaredLocalVariables, undeclaredLocalVariablesInStatement, exp))
			}

			if expressions.IsEmpty() {
				_ = iterator.Add(&es)
			}
		}
	}
}

func (f *Frame) splitMultiAssignment(toOffset int,
	undeclaredLocalVariablesInStatement util.ISet[ILocalVariable],
	expressions model.Expressions, expr model.IExpression) model.IExpression {

	if expr.IsBinaryOperatorExpression() && expr.GetOperator() == "=" {
		rightExpression := f.splitMultiAssignment(toOffset, undeclaredLocalVariablesInStatement, expressions, expr.GetRightExpression())

		if expr.GetLeftExpression().IsLocalVariableReferenceExpression() {
			lvre := expr.GetLeftExpression().(*ClassFileLocalVariableReferenceExpression)
			localVariable := lvre.LocalVariable().(ILocalVariable)

			if undeclaredLocalVariablesInStatement.Contains(localVariable) && (localVariable.ToOffset() <= toOffset) {
				// Split multi assignment
				if rightExpression == expr.GetRightExpression() {
					expressions.Add(expr)
				} else {
					expressions.Add(model.NewBinaryOperatorExpression(
						expr.GetLineNumber(), expr.GetType(), lvre, "=", rightExpression, expr.GetPriority()))
				}
				// Return local variable
				return lvre
			}
		}
	}

	return expr
}

func (f *Frame) newDeclarationStatement(undeclaredLocalVariables util.ISet[ILocalVariable],
	undeclaredLocalVariablesInStatement util.ISet[ILocalVariable], boe model.IExpression) model.ILocalVariableDeclarationStatement {
	reference := boe.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
	localVariable := reference.LocalVariable().(ILocalVariable)

	undeclaredLocalVariables.Remove(localVariable)
	undeclaredLocalVariablesInStatement.Remove(localVariable)
	localVariable.SetDeclared(true)

	typ := localVariable.Type()
	var variableInitializer model.IVariableInitializer

	if boe.RightExpression().IsNewInitializedArray() {
		if typ.IsObjectType() && typ.(*model.ObjectType).TypeArguments() != nil {
			variableInitializer = model.NewExpressionVariableInitializer(boe.RightExpression())
		} else {
			variableInitializer = boe.RightExpression().(model.INewInitializedArray).ArrayInitializer()
		}
	} else {
		variableInitializer = model.NewExpressionVariableInitializer(boe.RightExpression())
	}

	return model.NewLocalVariableDeclarationStatement(typ,
		srvdecl.NewClassFileLocalVariableDeclarator2(boe.LineNumber(),
			reference.LocalVariable().(ILocalVariable), variableInitializer))
}

func (f *Frame) createInlineDeclarations3(undeclaredLocalVariables util.ISet[ILocalVariable],
	undeclaredLocalVariablesInStatement util.ISet[ILocalVariable], fs intsrv.IClassFileForStatement) {
	init := fs.Init()

	if init != nil {
		expressions := model.NewExpressions()
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
	undeclaredLocalVariables util.ISet[ILocalVariable],
	undeclaredLocalVariablesInStatement util.ISet[ILocalVariable],
	forStatement intsrv.IClassFileForStatement, init model.IExpression) {

	if !init.IsBinaryOperatorExpression() {
		return
	}

	if !init.LeftExpression().IsLocalVariableReferenceExpression() {
		return
	}

	reference := init.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
	localVariable := reference.LocalVariable().(ILocalVariable)

	if localVariable.IsDeclared() || (localVariable.ToOffset() > forStatement.ToOffset()) {
		return
	}

	undeclaredLocalVariables.Remove(localVariable)
	undeclaredLocalVariablesInStatement.Remove(localVariable)
	localVariable.SetDeclared(true)

	var variableInitializer model.IVariableInitializer

	if init.RightExpression().IsNewInitializedArray() {
		variableInitializer = init.RightExpression().(model.INewInitializedArray).ArrayInitializer()
	} else {
		variableInitializer = model.NewExpressionVariableInitializer(init.RightExpression())
	}

	forStatement.SetDeclaration(model.NewLocalVariableDeclaration(localVariable.Type(),
		srvdecl.NewClassFileLocalVariableDeclarator2(init.LineNumber(),
			reference.LocalVariable().(ILocalVariable), variableInitializer)))
	forStatement.SetInit(nil)
}

func (f *Frame) updateForStatement2(
	variablesToDeclare util.ISet[ILocalVariable], foundVariables util.ISet[ILocalVariable],
	forStatement intsrv.IClassFileForStatement, init model.IExpressions) {
	boes := util.NewDefaultList[model.IExpression]()
	localVariables := util.NewDefaultList[ILocalVariable]()
	var type0 model.IType
	var type1 model.IType
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
			LocalVariable().(ILocalVariable)

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
				typ := GetCommonPrimitiveType(type1.(*model.PrimitiveType), type2.(*model.PrimitiveType))

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
		forStatement.SetDeclaration(model.NewLocalVariableDeclaration(type1, f.createDeclarators1(boes, false)))
	} else {
		forStatement.SetDeclaration(model.NewLocalVariableDeclaration(type0, f.createDeclarators1(boes, true)))
	}

	forStatement.SetInit(nil)
}

func (f *Frame) createDeclarators1(boes util.IList[model.IExpression], setDimension bool) model.ILocalVariableDeclarators {
	declarators := model.NewLocalVariableDeclarators()

	for _, boe := range boes.ToSlice() {
		reference := boe.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
		var variableInitializer model.IVariableInitializer
		if boe.RightExpression().IsNewInitializedArray() {
			variableInitializer = boe.RightExpression().(model.INewInitializedArray).ArrayInitializer()
		} else {
			variableInitializer = model.NewExpressionVariableInitializer(boe.RightExpression())
		}
		declarator := srvdecl.NewClassFileLocalVariableDeclarator2(boe.LineNumber(),
			reference.LocalVariable().(ILocalVariable), variableInitializer)

		if setDimension {
			declarator.SetDimension(reference.LocalVariable().(ILocalVariable).Dimension())
		}

		declarators.Add(declarator)
	}

	return declarators
}

func (f *Frame) createStartBlockDeclarations() {
	addIndex := -1
	i := len(f.LocalVariableArray)

	for i > 0 {
		i--
		lv := f.LocalVariableArray[i]
		for lv != nil {
			if lv.IsDeclared() {
				if addIndex == -1 {
					addIndex = f.AddIndex()
				}

				_ = f.Statements.AddAt(addIndex, model.NewLocalVariableDeclarationStatement(
					lv.Type(), srvdecl.NewClassFileLocalVariableDeclarator(lv)))
				lv.SetDeclared(true)
			}

			lv = lv.Next()
		}
	}
}

func (f *Frame) AddIndex() int {
	addIndex := 0

	if f.Parent.Parent() == nil {
		// Insert declarations after 'super' call invocation => Search index of SuperConstructorInvocationExpression.
		length := f.Statements.Size()

		for addIndex < length {
			state := f.Statements.Get(addIndex)
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
	size := f.Statements.Size()

	if size > 1 {
		declarations := util.NewDefaultList[model.ILocalVariableDeclarationStatement]()
		iterator := f.Statements.ListIterator()

		for iterator.HasNext() {
			previous := iterator.Next()

			for !previous.IsLocalVariableDeclarationStatement() && iterator.HasNext() {
				previous = iterator.Next()
			}

			if previous.IsLocalVariableDeclarationStatement() {
				lvds1 := previous.(model.ILocalVariableDeclarationStatement)
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

					lvds2 := stat.(model.ILocalVariableDeclarationStatement)
					lineNumber2 := lvds2.LocalVariableDeclarators().LineNumber()

					if lineNumber1 != lineNumber2 {
						iterator.Previous()
						break
					}

					lineNumber1 = lineNumber2
					type2 := lvds2.Type()

					if type1.IsPrimitiveType() && type2.IsPrimitiveType() {
						t := GetCommonPrimitiveType(type1.(*model.PrimitiveType), type2.(*model.PrimitiveType))

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
						_ = iterator.Set(model.NewLocalVariableDeclarationStatement(type1, f.createDeclarators2(declarations, false)))
					} else {
						_ = iterator.Set(model.NewLocalVariableDeclarationStatement(type0, f.createDeclarators2(declarations, true)))
					}

					iterator.Next()
				}
			}
		}
	}
}

func (f *Frame) createDeclarators2(declarations util.IList[model.ILocalVariableDeclarationStatement],
	setDimension bool) model.ILocalVariableDeclarators {
	declarators := model.NewLocalVariableDeclarators()

	for _, decl := range declarations.ToSlice() {
		declarator := decl.LocalVariableDeclarators().(model.ILocalVariableDeclarator)

		if setDimension {
			declarator.SetDimension(decl.Type().Dimension())
		}

		declarators.Add(declarator)
	}

	return declarators
}

func NewGenerateLocalVariableNameVisitor(blackListNames []string, types map[model.IType]bool) *GenerateLocalVariableNameVisitor {
	return &GenerateLocalVariableNameVisitor{
		blackListNames: blackListNames,
		types:          types,
	}
}

var IntegerNames = []string{"i", "j", "k", "m", "n"}

type GenerateLocalVariableNameVisitor struct {
	sb             string
	blackListNames []string
	types          map[model.IType]bool
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

func (c *GenerateLocalVariableNameVisitor) generate(typ model.IType) {
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

func (c *GenerateLocalVariableNameVisitor) VisitPrimitiveType(t *model.PrimitiveType) {
	c.sb = ""

	switch t.Dimension() {
	case model.FlagByte:
		c.sb += "b"
	case model.FlagChar:
		c.sb += "c"
	case model.FlagDouble:
		c.sb += "d"
	case model.FlagFloat:
		c.sb += "f"
	case model.FlagInt:
		for _, in := range IntegerNames {
			if !contains(c.blackListNames, in) {
				c.blackListNames = append(c.blackListNames, in)
				return
			}
		}
		c.sb += "i"
	case model.FlagLong:
		c.sb += "l"
	case model.FlagShort:
		c.sb += "s"
	case model.FlagBoolean:
		c.sb += "bool"
	default:
	}

	c.generate(t)
}

func (c *GenerateLocalVariableNameVisitor) Visit(t model.IType, str string) {
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

func (c *GenerateLocalVariableNameVisitor) VisitObjectType(t *model.ObjectType) {
	c.Visit(t.(model.IType), t.Name())
}

func (c *GenerateLocalVariableNameVisitor) VisitInnerObjectType(t model.IInnerObjectType) {
	c.Visit(t.(model.IType), t.Name())
}

func (c *GenerateLocalVariableNameVisitor) VisitGenericType(t model.IGenericType) {
	c.Visit(t.(model.IType), t.Name())
}

func (c *GenerateLocalVariableNameVisitor) VisitTypeArguments(_ model.ITypeArguments) {}

func (c *GenerateLocalVariableNameVisitor) VisitDiamondTypeArgument(_ model.IDiamondTypeArgument) {
}

func (c *GenerateLocalVariableNameVisitor) VisitWildcardExtendsTypeArgument(_ model.IWildcardExtendsTypeArgument) {
}

func (c *GenerateLocalVariableNameVisitor) VisitWildcardSuperTypeArgument(_ model.IWildcardSuperTypeArgument) {
}

func (c *GenerateLocalVariableNameVisitor) VisitWildcardTypeArgument(_ model.IWildcardTypeArgument) {
}

func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func retainAll(src, target []ILocalVariable) []ILocalVariable {
	// target의 값을 Set으로 저장
	targetSet := make(map[ILocalVariable]struct{})
	for _, v := range target {
		targetSet[v] = struct{}{}
	}

	// src 슬라이스에서 target에 포함된 값만 남기기
	result := make([]ILocalVariable, 0)
	for _, v := range src {
		if _, exists := targetSet[v]; exists {
			result = append(result, v)
		}
	}

	return result
}
