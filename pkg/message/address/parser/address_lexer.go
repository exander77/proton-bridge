// Code generated from pkg/message/address/parser/AddressLexer.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 42, 163,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 2,
	2, 42, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39,
	21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57,
	30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75,
	39, 77, 40, 79, 41, 81, 42, 3, 2, 5, 3, 2, 50, 59, 3, 2, 67, 92, 3, 2,
	99, 124, 2, 162, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2,
	2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3,
	2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39,
	3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2,
	47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2,
	2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2,
	2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2,
	2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3,
	2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 3, 83, 3, 2, 2, 2, 5, 85,
	3, 2, 2, 2, 7, 87, 3, 2, 2, 2, 9, 89, 3, 2, 2, 2, 11, 91, 3, 2, 2, 2, 13,
	93, 3, 2, 2, 2, 15, 95, 3, 2, 2, 2, 17, 97, 3, 2, 2, 2, 19, 99, 3, 2, 2,
	2, 21, 101, 3, 2, 2, 2, 23, 103, 3, 2, 2, 2, 25, 105, 3, 2, 2, 2, 27, 107,
	3, 2, 2, 2, 29, 109, 3, 2, 2, 2, 31, 111, 3, 2, 2, 2, 33, 113, 3, 2, 2,
	2, 35, 115, 3, 2, 2, 2, 37, 117, 3, 2, 2, 2, 39, 119, 3, 2, 2, 2, 41, 121,
	3, 2, 2, 2, 43, 123, 3, 2, 2, 2, 45, 125, 3, 2, 2, 2, 47, 127, 3, 2, 2,
	2, 49, 129, 3, 2, 2, 2, 51, 131, 3, 2, 2, 2, 53, 133, 3, 2, 2, 2, 55, 135,
	3, 2, 2, 2, 57, 137, 3, 2, 2, 2, 59, 139, 3, 2, 2, 2, 61, 141, 3, 2, 2,
	2, 63, 143, 3, 2, 2, 2, 65, 145, 3, 2, 2, 2, 67, 147, 3, 2, 2, 2, 69, 149,
	3, 2, 2, 2, 71, 151, 3, 2, 2, 2, 73, 153, 3, 2, 2, 2, 75, 155, 3, 2, 2,
	2, 77, 157, 3, 2, 2, 2, 79, 159, 3, 2, 2, 2, 81, 161, 3, 2, 2, 2, 83, 84,
	7, 11, 2, 2, 84, 4, 3, 2, 2, 2, 85, 86, 7, 12, 2, 2, 86, 6, 3, 2, 2, 2,
	87, 88, 7, 15, 2, 2, 88, 8, 3, 2, 2, 2, 89, 90, 7, 34, 2, 2, 90, 10, 3,
	2, 2, 2, 91, 92, 7, 35, 2, 2, 92, 12, 3, 2, 2, 2, 93, 94, 7, 36, 2, 2,
	94, 14, 3, 2, 2, 2, 95, 96, 7, 37, 2, 2, 96, 16, 3, 2, 2, 2, 97, 98, 7,
	38, 2, 2, 98, 18, 3, 2, 2, 2, 99, 100, 7, 39, 2, 2, 100, 20, 3, 2, 2, 2,
	101, 102, 7, 40, 2, 2, 102, 22, 3, 2, 2, 2, 103, 104, 7, 41, 2, 2, 104,
	24, 3, 2, 2, 2, 105, 106, 7, 42, 2, 2, 106, 26, 3, 2, 2, 2, 107, 108, 7,
	43, 2, 2, 108, 28, 3, 2, 2, 2, 109, 110, 7, 44, 2, 2, 110, 30, 3, 2, 2,
	2, 111, 112, 7, 45, 2, 2, 112, 32, 3, 2, 2, 2, 113, 114, 7, 46, 2, 2, 114,
	34, 3, 2, 2, 2, 115, 116, 7, 47, 2, 2, 116, 36, 3, 2, 2, 2, 117, 118, 7,
	48, 2, 2, 118, 38, 3, 2, 2, 2, 119, 120, 7, 49, 2, 2, 120, 40, 3, 2, 2,
	2, 121, 122, 9, 2, 2, 2, 122, 42, 3, 2, 2, 2, 123, 124, 7, 60, 2, 2, 124,
	44, 3, 2, 2, 2, 125, 126, 7, 61, 2, 2, 126, 46, 3, 2, 2, 2, 127, 128, 7,
	62, 2, 2, 128, 48, 3, 2, 2, 2, 129, 130, 7, 63, 2, 2, 130, 50, 3, 2, 2,
	2, 131, 132, 7, 64, 2, 2, 132, 52, 3, 2, 2, 2, 133, 134, 7, 65, 2, 2, 134,
	54, 3, 2, 2, 2, 135, 136, 7, 66, 2, 2, 136, 56, 3, 2, 2, 2, 137, 138, 9,
	3, 2, 2, 138, 58, 3, 2, 2, 2, 139, 140, 7, 93, 2, 2, 140, 60, 3, 2, 2,
	2, 141, 142, 7, 94, 2, 2, 142, 62, 3, 2, 2, 2, 143, 144, 7, 95, 2, 2, 144,
	64, 3, 2, 2, 2, 145, 146, 7, 96, 2, 2, 146, 66, 3, 2, 2, 2, 147, 148, 7,
	97, 2, 2, 148, 68, 3, 2, 2, 2, 149, 150, 7, 98, 2, 2, 150, 70, 3, 2, 2,
	2, 151, 152, 9, 4, 2, 2, 152, 72, 3, 2, 2, 2, 153, 154, 7, 125, 2, 2, 154,
	74, 3, 2, 2, 2, 155, 156, 7, 126, 2, 2, 156, 76, 3, 2, 2, 2, 157, 158,
	7, 127, 2, 2, 158, 78, 3, 2, 2, 2, 159, 160, 7, 128, 2, 2, 160, 80, 3,
	2, 2, 2, 161, 162, 7, 129, 2, 2, 162, 82, 3, 2, 2, 2, 3, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'\t'", "'\n'", "'\r'", "' '", "'!'", "'\"'", "'#'", "'$'", "'%'",
	"'&'", "'''", "'('", "')'", "'*'", "'+'", "','", "'-'", "'.'", "'/'", "",
	"':'", "';'", "'<'", "'='", "'>'", "'?'", "'@'", "", "'['", "'\\'", "']'",
	"'^'", "'_'", "'`'", "", "'{'", "'|'", "'}'", "'~'", "'\u007F'",
}

