package utils

import (
	"bitbucket.org/coontec/javaClass/class/api"
	"bitbucket.org/coontec/javaClass/class/model/classfile"
	"bitbucket.org/coontec/javaClass/class/model/classfile/attribute"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/class/service/deserializer"
	"errors"
	"fmt"
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

	typeTypes.ThisType = m.makeFromInternalTypeName(internalTypeName)
	attributeSignature := classFile.Attribute()["Signature"].(*attribute.AttributeSignature)

	if attributeSignature == nil {
		superTypeName := classFile.SuperTypeName()
		interfaceTypeNames := classFile.InterfaceTypeNames()

		if superTypeName != "class/lang/Object" {
			typeTypes.SuperType = m.makeFromInternalTypeName(superTypeName)
		}

		if interfaceTypeNames != nil {
			length := len(interfaceTypeNames)

			if length == 1 {
				typeTypes.Interfaces = m.makeFromInternalTypeName(interfaceTypeNames[0]).(_type.IType)
			} else {
				list := _type.NewUnmodifiableTypes()
				for _, interfaceTypeName := range interfaceTypeNames {
					list.Add(m.makeFromInternalTypeName(interfaceTypeName).(_type.IType))
				}
				typeTypes.Interfaces = list
			}
		}
	} else {
		reader := NewSignatureReader(attributeSignature.Signature())

		typeTypes.TypeParameters = m.parseTypeParameters(reader)
		typeTypes.SuperType = m.parseClassTypeSignature(reader, 0)

	}

	return typeTypes
}

func (m *TypeMaker) parseTypeArguments(reader *SignatureReader) _type.ITypeArgument {
	firstTypeArgument := m.parseTypeArgument(reader)
	if firstTypeArgument == nil {
		return nil
	}

	nextTypeArgument := m.parseTypeArgument(reader)

	if nextTypeArgument == nil {
		return firstTypeArgument
	} else {
		typeArguments := _type.NewTypeArguments()
		typeArguments.Add(firstTypeArgument)

		for nextTypeArgument != nil {
			typeArguments.Add(nextTypeArgument)
			nextTypeArgument = m.parseTypeArgument(reader)
		}

		return typeArguments
	}
}

func (m *TypeMaker) parseTypeArgument(reader *SignatureReader) _type.ITypeArgument {
	switch reader.Read() {
	case '+':
		return _type.NewWildcardExtendsTypeArgument(m.parseReferenceTypeSignature(reader))
	case '-':
		return _type.NewWildcardSuperTypeArgument(m.parseReferenceTypeSignature(reader))
	case '*':
		return _type.WildcardTypeArgumentEmpty
	default:
		reader.index--
		return m.parseReferenceTypeSignature(reader)
	}
}

func (m *TypeMaker) parseReferenceTypeSignature(reader *SignatureReader) _type.IType {
	if reader.Available() {
		dimension := 0
		c := reader.Read()

		for c == '[' {
			dimension++
			c = reader.Read()
		}

		switch c {
		case 'B':
			if dimension == 0 {
				return _type.PtTypeByte
			}
			return _type.PtTypeByte.CreateType(dimension)
		case 'C':
			if dimension == 0 {
				return _type.PtTypeChar
			}
			return _type.PtTypeChar.CreateType(dimension)
		case 'D':
			if dimension == 0 {
				return _type.PtTypeDouble
			}
			return _type.PtTypeDouble.CreateType(dimension)
		case 'F':
			if dimension == 0 {
				return _type.PtTypeFloat
			}
			return _type.PtTypeFloat.CreateType(dimension)
		case 'I':
			if dimension == 0 {
				return _type.PtTypeInt
			}
			return _type.PtTypeInt.CreateType(dimension)
		case 'J':
			if dimension == 0 {
				return _type.PtTypeLong
			}
			return _type.PtTypeLong.CreateType(dimension)
		case 'L':
			reader.index--
			return m.parseClassTypeSignature(reader, dimension).(_type.IType)
		case 'S':
			if dimension == 0 {
				return _type.PtTypeShort
			}
			return _type.PtTypeShort.CreateType(dimension)
		case 'T':
			index := reader.index

			if reader.Search(';') == false {
				return nil
			}

			identifier := reader.Substring(index)
			reader.index++

			return _type.NewGenericType(identifier, dimension)
		case 'V':
			if dimension == 0 {
				return _type.PtTypeVoid
			}
			return _type.PtTypeVoid.CreateType(dimension)
		case 'Z':
			if dimension == 0 {
				return _type.PtTypeBoolean
			}
			return _type.PtTypeBoolean.CreateType(dimension)
		default:
			reader.index--
			return nil
		}
	}
	return nil
}

