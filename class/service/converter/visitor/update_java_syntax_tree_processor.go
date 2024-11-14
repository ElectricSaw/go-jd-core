package visitor

import "bitbucket.org/coontec/javaClass/class/model/message"

type UpdateJavaSyntaxTreeProcessor struct {
}

func (p *UpdateJavaSyntaxTreeProcessor) Process(message message.Message) error {
	return nil
}
