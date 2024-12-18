package classpath

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

type IConstant interface {
	Tag() TAG
}

type IConstantValue interface {
	IConstant
	IsConstantValue() bool
}

type IConstantClass interface {
	Tag() TAG
	NameIndex() int
}

type IConstantDouble interface {
	IConstantValue

	Value() float64
}

type IConstantFloat interface {
	IConstantValue

	Value() float32
}

type IConstantInteger interface {
	IConstant

	Value() int
}

type IConstantLong interface {
	IConstantValue

	Value() int64
}

type IConstantMemberRef interface {
	IConstant
	ClassIndex() int
	NameAndTypeIndex() int
}

type IConstantMethodHandle interface {
	IConstant
	ReferenceKind() int
	ReferenceIndex() int
}

type IConstantMethodType interface {
	IConstant
	DescriptorIndex() int
}

type IConstantNameAndType interface {
	IConstant
	NameIndex() int
	DescriptorIndex() int
}

type IConstantString interface {
	IConstant
	StringIndex() int
}

type IConstantUtf8 interface {
	IConstantValue

	Value() string
}

type IConstantPool interface {
	Constant(index int) IConstant
	ConstantTypeName(index int) (string, bool)
	ConstantString(index int) (string, bool)
	ConstantUtf8(index int) (string, bool)
	ConstantValue(index int) IConstantValue
	String() string
}
