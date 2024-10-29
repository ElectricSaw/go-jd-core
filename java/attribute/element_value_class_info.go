package attribute

func NewElementValueClassInfo(classInfo string) ElementValueClassInfo {
	return ElementValueClassInfo{
		classInfo: classInfo,
	}
}

type ElementValueClassInfo struct {
	classInfo string
}

func (e ElementValueClassInfo) ClassInfo() string {
	return e.classInfo
}

func (e *ElementValueClassInfo) Accept(visitor ElementValueVisitor) {
	visitor.VisitClassInfo(e)
}
