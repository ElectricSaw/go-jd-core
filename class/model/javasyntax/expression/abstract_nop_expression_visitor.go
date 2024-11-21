package expression

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

type AbstractNopExpressionVisitor struct {
}

func (v *AbstractNopExpressionVisitor) VisitArrayExpression(expression intmod.IArrayExpression) {}
func (v *AbstractNopExpressionVisitor) VisitBinaryOperatorExpression(expression intmod.IBinaryOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitBooleanExpression(expression intmod.IBooleanExpression) {}
func (v *AbstractNopExpressionVisitor) VisitCastExpression(expression intmod.ICastExpression)       {}
func (v *AbstractNopExpressionVisitor) VisitCommentExpression(expression intmod.ICommentExpression) {}
func (v *AbstractNopExpressionVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitConstructorReferenceExpression(expression intmod.IConstructorReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitDoubleConstantExpression(expression intmod.IDoubleConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitEnumConstantReferenceExpression(expression intmod.IEnumConstantReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitExpressions(expression intmod.IExpressions) {}
func (v *AbstractNopExpressionVisitor) VisitFieldReferenceExpression(expression intmod.IFieldReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitFloatConstantExpression(expression intmod.IFloatConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitIntegerConstantExpression(expression intmod.IIntegerConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitInstanceOfExpression(expression intmod.IInstanceOfExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLambdaFormalParametersExpression(expression intmod.ILambdaFormalParametersExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLambdaIdentifiersExpression(expression intmod.ILambdaIdentifiersExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLengthExpression(expression intmod.ILengthExpression) {}
func (v *AbstractNopExpressionVisitor) VisitLocalVariableReferenceExpression(expression intmod.ILocalVariableReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLongConstantExpression(expression intmod.ILongConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitMethodInvocationExpression(expression intmod.IMethodInvocationExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitMethodReferenceExpression(expression intmod.IMethodReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitNewArray(expression intmod.INewArray)           {}
func (v *AbstractNopExpressionVisitor) VisitNewExpression(expression intmod.INewExpression) {}
func (v *AbstractNopExpressionVisitor) VisitNewInitializedArray(expression intmod.INewInitializedArray) {
}
func (v *AbstractNopExpressionVisitor) VisitNoExpression(expression intmod.INoExpression)     {}
func (v *AbstractNopExpressionVisitor) VisitNullExpression(expression intmod.INullExpression) {}
func (v *AbstractNopExpressionVisitor) VisitObjectTypeReferenceExpression(expression intmod.IObjectTypeReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitParenthesesExpression(expression intmod.IParenthesesExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitPostOperatorExpression(expression intmod.IPostOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitPreOperatorExpression(expression intmod.IPreOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitStringConstantExpression(expression intmod.IStringConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitSuperConstructorInvocationExpression(expression intmod.ISuperConstructorInvocationExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitSuperExpression(expression intmod.ISuperExpression) {}
func (v *AbstractNopExpressionVisitor) VisitTernaryOperatorExpression(expression intmod.ITernaryOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitThisExpression(expression intmod.IThisExpression) {}
func (v *AbstractNopExpressionVisitor) VisitTypeReferenceDotClassExpression(expression intmod.ITypeReferenceDotClassExpression) {
}
