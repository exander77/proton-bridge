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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 41, 264,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3,
	37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42,
	3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53,
	3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3,
	58, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63,
	3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 2, 2, 67, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15,
	29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24,
	47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33,
	65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 2,
	83, 2, 85, 2, 87, 2, 89, 2, 91, 2, 93, 2, 95, 2, 97, 2, 99, 2, 101, 2,
	103, 2, 105, 2, 107, 2, 109, 2, 111, 2, 113, 2, 115, 2, 117, 2, 119, 2,
	121, 2, 123, 2, 125, 2, 127, 2, 129, 2, 131, 2, 3, 2, 31, 3, 2, 50, 59,
	3, 2, 67, 92, 3, 2, 99, 124, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100,
	4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103,
	4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106,
	4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109,
	4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112,
	4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115,
	4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118,
	4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121,
	4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124,
	2, 237, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2,
	2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3,
	2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63,
	3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2,
	71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2,
	2, 79, 3, 2, 2, 2, 3, 133, 3, 2, 2, 2, 5, 136, 3, 2, 2, 2, 7, 138, 3, 2,
	2, 2, 9, 140, 3, 2, 2, 2, 11, 142, 3, 2, 2, 2, 13, 144, 3, 2, 2, 2, 15,
	146, 3, 2, 2, 2, 17, 148, 3, 2, 2, 2, 19, 150, 3, 2, 2, 2, 21, 152, 3,
	2, 2, 2, 23, 154, 3, 2, 2, 2, 25, 156, 3, 2, 2, 2, 27, 158, 3, 2, 2, 2,
	29, 160, 3, 2, 2, 2, 31, 162, 3, 2, 2, 2, 33, 164, 3, 2, 2, 2, 35, 166,
	3, 2, 2, 2, 37, 168, 3, 2, 2, 2, 39, 170, 3, 2, 2, 2, 41, 172, 3, 2, 2,
	2, 43, 174, 3, 2, 2, 2, 45, 176, 3, 2, 2, 2, 47, 178, 3, 2, 2, 2, 49, 180,
	3, 2, 2, 2, 51, 182, 3, 2, 2, 2, 53, 184, 3, 2, 2, 2, 55, 186, 3, 2, 2,
	2, 57, 188, 3, 2, 2, 2, 59, 190, 3, 2, 2, 2, 61, 192, 3, 2, 2, 2, 63, 194,
	3, 2, 2, 2, 65, 196, 3, 2, 2, 2, 67, 198, 3, 2, 2, 2, 69, 200, 3, 2, 2,
	2, 71, 202, 3, 2, 2, 2, 73, 204, 3, 2, 2, 2, 75, 206, 3, 2, 2, 2, 77, 208,
	3, 2, 2, 2, 79, 210, 3, 2, 2, 2, 81, 212, 3, 2, 2, 2, 83, 214, 3, 2, 2,
	2, 85, 216, 3, 2, 2, 2, 87, 218, 3, 2, 2, 2, 89, 220, 3, 2, 2, 2, 91, 222,
	3, 2, 2, 2, 93, 224, 3, 2, 2, 2, 95, 226, 3, 2, 2, 2, 97, 228, 3, 2, 2,
	2, 99, 230, 3, 2, 2, 2, 101, 232, 3, 2, 2, 2, 103, 234, 3, 2, 2, 2, 105,
	236, 3, 2, 2, 2, 107, 238, 3, 2, 2, 2, 109, 240, 3, 2, 2, 2, 111, 242,
	3, 2, 2, 2, 113, 244, 3, 2, 2, 2, 115, 246, 3, 2, 2, 2, 117, 248, 3, 2,
	2, 2, 119, 250, 3, 2, 2, 2, 121, 252, 3, 2, 2, 2, 123, 254, 3, 2, 2, 2,
	125, 256, 3, 2, 2, 2, 127, 258, 3, 2, 2, 2, 129, 260, 3, 2, 2, 2, 131,
	262, 3, 2, 2, 2, 133, 134, 7, 15, 2, 2, 134, 135, 7, 12, 2, 2, 135, 4,
	3, 2, 2, 2, 136, 137, 7, 129, 2, 2, 137, 6, 3, 2, 2, 2, 138, 139, 7, 11,
	2, 2, 139, 8, 3, 2, 2, 2, 140, 141, 7, 34, 2, 2, 141, 10, 3, 2, 2, 2, 142,
	143, 7, 35, 2, 2, 143, 12, 3, 2, 2, 2, 144, 145, 7, 36, 2, 2, 145, 14,
	3, 2, 2, 2, 146, 147, 7, 37, 2, 2, 147, 16, 3, 2, 2, 2, 148, 149, 7, 38,
	2, 2, 149, 18, 3, 2, 2, 2, 150, 151, 7, 39, 2, 2, 151, 20, 3, 2, 2, 2,
	152, 153, 7, 40, 2, 2, 153, 22, 3, 2, 2, 2, 154, 155, 7, 41, 2, 2, 155,
	24, 3, 2, 2, 2, 156, 157, 7, 42, 2, 2, 157, 26, 3, 2, 2, 2, 158, 159, 7,
	43, 2, 2, 159, 28, 3, 2, 2, 2, 160, 161, 7, 44, 2, 2, 161, 30, 3, 2, 2,
	2, 162, 163, 7, 45, 2, 2, 163, 32, 3, 2, 2, 2, 164, 165, 7, 46, 2, 2, 165,
	34, 3, 2, 2, 2, 166, 167, 7, 47, 2, 2, 167, 36, 3, 2, 2, 2, 168, 169, 7,
	48, 2, 2, 169, 38, 3, 2, 2, 2, 170, 171, 7, 49, 2, 2, 171, 40, 3, 2, 2,
	2, 172, 173, 9, 2, 2, 2, 173, 42, 3, 2, 2, 2, 174, 175, 7, 60, 2, 2, 175,
	44, 3, 2, 2, 2, 176, 177, 7, 61, 2, 2, 177, 46, 3, 2, 2, 2, 178, 179, 7,
	62, 2, 2, 179, 48, 3, 2, 2, 2, 180, 181, 7, 63, 2, 2, 181, 50, 3, 2, 2,
	2, 182, 183, 7, 64, 2, 2, 183, 52, 3, 2, 2, 2, 184, 185, 7, 65, 2, 2, 185,
	54, 3, 2, 2, 2, 186, 187, 7, 66, 2, 2, 187, 56, 3, 2, 2, 2, 188, 189, 9,
	3, 2, 2, 189, 58, 3, 2, 2, 2, 190, 191, 7, 93, 2, 2, 191, 60, 3, 2, 2,
	2, 192, 193, 7, 94, 2, 2, 193, 62, 3, 2, 2, 2, 194, 195, 7, 95, 2, 2, 195,
	64, 3, 2, 2, 2, 196, 197, 7, 96, 2, 2, 197, 66, 3, 2, 2, 2, 198, 199, 7,
	97, 2, 2, 199, 68, 3, 2, 2, 2, 200, 201, 7, 98, 2, 2, 201, 70, 3, 2, 2,
	2, 202, 203, 9, 4, 2, 2, 203, 72, 3, 2, 2, 2, 204, 205, 7, 125, 2, 2, 205,
	74, 3, 2, 2, 2, 206, 207, 7, 126, 2, 2, 207, 76, 3, 2, 2, 2, 208, 209,
	7, 127, 2, 2, 209, 78, 3, 2, 2, 2, 210, 211, 7, 128, 2, 2, 211, 80, 3,
	2, 2, 2, 212, 213, 9, 5, 2, 2, 213, 82, 3, 2, 2, 2, 214, 215, 9, 6, 2,
	2, 215, 84, 3, 2, 2, 2, 216, 217, 9, 7, 2, 2, 217, 86, 3, 2, 2, 2, 218,
	219, 9, 8, 2, 2, 219, 88, 3, 2, 2, 2, 220, 221, 9, 9, 2, 2, 221, 90, 3,
	2, 2, 2, 222, 223, 9, 10, 2, 2, 223, 92, 3, 2, 2, 2, 224, 225, 9, 11, 2,
	2, 225, 94, 3, 2, 2, 2, 226, 227, 9, 12, 2, 2, 227, 96, 3, 2, 2, 2, 228,
	229, 9, 13, 2, 2, 229, 98, 3, 2, 2, 2, 230, 231, 9, 14, 2, 2, 231, 100,
	3, 2, 2, 2, 232, 233, 9, 15, 2, 2, 233, 102, 3, 2, 2, 2, 234, 235, 9, 16,
	2, 2, 235, 104, 3, 2, 2, 2, 236, 237, 9, 17, 2, 2, 237, 106, 3, 2, 2, 2,
	238, 239, 9, 18, 2, 2, 239, 108, 3, 2, 2, 2, 240, 241, 9, 19, 2, 2, 241,
	110, 3, 2, 2, 2, 242, 243, 9, 20, 2, 2, 243, 112, 3, 2, 2, 2, 244, 245,
	9, 21, 2, 2, 245, 114, 3, 2, 2, 2, 246, 247, 9, 22, 2, 2, 247, 116, 3,
	2, 2, 2, 248, 249, 9, 23, 2, 2, 249, 118, 3, 2, 2, 2, 250, 251, 9, 24,
	2, 2, 251, 120, 3, 2, 2, 2, 252, 253, 9, 25, 2, 2, 253, 122, 3, 2, 2, 2,
	254, 255, 9, 26, 2, 2, 255, 124, 3, 2, 2, 2, 256, 257, 9, 27, 2, 2, 257,
	126, 3, 2, 2, 2, 258, 259, 9, 28, 2, 2, 259, 128, 3, 2, 2, 2, 260, 261,
	9, 29, 2, 2, 261, 130, 3, 2, 2, 2, 262, 263, 9, 30, 2, 2, 263, 132, 3,
	2, 2, 2, 3, 2, 2,
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
	"", "'\r\n'", "'\u007F'", "'\t'", "' '", "'!'", "'\"'", "'#'", "'$'", "'%'",
	"'&'", "'''", "'('", "')'", "'*'", "'+'", "','", "'-'", "'.'", "'/'", "",
	"':'", "';'", "'<'", "'='", "'>'", "'?'", "'@'", "", "'['", "'\\'", "']'",
	"'^'", "'_'", "'`'", "", "'{'", "'|'", "'}'", "'~'",
}

