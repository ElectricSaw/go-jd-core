package localvariable

import (
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
)

func NewRootFrame() intsrv.IRootFrame {
	return &RootFrame{
		Frame: *NewFrame(nil, nil).(*Frame),
	}
}

type RootFrame struct {
	Frame
}

func (f *RootFrame) LocalVariable(index int) intsrv.ILocalVariableReference {
	if index < len(f.LocalVariableArray) {
		return f.LocalVariableArray[index]
	}
	return nil
}

func (f *RootFrame) UpdateLocalVariableInForStatements(typeMarker intsrv.ITypeMaker) {
	if f.Children != nil {
		for _, child := range f.Children {
			child.UpdateLocalVariableInForStatements(typeMarker)
		}
	}
}

func (f *RootFrame) CreateDeclarations(containsLineNumber bool) {
	if f.Children != nil {
		for _, child := range f.Children {
			child.CreateDeclarations(containsLineNumber)
		}
	}
}
