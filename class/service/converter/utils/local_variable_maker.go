package utils

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/attribute"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/model/localvariable"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/visitor"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewLocalVariableMaker(typeMaker *TypeMaker, comd intsrv.IClassFileConstructorOrMethodDeclaration, constructor bool) *LocalVariableMaker {
	m := &LocalVariableMaker{
		typeMaker: typeMaker,
	}

	return m
}

type LocalVariableMaker struct {
	localVariableSet              intsrv.ILocalVariableSet
	names                         []string
	blackListNames                []string
	currentFrame                  intsrv.IRootFrame
	localVariableCache            []intsrv.ILocalVariable
	typeMaker                     *TypeMaker
	typeBounds                    map[string]intmod.IType
	formalParameters              declaration.FormalParameter
	populateBlackListNamesVisitor *visitor.PopulateBlackListNamesVisitor
	searchInTypeArgumentVisitor   *visitor.SearchInTypeArgumentVisitor
	createParameterVisitor        *visitor.CreateParameterVisitor
	createLocalVariableVisitor    *visitor.CreateLocalVariableVisitor
}

func (m *LocalVariableMaker) initLocalVariablesFromAttributes(method intmod.IMethod) {
	code := method.Attribute("Code").(*attribute.AttributeCode)

	// Init local variables from attributes
	if code != nil {
		localVariableTable := code.Attribute("LocalVariableTable").(*attribute.AttributeLocalVariableTable)

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
				m.names = append(m.names, name)
			}
		}

		localVariableTypeTable := code.Attribute("LocalVariableTypeTable").(*attribute.AttributeLocalVariableTypeTable)

		if localVariableTypeTable != nil {
			updateTypeVisitor := visitor.NewUpdateTypeVisitor(m.localVariableSet)

			for _, lv := range localVariableTypeTable.LocalVariableTypeTable() {
				updateTypeVisitor.SetLocalVariableType(lv)
				m.typeMaker.MakeFromSignature(lv.Signature()).AcceptTypeArgumentVisitor(updateTypeVisitor)
			}
		}
	}
}

func (m *LocalVariableMaker) initLocalVariablesFromParameterTypes(classFile intmod.IClassFile,
	parameterTypes intmod.IType, varargs bool, firstVariableIndex, lastParameterIndex int) {
	typeMap := make(map[intmod.IType]bool)
	t := sliceToDefaultList[intmod.IType](parameterTypes.ToSlice())

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
	generateParameterSuffixNameVisitor := visitor.NewGenerateParameterSuffixNameVisitor()

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

			for _, value := range m.names {
				if value == name {
					sb = sb[:length]
					sb += fmt.Sprintf("%d", counter)
					counter++
					name = sb
				}
			}

			m.names = append(m.names, name)
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
		lv = m.currentFrame.LocalVariable(index)
		//            assert lv != nil : "getLocalVariable : local variable not found";
		if lv == nil {
			lv = localvariable.NewObjectLocalVariable2(m.typeMaker, index, offset, _type.OtTypeObject.(intmod.IType),
				fmt.Sprintf("SYNTHETIC_LOCAL_VARIABLE_%d", index), true)
		}
	} else if lv.Frame() != m.currentFrame {
		frame := searchCommonParentFrame(lv.Frame(), currentFrame)
		frame.MergeLocalVariable(typeBounds, this, lv)

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
		lv = m.currentFrame.LocalVariable(index)
	} else {
		lv2 := m.currentFrame.LocalVariable(index)

		if lv2 != nil && lv.Type() == lv2.Type() {
			if lv.Name() == "" && lv2.Name() == nil || lv.Name() == lv2.Name() {
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
			lv = localvariable.NewObjectLocalVariable2(m.typeMaker, index, offset, t.(intmod.IType), nil, true)
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
		m.currentFrame.addLocalVariable(lv)
	}
}

func (m *LocalVariableMaker) ContainsName(name string) bool {
	for _, item := range m.names {
		if item == name {
			return true
		}
	}

	return false
}

func (m *LocalVariableMaker) Make(containsLineNumber bool, typeMaker *TypeMaker) {
	m.currentFrame.updateLocalVariableInForStatements(typeMaker)
	m.currentFrame.createNames(m.blackListNames)
	m.currentFrame.createDeclarations(containsLineNumber)
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

func searchCommonParentFrame(frame1, frame2 localvariable.IFrame) localvariable.IFrame {
	if frame1 == frame2 {
		return frame1
	}

	if frame2.Parent() == frame1 {
		return frame1
	}

	if frame1.Parent() == frame2 {
		return frame2
	}

	set := make([]localvariable.IFrame, 0)

	for frame1 != nil {
		set = append(set, frame1)
		frame1 = frame1.Parent()
	}

	for frame2 != nil {
		if ContainsFrame(set, frame2) {
			return frame2
		}
		frame2 = frame2.Parent()
	}

	return nil
}

func sliceToDefaultList[T any](slice []T) util.DefaultList[T] {
	ret := util.DefaultList[T]{}
	for _, item := range slice {
		ret.Add(item)
	}
	return ret
}

func ContainsFrame(slice []localvariable.IFrame, item localvariable.IFrame) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
