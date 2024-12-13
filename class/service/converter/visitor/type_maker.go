package visitor

import (
	"errors"
	"fmt"
	"github.com/ElectricSaw/go-jd-core/class/api"
	intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/classfile/attribute"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	"github.com/ElectricSaw/go-jd-core/class/service/deserializer"
	"hash/fnv"
	"os"
	"strings"
	"unicode"
)

var InternalNameToObjectPrimitiveType = map[string]intmod.IObjectType{
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

func NewTypeMaker(loader api.Loader) intsrv.ITypeMaker {
	t := &TypeMaker{
		signatureToType:                 make(map[string]intmod.IType),
		internalTypeNameFieldNameToType: make(map[string]intmod.IType),
		descriptorToObjectType:          make(map[string]intmod.IObjectType),
		internalTypeNameToObjectType:    make(map[string]intmod.IObjectType),
		internalTypeNameToTypeTypes:     make(map[string]intsrv.ITypeTypes),
		internalTypeNameMethodNameParameterCountToDeclaredParameterTypes: make(map[string][]intmod.IType),
		internalTypeNameMethodNameParameterCountToParameterTypes:         make(map[string][]intmod.IType),
		internalTypeNameMethodNameDescriptorToMethodTypes:                make(map[string]intsrv.IMethodTypes),
		signatureToMethodTypes:        make(map[string]intsrv.IMethodTypes),
		assignableRawTypes:            make(map[int64]bool),
		superParameterizedObjectTypes: make(map[int64]intmod.IObjectType),
		hierarchy:                     make(map[string][]string),
		loader:                        loader,
	}

	t.signatureToType["B"] = _type.PtTypeByte.(intmod.IType)
	t.signatureToType["C"] = _type.PtTypeChar.(intmod.IType)
	t.signatureToType["D"] = _type.PtTypeDouble.(intmod.IType)
	t.signatureToType["F"] = _type.PtTypeFloat.(intmod.IType)
	t.signatureToType["I"] = _type.PtTypeInt.(intmod.IType)
	t.signatureToType["J"] = _type.PtTypeLong.(intmod.IType)
	t.signatureToType["S"] = _type.PtTypeShort.(intmod.IType)
	t.signatureToType["V"] = _type.PtTypeVoid.(intmod.IType)
	t.signatureToType["Z"] = _type.PtTypeBoolean.(intmod.IType)

	t.signatureToType["Ljava/lang/Class;"] = _type.OtTypeClass.(intmod.IType)
	t.signatureToType["Ljava/lang/Exception;"] = _type.OtTypeException.(intmod.IType)
	t.signatureToType["Ljava/lang/Object;"] = _type.OtTypeObject.(intmod.IType)
	t.signatureToType["Ljava/lang/Throwable;"] = _type.OtTypeThrowable.(intmod.IType)
	t.signatureToType["Ljava/lang/String;"] = _type.OtTypeString.(intmod.IType)
	t.signatureToType["Ljava/lang/System;"] = _type.OtTypeSystem.(intmod.IType)

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
	signatureToType                                                  map[string]intmod.IType
	internalTypeNameFieldNameToType                                  map[string]intmod.IType
	descriptorToObjectType                                           map[string]intmod.IObjectType
	internalTypeNameToObjectType                                     map[string]intmod.IObjectType
	internalTypeNameToTypeTypes                                      map[string]intsrv.ITypeTypes
	internalTypeNameMethodNameParameterCountToDeclaredParameterTypes map[string][]intmod.IType
	internalTypeNameMethodNameParameterCountToParameterTypes         map[string][]intmod.IType
	internalTypeNameMethodNameDescriptorToMethodTypes                map[string]intsrv.IMethodTypes
	signatureToMethodTypes                                           map[string]intsrv.IMethodTypes
	assignableRawTypes                                               map[int64]bool
	superParameterizedObjectTypes                                    map[int64]intmod.IObjectType
	hierarchy                                                        map[string][]string
	classPathLoader                                                  intsrv.IClassPathLoader
	loader                                                           api.Loader
}

func (m *TypeMaker) ParseClassFileSignature(classFile intcls.IClassFile) intsrv.ITypeTypes {
	typeTypes := NewTypeTypes()
	internalTypeName := classFile.InternalTypeName()

	typeTypes.SetThisType(m.MakeFromInternalTypeName(internalTypeName))
	attributeSignature := classFile.Attributes()["Signature"].(*attribute.AttributeSignature)

	if attributeSignature == nil {
		superTypeName := classFile.SuperTypeName()
		interfaceTypeNames := classFile.InterfaceTypeNames()

		if superTypeName != "class/lang/Object" {
			typeTypes.SetSuperType(m.MakeFromInternalTypeName(superTypeName))
		}

		if interfaceTypeNames != nil {
			length := len(interfaceTypeNames)

			if length == 1 {
				typeTypes.SetInterfaces(m.MakeFromInternalTypeName(interfaceTypeNames[0]).(intmod.IType))
			} else {
				list := _type.NewUnmodifiableTypes()
				for _, interfaceTypeName := range interfaceTypeNames {
					list.Add(m.MakeFromInternalTypeName(interfaceTypeName).(intmod.IType))
				}

				if cast, ok := list.(intmod.IType); ok {
					typeTypes.SetInterfaces(cast)
				}
			}
		}
	} else {
		reader := NewSignatureReader(attributeSignature.Signature())

		typeTypes.SetTypeParameters(m.parseTypeParameters(reader))
		typeTypes.SetSuperType(m.parseClassTypeSignature(reader, 0))

	}

	return typeTypes
}

func (m *TypeMaker) ParseMethodSignature(classFile intcls.IClassFile, method intcls.IMethod) intsrv.IMethodTypes {
	key := classFile.InternalTypeName() + ":" + method.Name() + method.Descriptor()
	return m.parseMethodSignature(method, key)
}

func (m *TypeMaker) parseMethodSignature(method intcls.IMethod, key string) intsrv.IMethodTypes {
	attributeSignature := method.Attributes()["Signature"].(*attribute.AttributeSignature)
	exceptionTypeNames := getExceptionTypeNames(method)
	var methodTypes intsrv.IMethodTypes

	if attributeSignature == nil {
		methodTypes = m.parseMethodSignature2(method.Descriptor(), exceptionTypeNames)
	} else {
		methodTypes = m.parseMethodSignature3(method.Descriptor(), attributeSignature.Signature(), exceptionTypeNames)
	}

	m.internalTypeNameMethodNameDescriptorToMethodTypes[key] = methodTypes

	return methodTypes
}

func (m *TypeMaker) ParseFieldSignature(classFile intcls.IClassFile, field intcls.IField) intmod.IType {
	key := classFile.InternalTypeName() + ":" + field.Name()
	attributeSignature := field.Attributes()["Signature"].(*attribute.AttributeSignature)
	signature := ""

	if attributeSignature == nil {
		signature = field.Descriptor()
	} else {
		signature = attributeSignature.Signature()
	}

	typ := m.MakeFromSignature(signature)

	m.internalTypeNameFieldNameToType[key] = typ

	return typ
}

func (m *TypeMaker) MakeFromSignature(signature string) intmod.IType {
	typ := m.signatureToType[signature]
	if typ == nil {
		reader := NewSignatureReader(signature)
		typ = m.parseReferenceTypeSignature(reader)
		m.signatureToType[signature] = typ
	}

	return typ
}

func (m *TypeMaker) parseMethodSignature2(signature string, exceptionTypeNames []string) intsrv.IMethodTypes {
	cacheKey := signature
	containsThrowsSignature := strings.Index(signature, "^") != -1

	if !containsThrowsSignature && (exceptionTypeNames != nil) {
		sb := ""

		for _, exceptionTypeName := range exceptionTypeNames {
			sb += "^L" + exceptionTypeName + ";"
		}

		cacheKey = sb
	}

	methodTypes := m.signatureToMethodTypes[cacheKey]

	if methodTypes == nil {
		reader := NewSignatureReader(signature)
		methodTypes = NewMethodTypes2(m.parseTypeParameters(reader),
			nil, nil, nil)

		if reader.Read() != '(' {
			return nil
		}

		firstParameterType := m.parseReferenceTypeSignature(reader)

		if firstParameterType == nil {
			methodTypes.SetParameterTypes(nil)
		} else {
			nextParameterType := m.parseReferenceTypeSignature(reader)
			types := &_type.UnmodifiableTypes{}
			types.Add(firstParameterType)

			for nextParameterType != nil {
				types.Add(nextParameterType)
				nextParameterType = m.parseReferenceTypeSignature(reader)
			}

			methodTypes.SetParameterTypes(types)
		}

		if reader.Read() != ')' {
			return nil
		}

		methodTypes.SetReturnedType(m.parseReferenceTypeSignature(reader))
		firstException := m.parseExceptionSignature(reader)

		if firstException == nil {
			if exceptionTypeNames != nil {
				if len(exceptionTypeNames) == 1 {
					methodTypes.SetExceptionTypes(m.MakeFromInternalTypeName(exceptionTypeNames[0]).(intmod.IType))
				} else {
					list := &_type.UnmodifiableTypes{}

					for _, exceptionTypeName := range exceptionTypeNames {
						list.Add(m.MakeFromInternalTypeName(exceptionTypeName).(intmod.IType))
					}

					methodTypes.SetExceptionTypes(list)
				}
			}
		} else {
			nextException := m.parseExceptionSignature(reader)

			if nextException == nil {
				methodTypes.SetExceptionTypes(firstException)
			} else {
				list := &_type.UnmodifiableTypes{}
				list.Add(firstException)

				for nextException != nil {
					list.Add(nextException)
					nextException = m.parseExceptionSignature(reader)
				}

				methodTypes.SetExceptionTypes(list)
			}
		}

		m.signatureToMethodTypes[cacheKey] = methodTypes
	}

	return methodTypes
}

func (m *TypeMaker) parseMethodSignature3(descriptor, signature string, exceptionTypeNames []string) intsrv.IMethodTypes {
	if signature == "" {
		return m.parseMethodSignature2(descriptor, exceptionTypeNames)
	} else {
		mtDescriptor := m.parseMethodSignature2(descriptor, exceptionTypeNames)
		mtSignature := m.parseMethodSignature2(signature, exceptionTypeNames)

		if mtDescriptor.ParameterTypes() == nil {
			return mtSignature
		} else if mtSignature.ParameterTypes() == nil {
			return NewMethodTypes2(mtSignature.TypeParameters(),
				mtDescriptor.ParameterTypes(),
				mtSignature.ReturnedType(),
				mtSignature.ExceptionTypes())
		} else if mtDescriptor.ParameterTypes().Size() == mtSignature.ParameterTypes().Size() {
			return mtSignature
		} else {
			// TODO: 테스트 필요.
			parameterTypes := _type.NewUnmodifiableTypes()
			parameterTypes.Add(mtSignature.ParameterTypes())

			return NewMethodTypes2(mtSignature.TypeParameters(),
				parameterTypes, mtSignature.ReturnedType(),
				mtSignature.ExceptionTypes())
		}
	}
}

func (m *TypeMaker) parseTypeParameters(reader intsrv.ISignatureReader) intmod.ITypeParameter {
	if reader.NextEqualsTo('<') {
		reader.Inc()

		firstTypeParameter := m.parseTypeParameter(reader)

		if firstTypeParameter == nil {
			return nil
		}

		nextTypeParameter := m.parseTypeParameter(reader)
		var typeParameters intmod.ITypeParameter

		if nextTypeParameter == nil {
			typeParameters = firstTypeParameter
		} else {
			list := _type.NewTypeParameters()
			list.Add(firstTypeParameter)

			for nextTypeParameter != nil {
				list.Add(nextTypeParameter)
				nextTypeParameter = m.parseTypeParameter(reader)
			}

			// TODO
			typeParameters = list.(intmod.ITypeParameter)
		}

		if reader.Read() != '>' {
			return nil
		}

		return typeParameters
	}

	return nil
}

func (m *TypeMaker) parseTypeParameter(reader intsrv.ISignatureReader) intmod.ITypeParameter {
	firstIndex := reader.Index()

	if reader.Search(':') {
		identifier := reader.Substring(firstIndex)
		var firstBound intmod.IType
		var types intmod.IUnmodifiableTypes

		for reader.NextEqualsTo(':') {
			reader.Inc()

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
			return _type.NewTypeParameterWithTypeBounds(identifier, firstBound).(intmod.ITypeParameter)
		} else {
			return _type.NewTypeParameterWithTypeBounds(identifier, types.(intmod.IType)).(intmod.ITypeParameter)
		}
	}

	return nil
}

func (m *TypeMaker) parseExceptionSignature(reader intsrv.ISignatureReader) intmod.IType {
	if reader.NextEqualsTo('^') {
		reader.Inc()
		return m.parseReferenceTypeSignature(reader)
	}
	return nil
}

func (m *TypeMaker) parseClassTypeSignature(reader intsrv.ISignatureReader, dimension int) intmod.IObjectType {
	if reader.NextEqualsTo('L') {
		reader.Inc()
		index := reader.Index()
		endMarker := reader.SearchEndMarker()

		if endMarker == 0 {
			return nil
		}

		internalTypeName := reader.Substring(index)
		ot := m.MakeFromInternalTypeName(internalTypeName)

		if endMarker == '<' {
			reader.Inc()
			ot = ot.CreateTypeWithArgs(m.parseTypeArguments(reader))
			if reader.Read() != '>' {
				return nil
			}
		}

		for reader.NextEqualsTo('.') {
			reader.Inc()
			index = reader.Index()
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
				reader.Inc()

				typeArguments := m.parseTypeArguments(reader)
				if reader.Read() != '>' {
					return nil
				}

				ot = _type.NewInnerObjectTypeWithArgs(internalTypeName, qualitifedName, name, typeArguments, ot).(intmod.IObjectType)
			} else {
				ot = _type.NewInnerObjectType(internalTypeName, qualitifedName, name, ot).(intmod.IObjectType)
			}
		}

		reader.Inc()

		if dimension == 0 {
			return ot
		}

		return ot.CreateType(dimension).(intmod.IObjectType)
	}

	return nil
}

