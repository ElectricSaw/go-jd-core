package class

import (
	"github.com/ElectricSaw/go-jd-core/class/api"
	"github.com/ElectricSaw/go-jd-core/class/model/message"
	"github.com/ElectricSaw/go-jd-core/class/service/converter"
	"github.com/ElectricSaw/go-jd-core/class/service/deserializer"
	"github.com/ElectricSaw/go-jd-core/class/service/fragmenter"
	"github.com/ElectricSaw/go-jd-core/class/service/layouter"
	"github.com/ElectricSaw/go-jd-core/class/service/tokenizer"
	"github.com/ElectricSaw/go-jd-core/class/service/writer"
)

func NewClassFileToJavaSourceDecompiler() *ClassFileToJavaSourceDecompiler {
	return &ClassFileToJavaSourceDecompiler{
		dsr: deserializer.NewDeserializeClassFileProcessor(),
		cnv: converter.NewClassFileToJavaSyntaxProcessor(),
		frg: fragmenter.NewJavaSyntaxToJavaFragmentProcessor(),
		lyt: layouter.NewLayoutFragmentProcessor(),
		tkn: tokenizer.NewJavaFragmentToTokenProcessor(),
		wri: writer.NewWriteTokenProcessor(),
	}
}

type ClassFileToJavaSourceDecompiler struct {
	dsr *deserializer.DeserializeClassFileProcessor
	cnv *converter.ClassFileToJavaSyntaxProcessor
	frg *fragmenter.JavaSyntaxToJavaFragmentProcessor
	lyt *layouter.LayoutFragmentProcessor
	tkn *tokenizer.JavaFragmentToTokenProcessor
	wri *writer.WriteTokenProcessor
}

func (d *ClassFileToJavaSourceDecompiler) Decompiler(loader api.Loader,
	printer api.Printer, internalName string) error {
	msg := message.NewMessage()

	msg.Headers["mainInternalTypeName"] = internalName
	msg.Headers["loader"] = loader
	msg.Headers["printer"] = printer

	return d.DecompilerWithMessage(msg)
}

func (d *ClassFileToJavaSourceDecompiler) DecompilerWithConfig(loader api.Loader, printer api.Printer,
	internalName string, configuration map[string]interface{}) error {
	msg := message.NewMessage()

	msg.Headers["mainInternalTypeName"] = internalName
	msg.Headers["configuration"] = configuration
	msg.Headers["loader"] = loader
	msg.Headers["printer"] = printer

	return d.DecompilerWithMessage(msg)
}

func (d *ClassFileToJavaSourceDecompiler) DecompilerWithMessage(msg *message.Message) error {
	if err := d.dsr.Process(msg); err != nil {
		return err
	}
	if err := d.cnv.Process(msg); err != nil {
		return err
	}
	if err := d.frg.Process(msg); err != nil {
		return err
	}
	if err := d.lyt.Process(msg); err != nil {
		return err
	}
	if err := d.tkn.Process(msg); err != nil {
		return err
	}
	if err := d.wri.Process(msg); err != nil {
		return err
	}
	return nil
}
