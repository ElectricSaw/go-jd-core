package _type

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"

type AbstractNopTypeVisitor struct {
}

func (v *AbstractNopTypeVisitor) VisitPrimitiveType(y intsyn.IPrimitiveType)     {}
func (v *AbstractNopTypeVisitor) VisitObjectType(y intsyn.IObjectType)           {}
func (v *AbstractNopTypeVisitor) VisitInnerObjectType(y intsyn.IInnerObjectType) {}
func (v *AbstractNopTypeVisitor) VisitTypes(types intsyn.ITypes)                 {}
func (v *AbstractNopTypeVisitor) VisitGenericType(y intsyn.IGenericType)         {}