func (m *TypeMaker) parseTypeArguments(reader intsrv.ISignatureReader) intmod.ITypeArgument {
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

func (m *TypeMaker) parseReferenceTypeSignature(reader intsrv.ISignatureReader) intmod.IType {
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
				return _type.PtTypeByte.(intmod.IType)
			}
			return _type.PtTypeByte.CreateType(dimension)
		case 'C':
			if dimension == 0 {
				return _type.PtTypeChar.(intmod.IType)
			}
			return _type.PtTypeChar.CreateType(dimension)
		case 'D':
			if dimension == 0 {
				return _type.PtTypeDouble.(intmod.IType)
			}
			return _type.PtTypeDouble.CreateType(dimension)
		case 'F':
			if dimension == 0 {
				return _type.PtTypeFloat.(intmod.IType)
			}
			return _type.PtTypeFloat.CreateType(dimension)
		case 'I':
			if dimension == 0 {
				return _type.PtTypeInt.(intmod.IType)
			}
			return _type.PtTypeInt.CreateType(dimension)
		case 'J':
			if dimension == 0 {
				return _type.PtTypeLong.(intmod.IType)
			}
			return _type.PtTypeLong.CreateType(dimension)
		case 'L':
			reader.Dec()
			return m.parseClassTypeSignature(reader, dimension).(intmod.IType)
		case 'S':
			if dimension == 0 {
				return _type.PtTypeShort.(intmod.IType)
			}
			return _type.PtTypeShort.CreateType(dimension)
		case 'T':
			index := reader.Index()

			if reader.Search(';') == false {
				return nil
			}

			identifier := reader.Substring(index)
			reader.Inc()

			return _type.NewGenericTypeWithAll(identifier, dimension).(intmod.IType)
		case 'V':
			if dimension == 0 {
				return _type.PtTypeVoid.(intmod.IType)
			}
			return _type.PtTypeVoid.CreateType(dimension)
		case 'Z':
			if dimension == 0 {
				return _type.PtTypeBoolean.(intmod.IType)
			}
			return _type.PtTypeBoolean.CreateType(dimension)
		default:
			reader.Dec()
			return nil
		}
	}
	return nil
}

