package visitor

import _type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"

type BindTypesToTypesVisitor struct {
	_type.AbstractNodeVisitor

	typeArgumentToTypeVisitor
}
