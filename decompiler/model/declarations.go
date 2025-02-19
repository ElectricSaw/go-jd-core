package model

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/classfile"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

const (
	// Access flags for Class, Field, Method, Nested class, Module, Module Requires, Module Exports, Module Opens
	FlagPublic       = 0x0001 // C  F  M  N  .  .  .  .
	FlagPrivate      = 0x0002 // .  F  M  N  .  .  .  .
	FlagProtected    = 0x0004 // .  F  M  N  .  .  .  .
	FlagStatic       = 0x0008 // C  F  M  N  .  .  .  .
	FlagFinal        = 0x0010 // C  F  M  N  .  .  .  .
	FlagSynchronized = 0x0020 // .  .  M  .  .  .  .  .
	FlagSuper        = 0x0020 // C  .  .  .  .  .  .  .
	FlagOpen         = 0x0020 // .  .  .  .  Mo .  .  .
	FlagTransitive   = 0x0020 // .  .  .  .  .  MR .  .
	FlagVolatile     = 0x0040 // .  F  .  .  .  .  .  .
	FlagBridge       = 0x0040 // .  .  M  .  .  .  .  .
	FlagStaticPhase  = 0x0040 // .  .  .  .  .  MR .  .
	FlagTransient    = 0x0080 // .  F  .  .  .  .  .  .
	FlagVarArgs      = 0x0080 // .  .  M  .  .  .  .  .
	FlagNative       = 0x0100 // .  .  M  .  .  .  .  .
	FlagInterface    = 0x0200 // C  .  .  N  .  .  .  .
	FlagAnonymous    = 0x0200 // .  .  M  .  .  .  .  . // Custom flag
	FlagAbstract     = 0x0400 // C  .  M  N  .  .  .  .
	FlagStrict       = 0x0800 // .  .  M  .  .  .  .  .
	FlagSynthetic    = 0x1000 // C  F  M  N  Mo MR ME MO
	FlagAnnotation   = 0x2000 // C  .  .  N  .  .  .  .
	FlagEnum         = 0x4000 // C  F  .  N  .  .  .  .
	FlagModule       = 0x8000 // C  .  .  .  .  .  .  .
	FlagMandated     = 0x8000 // .  .  .  .  Mo MR ME MO

	// Extension
	FlagDefault = 0x10000 // .  .  M  .  .  .  .  .
)

/////////////////////////////////////////////////////////////////////////
//  Global Variable
/////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////
//  New Functions
/////////////////////////////////////////////////////////////////////////

func NewAnnotationDeclaration(annotationReferences *AnnotationReference, flags int,
	internalTypeName, name string, annotationDeclarators *FieldDeclarator, bodyDeclaration *BodyDeclaration) AnnotationDeclaration {
	d := AnnotationDeclaration{
		DefaultBase:           *util.NewDefaultBase[IMemberDeclaration]().(*util.DefaultBase[IMemberDeclaration]),
		AnnotationReferences:  annotationReferences,
		Flags:                 flags,
		InternalTypeName:      internalTypeName,
		Name:                  name,
		BodyDeclaration:       bodyDeclaration,
		AnnotationDeclarators: annotationDeclarators,
	}
	d.SetValue(&d)
	return d
}

func NewBodyDeclaration(internalTypeName string, memberDeclaration IMemberDeclaration) BodyDeclaration {
	d := BodyDeclaration{
		InternalTypeName:  internalTypeName,
		MemberDeclaration: memberDeclaration,
	}
	return d
}

func NewClassDeclaration(flags int, internalTypeName, name string,
	bodyDeclaration *BodyDeclaration) ClassDeclaration {
	return NewClassDeclarationWithAll(nil, flags, internalTypeName, name,
		bodyDeclaration, nil, nil, nil)
}

func NewClassDeclarationWithAll(annotationReferences *AnnotationReference, flags int,
	internalTypeName, name string, bodyDeclaration *BodyDeclaration,
	typeParameters ITypeParameter, interfaces IType, superType *ObjectType) ClassDeclaration {
	d := ClassDeclaration{
		DefaultBase:          *util.NewDefaultBase[IMemberDeclaration]().(*util.DefaultBase[IMemberDeclaration]),
		AnnotationReferences: annotationReferences,
		Flags:                flags,
		InternalTypeName:     internalTypeName,
		Name:                 name,
		BodyDeclaration:      bodyDeclaration,
		TypeParameters:       typeParameters,
		Interfaces:           interfaces,
		SuperType:            superType,
	}
	d.SetValue(&d)
	return d
}

func NewConstructorDeclaration(flags int, formalParameter *FormalParameter, descriptor string,
	statements IStatement) ConstructorDeclaration {
	return NewConstructorDeclarationWithAll(nil, flags, nil,
		formalParameter, nil, descriptor, statements)
}

