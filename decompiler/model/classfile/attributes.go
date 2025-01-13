package classfile

import "fmt"

func NewAttributeAnnotationDefault(defaultValue IElementValue) AttributeAnnotationDefault {
	return AttributeAnnotationDefault{
		DefaultValue: defaultValue,
	}
}

func NewAttributeBootstrapMethods(bootstrapMethod []BootstrapMethod) AttributeBootstrapMethods {
	return AttributeBootstrapMethods{
		BootstrapMethod: bootstrapMethod,
	}
}

func NewAttributeCode(maxStack int, maxLocals int, code []byte,
	exceptionTable []CodeException, attribute map[string]IAttribute) AttributeCode {
	return AttributeCode{
		MaxStack:       maxStack,
		MaxLocals:      maxLocals,
		Code:           code,
		ExceptionTable: exceptionTable,
		Attribute:      attribute,
	}
}

func NewAttributeConstantValue(constantValue IConstantValue) AttributeConstantValue {
	return AttributeConstantValue{
		ConstantValue: constantValue,
	}
}

func NewAttributeDeprecated() AttributeDeprecated {
	return AttributeDeprecated{}
}

func NewAttributeExceptions(exceptionTypeNames []string) AttributeExceptions {
	return AttributeExceptions{
		ExceptionTypeNames: exceptionTypeNames,
	}
}

func NewAttributeInnerClasses(classes []InnerClass) AttributeInnerClasses {
	return AttributeInnerClasses{
		Classes: classes,
	}
}

func NewAttributeLineNumberTable(lineNumberTable []LineNumber) AttributeLineNumberTable {
	return AttributeLineNumberTable{
		LineNumberTable: lineNumberTable,
	}
}

func NewAttributeLocalVariableTable(localVariableTable []LocalVariable) AttributeLocalVariableTable {
	return AttributeLocalVariableTable{
		LocalVariableTable: localVariableTable,
	}
}

func NewAttributeLocalVariableTypeTable(localVariableTypeTable []LocalVariableType) AttributeLocalVariableTypeTable {
	return AttributeLocalVariableTypeTable{
		LocalVariableTypeTable: localVariableTypeTable,
	}
}

func NewAttributeMethodParameters(parameters []MethodParameter) AttributeMethodParameters {
	return AttributeMethodParameters{
		Parameters: parameters,
	}
}

func NewAttributeModule(name string, flags int, version string,
	requires []ModuleInfo, exports []PackageInfo, opens []PackageInfo,
	uses []string, provides []ServiceInfo) AttributeModule {
	return AttributeModule{
		Name:     name,
		Flags:    flags,
		Version:  version,
		Requires: requires,
		Exports:  exports,
		Opens:    opens,
		Uses:     uses,
		Provides: provides,
	}
}

func NewAttributeModuleMainClass(mainClass ConstantClass) AttributeModuleMainClass {
	return AttributeModuleMainClass{
		MainClass: mainClass,
	}
}

func NewAttributeModulePackages(packageNames []string) AttributeModulePackages {
	return AttributeModulePackages{
		PackageNames: packageNames,
	}
}

func NewAttributeParameterAnnotations(parameterAnnotations []Annotations) AttributeParameterAnnotations {
	return AttributeParameterAnnotations{
		ParameterAnnotations: parameterAnnotations,
	}
}

func NewAttributeSignature(signature string) AttributeSignature {
	return AttributeSignature{
		Signature: signature,
	}
}

func NewAttributeSourceFile(sourceFile string) AttributeSourceFile {
	return AttributeSourceFile{
		SourceFile: sourceFile,
	}
}

func NewAttributeSynthetic() AttributeSynthetic {
	return AttributeSynthetic{}
}

func NewAttributeUnknown() AttributeUnknown {
	return AttributeUnknown{}
}

type IAttribute interface {
	String() string
}

type AttributeAnnotationDefault struct {
	DefaultValue IElementValue
}

func (a AttributeAnnotationDefault) String() string {
	return fmt.Sprintf("AttributeAnnotationDefault{ %s }", a.DefaultValue)
}

type AttributeBootstrapMethods struct {
	BootstrapMethod []BootstrapMethod
}

