package utils

import (
	"bitbucket.org/coontec/javaClass/java/api"
	"bitbucket.org/coontec/javaClass/java/model/classfile"
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/java/service/deserializer"
	"errors"
	"os"
	"strings"
	"unicode"
)

var InternalNameToObjectPrimitiveType = map[string]_type.IObjectType{
	_type.OtTypePrimitiveBoolean.InternalName(): _type.OtTypePrimitiveBoolean,
	_type.OtTypePrimitiveByte.InternalName():    _type.OtTypePrimitiveByte,
	_type.OtTypePrimitiveChar.InternalName():    _type.OtTypePrimitiveChar,
	_type.OtTypePrimitiveDouble.InternalName():  _type.OtTypePrimitiveDouble,
	_type.OtTypePrimitiveFloat.InternalName():   _type.OtTypePrimitiveFloat,
	_type.OtTypePrimitiveInt.InternalName():     _type.OtTypePrimitiveInt,
	_type.OtTypePrimitiveLong.InternalName():    _type.OtTypePrimitiveLong,
	_type.OtTypePrimitiveShort.InternalName():   _type.OtTypePrimitiveShort,
	_type.OtTypePrimitiveVoid.InternalName():    _type.OtTypePrimitiveVoid,
}

func NewTypeMake() *TypeMaker {
	t := &TypeMaker{}

	t.signatureToType["B"] = _type.PtTypeByte
	t.signatureToType["C"] = _type.PtTypeChar
	t.signatureToType["D"] = _type.PtTypeDouble
	t.signatureToType["F"] = _type.PtTypeFloat
	t.signatureToType["I"] = _type.PtTypeInt
	t.signatureToType["J"] = _type.PtTypeLong
	t.signatureToType["S"] = _type.PtTypeShort
	t.signatureToType["V"] = _type.PtTypeVoid
	t.signatureToType["Z"] = _type.PtTypeBoolean

	t.signatureToType["Ljava/lang/Class;"] = _type.OtTypeClass
	t.signatureToType["Ljava/lang/Exception;"] = _type.OtTypeException
	t.signatureToType["Ljava/lang/Object;"] = _type.OtTypeObject
	t.signatureToType["Ljava/lang/Throwable;"] = _type.OtTypeThrowable
	t.signatureToType["Ljava/lang/String;"] = _type.OtTypeString
	t.signatureToType["Ljava/lang/System;"] = _type.OtTypeSystem

	t.descriptorToObjectType["Ljava/lang/Class;"] = _type.OtTypeClass
	t.descriptorToObjectType["Ljava/lang/Exception;"] = _type.OtTypeException
	t.descriptorToObjectType["Ljava/lang/Object;"] = _type.OtTypeObject
	t.descriptorToObjectType["Ljava/lang/Throwable;"] = _type.OtTypeThrowable
	t.descriptorToObjectType["Ljava/lang/String;"] = _type.OtTypeString
	t.descriptorToObjectType["Ljava/lang/System;"] = _type.OtTypeSystem

	t.internalTypeNameToObjectType["Ljava/lang/Class;"] = _type.OtTypeClass
	t.internalTypeNameToObjectType["Ljava/lang/Exception;"] = _type.OtTypeException
	t.internalTypeNameToObjectType["Ljava/lang/Object;"] = _type.OtTypeObject
	t.internalTypeNameToObjectType["Ljava/lang/Throwable;"] = _type.OtTypeThrowable
	t.internalTypeNameToObjectType["Ljava/lang/String;"] = _type.OtTypeString
	t.internalTypeNameToObjectType["Ljava/lang/System;"] = _type.OtTypeSystem

	return t
}

type TypeMaker struct {
	signatureToType                                                  map[string]_type.IType
	internalTypeNameFieldNameToType                                  map[string]_type.IType
	descriptorToObjectType                                           map[string]_type.IObjectType
	internalTypeNameToObjectType                                     map[string]_type.IObjectType
	internalTypeNameToTypeTypes                                      map[string]TypeTypes
	internalTypeNameMethodNameParameterCountToDeclaredParameterTypes map[string][]_type.IType
	internalTypeNameMethodNameParameterCountToParameterTypes         map[string][]_type.IType
	internalTypeNameMethodNameDescriptorToMethodTypes                map[string]MethodTypes
	signatureToMethodTypes                                           map[string]MethodTypes
	assignableRawTypes                                               map[int64]bool
	superParameterizedObjectTypes                                    map[int64]_type.IObjectType
	hierarchy                                                        map[string][]string
	classPathLoader                                                  ClassPathLoader
	loader                                                           api.Loader
}