func NewConstructorDeclarationWithAll(annotationReferences *AnnotationReference, flags int,
	typeParameters ITypeParameter, formalParameter *FormalParameter, exceptionTypes IType,
	descriptor string, statements IStatement) ConstructorDeclaration {
	d := ConstructorDeclaration{
		DefaultBase:          *util.NewDefaultBase[IMemberDeclaration]().(*util.DefaultBase[IMemberDeclaration]),
		AnnotationReferences: annotationReferences,
		Flags:                flags,
		TypeParameters:       typeParameters,
		FormalParameters:     formalParameter,
		ExceptionTypes:       exceptionTypes,
		Descriptor:           descriptor,
		Statements:           statements,
	}
	d.SetValue(&d)
	return d
}

func NewArrayVariableInitializer(typ IType) ArrayVariableInitializer {
	d := ArrayVariableInitializer{
		DefaultList: *util.NewDefaultList[IVariableInitializer]().(*util.DefaultList[IVariableInitializer]),
		Type:        typ,
	}
	return d
}

func NewEnumDeclaration(flags int, internalTypeName, name string, constants util.DefaultList[*Constant],
	bodyDeclaration *BodyDeclaration) EnumDeclaration {
	return NewEnumDeclarationWithAll(nil, flags,
		internalTypeName, name, nil, constants, bodyDeclaration)
}

func NewEnumDeclarationWithAll(annotationReferences *AnnotationReference,
	flags int, internalTypeName, name string, interfaces IType,
	constants util.DefaultList[*Constant], bodyDeclaration *BodyDeclaration) EnumDeclaration {
	d := EnumDeclaration{
		DefaultBase:          *util.NewDefaultBase[IMemberDeclaration]().(*util.DefaultBase[IMemberDeclaration]),
		AnnotationReferences: annotationReferences,
		Flags:                flags,
		InternalTypeName:     internalTypeName,
		Name:                 name,
		BodyDeclaration:      bodyDeclaration,
		Interfaces:           interfaces,
		Constants:            constants,
	}
	d.SetValue(&d)
	return d
}

func NewConstant(name string) Constant {
	return NewConstant6(-1, nil, name, nil, nil)
}

func NewConstant2(lineNumber int, name string) Constant {
	return NewConstant6(lineNumber, nil, name, nil, nil)
}

func NewConstant3(name string, arguments IExpression) Constant {
	return NewConstant6(-1, nil, name, arguments, nil)
}

func NewConstant4(lineNumber int, name string, arguments IExpression) Constant {
	return NewConstant6(lineNumber, nil, name, arguments, nil)
}

func NewConstant5(lineNumber int, name string, arguments IExpression,
	bodyDeclaration *BodyDeclaration) Constant {
	return NewConstant6(lineNumber, nil, name, arguments, bodyDeclaration)
}

func NewConstant6(lineNumber int, annotationReferences *AnnotationReference, name string,
	arguments IExpression, bodyDeclaration *BodyDeclaration) Constant {
	c := Constant{
		LineNumber:           lineNumber,
		AnnotationReferences: annotationReferences,
		Name:                 name,
		Arguments:            arguments,
		BodyDeclaration:      bodyDeclaration,
	}
	return c
}

func NewExpressionVariableInitializer(expression IExpression) ExpressionVariableInitializer {
	return ExpressionVariableInitializer{
		Expression: expression,
	}
}

func NewFieldDeclaration(flags int, typ IType, fieldDeclaration *FieldDeclarator) FieldDeclaration {
	return NewFieldDeclarationWithAll(nil, flags, typ, fieldDeclaration)
}

func NewFieldDeclarationWithAll(annotationReferences *AnnotationReference, flags int,
	typ IType, fieldDeclaration *FieldDeclarator) FieldDeclaration {
	d := FieldDeclaration{
		AnnotationReferences: annotationReferences,
		Flags:                flags,
		Type:                 typ,
		FieldDeclarators:     fieldDeclaration,
	}
	d.SetValue(&d)
	return d
}

func NewFieldDeclarator(name string) FieldDeclarator {
	return NewFieldDeclarator3(name, 0, nil)
}

func NewFieldDeclarator2(name string, variableInitializer IVariableInitializer) FieldDeclarator {
	return NewFieldDeclarator3(name, 0, variableInitializer)
}

func NewFieldDeclarator3(name string, dimension int, variableInitializer IVariableInitializer) FieldDeclarator {
	d := FieldDeclarator{
		DefaultBase:         *util.NewDefaultBase[*FieldDeclarator]().(*util.DefaultBase[*FieldDeclarator]),
		Name:                name,
		Dimension:           dimension,
		VariableInitializer: variableInitializer,
	}
	d.SetValue(&d)
	return d
}

func NewFieldDeclarators() FieldDeclarators {
	return NewFieldDeclaratorsWithCapacity(0)
}

func NewFieldDeclaratorsWithCapacity(capacity int) FieldDeclarators {
	d := FieldDeclarators{
		DefaultList: *util.NewDefaultListWithCapacity[*FieldDeclarator](capacity).(*util.DefaultList[*FieldDeclarator]),
	}
	return d
}

func NewFieldDeclaratorsWithElements(fieldDeclarator ...*FieldDeclarator) FieldDeclarators {
	d := FieldDeclarators{
		DefaultList: *util.NewDefaultListWithElements[*FieldDeclarator](fieldDeclarator...).(*util.DefaultList[*FieldDeclarator]),
	}
	return d
}

