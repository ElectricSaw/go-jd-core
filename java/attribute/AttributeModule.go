package attribute

func NewAttributeModule(name string, flags int, version string, requires []ModuleInfo, exports []PackageInfo, opens []PackageInfo, uses []string, provides []ServiceInfo) AttributeModule {
	return AttributeModule{name, flags, version, requires, exports, opens, uses, provides}
}

type AttributeModule struct {
	name    string
	flags   int
	version string

	requires []ModuleInfo
	exports  []PackageInfo
	opens    []PackageInfo
	uses     []string
	provides []ServiceInfo
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

func (a AttributeModule) Requires() []ModuleInfo {
	return a.requires
}

func (a AttributeModule) Exports() []PackageInfo {
	return a.exports
}

func (a AttributeModule) Opens() []PackageInfo {
	return a.opens
}

func (a AttributeModule) Uses() []string {
	return a.uses
}

func (a AttributeModule) Provides() []ServiceInfo {
	return a.provides
}

func (a AttributeModule) attributeIgnoreFunc() {}
