package localvariable

import (
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
)

func NewRootFrame() *RootFrame {
	return &RootFrame{
		Frame: *NewFrame(nil, nil).(*Frame),
	}
}

type RootFrame struct {
	Frame
}

func (f *RootFrame) LocalVariable(index int) intsrv.ILocalVariableReference {
	if index < len(f.localVariableArray) {
		return f.localVariableArray[index]
	}
	return nil
}

func (f *RootFrame) UpdateLocalVariableInForStatements(typeMarker *utils.TypeMaker) {
	if f.children != nil {
		for _, child := range f.children {
			child.UpdateLocalVariableInForStatements(typeMarker)
		}
	}
}

func (f *RootFrame) CreateDeclarations(containsLineNumber bool) {
	if f.children != nil {
		for _, child := range f.children {
			child.CreateDeclarations(containsLineNumber)
		}
	}
}
