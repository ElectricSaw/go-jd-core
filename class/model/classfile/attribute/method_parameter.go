package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewMethodParameter(name string, access int) intcls.IMethodParameter {
	return &MethodParameter{name: name, access: access}
}

type MethodParameter struct {
	name   string
	access int
}

func (m MethodParameter) Name() string {
	return m.name
}

func (m MethodParameter) Access() int {
	return m.access
}
