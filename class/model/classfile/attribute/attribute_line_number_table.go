package attribute

func NewAttributeLineNumberTable(lineNumberTable []LineNumber) *AttributeLineNumberTable {
	return &AttributeLineNumberTable{lineNumberTable}
}

type AttributeLineNumberTable struct {
	lineNumberTable []LineNumber
}

func (a AttributeLineNumberTable) LineNumberTable() []LineNumber {
	return a.lineNumberTable
}

func (a AttributeLineNumberTable) attributeIgnoreFunc() {}
