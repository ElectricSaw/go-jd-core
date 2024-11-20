package visitor

import "bitbucket.org/coontec/go-jd-core/class/model/message"

type UpdateJavaSyntaxTreeStep1Visitor struct {
}

func (p *UpdateJavaSyntaxTreeStep1Visitor) Process(message message.Message) error {
	return nil
}
