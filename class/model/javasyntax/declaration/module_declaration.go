package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewModuleDeclaration(flags int, internalTypeName, name, version string,
	requires []intsyn.IModuleInfo, exports []intsyn.IPackageInfo, opens []intsyn.IPackageInfo,
	uses []string, provides []intsyn.IServiceInfo) intsyn.IModuleDeclaration {
	return &ModuleDeclaration{
		TypeDeclaration: *NewTypeDeclaration(nil, flags, internalTypeName, name, nil).(*TypeDeclaration),
		version:         version,
		requires:        requires,
		exports:         exports,
		opens:           opens,
		uses:            uses,
		provides:        provides,
	}
}

type ModuleDeclaration struct {
	TypeDeclaration

	version  string
	requires []intsyn.IModuleInfo
	exports  []intsyn.IPackageInfo
	opens    []intsyn.IPackageInfo
	uses     []string
	provides []intsyn.IServiceInfo
}

func (d *ModuleDeclaration) Version() string {
	return d.version
}

func (d *ModuleDeclaration) Requires() []intsyn.IModuleInfo {
	return d.requires
}

func (d *ModuleDeclaration) Exports() []intsyn.IPackageInfo {
	return d.exports
}

func (d *ModuleDeclaration) Opens() []intsyn.IPackageInfo {
	return d.opens
}

func (d *ModuleDeclaration) Uses() []string {
	return d.uses
}

func (d *ModuleDeclaration) Provides() []intsyn.IServiceInfo {
	return d.provides
}

func (d *ModuleDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitModuleDeclaration(d)
}

func (d *ModuleDeclaration) String() string {
	return fmt.Sprintf("ModuleDeclaration{%s}", d.internalTypeName)
}

func NewModuleInfo(name string, flags int, version string) intsyn.IModuleInfo {
	return &ModuleInfo{
		name:    name,
		flags:   flags,
		version: version,
	}
}

type ModuleInfo struct {
	name    string
	flags   int
	version string
}

func (i *ModuleInfo) Name() string {
	return i.name
}

func (i *ModuleInfo) Flags() int {
	return i.flags
}

func (i *ModuleInfo) Version() string {
	return i.version
}

func (i *ModuleInfo) String() string {
	msg := fmt.Sprintf("ModuleInfo{name=%s, flags=%d", i.name, i.flags)
	if i.version != "" {
		msg += fmt.Sprintf(", version=%s", i.version)
	}
	msg += "}"

	return msg
}

func NewPackageInfo(internalName string, flags int, moduleInfoNames []string) intsyn.IPackageInfo {
	return &PackageInfo{
		internalName:    internalName,
		flags:           flags,
		moduleInfoNames: moduleInfoNames,
	}
}

type PackageInfo struct {
	internalName    string
	flags           int
	moduleInfoNames []string
}

func (i *PackageInfo) InternalName() string {
	return i.internalName
}

func (i *PackageInfo) Flags() int {
	return i.flags
}

func (i *PackageInfo) ModuleInfoNames() []string {
	return i.moduleInfoNames
}

func (i *PackageInfo) String() string {
	msg := fmt.Sprintf("PackageInfo{internalName=%s, flags=%d", i.internalName, i.flags)
	if len(i.moduleInfoNames) > 0 {
		msg += fmt.Sprintf(", moduleInfoNames=%s", i.moduleInfoNames)
	}
	msg += "}"

	return msg
}

func NewServiceInfo(internalTypeName string, implementationTypeNames []string) intsyn.IServiceInfo {
	return &ServiceInfo{
		internalTypeName:        internalTypeName,
		implementationTypeNames: implementationTypeNames,
	}
}

type ServiceInfo struct {
	internalTypeName        string
	implementationTypeNames []string
}

func (i *ServiceInfo) InternalTypeName() string {
	return i.internalTypeName
}

func (i *ServiceInfo) ImplementationTypeNames() []string {
	return i.implementationTypeNames
}

func (i *ServiceInfo) String() string {
	msg := fmt.Sprintf("PackageInfo{internalTypeName=%s", i.internalTypeName)
	if len(i.implementationTypeNames) > 0 {
		msg += fmt.Sprintf(", implementationTypeNames=%s", i.implementationTypeNames)
	}
	msg += "}"

	return msg
}