func (m *TypeMaker) parseTypeArgument(reader intsrv.ISignatureReader) intmod.ITypeArgument {
	switch reader.Read() {
	case '+':
		return _type.NewWildcardExtendsTypeArgument(m.parseReferenceTypeSignature(reader)).(intmod.ITypeArgument)
	case '-':
		return _type.NewWildcardSuperTypeArgument(m.parseReferenceTypeSignature(reader)).(intmod.ITypeArgument)
	case '*':
		return _type.WildcardTypeArgumentEmpty.(intmod.ITypeArgument)
	default:
		reader.Dec()
		return m.parseReferenceTypeSignature(reader)
	}
}

func (m *TypeMaker) MakeFromDescriptorOrInternalTypeName(descriptorOrInternalTypeName string) intmod.IObjectType {
	if descriptorOrInternalTypeName[0] == '[' {
		return m.MakeFromDescriptor(descriptorOrInternalTypeName)
	}

	return m.MakeFromInternalTypeName(descriptorOrInternalTypeName)
}

func (m *TypeMaker) MakeFromDescriptor(descriptor string) intmod.IObjectType {
	ot := m.descriptorToObjectType[descriptor]

	if ot == nil {
		if descriptor[0] == '[' {
			dimension := 1
			for descriptor[dimension] == '[' {
				dimension++
			}
			ot = m.makeFromDescriptorWithoutBracket(descriptor[dimension:]).CreateType(dimension).(intmod.IObjectType)
		} else {
			ot = m.makeFromDescriptorWithoutBracket(descriptor)
		}

		m.descriptorToObjectType[descriptor] = ot
	}

	return ot
}

func (m *TypeMaker) makeFromDescriptorWithoutBracket(descriptor string) intmod.IObjectType {
	ot := InternalNameToObjectPrimitiveType[descriptor]

	if ot == nil {
		ot = m.MakeFromInternalTypeName(descriptor[1 : len(descriptor)-1])
	}

	return ot
}

func (m *TypeMaker) MakeFromInternalTypeName(internalTypeName string) intmod.IObjectType {
	ot := m.loadType(internalTypeName)

	if ot == nil {
		ot = m.create(internalTypeName)
	}

	return ot
}

func (m *TypeMaker) create(internalTypeName string) intmod.IObjectType {
	lastSlash := strings.LastIndex(internalTypeName, "/")
	lastDollar := strings.LastIndex(internalTypeName, "$")

	var ot intmod.IObjectType

	if lastSlash < lastDollar {
		outerTypeName := internalTypeName[:lastDollar]
		outerSot := m.create(outerTypeName)
		innerName := internalTypeName[len(outerTypeName)+1:]

		if innerName == "" {
			qualifiedName := strings.ReplaceAll(internalTypeName, "/", ".")
			name := qualifiedName[lastSlash+1:]
			ot = _type.NewObjectType(internalTypeName, qualifiedName, name)
		} else if unicode.IsDigit(rune(innerName[0])) {
			ot = _type.NewInnerObjectType(internalTypeName, "",
				extractLocalClassName(innerName), outerSot).(intmod.IObjectType)
		} else {
			qualifiedName := outerSot.QualifiedName() + "." + innerName
			ot = _type.NewInnerObjectType(internalTypeName, qualifiedName,
				innerName, outerSot).(intmod.IObjectType)
		}
	} else {
		qualifiedName := strings.ReplaceAll(internalTypeName, "/", ".")
		name := qualifiedName[lastSlash+1:]
		ot = _type.NewObjectType(internalTypeName, qualifiedName, name)
	}

	m.internalTypeNameToObjectType[internalTypeName] = ot

	return ot
}

