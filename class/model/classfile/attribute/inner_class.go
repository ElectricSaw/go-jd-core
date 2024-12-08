package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewInnerClass(innerTypeName string, outerTypeName string, innerName string, innerAccessFlags int) intcls.IInnerClass {
	return &InnerClass{innerTypeName, outerTypeName, innerName, innerAccessFlags}
}

type InnerClass struct {
	innerTypeName    string
	outerTypeName    string
	innerName        string
	innerAccessFlags int
}

func (i InnerClass) InnerTypeName() string {
	return i.innerTypeName
}

func (i InnerClass) OuterTypeName() string {
	return i.outerTypeName
}

func (i InnerClass) InnerName() string {
	return i.innerName
}

func (i InnerClass) InnerAccessFlags() int {
	return i.innerAccessFlags
}
