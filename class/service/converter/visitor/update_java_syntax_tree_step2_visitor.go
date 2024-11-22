package visitor

import (
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
)

var aggregateFieldsVisitor = NewAggregateFieldsVisitor()
var sortMembersVisitor = NewSortMembersVisitor()
var autoboxingVisitor = NewAutoboxingVisitor()

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
