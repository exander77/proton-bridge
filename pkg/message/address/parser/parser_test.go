package parser

import (
	"encoding/xml"
	"io"
	"os"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParserValid(t *testing.T) {
	tests := []string{
		`user@example.com`,
		`Gogh Fir <gf@example.com>`,
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

func _TestIsEmailStandardTestSet(t *testing.T) {
	f, err := os.Open("tests.xml")
	require.NoError(t, err)
	defer func() { require.NoError(t, err) }()

	for test := range readTestCases(f) {
		t.Run(test.id, func(t *testing.T) {
			msg := parseAddress(test.address)

			if test.valid {
				assert.Empty(t, msg)
			} else {
				assert.NotEmpty(t, msg)
			}
		})
	}
}

type testCase struct {
	id      string
	address string
	valid   bool
}

func readTestCases(r io.Reader) chan testCase {
	ch := make(chan testCase)

	var (
		test testCase
		data string
	)

	go func() {
		decoder := xml.NewDecoder(r)

		for token, err := decoder.Token(); err == nil; token, err = decoder.Token() {
			switch t := token.(type) {
			case xml.StartElement:
				switch t.Name.Local {
				case "test":
					test = testCase{
						id: t.Attr[0].Value,
					}
				}

			case xml.EndElement:
				switch t.Name.Local {
				case "test":
					ch <- test

				case "address":
					test.address = data

				case "category":
					test.valid = data != "ISEMAIL_ERR"
				}

			case xml.CharData:
				data = string(t)
			}
		}

		close(ch)
	}()

	return ch
}

func parseAddress(address string) string {
	stream := antlr.NewInputStream(address)
	lexer := NewAddressLexer(stream)

	/*
		for {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}

			fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
		}
	*/

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
