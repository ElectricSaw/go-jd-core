package model

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"reflect"
)

const (
	FlagBoolean = 1 << iota
	FlagChar
	FlagFloat
	FlagDouble
	FlagByte
	FlagShort
	FlagInt
	FlagLong
	FlagVoid
)

var (
	PtTypeBoolean              = NewPrimitiveType("boolean", FlagBoolean, FlagBoolean, FlagBoolean)
	PtTypeByte                 = NewPrimitiveType("byte", FlagByte, FlagByte, FlagByte|FlagInt|FlagShort)
	PtTypeChar                 = NewPrimitiveType("char", FlagChar, FlagChar, FlagChar|FlagInt)
	PtTypeDouble               = NewPrimitiveType("double", FlagDouble, FlagDouble, FlagDouble)
	PtTypeFloat                = NewPrimitiveType("float", FlagFloat, FlagFloat, FlagFloat)
	PtTypeInt                  = NewPrimitiveType("int", FlagInt, FlagInt|FlagByte|FlagChar|FlagShort, FlagInt)
	PtTypeLong                 = NewPrimitiveType("long", FlagLong, FlagLong, FlagLong)
	PtTypeShort                = NewPrimitiveType("short", FlagShort, FlagShort|FlagByte, FlagShort|FlagInt)
	PtTypeVoid                 = NewPrimitiveType("void", FlagVoid, FlagVoid, FlagVoid)
	PtMaybeCharType            = NewPrimitiveType("maybe_char", FlagChar|FlagInt, FlagChar|FlagInt, FlagChar|FlagInt)                                                                                                 //  32768 .. 65535
	PtMaybeShortType           = NewPrimitiveType("maybe_short", FlagChar|FlagShort|FlagInt, FlagChar|FlagShort|FlagInt, FlagChar|FlagShort|FlagInt)                                                                  //    128 .. 32767
	PtMaybeByteType            = NewPrimitiveType("maybe_byte", FlagByte|FlagChar|FlagShort|FlagInt, FlagByte|FlagChar|FlagShort|FlagInt, FlagByte|FlagChar|FlagShort|FlagInt)                                        //      2 .. 127
	PtMaybeBooleanType         = NewPrimitiveType("maybe_boolean", FlagBoolean|FlagByte|FlagChar|FlagShort|FlagInt, FlagBoolean|FlagByte|FlagChar|FlagShort|FlagInt, FlagBoolean|FlagByte|FlagChar|FlagShort|FlagInt) //      0 .. 1
	PtMaybeNegativeByteType    = NewPrimitiveType("maybe_negative_byte", FlagByte|FlagShort|FlagInt, FlagByte|FlagShort|FlagInt, FlagByte|FlagShort|FlagInt)                                                          //   -128 .. -1
	PtMaybeNegativeShortType   = NewPrimitiveType("maybe_negative_short", FlagShort|FlagInt, FlagShort|FlagInt, FlagShort|FlagInt)                                                                                    // -32768 .. -129
	PtMaybeIntType             = NewPrimitiveType("maybe_int", FlagInt, FlagInt, FlagInt)                                                                                                                             // Otherwise
	PtMaybeNegativeBooleanType = NewPrimitiveType("maybe_negative_boolean", FlagBoolean|FlagByte|FlagShort|FlagInt, FlagBoolean|FlagByte|FlagShort|FlagInt, FlagBoolean|FlagByte|FlagShort|FlagInt)                   // Boolean or negative

	descriptorToType = map[int]PrimitiveType{
		int('B') - int('B'): PtTypeByte,
		int('C') - int('B'): PtTypeChar,
		int('D') - int('B'): PtTypeDouble,
		int('F') - int('B'): PtTypeFloat,
		int('I') - int('B'): PtTypeInt,
		int('J') - int('B'): PtTypeLong,
		int('S') - int('B'): PtTypeShort,
		int('V') - int('B'): PtTypeVoid,
		int('Z') - int('B'): PtTypeBoolean,
	}
)

var (
	Diamond                   = NewDiamondTypeArgument()
	WildcardTypeArgumentEmpty = NewWildcardTypeArgument()
)

