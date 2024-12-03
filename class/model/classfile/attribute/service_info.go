package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewServiceInfo(interfaceTypeName string, implementationTypeNames []string) intcls.IServiceInfo {
	return &ServiceInfo{interfaceTypeName, implementationTypeNames}
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
