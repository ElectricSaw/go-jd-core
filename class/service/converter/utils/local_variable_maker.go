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
	return &LocalVariableMaker{
		typeMaker: typeMaker,
	}
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
	code := method.Attribute("Code").(*attribute.AttributeCode);

	// Init local variables from attributes
	if code != nil {
		localVariableTable := code.Attribute("LocalVariableTable").(*attribute.AttributeLocalVariableTable)

		if localVariableTable != nil {
			staticFlag := (method.AccessFlags() & intmod.FlagStatic) != 0;

			for _, localVariable := range localVariableTable.LocalVariableTable() {
				index := localVariable.Index();
				startPc := 0
				if !staticFlag && index==0 {
					startPc = localVariable.StartPc();
				}
				descriptor := localVariable.Descriptor();
				name := localVariable.Name();
				var lv intsrv.ILocalVariable;

				if descriptor[len(descriptor) - 1] == ';' {
					lv = localvariable.NewObjectLocalVariable(m.typeMaker, index, startPc,
						m.typeMaker.MakeFromDescriptor(descriptor).(intmod.IType), name);
				} else {
					dimension := CountDimension(descriptor);

					if dimension == 0 {
						lv = localvariable.NewPrimitiveLocalVariable(index, startPc,
							_type.GetPrimitiveType(int(descriptor[0])), name);
					} else {
						lv = localvariable.NewObjectLocalVariable(m.typeMaker, index, startPc,
							m.typeMaker.MakeFromSignature(descriptor[dimension:]).CreateType(dimension), name);
					}
				}

				m.localVariableSet.Add(index, lv);
				m.names = append(m.names, name);
			}
		}

		localVariableTypeTable := code.Attribute("LocalVariableTypeTable").(*attribute.AttributeLocalVariableTypeTable)

		if localVariableTypeTable != nil {
			updateTypeVisitor := visitor.NewUpdateTypeVisitor(m.localVariableSet);

			for _, lv := range localVariableTypeTable.LocalVariableTypeTable() {
				updateTypeVisitor.SetLocalVariableType(lv);
				m.typeMaker.MakeFromSignature(lv.Signature()).AcceptTypeArgumentVisitor(updateTypeVisitor);
			}
		}
	}
}

func (m *LocalVariableMaker) initLocalVariablesFromParameterTypes(classFile intmod.IClassFile,
	parameterTypes intmod.IType, varargs bool,  firstVariableIndex, lastParameterIndex int) {
	typeMap := make(map[intmod.IType]bool)
	t := sliceToDefaultList[intmod.IType](parameterTypes.List());

	for parameterIndex:=0; parameterIndex<=lastParameterIndex; parameterIndex++ {
		y := t.Get(parameterIndex);
		_, ok := typeMap[y]
		typeMap[y] = ok
	}

	parameterNamePrefix := "param";

	if classFile.OuterClassFile() != nil {
		innerTypeDepth := 1;
		y := m.typeMaker.MakeFromInternalTypeName(classFile.OuterClassFile().InternalTypeName()).(intmod.IType)

		for y != nil && y.IsInnerObjectType() {
			innerTypeDepth++
			y = y.OuterType().(intmod.IType)
		}

		parameterNamePrefix = fmt.Sprintf("%s%d", parameterNamePrefix, innerTypeDepth)
	}

	StringBuilder sb = new StringBuilder();
	GenerateParameterSuffixNameVisitor generateParameterSuffixNameVisitor = new GenerateParameterSuffixNameVisitor();

	for (int parameterIndex=0, variableIndex=firstVariableIndex; parameterIndex<=lastParameterIndex; parameterIndex++, variableIndex++) {
		Type type = t.(parameterIndex);
		AbstractLocalVariable lv = localVariableSet.root(variableIndex);

		if (lv == nil) {
			sb.SetLength(0);
			sb.append(parameterNamePrefix);

			if ((parameterIndex == lastParameterIndex) && varargs) {
				sb.append("VarArgs");
				//                } else if (type.Dimension() > 1) {
				//                    sb.append("ArrayOfArray");
			} else {
				if (type.Dimension() > 0) {
					sb.append("ArrayOf");
				}
				type.accept(generateParameterSuffixNameVisitor);
				sb.append(generateParameterSuffixNameVisitor.Suffix());
			}

			int length = sb.length();
			int counter = 1;

			if (typeMap.(type)) {
				sb.append(counter++);
			}

			String name = sb.toString();

			while (names.contains(name)) {
				sb.SetLength(length);
				sb.append(counter++);
				name = sb.toString();
			}

			names.add(name);
			createParameterVisitor.init(variableIndex, name);
			type.accept(createParameterVisitor);

			AbstractLocalVariable alv = createParameterVisitor.LocalVariable();

			alv.SetDeclared(true);
			localVariableSet.add(variableIndex, alv);
		}

		if (PrimitiveType.TYPE_LONG.equals(type) || PrimitiveType.TYPE_DOUBLE.equals(type)) {
variableIndex++;
}
}
}