var (
	OtTypeBoolean       = NewObjectType("class/lang/Boolean", "class.lang.Boolean", "Boolean")
	OtTypeByte          = NewObjectType("class/lang/Byte", "class.lang.Byte", "Byte")
	OtTypeCharacter     = NewObjectType("class/lang/Character", "class.lang.Character", "Character")
	OtTypeClass         = NewObjectType("class/lang/Class", "class.lang.Class", "Class")
	OtTypeClassWildcard = OtTypeClass.CreateTypeWithArgs(ITypeArgument(&WildcardTypeArgumentEmpty))

	OtTypeDouble           = NewObjectType("class/lang/Double", "class.lang.Double", "Double")
	OtTypeException        = NewObjectType("class/lang/Exception", "class.lang.Exception", "Exception")
	OtTypeFloat            = NewObjectType("class/lang/Float", "class.lang.Float", "Float")
	OtTypeInteger          = NewObjectType("class/lang/Integer", "class.lang.Integer", "Integer")
	OtTypeIterable         = NewObjectType("class/lang/Iterable", "class.lang.Iterable", "Iterable")
	OtTypeLong             = NewObjectType("class/lang/Long", "class.lang.Long", "Long")
	OtTypeMath             = NewObjectType("class/lang/Math", "class.lang.Math", "Math")
	OtTypeObject           = NewObjectType("class/lang/Object", "class.lang.Object", "Object")
	OtTypeRuntimeException = NewObjectType("class/lang/RuntimeException", "class.lang.RuntimeException", "RuntimeException")
	OtTypeShort            = NewObjectType("class/lang/Short", "class.lang.Short", "Short")
	OtTypeString           = NewObjectType("class/lang/String", "class.lang.String", "String")
	OtTypeStringBuffer     = NewObjectType("class/lang/StringBuffer", "class.lang.StringBuffer", "StringBuffer")
	OtTypeStringBuilder    = NewObjectType("class/lang/StringBuilder", "class.lang.StringBuilder", "StringBuilder")
	OtTypeSystem           = NewObjectType("class/lang/System", "class.lang.System", "System")
	OtTypeThread           = NewObjectType("class/lang/Thread", "class.lang.Thread", "Thread")
	OtTypeThrowable        = NewObjectType("class/lang/Throwable", "class.lang.Throwable", "Throwable")

	OtTypePrimitiveBoolean = NewObjectTypeWithDesc("Z")
	OtTypePrimitiveByte    = NewObjectTypeWithDesc("B")
	OtTypePrimitiveChar    = NewObjectTypeWithDesc("C")
	OtTypePrimitiveDouble  = NewObjectTypeWithDesc("D")
	OtTypePrimitiveFloat   = NewObjectTypeWithDesc("F")
	OtTypePrimitiveInt     = NewObjectTypeWithDesc("I")
	OtTypePrimitiveLong    = NewObjectTypeWithDesc("J")
	OtTypePrimitiveShort   = NewObjectTypeWithDesc("S")
	OtTypePrimitiveVoid    = NewObjectTypeWithDesc("V")

	OtTypeUndefinedObject = NewObjectType("class/lang/Object", "class.lang.Object", "Object")
)

func GetPrimitiveType(primitiveDescriptor int) PrimitiveType {
	return descriptorToType[primitiveDescriptor-66] // int('B')
}

func NewPrimitiveType(name string, flags, leftFlags, rightFlags int) PrimitiveType {
	t := PrimitiveType{
		Name:       name,
		Dimension:  0,
		Flags:      flags,
		LeftFlags:  leftFlags,
		RightFlags: rightFlags,
	}

	sb := ""

	if flags&FlagDouble != 0 {
		sb += "D"
	} else if flags&FlagFloat != 0 {
		sb += "F"
	} else if flags&FlagLong != 0 {
		sb += "J"
	} else if flags&FlagBoolean != 0 {
		sb += "Z"
	} else if flags&FlagByte != 0 {
		sb += "B"
	} else if flags&FlagChar != 0 {
		sb += "C"
	} else if flags&FlagShort != 0 {
		sb += "S"
	} else {
		sb += "I"
	}

	t.Descriptor = sb
	t.SetValue(&t)

	return t
}

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

func NewObjectType(internalName, qualifiedName, name string) ObjectType {
	return NewObjectTypeWithAll(internalName, qualifiedName, name, nil, 0)
}

