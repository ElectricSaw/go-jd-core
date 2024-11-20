package localvariable

import "bitbucket.org/coontec/go-jd-core/class/service/converter/utils"

func NewRootFrame() *RootFrame {
	return &RootFrame{
		Frame: *NewFrame(nil, nil),
	}
}

type RootFrame struct {
	Frame
}

func (f *RootFrame) LocalVariable(index int) ILocalVariableReference {
	if index < len(f.localVariableArray) {
		return f.localVariableArray[index]
	}
	return nil
}

func (f *RootFrame) UpdateLocalVariableInForStatements(typeMarker utils.TypeMaker) {

}

func (f *RootFrame) CreateDeclarations(containsLineNumber bool) {
	if f.children != nil {
		for _, child := range f.children {
			child.createDeclarations(containsLineNumber)
		}
	}
}
