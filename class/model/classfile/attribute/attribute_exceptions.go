package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewAttributeExceptions(exceptionTypeNames []string) intcls.IAttributeExceptions {
	return &AttributeExceptions{exceptionTypeNames}
}

type AttributeExceptions struct {
	exceptionTypeNames []string
}

func (a AttributeExceptions) ExceptionTypeNames() []string {
	return a.exceptionTypeNames
}

func (a AttributeExceptions) IsAttribute() bool {
	return true
}
