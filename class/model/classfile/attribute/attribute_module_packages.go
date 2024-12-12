package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewAttributeModulePackages(packageNames []string) intcls.IAttributeModulePackages {
	return &AttributeModulePackages{packageNames: packageNames}
}

type AttributeModulePackages struct {
	packageNames []string
}

func (a AttributeModulePackages) PackageNames() []string {
	return a.packageNames
}

func (a AttributeModulePackages) IsAttribute() bool {
	return true
}
