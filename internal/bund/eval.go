package bund

import (
	// "github.com/pieterclaerhout/go-log"

	"github.com/vulogov/Bund/internal/conf"
	"github.com/vulogov/Bund/internal/parse"
)

func Eval() {
	Init()
	if *conf.LexerPrint {
		parse.LexerPrint(*conf.Expr)
		return
	}
	if *conf.ParserPrint {
		parse.ParserPrint(*conf.Expr)
		return
	}
}
