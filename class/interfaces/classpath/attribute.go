package classpath

type IAttribute interface {
	IsAttribute() bool
}

type IAnnotation interface {
	Descriptor() string
	ElementValuePairs() []IElementValuePair
}

type IAnnotations interface {
	IAttribute

	Annotations() []IAnnotation
}

type IAttributeAnnotationDefault interface {
	IAttribute

	DefaultValue() IElementValue
}

type IAttributeBootstrapMethods interface {
	IAttribute

	BootstrapMethods() []IBootstrapMethod
}

type IAttributeCode interface {
	IAttribute

	MaxStack() int
	MaxLocals() int
	Code() []byte
	ExceptionTable() []ICodeException
	Attribute(name string) IAttribute
}

type IAttributeConstantValue interface {
	IAttribute

	ConstantValue() IConstantValue
}

type IAttributeDeprecated interface {
	IAttribute

	IsAttributeDeprecated() bool
}

type IAttributeExceptions interface {
	IAttribute

	ExceptionTypeNames() []string
}

type IAttributeInnerClasses interface {
	IAttribute

	InnerClasses() []IInnerClass
}

type IAttributeLineNumberTable interface {
	IAttribute

	LineNumberTable() []ILineNumber
}

type IAttributeLocalVariableTable interface {
	IAttribute

	LocalVariableTable() []ILocalVariable
}

type IAttributeLocalVariableTypeTable interface {
	IAttribute

	LocalVariableTypeTable() []ILocalVariableType
}

type IAttributeMethodParameters interface {
	IAttribute

	Parameters() []IMethodParameter
}

type IAttributeModule interface {
	IAttribute

	Name() string
	Flags() int
	Version() string
	Requires() []IModuleInfo
	Exports() []IPackageInfo
	Opens() []IPackageInfo
	Uses() []string
	Provides() []IServiceInfo
}

type IAttributeModuleMainClass interface {
	IAttribute

	MainClass() IConstantClass
}

type IAttributeModulePackages interface {
	IAttribute

	PackageNames() []string
}

type IAttributeParameterAnnotations interface {
	IAttribute

	ParameterAnnotations() []IAnnotations
}

type IAttributeSignature interface {
	IAttribute

	Signature() string
}

type IAttributeSourceFile interface {
	IAttribute

	SourceFile() string
}

type IAttributeSynthetic interface {
	IAttribute

	IsAttributeSynthetic() bool
}

type IUnknownAttribute interface {
	IAttribute

	IsUnknownAttribute() bool
}

type IBootstrapMethod interface {
	BootstrapMethodRef() int
	BootstrapArguments() []int
}

type ICodeException interface {
	Index() int
	StartPc() int
	EndPc() int
	HandlerPc() int
	CatchType() int
}

type IElementValue interface {
	Accept(attribute IElementValueVisitor)
}

type IElementValueVisitor interface {
	VisitPrimitiveType(elementValue IElementValuePrimitiveType)
	VisitClassInfo(elementValue IElementValueClassInfo)
	VisitAnnotationValue(elementValue IElementValueAnnotationValue)
	VisitEnumConstValue(elementValue IElementValueEnumConstValue)
	VisitArrayValue(elementValue IElementValueArrayValue)
}

type IElementValueAnnotationValue interface {
	IElementValue

	AnnotationValue() IAnnotation
}

type IElementValueArrayValue interface {
	IElementValue

	Values() []IElementValue
}

type IElementValueClassInfo interface {
	IElementValue

	ClassInfo() string
}

type IElementValueEnumConstValue interface {
	IElementValue
	
	Descriptor() string
	ConstName() string
}

type IElementValuePair interface {
	Name() string
	SetName(name string)
	Value() IElementValue
	SetValue(value IElementValue)
}

type IElementValuePrimitiveType interface {
	IElementValue

	Type() int
	SetType(type_ int)
	Value() IConstantValue
	SetValue(constValue IConstantValue)
}

type IInnerClass interface {
	InnerTypeName() string
	OuterTypeName() string
	InnerName() string
	InnerAccessFlags() int
}

type ILineNumber interface {
	StartPc() int
	LineNumber() int
}

type ILocalVariable interface {
	StartPc() int
	Length() int
	Name() string
	Descriptor() string
	Index() int
}

type ILocalVariableType interface {
	StartPc() int
	Length() int
	Name() string
	Signature() string
	Index() int
}

type IMethodParameter interface {
	Name() string
	Access() int
}

type IModuleInfo interface {
	Name() string
	Flags() int
	Version() string
}

type IPackageInfo interface {
	InternalName() string
	Flags() int
	ModuleInfoNames() []string
}

type IServiceInfo interface {
	InterfaceTypeName() string
	ImplementationTypeNames() []string
}
