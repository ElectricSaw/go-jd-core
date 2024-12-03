package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewAttributeLineNumberTable(lineNumberTable []intcls.ILineNumber) intcls.IAttributeLineNumberTable {
	return &AttributeLineNumberTable{lineNumberTable}
}

type AttributeLineNumberTable struct {
	lineNumberTable []intcls.ILineNumber
}

func (a AttributeLineNumberTable) LineNumberTable() []intcls.ILineNumber {
	return a.lineNumberTable
}

func (a AttributeLineNumberTable) IsAttribute() bool {
	return true
}
