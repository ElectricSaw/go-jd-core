package expression

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"

type AbstractNopExpressionVisitor struct {
}

func (v *AbstractNopExpressionVisitor) VisitArrayExpression(expression intsyn.IArrayExpression) {}
func (v *AbstractNopExpressionVisitor) VisitBinaryOperatorExpression(expression intsyn.IBinaryOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitBooleanExpression(expression intsyn.IBooleanExpression) {}
func (v *AbstractNopExpressionVisitor) VisitCastExpression(expression intsyn.ICastExpression)       {}
func (v *AbstractNopExpressionVisitor) VisitCommentExpression(expression intsyn.ICommentExpression) {}
func (v *AbstractNopExpressionVisitor) VisitConstructorInvocationExpression(expression intsyn.IConstructorInvocationExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitConstructorReferenceExpression(expression intsyn.IConstructorReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitDoubleConstantExpression(expression intsyn.IDoubleConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitEnumConstantReferenceExpression(expression intsyn.IEnumConstantReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitExpressions(expression intsyn.IExpressions) {}
func (v *AbstractNopExpressionVisitor) VisitFieldReferenceExpression(expression intsyn.IFieldReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitFloatConstantExpression(expression intsyn.IFloatConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitIntegerConstantExpression(expression intsyn.IIntegerConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitInstanceOfExpression(expression intsyn.IInstanceOfExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLambdaFormalParametersExpression(expression intsyn.ILambdaFormalParametersExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLambdaIdentifiersExpression(expression intsyn.ILambdaIdentifiersExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLengthExpression(expression intsyn.ILengthExpression) {}
func (v *AbstractNopExpressionVisitor) VisitLocalVariableReferenceExpression(expression intsyn.ILocalVariableReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLongConstantExpression(expression intsyn.ILongConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitMethodInvocationExpression(expression intsyn.IMethodInvocationExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitMethodReferenceExpression(expression intsyn.IMethodReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitNewArray(expression intsyn.INewArray)           {}
func (v *AbstractNopExpressionVisitor) VisitNewExpression(expression intsyn.INewExpression) {}
func (v *AbstractNopExpressionVisitor) VisitNewInitializedArray(expression intsyn.INewInitializedArray) {
}
func (v *AbstractNopExpressionVisitor) VisitNoExpression(expression intsyn.INoExpression)     {}
func (v *AbstractNopExpressionVisitor) VisitNullExpression(expression intsyn.INullExpression) {}
func (v *AbstractNopExpressionVisitor) VisitObjectTypeReferenceExpression(expression intsyn.IObjectTypeReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitParenthesesExpression(expression intsyn.IParenthesesExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitPostOperatorExpression(expression intsyn.IPostOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitPreOperatorExpression(expression intsyn.IPreOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitStringConstantExpression(expression intsyn.IStringConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitSuperConstructorInvocationExpression(expression intsyn.ISuperConstructorInvocationExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitSuperExpression(expression intsyn.ISuperExpression) {}
func (v *AbstractNopExpressionVisitor) VisitTernaryOperatorExpression(expression intsyn.ITernaryOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitThisExpression(expression intsyn.IThisExpression) {}
func (v *AbstractNopExpressionVisitor) VisitTypeReferenceDotClassExpression(expression intsyn.ITypeReferenceDotClassExpression) {
}
