package attribute

func NewAttributeLocalVariableTable(localVariableTable []LocalVariable) *AttributeLocalVariableTable {
	return &AttributeLocalVariableTable{localVariableTable}
}

type AttributeLocalVariableTable struct {
	localVariableTable []LocalVariable
}

func (a AttributeLocalVariableTable) LocalVariableTable() []LocalVariable {
	return a.localVariableTable
}

func (a AttributeLocalVariableTable) attributeIgnoreFunc() {}
