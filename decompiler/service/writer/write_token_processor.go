package writer

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model/message"
)

func NewWriteTokenProcessor() *WriteTokenProcessor {
	return &WriteTokenProcessor{}
}

type WriteTokenProcessor struct {
}

func (p *WriteTokenProcessor) Process(message *message.Message) error {
	return nil
}