func (m *TypeMaker) SearchSuperParameterizedType(superObjectType, objectType intmod.IObjectType) intmod.IObjectType {
	if superObjectType == _type.OtTypeUndefinedObject || superObjectType == _type.OtTypeObject || superObjectType == objectType {
		return objectType
	} else if superObjectType.Dimension() > 0 || objectType.Dimension() > 0 {
		return nil
	} else {
		superInternalTypeName := superObjectType.InternalName()
		superHashCode := hashCodeWithString(superInternalTypeName) * 31
		return m.searchSuperParameterizedType(superHashCode, superInternalTypeName, objectType)
	}
}

func (m *TypeMaker) IsAssignable(typeBounds map[string]intmod.IType, left, right intmod.IObjectType) bool {
	if left == _type.OtTypeUndefinedObject || right == _type.OtTypeUndefinedObject || left == _type.OtTypeObject || left == right {
		return true
	} else if left.Dimension() > 0 || right.Dimension() > 0 {
		return false
	} else {
		leftInternalTypeName := left.InternalName()
		leftHashCode := hashCodeWithString(leftInternalTypeName) * 31
		ot := m.searchSuperParameterizedType(leftHashCode, leftInternalTypeName, right)

		if ot != nil && leftInternalTypeName == ot.InternalName() {
			if left.TypeArguments() == nil || ot.TypeArguments() == nil {
				return true
			} else {
				return left.TypeArguments().IsTypeArgumentAssignableFrom(typeBounds, ot.TypeArguments())
			}
		}
	}

	return false
}

func (m *TypeMaker) searchSuperParameterizedType(leftHashCode int, leftInternalTypeName string, right intmod.IObjectType) intmod.IObjectType {
	if right == _type.OtTypeObject {
		return nil
	}

	key := int64(leftHashCode + right.HashCode())

	if v, ok := m.superParameterizedObjectTypes[key]; ok {
		return v
	}

	rightInternalTypeName := right.InternalName()

	if leftInternalTypeName == rightInternalTypeName {
		m.superParameterizedObjectTypes[key] = right
		return right
	}

	rightTypeTypes := m.MakeTypeTypes(rightInternalTypeName)

	if rightTypeTypes != nil {
		bindTypesToTypesVisitor := NewBindTypesToTypesVisitor()
		bindings := make(map[string]intmod.ITypeArgument)

		if (rightTypeTypes.TypeParameters() == nil) || (right.TypeArguments() == nil) {
			bindings = make(map[string]intmod.ITypeArgument)
		} else {
			bindings = make(map[string]intmod.ITypeArgument)

			if rightTypeTypes.TypeParameters().IsList() && right.TypeArguments().IsTypeArgumentList() {
				iteratorTypeParameter := rightTypeTypes.TypeParameters().ToSlice()
				iteratorTypeArgument := right.TypeArguments().TypeArgumentList()

				for i := 0; i < len(iteratorTypeParameter); i++ {
					bindings[iteratorTypeParameter[i].Identifier()] = iteratorTypeArgument[i]
				}
			} else {
				bindings[rightTypeTypes.TypeParameters().First().Identifier()] = right.TypeArguments().TypeArgumentFirst()
			}
		}

		bindTypesToTypesVisitor.SetBindings(bindings)

		if rightTypeTypes.SuperType() != nil {
			bindTypesToTypesVisitor.Init()
			rightTypeTypes.SuperType().AcceptTypeVisitor(bindTypesToTypesVisitor)
			ot := bindTypesToTypesVisitor.Type().(intmod.IObjectType)
			ot = m.searchSuperParameterizedType(leftHashCode, leftInternalTypeName, ot)

			if ot != nil {
				m.superParameterizedObjectTypes[key] = ot
				return ot
			}
		}

		if rightTypeTypes.Interfaces() != nil {
			for _, interfaze := range rightTypeTypes.Interfaces().ToSlice() {
				bindTypesToTypesVisitor.Init()
				interfaze.AcceptTypeVisitor(bindTypesToTypesVisitor)
				ot := bindTypesToTypesVisitor.Type().(intmod.IObjectType)
				ot = m.searchSuperParameterizedType(leftHashCode, leftInternalTypeName, ot)

				if ot != nil {
					m.superParameterizedObjectTypes[key] = ot
					return ot
				}
			}
		}
	}

	m.superParameterizedObjectTypes[key] = nil
	return nil
}

func (m *TypeMaker) IsRawTypeAssignable(left, right intmod.IObjectType) bool {
	if left == _type.OtTypeUndefinedObject || left == _type.OtTypeObject || left == right {
		return true
	} else if left.Dimension() > 0 || right.Dimension() > 0 {
		return false
	} else {
		leftInternalName := left.InternalName()
		rightInternalName := right.InternalName()

		if leftInternalName == rightInternalName {
			return true
		} else {
			return m.isRawTypeAssignable(hashCodeWithString(leftInternalName)*31, leftInternalName, rightInternalName)
		}
	}
}

func (m *TypeMaker) isRawTypeAssignable(leftHashCode int, leftInternalName, rightInternalName string) bool {
	if rightInternalName == "java/lang/Object" {
		return false
	}

	key := int64(leftHashCode + hashCodeWithString(rightInternalName))

	if value, ok := m.assignableRawTypes[key]; ok {
		return value
	}

	superClassAndInterfaceNames := m.hierarchy[rightInternalName]

	if superClassAndInterfaceNames == nil {
		m.loadType(rightInternalName)
		superClassAndInterfaceNames = m.hierarchy[rightInternalName]
	}

	if superClassAndInterfaceNames != nil {
		for _, name := range superClassAndInterfaceNames {
			if leftInternalName == name {
				m.assignableRawTypes[key] = true
				return true
			}
		}

		for _, name := range superClassAndInterfaceNames {
			if m.isRawTypeAssignable(leftHashCode, leftInternalName, name) {
				m.assignableRawTypes[key] = true
				return true
			}
		}
	}

	m.assignableRawTypes[key] = false
	return false
}

func (m *TypeMaker) MakeTypeTypes(internalTypeName string) intsrv.ITypeTypes {
	if value, ok := m.internalTypeNameToTypeTypes[internalTypeName]; ok {
		return value
	}

	var typeTypes intsrv.ITypeTypes

	if m.loader.CanLoad(internalTypeName) {
		data, err := m.loader.Load(internalTypeName)
		if err != nil {
			return nil
		}
		typeTypes = m.makeTypeTypes(internalTypeName, data)
		m.internalTypeNameToTypeTypes[internalTypeName] = typeTypes
	} else if m.classPathLoader.CanLoad(internalTypeName) {
		data, err := m.classPathLoader.Load(internalTypeName)
		if err != nil {
			return nil
		}
		typeTypes = m.makeTypeTypes(internalTypeName, data)
		m.internalTypeNameToTypeTypes[internalTypeName] = typeTypes
	}

	return typeTypes
}

