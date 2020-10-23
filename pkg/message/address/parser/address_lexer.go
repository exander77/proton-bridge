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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 48, 187,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41,
	3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3,
	46, 3, 47, 3, 47, 2, 2, 48, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18,
	35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27,
	53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36,
	71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45,
	89, 46, 91, 47, 93, 48, 3, 2, 5, 3, 2, 50, 59, 3, 2, 67, 92, 3, 2, 99,
	124, 2, 186, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2,
	63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2,
	2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2,
	2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2,
	2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3,
	2, 2, 2, 3, 95, 3, 2, 2, 2, 5, 97, 3, 2, 2, 2, 7, 99, 3, 2, 2, 2, 9, 101,
	3, 2, 2, 2, 11, 103, 3, 2, 2, 2, 13, 105, 3, 2, 2, 2, 15, 107, 3, 2, 2,
	2, 17, 109, 3, 2, 2, 2, 19, 111, 3, 2, 2, 2, 21, 113, 3, 2, 2, 2, 23, 115,
	3, 2, 2, 2, 25, 117, 3, 2, 2, 2, 27, 119, 3, 2, 2, 2, 29, 121, 3, 2, 2,
	2, 31, 123, 3, 2, 2, 2, 33, 125, 3, 2, 2, 2, 35, 127, 3, 2, 2, 2, 37, 129,
	3, 2, 2, 2, 39, 131, 3, 2, 2, 2, 41, 133, 3, 2, 2, 2, 43, 135, 3, 2, 2,
	2, 45, 137, 3, 2, 2, 2, 47, 139, 3, 2, 2, 2, 49, 141, 3, 2, 2, 2, 51, 143,
	3, 2, 2, 2, 53, 145, 3, 2, 2, 2, 55, 147, 3, 2, 2, 2, 57, 149, 3, 2, 2,
	2, 59, 151, 3, 2, 2, 2, 61, 153, 3, 2, 2, 2, 63, 155, 3, 2, 2, 2, 65, 157,
	3, 2, 2, 2, 67, 159, 3, 2, 2, 2, 69, 161, 3, 2, 2, 2, 71, 163, 3, 2, 2,
	2, 73, 165, 3, 2, 2, 2, 75, 167, 3, 2, 2, 2, 77, 169, 3, 2, 2, 2, 79, 171,
	3, 2, 2, 2, 81, 173, 3, 2, 2, 2, 83, 175, 3, 2, 2, 2, 85, 177, 3, 2, 2,
	2, 87, 179, 3, 2, 2, 2, 89, 181, 3, 2, 2, 2, 91, 183, 3, 2, 2, 2, 93, 185,
	3, 2, 2, 2, 95, 96, 7, 2, 2, 2, 96, 4, 3, 2, 2, 2, 97, 98, 4, 3, 10, 2,
	98, 6, 3, 2, 2, 2, 99, 100, 7, 11, 2, 2, 100, 8, 3, 2, 2, 2, 101, 102,
	7, 12, 2, 2, 102, 10, 3, 2, 2, 2, 103, 104, 7, 13, 2, 2, 104, 12, 3, 2,
	2, 2, 105, 106, 7, 14, 2, 2, 106, 14, 3, 2, 2, 2, 107, 108, 7, 15, 2, 2,
	108, 16, 3, 2, 2, 2, 109, 110, 4, 16, 33, 2, 110, 18, 3, 2, 2, 2, 111,
	112, 7, 34, 2, 2, 112, 20, 3, 2, 2, 2, 113, 114, 7, 35, 2, 2, 114, 22,
	3, 2, 2, 2, 115, 116, 7, 36, 2, 2, 116, 24, 3, 2, 2, 2, 117, 118, 7, 37,
	2, 2, 118, 26, 3, 2, 2, 2, 119, 120, 7, 38, 2, 2, 120, 28, 3, 2, 2, 2,
	121, 122, 7, 39, 2, 2, 122, 30, 3, 2, 2, 2, 123, 124, 7, 40, 2, 2, 124,
	32, 3, 2, 2, 2, 125, 126, 7, 41, 2, 2, 126, 34, 3, 2, 2, 2, 127, 128, 7,
	42, 2, 2, 128, 36, 3, 2, 2, 2, 129, 130, 7, 43, 2, 2, 130, 38, 3, 2, 2,
	2, 131, 132, 7, 44, 2, 2, 132, 40, 3, 2, 2, 2, 133, 134, 7, 45, 2, 2, 134,
	42, 3, 2, 2, 2, 135, 136, 7, 46, 2, 2, 136, 44, 3, 2, 2, 2, 137, 138, 7,
	47, 2, 2, 138, 46, 3, 2, 2, 2, 139, 140, 7, 48, 2, 2, 140, 48, 3, 2, 2,
	2, 141, 142, 7, 49, 2, 2, 142, 50, 3, 2, 2, 2, 143, 144, 9, 2, 2, 2, 144,
	52, 3, 2, 2, 2, 145, 146, 7, 60, 2, 2, 146, 54, 3, 2, 2, 2, 147, 148, 7,
	61, 2, 2, 148, 56, 3, 2, 2, 2, 149, 150, 7, 62, 2, 2, 150, 58, 3, 2, 2,
	2, 151, 152, 7, 63, 2, 2, 152, 60, 3, 2, 2, 2, 153, 154, 7, 64, 2, 2, 154,
	62, 3, 2, 2, 2, 155, 156, 7, 65, 2, 2, 156, 64, 3, 2, 2, 2, 157, 158, 7,
	66, 2, 2, 158, 66, 3, 2, 2, 2, 159, 160, 9, 3, 2, 2, 160, 68, 3, 2, 2,
	2, 161, 162, 7, 93, 2, 2, 162, 70, 3, 2, 2, 2, 163, 164, 7, 94, 2, 2, 164,
	72, 3, 2, 2, 2, 165, 166, 7, 95, 2, 2, 166, 74, 3, 2, 2, 2, 167, 168, 7,
	96, 2, 2, 168, 76, 3, 2, 2, 2, 169, 170, 7, 97, 2, 2, 170, 78, 3, 2, 2,
	2, 171, 172, 7, 98, 2, 2, 172, 80, 3, 2, 2, 2, 173, 174, 9, 4, 2, 2, 174,
	82, 3, 2, 2, 2, 175, 176, 7, 125, 2, 2, 176, 84, 3, 2, 2, 2, 177, 178,
	7, 126, 2, 2, 178, 86, 3, 2, 2, 2, 179, 180, 7, 127, 2, 2, 180, 88, 3,
	2, 2, 2, 181, 182, 7, 128, 2, 2, 182, 90, 3, 2, 2, 2, 183, 184, 7, 129,
	2, 2, 184, 92, 3, 2, 2, 2, 185, 186, 4, 130, 1, 2, 186, 94, 3, 2, 2, 2,
	3, 2, 2,
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
	"", "'\u0000'", "", "'\t'", "'\n'", "'\u000B'", "'\u000C'", "'\r'", "",
	"' '", "'!'", "'\"'", "'#'", "'$'", "'%'", "'&'", "'''", "'('", "')'",
	"'*'", "'+'", "','", "'-'", "'.'", "'/'", "", "':'", "';'", "'<'", "'='",
	"'>'", "'?'", "'@'", "", "'['", "'\\'", "']'", "'^'", "'_'", "'`'", "",
	"'{'", "'|'", "'}'", "'~'", "'\u007F'",
}

