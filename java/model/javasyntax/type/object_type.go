package _type

import "fmt"

var OtTypeBoolean = NewObjectType("java/lang/Boolean", "java.lang.Boolean", "Boolean")
var OtTypeByte = NewObjectType("java/lang/Byte", "java.lang.Byte", "Byte")
var OtTypeCharacter = NewObjectType("java/lang/Character", "java.lang.Character", "Character")
var OtTypeClass = NewObjectType("java/lang/Class", "java.lang.Class", "Class")
var OtTypeClassWildcard = OtTypeClass.CreateTypeWithArgs(WildcardTypeArgumentEmpty)

var OtTypeDouble = NewObjectType("java/lang/Double", "java.lang.Double", "Double")
var OtTypeException = NewObjectType("java/lang/Exception", "java.lang.Exception", "Exception")
var OtTypeFloat = NewObjectType("java/lang/Float", "java.lang.Float", "Float")
var OtTypeInteger = NewObjectType("java/lang/Integer", "java.lang.Integer", "Integer")
var OtTypeIterable = NewObjectType("java/lang/Iterable", "java.lang.Iterable", "Iterable")
var OtTypeLong = NewObjectType("java/lang/Long", "java.lang.Long", "Long")
var OtTypeMath = NewObjectType("java/lang/Math", "java.lang.Math", "Math")
var OtTypeObject = NewObjectType("java/lang/Object", "java.lang.Object", "Object")
var OtTypeRuntimeException = NewObjectType("java/lang/RuntimeException", "java.lang.RuntimeException", "RuntimeException")
var OtTypeShort = NewObjectType("java/lang/Short", "java.lang.Short", "Short")
var OtTypeString = NewObjectType("java/lang/String", "java.lang.String", "String")
var OtTypeStringBuffer = NewObjectType("java/lang/StringBuffer", "java.lang.StringBuffer", "StringBuffer")
var OtTypeStringBuilder = NewObjectType("java/lang/StringBuilder", "java.lang.StringBuilder", "StringBuilder")
var OtTypeSystem = NewObjectType("java/lang/System", "java.lang.System", "System")
var OtTypeThread = NewObjectType("java/lang/Thread", "java.lang.Thread", "Thread")
var OtTypeThrowable = NewObjectType("java/lang/Throwable", "java.lang.Throwable", "Throwable")

var OtTypePrimitiveBoolean = NewObjectTypeWithDesc("Z")
var OtTypePrimitiveByte = NewObjectTypeWithDesc("B")
var OtTypePrimitiveChar = NewObjectTypeWithDesc("C")
var OtTypePrimitiveDouble = NewObjectTypeWithDesc("D")
var OtTypePrimitiveFloat = NewObjectTypeWithDesc("F")
var OtTypePrimitiveInt = NewObjectTypeWithDesc("I")
var OtTypePrimitiveLong = NewObjectTypeWithDesc("J")
var OtTypePrimitiveShort = NewObjectTypeWithDesc("S")
var OtTypePrimitiveVoid = NewObjectTypeWithDesc("V")

var TypeUndefinedObject = NewObjectType("java/lang/Object", "java.lang.Object", "Object")

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

func NewObjectType(internalName, qualifiedName, name string) *ObjectType {
	return NewObjectTypeWithAll(internalName, qualifiedName, name, nil, 0)
}

func NewObjectTypeWithDim(internalName, qualifiedName, name string, dimension int) *ObjectType {
	return NewObjectTypeWithAll(internalName, qualifiedName, name, nil, dimension)
}

func NewObjectTypeWithArgs(internalName, qualifiedName, name string, typeArguments ITypeArgument) *ObjectType {
	return NewObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, 0)
}

func NewObjectTypeWithAll(internalName, qualifiedName, name string, typeArguments ITypeArgument, dimension int) *ObjectType {
	return &ObjectType{
		internalName:  internalName,
		qualifiedName: qualifiedName,
		name:          name,
		typeArguments: typeArguments,
		dimension:     dimension,
		descriptor:    createDescriptor(fmt.Sprintf("L%s;", internalName), dimension),
	}
}

func NewObjectTypeWithDesc(primitiveDescriptor string) *ObjectType {
	return NewObjectTypeWithDescAndDim(primitiveDescriptor, 0)
}

func NewObjectTypeWithDescAndDim(primitiveDescriptor string, dimension int) *ObjectType {
	return &ObjectType{
		internalName: primitiveDescriptor,
		//qualifiedName: qualifiedName,  // PrimitiveType.getPrimitiveType(primitiveDescriptor.charAt(0)).getName();
		//name: name,  //PrimitiveType.getPrimitiveType(primitiveDescriptor.charAt(0)).getName();
		dimension:  dimension,
		descriptor: createDescriptor(fmt.Sprintf("L%s;", primitiveDescriptor), dimension),
	}
}

type ObjectType struct {
	AbstractType
	AbstractTypeArgument

	internalName  string
	qualifiedName string
	name          string
	typeArguments ITypeArgument
	dimension     int
	descriptor    string
}

/////////////////////////////////////////////////////////////////////

func (t *ObjectType) QualifiedName() string {
	return t.qualifiedName
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

func (t *ObjectType) CreateType(dimension int) IType {
	if t.dimension == dimension {
		return t
	} else if t.descriptor[len(t.descriptor)-1] != ';' {
		if dimension == 0 {
			return GetPrimitiveType(int(t.descriptor[t.dimension]))
		} else {
			return NewObjectTypeWithDescAndDim(t.internalName, t.dimension)
		}
	} else {
		return NewObjectTypeWithAll(t.internalName, t.qualifiedName, t.name, t.typeArguments, dimension)
	}
}

func (t *ObjectType) IsObjectType() bool {
	return true
}

func (t *ObjectType) InternalName() string {
	return t.internalName
}

func (t *ObjectType) AcceptTypeVisitor(visitor TypeVisitor) {
	visitor.VisitObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *ObjectType) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	switch meta := typeArgument.(type) {
	case IObjectType:
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
		ot, ok := bt.(IObjectType)

		if ok {
			if t.internalName == ot.InternalName() {
				return true
			}
		}
	}

	return false
}

func (t *ObjectType) IsTypeArgumentAssignableFromWithObj(typeBounds map[string]IType, objectType ObjectType) bool {
	if t.dimension != objectType.dimension || t.internalName != objectType.internalName {
		return false
	}

	if objectType.TypeArguments() == nil {
		return t.typeArguments == nil
	} else if t.typeArguments == nil {
		return false
	} else {
		return t.typeArguments.IsTypeArgumentAssignableFrom(typeBounds, objectType.typeArguments)
	}
}

func (t *ObjectType) IsObjectTypeArgument() bool {
	return true
}

func (t *ObjectType) AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor) {
	visitor.VisitObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *ObjectType) TypeArguments() ITypeArgument {
	return t.typeArguments
}

func (t *ObjectType) CreateTypeWithArgs(typeArguments ITypeArgument) *ObjectType {
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
