package visitor

import "bitbucket.org/coontec/go-jd-core/class/model/message"

type UpdateJavaSyntaxTreeProcessor struct {
}

func (p *UpdateJavaSyntaxTreeProcessor) Process(message message.Message) error {
	return nil
}
