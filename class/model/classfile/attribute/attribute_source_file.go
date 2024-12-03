package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewAttributeSourceFile(sourceFile string) intcls.IAttributeSourceFile {
	return &AttributeSourceFile{sourceFile}
}

type AttributeSourceFile struct {
	sourceFile string
}

func (a AttributeSourceFile) SourceFile() string {
	return a.sourceFile
}

func (a AttributeSourceFile) IsAttribute() bool {
	return true
}