func NewFormalParameter(typ IType, name string) FormalParameter {
	return NewFormalParameter4(nil, typ, false, name)
}

func NewFormalParameter2(annotationReferences *AnnotationReference, typ IType, name string) FormalParameter {
	return NewFormalParameter4(annotationReferences, typ, false, name)
}

func NewFormalParameter3(typ IType, varargs bool, name string) FormalParameter {
	return NewFormalParameter4(nil, typ, varargs, name)
}

func NewFormalParameter4(annotationReferences *AnnotationReference, typ IType, varargs bool, name string) FormalParameter {
	p := FormalParameter{
		DefaultBase:          *util.NewDefaultBase[*FormalParameter]().(*util.DefaultBase[*FormalParameter]),
		AnnotationReferences: annotationReferences,
		Type:                 typ,
		Varargs:              varargs,
		Name:                 name,
	}
	p.SetValue(&p)
	return p
}

func NewFormalParameters() FormalParameters {
	return NewFormalParametersWithCapacity(0)
}

func NewFormalParametersWithCapacity(capacity int) FormalParameters {
	return FormalParameters{
		DefaultList: *util.NewDefaultListWithCapacity[*FormalParameter](capacity).(*util.DefaultList[*FormalParameter]),
	}
}

func NewFormalParametersWithElements(formalParameter ...*FormalParameter) FormalParameters {
	d := FormalParameters{
		DefaultList: *util.NewDefaultListWithElements[*FormalParameter](formalParameter...).(*util.DefaultList[*FormalParameter]),
	}
	return d
}

func NewInstanceInitializerDeclaration(description string, statements IStatement) InstanceInitializerDeclaration {
	d := InstanceInitializerDeclaration{
		DefaultBase: *util.NewDefaultBase[IMemberDeclaration]().(*util.DefaultBase[IMemberDeclaration]),
		Description: description,
		Statements:  statements,
	}
	d.SetValue(&d)
	return d
}

func NewInterfaceDeclaration(flags int, internalTypeName, name string, interfaces IType) InterfaceDeclaration {
	return NewInterfaceDeclarationWithAll(nil, flags, internalTypeName, name, nil, nil, interfaces)
}

func NewInterfaceDeclarationWithAll(annotationReferences *AnnotationReference, flags int,
	internalTypeName, name string, bodyDeclaration *BodyDeclaration,
	typeParameters ITypeParameter, interfaces IType) InterfaceDeclaration {
	d := InterfaceDeclaration{
		DefaultBase:          *util.NewDefaultBase[IMemberDeclaration]().(*util.DefaultBase[IMemberDeclaration]),
		AnnotationReferences: annotationReferences,
		Flags:                flags,
		InternalTypeName:     internalTypeName,
		Name:                 name,
		BodyDeclaration:      bodyDeclaration,
		TypeParameters:       typeParameters,
		Interfaces:           interfaces,
	}
	d.SetValue(&d)
	return d
}

func NewLocalVariableDeclaration(typ IType, localVariableDeclarators *LocalVariableDeclarator) LocalVariableDeclaration {
	return LocalVariableDeclaration{
		Type:                     typ,
		LocalVariableDeclarators: localVariableDeclarators,
	}
}

func NewLocalVariableDeclarator(name string) LocalVariableDeclarator {
	return NewLocalVariableDeclarator3(0, name, nil)
}

func NewLocalVariableDeclarator2(name string, variableInitializer IVariableInitializer) LocalVariableDeclarator {
	return NewLocalVariableDeclarator3(0, name, variableInitializer)
}

func NewLocalVariableDeclarator3(lineNumber int, name string, variableInitializer IVariableInitializer) LocalVariableDeclarator {
	d := LocalVariableDeclarator{
		DefaultBase:         *util.NewDefaultBase[ILocalVariableDeclarator]().(*util.DefaultBase[ILocalVariableDeclarator]),
		LineNumber:          lineNumber,
		Name:                name,
		VariableInitializer: variableInitializer,
	}
	d.SetValue(&d)
	return d
}

func NewLocalVariableDeclarators() LocalVariableDeclarators {
	return NewLocalVariableDeclaratorsWithCapacity(0)
}

func NewLocalVariableDeclaratorsWithCapacity(capacity int) LocalVariableDeclarators {
	return LocalVariableDeclarators{
		DefaultList: *util.NewDefaultListWithCapacity[ILocalVariableDeclarator](capacity).(*util.DefaultList[ILocalVariableDeclarator]),
	}
}

func NewLocalVariableDeclaratorsWithElements(localVariableDeclarator ...ILocalVariableDeclarator) LocalVariableDeclarators {
	return LocalVariableDeclarators{
		DefaultList: *util.NewDefaultListWithElements[ILocalVariableDeclarator](localVariableDeclarator...).(*util.DefaultList[ILocalVariableDeclarator]),
	}
}

func NewMemberDeclarations() MemberDeclarations {
	return NewMemberDeclarationsWithCapacity(0)
}

func NewMemberDeclarationsWithCapacity(capacity int) MemberDeclarations {
	return MemberDeclarations{
		DefaultList: *util.NewDefaultListWithCapacity[IMemberDeclaration](capacity).(*util.DefaultList[IMemberDeclaration]),
	}
}