func NewObjectTypeWithDim(internalName, qualifiedName, name string, dimension int) ObjectType {
	return NewObjectTypeWithAll(internalName, qualifiedName, name, nil, dimension)
}

func NewObjectTypeWithArgs(internalName, qualifiedName, name string, typeArguments ITypeArgument) ObjectType {
	return NewObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, 0)
}

func NewObjectTypeWithAll(internalName, qualifiedName, name string, typeArguments ITypeArgument, dimension int) ObjectType {
	t := ObjectType{
		InternalName:  internalName,
		QualifiedName: qualifiedName,
		Name:          name,
		TypeArguments: typeArguments,
		Dimension:     dimension,
		Descriptor:    createDescriptor(fmt.Sprintf("L%s;", internalName), dimension),
	}
	t.SetValue(&t)
	return t
}

func NewObjectTypeWithDesc(primitiveDescriptor string) ObjectType {
	return NewObjectTypeWithDescAndDim(primitiveDescriptor, 0)
}

func NewObjectTypeWithDescAndDim(primitiveDescriptor string, dimension int) ObjectType {
	t := ObjectType{
		InternalName:  primitiveDescriptor,
		QualifiedName: GetPrimitiveType(int(primitiveDescriptor[0])).Name,
		Dimension:     dimension,
		Descriptor:    createDescriptor(fmt.Sprintf("L%s;", primitiveDescriptor), dimension),
	}
	t.SetValue(&t)
	return t
}

func NewInnerObjectType(internalName, qualifiedName, name string, outerType *ObjectType) InnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, nil, 0, outerType)
}

func NewInnerObjectTypeWithDim(internalName, qualifiedName, name string, dimension int, outerType *ObjectType) InnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, nil, dimension, outerType)
}

func NewInnerObjectTypeWithArgs(internalName, qualifiedName, name string, typeArguments ITypeArgument, outerType *ObjectType) InnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, 0, outerType)
}

func NewInnerObjectTypeWithAll(internalName, qualifiedName, name string, typeArguments ITypeArgument, dimension int, outerType *ObjectType) InnerObjectType {
	t := InnerObjectType{
		InternalName:  internalName,
		QualifiedName: qualifiedName,
		Name:          name,
		TypeArguments: typeArguments,
		Dimension:     dimension,
		Descriptor:    createDescriptor(fmt.Sprintf("L%s;", internalName), dimension),
		OuterType:     outerType,
	}
	t.SetValue(&t)
	return t
}

func NewTypes() Types {
	return NewTypesWithSlice(make([]IType, 0)...)
}

func NewTypesWithSlice(types ...IType) Types {
	return Types{
		DefaultList: *util.NewDefaultListWithSlice[IType](types).(*util.DefaultList[IType]),
	}
}

func NewGenericType(name string) GenericType {
	return NewGenericTypeWithAll(name, 0)
}

func NewGenericTypeWithAll(name string, dimension int) GenericType {
	t := GenericType{
		Name:       name,
		Descriptor: name,
		Dimension:  dimension,
	}
	t.SetValue(&t)
	return t
}

func NewUnmodifiableTypes(types ...IType) UnmodifiableTypes {
	return NewUnmodifiableTypesWithSlice(types)
}

func NewUnmodifiableTypesWithSlice(types []IType) UnmodifiableTypes {
	t := UnmodifiableTypes{}
	t.AddAll(types)
	return t
}

type IType interface {
	GetDimension() int

	CreateType(dimension int) IType

	IsGenericType() bool
	IsInnerObjectType() bool
	IsObjectType() bool
	IsPrimitiveType() bool
	IsTypes() bool

	GetOuterType() *ObjectType
	GetInternalName() string

	AcceptTypeVisitor(visitor ITypeVisitor)
	HashCode() int
	Equals(o interface{}) bool
	String() string
}

type ITypeVisitor interface {
	VisitPrimitiveType(y *PrimitiveType)
	VisitObjectType(y IType)
	VisitInnerObjectType(y IType)
	VisitTypes(types *Types)
	VisitGenericType(y *GenericType)
}

type ITypeVisitable interface {
	AcceptTypeVisitor(visitor ITypeVisitor)
}

