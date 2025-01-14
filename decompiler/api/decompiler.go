package api

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

type Decompiler interface {
	Decompiler(loader Loader, printer Printer, internalName string) error
	DecompilerWithConfig(loader Loader, printer Printer, internalName string, configuration map[string]interface{}) error
	DecompilerWithMessage(msg *model.Message) error
}
