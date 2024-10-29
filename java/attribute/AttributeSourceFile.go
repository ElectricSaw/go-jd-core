package attribute

func NewAttributeSourceFile(sourceFile string) AttributeSourceFile {
	return AttributeSourceFile{sourceFile}
}

type AttributeSourceFile struct {
	sourceFile string
}

func (a AttributeSourceFile) SourceFile() string {
	return a.sourceFile
}

func (a AttributeSourceFile) attributeIgnoreFunc() {}
