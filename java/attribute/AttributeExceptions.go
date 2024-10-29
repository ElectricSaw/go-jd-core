package attribute

func NewAttributeExceptions(exceptionTypeNames []string) AttributeExceptions {
	return AttributeExceptions{exceptionTypeNames}
}

type AttributeExceptions struct {
	exceptionTypeNames []string
}

func (a AttributeExceptions) ExceptionTypeNames() []string {
	return a.exceptionTypeNames
}

func (a AttributeExceptions) attributeIgnoreFunc() {}
