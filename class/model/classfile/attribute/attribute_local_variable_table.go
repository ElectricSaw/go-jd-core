package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewAttributeLocalVariableTable(localVariableTable []intcls.ILocalVariable) intcls.IAttributeLocalVariableTable {
	return &AttributeLocalVariableTable{localVariableTable}
}

type AttributeLocalVariableTable struct {
	localVariableTable []intcls.ILocalVariable
}

func (a AttributeLocalVariableTable) LocalVariableTable() []intcls.ILocalVariable {
	return a.localVariableTable
}

func (a AttributeLocalVariableTable) IsAttribute() bool {
	return true
}
