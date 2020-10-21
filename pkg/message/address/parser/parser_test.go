package parser

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func TestParserValid(t *testing.T) {
	tests := []string{
		`user@example.com`,
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			assert.Empty(t, parseAddress(input))
		})
	}
}

func TestParserBad(t *testing.T) {
	tests := []string{
		`this sucks`,
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			assert.NotEmpty(t, parseAddress(input))
		})
	}
}

func parseAddress(address string) string {
	stream := antlr.NewInputStream(address)
	lexer := NewAddressLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	l := &errorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}

	p := NewAddressParser(tokens)
	p.AddErrorListener(l)

	antlr.ParseTreeWalkerDefault.Walk(&BaseAddressParserListener{}, p.Address())

	return l.msg
}

type errorListener struct {
	*antlr.DefaultErrorListener
	msg string
}

func (l *errorListener) SyntaxError(
	_ antlr.Recognizer,
	_ interface{},
	_, _ int,
	msg string,
	_ antlr.RecognitionException,
) {
	l.msg = msg
}
