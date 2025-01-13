package classfile

import "fmt"

type TAG byte

const (
	ConstTagUnknown            TAG = 0
	ConstTagUtf8                   = 1
	ConstTagInteger                = 3
	ConstTagFloat                  = 4
	ConstTagLong                   = 5
	ConstTagDouble                 = 6
	ConstTagClass                  = 7
	ConstTagString                 = 8
	ConstTagFieldRef               = 9
	ConstTagMethodRef              = 10
	ConstTagInterfaceMethodRef     = 11
	ConstTagNameAndType            = 12
	ConstTagMethodHandle           = 15
	ConstTagMethodType             = 16
	ConstTagInvokeDynamic          = 18
	ConstTagMemberRef              = 19
)

func NewConstantClass(nameIndex int) ConstantClass {
	return ConstantClass{
		Tag:       ConstTagClass,
		NameIndex: nameIndex,
	}
}

func NewConstantDouble(value float64) ConstantDouble {
	return ConstantDouble{
		Tag:   ConstTagDouble,
		Value: value,
	}
}

func NewConstantFloat(value float32) ConstantFloat {
	return ConstantFloat{
		Tag:   ConstTagFloat,
		Value: value,
	}
}

func NewConstantInteger(value int) ConstantInteger {
	return ConstantInteger{
		Tag:   ConstTagInteger,
		Value: value,
	}
}

func NewConstantLong(value int64) ConstantLong {
	return ConstantLong{
		Tag:   ConstTagLong,
		Value: value,
	}
}

func NewConstantMemberRef(classIndex int, nameAndTypeIndex int) ConstantMemberRef {
	return ConstantMemberRef{
		Tag:              ConstTagMemberRef,
		ClassIndex:       classIndex,
		NameAndTypeIndex: nameAndTypeIndex,
	}
}

func NewConstantMethodHandle(referenceKind int, referenceIndex int) ConstantMethodHandle {
	return ConstantMethodHandle{
		Tag:            ConstTagMethodHandle,
		ReferenceKind:  referenceKind,
		ReferenceIndex: referenceIndex,
	}
}

func NewConstantMethodType(descriptorIndex int) ConstantMethodType {
	return ConstantMethodType{
		Tag:             ConstTagMethodType,
		DescriptorIndex: descriptorIndex,
	}
}

func NewConstantNameAndType(nameIndex int, descriptorIndex int) ConstantNameAndType {
	return ConstantNameAndType{
		Tag:             ConstTagNameAndType,
		NameIndex:       nameIndex,
		DescriptorIndex: descriptorIndex,
	}
}

func NewConstantString(stringIndex int) ConstantString {
	return ConstantString{
		Tag:         ConstTagString,
		StringIndex: stringIndex,
	}
}

func NewConstantUtf8(value string) ConstantUtf8 {
	return ConstantUtf8{
		Tag:   ConstTagUtf8,
		Value: value,
	}
}

type IConstant interface {
	GetTag() TAG
	String() string
}

type IConstantValue interface {
	GetTag() TAG
	GetValue() interface{}
	String() string
}

type ConstantClass struct {
	Tag       TAG
	NameIndex int
}

func (c ConstantClass) GetTag() TAG {
	return c.Tag
}

func (c ConstantClass) String() string {
	return fmt.Sprintf("ConstantClass{ tag: %d, nameIndex: %d }", c.Tag, c.NameIndex)
}

type ConstantDouble struct {
	Tag   TAG
	Value float64
}

func (c ConstantDouble) GetTag() TAG {
	return c.Tag
}

func (c ConstantDouble) GetValue() interface{} {
	return c.Value
}

func (c ConstantDouble) String() string {
	return fmt.Sprintf("ConstantDouble{ tag: %d, value: %f }", c.Tag, c.Value)
}

type ConstantFloat struct {
	Tag   TAG
	Value float32
}

func (c ConstantFloat) GetTag() TAG {
	return c.Tag
}

func (c ConstantFloat) GetValue() interface{} {
	return c.Value
}

func (c ConstantFloat) String() string {
	return fmt.Sprintf("ConstantClass{ tag: %d, value: %f }", c.Tag, c.Value)
}

type ConstantInteger struct {
	Tag   TAG
	Value int
}

func (c ConstantInteger) GetTag() TAG {
	return c.Tag
}

func (c ConstantInteger) GetValue() interface{} {
	return c.Value
}

func (c ConstantInteger) String() string {
	return fmt.Sprintf("ConstantInteger{ tag: %d, value: %d }", c.Tag, c.Value)
}

type ConstantLong struct {
	Tag   TAG
	Value int64
}

func (c ConstantLong) GetTag() TAG {
	return c.Tag
}

func (c ConstantLong) GetValue() interface{} {
	return c.Value
}

func (c ConstantLong) String() string {
	return fmt.Sprintf("ConstantLong{ tag: %d, value: %d }", c.Tag, c.Value)
}

type ConstantMemberRef struct {
	Tag              TAG
	ClassIndex       int
	NameAndTypeIndex int
}

func (c ConstantMemberRef) GetTag() TAG {
	return c.Tag
}

func (c ConstantMemberRef) String() string {
	return fmt.Sprintf("ConstantMemberRef{ tag: %d, nameIndex: %d, nameAndTypeIndex: %d }",
		c.Tag, c.ClassIndex, c.NameAndTypeIndex)
}

type ConstantMethodHandle struct {
	Tag            TAG
	ReferenceKind  int
	ReferenceIndex int
}

func (c ConstantMethodHandle) GetTag() TAG {
	return c.Tag
}

func (c ConstantMethodHandle) String() string {
	return fmt.Sprintf("ConstantMethodHandle{ tag: %d, referenceKind: %d, referenceIndex: %d }",
		c.Tag, c.ReferenceKind, c.ReferenceIndex)
}

type ConstantMethodType struct {
	Tag             TAG
	DescriptorIndex int
}

func (c ConstantMethodType) GetTag() TAG {
	return c.Tag
}

func (c ConstantMethodType) String() string {
	return fmt.Sprintf("ConstantMethodType{ tag: %d, descriptorIndex: %d }", c.Tag, c.DescriptorIndex)
}

type ConstantNameAndType struct {
	Tag             TAG
	NameIndex       int
	DescriptorIndex int
}

func (c ConstantNameAndType) GetTag() TAG {
	return c.Tag
}

func (c ConstantNameAndType) String() string {
	return fmt.Sprintf("ConstantNameAndType{ tag: %d, nameIndex: %d, descriptorIndex: %d }",
		c.Tag, c.NameIndex, c.DescriptorIndex)
}

type ConstantString struct {
	Tag         TAG
	StringIndex int
}

func (c ConstantString) GetTag() TAG {
	return c.Tag
}

func (c ConstantString) String() string {
	return fmt.Sprintf("ConstantString{ tag: %d, nameIndex: %d }", c.Tag, c.StringIndex)
}

type ConstantUtf8 struct {
	Tag   TAG
	Value string
}

func (c ConstantUtf8) GetTag() TAG {
	return c.Tag
}

func (c ConstantUtf8) GetValue() interface{} {
	return c.Value
}

func (c ConstantUtf8) String() string {
	return fmt.Sprintf("ConstantUtf8{ tag: %d, nameIndex: %s }", c.Tag, c.Value)
}
