package expression

type AbstractNopExpressionVisitor struct {
}

func (v *AbstractNopExpressionVisitor) VisitArrayExpression(expression *ArrayExpression) {}
func (v *AbstractNopExpressionVisitor) VisitBinaryOperatorExpression(expression *BinaryOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitBooleanExpression(expression *BooleanExpression) {}
func (v *AbstractNopExpressionVisitor) VisitCastExpression(expression *CastExpression)       {}
func (v *AbstractNopExpressionVisitor) VisitCommentExpression(expression *CommentExpression) {}
func (v *AbstractNopExpressionVisitor) VisitConstructorInvocationExpression(expression *ConstructorInvocationExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitConstructorReferenceExpression(expression *ConstructorReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitDoubleConstantExpression(expression *DoubleConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitEnumConstantReferenceExpression(expression *EnumConstantReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitExpressions(expression *Expressions) {}
func (v *AbstractNopExpressionVisitor) VisitFieldReferenceExpression(expression *FieldReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitFloatConstantExpression(expression *FloatConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitIntegerConstantExpression(expression *IntegerConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitInstanceOfExpression(expression *InstanceOfExpression) {}
func (v *AbstractNopExpressionVisitor) VisitLambdaFormalParametersExpression(expression *LambdaFormalParametersExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLambdaIdentifiersExpression(expression *LambdaIdentifiersExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLengthExpression(expression *LengthExpression) {}
func (v *AbstractNopExpressionVisitor) VisitLocalVariableReferenceExpression(expression *LocalVariableReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitLongConstantExpression(expression *LongConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitMethodInvocationExpression(expression *MethodInvocationExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitMethodReferenceExpression(expression *MethodReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitNewArray(expression *NewArray)                       {}
func (v *AbstractNopExpressionVisitor) VisitNewExpression(expression *NewExpression)             {}
func (v *AbstractNopExpressionVisitor) VisitNewInitializedArray(expression *NewInitializedArray) {}
func (v *AbstractNopExpressionVisitor) VisitNoExpression(expression *NoExpression)               {}
func (v *AbstractNopExpressionVisitor) VisitNullExpression(expression *NullExpression)           {}
func (v *AbstractNopExpressionVisitor) VisitObjectTypeReferenceExpression(expression *ObjectTypeReferenceExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitParenthesesExpression(expression *ParenthesesExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitPostOperatorExpression(expression *PostOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitPreOperatorExpression(expression *PreOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitStringConstantExpression(expression *StringConstantExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitSuperConstructorInvocationExpression(expression *SuperConstructorInvocationExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitSuperExpression(expression *SuperExpression) {}
func (v *AbstractNopExpressionVisitor) VisitTernaryOperatorExpression(expression *TernaryOperatorExpression) {
}
func (v *AbstractNopExpressionVisitor) VisitThisExpression(expression *ThisExpression) {}
func (v *AbstractNopExpressionVisitor) VisitTypeReferenceDotClassExpression(expression *TypeReferenceDotClassExpression) {
}
