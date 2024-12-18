package attribute

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewCodeException(index int, startPc int, endPc int, handlerPc int, catchType int) intcls.ICodeException {
	return &CodeException{index, startPc, endPc, handlerPc, catchType}
}

type CodeException struct {
	index     int
	startPc   int
	endPc     int
	handlerPc int
	catchType int
}

func (c CodeException) Index() int {
	return c.index
}

func (c CodeException) StartPc() int {
	return c.startPc
}

func (c CodeException) EndPc() int {
	return c.endPc
}

func (c CodeException) HandlerPc() int {
	return c.handlerPc
}

func (c CodeException) CatchType() int {
	return c.catchType
}
