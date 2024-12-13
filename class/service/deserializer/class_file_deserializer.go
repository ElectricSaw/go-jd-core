package deserializer

import (
	"errors"
	"fmt"
	"github.com/ElectricSaw/go-jd-core/class/api"
	intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/classfile"
	"github.com/ElectricSaw/go-jd-core/class/model/classfile/attribute"
	"github.com/ElectricSaw/go-jd-core/class/model/classfile/constant"
	"log"
	"strings"
	"unicode"
)

var EmptyIntArray = []int{}

func GetConstantTypeName(constants intcls.IConstantPool, index int, msg string) string {
	name, ok := constants.ConstantTypeName(index)
	if !ok {
		log.Fatalf(msg)
		return ""
	}
	return name
}

type ClassFileDeserializer struct {
}

func (d *ClassFileDeserializer) LoadClassFileWithRaw(loader api.Loader, internalTypeName string) (intmod.IClassFile, error) {
	classFile, err := d.InnerLoadClassFile(loader, internalTypeName)
	if err != nil {
		return nil, err
	}
	return classFile, nil
}

func (d *ClassFileDeserializer) InnerLoadClassFile(loader api.Loader, internalTypeName string) (intmod.IClassFile, error) {
	if !loader.CanLoad(internalTypeName) {
		return nil, errors.New("Can't load '" + internalTypeName + "'")
	}

	data, err := loader.Load(internalTypeName)
	if err != nil {
		return nil, err
	}

	reader := NewClassFileReader(data)
	classFile, err := d.LoadClassFile(reader)
	if err != nil {
		return nil, err
	}

	var aic *attribute.AttributeInnerClasses
	if v, ok := classFile.Attributes()["InnerClasses"]; ok {
		aic = v.(*attribute.AttributeInnerClasses)
	}

	if aic != nil {
		var innerClassFiles []intmod.IClassFile
		innerTypePrefix := internalTypeName + "$"

		for _, ic := range aic.InnerClasses() {
			innerTypeName := ic.InnerTypeName()

			if !(internalTypeName == innerTypeName) {
				if (internalTypeName == ic.OuterTypeName()) || (strings.HasPrefix(innerTypeName, innerTypePrefix)) {
					innerClassFile, err := d.InnerLoadClassFile(loader, innerTypeName)
					if err != nil {
						fmt.Printf("inner class file load failed: %v\n", innerClassFile)
						continue
					}

					flags := ic.InnerAccessFlags()
					var length int

					if strings.HasPrefix(innerTypeName, innerTypePrefix) {
						length = len(internalTypeName) + 1
					} else {
						length = strings.Index(innerTypeName, "$") + 1
					}

					if unicode.IsDigit(rune(innerTypeName[length])) {
						flags |= intmod.AccSynthetic
					}

					if innerClassFile == nil {
						innerClassFile = classfile.NewClassFile(
							classFile.MajorVersion(),
							classFile.MinorVersion(),
							0,
							innerTypeName,
							"class/lang/Object",
							nil,
							nil,
							nil,
							nil,
						)
					}

					innerClassFile.SetOuterClassFile(classFile)
					innerClassFile.SetAccessFlags(flags)
					innerClassFiles = append(innerClassFiles, innerClassFile)
				}
			}
		}

		if len(innerClassFiles) != 0 {
			classFile.SetInnerClassFiles(innerClassFiles)
		}
	}

	return classFile, nil
}

