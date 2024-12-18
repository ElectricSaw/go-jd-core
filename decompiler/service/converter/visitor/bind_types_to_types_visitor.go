package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

func NewBindTypesToTypesVisitor() intsrv.IBindTypesToTypesVisitor {
	return &BindTypesToTypesVisitor{
		typeArgumentToTypeVisitor:               NewTypeArgumentToTypeVisitor(),
		bindTypeArgumentsToTypeArgumentsVisitor: NewBindTypeArgumentsToTypeArgumentsVisitor(),
	}
}

type BindTypesToTypesVisitor struct {
	_type.AbstractNopTypeVisitor

	typeArgumentToTypeVisitor               intsrv.ITypeArgumentToTypeVisitor
	bindTypeArgumentsToTypeArgumentsVisitor intsrv.IBindTypeArgumentsToTypeArgumentsVisitor
	bindings                                map[string]intmod.ITypeArgument
	result                                  intmod.IType
}

func (v *BindTypesToTypesVisitor) SetBindings(bindings map[string]intmod.ITypeArgument) {
	v.bindings = bindings
	v.bindTypeArgumentsToTypeArgumentsVisitor.SetBindings(v.bindings)
}

func (v *BindTypesToTypesVisitor) Init() {
	v.result = nil
}

func (v *BindTypesToTypesVisitor) Type() intmod.IType {
	return v.result
}

func (v *BindTypesToTypesVisitor) VisitPrimitiveType(t intmod.IPrimitiveType) {
	v.result = t
}

func (v *BindTypesToTypesVisitor) VisitObjectType(t intmod.IObjectType) {
	typeArguments := t.TypeArguments()

	if typeArguments == nil {
		v.result = t.(intmod.IType)
	} else {
		v.bindTypeArgumentsToTypeArgumentsVisitor.Init()
		typeArguments.AcceptTypeArgumentVisitor(v.bindTypeArgumentsToTypeArgumentsVisitor)
		ta := v.bindTypeArgumentsToTypeArgumentsVisitor.TypeArgument()
		if typeArguments == ta {
			v.result = t.(intmod.IType)
		} else {
			v.result = t.CreateTypeWithArgs(ta).(intmod.IType)
		}
	}
}

func (v *BindTypesToTypesVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {
	t.OuterType().AcceptTypeVisitor(v)
	typeArguments := t.TypeArguments()

	if t.OuterType().(intmod.IType) == v.result {
		if typeArguments == nil {
			v.result = t.(intmod.IType)
		} else {
			v.bindTypeArgumentsToTypeArgumentsVisitor.Init()
			typeArguments.AcceptTypeArgumentVisitor(v.bindTypeArgumentsToTypeArgumentsVisitor)
			ta := v.bindTypeArgumentsToTypeArgumentsVisitor.TypeArgument()
			if typeArguments == ta {
				v.result = t.(intmod.IType)
			} else {
				v.result = t.CreateTypeWithArgs(ta).(intmod.IType)
			}
		}
	} else {
		outerObjectType := v.result.(intmod.IObjectType)

		if typeArguments != nil {
			v.bindTypeArgumentsToTypeArgumentsVisitor.Init()
			typeArguments.AcceptTypeArgumentVisitor(v.bindTypeArgumentsToTypeArgumentsVisitor)
			typeArguments = v.bindTypeArgumentsToTypeArgumentsVisitor.TypeArgument()

			if _type.WildcardTypeArgumentEmpty == typeArguments {
				typeArguments = nil
			}
		}

		v.result = _type.NewInnerObjectTypeWithAll(t.InternalName(), t.QualifiedName(),
			t.Name(), typeArguments, t.Dimension(), outerObjectType).(intmod.IType)
	}
}

func (v *BindTypesToTypesVisitor) VisitGenericType(t intmod.IGenericType) {
	ta := v.bindings[t.Name()]

	if ta == nil {
		v.result = _type.OtTypeObject.CreateType(t.Dimension())
	} else if ta == _type.WildcardTypeArgumentEmpty {
		v.result = _type.OtTypeObject.CreateType(t.Dimension())
	} else {
		v.typeArgumentToTypeVisitor.Init()
		ta.AcceptTypeArgumentVisitor(v.typeArgumentToTypeVisitor)
		t2 := v.typeArgumentToTypeVisitor.Type()
		v.result = t2.CreateType(t2.Dimension() + t.Dimension())
	}
}

func (v *BindTypesToTypesVisitor) VisitTypes(types intmod.ITypes) {
	size := types.Size()
	i := 0

	for i = 0; i < size; i++ {
		t := types.Get(i)
		t.AcceptTypeVisitor(v)
		if v.result != t {
			break
		}
	}

	if i == size {
		v.result = types
	} else {
		newTypes := _type.NewTypes()
		newTypes.AddAll(types.ToSlice()[:i])
		newTypes.Add(v.result)

		for i++; i < size; i++ {
			t := types.Get(i)
			t.AcceptTypeVisitor(v)
			newTypes.Add(v.result)
		}

		v.result = newTypes
	}
}
