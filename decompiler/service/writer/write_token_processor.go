package writer

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewWriteTokenProcessor() *WriteTokenProcessor {
	return &WriteTokenProcessor{}
}

type WriteTokenProcessor struct {
}

func (p *WriteTokenProcessor) Process(message *model.Message) error {
	return nil
}