var lexerSymbolicNames = []string{
	"", "U_00", "U_01_08", "TAB", "LF", "U_0B", "U_0C", "CR", "U_0E_1F", "SP",
	"Exclamation", "DQuote", "Hash", "Dollar", "Percent", "Ampersand", "SQuote",
	"LParens", "RParens", "Asterisk", "Plus", "Comma", "Minus", "Period", "Slash",
	"Digit", "Colon", "Semicolon", "Less", "Equal", "Greater", "Question",
	"At", "AlphaUpper", "LBracket", "Backslash", "RBracket", "Caret", "Underscore",
	"Backtick", "AlphaLower", "LCurly", "Pipe", "RCurly", "Tilde", "Delete",
	"UTF8NonAscii",
}

var lexerRuleNames = []string{
	"U_00", "U_01_08", "TAB", "LF", "U_0B", "U_0C", "CR", "U_0E_1F", "SP",
	"Exclamation", "DQuote", "Hash", "Dollar", "Percent", "Ampersand", "SQuote",
	"LParens", "RParens", "Asterisk", "Plus", "Comma", "Minus", "Period", "Slash",
	"Digit", "Colon", "Semicolon", "Less", "Equal", "Greater", "Question",
	"At", "AlphaUpper", "LBracket", "Backslash", "RBracket", "Caret", "Underscore",
	"Backtick", "AlphaLower", "LCurly", "Pipe", "RCurly", "Tilde", "Delete",
	"UTF8NonAscii",
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
	AddressLexerU_00         = 1
	AddressLexerU_01_08      = 2
	AddressLexerTAB          = 3
	AddressLexerLF           = 4
	AddressLexerU_0B         = 5
	AddressLexerU_0C         = 6
	AddressLexerCR           = 7
	AddressLexerU_0E_1F      = 8
	AddressLexerSP           = 9
	AddressLexerExclamation  = 10
	AddressLexerDQuote       = 11
	AddressLexerHash         = 12
	AddressLexerDollar       = 13
	AddressLexerPercent      = 14
	AddressLexerAmpersand    = 15
	AddressLexerSQuote       = 16
	AddressLexerLParens      = 17
	AddressLexerRParens      = 18
	AddressLexerAsterisk     = 19
	AddressLexerPlus         = 20
	AddressLexerComma        = 21
	AddressLexerMinus        = 22
	AddressLexerPeriod       = 23
	AddressLexerSlash        = 24
	AddressLexerDigit        = 25
	AddressLexerColon        = 26
	AddressLexerSemicolon    = 27
	AddressLexerLess         = 28
	AddressLexerEqual        = 29
	AddressLexerGreater      = 30
	AddressLexerQuestion     = 31
	AddressLexerAt           = 32
	AddressLexerAlphaUpper   = 33
	AddressLexerLBracket     = 34
	AddressLexerBackslash    = 35
	AddressLexerRBracket     = 36
	AddressLexerCaret        = 37
	AddressLexerUnderscore   = 38
	AddressLexerBacktick     = 39
	AddressLexerAlphaLower   = 40
	AddressLexerLCurly       = 41
	AddressLexerPipe         = 42
	AddressLexerRCurly       = 43
	AddressLexerTilde        = 44
	AddressLexerDelete       = 45
	AddressLexerUTF8NonAscii = 46
)
