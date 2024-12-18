package processor

import "github.com/ElectricSaw/go-jd-core/decompiler/model/message"

type Processor interface {
	Process(message *message.Message) error
}
