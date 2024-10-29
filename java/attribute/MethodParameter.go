package attribute

func NewMethodParameter(name string, access int) MethodParameter {
	return MethodParameter{name: name, access: access}
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
