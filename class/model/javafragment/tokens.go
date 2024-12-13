package javafragment

import (
	"github.com/ElectricSaw/go-jd-core/class/model/token"
)

// -------- StartMovableJavaBlockFragment --------

var StartMovableTypeBlock = NewStartMovableJavaBlockFragment(1)
var StartMovableFieldBlock = NewStartMovableJavaBlockFragment(2)
var StartMovableMethodBlock = NewStartMovableJavaBlockFragment(3)

// -------- TokensFragment --------

var Comma = NewTokensFragment(token.Comma)
var Semicolon = NewTokensFragment(token.Semicolon)
var StartDeclarationOrStatementBlock = NewTokensFragment(token.StartDeclarationOrStatementBlock)
var EndDeclarationOrStatementBlock = NewTokensFragment(token.EndDeclarationOrStatementBlock)
var EndDeclarationOrStatementBlockSemicolon = NewTokensFragment(token.EndDeclarationOrStatementBlock, token.Semicolon)
var ReturnSemicolon = NewTokensFragment(token.Return, token.Semicolon)
