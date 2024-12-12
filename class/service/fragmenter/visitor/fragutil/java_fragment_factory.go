package fragutil

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/model/javafragment"
	"github.com/ElectricSaw/go-jd-core/class/util"
	"math"
)

func AddSpacerAfterPackage(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpacerFragment(0, 1, 1, 0, "Spacer after package"))
	fragments.Add(javafragment.NewSpacerFragment(0, 1, 1, 1, "Second spacer after package"))
}

func AddSpacerAfterImports(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpacerFragment(0, 1, 1, 1, "Spacer after imports"))
}

func AddSpacerBeforeMainDeclaration(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 5, "Spacer before main declaration"))
}

func AddEndArrayInitializerInParameter(fragments util.IList[intmod.IFragment], start intmod.IStartBlockFragment) {
	fragments.Add(javafragment.NewEndBlockInParameterFragment(0, 0, 1, 20, "End array initializer", start))
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 0, math.MaxInt32, 21, "End array initializer spacer in parameter"))
}

func AddEndArrayInitializer(fragments util.IList[intmod.IFragment], start intmod.IStartBlockFragment) {
	fragments.Add(javafragment.NewEndBlockFragment(0, 0, math.MaxInt32, 20, "End array initializer", start))
}

func AddEndSingleStatementMethodBody(fragments util.IList[intmod.IFragment], start intmod.IStartBodyFragment) {
	fragments.Add(javafragment.NewEndBodyFragment(0, 1, 1, 8, "End single statement method body", start))
}

func AddEndMethodBody(fragments util.IList[intmod.IFragment], start intmod.IStartBodyFragment) {
	fragments.Add(javafragment.NewEndBodyFragment(0, 1, 1, 8, "End method body", start))
}

func AddEndInstanceInitializerBlock(fragments util.IList[intmod.IFragment], start intmod.IStartBlockFragment) {
	fragments.Add(javafragment.NewEndBlockFragment(0, 1, 1, 8, "End anonymous method body", start))
}

func AddEndTypeBody(fragments util.IList[intmod.IFragment], start intmod.IStartBodyFragment) {
	fragments.Add(javafragment.NewEndBodyFragment(0, 1, 1, 3, "End type body", start))
}

func AddEndSubTypeBodyInParameter(fragments util.IList[intmod.IFragment], start intmod.IStartBodyFragment) {
	fragments.Add(javafragment.NewEndBodyInParameterFragment(0, 1, 1, 10, "End sub type body in parameter", start))
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 0, math.MaxInt32, 13, "End sub type body spacer in parameter"))
}

func AddEndSubTypeBody(fragments util.IList[intmod.IFragment], start intmod.IStartBodyFragment) {
	fragments.Add(javafragment.NewEndBodyFragment(0, 1, 1, 10, "End sub type body", start))
}

func AddEndSingleStatementBlock(fragments util.IList[intmod.IFragment], start intmod.IStartSingleStatementBlockFragment) {
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 19, "End single statement block spacer"))
	//fragments.Add(javafragment.NewEndSingleStatementBlockFragment(0, 1, 2, 15, "End single statement block", start));
	fragments.Add(javafragment.NewEndSingleStatementBlockFragment(0, 0, 1, 6, "End single statement block", start))
}

func AddEndStatementsBlock(fragments util.IList[intmod.IFragment], group intmod.IStartStatementsBlockFragmentGroup) {
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 19, "End statement block spacer"))
	fragments.Add(javafragment.NewEndStatementsBlockFragment(0, 1, 2, 15, "End statement block", group))
}

func AddSpacerAfterEndStatementsBlock(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 11, "Spacer after end statement block"))
}

func AddEndStatementsInLambdaBlockInParameter(fragments util.IList[intmod.IFragment], start intmod.IStartBlockFragment) {
	fragments.Add(javafragment.NewEndBlockInParameterFragment(0, 1, 2, 15, "End statements in lambda block spacer in parameter", start))
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 0, math.MaxInt32, 15, "End statements in lambda block spacer in parameter"))
}

func AddEndStatementsInLambdaBlock(fragments util.IList[intmod.IFragment], start intmod.IStartBlockFragment) {
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 15, "End statements in lambda block spacer"))
	fragments.Add(javafragment.NewEndBlockFragment(0, 1, 2, 15, "End statements in lambda block spacer", start))
}

func AddSpacerAfterMemberAnnotations(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 1, 1, 10, "Spacer after member annotations"))
}

func AddSpacerAfterSwitchLabel(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.StartDeclarationOrStatementBlock)
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 1, 1, 16, "Spacer after switch label"))
}

func AddSpacerBetweenSwitchLabels(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 1, 1, 16, "Spacer between switch label"))
}

func AddSpacerBeforeExtends(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 0, 1, 2, "Spacer before extends"))
}

func AddSpacerBeforeImplements(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 0, 1, 2, "Spacer before implements"))
}

func AddSpacerBetweenEnumValues(fragments util.IList[intmod.IFragment], preferredLineCount int) {
	fragments.Add(javafragment.Comma)
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, preferredLineCount, 1, 10, "Spacer between enum values"))
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 24, "Second spacer between enum values"))
}

