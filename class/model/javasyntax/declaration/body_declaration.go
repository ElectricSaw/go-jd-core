package declaration

import intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewBodyDeclaration(internalTypeName string, memberDeclaration intsyn.IMemberDeclaration) intsyn.IBodyDeclaration {
	return &BodyDeclaration{
		internalTypeName:   internalTypeName,
		memberDeclarations: memberDeclaration,
	}
}

type BodyDeclaration struct {
	internalTypeName   string
	memberDeclarations intsyn.IMemberDeclaration
}

func (d *BodyDeclaration) InternalTypeName() string {
	return d.internalTypeName
}

func (d *BodyDeclaration) MemberDeclarations() intsyn.IMemberDeclaration {
	return d.memberDeclarations
}

func (d *BodyDeclaration) SetMemberDeclarations(memberDeclaration intsyn.IMemberDeclaration) {
	d.memberDeclarations = memberDeclaration
}

func (d *BodyDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}