func (d *ClassFileDeserializer) LoadClassFile(reader intsrv.IClassFileReader) (intmod.IClassFile, error) {
	magic := reader.ReadMagic()

	if magic != intsrv.JavaMagicNumber {
		return nil, errors.New("invalid CLASS file")
	}

	minorVersion := reader.ReadUnsignedShort()
	majorVersion := reader.ReadUnsignedShort()

	constantArray, err := d.LoadConstants(reader)
	if err != nil {
		return nil, err
	}
	constants := classfile.NewConstantPool(constantArray)

	accessFlags := reader.ReadUnsignedShort()
	thisClassIndex := reader.ReadUnsignedShort()
	superClassIndex := reader.ReadUnsignedShort()

	internalTypeName := GetConstantTypeName(constants, thisClassIndex, "internal type name failed.")
	superTypeName := ""
	if superClassIndex != 0 {
		superTypeName = GetConstantTypeName(constants, superClassIndex, "super type name failed.")
	}
	interfaceTypeNames, err := d.LoadInterfaces(reader, constants)
	if err != nil {
		return nil, err
	}
	fields, err := d.LoadFields(reader, constants)
	if err != nil {
		return nil, err
	}
	methods, err := d.LoadMethods(reader, constants)
	if err != nil {
		return nil, err
	}
	attributes, err := d.LoadAttributes(reader, constants)
	if err != nil {
		return nil, err
	}

	return classfile.NewClassFile(majorVersion, minorVersion, accessFlags, internalTypeName, superTypeName, interfaceTypeNames, fields, methods, attributes), nil
}

func (d *ClassFileDeserializer) LoadConstants(reader intsrv.IClassFileReader) ([]intcls.IConstant, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	constants := make([]intcls.IConstant, count)

	for i := 1; i < count; i++ {
		tag := reader.Read()

		switch tag {
		case 1:
			constants[i] = constant.NewConstantUtf8(reader.ReadUTF8())
		case 3:
			constants[i] = constant.NewConstantInteger(reader.ReadInt())
		case 4:
			constants[i] = constant.NewConstantFloat(reader.ReadFloat())
		case 5:
			constants[i] = constant.NewConstantLong(reader.ReadLong())
		case 6:
			constants[i] = constant.NewConstantDouble(reader.ReadDouble())
		case 7, 19, 20:
			constants[i] = constant.NewConstantClass(reader.ReadUnsignedShort())
		case 8:
			constants[i] = constant.NewConstantString(reader.ReadUnsignedShort())
		case 9, 10, 11, 17, 18:
			constants[i] = constant.NewConstantMemberRef(reader.ReadUnsignedShort(), reader.ReadUnsignedShort())
		case 12:
			constants[i] = constant.NewConstantNameAndType(reader.ReadUnsignedShort(), reader.ReadUnsignedShort())
		case 15:
			constants[i] = constant.NewConstantMethodHandle(int(reader.Read()), reader.ReadUnsignedShort())
		case 16:
			constants[i] = constant.NewConstantMethodType(reader.ReadUnsignedShort())
		default:
			return nil, errors.New("invalid constant pool entry")
		}
	}

	return constants, nil
}

func (d *ClassFileDeserializer) LoadInterfaces(reader intsrv.IClassFileReader, constants intcls.IConstantPool) ([]string, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	interfaceTypeNames := make([]string, count)

	for i := 0; i < count; i++ {
		index := reader.ReadUnsignedShort()
		interfaceTypeNames[i], _ = constants.ConstantTypeName(index)
	}

	return interfaceTypeNames, nil
}

func (d *ClassFileDeserializer) LoadFields(reader intsrv.IClassFileReader, constants intcls.IConstantPool) ([]intcls.IField, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	fields := make([]intcls.IField, count)

	for i := 0; i < count; i++ {
		accessFlags := reader.ReadUnsignedShort()
		nameIndex := reader.ReadUnsignedShort()
		descriptorIndex := reader.ReadUnsignedShort()
		attr, _ := d.LoadAttributes(reader, constants)

		name, _ := constants.ConstantUtf8(nameIndex)
		descriptor, _ := constants.ConstantUtf8(descriptorIndex)

		fields[i] = classfile.NewField(accessFlags, name, descriptor, attr)
	}

	return fields, nil
}

