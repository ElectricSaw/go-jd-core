package processor

import "bitbucket.org/coontec/javaClass/class/model/message"

type Processor interface {
	Process(message message.Message) error
}
