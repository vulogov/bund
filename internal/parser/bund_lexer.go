// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 66, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 7, 4, 28, 10, 4, 12, 4, 14, 4, 31, 11, 4, 3, 5, 3, 5, 3, 5, 5, 5,
	36, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 46, 10,
	8, 3, 8, 5, 8, 49, 10, 8, 3, 8, 3, 8, 5, 8, 53, 10, 8, 3, 9, 3, 9, 7, 9,
	57, 10, 9, 12, 9, 14, 9, 60, 11, 9, 3, 10, 6, 10, 63, 10, 10, 13, 10, 14,
	10, 64, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 2, 13, 2, 15, 2, 17, 2, 19,
	2, 3, 2, 6, 3, 2, 51, 59, 3, 2, 50, 59, 4, 2, 12, 12, 14, 15, 4, 2, 11,
	11, 34, 34, 2, 68, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3, 2, 2,
	2, 9, 35, 3, 2, 2, 2, 11, 39, 3, 2, 2, 2, 13, 41, 3, 2, 2, 2, 15, 43, 3,
	2, 2, 2, 17, 54, 3, 2, 2, 2, 19, 62, 3, 2, 2, 2, 21, 22, 5, 5, 3, 2, 22,
	4, 3, 2, 2, 2, 23, 24, 5, 7, 4, 2, 24, 6, 3, 2, 2, 2, 25, 29, 5, 11, 6,
	2, 26, 28, 5, 13, 7, 2, 27, 26, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27,
	3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 8, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2,
	32, 36, 5, 19, 10, 2, 33, 36, 5, 17, 9, 2, 34, 36, 5, 15, 8, 2, 35, 32,
	3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2,
	37, 38, 8, 5, 2, 2, 38, 10, 3, 2, 2, 2, 39, 40, 9, 2, 2, 2, 40, 12, 3,
	2, 2, 2, 41, 42, 9, 3, 2, 2, 42, 14, 3, 2, 2, 2, 43, 45, 7, 94, 2, 2, 44,
	46, 5, 19, 10, 2, 45, 44, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 52, 3, 2,
	2, 2, 47, 49, 7, 15, 2, 2, 48, 47, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49,
	50, 3, 2, 2, 2, 50, 53, 7, 12, 2, 2, 51, 53, 4, 14, 15, 2, 52, 48, 3, 2,
	2, 2, 52, 51, 3, 2, 2, 2, 53, 16, 3, 2, 2, 2, 54, 58, 7, 37, 2, 2, 55,
	57, 10, 4, 2, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2,
	2, 2, 58, 59, 3, 2, 2, 2, 59, 18, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 63,
	9, 5, 2, 2, 62, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2,
	64, 65, 3, 2, 2, 2, 65, 20, 3, 2, 2, 2, 10, 2, 29, 35, 45, 48, 52, 58,
	64, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames []string

var lexerSymbolicNames = []string{
	"", "NUMBER", "INTEGER", "DECIMAL_INTEGER", "SKIP_",
}

var lexerRuleNames = []string{
	"NUMBER", "INTEGER", "DECIMAL_INTEGER", "SKIP_", "NON_ZERO_DIGIT", "DIGIT",
	"LINE_JOINING", "COMMENT", "SPACES",
}

type BundLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewBundLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *BundLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBundLexer(input antlr.CharStream) *BundLexer {
	l := new(BundLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Bund.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BundLexer tokens.
const (
	BundLexerNUMBER          = 1
	BundLexerINTEGER         = 2
	BundLexerDECIMAL_INTEGER = 3
	BundLexerSKIP_           = 4
)
