package utils

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/visitor"
)

type LocalVariableMaker struct {
	localVariableSet              intsrv.ILocalVariableSet
	names                         []string
	blackListNames                []string
	currentFrame                  intsrv.IRootFrame
	localVariableCache            []intsrv.ILocalVariable
	typeMaker                     TypeMaker
	typeBounds                    map[string]intmod.IType
	formalParameters              declaration.FormalParameter
	populateBlackListNamesVisitor *visitor.PopulateBlackListNamesVisitor
	searchInTypeArgumentVisitor   *visitor.SearchInTypeArgumentVisitor
	createParameterVisitor        *visitor.CreateParameterVisitor
	createLocalVariableVisitor    *visitor.CreateLocalVariableVisitor
}
