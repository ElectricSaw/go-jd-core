package utils

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/model/localvariable"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/visitor"
)

type LocalVariableMaker struct {
	localVariableSet              localvariable.LocalVariableSet
	names                         []string
	blackListNames                []string
	currentFrame                  *localvariable.RootFrame
	localVariableCache            []localvariable.AbstractLocalVariable
	typeMaker                     TypeMaker
	typeBounds                    map[string]intmod.IType
	formalParameters              declaration.FormalParameter
	populateBlackListNamesVisitor *visitor.PopulateBlackListNamesVisitor
	searchInTypeArgumentVisitor   *visitor.SearchInTypeArgumentVisitor
	createParameterVisitor        *visitor.CreateParameterVisitor
	createLocalVariableVisitor    *visitor.CreateLocalVariableVisitor
}