type PrimitiveType struct {
	util.DefaultBase[IType]
	Name       string
	Dimension  int
	Flags      int
	LeftFlags  int
	RightFlags int
	Descriptor string
}

func (t *PrimitiveType) GetDimension() int {
	return t.Dimension
}

func (t *PrimitiveType) CreateType(dimension int) IType {
	if dimension == 0 {
		return t
	} else {
		tmp := NewObjectTypeWithDescAndDim(t.Descriptor, dimension)
		return &tmp
	}
}

func (t *PrimitiveType) IsGenericType() bool {
	return false
}

func (t *PrimitiveType) IsInnerObjectType() bool {
	return false
}

func (t *PrimitiveType) IsObjectType() bool {
	return false
}

func (t *PrimitiveType) IsPrimitiveType() bool {
	return true
}

func (t *PrimitiveType) IsTypes() bool {
	return false
}

func (t *PrimitiveType) GetOuterType() *ObjectType {
	return &OtTypeUndefinedObject
}

func (t *PrimitiveType) GetInternalName() string {
	return ""
}

func (t *PrimitiveType) HashCode() int {
	return 750039781 + t.Flags
}

func (t *PrimitiveType) AcceptTypeVisitor(visitor ITypeVisitor) {
	visitor.VisitPrimitiveType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *PrimitiveType) TypeArgumentFirst() ITypeArgument {
	return t
}

func (t *PrimitiveType) TypeArgumentList() util.IList[ITypeArgument] {
	return util.NewDefaultListWithElements[ITypeArgument](t)
}

func (t *PrimitiveType) TypeArgumentSize() int {
	return 1
}

func (t *PrimitiveType) GetType() IType {
	return &OtTypeUndefinedObject
}

func (t *PrimitiveType) IsTypeArgumentAssignableFrom(_ map[string]IType, typeArgument ITypeArgument) bool {
	if o, ok := typeArgument.(*PrimitiveType); ok {
		return t.Equals(o)
	}

	return false
}

func (t *PrimitiveType) IsTypeArgumentList() bool {
	return false
}

func (t *PrimitiveType) IsGenericTypeArgument() bool {
	return false
}

func (t *PrimitiveType) IsInnerObjectTypeArgument() bool {
	return false
}

func (t *PrimitiveType) IsObjectTypeArgument() bool {
	return false
}

func (t *PrimitiveType) IsPrimitiveTypeArgument() bool {
	return true
}

func (t *PrimitiveType) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (t *PrimitiveType) IsWildcardSuperTypeArgument() bool {
	return false
}

func (t *PrimitiveType) IsWildcardTypeArgument() bool {
	return false
}

func (t *PrimitiveType) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitPrimitiveType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *PrimitiveType) JavaPrimitiveFlags() int {
	if t.Flags&FlagBoolean != 0 {
		return FlagBoolean
	} else if t.Flags&FlagInt != 0 {
		return FlagInt
	} else if t.Flags&FlagChar != 0 {
		return FlagChar
	} else if t.Flags&FlagShort != 0 {
		return FlagShort
	} else if t.Flags&FlagByte != 0 {
		return FlagByte
	}
	return t.Flags
}

func (t *PrimitiveType) Equals(o interface{}) bool {
	if t == o {
		return true
	}

	if o == nil {
		return false
	}

	var other *PrimitiveType

	switch o := o.(type) {
	case PrimitiveType:
		other = &o
	case *PrimitiveType:
		other = o
	default:
		return false
	}

	if t.Flags != other.Flags {
		return false
	}

	return true
}

func (t *PrimitiveType) String() string {
	return "PrimitiveType { primitive=" + t.Name + " }"
}

type ObjectType struct {
	util.DefaultBase[IType]

	InternalName  string
	QualifiedName string
	Name          string
	TypeArguments ITypeArgument
	Dimension     int
	Descriptor    string
}

func (t *ObjectType) GetDimension() int {
	return t.Dimension
}

func (t *ObjectType) CreateType(dimension int) IType {
	if t.Dimension == dimension {
		return t
	} else if t.Descriptor[len(t.Descriptor)-1] != ';' {
		if dimension == 0 {
			tmp := GetPrimitiveType(int(t.Descriptor[t.Dimension]))
			return &tmp
		} else {
			tmp := NewObjectTypeWithDescAndDim(t.InternalName, t.Dimension)
			return &tmp
		}
	} else {
		tmp := NewObjectTypeWithAll(t.InternalName, t.QualifiedName, t.Name, t.TypeArguments, dimension)
		return &tmp
	}
}

