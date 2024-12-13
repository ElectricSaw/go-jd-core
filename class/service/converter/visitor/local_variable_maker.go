package visitor

import (
	"fmt"
	intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	"github.com/ElectricSaw/go-jd-core/class/service/converter/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/class/service/converter/model/localvariable"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewLocalVariableMaker(typeMaker intsrv.ITypeMaker,
	comd intsrv.IClassFileConstructorOrMethodDeclaration,
	constructor bool) intsrv.ILocalVariableMaker {
	classFile := comd.ClassFile()
	method := comd.Method()
	parameterTypes := comd.ParameterTypes()

	m := &LocalVariableMaker{
		localVariableSet:           localvariable.NewLocalVariableSet(),
		names:                      util.NewSet[string](),
		blackListNames:             util.NewSet[string](),
		currentFrame:               localvariable.NewRootFrame(),
		typeMaker:                  typeMaker,
		typeBounds:                 comd.TypeBounds(),
		createParameterVisitor:     NewCreateParameterVisitor(typeMaker),
		createLocalVariableVisitor: NewCreateLocalVariableVisitor(typeMaker),
	}

	m.populateBlackListNamesVisitor = NewPopulateBlackListNamesVisitor(m.blackListNames)
	m.searchInTypeArgumentVisitor = NewSearchInTypeArgumentVisitor()

	if classFile.Fields() != nil {
		for _, field := range classFile.Fields() {
			descriptor := field.Descriptor()
			if descriptor[len(descriptor)-1] == ';' {
				m.typeMaker.MakeFromDescriptor(descriptor).AcceptTypeArgumentVisitor(m.populateBlackListNamesVisitor)
			}
		}
	}

	m.typeMaker.MakeFromInternalTypeName(classFile.InternalTypeName()).AcceptTypeArgumentVisitor(m.populateBlackListNamesVisitor)

	if classFile.SuperTypeName() != "" {
		m.typeMaker.MakeFromInternalTypeName(classFile.SuperTypeName()).AcceptTypeArgumentVisitor(m.populateBlackListNamesVisitor)
	}

	if classFile.InterfaceTypeNames() != nil {
		for _, name := range classFile.InterfaceTypeNames() {
			m.typeMaker.MakeFromInternalTypeName(name).AcceptTypeArgumentVisitor(m.populateBlackListNamesVisitor)
		}
	}

	if parameterTypes != nil {
		if parameterTypes.IsList() {
			for _, t := range parameterTypes.ToSlice() {
				t.AcceptTypeArgumentVisitor(m.populateBlackListNamesVisitor)
			}
		} else {
			parameterTypes.First().AcceptTypeArgumentVisitor(m.populateBlackListNamesVisitor)
		}
	}

	m.initLocalVariablesFromAttributes(method)

	firstVariableIndex := 0

	if method.AccessFlags()&intmod.FlagStatic == 0 {
		if m.localVariableSet.Root(0) == nil {
			m.localVariableSet.Add(0, localvariable.NewObjectLocalVariable(
				m.typeMaker, 0, 0, m.typeMaker.MakeFromInternalTypeName(
					classFile.InternalTypeName()), "this"))
		}
		firstVariableIndex = 1
	}

	if constructor {
		if classFile.IsEnum() {
			if m.localVariableSet.Root(1) == nil {
				m.localVariableSet.Add(1, localvariable.NewObjectLocalVariable(m.typeMaker, 1, 0,
					_type.OtTypeString, "this$enum$name"))
			}
			if m.localVariableSet.Root(2) == nil {
				m.localVariableSet.Add(2, localvariable.NewPrimitiveLocalVariable(2, 0,
					_type.PtTypeInt, "this$enum$index"))
			}
		} else if classFile.OuterClassFile() != nil && !classFile.IsStatic() {
			if m.localVariableSet.Root(1) == nil {
				m.localVariableSet.Add(1, localvariable.NewObjectLocalVariable(m.typeMaker, 1, 0,
					m.typeMaker.MakeFromInternalTypeName(classFile.InternalTypeName()), "this$0"))
			}
		}
	}

	if parameterTypes != nil {
		lastParameterIndex := parameterTypes.Size() - 1
		varargs := method.AccessFlags()&intmod.FlagVarArgs != 0

		m.initLocalVariablesFromParameterTypes(classFile, parameterTypes, varargs, firstVariableIndex, lastParameterIndex)

		rvpa := method.Attribute("RuntimeVisibleParameterAnnotations").(intcls.IAttributeParameterAnnotations)
		ripa := method.Attribute("RuntimeInvisibleParameterAnnotations").(intcls.IAttributeParameterAnnotations)

		if rvpa == nil && ripa == nil {
			variableIndex := firstVariableIndex
			for parameterIndex := 0; parameterIndex <= lastParameterIndex; parameterIndex++ {
				lv := m.localVariableSet.Root(variableIndex)
				m.formalParameters.Add(declaration.NewClassFileFormalParameter2(lv,
					varargs && parameterIndex == lastParameterIndex))

				if _type.PtTypeLong == lv.Type() || _type.PtTypeDouble == lv.Type() {
					variableIndex++
				}

				variableIndex++
			}
		} else {
			var visiblesArray []intcls.IAnnotations = nil
			var invisiblesArray []intcls.IAnnotations = nil
			annotationConverter := NewAnnotationConverter(m.typeMaker)

			if rvpa != nil {
				visiblesArray = rvpa.ParameterAnnotations()
			}
			if ripa != nil {
				invisiblesArray = ripa.ParameterAnnotations()
			}

			variableIndex := firstVariableIndex
			for parameterIndex := 0; parameterIndex <= lastParameterIndex; parameterIndex++ {
				lv := m.localVariableSet.Root(variableIndex)

				var visibles intcls.IAnnotations
				var invisibles intcls.IAnnotations

				if visiblesArray == nil || len(visiblesArray) <= parameterIndex {
					visibles = nil
				} else {
					visibles = visiblesArray[parameterIndex]
				}

				if invisiblesArray == nil || len(invisiblesArray) <= parameterIndex {
					invisibles = nil
				} else {
					invisibles = invisiblesArray[parameterIndex]
				}
				annotationReferences := annotationConverter.ConvertWithAnnotations2(visibles, invisibles)

				m.formalParameters.Add(declaration.NewClassFileFormalParameter3(annotationReferences, lv, varargs && (parameterIndex == lastParameterIndex)))

				if _type.PtTypeLong == lv.Type() || _type.PtTypeDouble == lv.Type() {
					variableIndex++
				}

				variableIndex++
			}
		}
	}

	m.localVariableCache = m.localVariableSet.Initialize(m.currentFrame)

	return m
}

type LocalVariableMaker struct {
	localVariableSet              intsrv.ILocalVariableSet
	names                         util.ISet[string]
	blackListNames                util.ISet[string]
	currentFrame                  intsrv.IRootFrame
	localVariableCache            []intsrv.ILocalVariable
	typeMaker                     intsrv.ITypeMaker
	typeBounds                    map[string]intmod.IType
	formalParameters              intmod.IFormalParameters
	populateBlackListNamesVisitor intsrv.IPopulateBlackListNamesVisitor
	searchInTypeArgumentVisitor   intsrv.ISearchInTypeArgumentVisitor
	createParameterVisitor        intsrv.ICreateParameterVisitor
	createLocalVariableVisitor    intsrv.ICreateLocalVariableVisitor
}

func (m *LocalVariableMaker) initLocalVariablesFromAttributes(method intcls.IMethod) {
	code := method.Attribute("Code").(intcls.IAttributeCode)

	// Init local variables from attributes
	if code != nil {
		localVariableTable := code.Attribute("LocalVariableTable").(intcls.IAttributeLocalVariableTable)

		if localVariableTable != nil {
			staticFlag := (method.AccessFlags() & intmod.FlagStatic) != 0

			for _, localVariable := range localVariableTable.LocalVariableTable() {
				index := localVariable.Index()
				startPc := 0
				if !staticFlag && index == 0 {
					startPc = localVariable.StartPc()
				}
				descriptor := localVariable.Descriptor()
				name := localVariable.Name()
				var lv intsrv.ILocalVariable

				if descriptor[len(descriptor)-1] == ';' {
					lv = localvariable.NewObjectLocalVariable(m.typeMaker, index, startPc,
						m.typeMaker.MakeFromDescriptor(descriptor).(intmod.IType), name)
				} else {
					dimension := CountDimension(descriptor)

					if dimension == 0 {
						lv = localvariable.NewPrimitiveLocalVariable(index, startPc,
							_type.GetPrimitiveType(int(descriptor[0])), name)
					} else {
						lv = localvariable.NewObjectLocalVariable(m.typeMaker, index, startPc,
							m.typeMaker.MakeFromSignature(descriptor[dimension:]).CreateType(dimension), name)
					}
				}

				m.localVariableSet.Add(index, lv)
				m.names.Add(name)
			}
		}

		localVariableTypeTable := code.Attribute("LocalVariableTypeTable").(intcls.IAttributeLocalVariableTypeTable)

		if localVariableTypeTable != nil {
			updateTypeVisitor := NewUpdateTypeVisitor(m.localVariableSet)

			for _, lv := range localVariableTypeTable.LocalVariableTypeTable() {
				updateTypeVisitor.SetLocalVariableType(lv)
				m.typeMaker.MakeFromSignature(lv.Signature()).AcceptTypeArgumentVisitor(updateTypeVisitor)
			}
		}
	}
}

func (m *LocalVariableMaker) initLocalVariablesFromParameterTypes(classFile intcls.IClassFile,
	parameterTypes intmod.IType, varargs bool, firstVariableIndex, lastParameterIndex int) {
	typeMap := make(map[intmod.IType]bool)
	t := util.NewDefaultListWithSlice[intmod.IType](parameterTypes.ToSlice())

	for parameterIndex := 0; parameterIndex <= lastParameterIndex; parameterIndex++ {
		y := t.Get(parameterIndex)
		_, ok := typeMap[y]
		typeMap[y] = ok
	}

	parameterNamePrefix := "param"

	if classFile.OuterClassFile() != nil {
		innerTypeDepth := 1
		y := m.typeMaker.MakeFromInternalTypeName(classFile.OuterClassFile().InternalTypeName()).(intmod.IType)

		for y != nil && y.IsInnerObjectType() {
			innerTypeDepth++
			y = y.OuterType().(intmod.IType)
		}

		parameterNamePrefix = fmt.Sprintf("%s%d", parameterNamePrefix, innerTypeDepth)
	}

	sb := ""
	generateParameterSuffixNameVisitor := NewGenerateParameterSuffixNameVisitor()

	variableIndex := firstVariableIndex
	for parameterIndex := 0; parameterIndex <= lastParameterIndex; parameterIndex++ {
		typ := t.Get(parameterIndex)
		lv := m.localVariableSet.Root(variableIndex).(intsrv.ILocalVariable)

		if lv == nil {
			sb = ""
			sb += parameterNamePrefix

			if (parameterIndex == lastParameterIndex) && varargs {
				sb += "VarArgs"
				//                } else if (type.Dimension() > 1) {
				//                    sb.append("ArrayOfArray");
			} else {
				if typ.Dimension() > 0 {
					sb += "ArrayOf"
				}
				typ.AcceptTypeArgumentVisitor(generateParameterSuffixNameVisitor)
				sb += generateParameterSuffixNameVisitor.Suffix()
			}

			length := len(sb)
			counter := 1

			if typeMap[typ] {
				sb += fmt.Sprintf("%d", counter)
				counter++
			}

			name := sb

			for _, value := range m.names.ToSlice() {
				if value == name {
					sb = sb[:length]
					sb += fmt.Sprintf("%d", counter)
					counter++
					name = sb
				}
			}

			m.names.Add(name)
			m.createParameterVisitor.Init(variableIndex, name)
			typ.AcceptTypeArgumentVisitor(m.createParameterVisitor)
			alv := m.createParameterVisitor.LocalVariable().(intsrv.ILocalVariable)
			alv.SetDeclared(true)
			m.localVariableSet.Add(variableIndex, alv)
		}

		if _type.PtTypeLong == typ || _type.PtTypeDouble == typ {
			variableIndex++
		}
		variableIndex++
	}
}

func (m *LocalVariableMaker) LocalVariable(index, offset int) intsrv.ILocalVariable {
	lv := m.localVariableCache[index]

	if lv == nil {
		lv = m.currentFrame.LocalVariable(index).(intsrv.ILocalVariable)
		//            assert lv != nil : "getLocalVariable : local variable not found";
		if lv == nil {
			lv = localvariable.NewObjectLocalVariable2(m.typeMaker, index, offset, _type.OtTypeObject.(intmod.IType),
				fmt.Sprintf("SYNTHETIC_LOCAL_VARIABLE_%d", index), true)
		}
	} else if lv.Frame() != m.currentFrame {
		frame := searchCommonParentFrame(lv.Frame(), m.currentFrame)
		frame.MergeLocalVariable(m.typeBounds, m, lv)

		if lv.Frame() != frame {
			lv.Frame().RemoveLocalVariable(lv)
			frame.AddLocalVariable(lv)
		}
	}

	lv.SetToOffset(offset)

	return lv
}

func (m *LocalVariableMaker) searchLocalVariable(index, offset int) intsrv.ILocalVariable {
	lv := m.localVariableSet.Get(index, offset)

	if lv == nil {
		lv = m.currentFrame.LocalVariable(index).(intsrv.ILocalVariable)
	} else {
		lv2 := m.currentFrame.LocalVariable(index).(intsrv.ILocalVariable)

		if lv2 != nil && lv.Type() == lv2.Type() {
			if lv.Name() == "" && lv2.Name() == "" || lv.Name() == lv2.Name() {
				lv = lv2
			}
		}

		m.localVariableSet.Remove(index, offset)
	}

	return lv
}

func (m *LocalVariableMaker) IsCompatible(lv intsrv.ILocalVariable, valueType intmod.IType) bool {
	if valueType == _type.OtTypeUndefinedObject.(intmod.IType) {
		return true
	} else if valueType.IsObjectType() && (lv.Type().Dimension() == valueType.Dimension()) {
		valueObjectType := valueType.(intmod.IObjectType)

		if lv.Type().IsObjectType() {
			lvObjectType := lv.Type().(intmod.IObjectType)

			lvTypeArguments := lvObjectType.TypeArguments()
			valueTypeArguments := valueObjectType.TypeArguments()

			if (lvTypeArguments == nil) || (valueTypeArguments == nil) || (valueTypeArguments == _type.WildcardTypeArgumentEmpty) {
				return m.typeMaker.IsRawTypeAssignable(lvObjectType, valueObjectType)
			}

			m.searchInTypeArgumentVisitor.Init()
			lvTypeArguments.AcceptTypeArgumentVisitor(m.searchInTypeArgumentVisitor)

			if !m.searchInTypeArgumentVisitor.ContainsGeneric() {
				m.searchInTypeArgumentVisitor.Init()
				valueTypeArguments.AcceptTypeArgumentVisitor(m.searchInTypeArgumentVisitor)

				if m.searchInTypeArgumentVisitor.ContainsGeneric() {
					return m.typeMaker.IsRawTypeAssignable(lvObjectType, valueObjectType)
				}
			}
		} else if lv.Type().IsGenericType() && valueObjectType.InternalName() == _type.OtTypeObject.InternalName() {
			return true
		}
	}

	return false
}

func (m *LocalVariableMaker) LocalVariableInAssignment(typeBounds map[string]intmod.IType, index, offset int, valueType intmod.IType) intsrv.ILocalVariable {
	lv := m.searchLocalVariable(index, offset)

	if lv == nil {
		// Create a new local variable
		m.createLocalVariableVisitor.Init(index, offset)
		valueType.AcceptTypeArgumentVisitor(m.createLocalVariableVisitor)
		lv = m.createLocalVariableVisitor.LocalVariable()
	} else if lv.IsAssignableFrom(typeBounds, valueType) || m.IsCompatible(lv, valueType) {
		// Assignable, reduce type
		lv.TypeOnRight(typeBounds, valueType)
	} else if !lv.Type().IsGenericType() || _type.OtTypeObject.(intmod.IType) != valueType {
		// Not assignable -> Create a new local variable
		m.createLocalVariableVisitor.Init(index, offset)
		valueType.AcceptTypeArgumentVisitor(m.createLocalVariableVisitor)
		lv = m.createLocalVariableVisitor.LocalVariable()
	}

	lv.SetToOffset(offset)
	m.store(lv)

	return lv
}

func (m *LocalVariableMaker) LocalVariableInNullAssignment(index, offset int, valueType intmod.IType) intsrv.ILocalVariable {
	lv := m.searchLocalVariable(index, offset)

	if lv == nil {
		// Create a new local variable
		m.createLocalVariableVisitor.Init(index, offset)
		valueType.AcceptTypeArgumentVisitor(m.createLocalVariableVisitor)
		lv = m.createLocalVariableVisitor.LocalVariable()
	} else {
		t := lv.Type()

		if (t.Dimension() == 0) && t.IsPrimitiveType() {
			// Not assignable -> Create a new local variable
			m.createLocalVariableVisitor.Init(index, offset)
			valueType.AcceptTypeArgumentVisitor(m.createLocalVariableVisitor)
			lv = m.createLocalVariableVisitor.LocalVariable()
		}
	}

	lv.SetToOffset(offset)
	m.store(lv)

	return lv
}

func (m *LocalVariableMaker) LocalVariableInAssignmentWithLocalVariable(typeBounds map[string]intmod.IType, index, offset int, valueLocalVariable intsrv.ILocalVariable) intsrv.ILocalVariable {
	lv := m.searchLocalVariable(index, offset)

	if lv == nil {
		// Create a new local variable
		m.createLocalVariableVisitor.Init(index, offset)
		valueLocalVariable.Accept(m.createLocalVariableVisitor)
		lv = m.createLocalVariableVisitor.LocalVariable()
	} else if lv.IsAssignableFrom(typeBounds, valueLocalVariable.(intmod.IType)) || m.IsCompatible(lv, valueLocalVariable.Type()) {
		// Assignable
	} else if !lv.Type().IsGenericType() || _type.OtTypeObject.(intmod.IType) != valueLocalVariable.Type() {
		// Not assignable -> Create a new local variable
		m.createLocalVariableVisitor.Init(index, offset)
		valueLocalVariable.Accept(m.createLocalVariableVisitor)
		lv = m.createLocalVariableVisitor.LocalVariable()
	}

	lv.VariableOnRight(typeBounds, valueLocalVariable)
	lv.SetToOffset(offset)
	m.store(lv)

	return lv
}

func (m *LocalVariableMaker) ExceptionLocalVariable(index, offset int, t intmod.IObjectType) intsrv.ILocalVariable {
	var lv intsrv.ILocalVariable

	if index == -1 {
		lv = localvariable.NewObjectLocalVariable2(m.typeMaker, index, offset, t.(intmod.IType), "", true)
		m.currentFrame.SetExceptionLocalVariable(lv)
	} else {
		lv = m.localVariableSet.Remove(index, offset)

		if lv == nil {
			lv = localvariable.NewObjectLocalVariable2(m.typeMaker, index, offset, t.(intmod.IType), "", true)
		} else {
			lv.SetDeclared(true)
		}

		m.currentFrame.AddLocalVariable(lv)
	}

	return lv
}

func (m *LocalVariableMaker) RemoveLocalVariable(lv intsrv.ILocalVariable) {
	index := lv.Index()

	if index < len(m.localVariableCache) {
		// Remove from cache
		m.localVariableCache[index] = nil
		// Remove from current frame
		m.currentFrame.RemoveLocalVariable(lv)
	}
}

func (m *LocalVariableMaker) store(lv intsrv.ILocalVariable) {
	// Store to cache
	index := lv.Index()

	if index >= len(m.localVariableCache) {
		tmp := m.localVariableCache
		m.localVariableCache = make([]intsrv.ILocalVariable, 0, len(tmp)*2)
		for _, item := range tmp {
			m.localVariableCache = append(m.localVariableCache, item)
		}
	}

	// FIXME: 오류 발생 예상 지점.
	m.localVariableCache[index] = lv

	// Store to current frame
	if lv.Frame() == nil {
		m.currentFrame.AddLocalVariable(lv)
	}
}

func (m *LocalVariableMaker) ContainsName(name string) bool {
	for _, item := range m.names.ToSlice() {
		if item == name {
			return true
		}
	}

	return false
}

func (m *LocalVariableMaker) Make(containsLineNumber bool, typeMaker intsrv.ITypeMaker) {
	m.currentFrame.UpdateLocalVariableInForStatements(typeMaker)
	m.currentFrame.CreateNames(m.blackListNames.ToSlice())
	m.currentFrame.CreateDeclarations(containsLineNumber)
}

func (m *LocalVariableMaker) FormalParameters() intmod.IFormalParameter {
	return m.formalParameters
}

func (m *LocalVariableMaker) PushFrame(statements intmod.IStatements) {
	parent := m.currentFrame
	m.currentFrame = localvariable.NewFrame(m.currentFrame, statements)
	parent.AddChild(m.currentFrame)
}

func (m *LocalVariableMaker) PopFrame() {
	m.currentFrame.Close()
	m.currentFrame = m.currentFrame.Parent()
}

func (m *LocalVariableMaker) ChangeFrame(localVariable intsrv.ILocalVariable) {
	frame := searchCommonParentFrame(localVariable.Frame(), m.currentFrame)

	if localVariable.Frame() != frame {
		localVariable.Frame().RemoveLocalVariable(localVariable)
		frame.AddLocalVariable(localVariable)
	}
}

func searchCommonParentFrame(frame1, frame2 intsrv.IFrame) intsrv.IFrame {
	if frame1 == frame2 {
		return frame1
	}

	if frame2.Parent() == frame1 {
		return frame1
	}

	if frame1.Parent() == frame2 {
		return frame2
	}

	set := util.NewDefaultList[intsrv.IFrame]()

	for frame1 != nil {
		set.Add(frame1)
		frame1 = frame1.Parent()
	}

	for frame2 != nil {
		if set.Contains(frame2) {
			return frame2
		}
		frame2 = frame2.Parent()
	}

	return nil
}
