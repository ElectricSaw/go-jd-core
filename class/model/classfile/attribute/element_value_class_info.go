package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewElementValueClassInfo(classInfo string) intcls.IElementValueClassInfo {
	return &ElementValueClassInfo{
		classInfo: classInfo,
	}
}

type ElementValueClassInfo struct {
	classInfo string
}

func (e *ElementValueClassInfo) ClassInfo() string {
	return e.classInfo
}

func (e *ElementValueClassInfo) Accept(visitor intcls.IElementValueVisitor) {
	visitor.VisitClassInfo(e)
}