func NewMemberDeclarationsWithElements(memberDeclaration ...IMemberDeclaration) MemberDeclarations {
	return MemberDeclarations{
		DefaultList: *util.NewDefaultListWithElements[IMemberDeclaration](memberDeclaration...).(*util.DefaultList[IMemberDeclaration]),
	}
}

func NewMethodDeclaration(flags int, name string, returnedType IType,
	descriptor string) MethodDeclaration {
	return NewMethodDeclaration6(nil, flags, name, nil,
		returnedType, nil, nil, descriptor, nil, nil)
}

func NewMethodDeclaration2(flags int, name string, returnedType IType,
	descriptor string, statements IStatement) MethodDeclaration {
	return NewMethodDeclaration6(nil, flags, name, nil,
		returnedType, nil, nil, descriptor, statements, nil)
}

func NewMethodDeclaration3(flags int, name string, returnedType IType,
	descriptor string, defaultAnnotationValue IElementValue) MethodDeclaration {
	return NewMethodDeclaration6(nil, flags, name, nil,
		returnedType, nil, nil, descriptor, nil, defaultAnnotationValue)
}

func NewMethodDeclaration4(flags int, name string, returnedType IType,
	formalParameter *FormalParameter, descriptor string, statements IStatement) MethodDeclaration {
	return NewMethodDeclaration6(nil, flags, name, nil,
		returnedType, formalParameter, nil, descriptor, statements, nil)
}

func NewMethodDeclaration5(flags int, name string, returnedType IType,
	formalParameter *FormalParameter, descriptor string,
	defaultAnnotationValue IElementValue) MethodDeclaration {
	return NewMethodDeclaration6(nil, flags, name, nil,
		returnedType, formalParameter, nil, descriptor, nil, defaultAnnotationValue)
}

func NewMethodDeclaration6(annotationReferences *AnnotationReference,
	flags int, name string, typeParameters ITypeParameter, returnedType IType,
	formalParameter *FormalParameter, exceptionTypes IType, descriptor string,
	statements IStatement, defaultAnnotationValue IElementValue) MethodDeclaration {
	d := MethodDeclaration{
		DefaultBase:            *util.NewDefaultBase[IMemberDeclaration]().(*util.DefaultBase[IMemberDeclaration]),
		AnnotationReferences:   annotationReferences,
		Flags:                  flags,
		Name:                   name,
		TypeParameters:         typeParameters,
		ReturnedType:           returnedType,
		FormalParameter:        formalParameter,
		ExceptionTypes:         exceptionTypes,
		Descriptor:             descriptor,
		Statements:             statements,
		DefaultAnnotationValue: defaultAnnotationValue,
	}
	d.SetValue(&d)
	return d
}

func NewModuleDeclaration(flags int, internalTypeName, name, version string,
	requires util.IList[ModuleInfo], exports util.IList[PackageInfo],
	opens util.IList[PackageInfo], uses util.IList[string],
	provides util.IList[ServiceInfo]) ModuleDeclaration {
	d := ModuleDeclaration{
		DefaultBase:      *util.NewDefaultBase[IMemberDeclaration]().(*util.DefaultBase[IMemberDeclaration]),
		Flags:            flags,
		InternalTypeName: internalTypeName,
		Name:             name,
		Version:          version,
		Requires:         requires,
		Exports:          exports,
		Opens:            opens,
		Uses:             uses,
		Provides:         provides,
	}
	d.SetValue(&d)
	return d
}

func NewModuleInfo(name string, flags int, version string) ModuleInfo {
	return ModuleInfo{
		Name:    name,
		Flags:   flags,
		Version: version,
	}
}

func NewPackageInfo(internalName string, flags int, moduleInfoNames []string) PackageInfo {
	return NewPackageInfoWithList(internalName, flags, util.NewDefaultListWithElements[string](moduleInfoNames...))
}

func NewPackageInfoWithList(internalName string, flags int, moduleInfoNames util.IList[string]) PackageInfo {
	return PackageInfo{
		InternalName:    internalName,
		Flags:           flags,
		ModuleInfoNames: moduleInfoNames,
	}
}

func NewServiceInfo(internalTypeName string, implementationTypeNames []string) ServiceInfo {
	return NewServiceInfoWithList(internalTypeName, util.NewDefaultListWithElements[string](implementationTypeNames...))
}

func NewServiceInfoWithList(internalTypeName string, implementationTypeNames util.IList[string]) ServiceInfo {
	return ServiceInfo{
		InternalTypeName:        internalTypeName,
		ImplementationTypeNames: implementationTypeNames,
	}
}

func NewStaticInitializerDeclaration(descriptor string, statements IStatement) StaticInitializerDeclaration {
	d := StaticInitializerDeclaration{
		Descriptor: descriptor,
		Statements: statements,
	}
	d.SetValue(&d)
	return d
}

func NewTypeDeclarations() TypeDeclarations {
	return NewTypeDeclarationsWithCapacity(0)
}

