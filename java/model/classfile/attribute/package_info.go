package attribute

func NewPackageInfo(internalName string, flags int, moduleInfoNames []string) PackageInfo {
	return PackageInfo{internalName, flags, moduleInfoNames}
}

type PackageInfo struct {
	internalName    string
	flags           int
	moduleInfoNames []string
}

func (p PackageInfo) InternalName() string {
	return p.internalName
}

func (p PackageInfo) Flags() int {
	return p.flags
}

func (p PackageInfo) ModuleInfoNames() []string {
	return p.moduleInfoNames
}
