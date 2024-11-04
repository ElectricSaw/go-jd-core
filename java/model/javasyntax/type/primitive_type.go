package _type

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

var PtTypeBoolean = NewPrimitiveType("boolean", FlagBoolean, FlagBoolean, FlagBoolean)
var PtTypeByte = NewPrimitiveType("byte", FlagByte, FlagByte, FlagByte|FlagInt|FlagShort)
var PtTypeChar = NewPrimitiveType("char", FlagChar, FlagChar, FlagChar|FlagInt)
var PtTypeDouble = NewPrimitiveType("double", FlagDouble, FlagDouble, FlagDouble)
var PtTypeFloat = NewPrimitiveType("float", FlagFloat, FlagFloat, FlagFloat)
var PtTypeInt = NewPrimitiveType("int", FlagInt, FlagInt|FlagByte|FlagChar|FlagShort, FlagInt)
var PtTypeLong = NewPrimitiveType("long", FlagLong, FlagLong, FlagLong)
var PtTypeShort = NewPrimitiveType("short", FlagShort, FlagShort|FlagByte, FlagShort|FlagInt)
var PtTypeVoid = NewPrimitiveType("void", FlagVoid, FlagVoid, FlagVoid)
var PtMaybeCharType = NewPrimitiveType("maybe_char", FlagChar|FlagInt, FlagChar|FlagInt, FlagChar|FlagInt)                                                                                                    //  32768 .. 65535
var PtMaybeShortType = NewPrimitiveType("maybe_short", FlagChar|FlagShort|FlagInt, FlagChar|FlagShort|FlagInt, FlagChar|FlagShort|FlagInt)                                                                    //    128 .. 32767
var PtMaybeByteType = NewPrimitiveType("maybe_byte", FlagByte|FlagChar|FlagShort|FlagInt, FlagByte|FlagChar|FlagShort|FlagInt, FlagByte|FlagChar|FlagShort|FlagInt)                                           //      2 .. 127
var PtMaybeBooleanType = NewPrimitiveType("maybe_boolean", FlagBoolean|FlagByte|FlagChar|FlagShort|FlagInt, FlagBoolean|FlagByte|FlagChar|FlagShort|FlagInt, FlagBoolean|FlagByte|FlagChar|FlagShort|FlagInt) //      0 .. 1
var PtMaybeNegativeByteType = NewPrimitiveType("maybe_negative_byte", FlagByte|FlagShort|FlagInt, FlagByte|FlagShort|FlagInt, FlagByte|FlagShort|FlagInt)                                                     //   -128 .. -1
var PtMaybeNegativeShortType = NewPrimitiveType("maybe_negative_short", FlagShort|FlagInt, FlagShort|FlagInt, FlagShort|FlagInt)                                                                              // -32768 .. -129
var PtMaybeIntType = NewPrimitiveType("maybe_int", FlagInt, FlagInt, FlagInt)                                                                                                                                 // Otherwise
var PtMaybeNegativeBooleanType = NewPrimitiveType("maybe_negative_boolean", FlagBoolean|FlagByte|FlagShort|FlagInt, FlagBoolean|FlagByte|FlagShort|FlagInt, FlagBoolean|FlagByte|FlagShort|FlagInt)           // Boolean or negative

var descriptorToType = []*PrimitiveType{
	PtTypeByte,
	PtTypeChar,
	PtTypeDouble,
	PtTypeFloat,
	PtTypeInt,
	PtTypeLong,
	PtTypeShort,
	PtTypeVoid,
	PtTypeBoolean,
}

func GetPrimitiveType(primitiveDescriptor int) *PrimitiveType {
	return descriptorToType[primitiveDescriptor-66] // int('B')
}

func NewPrimitiveType(name string, flags, leftFlags, rightFlags int) *PrimitiveType {
	t := &PrimitiveType{
		name:       name,
		flags:      flags,
		leftFlags:  leftFlags,
		rightFlags: rightFlags,
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

	t.descriptor = sb

	return t
}

type PrimitiveType struct {
	AbstractType
	AbstractTypeArgument

	name       string
	flags      int
	leftFlags  int
	rightFlags int
	descriptor string
}

/////////////////////////////////////////////////////////////////////

func (t *PrimitiveType) GetName() string {
	return t.name
}

func (t *PrimitiveType) GetDescriptor() string {
	return t.descriptor
}

func (t *PrimitiveType) GetDimension() int {
	return 0
}

func (t *PrimitiveType) CreateType(dimension int) IType {
	if dimension == 0 {
		return t
	} else {
		return NewObjectTypeWithDescAndDim(t.descriptor, dimension)
	}
}

func (t *PrimitiveType) IsPrimitiveType() bool {
	return true
}

func (t *PrimitiveType) AcceptTypeVisitor(visitor TypeVisitor) {
	visitor.VisitPrimitiveType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *PrimitiveType) IsTypeArgumentAssignableFrom(_ map[string]IType, typeArgument ITypeArgument) bool {
	return t.Equals(typeArgument)
}

func (t *PrimitiveType) IsPrimitiveTypeArgument() bool {
	return true
}

func (t *PrimitiveType) AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor) {
	visitor.VisitPrimitiveType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *PrimitiveType) Equals(o ITypeArgument) bool {
	if t == o {
		return true
	}

	if o == nil {
		return false
	}

	that, ok := o.(*PrimitiveType)
	if !ok {
		return false
	}

	if t.flags != that.flags {
		return false
	}

	return true
}

func (t *PrimitiveType) GetFlags() int {
	return t.flags
}

func (t *PrimitiveType) GetLeftFlags() int {
	return t.leftFlags
}

func (t *PrimitiveType) GetRightFlags() int {
	return t.rightFlags
}

func (t *PrimitiveType) GetJavaPrimitiveFlags() int {
	if t.flags&FlagBoolean != 0 {
		return FlagBoolean
	} else if t.flags&FlagInt != 0 {
		return FlagInt
	} else if t.flags&FlagChar != 0 {
		return FlagChar
	} else if t.flags&FlagShort != 0 {
		return FlagShort
	} else if t.flags&FlagByte != 0 {
		return FlagByte
	}
	return t.flags
}

func (t *PrimitiveType) String() string {
	return "PrimitiveType { primitive=" + t.name + " }"
}