func NewTypeDeclarationsWithCapacity(capacity int) TypeDeclarations {
	return TypeDeclarations{
		DefaultList: *util.NewDefaultListWithCapacity[IMemberDeclaration](capacity).(*util.DefaultList[IMemberDeclaration]),
	}
}

func NewTypeDeclarationsWithElements(memberDeclaration ...IMemberDeclaration) TypeDeclarations {
	return TypeDeclarations{
		DefaultList: *util.NewDefaultListWithElements[IMemberDeclaration](memberDeclaration...).(*util.DefaultList[IMemberDeclaration]),
	}
}

/////////////////////////////////////////////////////////////////////////
//  Interfaces
/////////////////////////////////////////////////////////////////////////

type IDeclaration interface {
	AcceptDeclaration(visitor IDeclarationVisitor)
	String() string
}

type IDeclarationVisitor interface {
	VisitAnnotationDeclaration(declaration *AnnotationDeclaration)
	VisitArrayVariableInitializer(declaration *ArrayVariableInitializer)
	VisitBodyDeclaration(declaration *BodyDeclaration)
	VisitClassDeclaration(declaration *ClassDeclaration)
	VisitConstructorDeclaration(declaration *ConstructorDeclaration)
	VisitEnumDeclaration(declaration *EnumDeclaration)
	VisitEnumDeclarationConstant(declaration *Constant)
	VisitExpressionVariableInitializer(declaration *ExpressionVariableInitializer)
	VisitFieldDeclaration(declaration *FieldDeclaration)
	VisitFieldDeclarator(declaration *FieldDeclarator)
	VisitFieldDeclarators(declarations *FieldDeclarators)
	VisitFormalParameter(declaration *FormalParameter)
	VisitFormalParameters(declarations *FormalParameters)
	VisitInstanceInitializerDeclaration(declaration *InstanceInitializerDeclaration)
	VisitInterfaceDeclaration(declaration *InterfaceDeclaration)
	VisitLocalVariableDeclaration(declaration ILocalVariableDeclaration)
	VisitLocalVariableDeclarator(declarator *LocalVariableDeclarator)
	VisitLocalVariableDeclarators(declarators *LocalVariableDeclarators)
	VisitMethodDeclaration(declaration *MethodDeclaration)
	VisitMemberDeclarations(declarations *MemberDeclarations)
	VisitModuleDeclaration(declarations *ModuleDeclaration)
	VisitStaticInitializerDeclaration(declaration *StaticInitializerDeclaration)
	VisitTypeDeclarations(declarations *TypeDeclarations)
}

type IMemberDeclaration interface {
	IsClassDeclaration() bool
	AcceptDeclaration(visitor IDeclarationVisitor)
	String() string
}

type ITypeDeclaration interface {
	GetAnnotationReferences() *AnnotationReference
	GetFlag() int
	GetInternalTypeName() string
	GetName() string
	GetBodyDeclaration() *BodyDeclaration
	IsClassDeclaration() bool
	AcceptDeclaration(visitor IDeclarationVisitor)
	String() string
}

type ILocalVariableDeclarator interface {
	GetLineNumber() int
	AcceptDeclaration(visitor IDeclarationVisitor)
	String() string
}

type IVariableInitializer interface {
	GetLineNumber() int
	GetExpression() IExpression
	IsExpressionVariableInitializer() bool
	AcceptDeclaration(visitor IDeclarationVisitor)
	String() string
}

type IFormalParameter interface {
	AcceptDeclaration(visitor IDeclarationVisitor)
	String() string
}

type IFieldDeclarator interface {
	SetFieldDeclaration(declaration *FieldDeclaration)
	AcceptDeclaration(visitor IDeclarationVisitor)
	String() string
}

type ILocalVariableDeclaration interface {
	IsFinal() bool
	GetType() IType
	GetLocalVariableDeclarators() ILocalVariableDeclarator
	String() string
}

/////////////////////////////////////////////////////////////////////////
//  Structures
/////////////////////////////////////////////////////////////////////////

type AnnotationDeclaration struct {
	util.DefaultBase[IMemberDeclaration]

	AnnotationReferences  *AnnotationReference
	Flags                 int
	InternalTypeName      string
	Name                  string
	BodyDeclaration       *BodyDeclaration
	AnnotationDeclarators *FieldDeclarator
}

func (d *AnnotationDeclaration) GetAnnotationReferences() *AnnotationReference {
	return d.AnnotationReferences
}

func (d *AnnotationDeclaration) GetFlag() int {
	return d.Flags
}

func (d *AnnotationDeclaration) GetInternalTypeName() string {
	return d.InternalTypeName
}

func (d *AnnotationDeclaration) GetName() string {
	return d.Name
}

func (d *AnnotationDeclaration) GetBodyDeclaration() *BodyDeclaration {
	return d.BodyDeclaration
}

func (d *AnnotationDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *AnnotationDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitAnnotationDeclaration(d)
}

func (d *AnnotationDeclaration) String() string {
	return fmt.Sprintf("AnnotationDeclaration { %s }", d.InternalTypeName)
}

type ArrayVariableInitializer struct {
	util.DefaultList[IVariableInitializer]

	Type IType
}

func (i *ArrayVariableInitializer) GetLineNumber() int {
	if i.Size() == 0 {
		return UnknownLineNumber
	}
	return i.Get(0).GetLineNumber()
}

