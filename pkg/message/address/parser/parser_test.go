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
		`John Doe <jdoe@machine.example>`,
		`Mary Smith <mary@example.net>`,
		`"Joe Q. Public" <john.q.public@example.com>`,
		`Mary Smith <mary@x.test>`,
		`jdoe@example.org`,
		`Who? <one@y.test>`,
		`<boss@nil.test>`,
		`"Giant; \"Big\" Box" <sysservices@example.net>`,
		`Pete <pete@silly.example>`,
		`A Group:Ed Jones <c@a.test>,joe@where.test,John <jdoe@one.test>;`,
		`Undisclosed recipients:;`,
		`"Mary Smith: Personal Account" <smith@home.example>`,
		`Pete(A nice \) chap) <pete(his account)@silly.test(his host)>`,
		`(Empty list)(start)Hidden recipients  :(nobody(that I know))  ;`,
		`Gogh Fir <gf@example.com>`,
	}
	for _, input := range tests {
		input := input

		t.Run(input, func(t *testing.T) {
			assert.Empty(t, parseAddress(input))
		})
	}
}

func TestParserBad(t *testing.T) {
	tests := []string{
		`this address sucks`,
	}

	for _, input := range tests {
		input := input

		t.Run(input, func(t *testing.T) {
			assert.NotEmpty(t, parseAddress(input))
		})
	}
}

func TestIsEmailValidCategory(t *testing.T) {
	f, err := os.Open("tests.xml")
	require.NoError(t, err)
	defer func() { require.NoError(t, err) }()

	for test := range readTestCases(f) {
		test := test

		if test.category != "ISEMAIL_VALID_CATEGORY" {
			continue
		}

		t.Run(test.id, func(t *testing.T) {
			assert.Empty(t, parseAddress(test.address))
		})
	}
}

type testCase struct {
	id        string
	address   string
	category  string
	diagnosis string
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
				if t.Name.Local == "test" {
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
					test.category = data

				case "diagnosis":
					test.diagnosis = data
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