func (m *TypeMaker) parseTypeParameters(reader *SignatureReader) _type.ITypeParameter {
	if reader.NextEqualsTo('<') {
		reader.index++

		firstTypeParameter := m.parseTypeParameter(reader)

		if firstTypeParameter == nil {
			return nil
		}

		nextTypeParameter := m.parseTypeParameter(reader)
		var typeParameters _type.ITypeParameter

		if nextTypeParameter == nil {
			typeParameters = firstTypeParameter
		} else {
			list := _type.NewTypeParameters()
			list.Add(firstTypeParameter)

			for nextTypeParameter != nil {
				list.Add(nextTypeParameter)
				nextTypeParameter = m.parseTypeParameter(reader)
			}

			typeParameters = list
		}

		if reader.Read() != '>' {
			return nil
		}

		return typeParameters
	}

	return nil
}

func (m *TypeMaker) parseTypeParameter(reader *SignatureReader) _type.ITypeParameter {
	firstIndex := reader.index

	if reader.Search(':') {
		identifier := reader.Substring(firstIndex)
		var firstBound _type.IType
		var types *_type.UnmodifiableTypes

		for reader.NextEqualsTo(':') {
			reader.index++

			bound := m.parseReferenceTypeSignature(reader)

			if bound != nil && !(bound.Descriptor() == "Ljava/lang/Object;") {
				if firstBound == nil {
					firstBound = bound
				} else if types == nil {
					types = _type.NewUnmodifiableTypes()
					types.Add(firstBound)
					types.Add(bound)
				} else {
					types.Add(bound)
				}
			}
		}

		if firstBound == nil {
			return _type.NewTypeParameter(identifier)
		} else if types == nil {
			return _type.NewTypeParameterWithTypeBounds(identifier, firstBound)
		} else {
			return _type.NewTypeParameterWithTypeBounds(identifier, types)
		}
	}

	return nil
}

func (m *TypeMaker) parseClassTypeSignature(reader *SignatureReader, dimension int) _type.IObjectType {
	if reader.NextEqualsTo('L') {
		reader.index++
		index := reader.index
		endMarker := reader.SearchEndMarker()

		if endMarker == 0 {
			return nil
		}

		internalTypeName := reader.Substring(index)
		ot := m.makeFromInternalTypeName(internalTypeName)

		if endMarker == '<' {
			reader.index++
			ot = ot.CreateTypeWithArgs(m.parseTypeArguments(reader))
			if reader.Read() != '>' {
				return nil
			}
		}

		for reader.NextEqualsTo('.') {
			reader.index++
			index = reader.index
			endMarker = reader.SearchEndMarker()

			if endMarker == 0 {
				return nil
			}

			name := reader.Substring(index)
			internalTypeName += "$" + name
			var qualitifedName string

			if unicode.IsDigit(rune(name[0])) {
				name = extractLocalClassName(name)
				qualitifedName = ""
			} else {
				qualitifedName = ot.QualifiedName() + "." + name
			}

			if endMarker == '<' {
				reader.index++

				typeArguments := m.parseTypeArguments(reader)
				if reader.Read() != '>' {
					return nil
				}

				ot = _type.NewInnerObjectTypeWithArgs(internalTypeName, qualitifedName, name, typeArguments, ot)
			} else {
				ot = _type.NewInnerObjectType(internalTypeName, qualitifedName, name, ot)
			}
		}

		reader.index++

		if dimension == 0 {
			return ot
		}

		return ot.CreateType(dimension).(_type.IObjectType)
	}

	return nil
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

func NewSignatureReader(signature string) *SignatureReader {
	return NewSignatureReaderWithAll(signature, 0)
}

func NewSignatureReaderWithAll(signature string, index int) *SignatureReader {
	array := []byte(signature)
	return &SignatureReader{
		signature: signature,
		array:     array,
		length:    len(array),
		index:     index,
	}
}

type SignatureReader struct {
	signature string
	array     []byte
	length    int
	index     int
}

func (r *SignatureReader) Read() byte {
	ret := r.array[r.index]
	r.index++
	return ret
}

func (r *SignatureReader) NextEqualsTo(c byte) bool {
	return r.index < r.length && r.array[r.index] == c
}

func (r *SignatureReader) Search(c byte) bool {
	length := r.length
	for i := r.index; i < length; i++ {
		if r.array[i] == c {
			r.index = i
			return true
		}
	}
	return false
}

func (r *SignatureReader) SearchEndMarker() byte {
	length := r.length

	for r.index < length {
		c := r.array[r.index]

		if c == byte(';') || c == byte('<') || c == byte('.') {
			return c
		}
		r.index++
	}
	return 0
}

func (r *SignatureReader) Available() bool {
	return r.index < r.length
}

func (r *SignatureReader) Substring(beginIndex int) string {
	return string(r.array[beginIndex : r.index-beginIndex])
}

func (r *SignatureReader) String() string {
	return fmt.Sprintf("SignatureReader{index=%d, nextChars=%s", r.index, string(r.array[r.index:r.length-r.index]))
}

type TypeTypes struct {
	ThisType       _type.IObjectType
	TypeParameters _type.ITypeParameter
	SuperType      _type.IObjectType
	Interfaces     _type.IType
}

type MethodTypes struct {
	TypeParameters _type.ITypeParameter
	ParameterTypes _type.IType
	ReturnedType   _type.IType
	ExceptionTypes _type.IType
}
