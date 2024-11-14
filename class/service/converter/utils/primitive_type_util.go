package utils

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"math"
)

func GetPrimitiveTypeFromDescriptor(descriptor string) _type.IType {
	dimension := 0

	for descriptor[dimension] == '[' {
		dimension++
	}

	if dimension == 0 {
		return _type.GetPrimitiveType(int(descriptor[dimension]))
	} else {
		return _type.NewObjectTypeWithDescAndDim(descriptor[dimension:], dimension)
	}
}

func GetPrimitiveTypeFromValue(value int) *_type.PrimitiveType {
	if value >= 0 {
		if value <= 1 {
			return _type.PtMaybeBooleanType
		}
		if value <= math.MaxInt8 {
			return _type.PtMaybeByteType
		}
		if value <= math.MaxInt16 {
			return _type.PtMaybeShortType
		}
		if value <= '\uFFFF' {
			return _type.PtMaybeCharType
		}
	} else {
		if value >= math.MinInt8 {
			return _type.PtMaybeNegativeByteType
		}
		if value >= math.MinInt16 {
			return _type.PtMaybeNegativeShortType
		}
	}
	return _type.PtMaybeIntType
}

func GetCommonPrimitiveType(pt1, pt2 *_type.PrimitiveType) *_type.PrimitiveType {
	return GetPrimitiveTypeFromFlags(pt1.Flags() & pt2.Flags())
}

func GetPrimitiveTypeFromFlags(flags int) *_type.PrimitiveType {
	switch flags {
	case _type.FlagBoolean:
		return _type.PtTypeBoolean
	case _type.FlagChar:
		return _type.PtTypeChar
	case _type.FlagFloat:
		return _type.PtTypeFloat
	case _type.FlagDouble:
		return _type.PtTypeDouble
	case _type.FlagByte:
		return _type.PtTypeByte
	case _type.FlagShort:
		return _type.PtTypeShort
	case _type.FlagInt:
		return _type.PtTypeInt
	case _type.FlagLong:
		return _type.PtTypeLong
	case _type.FlagVoid:
		return _type.PtTypeVoid
	default:
		if flags == _type.FlagChar|_type.FlagInt {
			return _type.PtMaybeCharType
		}
		if flags == _type.FlagChar|_type.FlagShort|_type.FlagInt {
			return _type.PtMaybeShortType
		}
		if flags == _type.FlagByte|_type.FlagChar|_type.FlagShort|_type.FlagInt {
			return _type.PtMaybeByteType
		}
		if flags == _type.FlagBoolean|_type.FlagByte|_type.FlagChar|_type.FlagShort|_type.FlagInt {
			return _type.PtMaybeBooleanType
		}
		if flags == _type.FlagByte|_type.FlagShort|_type.FlagInt {
			return _type.PtMaybeNegativeByteType
		}
		if flags == _type.FlagShort|_type.FlagInt {
			return _type.PtMaybeNegativeShortType
		}
		if flags == _type.FlagBoolean|_type.FlagByte|_type.FlagShort|_type.FlagInt {
			return _type.PtMaybeNegativeBooleanType
		}
	}

	return nil
}

func GetPrimitiveTypeFromTag(tag int) _type.IType {
	switch tag {
	case 4:
		return _type.PtTypeBoolean
	case 5:
		return _type.PtTypeChar
	case 6:
		return _type.PtTypeFloat
	case 7:
		return _type.PtTypeDouble
	case 8:
		return _type.PtTypeByte
	case 9:
		return _type.PtTypeShort
	case 10:
		return _type.PtTypeInt
	case 11:
		return _type.PtTypeLong
	default:
		return nil
	}
}
