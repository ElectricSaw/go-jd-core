package constant

type ACC byte

const (
	ConstTagUnknown            ACC = 0
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

type Constant interface {
	Tag() ACC
}
