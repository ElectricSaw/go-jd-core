package _type

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"bitbucket.org/coontec/javaClass/class/util"
)

func NewTypes() intsyn.ITypes {
	return &Types{}
}

type Types struct {
	AbstractType
	util.DefaultList[intsyn.IType]
}

func (t *Types) Size() int {
	return t.DefaultList.Size()
}

func (t *Types) IsTypes() bool {
	return true
}

func (t *Types) AcceptTypeVisitor(visitor intsyn.ITypeVisitor) {
	visitor.VisitTypes(t)
}