func (i *ArrayVariableInitializer) GetExpression() IExpression {
	return &NeNoExpression
}

func (i *ArrayVariableInitializer) IsExpressionVariableInitializer() bool {
	return false
}

func (i *ArrayVariableInitializer) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitArrayVariableInitializer(i)
}

func (i *ArrayVariableInitializer) String() string {
	return fmt.Sprintf("ArrayVariableInitializer{ type=%s, size=%d }", i.Type, i.Size())
}

type BodyDeclaration struct {
	InternalTypeName  string
	MemberDeclaration IMemberDeclaration
}

func (d *BodyDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}

func (d *BodyDeclaration) String() string {
	return fmt.Sprintf("BodyDeclaration{ internal-type-name=%s, member=%s }", d.InternalTypeName, d.MemberDeclaration)
}

type ClassDeclaration struct {
	util.DefaultBase[IMemberDeclaration]

	AnnotationReferences *AnnotationReference
	Flags                int
	InternalTypeName     string
	Name                 string
	BodyDeclaration      *BodyDeclaration
	TypeParameters       ITypeParameter
	Interfaces           IType
	SuperType            *ObjectType
}

func (d *ClassDeclaration) GetAnnotationReferences() *AnnotationReference {
	return d.AnnotationReferences
}

func (d *ClassDeclaration) GetFlag() int {
	return d.Flags
}

func (d *ClassDeclaration) GetInternalTypeName() string {
	return d.InternalTypeName
}

func (d *ClassDeclaration) GetName() string {
	return d.Name
}

func (d *ClassDeclaration) GetBodyDeclaration() *BodyDeclaration {
	return d.BodyDeclaration
}

func (d *ClassDeclaration) IsClassDeclaration() bool {
	return true
}

func (d *ClassDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitClassDeclaration(d)
}

func (d *ClassDeclaration) String() string {
	return fmt.Sprintf("ClassDeclaration{%s}", d.InternalTypeName)
}

type ConstructorDeclaration struct {
	util.DefaultBase[IMemberDeclaration]

	AnnotationReferences *AnnotationReference
	Flags                int
	TypeParameters       ITypeParameter
	FormalParameters     *FormalParameter
	ExceptionTypes       IType
	Descriptor           string
	Statements           IStatement
}

func (d *ConstructorDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ConstructorDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitConstructorDeclaration(d)
}

func (d *ConstructorDeclaration) String() string {
	return fmt.Sprintf("ConstructorDeclaration{ %s }", d.Descriptor)
}

type EnumDeclaration struct {
	util.DefaultBase[IMemberDeclaration]

	AnnotationReferences *AnnotationReference
	Flags                int
	InternalTypeName     string
	Name                 string
	BodyDeclaration      *BodyDeclaration
	Interfaces           IType
	Constants            util.DefaultList[*Constant]
}

func (d *EnumDeclaration) GetAnnotationReferences() *AnnotationReference {
	return d.AnnotationReferences
}

func (d *EnumDeclaration) GetFlag() int {
	return d.Flags
}

func (d *EnumDeclaration) GetInternalTypeName() string {
	return d.InternalTypeName
}

func (d *EnumDeclaration) GetName() string {
	return d.Name
}

func (d *EnumDeclaration) GetBodyDeclaration() *BodyDeclaration {
	return d.BodyDeclaration
}

func (d *EnumDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *EnumDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitEnumDeclaration(d)
}

func (d *EnumDeclaration) String() string {
	return fmt.Sprintf("EnumDeclaration{%s}", d.InternalTypeName)
}

type ExpressionVariableInitializer struct {
	Expression IExpression
}

func (i *ExpressionVariableInitializer) GetLineNumber() int {
	return i.Expression.GetLineNumber()
}

func (i *ExpressionVariableInitializer) GetExpression() IExpression {
	return i.Expression
}

func (i *ExpressionVariableInitializer) IsExpressionVariableInitializer() bool {
	return true
}

func (i *ExpressionVariableInitializer) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitExpressionVariableInitializer(i)
}

func (i *ExpressionVariableInitializer) String() string {
	return fmt.Sprintf("")
}

type FieldDeclaration struct {
	util.DefaultBase[IMemberDeclaration]

	AnnotationReferences *AnnotationReference
	Flags                int
	Type                 IType
	FieldDeclarators     *FieldDeclarator
}

func (d *FieldDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *FieldDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitFieldDeclaration(d)
}

func (d *FieldDeclaration) HashCode() int {
	result := 327494460 + d.Flags
	if d.AnnotationReferences != nil {
		result = 31*result + (d.AnnotationReferences.HashCode())
	} else {
		result = 31*result + 0
	}
	result = 31*result + d.Type.HashCode()
	result = 31*result + d.FieldDeclarators.HashCode()
	return result
}

func (d *FieldDeclaration) String() string {
	return fmt.Sprintf("FieldDeclaration{ type=%s, field-declarators=%s }", d.Type, d.FieldDeclarators)
}

