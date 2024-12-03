package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewPackageInfo(internalName string, flags int, moduleInfoNames []string) intcls.IPackageInfo {
	return &PackageInfo{internalName, flags, moduleInfoNames}
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