func (m *TypeMaker) makeTypeTypes(internalTypeName string, data []byte) intsrv.ITypeTypes {
	if data == nil {
		return nil
	}

	reader := deserializer.NewClassFileReader(data)
	constants, err := m.loadClassFile(internalTypeName, reader)
	if err != nil {
		return nil
	}

	// Skip fields
	skipMembers(reader)

	// Skip methods
	skipMembers(reader)

	// Load attributes
	var signature string
	count := reader.ReadUnsignedShort()

	for j := 0; j < count; j++ {
		attributeNameIndex := reader.ReadUnsignedShort()
		attributeLength := reader.ReadInt()

		if constants[attributeNameIndex] == "Signature" {
			signature = constants[reader.ReadUnsignedShort()].(string)
			break
		} else {
			reader.Skip(attributeLength)
		}
	}

	superClassAndInterfaceNames := m.hierarchy[internalTypeName]
	typeTypes := &TypeTypes{}

	typeTypes.SetThisType(m.MakeFromInternalTypeName(internalTypeName))

	if signature == "" {
		superTypeName := superClassAndInterfaceNames[0]

		if superTypeName == "" {
			typeTypes.SetSuperType(nil)
		} else {
			typeTypes.SetSuperType(m.MakeFromInternalTypeName(superTypeName))
		}

		switch len(superClassAndInterfaceNames) {
		case 0, 1:
		case 2:
			typeTypes.SetInterfaces(m.MakeFromInternalTypeName(superClassAndInterfaceNames[1]).(intmod.IType))
		default:
			length := len(superClassAndInterfaceNames)
			list := _type.NewUnmodifiableTypes()
			for i := 1; i < length; i++ {
				list.Add(m.MakeFromInternalTypeName(superClassAndInterfaceNames[i]).(intmod.IType))
			}
			typeTypes.SetInterfaces(list)
			break
		}
	} else {
		// Parse 'signature' attribute
		signatureReader := &SignatureReader{signature: signature}

		typeTypes.SetTypeParameters(m.parseTypeParameters(signatureReader))
		typeTypes.SetSuperType(m.parseClassTypeSignature(signatureReader, 0))

		firstInterface := m.parseClassTypeSignature(signatureReader, 0)

		if firstInterface != nil {
			nextInterface := m.parseClassTypeSignature(signatureReader, 0)

			if nextInterface == nil {
				typeTypes.SetInterfaces(firstInterface.(intmod.IType))
			} else {
				// length := len(superClassAndInterfaceNames)
				list := _type.NewUnmodifiableTypes()
				list.Add(firstInterface.(intmod.IType))

				for nextInterface != nil {
					list.Add(nextInterface.(intmod.IType))
					nextInterface = m.parseClassTypeSignature(signatureReader, 0)
				}

				typeTypes.SetInterfaces(list)
			}
		}
	}

	return typeTypes
}

func (m *TypeMaker) SetFieldType(internalTypeName, fieldName string, typ intmod.IType) {
	key := internalTypeName + ":" + fieldName
	m.internalTypeNameFieldNameToType[key] = typ
}

func (m *TypeMaker) MakeFieldType(internalTypeName, fieldName, descriptor string) intmod.IType {
	typ := m.loadFieldType(internalTypeName, fieldName, descriptor)

	if typ == nil {
		key := internalTypeName + ":" + fieldName
		typ = m.MakeFromSignature(descriptor)
		m.internalTypeNameFieldNameToType[key] = typ
	}

	return typ
}

func (m *TypeMaker) loadFieldType(internalTypeName, fieldName, descriptor string) intmod.IType {
	key := internalTypeName + ":" + fieldName
	typ := m.internalTypeNameFieldNameToType[key]

	if typ == nil {
		// Load fields
		if m.loadFieldsAndMethods(internalTypeName) {
			typ = m.internalTypeNameFieldNameToType[key]

			if typ == nil {
				typeTypes := m.MakeTypeTypes(internalTypeName)

				if typeTypes != nil {
					if typeTypes.SuperType() != nil {
						typ = m.loadFieldType2(typeTypes.SuperType(), fieldName, descriptor)
					}

					if (typ == nil) && (typeTypes.Interfaces() != nil) {
						if typeTypes.Interfaces().IsList() {
							for _, interfaze := range typeTypes.Interfaces().ToSlice() {
								typ = m.loadFieldType2(interfaze.(intmod.IObjectType),
									fieldName, descriptor)
								if typ != nil {
									break
								}
							}
						} else {
							typ = m.loadFieldType2(typeTypes.Interfaces().First().(intmod.IObjectType), fieldName, descriptor)
						}
					}
				}
			}

			if typ != nil {
				m.internalTypeNameFieldNameToType[key] = typ
			}
		}
	}

	return typ
}

func (m *TypeMaker) loadFieldType2(objectType intmod.IObjectType, fieldName, descriptor string) intmod.IType {
	internalTypeName := objectType.InternalName()
	typeArguments := objectType.TypeArguments()
	typ := m.loadFieldType(internalTypeName, fieldName, descriptor)

	if (typ != nil) && (typeArguments != nil) {
		typeTypes := m.MakeTypeTypes(internalTypeName)

		if (typeTypes != nil) && (typeTypes.TypeParameters() != nil) {
			bindTypesToTypesVisitor := NewBindTypesToTypesVisitor()
			bindings := make(map[string]intmod.ITypeArgument)

			if typeTypes.TypeParameters().IsList() && typeArguments.IsTypeArgumentList() {
				iteratorTypeParameter := typeTypes.TypeParameters().ToSlice()
				iteratorTypeArgument := typeArguments.TypeArgumentList()

				for i := 0; i < len(iteratorTypeParameter); i++ {
					bindings[iteratorTypeParameter[i].Identifier()] = iteratorTypeArgument[i]
				}
			} else {
				bindings[typeTypes.TypeParameters().First().Identifier()] = typeArguments.TypeArgumentFirst()
			}

			bindTypesToTypesVisitor.SetBindings(bindings)
			bindTypesToTypesVisitor.Init()

			typ.AcceptTypeVisitor(bindTypesToTypesVisitor)
			typ = bindTypesToTypesVisitor.Type()
		}
	}

	return typ
}

func (m *TypeMaker) SetMethodReturnedType(internalTypeName, methodName, descriptor string, typ intmod.IType) {
	m.MakeMethodTypes2(internalTypeName, methodName, descriptor).SetReturnedType(typ)
}

func (m *TypeMaker) MakeMethodTypes(descriptor string) intsrv.IMethodTypes {
	return m.parseMethodSignature2(descriptor, nil)
}

