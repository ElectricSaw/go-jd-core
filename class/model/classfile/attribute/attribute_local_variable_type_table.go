package attribute

func NewAttributeLocalVariableTypeTable(localVariableTypeTable []LocalVariableType) *AttributeLocalVariableTypeTable {
	return &AttributeLocalVariableTypeTable{localVariableTypeTable: localVariableTypeTable}
}

type AttributeLocalVariableTypeTable struct {
	localVariableTypeTable []LocalVariableType
}

func (a AttributeLocalVariableTypeTable) LocalVariableTypeTable() []LocalVariableType {
	return a.localVariableTypeTable
}

func (a AttributeLocalVariableTypeTable) attributeIgnoreFunc() {}
