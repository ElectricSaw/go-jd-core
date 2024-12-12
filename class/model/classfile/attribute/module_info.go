package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewModuleInfo(name string, flags int, version string) intcls.IModuleInfo {
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