var lexerSymbolicNames = []string{
	"", "CRLF", "DEL", "HTAB", "SP", "Exclamation", "DQuote", "Hash", "Dollar",
	"Percent", "Ampersand", "SQuote", "LParens", "RParens", "Asterisk", "Plus",
	"Comma", "Minus", "Period", "Slash", "Digit", "Colon", "Semicolon", "Less",
	"Equal", "Greater", "Question", "At", "AlphaUpper", "LBracket", "Backslash",
	"RBracket", "Caret", "Underscore", "Backtick", "AlphaLower", "LCurly",
	"Pipe", "RCurly", "Tilde",
}

var lexerRuleNames = []string{
	"CRLF", "DEL", "HTAB", "SP", "Exclamation", "DQuote", "Hash", "Dollar",
	"Percent", "Ampersand", "SQuote", "LParens", "RParens", "Asterisk", "Plus",
	"Comma", "Minus", "Period", "Slash", "Digit", "Colon", "Semicolon", "Less",
	"Equal", "Greater", "Question", "At", "AlphaUpper", "LBracket", "Backslash",
	"RBracket", "Caret", "Underscore", "Backtick", "AlphaLower", "LCurly",
	"Pipe", "RCurly", "Tilde", "A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
	"Y", "Z",
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
	AddressLexerCRLF        = 1
	AddressLexerDEL         = 2
	AddressLexerHTAB        = 3
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
)
