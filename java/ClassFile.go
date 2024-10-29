package java

type ClassFile struct {
	MajorVersion       int
	MinorVersion       int
	AccessFlags        int
	InternalTypeName   string
	SuperTypeName      string
	InterfaceTypeNames []string
	Field              []interface{}          // []Field
	Method             []interface{}          // []Method
	Attributes         map[string]interface{} // map[string]Attribute
}
