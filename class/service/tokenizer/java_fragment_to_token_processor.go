package tokenizer

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/message"
	"bitbucket.org/coontec/go-jd-core/class/service/tokenizer/visitor"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

type JavaFragmentToTokenProcessor struct {
}

func (p *JavaFragmentToTokenProcessor) Process(message *message.Message) error {
	fragments := message.Body.(util.IList[intmod.IJavaFragment])
	visit := visitor.NewTokenizeJavaFragmentVisitor(fragments.Size() * 3)

	for _, fragment := range fragments.ToSlice() {
		fragment.Accept(visit)
	}

	return nil
}
