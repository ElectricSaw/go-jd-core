package processor

import "bitbucket.org/coontec/javaClass/java/model/message"

type Processor interface {
	Process(message message.Message) error
}