var lexerSymbolicNames = []string{
	"", "TAB", "LF", "CR", "SP", "Exclamation", "DQuote", "Hash", "Dollar",
	"Percent", "Ampersand", "SQuote", "LParens", "RParens", "Asterisk", "Plus",
	"Comma", "Minus", "Period", "Slash", "Digit", "Colon", "Semicolon", "Less",
	"Equal", "Greater", "Question", "At", "AlphaUpper", "LBracket", "Backslash",
	"RBracket", "Caret", "Underscore", "Backtick", "AlphaLower", "LCurly",
	"Pipe", "RCurly", "Tilde", "DEL",
}

var lexerRuleNames = []string{
	"TAB", "LF", "CR", "SP", "Exclamation", "DQuote", "Hash", "Dollar", "Percent",
	"Ampersand", "SQuote", "LParens", "RParens", "Asterisk", "Plus", "Comma",
	"Minus", "Period", "Slash", "Digit", "Colon", "Semicolon", "Less", "Equal",
	"Greater", "Question", "At", "AlphaUpper", "LBracket", "Backslash", "RBracket",
	"Caret", "Underscore", "Backtick", "AlphaLower", "LCurly", "Pipe", "RCurly",
	"Tilde", "DEL",
}

type AddressLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewAddressLexer(input antlr.CharStream) *AddressLexer {

	l := new(AddressLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "AddressLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AddressLexer tokens.
const (
	AddressLexerTAB         = 1
	AddressLexerLF          = 2
	AddressLexerCR          = 3
	AddressLexerSP          = 4
	AddressLexerExclamation = 5
	AddressLexerDQuote      = 6
	AddressLexerHash        = 7
	AddressLexerDollar      = 8
	AddressLexerPercent     = 9
	AddressLexerAmpersand   = 10
	AddressLexerSQuote      = 11
	AddressLexerLParens     = 12
	AddressLexerRParens     = 13
	AddressLexerAsterisk    = 14
	AddressLexerPlus        = 15
	AddressLexerComma       = 16
	AddressLexerMinus       = 17
	AddressLexerPeriod      = 18
	AddressLexerSlash       = 19
	AddressLexerDigit       = 20
	AddressLexerColon       = 21
	AddressLexerSemicolon   = 22
	AddressLexerLess        = 23
	AddressLexerEqual       = 24
	AddressLexerGreater     = 25
	AddressLexerQuestion    = 26
	AddressLexerAt          = 27
	AddressLexerAlphaUpper  = 28
	AddressLexerLBracket    = 29
	AddressLexerBackslash   = 30
	AddressLexerRBracket    = 31
	AddressLexerCaret       = 32
	AddressLexerUnderscore  = 33
	AddressLexerBacktick    = 34
	AddressLexerAlphaLower  = 35
	AddressLexerLCurly      = 36
	AddressLexerPipe        = 37
	AddressLexerRCurly      = 38
	AddressLexerTilde       = 39
	AddressLexerDEL         = 40
)