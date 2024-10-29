package attribute

func NewAttributeModulePackages(packageNames []string) AttributeModulePackages {
	return AttributeModulePackages{packageNames: packageNames}
}

type AttributeModulePackages struct {
	packageNames []string
}

func (a AttributeModulePackages) PackageNames() []string {
	return a.packageNames
}

func (a AttributeModulePackages) attributeIgnoreFunc() {}
