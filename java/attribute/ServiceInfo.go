package attribute

func NewServiceInfo(interfaceTypeName string, implementationTypeNames []string) ServiceInfo {
	return ServiceInfo{interfaceTypeName, implementationTypeNames}
}

type ServiceInfo struct {
	interfaceTypeName       string
	implementationTypeNames []string
}

func (s ServiceInfo) InterfaceTypeName() string {
	return s.interfaceTypeName
}

func (s ServiceInfo) ImplementationTypeNames() []string {
	return s.implementationTypeNames
}
