package utils

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/class/service/converter/model/localvariable"
	"bitbucket.org/coontec/javaClass/class/service/converter/visitor"
)

type LocalVariableMaker struct {
	localVariableSet              localvariable.LocalVariableSet
	names                         []string
	blackListNames                []string
	currentFrame                  *localvariable.RootFrame
	localVariableCache            []localvariable.AbstractLocalVariable
	typeMaker                     TypeMaker
	typeBounds                    map[string]_type.IType
	formalParameters              declaration.FormalParameter
	populateBlackListNamesVisitor *visitor.PopulateBlackListNamesVisitor
	searchInTypeArgumentVisitor   *visitor.SearchInTypeArgumentVisitor
	createParameterVisitor        *visitor.CreateParameterVisitor
	createLocalVariableVisitor    *visitor.CreateLocalVariableVisitor
}