func (d *ClassFileDeserializer) LoadMethods(reader intsrv.IClassFileReader, constants intcls.IConstantPool) ([]intcls.IMethod, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	methods := make([]intcls.IMethod, count)

	for i := 0; i < count; i++ {
		accessFlags := reader.ReadUnsignedShort()
		nameIndex := reader.ReadUnsignedShort()
		descriptorIndex := reader.ReadUnsignedShort()
		attr, _ := d.LoadAttributes(reader, constants)

		name, _ := constants.ConstantUtf8(nameIndex)
		descriptor, _ := constants.ConstantUtf8(descriptorIndex)

		methods[i] = classfile.NewMethod(accessFlags, name, descriptor, attr, constants)
	}

	return methods, nil
}

func (d *ClassFileDeserializer) LoadAttributes(reader intsrv.IClassFileReader, constants intcls.IConstantPool) (map[string]intcls.IAttribute, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	attributes := make(map[string]intcls.IAttribute)

	for i := 0; i < count; i++ {
		attributeNameIndex := reader.ReadUnsignedShort()
		attributeLength := reader.ReadInt()
		c0nst := constants.Constant(attributeNameIndex)

		if c0nst.Tag() == intcls.ConstTagUtf8 {
			name := c0nst.(*constant.ConstantUtf8).Value()

			switch name {
			case "AnnotationDefault":
				v, _ := d.LoadElementValue(reader, constants)
				attributes[name] = attribute.NewAttributeAnnotationDefault(v)
			case "BootstrapMethods":
				attributes[name] = attribute.NewAttributeBootstrapMethods(d.LoadBootstrapMethods(reader))
			case "Code":
				maxStack := reader.ReadUnsignedShort()
				maxLocals := reader.ReadUnsignedShort()
				code := d.LoadCode(reader)
				exceptionTable := d.LoadCodeExceptions(reader)
				attr, _ := d.LoadAttributes(reader, constants)
				attributes[name] = attribute.NewAttributeCode(maxStack, maxLocals, code, exceptionTable, attr)
			case "ConstantValue":
				if attributeLength != 2 {
					return nil, errors.New("invalid attribute length")
				}
				attributes[name] = attribute.NewAttributeConstantValue(d.LoadConstantValue(reader, constants))
			case "Deprecated":
				if attributeLength != 0 {
					return nil, errors.New("invalid attribute length")
				}
				attributes[name] = attribute.NewAttributeDeprecated()
			case "Exceptions":
				attributes[name] = attribute.NewAttributeExceptions(d.LoadExceptionTypeNames(reader, constants))
			case "InnerClasses":
				attributes[name] = attribute.NewAttributeInnerClasses(d.LoadInnerClasses(reader, constants))
			case "LocalVariableTable":
				v := d.LoadLocalVariables(reader, constants)
				if v != nil {
					attributes[name] = attribute.NewAttributeLocalVariableTable(v)
				}
			case "LocalVariableTypeTable":
				attributes[name] = attribute.NewAttributeLocalVariableTypeTable(d.LoadLocalVariableTypes(reader, constants))
			case "LineNumberTable":
				attributes[name] = attribute.NewAttributeLineNumberTable(d.LoadLineNumbers(reader))
			case "MethodParameters":
				attributes[name] = attribute.NewAttributeMethodParameters(d.LoadParameters(reader, constants))
			case "Module":
				moduleName, _ := constants.ConstantTypeName(reader.ReadUnsignedShort())
				version, _ := constants.ConstantUtf8(reader.ReadUnsignedShort())

				attributes[name] = attribute.NewAttributeModule(moduleName, reader.ReadUnsignedShort(), version,
					d.LoadModuleInfos(reader, constants),
					d.LoadPackageInfos(reader, constants),
					d.LoadPackageInfos(reader, constants),
					d.LoadConstantClassNames(reader, constants),
					d.LoadServiceInfos(reader, constants),
				)
			case "ModulePackages":
				attributes[name] = attribute.NewAttributeModulePackages(d.LoadConstantClassNames(reader, constants))
			case "ModuleMainClass":
				attributes[name] = attribute.NewAttributeModuleMainClass(constants.Constant(reader.ReadUnsignedShort()).(constant.ConstantClass))
			case "RuntimeInvisibleAnnotations", "RuntimeVisibleAnnotations":
				a := d.LoadAnnotations(reader, constants)
				if a != nil {
					attributes[name] = attribute.NewAnnotations(a)
				}
			case "RuntimeInvisibleParameterAnnotations", "RuntimeVisibleParameterAnnotations":
				attributes[name] = attribute.NewAttributeParameterAnnotations(d.LoadParameterAnnotations(reader, constants))
			case "Signature":
				if attributeLength != 2 {
					return nil, errors.New("invalid attribute length")
				}
				v, _ := constants.ConstantUtf8(reader.ReadUnsignedShort())
				attributes[name] = attribute.NewAttributeSignature(v)
			case "SourceFile":
				if attributeLength != 2 {
					return nil, errors.New("invalid attribute length")
				}
				v, _ := constants.ConstantUtf8(reader.ReadUnsignedShort())
				attributes[name] = attribute.NewAttributeSourceFile(v)
			case "Synthetic":
				if attributeLength != 0 {
					return nil, errors.New("invalid attribute length")
				}
				attributes[name] = attribute.NewAttributeSynthetic()
			default:
				attributes[name] = attribute.NewUnknownAttribute()
				reader.Skip(attributeLength)
			}
		} else {
			return nil, errors.New("invalid attribute")
		}
	}

	return attributes, nil
}

