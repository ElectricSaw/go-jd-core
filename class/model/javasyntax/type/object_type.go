package _type

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

var OtTypeBoolean = NewObjectType("class/lang/Boolean", "class.lang.Boolean", "Boolean")
var OtTypeByte = NewObjectType("class/lang/Byte", "class.lang.Byte", "Byte")
var OtTypeCharacter = NewObjectType("class/lang/Character", "class.lang.Character", "Character")
var OtTypeClass = NewObjectType("class/lang/Class", "class.lang.Class", "Class")
var OtTypeClassWildcard = OtTypeClass.CreateTypeWithArgs(WildcardTypeArgumentEmpty.(intsyn.ITypeArgument))

var OtTypeDouble = NewObjectType("class/lang/Double", "class.lang.Double", "Double")
var OtTypeException = NewObjectType("class/lang/Exception", "class.lang.Exception", "Exception")
var OtTypeFloat = NewObjectType("class/lang/Float", "class.lang.Float", "Float")
var OtTypeInteger = NewObjectType("class/lang/Integer", "class.lang.Integer", "Integer")
var OtTypeIterable = NewObjectType("class/lang/Iterable", "class.lang.Iterable", "Iterable")
var OtTypeLong = NewObjectType("class/lang/Long", "class.lang.Long", "Long")
var OtTypeMath = NewObjectType("class/lang/Math", "class.lang.Math", "Math")
var OtTypeObject = NewObjectType("class/lang/Object", "class.lang.Object", "Object")
var OtTypeRuntimeException = NewObjectType("class/lang/RuntimeException", "class.lang.RuntimeException", "RuntimeException")
var OtTypeShort = NewObjectType("class/lang/Short", "class.lang.Short", "Short")
var OtTypeString = NewObjectType("class/lang/String", "class.lang.String", "String")
var OtTypeStringBuffer = NewObjectType("class/lang/StringBuffer", "class.lang.StringBuffer", "StringBuffer")
var OtTypeStringBuilder = NewObjectType("class/lang/StringBuilder", "class.lang.StringBuilder", "StringBuilder")
var OtTypeSystem = NewObjectType("class/lang/System", "class.lang.System", "System")
var OtTypeThread = NewObjectType("class/lang/Thread", "class.lang.Thread", "Thread")
var OtTypeThrowable = NewObjectType("class/lang/Throwable", "class.lang.Throwable", "Throwable")

var OtTypePrimitiveBoolean = NewObjectTypeWithDesc("Z")
var OtTypePrimitiveByte = NewObjectTypeWithDesc("B")
var OtTypePrimitiveChar = NewObjectTypeWithDesc("C")
var OtTypePrimitiveDouble = NewObjectTypeWithDesc("D")
var OtTypePrimitiveFloat = NewObjectTypeWithDesc("F")
var OtTypePrimitiveInt = NewObjectTypeWithDesc("I")
var OtTypePrimitiveLong = NewObjectTypeWithDesc("J")
var OtTypePrimitiveShort = NewObjectTypeWithDesc("S")
var OtTypePrimitiveVoid = NewObjectTypeWithDesc("V")

var OtTypeUndefinedObject = NewObjectType("class/lang/Object", "class.lang.Object", "Object")

func createDescriptor(descriptor string, dimension int) string {
	switch dimension {
	case 0:
		return descriptor
	case 1:
		return "[" + descriptor
	case 2:
		return "[[" + descriptor
	default:
		ret := ""
		for i := 0; i < dimension; i++ {
			ret += "["
		}
		return ret + descriptor
	}
}

func NewObjectType(internalName, qualifiedName, name string) intsyn.IObjectType {
	return NewObjectTypeWithAll(internalName, qualifiedName, name, nil, 0)
}

func NewObjectTypeWithDim(internalName, qualifiedName, name string, dimension int) intsyn.IObjectType {
	return NewObjectTypeWithAll(internalName, qualifiedName, name, nil, dimension)
}

func NewObjectTypeWithArgs(internalName, qualifiedName, name string, typeArguments intsyn.ITypeArgument) intsyn.IObjectType {
	return NewObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, 0)
}

func NewObjectTypeWithAll(internalName, qualifiedName, name string, typeArguments intsyn.ITypeArgument, dimension int) intsyn.IObjectType {
	return &ObjectType{
		internalName:  internalName,
		qualifiedName: qualifiedName,
		name:          name,
		typeArguments: typeArguments,
		dimension:     dimension,
		descriptor:    createDescriptor(fmt.Sprintf("L%s;", internalName), dimension),
	}
}