func (m *LocalVariableMaker) getLocalVariable(int index, int offset)  AbstractLocalVariable {
	AbstractLocalVariable lv = localVariableCache[index];

	if (lv == nil) {
		lv = currentFrame.LocalVariable(index);
		//            assert lv != nil : "getLocalVariable : local variable not found";
		if (lv == nil) {
			lv = new ObjectLocalVariable(typeMaker, index, offset, ObjectType.TYPE_OBJECT, "SYNTHETIC_LOCAL_VARIABLE_"+index, true);
		}
	} else if (lv.Frame() != currentFrame) {
		Frame frame = searchCommonParentFrame(lv.Frame(), currentFrame);
		frame.mergeLocalVariable(typeBounds, this, lv);

		if (lv.Frame() != frame) {
			lv.Frame().removeLocalVariable(lv);
			frame.addLocalVariable(lv);
		}
	}

	lv.SetToOffset(offset);

	return lv;
}

func (m *LocalVariableMaker) searchLocalVariable(int index, int offset)  AbstractLocalVariable {
	AbstractLocalVariable lv = localVariableSet.(index, offset);

	if (lv == nil) {
		lv = currentFrame.LocalVariable(index);
	} else {
		AbstractLocalVariable lv2 = currentFrame.LocalVariable(index);

		if ((lv2 != nil) && ((lv.Name() == nil) ? (lv2.Name() == nil) : lv.Name().equals(lv2.Name())) && lv.Type().equals(lv2.Type())) {
			lv = lv2;
		}

		localVariableSet.remove(index, offset);
	}

	return lv;
}

func (m *LocalVariableMaker) isCompatible(AbstractLocalVariable lv, Type valueType)  boolean {
	if (valueType == ObjectType.TYPE_UNDEFINED_OBJECT) {
		return true;
	} else if (valueType.isObjectType() && (lv.Type().Dimension() == valueType.Dimension())) {
		ObjectType valueObjectType = (ObjectType) valueType;

		if (lv.Type().isObjectType()) {
			ObjectType lvObjectType = (ObjectType) lv.Type();

			BaseTypeArgument lvTypeArguments = lvObjectType.TypeArguments();
			BaseTypeArgument valueTypeArguments = valueObjectType.TypeArguments();

			if ((lvTypeArguments == nil) || (valueTypeArguments == nil) || (valueTypeArguments == WildcardTypeArgument.WILDCARD_TYPE_ARGUMENT)) {
				return typeMaker.isRawTypeAssignable(lvObjectType, valueObjectType);
			}

			searchInTypeArgumentVisitor.init();
			lvTypeArguments.accept(searchInTypeArgumentVisitor);

			if (!searchInTypeArgumentVisitor.containsGeneric()) {
				searchInTypeArgumentVisitor.init();
				valueTypeArguments.accept(searchInTypeArgumentVisitor);

				if (searchInTypeArgumentVisitor.containsGeneric()) {
					return typeMaker.isRawTypeAssignable(lvObjectType, valueObjectType);
				}
			}
		} else if (lv.Type().isGenericType() && valueObjectType.InternalName().equals(ObjectType.TYPE_OBJECT.InternalName())) {
			return true;
		}
	}

	return false;
}

func (m *LocalVariableMaker) getLocalVariableInAssignment(Map<String, BaseType> typeBounds, int index, int offset, Type valueType)  AbstractLocalVariable {
	AbstractLocalVariable lv = searchLocalVariable(index, offset);

	if (lv == nil) {
		// Create a new local variable
		createLocalVariableVisitor.init(index, offset);
		valueType.accept(createLocalVariableVisitor);
		lv = createLocalVariableVisitor.LocalVariable();
	} else if (lv.isAssignableFrom(typeBounds, valueType) || isCompatible(lv, valueType)) {
		// Assignable, reduce type
		lv.typeOnRight(typeBounds, valueType);
	} else if (!lv.Type().isGenericType() || (ObjectType.TYPE_OBJECT != valueType)) {
		// Not assignable -> Create a new local variable
		createLocalVariableVisitor.init(index, offset);
		valueType.accept(createLocalVariableVisitor);
		lv = createLocalVariableVisitor.LocalVariable();
	}

	lv.SetToOffset(offset);
	store(lv);

	return lv;
}

func (m *LocalVariableMaker) getLocalVariableInnilAssignment(int index, int offset, Type valueType)  AbstractLocalVariable {
	AbstractLocalVariable lv = searchLocalVariable(index, offset);

	if (lv == nil) {
		// Create a new local variable
		createLocalVariableVisitor.init(index, offset);
		valueType.accept(createLocalVariableVisitor);
		lv = createLocalVariableVisitor.LocalVariable();
	} else {
		Type type = lv.Type();

		if ((type.Dimension() == 0) && type.isPrimitiveType()) {
			// Not assignable -> Create a new local variable
			createLocalVariableVisitor.init(index, offset);
			valueType.accept(createLocalVariableVisitor);
			lv = createLocalVariableVisitor.LocalVariable();
		}
	}

	lv.SetToOffset(offset);
	store(lv);

	return lv;
}