func (m *TypeMaker) ParseClassFileSignature(classFile classfile.ClassFile) TypeTypes {
	typeTypes := TypeTypes{}
	internalTypeName := classFile.InternalTypeName()

	typeTypes.thisType = m.makeFromInternalTypeName(internalTypeName)
	attributeSignature := classFile.Attribute()["Signature"]

	if attributeSignature == nil {
		superTypeName := classFile.SuperTypeName()
		interfaceTypeNames := classFile.InterfaceTypeNames()

		if superTypeName != "java/lang/Object" {
			typeTypes.superType = m.makeFromInternalTypeName(superTypeName)
		}

		if interfaceTypeNames != nil {
			length := len(interfaceTypeNames)

			if length == 1 {
				typeTypes.interfaces = m.makeFromInternalTypeName(interfaceTypeNames[0])
			} else {
				list := _type.NewUnmodifiableTypes()
				for _, interfaceTypeName := range interfaceTypeNames {
					list.Add(m.makeFromInternalTypeName(interfaceTypeName))
				}
			}
		}
	} else {
	}

	return typeTypes
}

func (m *TypeMaker) makeFromDescriptorWithoutBracket(descriptor string) _type.IObjectType {
	ot := InternalNameToObjectPrimitiveType[descriptor]

	if ot == nil {
		ot = m.makeFromInternalTypeName(descriptor[1 : len(descriptor)-1])
	}

	return ot
}

func (m *TypeMaker) makeFromInternalTypeName(internalTypeName string) _type.IObjectType {
	ot := m.loadType(internalTypeName)

	if ot == nil {
		ot = m.create(internalTypeName)
	}

	return ot
}

func (m *TypeMaker) create(internalTypeName string) _type.IObjectType {
	lastSlash := strings.LastIndex(internalTypeName, "/")
	lastDollar := strings.LastIndex(internalTypeName, "$")

	var ot _type.IObjectType

	if lastSlash < lastDollar {
		outerTypeName := internalTypeName[:lastDollar]
		outerSot := m.create(outerTypeName)
		innerName := internalTypeName[len(outerTypeName)+1:]

		if innerName == "" {
			qualifiedName := strings.ReplaceAll(internalTypeName, "/", ".")
			name := qualifiedName[lastSlash+1:]
			ot = _type.NewObjectType(internalTypeName, qualifiedName, name)
		} else if unicode.IsDigit(rune(innerName[0])) {
			ot = _type.NewInnerObjectType(internalTypeName, "", extractLocalClassName(innerName), outerSot)
		} else {
			qualifiedName := outerSot.QualifiedName() + "." + innerName
			ot = _type.NewInnerObjectType(internalTypeName, qualifiedName, innerName, outerSot)
		}
	} else {
		qualifiedName := strings.ReplaceAll(internalTypeName, "/", ".")
		name := qualifiedName[lastSlash+1:]
		ot = _type.NewObjectType(internalTypeName, qualifiedName, name)
	}

	m.internalTypeNameToObjectType[internalTypeName] = ot

	return ot
}

func (m *TypeMaker) loadType(internalTypeName string) _type.IObjectType {
	ot := m.internalTypeNameToObjectType[internalTypeName]
	if ot == nil {
		if m.loader.CanLoad(internalTypeName) {
			data, _ := m.loader.Load(internalTypeName)
			ot = m.loadType2(internalTypeName, data)
		} else if m.classPathLoader.CanLoad(internalTypeName) {
			data, _ := m.classPathLoader.Load(internalTypeName)
			ot = m.loadType2(internalTypeName, data)
		}
	}
	return ot
}

func (m *TypeMaker) loadType2(internalTypeName string, data []byte) _type.IObjectType {
	if data == nil {
		return nil
	}

	reader := deserializer.NewClassFileReader(data)
	constants, err := m.loadClassFile(internalTypeName, reader)
	if err != nil {
		return nil
	}

	// Skip fields
	m.skipMembers(reader)

	// Skip fields
	m.skipMembers(reader)

	outerTypeName := ""
	var outerObjectType _type.IObjectType

	// Load attributes
	count := reader.ReadUnsignedShort()
	for i := 0; i < count; i++ {
		attributeNameIndex := reader.ReadUnsignedShort()
		attributeLength := reader.ReadInt()

		if str, ok := constants[attributeNameIndex].(string); ok && str == "InnerClasses" {
			innerClassCount := reader.ReadUnsignedShort()

			for j := 0; j < innerClassCount; j++ {
				innerTypeIndex := reader.ReadUnsignedShort()
				outerTypeIndex := reader.ReadUnsignedShort()

				reader.Skip(2 * 2)

				cc := constants[innerTypeIndex].(int)
				innerTypeName := constants[cc].(string)

				if innerTypeName == internalTypeName {
					if outerTypeIndex == 0 {
						lastDollar := strings.LastIndex(internalTypeName, "$")
						if lastDollar != -1 {
							outerTypeName = internalTypeName[:lastDollar]
							outerObjectType = m.loadType(outerTypeName)
						}
					} else {
						cc = constants[outerTypeIndex].(int)
						outerTypeName = constants[cc].(string)
						outerObjectType = m.loadType(outerTypeName)
					}
					break
				}
			}
			break
		} else {
			reader.Skip(attributeLength)
		}
	}

	if outerObjectType == nil {
		lastSlash := strings.LastIndex(internalTypeName, "/")
		qualifiedName := strings.ReplaceAll(internalTypeName, "/", ".")
		name := internalTypeName[lastSlash:]

		return _type.NewObjectType(internalTypeName, qualifiedName, name)
	} else {
		var index int

		if len(internalTypeName) > len(outerTypeName)+1 {
			index = len(outerTypeName)
		} else {
			index = strings.LastIndex(internalTypeName, "$")
		}

		innerName := internalTypeName[index:]

		if unicode.IsDigit(rune(innerName[0])) {
			return _type.NewInnerObjectType(internalTypeName, "", extractLocalClassName(innerName), outerObjectType)
		} else {
			qualifiedName := outerObjectType.QualifiedName() + "." + innerName
			return _type.NewInnerObjectType(internalTypeName, qualifiedName, innerName, outerObjectType)
		}
	}
}