func (m *TypeMaker) MakeMethodTypes2(internalTypeName, methodName, descriptor string) intsrv.IMethodTypes {
	methodTypes := m.loadMethodTypes(internalTypeName, methodName, descriptor)

	if methodTypes == nil {
		key := internalTypeName + ":" + methodName + descriptor
		methodTypes = m.parseMethodSignature2(descriptor, nil)
		m.internalTypeNameMethodNameDescriptorToMethodTypes[key] = methodTypes
	}

	return methodTypes
}

func (m *TypeMaker) loadMethodTypes(internalTypeName, methodName, descriptor string) intsrv.IMethodTypes {
	key := internalTypeName + ":" + methodName + descriptor
	methodTypes := m.internalTypeNameMethodNameDescriptorToMethodTypes[key]

	if methodTypes == nil {
		// Load method
		if m.loadFieldsAndMethods(internalTypeName) {
			methodTypes = m.internalTypeNameMethodNameDescriptorToMethodTypes[key]

			if methodTypes == nil {
				typeTypes := m.MakeTypeTypes(internalTypeName)

				if typeTypes != nil {
					if typeTypes.SuperType() != nil {
						methodTypes = m.loadMethodTypes2(typeTypes.SuperType(), methodName, descriptor)
					}

					if (methodTypes == nil) && (typeTypes.Interfaces() != nil) {
						if typeTypes.Interfaces().IsList() {
							for _, interfaze := range typeTypes.Interfaces().ToSlice() {
								methodTypes = m.loadMethodTypes2(interfaze.(intmod.IObjectType), methodName, descriptor)
								if methodTypes != nil {
									break
								}
							}
						} else {
							methodTypes = m.loadMethodTypes2(typeTypes.Interfaces().First().(intmod.IObjectType), methodName, descriptor)
						}
					}
				}
			}

			if methodTypes != nil {
				m.internalTypeNameMethodNameDescriptorToMethodTypes[key] = methodTypes
			}
		}
	}

	return methodTypes
}

func (m *TypeMaker) loadMethodTypes2(objectType intmod.IObjectType, methodName, descriptor string) intsrv.IMethodTypes {
	internalTypeName := objectType.InternalName()
	typeArguments := objectType.TypeArguments()
	methodTypes := m.loadMethodTypes(internalTypeName, methodName, descriptor)

	if (methodTypes != nil) && (typeArguments != nil) {
		typeTypes := m.MakeTypeTypes(internalTypeName)

		if (typeTypes != nil) && (typeTypes.TypeParameters() != nil) {
			bindTypesToTypesVisitor := NewBindTypesToTypesVisitor()
			bindings := make(map[string]intmod.ITypeArgument)
			newMethodTypes := &MethodTypes{}

			if typeTypes.TypeParameters().IsList() && typeArguments.IsTypeArgumentList() {
				iteratorTypeParameter := typeTypes.TypeParameters().ToSlice()
				iteratorTypeArgument := typeArguments.TypeArgumentList()

				for i := 0; i < len(iteratorTypeParameter); i++ {
					bindings[iteratorTypeParameter[i].Identifier()] = iteratorTypeArgument[i]
				}
			} else {
				bindings[typeTypes.TypeParameters().First().Identifier()] = typeArguments.TypeArgumentFirst()
			}

			bindTypesToTypesVisitor.SetBindings(bindings)

			if methodTypes.ParameterTypes() == nil {
				newMethodTypes.SetParameterTypes(nil)
			} else {
				bindTypesToTypesVisitor.Init()
				methodTypes.ParameterTypes().AcceptTypeVisitor(bindTypesToTypesVisitor)
				baseType := bindTypesToTypesVisitor.Type()

				if baseType.IsList() && baseType.IsTypes() {
					baseType = _type.NewUnmodifiableTypes(baseType.ToSlice()...)
				}

				newMethodTypes.SetParameterTypes(baseType)
			}

			bindTypesToTypesVisitor.Init()
			methodTypes.ReturnedType().AcceptTypeVisitor(bindTypesToTypesVisitor)
			newMethodTypes.SetReturnedType(bindTypesToTypesVisitor.Type())

			newMethodTypes.SetTypeParameters(nil)
			newMethodTypes.SetExceptionTypes(methodTypes.ExceptionTypes())

			methodTypes = newMethodTypes
		}
	}

	return methodTypes
}

func (m *TypeMaker) loadType(internalTypeName string) intmod.IObjectType {
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

func (m *TypeMaker) loadType2(internalTypeName string, data []byte) intmod.IObjectType {
	if data == nil {
		return nil
	}

	reader := deserializer.NewClassFileReader(data)
	constants, err := m.loadClassFile(internalTypeName, reader)
	if err != nil {
		return nil
	}

	// Skip fields
	skipMembers(reader)

	// Skip fields
	skipMembers(reader)

	outerTypeName := ""
	var outerObjectType intmod.IObjectType

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
			return _type.NewInnerObjectType(internalTypeName, "",
				extractLocalClassName(innerName), outerObjectType).(intmod.IObjectType)
		} else {
			qualifiedName := outerObjectType.QualifiedName() + "." + innerName
			return _type.NewInnerObjectType(internalTypeName, qualifiedName,
				innerName, outerObjectType).(intmod.IObjectType)
		}
	}
}

func (m *TypeMaker) loadFieldsAndMethods(internalTypeName string) bool {
	if m.loader.CanLoad(internalTypeName) {
		data, err := m.loader.Load(internalTypeName)
		if err != nil {
			return false
		}
		m.loadFieldsAndMethods2(internalTypeName, data)
		return true
	} else if m.classPathLoader.CanLoad(internalTypeName) {
		data, err := m.classPathLoader.Load(internalTypeName)
		if err != nil {
			return false
		}
		m.loadFieldsAndMethods2(internalTypeName, data)
		return true
	}
	return false
}

