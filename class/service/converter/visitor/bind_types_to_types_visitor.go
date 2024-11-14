package visitor

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

type BindTypesToTypesVisitor struct {
	_type.AbstractNodeVisitor

	typeArgumentToTypeVisitor
}
