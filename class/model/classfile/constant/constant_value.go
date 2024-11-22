package constant

type ConstantValue interface {
	Tag() TAG
	constantValueIgnoreFunc()
}