func (m *TypeMaker) loadFieldsAndMethods2(internalTypeName string, data []byte) bool {
	if data != nil {
		reader := deserializer.NewClassFileReader(data)
		constants, err := m.loadClassFile(internalTypeName, reader)
		if err != nil {
			return false
		}

		count := reader.ReadUnsignedShort()
		for i := 0; i < count; i++ {
			// skip 'accessFlags'
			reader.Skip(2)

			nameIndex := reader.ReadUnsignedShort()
			descriptorIndex := reader.ReadUnsignedShort()

			// Load attributes
			signature := ""
			attributeCount := reader.ReadUnsignedShort()

			for j := 0; j < attributeCount; j++ {
				attributeNameIndex := reader.ReadUnsignedShort()
				attributeLength := reader.ReadInt()

				if value, ok := constants[attributeNameIndex].(string); ok && value == "Signature" {
					signature = constants[reader.ReadUnsignedShort()].(string)
				} else {
					reader.Skip(attributeLength)
				}
			}

			name := constants[nameIndex].(string)
			descriptor := constants[descriptorIndex].(string)
			key := internalTypeName + ":" + name

			if signature == "" {
				m.internalTypeNameFieldNameToType[key] = m.MakeFromSignature(descriptor)
			} else {
				m.internalTypeNameFieldNameToType[key] = m.MakeFromSignature(signature)
			}
		}

		// Load methods
		count = reader.ReadUnsignedShort()
		for i := 0; i < count; i++ {
			// skip 'accessFlags'
			reader.Skip(2)

			nameIndex := reader.ReadUnsignedShort()
			descriptorIndex := reader.ReadUnsignedShort()

			// Load attributes
			signature := ""
			var exceptionTypeNames []string
			attributeCount := reader.ReadUnsignedShort()

			for j := 0; j < attributeCount; j++ {
				attributeNameIndex := reader.ReadUnsignedShort()
				attributeLength := reader.ReadInt()
				name := constants[attributeNameIndex].(string)

				switch name {
				case "Signature":
					signature = constants[reader.ReadUnsignedShort()].(string)
				case "Exceptions":
					exceptionCount := reader.ReadUnsignedShort()
					if exceptionCount > 0 {
						exceptionTypeNames = make([]string, 0, exceptionCount)

						for k := 0; k < exceptionCount; k++ {
							exceptionClassIndex := reader.ReadUnsignedShort()
							cc := constants[exceptionClassIndex].(int)
							exceptionTypeNames[k] = constants[cc].(string)
						}
					}
					break
				default:
					reader.Skip(attributeLength)
					break
				}
			}

			name := constants[nameIndex].(string)
			descriptor := constants[descriptorIndex].(string)
			key := internalTypeName + ":" + name + descriptor
			var methodTypes intsrv.IMethodTypes

			if signature == "" {
				methodTypes = m.parseMethodSignature2(descriptor, exceptionTypeNames)
			} else {
				methodTypes = m.parseMethodSignature3(descriptor, signature, exceptionTypeNames)
			}

			m.internalTypeNameMethodNameDescriptorToMethodTypes[key] = methodTypes

			parameterCount := 0
			if methodTypes.ParameterTypes() != nil {
				parameterCount = methodTypes.ParameterTypes().Size()
			}

			key = fmt.Sprintf("%s:%s:%d", internalTypeName, name, parameterCount)
			set := m.internalTypeNameMethodNameParameterCountToDeclaredParameterTypes[key]

			if parameterCount > 0 {
				if set == nil {
					set = make([]intmod.IType, 0)
					m.internalTypeNameMethodNameParameterCountToDeclaredParameterTypes[key] = set
				}
				set = append(set, methodTypes.ParameterTypes())
			} else if set == nil {
				m.internalTypeNameMethodNameParameterCountToDeclaredParameterTypes[key] = make([]intmod.IType, 0)
			}
		}
	}

	return true
}

