package visitor

import (
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
)

var aggregateFieldsVisitor = NewAggregateFieldsVisitor()
var sortMembersVisitor = NewSortMembersVisitor()
var autoboxingVisitor = NewAutoboxingVisitor()

func NewUpdateJavaSyntaxTreeStep2Visitor(typeMaker *utils.TypeMaker) *UpdateJavaSyntaxTreeStep2Visitor {
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