func (t *ObjectType) IsGenericType() bool {
	return false
}

func (t *ObjectType) IsInnerObjectType() bool {
	return false
}

func (t *ObjectType) IsObjectType() bool {
	return true
}

func (t *ObjectType) IsPrimitiveType() bool {
	return false
}

func (t *ObjectType) IsTypes() bool {
	return false
}

func (t *ObjectType) GetOuterType() *ObjectType {
	return &OtTypeUndefinedObject
}

func (t *ObjectType) GetInternalName() string {
	return t.InternalName
}

func (t *ObjectType) HashCode() int {
	result := 735485092 + hashCodeWithString(t.InternalName)
	result *= 31
	if t.TypeArguments != nil {
		result += t.TypeArguments.HashCode()
	}
	result = 31*result + t.Dimension
	return result
}

func (t *ObjectType) AcceptTypeVisitor(visitor ITypeVisitor) {
	visitor.VisitObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *ObjectType) TypeArgumentFirst() ITypeArgument {
	return t
}

func (t *ObjectType) TypeArgumentList() util.IList[ITypeArgument] {
	return util.NewDefaultListWithElements[ITypeArgument](t)
}

func (t *ObjectType) TypeArgumentSize() int {
	return 1
}

func (t *ObjectType) GetType() IType {
	return &OtTypeUndefinedObject
}

func (t *ObjectType) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	switch meta := typeArgument.(type) {
	case *ObjectType:
		if t.Dimension != meta.Dimension || t.InternalName != meta.GetInternalName() {
			return false
		}

		if meta.TypeArguments == nil {
			return t.TypeArguments == nil
		} else if t.TypeArguments == nil {
			return false
		} else {
			return t.TypeArguments.IsTypeArgumentAssignableFrom(typeBounds, meta.TypeArguments)
		}
	case *InnerObjectType:
		if t.Dimension != meta.Dimension || t.InternalName != meta.GetInternalName() {
			return false
		}

		if meta.TypeArguments == nil {
			return t.TypeArguments == nil
		} else if t.TypeArguments == nil {
			return false
		} else {
			return t.TypeArguments.IsTypeArgumentAssignableFrom(typeBounds, meta.TypeArguments)
		}
	case *GenericType:
		bt := typeBounds[meta.Name]
		ot, ok := bt.(*ObjectType)

		if ok {
			if t.InternalName == ot.GetInternalName() {
				return true
			}
		}
	}

	return false
}

func (t *ObjectType) IsTypeArgumentList() bool {
	return false
}

func (t *ObjectType) IsGenericTypeArgument() bool {
	return false
}

func (t *ObjectType) IsInnerObjectTypeArgument() bool {
	return false
}

func (t *ObjectType) IsObjectTypeArgument() bool {
	return true
}

func (t *ObjectType) IsPrimitiveTypeArgument() bool {
	return false
}

func (t *ObjectType) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (t *ObjectType) IsWildcardSuperTypeArgument() bool {
	return false
}

func (t *ObjectType) IsWildcardTypeArgument() bool {
	return false
}

func (t *ObjectType) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *ObjectType) Equals(o interface{}) bool {
	if t == o {
		return true
	}

	if o == nil {
		return false
	}

	var other *ObjectType

	switch o := o.(type) {
	case ObjectType:
		other = &o
	case *ObjectType:
		other = o
	default:
		return false
	}

	if t.Dimension != other.Dimension {
		return false
	}

	if t.InternalName != other.InternalName {
		return false
	}

	if t.InternalName == "jara/lang/Class" {
		wildcard1 := (t.TypeArguments == nil) || (reflect.TypeOf(t.TypeArguments) == reflect.TypeOf(WildcardTypeArgument{}))
		wildcard2 := (other.TypeArguments == nil) || (reflect.TypeOf(other.TypeArguments) == reflect.TypeOf(WildcardTypeArgument{}))

		if wildcard1 || wildcard2 {
			return true
		}
	}

	if t.TypeArguments != nil {
		return t.TypeArguments.Equals(other.TypeArguments)
	}

	return t.TypeArguments == nil
}

