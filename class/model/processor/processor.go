package processor

import "bitbucket.org/coontec/go-jd-core/class/model/message"

type Processor interface {
	Process(message *message.Message) error
}