func NewObjectTypeWithDesc(primitiveDescriptor string) intsyn.IObjectType {
	return NewObjectTypeWithDescAndDim(primitiveDescriptor, 0)
}

func NewObjectTypeWithDescAndDim(primitiveDescriptor string, dimension int) intsyn.IObjectType {
	return &ObjectType{
		internalName:  primitiveDescriptor,
		qualifiedName: GetPrimitiveType(int(primitiveDescriptor[0])).Name(),
		dimension:     dimension,
		descriptor:    createDescriptor(fmt.Sprintf("L%s;", primitiveDescriptor), dimension),
	}
}

type ObjectType struct {
	AbstractType
	AbstractTypeArgument

	internalName  string
	qualifiedName string
	name          string
	typeArguments intsyn.ITypeArgument
	dimension     int
	descriptor    string
}

/////////////////////////////////////////////////////////////////////

func (t *ObjectType) QualifiedName() string {
	return t.qualifiedName
}

func (t *ObjectType) HashCode() int {
	result := 735485092 + hashCodeWithString(t.internalName)
	result *= 31
	if t.typeArguments != nil {
		result += t.typeArguments.HashCode()
	}
	result = 31*result + t.dimension
	return result
}

/////////////////////////////////////////////////////////////////////

func (t *ObjectType) Name() string {
	return t.name
}

func (t *ObjectType) Descriptor() string {
	return t.descriptor
}

func (t *ObjectType) Dimension() int {
	return t.dimension
}

func (t *ObjectType) CreateType(dimension int) intsyn.IType {
	if t.dimension == dimension {
		return t
	} else if t.descriptor[len(t.descriptor)-1] != ';' {
		if dimension == 0 {
			return GetPrimitiveType(int(t.descriptor[t.dimension])).(intsyn.IType)
		} else {
			return NewObjectTypeWithDescAndDim(t.internalName, t.dimension).(intsyn.IType)
		}
	} else {
		return NewObjectTypeWithAll(t.internalName, t.qualifiedName, t.name, t.typeArguments, dimension).(intsyn.IType)
	}
}

func (t *ObjectType) IsObjectType() bool {
	return true
}

func (t *ObjectType) InternalName() string {
	return t.internalName
}

func (t *ObjectType) AcceptTypeVisitor(visitor intsyn.ITypeVisitor) {
	visitor.VisitObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *ObjectType) IsTypeArgumentAssignableFrom(typeBounds map[string]intsyn.IType, typeArgument intsyn.ITypeArgument) bool {
	switch meta := typeArgument.(type) {
	case intsyn.IObjectType:
		if t.dimension != meta.Dimension() || t.internalName != meta.InternalName() {
			return false
		}

		if meta.TypeArguments() == nil {
			return t.typeArguments == nil
		} else if t.typeArguments == nil {
			return false
		} else {
			return t.typeArguments.IsTypeArgumentAssignableFrom(typeBounds, meta.TypeArguments())
		}
	case *GenericType:
		bt := typeBounds[meta.Name()]
		ot, ok := bt.(intsyn.IObjectType)

		if ok {
			if t.internalName == ot.InternalName() {
				return true
			}
		}
	}

	return false
}

func (t *ObjectType) IsTypeArgumentAssignableFromWithObj(typeBounds map[string]intsyn.IType, objectType intsyn.IObjectType) bool {
	if t.dimension != objectType.Dimension() || t.internalName != objectType.InternalName() {
		return false
	}

	if objectType.TypeArguments() == nil {
		return t.typeArguments == nil
	} else if t.typeArguments == nil {
		return false
	} else {
		return t.typeArguments.IsTypeArgumentAssignableFrom(typeBounds, objectType.TypeArguments())
	}
}

func (t *ObjectType) IsObjectTypeArgument() bool {
	return true
}

func (t *ObjectType) AcceptTypeArgumentVisitor(visitor intsyn.ITypeArgumentVisitor) {
	visitor.VisitObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *ObjectType) TypeArguments() intsyn.ITypeArgument {
	return t.typeArguments
}

func (t *ObjectType) CreateTypeWithArgs(typeArguments intsyn.ITypeArgument) intsyn.IObjectType {
	if t.typeArguments == typeArguments {
		return t
	} else {
		return NewObjectTypeWithAll(t.internalName, t.qualifiedName, t.name, typeArguments, t.dimension)
	}
}

func (t *ObjectType) String() string {
	msg := fmt.Sprintf("ObjectType{ %s", t.internalName)
	if t.typeArguments != nil {
		msg += fmt.Sprintf("<%s>", t.typeArguments)
	}
	if t.dimension > 0 {
		msg += fmt.Sprintf(", %d", t.dimension)
	}
	msg += " }"
	return msg
}
