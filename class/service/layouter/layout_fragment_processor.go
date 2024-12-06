package layouter

import (
	"bitbucket.org/coontec/go-jd-core/class/api"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/message"
	visitor2 "bitbucket.org/coontec/go-jd-core/class/service/layouter/visitor"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"math"
)

func NewLayoutFragmentProcessor() *LayoutFragmentProcessor {
	return &LayoutFragmentProcessor{}
}

type LayoutFragmentProcessor struct {
}

func (p *LayoutFragmentProcessor) Process(message *message.Message) error {
	var maxLineNumber int
	var ok, containsByteCode, showBridgeAndSynthetic bool
	var realignLineNumbers bool
	var realignLineNumbersConfiguration interface{}

	if maxLineNumber, ok = message.Headers["maxLineNumber"].(int); !ok {
		maxLineNumber = api.UnknownLineNumber
	}

	if containsByteCode, ok = message.Headers["containsByteCode"].(bool); !ok {
		containsByteCode = false
	}

	if containsByteCode, ok = message.Headers["showBridgeAndSynthetic"].(bool); !ok {
		containsByteCode = false
	}

	configuration := message.Headers["configuration"].(map[string]interface{})
	if configuration == nil {
		realignLineNumbersConfiguration = "false"
	} else {
		realignLineNumbersConfiguration = configuration["realignLineNumbers"].(bool)
	}

	if realignLineNumbersConfiguration == nil {
		switch meta := realignLineNumbersConfiguration.(type) {
		case string:
			realignLineNumbers = "true" == meta
		case bool:
			realignLineNumbers = true == meta
		}
	} else {
		realignLineNumbers = false
	}

	fragments := message.Body.(util.IList[intmod.IFragment])

	if maxLineNumber != api.UnknownLineNumber && !containsByteCode && !showBridgeAndSynthetic && realignLineNumbers {
		buildSectionsVisitor := visitor2.NewBuildSectionsVisitor()

		for _, fragment := range fragments.ToSlice() {
			fragment.AcceptFragmentVisitor(buildSectionsVisitor)
		}

		sections := buildSectionsVisitor.Sections()
		holder := visitor2.NewVisitorsHolder()
		visitor0 := visitor2.NewUpdateSpacerBetweenMovableBlocksVisitor()

		sumOfRates := math.MaxInt
		maximum := sections.Size() * 2

		if maximum > 20 {
			maximum = 20
		}

		for loop := 0; loop < maximum; loop++ {
			visitor0.Reset()

			for _, section := range sections.ToSlice() {
				for _, fragment := range section.FlexibleFragments().ToSlice() {
					fragment.AcceptFragmentVisitor(visitor0)
				}

				if section.FixedFragment() != nil {
					section.FixedFragment().AcceptFragmentVisitor(visitor0)
				}
			}

			for redo := 0; redo < 10; redo++ {
				changed := false
				for _, section := range sections.ToSlice() {
					changed = changed || section.Layout(false)
				}

				if !changed {
					break
				}
			}

			newSumOfRates := 0
			mostConstrainedSection := sections.Get(0)

			for _, section := range sections.ToSlice() {
				section.UpdateRate()

				if mostConstrainedSection.Rate() < section.Rate() {
					mostConstrainedSection = section
				}

				newSumOfRates += section.Rate()
			}

			if mostConstrainedSection.Rate() == 0 {
				break
			}

			if sumOfRates > newSumOfRates {
				sumOfRates = newSumOfRates
			} else {
				break
			}

			if !mostConstrainedSection.ReleaseConstraints(holder) {
				break
			}
		}

		for _, section := range sections.ToSlice() {
			section.Layout(true)
		}

		fragments.Clear()

		for _, section := range sections.ToSlice() {
			tmp := make([]intmod.IFragment, 0)
			for _, fragment := range section.FlexibleFragments().ToSlice() {
				tmp = append(tmp, fragment)
			}
			fragments.AddAll(tmp)

			fixedFragment := section.FixedFragment()

			if fixedFragment != nil {
				fragments.Add(fixedFragment)
			}
		}
	}

	return nil
}