func (t *ObjectType) String() string {
	msg := fmt.Sprintf("ObjectType{ %s", t.InternalName)
	if t.TypeArguments != nil {
		msg += fmt.Sprintf("<%s>", t.TypeArguments)
	}
	if t.Dimension > 0 {
		msg += fmt.Sprintf(", %d", t.Dimension)
	}
	msg += " }"
	return msg
}

/////////////////////////////////////////////////////////////////////

func (t *ObjectType) IsTypeArgumentAssignableFromWithObj(typeBounds map[string]IType, objectType ObjectType) bool {
	if t.Dimension != objectType.Dimension || t.InternalName != objectType.GetInternalName() {
		return false
	}

	if objectType.TypeArguments == nil {
		return t.TypeArguments == nil
	} else if t.TypeArguments == nil {
		return false
	} else {
		return t.TypeArguments.IsTypeArgumentAssignableFrom(typeBounds, objectType.TypeArguments)
	}
}

func (t *ObjectType) CreateTypeWithArgs(typeArguments ITypeArgument) IType {
	if t.TypeArguments == typeArguments {
		return t
	} else {
		tmp := NewObjectTypeWithAll(t.InternalName, t.QualifiedName, t.Name, typeArguments, t.Dimension)
		return &tmp
	}
}

type InnerObjectType struct {
	util.DefaultBase[IType]

	InternalName  string
	QualifiedName string
	Name          string
	TypeArguments ITypeArgument
	Dimension     int
	Descriptor    string
	OuterType     *ObjectType
}

func (t *InnerObjectType) GetDimension() int {
	return t.Dimension
}

func (t *InnerObjectType) CreateType(dimension int) IType {
	tmp := NewInnerObjectTypeWithAll(t.InternalName, t.QualifiedName, t.Name, t.TypeArguments, dimension, t.GetOuterType())
	return &tmp
}

func (t *InnerObjectType) IsGenericType() bool {
	return false
}

func (t *InnerObjectType) IsInnerObjectType() bool {
	return true
}

func (t *InnerObjectType) IsObjectType() bool {
	return true
}

func (t *InnerObjectType) IsPrimitiveType() bool {
	return false
}

func (t *InnerObjectType) IsTypes() bool {
	return false
}

func (t *InnerObjectType) GetOuterType() *ObjectType {
	return t.OuterType
}

func (t *InnerObjectType) GetInternalName() string {
	return t.InternalName
}

func (t *InnerObjectType) HashCode() int {
	result := 735485092 + hashCodeWithString(t.InternalName)
	result *= 31
	if t.TypeArguments != nil {
		result += t.TypeArguments.HashCode()
	}
	result = 31*result + t.Dimension
	result = 111476860 + result
	result = 31*result + t.GetOuterType().HashCode()
	return result
}

func (t *InnerObjectType) AcceptTypeVisitor(visitor ITypeVisitor) {
	visitor.VisitInnerObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) TypeArgumentFirst() ITypeArgument {
	return t
}

func (t *InnerObjectType) TypeArgumentList() util.IList[ITypeArgument] {
	return util.NewDefaultListWithElements[ITypeArgument](t)
}

func (t *InnerObjectType) TypeArgumentSize() int {
	return 1
}

func (t *InnerObjectType) GetType() IType {
	return &OtTypeUndefinedObject
}

func (t *InnerObjectType) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	switch meta := typeArgument.(type) {
	case *ObjectType:
		if t.Dimension != meta.Dimension || t.InternalName != meta.GetInternalName() {
			return false
		}

		if meta.TypeArguments == nil {
			return t.TypeArguments == nil
		} else if t.TypeArguments == nil {
			return false
		} else {
			return t.TypeArguments.IsTypeArgumentAssignableFrom(typeBounds, meta.TypeArguments)
		}
	case *InnerObjectType:
		if t.Dimension != meta.Dimension || t.InternalName != meta.GetInternalName() {
			return false
		}

		if meta.TypeArguments == nil {
			return t.TypeArguments == nil
		} else if t.TypeArguments == nil {
			return false
		} else {
			return t.TypeArguments.IsTypeArgumentAssignableFrom(typeBounds, meta.TypeArguments)
		}
	case *GenericType:
		bt := typeBounds[meta.Name]
		ot, ok := bt.(*ObjectType)

		if ok {
			if t.InternalName == ot.GetInternalName() {
				return true
			}
		}
	}

	return false
}

