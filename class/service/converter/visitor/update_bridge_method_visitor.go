package visitor

import (
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
)

func NewUpdateBridgeMethodVisitor(typeMaker *utils.TypeMaker) *UpdateBridgeMethodVisitor {
	return &UpdateBridgeMethodVisitor{
		typeMaker: typeMaker,
	}
}

type UpdateBridgeMethodVisitor struct {
	AbstractUpdateExpressionVisitor

	bodyDeclarationsVisitor  *BodyDeclarationsVisitor
	bridgeMethodDeclarations map[string]map[string]intsrv.IClassFileMethodDeclaration
	typeMaker                *utils.TypeMaker
}
