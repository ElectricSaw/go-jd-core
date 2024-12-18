package api

type Loader interface {
	Load(internalName string) ([]byte, error)
	CanLoad(internalName string) bool
}