func (t *InnerObjectType) IsTypeArgumentList() bool {
	return false
}

func (t *InnerObjectType) IsGenericTypeArgument() bool {
	return false
}

func (t *InnerObjectType) IsInnerObjectTypeArgument() bool {
	return true
}

func (t *InnerObjectType) IsObjectTypeArgument() bool {
	return true
}

func (t *InnerObjectType) IsPrimitiveTypeArgument() bool {
	return false
}

func (t *InnerObjectType) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (t *InnerObjectType) IsWildcardSuperTypeArgument() bool {
	return false
}

func (t *InnerObjectType) IsWildcardTypeArgument() bool {
	return false
}

func (t *InnerObjectType) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitInnerObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) Equals(o interface{}) bool {
	if t == o {
		return true
	}

	if o == nil {
		return false
	}

	var other *InnerObjectType

	switch o := o.(type) {
	case InnerObjectType:
		other = &o
	case *InnerObjectType:
		other = o
	default:
		return false
	}

	if !t.OuterType.Equals(other.Dimension) {
		return false
	}

	return true
}

func (t *InnerObjectType) String() string {
	if t.TypeArguments == nil {
		return fmt.Sprintf("InnerObjectType { %s.%s }", t.OuterType, t.Descriptor)
	} else {
		return fmt.Sprintf("InnerObjectType { %s.%s<%s> }", t.OuterType, t.Descriptor, t.TypeArguments)
	}
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) CreateTypeWithArg(typeArguments ITypeArgument) IType {
	tmp := NewInnerObjectTypeWithAll(t.InternalName, t.QualifiedName, t.Name, typeArguments, t.Dimension, t.OuterType)
	return &tmp
}

type Types struct {
	util.DefaultList[IType]
}

func (t *Types) GetDimension() int {
	return -1
}

func (t *Types) CreateType(_ int) IType {
	return nil
}

func (t *Types) IsGenericType() bool {
	return false
}

func (t *Types) IsInnerObjectType() bool {
	return false
}

func (t *Types) IsObjectType() bool {
	return false
}

func (t *Types) IsPrimitiveType() bool {
	return false
}

func (t *Types) IsTypes() bool {
	return true
}

func (t *Types) GetOuterType() *ObjectType {
	return &OtTypeUndefinedObject
}

func (t *Types) GetInternalName() string {
	return ""
}

func (t *Types) AcceptTypeVisitor(visitor ITypeVisitor) {
	visitor.VisitTypes(t)
}

func (t *Types) Equals(o interface{}) bool {
	if o == t {
		return true
	}

	if o == nil {
		return false
	}

	var other *Types

	switch o := o.(type) {
	case Types:
		other = &o
	case *Types:
		other = o
	default:
		return false
	}

	size := t.Size()
	equals := size == other.Size()

	if equals {
		es := t.ToSlice()
		otherEs := other.ToSlice()
		for i := 0; i < size; i++ {
			if es[i] != otherEs[i] {
				return false
			}
		}
	}

	return equals
}

func (t *Types) HashCode() int {
	result := 1
	for _, item := range t.ToSlice() {
		result = 31*result + item.HashCode()
	}
	return result
}

func (t *Types) String() string {
	return "Types {}"
}

/////////////////////////////////////////////////////////////////////

type GenericType struct {
	util.DefaultBase[IType]
	Name       string
	Descriptor string
	Dimension  int
}

func (t *GenericType) GetDimension() int {
	return t.Dimension
}

func (t *GenericType) CreateType(dimension int) IType {
	if t.Dimension == dimension {
		return t
	} else {
		tmp := NewGenericTypeWithAll(t.Name, dimension)
		return &tmp
	}
}

func (t *GenericType) IsGenericType() bool {
	return true
}

func (t *GenericType) IsInnerObjectType() bool {
	return false
}

func (t *GenericType) IsObjectType() bool {
	return false
}

func (t *GenericType) IsPrimitiveType() bool {
	return false
}

func (t *GenericType) IsTypes() bool {
	return false
}

