package tokenizer

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/message"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/tokenizer/visitor"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewJavaFragmentToTokenProcessor() *JavaFragmentToTokenProcessor {
	return &JavaFragmentToTokenProcessor{}
}

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
