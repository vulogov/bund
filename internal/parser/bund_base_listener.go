// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBundListener is a complete listener for a parse tree produced by BundParser.
type BaseBundListener struct{}

var _ BundListener = &BaseBundListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBundListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBundListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBundListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBundListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseBundListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseBundListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseBundListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseBundListener) ExitTerm(ctx *TermContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBundListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBundListener) ExitBlock(ctx *BlockContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseBundListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseBundListener) ExitPair(ctx *PairContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseBundListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseBundListener) ExitDirective(ctx *DirectiveContext) {}

// EnterCmd is called when production cmd is entered.
func (s *BaseBundListener) EnterCmd(ctx *CmdContext) {}

// ExitCmd is called when production cmd is exited.
func (s *BaseBundListener) ExitCmd(ctx *CmdContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseBundListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseBundListener) ExitLambda(ctx *LambdaContext) {}
