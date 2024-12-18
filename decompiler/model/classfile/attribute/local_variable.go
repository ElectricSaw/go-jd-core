package attribute

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewLocalVariable(startPc int, length int, name string, descriptor string, index int) intcls.ILocalVariable {
	return &LocalVariable{startPc, length, name, descriptor, index}
}

type LocalVariable struct {
	startPc    int
	length     int
	name       string
	descriptor string
	index      int
}

func (l LocalVariable) StartPc() int {
	return l.startPc
}

func (l LocalVariable) Length() int {
	return l.length
}

func (l LocalVariable) Name() string {
	return l.name
}

func (l LocalVariable) Descriptor() string {
	return l.descriptor
}

func (l LocalVariable) Index() int {
	return l.index
}