func (m *TypeMaker) loadClassFile(internalTypeName string, reader intsrv.IClassFileReader) ([]interface{}, error) {
	magic := reader.ReadMagic()

	if magic != intsrv.JavaMagicNumber {
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

func (m *TypeMaker) loadConstants(reader intsrv.IClassFileReader) ([]interface{}, error) {
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

func (m *TypeMaker) MatchCount(internalTypeName, name string, parameterCount int, constructor bool) int {
	suffixKey := fmt.Sprintf(":%s:%d", name, parameterCount)
	return len(m.getSetOfParameterTypes(internalTypeName, suffixKey, constructor))
}

func (m *TypeMaker) MatchCount2(typeBounds map[string]intmod.IType, internalTypeName, name string, parameters intmod.IExpression, constructor bool) int {
	parameterCount := parameters.Size()
	suffixKey := fmt.Sprintf(":%s:%d", name, parameterCount)
	setOfParameterTypes := m.getSetOfParameterTypes(internalTypeName, suffixKey, constructor)

	if parameterCount == 0 {
		return len(setOfParameterTypes)
	}

	if len(setOfParameterTypes) <= 1 {
		return len(setOfParameterTypes)
	} else {
		counter := 0

		for _, paramterTypes := range setOfParameterTypes {
			if m.match(typeBounds, paramterTypes, parameters) {
				counter++
			}
		}

		return counter
	}
}

func (m *TypeMaker) getSetOfParameterTypes(internalTypeName, suffixKey string, constructor bool) []intmod.IType {
	key := internalTypeName + suffixKey
	setOfParameterTypes := m.internalTypeNameMethodNameParameterCountToParameterTypes[key]

	if setOfParameterTypes == nil {
		setOfParameterTypes = make([]intmod.IType, 0)

		if !constructor {
			typeTypes := m.MakeTypeTypes(internalTypeName)

			if typeTypes != nil && typeTypes.SuperType() != nil {
				setOfParameterTypes = append(setOfParameterTypes,
					m.getSetOfParameterTypes(typeTypes.SuperType().InternalName(), suffixKey, constructor)...)
			}
		}

		declaredParameterTypes := m.internalTypeNameMethodNameParameterCountToDeclaredParameterTypes[key]
		if declaredParameterTypes == nil && m.loadFieldsAndMethods(internalTypeName) {
			declaredParameterTypes = m.internalTypeNameMethodNameParameterCountToDeclaredParameterTypes[key]
		}
		if declaredParameterTypes != nil {
			setOfParameterTypes = append(setOfParameterTypes, declaredParameterTypes...)
		}

		m.internalTypeNameMethodNameParameterCountToParameterTypes[key] = setOfParameterTypes
	}

	return setOfParameterTypes
}

func (m *TypeMaker) match(typeBounds map[string]intmod.IType, parameterTypes intmod.IType, parameters intmod.IExpression) bool {
	if parameterTypes.Size() != parameters.Size() {
		return false
	}

	switch parameterTypes.Size() {
	case 0:
		return true
	case 1:
		return m.match2(typeBounds, parameterTypes.Type(), parameters.Type())
	default:
		for i := 0; i < parameterTypes.Size(); i++ {
			if !m.match2(typeBounds, parameterTypes.Type(), parameters.Type()) {
				return false
			}
		}
		return true
	}
}

func (m *TypeMaker) match2(typeBounds map[string]intmod.IType, leftType intmod.IType, rightType intmod.IType) bool {
	if leftType == rightType {
		return true
	}

	if leftType.IsPrimitiveType() && rightType.IsPrimitiveType() {
		flags := leftType.(intmod.IPrimitiveType).Flags() | rightType.(intmod.IPrimitiveType).Flags()
		return flags != 0
	}

	if leftType.IsObjectType() && rightType.IsObjectType() {
		return m.IsAssignable(typeBounds, leftType.(intmod.IObjectType), rightType.(intmod.IObjectType))
	}

	return false
}

func NewClassPathLoader(buffer []byte) intsrv.IClassPathLoader {
	return &ClassPathLoader{
		buffer: buffer,
	}
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

func NewSignatureReader(signature string) intsrv.ISignatureReader {
	return NewSignatureReaderWithAll(signature, 0)
}

func NewSignatureReaderWithAll(signature string, index int) intsrv.ISignatureReader {
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

func (r *SignatureReader) Signature() string {
	return r.signature
}

func (r *SignatureReader) Array() []byte {
	return r.array
}

func (r *SignatureReader) Length() int {
	return r.length
}

func (r *SignatureReader) Index() int {
	return r.index
}

func (r *SignatureReader) Inc() {
	r.index++
}

func (r *SignatureReader) Dec() {
	r.index--
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

func skipMembers(reader intsrv.IClassFileReader) {
	count := reader.ReadUnsignedShort()
	for i := 0; i < count; i++ {
		reader.Skip(3 * 2)
		skipAttributes(reader)
	}
}

func skipAttributes(reader intsrv.IClassFileReader) {
	count := reader.ReadUnsignedShort()
	for i := 0; i < count; i++ {
		reader.Skip(2)
		attributeLength := reader.ReadInt()
		reader.Skip(attributeLength)
	}
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

func NewTypeTypes() intsrv.ITypeTypes {
	return NewTypeTypes2(nil, nil, nil, nil)
}

func NewTypeTypes2(thisType intmod.IObjectType, typeParameters intmod.ITypeParameter,
	superType intmod.IObjectType, interfaces intmod.IType) intsrv.ITypeTypes {
	return &TypeTypes{
		thisType:       thisType,
		typeParameters: typeParameters,
		superType:      superType,
		interfaces:     interfaces,
	}
}

type TypeTypes struct {
	thisType       intmod.IObjectType
	typeParameters intmod.ITypeParameter
	superType      intmod.IObjectType
	interfaces     intmod.IType
}

func (t *TypeTypes) ThisType() intmod.IObjectType {
	return t.thisType
}

func (t *TypeTypes) SetThisType(thisType intmod.IObjectType) {
	t.thisType = thisType
}

func (t *TypeTypes) TypeParameters() intmod.ITypeParameter {
	return t.typeParameters
}

func (t *TypeTypes) SetTypeParameters(typeParameters intmod.ITypeParameter) {
	t.typeParameters = typeParameters
}

func (t *TypeTypes) SuperType() intmod.IObjectType {
	return t.superType
}

func (t *TypeTypes) SetSuperType(superType intmod.IObjectType) {
	t.superType = superType
}

func (t *TypeTypes) Interfaces() intmod.IType {
	return t.interfaces
}

func (t *TypeTypes) SetInterfaces(interfaces intmod.IType) {
	t.interfaces = interfaces
}

func NewMethodTypes() intsrv.IMethodTypes {
	return NewMethodTypes2(nil, nil, nil, nil)
}

func NewMethodTypes2(typeParameters intmod.ITypeParameter, parameterTypes intmod.IType,
	returnedType intmod.IType, exceptionTypes intmod.IType) intsrv.IMethodTypes {
	return &MethodTypes{
		typeParameters: typeParameters,
		parameterTypes: parameterTypes,
		returnedType:   returnedType,
		exceptionTypes: exceptionTypes,
	}
}

type MethodTypes struct {
	typeParameters intmod.ITypeParameter
	parameterTypes intmod.IType
	returnedType   intmod.IType
	exceptionTypes intmod.IType
}

func (t *MethodTypes) TypeParameters() intmod.ITypeParameter {
	return t.typeParameters
}

func (t *MethodTypes) SetTypeParameters(typeParameters intmod.ITypeParameter) {
	t.typeParameters = typeParameters
}

func (t *MethodTypes) ParameterTypes() intmod.IType {
	return t.parameterTypes
}

func (t *MethodTypes) SetParameterTypes(parameterTypes intmod.IType) {
	t.parameterTypes = parameterTypes
}

func (t *MethodTypes) ReturnedType() intmod.IType {
	return t.returnedType
}
func (t *MethodTypes) SetReturnedType(returnedType intmod.IType) {
	t.returnedType = returnedType
}

func (t *MethodTypes) ExceptionTypes() intmod.IType {
	return t.exceptionTypes
}

func (t *MethodTypes) SetExceptionTypes(exceptionTypes intmod.IType) {
	t.exceptionTypes = exceptionTypes
}

func CountDimension(descriptor string) int {
	count := 0
	length := len(descriptor)

	for i := 0; i < length && descriptor[i] == '['; i++ {
		count++
	}
	return count
}

func getExceptionTypeNames(method intcls.IMethod) []string {
	if method != nil {
		attributeExceptions := method.Attributes()["Exceptions"].(*attribute.AttributeExceptions)

		if attributeExceptions != nil {
			return attributeExceptions.ExceptionTypeNames()
		}
	}

	return nil
}

func isAReferenceTypeSignature(reader intsrv.ISignatureReader) bool {
	if reader.Available() {
		c := reader.Read()

		for c == '[' {
			c = reader.Read()
		}

		switch c {
		case 'B', 'C', 'D', 'F', 'I', 'J':
			return true
		case 'L':
			// Unread 'L'
			reader.Dec()
			return isAClassTypeSignature(reader)
		case 'S':
			return true
		case 'T':
			reader.SearchEndMarker()
			return true
		case 'V', 'Z':
			return true
		default:
			// Unread 'c'
			reader.Dec()
			return false
		}
	}
	return false
}

func isAClassTypeSignature(reader intsrv.ISignatureReader) bool {
	if reader.NextEqualsTo('L') {
		reader.Inc()
		endMarker := reader.SearchEndMarker()

		if endMarker == 0 {
			return false
		}

		if endMarker == '<' {
			reader.Inc()
			isATypeArguments(reader)
			if reader.Read() != '>' {
				return false
			}
		}

		for reader.NextEqualsTo('.') {
			reader.Inc()
			endMarker = reader.SearchEndMarker()

			if endMarker == 0 {
				return false
			}

			if endMarker == '<' {
				reader.Inc()
				isATypeArguments(reader)
				if reader.Read() != '>' {
					return false
				}
			}
		}

		reader.Inc()

		return true
	}

	return false
}

func isATypeArguments(reader intsrv.ISignatureReader) bool {
	if !isATypeArgument(reader) {
		return false
	}

	for isATypeArgument(reader) {
	}

	return false
}

func isATypeArgument(reader intsrv.ISignatureReader) bool {
	switch reader.Read() {
	case '+', '-':
		return isAReferenceTypeSignature(reader)
	case '*':
		return true
	default:
		reader.Dec()
		return false
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

func hashCodeWithString(str string) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(str))
	if err != nil {
		return -1
	}
	return int(h.Sum32())
}
