package declaration

import (
	"fmt"
)

func NewModuleDeclaration(flags int, internalTypeName, name, version string, requires []ModuleInfo, exports []PackageInfo, opens []PackageInfo, uses []string, provides []ServiceInfo) *ModuleDeclaration {
	return &ModuleDeclaration{
		TypeDeclaration: *NewTypeDeclaration(nil, flags, internalTypeName, name, nil),
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
	requires []ModuleInfo
	exports  []PackageInfo
	opens    []PackageInfo
	uses     []string
	provides []ServiceInfo
}

func (d *ModuleDeclaration) GetVersion() string {
	return d.version
}

func (d *ModuleDeclaration) GetRequires() []ModuleInfo {
	return d.requires
}

func (d *ModuleDeclaration) GetExports() []PackageInfo {
	return d.exports
}

func (d *ModuleDeclaration) GetOpens() []PackageInfo {
	return d.opens
}

func (d *ModuleDeclaration) GetUses() []string {
	return d.uses
}

func (d *ModuleDeclaration) GetProvides() []ServiceInfo {
	return d.provides
}

func (d *ModuleDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitModuleDeclaration(d)
}

func (d *ModuleDeclaration) String() string {
	return fmt.Sprintf("ModuleDeclaration{%s}", d.internalTypeName)
}

func NewModuleInfo(name string, flags int, version string) *ModuleInfo {
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

func (i *ModuleInfo) GetName() string {
	return i.name
}

func (i *ModuleInfo) GetFlags() int {
	return i.flags
}

func (i *ModuleInfo) GetVersion() string {
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

func NewPackageInfo(internalName string, flags int, moduleInfoNames []string) *PackageInfo {
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

func (i *PackageInfo) GetInternalName() string {
	return i.internalName
}

func (i *PackageInfo) GetFlags() int {
	return i.flags
}

func (i *PackageInfo) GetModuleInfoNames() []string {
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

func NewServiceInfo(internalTypeName string, implementationTypeNames []string) *ServiceInfo {
	return &ServiceInfo{
		internalTypeName:        internalTypeName,
		implementationTypeNames: implementationTypeNames,
	}
}

type ServiceInfo struct {
	internalTypeName        string
	implementationTypeNames []string
}

func (i *ServiceInfo) GetInternalTypeName() string {
	return i.internalTypeName
}

func (i *ServiceInfo) GetImplementationTypeNames() []string {
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
