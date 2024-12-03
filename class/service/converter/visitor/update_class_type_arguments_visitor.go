package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

type UpdateClassTypeArgumentsVisitor struct {
	AbstractUpdateExpressionVisitor

	result intmod.ITypeArgument
}

func (v *UpdateClassTypeArgumentsVisitor) Init() {
	v.result = nil
}

func (v *UpdateClassTypeArgumentsVisitor) TypeArgument() intmod.ITypeArgument {
	return v.result
}

func (v *UpdateClassTypeArgumentsVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
	t := argument.Type()
	t.AcceptTypeArgumentVisitor(v)

	if v.result == t {
		v.result = argument
	} else {
		v.result = _type.NewWildcardExtendsTypeArgument(v.result.(intmod.IType))
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
	t := argument.Type()
	t.AcceptTypeArgumentVisitor(v)

	if v.result == t {
		v.result = argument
	} else {
		v.result = _type.NewWildcardSuperTypeArgument(v.result.(intmod.IType))
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {
	v.result = argument
}
func (v *UpdateClassTypeArgumentsVisitor) VisitWildcardTypeArgument(argument intmod.IWildcardTypeArgument) {
	v.result = argument
}

func (v *UpdateClassTypeArgumentsVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) { v.result = t }
func (v *UpdateClassTypeArgumentsVisitor) VisitGenericType(t intmod.IGenericType)     { v.result = t }

func (v *UpdateClassTypeArgumentsVisitor) VisitObjectType(t intmod.IObjectType) {
	typeArguments := t.TypeArguments()

	if typeArguments == nil {
		if t.InternalName() == _type.OtTypeClass.InternalName() {
			v.result = _type.OtTypeClassWildcard
		} else {
			v.result = t
		}
	} else {
		typeArguments.AcceptTypeArgumentVisitor(v)

		if v.result == typeArguments {
			v.result = t
		} else {
			v.result = t.CreateTypeWithArgs(typeArguments)
		}
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	t.OuterType().AcceptTypeArgumentVisitor(v)

	typeArguments := t.TypeArguments()

	if t.OuterType() == v.result {
		if typeArguments == nil {
			v.result = t
		} else {
			typeArguments.AcceptTypeArgumentVisitor(v)
			if v.result == typeArguments {
				v.result = t
			} else {
				v.result = t.CreateTypeWithArgs(v.result)
			}
		}
	} else {
		outerObjectType := v.result.(intmod.IObjectType)

		if typeArguments != nil {
			typeArguments.AcceptTypeArgumentVisitor(v)
			typeArguments = v.result
		}

		v.result = _type.NewInnerObjectTypeWithAll(t.InternalName(), t.QualifiedName(), t.Name(), typeArguments, t.Dimension(), outerObjectType)
	}
}

func (v *UpdateClassTypeArgumentsVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {
	size := arguments.Size()
	i := 0

	for i = 0; i < size; i++ {
		ta := arguments.Get(i).(intmod.ITypeArgument)
		ta.AcceptTypeArgumentVisitor(v)
		if v.result != ta {
			break
		}
	}

	if v.result != nil {
		if i == size {
			v.result = arguments
		} else {
			newTypes := _type.NewTypeArgumentsWithCapacity(size)
			newTypes.AddAll(arguments.ToSlice()[:i])
			newTypes.Add(v.result.(intmod.ITypeArgument))

			for i++; i < size; i++ {
				ta := arguments.Get(i).(intmod.ITypeArgument)
				ta.AcceptTypeArgumentVisitor(v)
				newTypes.Add(v.result.(intmod.ITypeArgument))
			}

			v.result = newTypes
		}
	}
}
