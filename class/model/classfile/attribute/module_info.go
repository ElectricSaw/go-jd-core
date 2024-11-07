package attribute

func NewModuleInfo(name string, flags int, version string) *ModuleInfo {
	return &ModuleInfo{name: name, flags: flags, version: version}
}

type ModuleInfo struct {
	name    string
	flags   int
	version string
}

func (m ModuleInfo) Name() string {
	return m.name
}

func (m ModuleInfo) Flags() int {
	return m.flags
}

func (m ModuleInfo) Version() string {
	return m.version
}