func (a AttributeBootstrapMethods) String() string {
	return fmt.Sprintf("AttributeBootstrapMethods{ bootstrapMethod: %d }", len(a.BootstrapMethod))
}

type AttributeCode struct {
	MaxStack       int
	MaxLocals      int
	Code           []byte
	ExceptionTable []CodeException
	Attribute      map[string]IAttribute
}

func (a AttributeCode) String() string {
	return fmt.Sprintf("AttributeCode{ maxStack: %d, maxLocals: %d }", a.MaxStack, a.MaxLocals)
}

type AttributeConstantValue struct {
	ConstantValue IConstantValue
}

func (a AttributeConstantValue) String() string {
	return fmt.Sprintf("AttributeConstantValue{ %s }", a.ConstantValue)
}

type AttributeDeprecated struct {
}

func (a AttributeDeprecated) String() string {
	return fmt.Sprintf("AttributeDeprecated{}")
}

type AttributeExceptions struct {
	ExceptionTypeNames []string
}

func (a AttributeExceptions) String() string {
	return fmt.Sprintf("AttributeExceptions{ exceptionTypeNames: %d }", len(a.ExceptionTypeNames))
}

type AttributeInnerClasses struct {
	Classes []InnerClass
}

func (a AttributeInnerClasses) String() string {
	return fmt.Sprintf("AttributeInnerClasses{ classes: %d }", len(a.Classes))
}

type AttributeLineNumberTable struct {
	LineNumberTable []LineNumber
}

func (a AttributeLineNumberTable) String() string {
	return fmt.Sprintf("AttributeLineNumberTable{ lineNumberTable: %d }", len(a.LineNumberTable))
}

type AttributeLocalVariableTable struct {
	LocalVariableTable []LocalVariable
}

func (a AttributeLocalVariableTable) String() string {
	return fmt.Sprintf("AttributeLocalVariableTable{ localVariableTable: %d }", len(a.LocalVariableTable))
}

type AttributeLocalVariableTypeTable struct {
	LocalVariableTypeTable []LocalVariableType
}

func (a AttributeLocalVariableTypeTable) String() string {
	return fmt.Sprintf("AttributeLocalVariableTypeTable{ localVariableTypeTable: %d }", len(a.LocalVariableTypeTable))
}

type AttributeMethodParameters struct {
	Parameters []MethodParameter
}

func (a AttributeMethodParameters) String() string {
	return fmt.Sprintf("AttributeMethodParameters{ parameters: %d }", len(a.Parameters))
}

type AttributeModule struct {
	Name    string
	Flags   int
	Version string

	Requires []ModuleInfo
	Exports  []PackageInfo
	Opens    []PackageInfo
	Uses     []string
	Provides []ServiceInfo
}

func (a AttributeModule) String() string {
	return fmt.Sprintf("AttributeModule{ name: %s, flags: %d, version: %s }", a.Name, a.Flags, a.Version)
}

type AttributeModuleMainClass struct {
	MainClass ConstantClass
}

func (a AttributeModuleMainClass) String() string {
	return fmt.Sprintf("AttributeModuleMainClass{ %s }", a.MainClass)
}

type AttributeModulePackages struct {
	PackageNames []string
}

func (a AttributeModulePackages) String() string {
	return fmt.Sprintf("AttributeModulePackages{ packageNames: %d }", len(a.PackageNames))
}

type AttributeParameterAnnotations struct {
	ParameterAnnotations []Annotations
}

func (a AttributeParameterAnnotations) String() string {
	return fmt.Sprintf("AttributeAnnotationDefault{ parameterAnnotations: %d }", len(a.ParameterAnnotations))
}

type AttributeSignature struct {
	Signature string
}

func (a AttributeSignature) String() string {
	return fmt.Sprintf("AttributeSignature{ %s }", a.Signature)
}

type AttributeSourceFile struct {
	SourceFile string
}

func (a AttributeSourceFile) String() string {
	return fmt.Sprintf("AttributeSourceFile{ %s }", a.SourceFile)
}

type AttributeSynthetic struct {
}

func (a AttributeSynthetic) String() string {
	return fmt.Sprintf("AttributeSynthetic{}")
}

type AttributeUnknown struct {
}

func (a AttributeUnknown) String() string {
	return fmt.Sprintf("AttributeUnknown{}")
}
