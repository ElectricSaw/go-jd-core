package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"math"
)

func NewPopulateBindingsWithTypeArgumentVisitor(typeMaker intsrv.ITypeMaker) *PopulateBindingsWithTypeArgumentVisitor {
	return &PopulateBindingsWithTypeArgumentVisitor{
		typeArgumentToTypeVisitor: NewTypeArgumentToTypeVisitor(),
		typeMaker:                 typeMaker,
		current:                   nil,
	}
}

type PopulateBindingsWithTypeArgumentVisitor struct {
	typeArgumentToTypeVisitor *TypeArgumentToTypeVisitor
	typeMaker                 intsrv.ITypeMaker
	contextualTypeBounds      map[string]intmod.IType
	bindings                  map[string]intmod.ITypeArgument
	typeBounds                map[string]intmod.IType
	current                   intmod.ITypeArgument
}

func (v *PopulateBindingsWithTypeArgumentVisitor) Init(
	contextualTypeBounds map[string]intmod.IType,
	bindings map[string]intmod.ITypeArgument,
	typeBounds map[string]intmod.IType,
	typeArgument intmod.ITypeArgument) {
	v.contextualTypeBounds = contextualTypeBounds
	v.bindings = bindings
	v.typeBounds = typeBounds
	v.current = typeArgument
}

func (v *PopulateBindingsWithTypeArgumentVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {
	if (v.current != nil) && v.current.IsTypeArgumentList() {
		typeArgumentIterator := arguments.ToSlice()
		typeGenericArgumentIterator := v.current.TypeArgumentList()
		MaxLength := int(math.Min(float64(len(typeArgumentIterator)), float64(len(typeGenericArgumentIterator))))

		for i := 0; i < MaxLength; i++ {
			v.current = typeGenericArgumentIterator[i]
			typeArgumentIterator[i].AcceptTypeArgumentVisitor(v)
		}
	}
}

func (v *PopulateBindingsWithTypeArgumentVisitor) VisitGenericType(typ intmod.IGenericType) {
	typeName := typ.Name()

	if value, ok := v.bindings[typeName]; ok {
		typeArguement := value

		if v.current != nil {
			if v.current.IsGenericTypeArgument() &&
				!equals(v.contextualTypeBounds[typeName],
					v.typeBounds[v.current.(intmod.IGenericType).Name()]) {
				return
			}

			if typeArguement == nil {
				v.bindings[typeName] = v.checkTypeClassCheckDimensionAndReturnCurrentAsTypeArgument(typ)
			} else if v.current != typeArguement {
				v.typeArgumentToTypeVisitor.Init()
				typeArguement.AcceptTypeArgumentVisitor(v.typeArgumentToTypeVisitor)
				t1 := v.typeArgumentToTypeVisitor.Type()

				v.typeArgumentToTypeVisitor.Init()
				v.current.AcceptTypeArgumentVisitor(v.typeArgumentToTypeVisitor)
				t2 := v.typeArgumentToTypeVisitor.Type()

				if t1.CreateType(0) != t2.CreateType(0) {
					if t1.IsObjectType() && t2.IsObjectType() {
						ot1 := t1.(intmod.IObjectType)
						ot2 := t2.(intmod.IObjectType).CreateType(t2.Dimension() - t1.Dimension()).(intmod.IObjectType)

						if !v.typeMaker.IsAssignable(v.typeBounds, ot1, ot2) {
							if v.typeMaker.IsAssignable(v.typeBounds, ot2, ot1) {
								v.bindings[typeName] = v.checkTypeClassCheckDimensionAndReturnCurrentAsTypeArgument(typ)
							} else {
								v.bindings[typeName] = _type.WildcardTypeArgumentEmpty
							}
						}
					}
				}
			}
		}
	}
}

func (v *PopulateBindingsWithTypeArgumentVisitor) checkTypeClassCheckDimensionAndReturnCurrentAsTypeArgument(typ intmod.IGenericType) intmod.ITypeArgument {
	if v.current != nil {
		if v.current.IsObjectTypeArgument() {
			ot := v.current.(intmod.IObjectType)
			if ot.TypeArguments() == nil && ot.InternalName() == _type.OtTypeClass.InternalName() {
				return _type.OtTypeClassWildcard.CreateType(ot.Dimension() - typ.Dimension())
			}
			return ot.CreateType(ot.Dimension() - typ.Dimension())
		} else if v.current.IsInnerObjectTypeArgument() || v.current.IsGenericTypeArgument() || v.current.IsPrimitiveTypeArgument() {
			t := v.current.(intmod.IType)
			return t.CreateType(t.Dimension() - typ.Dimension())
		}
	}

	return v.current.TypeArgumentFirst()
}

func (v *PopulateBindingsWithTypeArgumentVisitor) VisitWildcardExtendsTypeArgument(typ intmod.IWildcardExtendsTypeArgument) {
	if v.current != nil {
		if v.current.IsWildcardExtendsTypeArgument() {
			v.current = v.current.Type()
			typ.Type().AcceptTypeArgumentVisitor(v)
		} else {
			typ.Type().AcceptTypeArgumentVisitor(v)
		}
	}
}

func (v *PopulateBindingsWithTypeArgumentVisitor) VisitWildcardSuperTypeArgument(typ intmod.IWildcardSuperTypeArgument) {
	if v.current != nil {
		if v.current.IsWildcardSuperTypeArgument() {
			v.current = v.current.Type()
			typ.Type().AcceptTypeArgumentVisitor(v)
		} else {
			typ.Type().AcceptTypeArgumentVisitor(v)
		}
	}
}

func (v *PopulateBindingsWithTypeArgumentVisitor) VisitObjectType(typ intmod.IObjectType) {
	if v.current != nil && typ.TypeArguments() != nil {
		if v.current.IsObjectTypeArgument() || v.current.IsInnerObjectTypeArgument() {
			v.current = v.current.(intmod.IObjectType).TypeArguments()
			typ.TypeArguments().AcceptTypeArgumentVisitor(v)
		}
	}
}

func (v *PopulateBindingsWithTypeArgumentVisitor) VisitInnerObjectType(typ intmod.IInnerObjectType) {
	if v.current != nil && typ.TypeArguments() != nil && v.current.IsInnerObjectTypeArgument() {
		v.current = v.current.(intmod.IInnerObjectType).TypeArguments()
		typ.TypeArguments().AcceptTypeArgumentVisitor(v)
	}
}

func (v *PopulateBindingsWithTypeArgumentVisitor) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {
}
func (v *PopulateBindingsWithTypeArgumentVisitor) VisitWildcardTypeArgument(typ intmod.IWildcardTypeArgument) {
}
func (v *PopulateBindingsWithTypeArgumentVisitor) VisitPrimitiveType(typ intmod.IPrimitiveType) {}

func equals(bt1, bt2 intmod.IType) bool {
	return bt2 == nil || bt2 == bt1
}