func (t *GenericType) GetOuterType() *ObjectType {
	return &OtTypeUndefinedObject
}

func (t *GenericType) GetInternalName() string {
	return ""
}

func (t *GenericType) HashCode() int {
	result := 991890290 + hashCodeWithString(t.Name)
	result = 31*result + t.Dimension
	return result
}

func (t *GenericType) AcceptTypeVisitor(visitor ITypeVisitor) {
	visitor.VisitGenericType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *GenericType) TypeArgumentFirst() ITypeArgument {
	return t
}

func (t *GenericType) TypeArgumentList() util.IList[ITypeArgument] {
	return util.NewDefaultListWithElements[ITypeArgument](t)
}

func (t *GenericType) TypeArgumentSize() int {
	return 1
}

func (t *GenericType) GetType() IType {
	return &OtTypeUndefinedObject
}

func (t *GenericType) IsTypeArgumentAssignableFrom(_ map[string]IType, typeArgument ITypeArgument) bool {
	if o, ok := typeArgument.(*GenericType); ok {
		return t.Equals(o)
	}

	return false
}

func (t *GenericType) IsTypeArgumentList() bool {
	return false
}

func (t *GenericType) IsGenericTypeArgument() bool {
	return true
}

func (t *GenericType) IsInnerObjectTypeArgument() bool {
	return false
}

func (t *GenericType) IsObjectTypeArgument() bool {
	return false
}

func (t *GenericType) IsPrimitiveTypeArgument() bool {
	return false
}

func (t *GenericType) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (t *GenericType) IsWildcardSuperTypeArgument() bool {
	return false
}

func (t *GenericType) IsWildcardTypeArgument() bool {
	return false
}

func (t *GenericType) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitGenericType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *GenericType) Equals(o interface{}) bool {
	if o == t {
		return true
	}

	if o == nil {
		return false
	}

	var other *GenericType

	switch o := o.(type) {
	case GenericType:
		other = &o
	case *GenericType:
		other = o
	default:
		return false
	}

	if t.Dimension != other.Dimension {
		return false
	}
	if t.Name != other.Name {
		return false
	}
	return true
}

func (t *GenericType) String() string {
	msg := fmt.Sprintf("GenericType{ %s", t.Name)
	if t.Dimension > 0 {
		msg += fmt.Sprintf(", dimension: %d", t.Dimension)
	}
	msg += "}"

	return msg
}

type UnmodifiableTypes struct {
	util.DefaultList[IType]
}

func (t *UnmodifiableTypes) GetDimension() int {
	return -1
}

func (t *UnmodifiableTypes) CreateType(_ int) IType {
	return nil
}

func (t *UnmodifiableTypes) IsGenericType() bool {
	return false
}

func (t *UnmodifiableTypes) IsInnerObjectType() bool {
	return false
}

func (t *UnmodifiableTypes) IsObjectType() bool {
	return false
}

func (t *UnmodifiableTypes) IsPrimitiveType() bool {
	return false
}

func (t *UnmodifiableTypes) IsTypes() bool {
	return true
}

func (t *UnmodifiableTypes) GetOuterType() *ObjectType {
	return &OtTypeUndefinedObject
}

func (t *UnmodifiableTypes) GetInternalName() string {
	return ""
}

func (t *UnmodifiableTypes) AcceptTypeVisitor(visitor ITypeVisitor) {
	visitor.VisitTypes((*Types)(t))
}

func (t *UnmodifiableTypes) Equals(o interface{}) bool {
	if o == t {
		return true
	}

	if o == nil {
		return false
	}

	var other *Types

	switch o := o.(type) {
	case Types:
		other = &o
	case *Types:
		other = o
	default:
		return false
	}

	size := t.Size()
	equals := size == other.Size()

	if equals {
		es := t.ToSlice()
		otherEs := other.ToSlice()
		for i := 0; i < size; i++ {
			if es[i] != otherEs[i] {
				return false
			}
		}
	}

	return equals
}

func (t *UnmodifiableTypes) HashCode() int {
	result := 1
	for _, item := range t.ToSlice() {
		result = 31*result + item.HashCode()
	}
	return result
}

func (t *UnmodifiableTypes) String() string {
	return "UnmodifiableTypes{}"
}
