package attribute

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

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
