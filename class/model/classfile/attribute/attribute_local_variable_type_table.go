package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewAttributeLocalVariableTypeTable(localVariableTypeTable []intcls.ILocalVariableType) intcls.IAttributeLocalVariableTypeTable {
	return &AttributeLocalVariableTypeTable{localVariableTypeTable: localVariableTypeTable}
}

type AttributeLocalVariableTypeTable struct {
	localVariableTypeTable []intcls.ILocalVariableType
}

func (a AttributeLocalVariableTypeTable) LocalVariableTypeTable() []intcls.ILocalVariableType {
	return a.localVariableTypeTable
}

func (a AttributeLocalVariableTypeTable) IsAttribute() bool {
	return true
}
