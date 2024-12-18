package classpath

// Access flags for Class, Field, Method, Nested class, Module, Module Requires, Module Exports, Module Opens
const (
	AccPublic       = 0x0001 // C  F  M  N  .  .  .  .
	AccPrivate      = 0x0002 // .  F  M  N  .  .  .  .
	AccProtected    = 0x0004 // .  F  M  N  .  .  .  .
	AccStatic       = 0x0008 // C  F  M  N  .  .  .  .
	AccFinal        = 0x0010 // C  F  M  N  .  .  .  .
	AccSynchronized = 0x0020 // .  .  M  .  .  .  .  .
	AccSuper        = 0x0020 // C  .  .  .  .  .  .  .
	AccOpen         = 0x0020 // .  .  .  .  Mo .  .  .
	AccTransitive   = 0x0020 // .  .  .  .  .  MR .  .
	AccVolatile     = 0x0040 // .  F  .  .  .  .  .  .
	AccBridge       = 0x0040 // .  .  M  .  .  .  .  .
	AccStaticPhase  = 0x0040 // .  .  .  .  .  MR .  .
	AccTransient    = 0x0080 // .  F  .  .  .  .  .  .
	AccVarArgs      = 0x0080 // .  .  M  .  .  .  .  .
	AccNative       = 0x0100 // .  .  M  .  .  .  .  .
	AccInterface    = 0x0200 // C  .  .  N  .  .  .  .
	AccAbstract     = 0x0400 // C  .  M  N  .  .  .  .
	AccStrict       = 0x0800 // .  .  M  .  .  .  .  .
	AccSynthetic    = 0x1000 // C  F  M  N  Mo MR ME MO
	AccAnnotation   = 0x2000 // C  .  .  N  .  .  .  .
	AccEnum         = 0x4000 // C  F  .  N  .  .  .  .
	AccModule       = 0x8000 // C  .  .  .  .  .  .  .
	AccMandated     = 0x8000 // .  .  .  .  Mo MR ME MO
)

type IClassFile interface {
	MajorVersion() int
	MinorVersion() int
	IsEnum() bool
	IsAnnotation() bool
	IsInterface() bool
	IsModule() bool
	IsStatic() bool
	AccessFlags() int
	SetAccessFlags(accessFlags int)
	InternalTypeName() string
	SuperTypeName() string
	InterfaceTypeNames() []string
	Fields() []IField
	Methods() []IMethod
	Attributes() map[string]IAttribute
	Attribute(name string) IAttribute
	OuterClassFile() IClassFile
	SetOuterClassFile(outerClassFile IClassFile)
	InnerClassFiles() []IClassFile
	SetInnerClassFiles(innerClassFiles []IClassFile)
}

type IField interface {
	AccessFlags() int
	Name() string
	Descriptor() string
	Attributes() map[string]IAttribute
	Attribute(name string) IAttribute
	String() string
}

type IMethod interface {
	AccessFlags() int
	Name() string
	Descriptor() string
	Attributes() map[string]IAttribute
	Attribute(name string) IAttribute
	Constants() IConstantPool
	String() string
}
