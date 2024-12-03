package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewBootstrapMethod(bootstrapMethodRef int, bootstrapArguments []int) intcls.IBootstrapMethod {
	return &BootstrapMethod{
		bootstrapMethodRef: bootstrapMethodRef,
		bootstrapArguments: bootstrapArguments,
	}
}

type BootstrapMethod struct {
	bootstrapMethodRef int
	bootstrapArguments []int
}

func (b BootstrapMethod) BootstrapMethodRef() int {
	return b.bootstrapMethodRef
}

func (b BootstrapMethod) BootstrapArguments() []int {
	return b.bootstrapArguments
}
