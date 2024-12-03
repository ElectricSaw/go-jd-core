package _type

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

var PtTypeBoolean = NewPrimitiveType("boolean", intmod.FlagBoolean, intmod.FlagBoolean, intmod.FlagBoolean)
var PtTypeByte = NewPrimitiveType("byte", intmod.FlagByte, intmod.FlagByte, intmod.FlagByte|intmod.FlagInt|intmod.FlagShort)
var PtTypeChar = NewPrimitiveType("char", intmod.FlagChar, intmod.FlagChar, intmod.FlagChar|intmod.FlagInt)
var PtTypeDouble = NewPrimitiveType("double", intmod.FlagDouble, intmod.FlagDouble, intmod.FlagDouble)
var PtTypeFloat = NewPrimitiveType("float", intmod.FlagFloat, intmod.FlagFloat, intmod.FlagFloat)
var PtTypeInt = NewPrimitiveType("int", intmod.FlagInt, intmod.FlagInt|intmod.FlagByte|intmod.FlagChar|intmod.FlagShort, intmod.FlagInt)
var PtTypeLong = NewPrimitiveType("long", intmod.FlagLong, intmod.FlagLong, intmod.FlagLong)
var PtTypeShort = NewPrimitiveType("short", intmod.FlagShort, intmod.FlagShort|intmod.FlagByte, intmod.FlagShort|intmod.FlagInt)
var PtTypeVoid = NewPrimitiveType("void", intmod.FlagVoid, intmod.FlagVoid, intmod.FlagVoid)
var PtMaybeCharType = NewPrimitiveType("maybe_char", intmod.FlagChar|intmod.FlagInt, intmod.FlagChar|intmod.FlagInt, intmod.FlagChar|intmod.FlagInt)                                                                                                                                                                   //  32768 .. 65535
var PtMaybeShortType = NewPrimitiveType("maybe_short", intmod.FlagChar|intmod.FlagShort|intmod.FlagInt, intmod.FlagChar|intmod.FlagShort|intmod.FlagInt, intmod.FlagChar|intmod.FlagShort|intmod.FlagInt)                                                                                                              //    128 .. 32767
var PtMaybeByteType = NewPrimitiveType("maybe_byte", intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt, intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt, intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt)                                                                //      2 .. 127
var PtMaybeBooleanType = NewPrimitiveType("maybe_boolean", intmod.FlagBoolean|intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt, intmod.FlagBoolean|intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt, intmod.FlagBoolean|intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt) //      0 .. 1
var PtMaybeNegativeByteType = NewPrimitiveType("maybe_negative_byte", intmod.FlagByte|intmod.FlagShort|intmod.FlagInt, intmod.FlagByte|intmod.FlagShort|intmod.FlagInt, intmod.FlagByte|intmod.FlagShort|intmod.FlagInt)                                                                                               //   -128 .. -1
var PtMaybeNegativeShortType = NewPrimitiveType("maybe_negative_short", intmod.FlagShort|intmod.FlagInt, intmod.FlagShort|intmod.FlagInt, intmod.FlagShort|intmod.FlagInt)                                                                                                                                             // -32768 .. -129
var PtMaybeIntType = NewPrimitiveType("maybe_int", intmod.FlagInt, intmod.FlagInt, intmod.FlagInt)                                                                                                                                                                                                                     // Otherwise
var PtMaybeNegativeBooleanType = NewPrimitiveType("maybe_negative_boolean", intmod.FlagBoolean|intmod.FlagByte|intmod.FlagShort|intmod.FlagInt, intmod.FlagBoolean|intmod.FlagByte|intmod.FlagShort|intmod.FlagInt, intmod.FlagBoolean|intmod.FlagByte|intmod.FlagShort|intmod.FlagInt)                                // Boolean or negative

var descriptorToType = []intmod.IPrimitiveType{
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

func GetPrimitiveType(primitiveDescriptor int) intmod.IPrimitiveType {
	return descriptorToType[primitiveDescriptor-66] // int('B')
}

func NewPrimitiveType(name string, flags, leftFlags, rightFlags int) intmod.IPrimitiveType {
	t := &PrimitiveType{
		name:       name,
		flags:      flags,
		leftFlags:  leftFlags,
		rightFlags: rightFlags,
	}

	sb := ""

	if flags&intmod.FlagDouble != 0 {
		sb += "D"
	} else if flags&intmod.FlagFloat != 0 {
		sb += "F"
	} else if flags&intmod.FlagLong != 0 {
		sb += "J"
	} else if flags&intmod.FlagBoolean != 0 {
		sb += "Z"
	} else if flags&intmod.FlagByte != 0 {
		sb += "B"
	} else if flags&intmod.FlagChar != 0 {
		sb += "C"
	} else if flags&intmod.FlagShort != 0 {
		sb += "S"
	} else {
		sb += "I"
	}

	t.descriptor = sb
	t.SetValue(t)

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

func (t *PrimitiveType) CreateType(dimension int) intmod.IType {
	if dimension == 0 {
		return t
	} else {
		return NewObjectTypeWithDescAndDim(t.descriptor, dimension).(intmod.IType)
	}
}

func (t *PrimitiveType) IsPrimitiveType() bool {
	return true
}

func (t *PrimitiveType) AcceptTypeVisitor(visitor intmod.ITypeVisitor) {
	visitor.VisitPrimitiveType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *PrimitiveType) IsTypeArgumentAssignableFrom(_ map[string]intmod.IType, typeArgument intmod.ITypeArgument) bool {
	return t.Equals(typeArgument)
}

func (t *PrimitiveType) IsPrimitiveTypeArgument() bool {
	return true
}

func (t *PrimitiveType) AcceptTypeArgumentVisitor(visitor intmod.ITypeArgumentVisitor) {
	visitor.VisitPrimitiveType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *PrimitiveType) Equals(o intmod.ITypeArgument) bool {
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
	if t.flags&intmod.FlagBoolean != 0 {
		return intmod.FlagBoolean
	} else if t.flags&intmod.FlagInt != 0 {
		return intmod.FlagInt
	} else if t.flags&intmod.FlagChar != 0 {
		return intmod.FlagChar
	} else if t.flags&intmod.FlagShort != 0 {
		return intmod.FlagShort
	} else if t.flags&intmod.FlagByte != 0 {
		return intmod.FlagByte
	}
	return t.flags
}

func (t *PrimitiveType) String() string {
	return "PrimitiveType { primitive=" + t.name + " }"
}