func (d *ClassFileDeserializer) LoadElementValue(reader intsrv.IClassFileReader, constants intcls.IConstantPool) (intcls.IElementValue, error) {
	t := reader.Read()

	switch t {
	case 'B', 'D', 'F', 'I', 'J', 'S', 'Z', 'C', 's':
		constValueIndex := reader.ReadUnsignedShort()
		constValue := constants.Constant(reader.ReadUnsignedShort()).(intcls.IConstantValue)
		return attribute.NewElementValuePrimitiveType(constValueIndex, constValue), nil
	case 'e':
		descriptorIndex := reader.ReadUnsignedShort()
		descriptor, _ := constants.ConstantUtf8(descriptorIndex)
		constNameIndex := reader.ReadUnsignedShort()
		constName, _ := constants.ConstantUtf8(constNameIndex)
		return attribute.NewElementValueEnumConstValue(descriptor, constName), nil
	case 'c':
		classInfoIndex := reader.ReadUnsignedShort()
		classInfo, _ := constants.ConstantUtf8(classInfoIndex)
		return attribute.NewElementValueClassInfo(classInfo), nil
	case '@':
		typeIndex := reader.ReadUnsignedShort()
		descriptor, _ := constants.ConstantUtf8(typeIndex)
		return attribute.NewElementValueAnnotationValue(attribute.NewAnnotation(descriptor, d.LoadElementValuePairs(reader, constants))), nil
	case '[':
		return attribute.NewElementValueArrayValue(d.LoadElementValues(reader, constants)), nil
	default:
		return nil, errors.New(fmt.Sprintf("invalid element value type: %d", t))
	}
}