func (m *LocalVariableMaker) getLocalVariableInAssignment(Map<String, BaseType> typeBounds, int index, int offset, AbstractLocalVariable valueLocalVariable)  AbstractLocalVariable {
	AbstractLocalVariable lv = searchLocalVariable(index, offset);

	if (lv == nil) {
		// Create a new local variable
		createLocalVariableVisitor.init(index, offset);
		valueLocalVariable.accept(createLocalVariableVisitor);
		lv = createLocalVariableVisitor.LocalVariable();
	} else if (lv.isAssignableFrom(typeBounds, valueLocalVariable) || isCompatible(lv, valueLocalVariable.Type())) {
		// Assignable
	} else if (!lv.Type().isGenericType() || (ObjectType.TYPE_OBJECT != valueLocalVariable.Type())) {
		// Not assignable -> Create a new local variable
		createLocalVariableVisitor.init(index, offset);
		valueLocalVariable.accept(createLocalVariableVisitor);
		lv = createLocalVariableVisitor.LocalVariable();
	}

	lv.variableOnRight(typeBounds, valueLocalVariable);
	lv.SetToOffset(offset);
	store(lv);

	return lv;
}

func (m *LocalVariableMaker) getExceptionLocalVariable(int index, int offset, ObjectType type)  AbstractLocalVariable {
AbstractLocalVariable lv;

if (index == -1) {
currentFrame.SetExceptionLocalVariable(lv = new ObjectLocalVariable(typeMaker, index, offset, type, nil, true));
} else {
lv = localVariableSet.remove(index, offset);

if (lv == nil) {
lv = new ObjectLocalVariable(typeMaker, index, offset, type, nil, true);
} else {
lv.SetDeclared(true);
}

currentFrame.addLocalVariable(lv);
}

return lv;
}

func (m *LocalVariableMaker) removeLocalVariable(AbstractLocalVariable lv)  void {
	int index = lv.Index();

	if (index < localVariableCache.length) {
		// Remove from cache
		localVariableCache[index] = nil;
		// Remove from current frame
		currentFrame.removeLocalVariable(lv);
	}
}

func (m *LocalVariableMaker) store(lv AbstractLocalVariable) {
	// Store to cache
	int index = lv.Index();

	if (index >= localVariableCache.length) {
		AbstractLocalVariable[] tmp = localVariableCache;
		localVariableCache = new AbstractLocalVariable[index * 2];
		System.arraycopy(tmp, 0, localVariableCache, 0, tmp.length);
	}

	localVariableCache[index] = lv;

	// Store to current frame
	if (lv.Frame() == nil) {
		currentFrame.addLocalVariable(lv);
	}
}

func (m *LocalVariableMaker) containsName(String name)  boolean {
	return names.contains(name);
}

func (m *LocalVariableMaker) make(boolean containsLineNumber, TypeMaker typeMaker)  void {
	currentFrame.updateLocalVariableInForStatements(typeMaker);
	currentFrame.createNames(blackListNames);
	currentFrame.createDeclarations(containsLineNumber);
}

func (m *LocalVariableMaker) getFormalParameters()  BaseFormalParameter {
	return formalParameters;
}

func (m *LocalVariableMaker) pushFrame(Statements statements)  void {
	Frame parent = currentFrame;
	currentFrame = new Frame(currentFrame, statements);
	parent.addChild(currentFrame);
}

func (m *LocalVariableMaker) popFrame()  void {
	currentFrame.close();
	currentFrame = currentFrame.Parent();
}

func (m *LocalVariableMaker) Frame searchCommonParentFrame(Frame frame1, Frame frame2)  static {
	if (frame1 == frame2) {
		return frame1;
	}

	if (frame2.Parent() == frame1) {
		return frame1;
	}

	if (frame1.Parent() == frame2) {
		return frame2;
	}

	HashSet<Frame> set = new HashSet<>();

	while (frame1 != nil) {
		set.add(frame1);
		frame1 = frame1.Parent();
	}

	while (frame2 != nil) {
		if (set.contains(frame2)) {
			return frame2;
		}
		frame2 = frame2.Parent();
	}

	return nil;
}

func (m *LocalVariableMaker) changeFrame(AbstractLocalVariable localVariable)  void {
	Frame frame = LocalVariableMaker.searchCommonParentFrame(localVariable.Frame(), currentFrame);

	if (localVariable.Frame() != frame) {
		localVariable.Frame().removeLocalVariable(localVariable);
		frame.addLocalVariable(localVariable);
	}
}

func sliceToDefaultList[T any](slice []T) util.DefaultList[T] {
	ret := util.DefaultList[T]{}
	for _, item := range slice {
		ret.Add(item)
	}
	return ret
}