type FieldDeclarator struct {
	util.DefaultBase[*FieldDeclarator]

	FieldDeclaration    *FieldDeclaration
	Name                string
	Dimension           int
	VariableInitializer IVariableInitializer
}

func (d *FieldDeclarator) SetFieldDeclaration(fieldDeclaration *FieldDeclaration) {
	d.FieldDeclaration = fieldDeclaration
}

func (d *FieldDeclarator) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitFieldDeclarator(d)
}

func (d *FieldDeclarator) HashCode() int {
	return hashCodeWithStruct(d)
}

func (d *FieldDeclarator) String() string {
	return fmt.Sprintf("FieldDeclarator{ %s }", d.Name)
}

type FieldDeclarators struct {
	util.DefaultList[*FieldDeclarator]
}

func (d *FieldDeclarators) SetFieldDeclaration(fieldDeclaration *FieldDeclaration) {
	for _, fieldDeclarator := range d.ToSlice() {
		fieldDeclarator.SetFieldDeclaration(fieldDeclaration)
	}
}

func (d *FieldDeclarators) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitFieldDeclarators(d)
}

func (d *FieldDeclarators) String() string {
	return fmt.Sprintf("FieldDeclarators{}")
}

type FormalParameter struct {
	util.DefaultBase[*FormalParameter]

	AnnotationReferences *AnnotationReference
	Final                bool
	Type                 IType
	Varargs              bool
	Name                 string
}

func (d *FormalParameter) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitFormalParameter(d)
}

func (d *FormalParameter) String() string {
	sb := "FormalParameter{"

	if d.AnnotationReferences != nil {
		sb += fmt.Sprintf("%v ", d.AnnotationReferences)
	}

	if d.Varargs {
		sb += fmt.Sprintf("%v... ", d.Type.CreateType(d.Type.GetDimension()-1))
	} else {
		sb += fmt.Sprintf("%v ", d.Type)
	}
	sb += "}"

	return sb
}

type FormalParameters struct {
	util.DefaultList[*FormalParameter]
}

func (d *FormalParameters) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitFormalParameters(d)
}

func (d *FormalParameters) String() string {
	return fmt.Sprintf("FormalParameters{ }")
}

type InstanceInitializerDeclaration struct {
	util.DefaultBase[IMemberDeclaration]

	Description string
	Statements  IStatement
}

func (d *InstanceInitializerDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *InstanceInitializerDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitInstanceInitializerDeclaration(d)
}

func (d *InstanceInitializerDeclaration) String() string {
	return "InstanceInitializerDeclaration{}"
}

type InterfaceDeclaration struct {
	util.DefaultBase[IMemberDeclaration]

	AnnotationReferences *AnnotationReference
	Flags                int
	InternalTypeName     string
	Name                 string
	BodyDeclaration      *BodyDeclaration
	TypeParameters       ITypeParameter
	Interfaces           IType
}

func (d *InterfaceDeclaration) GetAnnotationReferences() *AnnotationReference {
	return d.AnnotationReferences
}

func (d *InterfaceDeclaration) GetFlag() int {
	return d.Flags
}

func (d *InterfaceDeclaration) GetInternalTypeName() string {
	return d.InternalTypeName
}

func (d *InterfaceDeclaration) GetName() string {
	return d.Name
}

func (d *InterfaceDeclaration) GetBodyDeclaration() *BodyDeclaration {
	return d.BodyDeclaration
}

func (d *InterfaceDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *InterfaceDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitInterfaceDeclaration(d)
}

func (d *InterfaceDeclaration) String() string {
	return fmt.Sprintf("InterfaceDeclaration{%v}", *d)
}

type LocalVariableDeclaration struct {
	Final                    bool
	Type                     IType
	LocalVariableDeclarators ILocalVariableDeclarator
}

func (d *LocalVariableDeclaration) IsFinal() bool {
	return d.Final
}

func (d *LocalVariableDeclaration) GetType() IType {
	return d.Type
}

func (d *LocalVariableDeclaration) GetLocalVariableDeclarators() ILocalVariableDeclarator {
	return d.LocalVariableDeclarators
}

func (d *LocalVariableDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitLocalVariableDeclaration(d)
}

func (d *LocalVariableDeclaration) String() string {
	return fmt.Sprintf("LocalVariableDeclaration{ final=%s, type=%s }")
}

type LocalVariableDeclarator struct {
	util.DefaultBase[ILocalVariableDeclarator]

	LineNumber          int
	Name                string
	Dimension           int
	VariableInitializer IVariableInitializer
}

func (d *LocalVariableDeclarator) GetLineNumber() int {
	return d.LineNumber
}

func (d *LocalVariableDeclarator) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitLocalVariableDeclarator(d)

}

func (d *LocalVariableDeclarator) String() string {
	return fmt.Sprintf("LocalVariableDeclarator{ name=%s, dimension=%d, variable-initializer=%v }", d.Name, d.Dimension, d.VariableInitializer)
}

type LocalVariableDeclarators struct {
	util.DefaultList[ILocalVariableDeclarator]
}

func (d *LocalVariableDeclarators) GetLineNumber() int {
	if d.Size() == 0 {
		return UnknownLineNumber
	}
	return d.Get(0).GetLineNumber()
}

