package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

func NewBindTypeArgumentsToTypeArgumentsVisitor() intsrv.IBindTypeArgumentsToTypeArgumentsVisitor {
	return &BindTypeArgumentsToTypeArgumentsVisitor{
		typeArgumentToTypeVisitor: NewTypeArgumentToTypeVisitor(),
	}
}

type BindTypeArgumentsToTypeArgumentsVisitor struct {
	typeArgumentToTypeVisitor intsrv.ITypeArgumentToTypeVisitor
	bindings                  map[string]intmod.ITypeArgument
	result                    intmod.ITypeArgument
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) Init() {
	v.result = nil
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) SetBindings(bindings map[string]intmod.ITypeArgument) {
	v.bindings = bindings
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) TypeArgument() intmod.ITypeArgument {
	if v.result == nil || _type.OtTypeObject == v.result {
		return nil
	}
	return v.result
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {
	size := arguments.Size()
	i := 0

	for i = 0; i < size; i++ {
		ta := arguments.Get(i)
		ta.AcceptTypeArgumentVisitor(v)
		if v.result != ta {
			break
		}
	}

	if v.result != nil {
		if i == size {
			v.result = arguments
		} else {
			newTypes := _type.NewTypeArguments()

			newTypes.AddAll(arguments.ToSlice()[:i])
			newTypes.Add(v.result)

			for i++; i < size; i++ {
				ta := arguments.Get(i)
				ta.AcceptTypeArgumentVisitor(v)
				if v.result == nil {
					return
				}
				newTypes.Add(v.result)
			}

			v.result = newTypes
		}
	}
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {
	v.result = argument
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
	argument.Type().AcceptTypeArgumentVisitor(v)

	if v.result == _type.WildcardTypeArgumentEmpty {
		v.result = _type.WildcardTypeArgumentEmpty
	} else if v.result == argument.Type() {
		v.result = argument
	} else if _type.OtTypeObject == v.result {
		v.result = _type.WildcardTypeArgumentEmpty
	} else if v.result != nil {
		v.typeArgumentToTypeVisitor.Init()
		v.result.AcceptTypeArgumentVisitor(v.typeArgumentToTypeVisitor)
		bt := v.typeArgumentToTypeVisitor.Type()

		if _type.OtTypeObject.(intmod.IType) == bt {
			v.result = _type.WildcardTypeArgumentEmpty
		} else {
			v.result = _type.NewWildcardExtendsTypeArgument(bt)
		}
	}
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) {
	v.result = t
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) VisitObjectType(t intmod.IObjectType) {
	typeArguments := t.TypeArguments()

	if typeArguments == nil {
		v.result = t
	} else {
		typeArguments.AcceptTypeArgumentVisitor(v)

		if typeArguments == v.result {
			v.result = t
		} else if v.result != nil {
			v.result = t.CreateTypeWithArgs(v.result)
		}
	}
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	t.OuterType().AcceptTypeArgumentVisitor(v)

	typeArguments := t.TypeArguments()

	if t.OuterType() == v.result {
		if typeArguments == nil {
			v.result = t
		} else {
			typeArguments.AcceptTypeArgumentVisitor(v)

			if typeArguments == v.result {
				v.result = t
			} else if v.result != nil {
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

func (v *BindTypeArgumentsToTypeArgumentsVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
	argument.Type().AcceptTypeArgumentVisitor(v)

	if v.result == _type.WildcardTypeArgumentEmpty {
		v.result = _type.WildcardTypeArgumentEmpty
	} else if v.result == argument.Type() {
		v.result = argument
	} else if v.result != nil {
		v.typeArgumentToTypeVisitor.Init()
		v.result.AcceptTypeArgumentVisitor(v.typeArgumentToTypeVisitor)
		v.result = _type.NewWildcardSuperTypeArgument(v.typeArgumentToTypeVisitor.Type())
	}
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) VisitGenericType(t intmod.IGenericType) {
	ta := v.bindings[t.Name()]

	if ta == nil {
		v.result = nil
	} else if _, ok := ta.(intmod.IType); ok {
		v.result = ta.(intmod.IType).CreateType(t.Dimension())
	} else {
		v.result = ta
	}
}

func (v *BindTypeArgumentsToTypeArgumentsVisitor) VisitWildcardTypeArgument(argument intmod.IWildcardTypeArgument) {
	v.result = argument
}
