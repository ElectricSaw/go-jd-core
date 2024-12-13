package utils

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func SliceInItemRemove(list []intsrv.ILocalVariable, item intsrv.ILocalVariable) []intsrv.ILocalVariable {
	for i := 0; i < len(list); i++ {
		if list[i] == item {
			newList := make([]intsrv.ILocalVariable, 0, len(list)-1)
			newList = append(newList, list[:i]...)
			newList = append(newList, list[i+1:]...)
			return newList
		}
	}
	return list
}

func ConvertField(list []intsrv.IClassFileFieldDeclaration) util.IList[intsrv.IClassFileMemberDeclaration] {
	ret := util.NewDefaultListWithCapacity[intsrv.IClassFileMemberDeclaration](len(list))
	for _, item := range list {
		ret.Add(item)
	}
	return ret
}

func ConvertMethod(list []intsrv.IClassFileConstructorOrMethodDeclaration) util.IList[intsrv.IClassFileMemberDeclaration] {
	ret := util.NewDefaultListWithCapacity[intsrv.IClassFileMemberDeclaration](len(list))
	for _, item := range list {
		ret.Add(item)
	}
	return ret
}

func ConvertTypes(list []intsrv.IClassFileTypeDeclaration) util.IList[intsrv.IClassFileMemberDeclaration] {
	ret := util.NewDefaultListWithCapacity[intsrv.IClassFileMemberDeclaration](len(list))
	for _, item := range list {
		ret.Add(item)
	}
	return ret
}

func ConvertTo(expr intmod.IExpression) intsrv.ILocalVariable {
	return expr.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)
}