func (d *ClassFileDeserializer) LoadElementValuePairs(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.IElementValuePair {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	pairs := make([]intcls.IElementValuePair, count)

	for i := 0; i < count; i++ {
		elementNameIndex := reader.ReadUnsignedShort()
		elementName, _ := constants.ConstantUtf8(elementNameIndex)
		v, _ := d.LoadElementValue(reader, constants)
		pairs[i] = attribute.NewElementValuePair(elementName, v)
	}

	return pairs
}

func (d *ClassFileDeserializer) LoadElementValues(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.IElementValue {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	values := make([]intcls.IElementValue, count)

	for i := 0; i < count; i++ {
		values[i], _ = d.LoadElementValue(reader, constants)
	}

	return values
}

func (d *ClassFileDeserializer) LoadBootstrapMethods(reader intsrv.IClassFileReader) []intcls.IBootstrapMethod {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	values := make([]intcls.IBootstrapMethod, count)

	for i := 0; i < count; i++ {
		bootstrapMethodRef := reader.ReadUnsignedShort()
		numBootstrapArguments := reader.ReadUnsignedShort()
		bootstrapArguments := make([]int, 0)

		if numBootstrapArguments == 0 {
			bootstrapArguments = EmptyIntArray
		} else {
			bootstrapArguments = make([]int, numBootstrapArguments)
			for j := 0; j < numBootstrapArguments; j++ {
				bootstrapArguments[j] = reader.ReadUnsignedShort()
			}
		}

		values[i] = attribute.NewBootstrapMethod(bootstrapMethodRef, bootstrapArguments)
	}

	return values
}

func (d *ClassFileDeserializer) LoadCode(reader intsrv.IClassFileReader) []byte {
	codeLength := reader.ReadInt()
	if codeLength == 0 {
		return nil
	}

	return reader.ReadFully(codeLength)
}

func (d *ClassFileDeserializer) LoadCodeExceptions(reader intsrv.IClassFileReader) []intcls.ICodeException {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	codeExceptions := make([]intcls.ICodeException, count)

	for i := 0; i < count; i++ {
		codeExceptions[i] = attribute.NewCodeException(i,
			reader.ReadUnsignedShort(),
			reader.ReadUnsignedShort(),
			reader.ReadUnsignedShort(),
			reader.ReadUnsignedShort(),
		)
	}

	return codeExceptions
}

func (d *ClassFileDeserializer) LoadConstantValue(reader intsrv.IClassFileReader, constants intcls.IConstantPool) intcls.IConstantValue {
	constantValueIndex := reader.ReadUnsignedShort()
	return constants.ConstantValue(constantValueIndex)
}

func (d *ClassFileDeserializer) LoadExceptionTypeNames(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []string {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	exceptionTypeNames := make([]string, count)

	for i := 0; i < count; i++ {
		exceptionClassIndex := reader.ReadUnsignedShort()
		exceptionTypeNames[i], _ = constants.ConstantTypeName(exceptionClassIndex)
	}

	return exceptionTypeNames
}

func (d *ClassFileDeserializer) LoadInnerClasses(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.IInnerClass {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	innerClasses := make([]intcls.IInnerClass, count)

	for i := 0; i < count; i++ {
		innerTypeIndex := reader.ReadUnsignedShort()
		outerTypeIndex := reader.ReadUnsignedShort()
		innerNameIndex := reader.ReadUnsignedShort()
		innerAccessFlags := reader.ReadUnsignedShort()

		innerTypeName, _ := constants.ConstantTypeName(innerTypeIndex)
		outerTypeName := ""
		if outerTypeIndex != 0 {
			outerTypeName, _ = constants.ConstantTypeName(outerTypeIndex)
		}
		innerName := ""
		if innerNameIndex != 0 {
			innerName, _ = constants.ConstantTypeName(innerNameIndex)
		}

		innerClasses[i] = attribute.NewInnerClass(innerTypeName, outerTypeName, innerName, innerAccessFlags)
	}

	return innerClasses
}

func (d *ClassFileDeserializer) LoadLocalVariables(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.ILocalVariable {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	localVariables := make([]intcls.ILocalVariable, count)

	for i := 0; i < count; i++ {
		startPc := reader.ReadUnsignedShort()
		length := reader.ReadUnsignedShort()
		nameIndex := reader.ReadUnsignedShort()
		descriptorIndex := reader.ReadUnsignedShort()
		index := reader.ReadUnsignedShort()

		name, _ := constants.ConstantUtf8(nameIndex)
		descriptor, _ := constants.ConstantUtf8(descriptorIndex)

		localVariables[i] = attribute.NewLocalVariable(startPc, length, name, descriptor, index)
	}

	return localVariables
}

func (d *ClassFileDeserializer) LoadLocalVariableTypes(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.ILocalVariableType {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	localVariables := make([]intcls.ILocalVariableType, count)

	for i := 0; i < count; i++ {
		startPc := reader.ReadUnsignedShort()
		length := reader.ReadUnsignedShort()
		nameIndex := reader.ReadUnsignedShort()
		descriptorIndex := reader.ReadUnsignedShort()
		index := reader.ReadUnsignedShort()

		name, _ := constants.ConstantUtf8(nameIndex)
		descriptor, _ := constants.ConstantUtf8(descriptorIndex)

		localVariables[i] = attribute.NewLocalVariableType(startPc, length, name, descriptor, index)
	}

	return localVariables
}

func (d *ClassFileDeserializer) LoadLineNumbers(reader intsrv.IClassFileReader) []intcls.ILineNumber {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	lineNumbers := make([]intcls.ILineNumber, count)

	for i := 0; i < count; i++ {
		lineNumbers[i] = attribute.NewLineNumber(reader.ReadUnsignedShort(), reader.ReadUnsignedShort())
	}

	return lineNumbers
}

func (d *ClassFileDeserializer) LoadParameters(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.IMethodParameter {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	parameters := make([]intcls.IMethodParameter, count)

	for i := 0; i < count; i++ {
		nameIndex := reader.ReadUnsignedShort()
		name, _ := constants.ConstantUtf8(nameIndex)
		parameters[i] = attribute.NewMethodParameter(name, reader.ReadUnsignedShort())
	}

	return parameters
}

func (d *ClassFileDeserializer) LoadModuleInfos(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.IModuleInfo {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	moduleInfos := make([]intcls.IModuleInfo, count)

	for i := 0; i < count; i++ {
		moduleInfoIndex := reader.ReadUnsignedShort()
		moduleFlags := reader.ReadUnsignedShort()
		moduleVersionIndex := reader.ReadUnsignedShort()

		moduleInfoName, _ := constants.ConstantTypeName(moduleInfoIndex)
		moduleVersion := ""
		if moduleVersionIndex == 0 {
			moduleVersion, _ = constants.ConstantUtf8(moduleVersionIndex)
		}

		moduleInfos[i] = attribute.NewModuleInfo(moduleInfoName, moduleFlags, moduleVersion)
	}

	return moduleInfos
}

func (d *ClassFileDeserializer) LoadPackageInfos(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.IPackageInfo {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	packageInfos := make([]intcls.IPackageInfo, count)

	for i := 0; i < count; i++ {
		packageInfoIndex := reader.ReadUnsignedShort()
		packageFlags := reader.ReadUnsignedShort()
		packageInfoName, _ := constants.ConstantTypeName(packageInfoIndex)

		packageInfos[i] = attribute.NewPackageInfo(packageInfoName, packageFlags, d.LoadConstantClassNames(reader, constants))
	}

	return packageInfos
}

func (d *ClassFileDeserializer) LoadConstantClassNames(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []string {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	names := make([]string, count)

	for i := 0; i < count; i++ {
		names[i], _ = constants.ConstantTypeName(reader.ReadUnsignedShort())
	}

	return names
}

func (d *ClassFileDeserializer) LoadServiceInfos(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.IServiceInfo {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	services := make([]intcls.IServiceInfo, count)

	for i := 0; i < count; i++ {
		name, _ := constants.ConstantTypeName(reader.ReadUnsignedShort())
		services[i] = attribute.NewServiceInfo(name, d.LoadConstantClassNames(reader, constants))
	}

	return services
}

func (d *ClassFileDeserializer) LoadAnnotations(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.IAnnotation {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	annotations := make([]intcls.IAnnotation, count)

	for i := 0; i < count; i++ {
		descriptorIndex := reader.ReadUnsignedShort()
		descriptor, _ := constants.ConstantUtf8(descriptorIndex)
		annotations[i] = attribute.NewAnnotation(descriptor, d.LoadElementValuePairs(reader, constants))
	}

	return annotations
}

func (d *ClassFileDeserializer) LoadParameterAnnotations(reader intsrv.IClassFileReader, constants intcls.IConstantPool) []intcls.IAnnotations {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil
	}

	parameterAnnotations := make([]intcls.IAnnotations, count)

	for i := 0; i < count; i++ {
		annotations := d.LoadAnnotations(reader, constants)
		if annotations != nil {
			parameterAnnotations[i] = attribute.NewAnnotations(annotations)
		}
	}

	return parameterAnnotations
}