func AddSpacerBetweenFieldDeclarators(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.Comma)
	fragments.Add(javafragment.NewSpacerFragment(0, 0, 1, 10, "Spacer between field declarators"))
}

func AddSpacerBetweenMemberAnnotations(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 1, 1, 10, "Spacer between member annotations"))
}

func AddSpacerBetweenMembers(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpacerBetweenMembersFragment(0, 2, math.MaxInt32, 7, "Spacer between members"))
}

func AddSpacerBetweenStatements(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 1, math.MaxInt32, 12, "Spacer between statements"))
}

func AddSpacerBetweenSwitchLabelBlock(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpacerFragment(0, 1, 1, 17, "Spacer between switch label block"))
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 11, "Spacer between switch label block 2"))
}

func AddSpacerAfterSwitchBlock(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.EndDeclarationOrStatementBlock)
}

func AddStartArrayInitializerBlock(fragments util.IList[intmod.IFragment]) intmod.IStartBlockFragment {
	fragment := javafragment.NewStartBlockFragment(0, 0, math.MaxInt32, 20, "Start array initializer block")
	fragments.Add(fragment)
	return fragment
}

func AddSpacerBetweenArrayInitializerBlock(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.Comma)
	fragments.Add(javafragment.NewSpaceSpacerFragment(0, 0, math.MaxInt32, 20, "Spacer between array initializer block"))
}

func AddNewLineBetweenArrayInitializerBlock(fragments util.IList[intmod.IFragment]) {
	fragments.Add(javafragment.NewSpacerFragment(0, 1, 1, 22, "fragment.Newline between array initializer block"))
}

func AddStartSingleStatementMethodBody(fragments util.IList[intmod.IFragment]) intmod.IStartBodyFragment {
	fragment := javafragment.NewStartBodyFragment(0, 1, 2, 9, "Start single statement method body")
	fragments.Add(fragment)
	return fragment
}

func AddStartMethodBody(fragments util.IList[intmod.IFragment]) intmod.IStartBodyFragment {
	fragment := javafragment.NewStartBodyFragment(0, 1, 2, 9, "Start method body")
	fragments.Add(fragment)
	return fragment
}

func AddStartInstanceInitializerBlock(fragments util.IList[intmod.IFragment]) intmod.IStartBlockFragment {
	fragment := javafragment.NewStartBlockFragment(0, 1, 2, 9, "Start anonymous method body")
	fragments.Add(fragment)
	return fragment
}

func AddStartTypeBody(fragments util.IList[intmod.IFragment]) intmod.IStartBodyFragment {
	fragment := javafragment.NewStartBodyFragment(0, 1, 2, 4, "Start type body")
	fragments.Add(fragment)
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 13, "Start type body spacer"))
	return fragment
}

func AddStartSingleStatementBlock(fragments util.IList[intmod.IFragment]) intmod.IStartSingleStatementBlockFragment {
	fragment := javafragment.NewStartSingleStatementBlockFragment(0, 1, 2, 18, "Start single statement block")
	fragments.Add(fragment)
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 23, "Start single statement block spacer"))
	return fragment
}

func AddStartStatementsBlock(fragments util.IList[intmod.IFragment]) intmod.IStartStatementsBlockFragmentGroup {
	fragment := javafragment.NewStartStatementsBlockFragment(0, 1, 2, 14, "Start statements block")
	fragments.Add(fragment)
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 19, "Start statements block spacer"))
	return fragment.Group()
}

func AddStartStatementsInLambdaBlock(fragments util.IList[intmod.IFragment]) intmod.IStartBlockFragment {
	fragment := javafragment.NewStartBlockFragment(0, 1, 2, 14, "Start statements in lambda block")
	fragments.Add(fragment)
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 14, "Start statements in lambda block spacer"))
	return fragment
}

func AddStartStatementsDoWhileBlock(fragments util.IList[intmod.IFragment]) intmod.IStartStatementsBlockFragmentGroup {
	fragment := javafragment.NewStartStatementsDoWhileBlockFragment(0, 1, 2, 14, "Start statements do-while block")
	fragments.Add(fragment)
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 14, "Start statements do-while block spacer"))
	return fragment.Group()
}

func AddStartStatementsTryBlock(fragments util.IList[intmod.IFragment]) intmod.IStartStatementsBlockFragmentGroup {
	fragment := javafragment.NewStartStatementsTryBlockFragment(0, 1, 2, 14, "Start statements try block")
	fragments.Add(fragment)
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 14, "Start statements try block spacer"))
	return fragment.Group()
}

func AddStartStatementsBlock2(fragments util.IList[intmod.IFragment], group intmod.IStartStatementsBlockFragmentGroup) {
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 23, "Start statements block pre spacer"))
	fragments.Add(javafragment.NewStartStatementsBlockFragmentWithGroup(0, 1, 2, 14, "Start statements block", group))
	fragments.Add(javafragment.NewSpacerFragment(0, 0, math.MaxInt32, 19, "Start statements block post spacer"))
}

func newImportsFragment() intmod.IImportsFragment {
	return javafragment.NewImportsFragment(0)
}
