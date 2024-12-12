package processor

import "github.com/ElectricSaw/go-jd-core/class/model/message"

type Processor interface {
	Process(message *message.Message) error
}
