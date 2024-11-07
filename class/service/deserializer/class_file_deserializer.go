package deserializer

import (
	"bitbucket.org/coontec/javaClass/class/api"
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"bitbucket.org/coontec/javaClass/class/model/classfile/attribute"
	"bitbucket.org/coontec/javaClass/class/model/classfile/constant"
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode"
)

var EmptyIntArray = []int{}

func GetConstantTypeName(constants *classfile.ConstantPool, index int, msg string) string {
	name, ok := constants.GetConstantTypeName(index)
	if !ok {
		log.Fatalf(msg)
		return ""
	}
	return name
}

type ClassFileDeserializer struct {
}

func (d *ClassFileDeserializer) LoadClassFileWithRaw(loader api.Loader, internalTypeName string) (*classfile.ClassFile, error) {
	classFile, err := d.InnerLoadClassFile(loader, internalTypeName)
	if err != nil {
		return nil, err
	}
	return classFile, nil
}

func (d *ClassFileDeserializer) InnerLoadClassFile(loader api.Loader, internalTypeName string) (*classfile.ClassFile, error) {
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

	aic := classFile.Attribute()["InnerClasses"].(*attribute.AttributeInnerClasses)

	if aic != nil {
		var innerClassFiles []classfile.IClassFile
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
						flags |= classfile.AccSynthetic
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

					innerClassFile.OuterClassFile = classFile
					innerClassFile.SetAccessFlags(flags)
					innerClassFiles = append(innerClassFiles, innerClassFile)
				}
			}
		}

		if len(innerClassFiles) != 0 {
			classFile.InnerClassFiles = innerClassFiles
		}
	}

	return classFile, nil
}

func (d *ClassFileDeserializer) LoadClassFile(reader *ClassFileReader) (*classfile.ClassFile, error) {
	magic := reader.ReadInt()

	if magic != JavaMagicNumber {
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
		superTypeName = GetConstantTypeName(constants, thisClassIndex, "super type name failed.")
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

func (d *ClassFileDeserializer) LoadConstants(reader *ClassFileReader) ([]constant.Constant, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	constants := make([]constant.Constant, count)

	for i := 0; i < count; i++ {
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

func (d *ClassFileDeserializer) LoadInterfaces(reader *ClassFileReader, constants *classfile.ConstantPool) ([]string, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	interfaceTypeNames := make([]string, count)

	for i := 0; i < count; i++ {
		index := reader.ReadUnsignedShort()
		interfaceTypeNames[i], _ = constants.GetConstantTypeName(index)
	}

	return interfaceTypeNames, nil
}

func (d *ClassFileDeserializer) LoadFields(reader *ClassFileReader, constants *classfile.ConstantPool) ([]classfile.Field, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadMethods(reader *ClassFileReader, constants *classfile.ConstantPool) ([]classfile.Method, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadAttributes(reader *ClassFileReader, constants *classfile.ConstantPool) (map[string]attribute.Attribute, error) {
	return nil, nil
}

func (d *ClassFileDeserializer) LoadElementValue(reader *ClassFileReader, constants *classfile.ConstantPool) (attribute.ElementValue, error) {
	return nil, nil
}

func (d *ClassFileDeserializer) LoadElementValuePairs(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.ElementValuePair, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadElementValues(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.ElementValue, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadBootstrapMethods(reader *ClassFileReader) ([]attribute.BootstrapMethod, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadCode(reader *ClassFileReader) ([]byte, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadCodeExceptions(reader *ClassFileReader) ([]attribute.CodeException, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadConstantValue(reader *ClassFileReader, constants *classfile.ConstantPool) ([]constant.ConstantValue, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadExceptionTypeNames(reader *ClassFileReader, constants *classfile.ConstantPool) ([]string, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadInnerClasses(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.InnerClass, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadLocalVariables(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.LocalVariable, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadLocalVariableTypes(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.LocalVariableType, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadLineNumbers(reader *ClassFileReader) ([]attribute.LineNumber, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadParameters(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.MethodParameter, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadModuleInfos(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.ModuleInfo, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadPackageInfos(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.PackageInfo, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadConstantClassNames(reader *ClassFileReader, constants *classfile.ConstantPool) ([]string, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadServiceInfos(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.ServiceInfo, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadAnnotations(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.Annotation, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}

func (d *ClassFileDeserializer) LoadParameterAnnotations(reader *ClassFileReader, constants *classfile.ConstantPool) ([]attribute.Annotations, error) {
	count := reader.ReadUnsignedShort()
	if count == 0 {
		return nil, nil
	}

	return nil, nil
}