func (d *LocalVariableDeclarators) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitLocalVariableDeclarators(d)
}

func (d *LocalVariableDeclarators) String() string {
	return "LocalVariableDeclarators{}"
}

type MemberDeclarations struct {
	util.DefaultList[IMemberDeclaration]
}

func (d *MemberDeclarations) IsClassDeclaration() bool {
	return false
}

func (d *MemberDeclarations) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitMemberDeclarations(d)
}

func (d *MemberDeclarations) String() string {
	return "MemberDeclarations{}"
}

type MethodDeclaration struct {
	util.DefaultBase[IMemberDeclaration]

	AnnotationReferences   *AnnotationReference
	Flags                  int
	Name                   string
	TypeParameters         ITypeParameter
	ReturnedType           IType
	FormalParameter        *FormalParameter
	ExceptionTypes         IType
	Descriptor             string
	Statements             IStatement
	DefaultAnnotationValue IElementValue
}

func (d *MethodDeclaration) IsStatic() bool {
	return d.Flags&classfile.AccStatic != 0
}

func (d *MethodDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *MethodDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitMethodDeclaration(d)
}

func (d *MethodDeclaration) String() string {
	return fmt.Sprintf("MethodDeclaration{ name=%s, descriptor=%s }", d.Name, d.Descriptor)
}

type ModuleDeclaration struct {
	util.DefaultBase[IMemberDeclaration]

	AnnotationReferences *AnnotationReference
	Flags                int
	InternalTypeName     string
	Name                 string
	BodyDeclaration      *BodyDeclaration

	Version  string
	Requires util.IList[ModuleInfo]
	Exports  util.IList[PackageInfo]
	Opens    util.IList[PackageInfo]
	Uses     util.IList[string]
	Provides util.IList[ServiceInfo]
}

func (d *ModuleDeclaration) GetAnnotationReferences() *AnnotationReference {
	return d.AnnotationReferences
}

func (d *ModuleDeclaration) GetFlag() int {
	return d.Flags
}

func (d *ModuleDeclaration) GetInternalTypeName() string {
	return d.InternalTypeName
}

func (d *ModuleDeclaration) GetName() string {
	return d.Name
}

func (d *ModuleDeclaration) GetBodyDeclaration() *BodyDeclaration {
	return d.BodyDeclaration
}

func (d *ModuleDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ModuleDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitModuleDeclaration(d)
}

func (d *ModuleDeclaration) String() string {
	return fmt.Sprintf("ModuleDeclaration{ %s }", d.InternalTypeName)
}

type StaticInitializerDeclaration struct {
	util.DefaultBase[IMemberDeclaration]

	Descriptor string
	Statements IStatement
}

func (d *StaticInitializerDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *StaticInitializerDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitStaticInitializerDeclaration(d)
}

func (d *StaticInitializerDeclaration) String() string {
	return "StaticInitializerDeclaration{}"
}

type TypeDeclarations struct {
	util.DefaultList[IMemberDeclaration]
}

func (d *TypeDeclarations) IsClassDeclaration() bool {
	return false
}

func (d *TypeDeclarations) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitTypeDeclarations(d)
}

func (d *TypeDeclarations) String() string {
	return fmt.Sprintf("")
}

/////////////////////////////////////////////////////////////////////////
//  Additional Structures
/////////////////////////////////////////////////////////////////////////

type Constant struct {
	LineNumber           int
	AnnotationReferences *AnnotationReference
	Name                 string
	Arguments            IExpression
	BodyDeclaration      *BodyDeclaration
}

func (c *Constant) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitEnumDeclarationConstant(c)
}

func (c *Constant) String() string {
	return fmt.Sprintf("Constant{ %s }", c.Name)
}

type ModuleInfo struct {
	Name    string
	Flags   int
	Version string
}

func (i *ModuleInfo) String() string {
	msg := fmt.Sprintf("ModuleInfo{name=%s, flags=%d", i.Name, i.Flags)
	if i.Version != "" {
		msg += fmt.Sprintf(", version=%s", i.Version)
	}
	msg += "}"

	return msg
}

type PackageInfo struct {
	InternalName    string
	Flags           int
	ModuleInfoNames util.IList[string]
}

func (i *PackageInfo) String() string {
	msg := fmt.Sprintf("PackageInfo{InternalName=%s, flags=%d", i.InternalName, i.Flags)
	if i.ModuleInfoNames.Size() > 0 {
		msg += fmt.Sprintf(", moduleInfoNames=%s", i.ModuleInfoNames.ToSlice())
	}
	msg += "}"

	return msg
}

type ServiceInfo struct {
	InternalTypeName        string
	ImplementationTypeNames util.IList[string]
}

func (i *ServiceInfo) String() string {
	msg := fmt.Sprintf("PackageInfo{internalTypeName=%s", i.InternalTypeName)
	if i.ImplementationTypeNames.Size() > 0 {
		msg += fmt.Sprintf(", implementationTypeNames=%s", i.ImplementationTypeNames.ToSlice())
	}
	msg += "}"

	return msg
}

/////////////////////////////////////////////////////////////////////////
//  Functions
/////////////////////////////////////////////////////////////////////////
