package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewAttributeModule(name string, flags int, version string,
	requires []intcls.IModuleInfo, exports []intcls.IPackageInfo, opens []intcls.IPackageInfo,
	uses []string, provides []intcls.IServiceInfo) intcls.IAttributeModule {
	return &AttributeModule{name, flags, version, requires, exports, opens, uses, provides}
}

type AttributeModule struct {
	name    string
	flags   int
	version string

	requires []intcls.IModuleInfo
	exports  []intcls.IPackageInfo
	opens    []intcls.IPackageInfo
	uses     []string
	provides []intcls.IServiceInfo
}

func (a AttributeModule) Name() string {
	return a.name
}

func (a AttributeModule) Flags() int {
	return a.flags
}

func (a AttributeModule) Version() string {
	return a.version
}

func (a AttributeModule) Requires() []intcls.IModuleInfo {
	return a.requires
}

func (a AttributeModule) Exports() []intcls.IPackageInfo {
	return a.exports
}

func (a AttributeModule) Opens() []intcls.IPackageInfo {
	return a.opens
}

func (a AttributeModule) Uses() []string {
	return a.uses
}

func (a AttributeModule) Provides() []intcls.IServiceInfo {
	return a.provides
}

func (a AttributeModule) IsAttribute() bool {
	return true
}
