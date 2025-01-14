package classfile

func NewAnnotation(descriptor string, elementValuePairs []ElementValuePair) Annotation {
	return Annotation{
		Descriptor:        descriptor,
		ElementValuePairs: elementValuePairs,
	}
}

func NewAnnotations(annotations []Annotation) Annotations {
	return Annotations{
		Annotations: annotations,
	}
}

func NewBootstrapMethod(bootstrapMethodRef int, bootstrapArguments []int) BootstrapMethod {
	return BootstrapMethod{
		BootstrapMethodRef: bootstrapMethodRef,
		BootstrapArguments: bootstrapArguments,
	}
}

func NewCodeException(index int, startPc int, endPc int, handlerPc int, catchType int) CodeException {
	return CodeException{
		Index:     index,
		StartPc:   startPc,
		EndPc:     endPc,
		HandlerPc: handlerPc,
		CatchType: catchType,
	}
}

func NewInnerClass(innerTypeName string, outerTypeName string, innerName string, innerAccessFlags int) InnerClass {
	return InnerClass{
		InnerTypeName:    innerTypeName,
		OuterTypeName:    outerTypeName,
		InnerName:        innerName,
		InnerAccessFlags: innerAccessFlags,
	}
}

func NewLineNumber(startPc int, lineNumber int) LineNumber {
	return LineNumber{
		StartPc:    startPc,
		LineNumber: lineNumber,
	}
}

func NewLocalVariable(startPc int, length int, name string, descriptor string, index int) LocalVariable {
	return LocalVariable{
		StartPc:    startPc,
		Length:     length,
		Name:       name,
		Descriptor: descriptor,
		Index:      index,
	}
}

func NewLocalVariableType(startPc int, length int, name string, signature string, index int) LocalVariableType {
	return LocalVariableType{
		StartPc:   startPc,
		Length:    length,
		Name:      name,
		Signature: signature,
		Index:     index,
	}
}

func NewMethodParameter(name string, access int) MethodParameter {
	return MethodParameter{
		Name:   name,
		Access: access,
	}
}

func NewModuleInfo(name string, flags int, version string) ModuleInfo {
	return ModuleInfo{
		Name:    name,
		Flags:   flags,
		Version: version,
	}
}

func NewPackageInfo(internalName string, flags int, moduleInfoNames []string) PackageInfo {
	return PackageInfo{
		InternalName:    internalName,
		Flags:           flags,
		ModuleInfoNames: moduleInfoNames,
	}
}

func NewServiceInfo(interfaceTypeName string, implementationTypeNames []string) ServiceInfo {
	return ServiceInfo{
		InterfaceTypeName:       interfaceTypeName,
		ImplementationTypeNames: implementationTypeNames,
	}
}

type Annotation struct {
	Descriptor        string
	ElementValuePairs []ElementValuePair
}

type Annotations struct {
	Annotations []Annotation
}

type BootstrapMethod struct {
	BootstrapMethodRef int
	BootstrapArguments []int
}

type CodeException struct {
	Index     int
	StartPc   int
	EndPc     int
	HandlerPc int
	CatchType int
}

type InnerClass struct {
	InnerTypeName    string
	OuterTypeName    string
	InnerName        string
	InnerAccessFlags int
}

type LineNumber struct {
	StartPc    int
	LineNumber int
}

type LocalVariable struct {
	StartPc    int
	Length     int
	Name       string
	Descriptor string
	Index      int
}

type LocalVariableType struct {
	StartPc   int
	Length    int
	Name      string
	Signature string
	Index     int
}

type MethodParameter struct {
	Name   string
	Access int
}

type ModuleInfo struct {
	Name    string
	Flags   int
	Version string
}

type PackageInfo struct {
	InternalName    string
	Flags           int
	ModuleInfoNames []string
}

type ServiceInfo struct {
	InterfaceTypeName       string
	ImplementationTypeNames []string
}
