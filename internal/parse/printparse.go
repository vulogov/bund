package parse

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/vulogov/Bund/internal/parser"
)

type bundListener struct {
	*parser.BaseBundListener
}

func ParserPrint(code string) {
	_input := antlr.NewInputStream(code)
	lexer := parser.NewBundLexer(_input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBundParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&bundListener{}, p.Expressions())
}

func (l *bundListener) ExitTerm(c *parser.TermContext) {
	fmt.Printf(">>> %s  \n", c.GetText())
}

func (l *bundListener) ExitPair(c *parser.PairContext) {
	fmt.Printf("pair: %s %+v \n", c.GetText(), c.GetHead())
}

func (l *bundListener) ExitDirective(c *parser.DirectiveContext) {
	fmt.Printf("directive: %s \n", c.GetText())
}

func (l *bundListener) ExitCmd(c *parser.CmdContext) {
	fmt.Printf("command: %s \n", c.GetText())
}
