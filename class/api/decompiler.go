package api

type Decompiler interface {
	Decompiler(loader Loader, printer Printer, internalName string) error
	DecompilerWithConfig(loader Loader, printer Printer, internalName string, configuration map[string]interface{}) error
}
