package token

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

type AbstractNopTokenVisitor struct {
}

func (v *AbstractNopTokenVisitor) VisitBooleanConstantToken(_ intmod.IBooleanConstantToken) {}

func (v *AbstractNopTokenVisitor) VisitCharacterConstantToken(_ intmod.ICharacterConstantToken) {}

func (v *AbstractNopTokenVisitor) VisitDeclarationToken(_ intmod.IDeclarationToken) {}

func (v *AbstractNopTokenVisitor) VisitEndBlockToken(_ intmod.IEndBlockToken) {}

func (v *AbstractNopTokenVisitor) VisitEndMarkerToken(_ intmod.IEndMarkerToken) {}

func (v *AbstractNopTokenVisitor) VisitKeywordToken(_ intmod.IKeywordToken) {}

func (v *AbstractNopTokenVisitor) VisitLineNumberToken(_ intmod.ILineNumberToken) {}

func (v *AbstractNopTokenVisitor) VisitNewLineToken(_ intmod.INewLineToken) {}

func (v *AbstractNopTokenVisitor) VisitNumericConstantToken(_ intmod.INumericConstantToken) {}

func (v *AbstractNopTokenVisitor) VisitReferenceToken(_ intmod.IReferenceToken) {}

func (v *AbstractNopTokenVisitor) VisitStartBlockToken(_ intmod.IStartBlockToken) {}

func (v *AbstractNopTokenVisitor) VisitStartMarkerToken(_ intmod.IStartMarkerToken) {}

func (v *AbstractNopTokenVisitor) VisitStringConstantToken(_ intmod.IStringConstantToken) {}

func (v *AbstractNopTokenVisitor) VisitTextToken(_ intmod.ITextToken) {}