func extractLocalClassName(name string) string {
	if unicode.IsDigit(rune(name[0])) {
		i := 0
		length := len(name)

		for i < length && unicode.IsDigit(rune(name[i])) {
			i++
		}

		if i == length {
			return ""
		}

		return name[i:]
	}

	return name
}

func (m *TypeMaker) skipMembers(reader *deserializer.ClassFileReader) {
	count := reader.ReadUnsignedShort()
	for i := 0; i < count; i++ {
		reader.Skip(3 * 2)
		m.skipAttributes(reader)
	}
}

func (m *TypeMaker) skipAttributes(reader *deserializer.ClassFileReader) {
	count := reader.ReadUnsignedShort()
	for i := 0; i < count; i++ {
		reader.Skip(2)
		attributeLength := reader.ReadInt()
		reader.Skip(attributeLength)
	}
}

func (m *TypeMaker) loadClassFile(internalTypeName string, reader *deserializer.ClassFileReader) ([]interface{}, error) {
	magic := reader.ReadInt()

	if magic != deserializer.JavaMagicNumber {
		return nil, errors.New("invalid CLASS file")
	}

	// Skip 'minorVersion', 'majorVersion'
	reader.Skip(2 * 2)

	constants, err := m.loadConstants(reader)
	if err != nil {
		return nil, err
	}

	// Skip 'accessFlags' & 'thisClassIndex'
	reader.Skip(2 * 2)

	superClassIndex := reader.ReadUnsignedShort()
	var superClassName string

	if superClassIndex == 0 {
		superClassName = ""
	} else {
		if cc, ok := constants[superClassIndex].(int); ok {
			superClassName = constants[cc].(string)
		}
	}

	count := reader.ReadUnsignedShort()
	superClassAndInterfaceNames := make([]string, 0, count+1)
	superClassAndInterfaceNames[0] = superClassName

	for i := 1; i <= count; i++ {
		interfaceIndex := reader.ReadUnsignedShort()
		if cc, ok := constants[interfaceIndex].(int); ok {
			superClassAndInterfaceNames[i] = constants[cc].(string)
		}
	}

	m.hierarchy[internalTypeName] = superClassAndInterfaceNames

	return constants, nil
}

func (m *TypeMaker) loadConstants(reader *deserializer.ClassFileReader) ([]interface{}, error) {
	count := reader.ReadUnsignedShort()

	if count == 0 {
		return nil, nil
	}

	constants := make([]interface{}, count)

	for i := 1; i < count; i++ {
		tag := reader.Read()

		switch tag {
		case 1:
			constants[i] = reader.ReadUTF8()
		case 7:
			constants[i] = int(reader.ReadUnsignedShort())
		case 8, 16, 19, 20:
			reader.Skip(2)
		case 15:
			reader.Skip(3)
		case 3, 4, 9, 10, 11, 12, 17, 18:
			reader.Skip(4)
		case 5, 6:
			reader.Skip(8)
			i++
		default:
			return nil, errors.New("Invalid constant pool entry")
		}
	}

	return constants, nil
}

type ClassPathLoader struct {
	buffer []byte
}

func (l *ClassPathLoader) Load(internalName string) ([]byte, error) {
	data, err := os.ReadFile(internalName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *ClassPathLoader) CanLoad(internalName string) bool {
	if _, err := os.Stat(internalName); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

type TypeTypes struct {
	thisType       _type.IObjectType
	typeParameters _type.TypeParameter
	superType      _type.IObjectType
	interfaces     _type.IType
}

type MethodTypes struct {
	typeParameters _type.TypeParameter
	parameterTypes _type.IType
	returnedType   _type.IType
	exceptionTypes _type.IType
}
