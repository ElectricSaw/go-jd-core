package utils

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"math"
)

func GetPrimitiveTypeFromDescriptor(descriptor string) intmod.IType {
	dimension := 0

	for descriptor[dimension] == '[' {
		dimension++
	}

	if dimension == 0 {
		return _type.GetPrimitiveType(int(descriptor[dimension])).(intmod.IType)
	} else {
		return _type.NewObjectTypeWithDescAndDim(descriptor[dimension:], dimension).(intmod.IType)
	}
}

func GetPrimitiveTypeFromValue(value int) intmod.IPrimitiveType {
	if value >= 0 {
		if value <= 1 {
			return _type.PtMaybeBooleanType.(intmod.IPrimitiveType)
		}
		if value <= math.MaxInt8 {
			return _type.PtMaybeByteType.(intmod.IPrimitiveType)
		}
		if value <= math.MaxInt16 {
			return _type.PtMaybeShortType.(intmod.IPrimitiveType)
		}
		if value <= '\uFFFF' {
			return _type.PtMaybeCharType.(intmod.IPrimitiveType)
		}
	} else {
		if value >= math.MinInt8 {
			return _type.PtMaybeNegativeByteType.(intmod.IPrimitiveType)
		}
		if value >= math.MinInt16 {
			return _type.PtMaybeNegativeShortType.(intmod.IPrimitiveType)
		}
	}
	return _type.PtMaybeIntType
}

func GetCommonPrimitiveType(pt1, pt2 intmod.IPrimitiveType) intmod.IPrimitiveType {
	return GetPrimitiveTypeFromFlags(pt1.Flags() & pt2.Flags())
}

func GetPrimitiveTypeFromFlags(flags int) intmod.IPrimitiveType {
	switch flags {
	case intmod.FlagBoolean:
		return _type.PtTypeBoolean
	case intmod.FlagChar:
		return _type.PtTypeChar
	case intmod.FlagFloat:
		return _type.PtTypeFloat
	case intmod.FlagDouble:
		return _type.PtTypeDouble
	case intmod.FlagByte:
		return _type.PtTypeByte
	case intmod.FlagShort:
		return _type.PtTypeShort
	case intmod.FlagInt:
		return _type.PtTypeInt
	case intmod.FlagLong:
		return _type.PtTypeLong
	case intmod.FlagVoid:
		return _type.PtTypeVoid
	default:
		if flags == intmod.FlagChar|intmod.FlagInt {
			return _type.PtMaybeCharType
		}
		if flags == intmod.FlagChar|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeShortType
		}
		if flags == intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeByteType
		}
		if flags == intmod.FlagBoolean|intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeBooleanType
		}
		if flags == intmod.FlagByte|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeNegativeByteType
		}
		if flags == intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeNegativeShortType
		}
		if flags == intmod.FlagBoolean|intmod.FlagByte|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeNegativeBooleanType
		}
	}

	return nil
}

func GetPrimitiveTypeFromTag(tag int) intmod.IType {
	switch tag {
	case 4:
		return _type.PtTypeBoolean.(intmod.IType)
	case 5:
		return _type.PtTypeChar.(intmod.IType)
	case 6:
		return _type.PtTypeFloat.(intmod.IType)
	case 7:
		return _type.PtTypeDouble.(intmod.IType)
	case 8:
		return _type.PtTypeByte.(intmod.IType)
	case 9:
		return _type.PtTypeShort.(intmod.IType)
	case 10:
		return _type.PtTypeInt.(intmod.IType)
	case 11:
		return _type.PtTypeLong.(intmod.IType)
	default:
		return nil
	}
}
