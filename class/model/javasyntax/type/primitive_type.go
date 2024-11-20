package _type

import intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

var PtTypeBoolean = NewPrimitiveType("boolean", intsyn.FlagBoolean, intsyn.FlagBoolean, intsyn.FlagBoolean)
var PtTypeByte = NewPrimitiveType("byte", intsyn.FlagByte, intsyn.FlagByte, intsyn.FlagByte|intsyn.FlagInt|intsyn.FlagShort)
var PtTypeChar = NewPrimitiveType("char", intsyn.FlagChar, intsyn.FlagChar, intsyn.FlagChar|intsyn.FlagInt)
var PtTypeDouble = NewPrimitiveType("double", intsyn.FlagDouble, intsyn.FlagDouble, intsyn.FlagDouble)
var PtTypeFloat = NewPrimitiveType("float", intsyn.FlagFloat, intsyn.FlagFloat, intsyn.FlagFloat)
var PtTypeInt = NewPrimitiveType("int", intsyn.FlagInt, intsyn.FlagInt|intsyn.FlagByte|intsyn.FlagChar|intsyn.FlagShort, intsyn.FlagInt)
var PtTypeLong = NewPrimitiveType("long", intsyn.FlagLong, intsyn.FlagLong, intsyn.FlagLong)
var PtTypeShort = NewPrimitiveType("short", intsyn.FlagShort, intsyn.FlagShort|intsyn.FlagByte, intsyn.FlagShort|intsyn.FlagInt)
var PtTypeVoid = NewPrimitiveType("void", intsyn.FlagVoid, intsyn.FlagVoid, intsyn.FlagVoid)
var PtMaybeCharType = NewPrimitiveType("maybe_char", intsyn.FlagChar|intsyn.FlagInt, intsyn.FlagChar|intsyn.FlagInt, intsyn.FlagChar|intsyn.FlagInt)                                                                                                                                                                   //  32768 .. 65535
var PtMaybeShortType = NewPrimitiveType("maybe_short", intsyn.FlagChar|intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagChar|intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagChar|intsyn.FlagShort|intsyn.FlagInt)                                                                                                              //    128 .. 32767
var PtMaybeByteType = NewPrimitiveType("maybe_byte", intsyn.FlagByte|intsyn.FlagChar|intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagByte|intsyn.FlagChar|intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagByte|intsyn.FlagChar|intsyn.FlagShort|intsyn.FlagInt)                                                                //      2 .. 127
var PtMaybeBooleanType = NewPrimitiveType("maybe_boolean", intsyn.FlagBoolean|intsyn.FlagByte|intsyn.FlagChar|intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagBoolean|intsyn.FlagByte|intsyn.FlagChar|intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagBoolean|intsyn.FlagByte|intsyn.FlagChar|intsyn.FlagShort|intsyn.FlagInt) //      0 .. 1
var PtMaybeNegativeByteType = NewPrimitiveType("maybe_negative_byte", intsyn.FlagByte|intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagByte|intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagByte|intsyn.FlagShort|intsyn.FlagInt)                                                                                               //   -128 .. -1
var PtMaybeNegativeShortType = NewPrimitiveType("maybe_negative_short", intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagShort|intsyn.FlagInt)                                                                                                                                             // -32768 .. -129
var PtMaybeIntType = NewPrimitiveType("maybe_int", intsyn.FlagInt, intsyn.FlagInt, intsyn.FlagInt)                                                                                                                                                                                                                     // Otherwise
var PtMaybeNegativeBooleanType = NewPrimitiveType("maybe_negative_boolean", intsyn.FlagBoolean|intsyn.FlagByte|intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagBoolean|intsyn.FlagByte|intsyn.FlagShort|intsyn.FlagInt, intsyn.FlagBoolean|intsyn.FlagByte|intsyn.FlagShort|intsyn.FlagInt)                                // Boolean or negative

var descriptorToType = []intsyn.IPrimitiveType{
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

func GetPrimitiveType(primitiveDescriptor int) intsyn.IPrimitiveType {
	return descriptorToType[primitiveDescriptor-66] // int('B')
}

func NewPrimitiveType(name string, flags, leftFlags, rightFlags int) intsyn.IPrimitiveType {
	t := &PrimitiveType{
		name:       name,
		flags:      flags,
		leftFlags:  leftFlags,
		rightFlags: rightFlags,
	}

	sb := ""

	if flags&intsyn.FlagDouble != 0 {
		sb += "D"
	} else if flags&intsyn.FlagFloat != 0 {
		sb += "F"
	} else if flags&intsyn.FlagLong != 0 {
		sb += "J"
	} else if flags&intsyn.FlagBoolean != 0 {
		sb += "Z"
	} else if flags&intsyn.FlagByte != 0 {
		sb += "B"
	} else if flags&intsyn.FlagChar != 0 {
		sb += "C"
	} else if flags&intsyn.FlagShort != 0 {
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

func (t *PrimitiveType) HashCode() int {
	return 750039781 + t.flags
}

/////////////////////////////////////////////////////////////////////

func (t *PrimitiveType) Name() string {
	return t.name
}

func (t *PrimitiveType) Descriptor() string {
	return t.descriptor
}

func (t *PrimitiveType) Dimension() int {
	return 0
}

func (t *PrimitiveType) CreateType(dimension int) intsyn.IType {
	if dimension == 0 {
		return t
	} else {
		return NewObjectTypeWithDescAndDim(t.descriptor, dimension).(intsyn.IType)
	}
}

func (t *PrimitiveType) IsPrimitiveType() bool {
	return true
}

func (t *PrimitiveType) AcceptTypeVisitor(visitor intsyn.ITypeVisitor) {
	visitor.VisitPrimitiveType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *PrimitiveType) IsTypeArgumentAssignableFrom(_ map[string]intsyn.IType, typeArgument intsyn.ITypeArgument) bool {
	return t.Equals(typeArgument)
}

func (t *PrimitiveType) IsPrimitiveTypeArgument() bool {
	return true
}

func (t *PrimitiveType) AcceptTypeArgumentVisitor(visitor intsyn.ITypeArgumentVisitor) {
	visitor.VisitPrimitiveType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *PrimitiveType) Equals(o intsyn.ITypeArgument) bool {
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

func (t *PrimitiveType) Flags() int {
	return t.flags
}

func (t *PrimitiveType) LeftFlags() int {
	return t.leftFlags
}

func (t *PrimitiveType) RightFlags() int {
	return t.rightFlags
}

func (t *PrimitiveType) JavaPrimitiveFlags() int {
	if t.flags&intsyn.FlagBoolean != 0 {
		return intsyn.FlagBoolean
	} else if t.flags&intsyn.FlagInt != 0 {
		return intsyn.FlagInt
	} else if t.flags&intsyn.FlagChar != 0 {
		return intsyn.FlagChar
	} else if t.flags&intsyn.FlagShort != 0 {
		return intsyn.FlagShort
	} else if t.flags&intsyn.FlagByte != 0 {
		return intsyn.FlagByte
	}
	return t.flags
}

func (t *PrimitiveType) String() string {
	return "PrimitiveType { primitive=" + t.name + " }"
}
