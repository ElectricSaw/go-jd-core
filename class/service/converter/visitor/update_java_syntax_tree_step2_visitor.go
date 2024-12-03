package visitor

import (
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
)

var aggregateFieldsVisitor = NewAggregateFieldsVisitor()
var sortMembersVisitor = NewSortMembersVisitor()
var autoboxingVisitor = NewAutoboxingVisitor()

func NewUpdateJavaSyntaxTreeStep2Visitor(typeMaker intsrv.ITypeMaker) *UpdateJavaSyntaxTreeStep2Visitor {
	return &UpdateJavaSyntaxTreeStep2Visitor{
		updateOuterFieldTypeVisitor:   *NewUpdateOuterFieldTypeVisitor(typeMaker),
		updateBridgeMethodTypeVisitor: *NewUpdateBridgeMethodTypeVisitor(typeMaker),
	}
}

type UpdateJavaSyntaxTreeStep2Visitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	initStaticFieldVisitor          *InitStaticFieldVisitor
	initInstanceFieldVisitor        *InitInstanceFieldVisitor
	initEnumVisitor                 *InitEnumVisitor
	removeDefaultConstructorVisitor *RemoveDefaultConstructorVisitor
	replaceBridgeMethodVisitor      *UpdateBridgeMethodVisitor
	initInnerClassStep2Visitor      *UpdateNewExpressionVisitor
	addCastExpressionVisitor        *AddCastExpressionVisitor
}
