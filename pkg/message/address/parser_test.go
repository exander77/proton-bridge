package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input              string
		wantName, wantAddr string
	}{
		{
			input:    `user@example.com`,
			wantAddr: `user@example.com`,
		},
		{
			input:    `John Doe <jdoe@machine.example>`,
			wantName: `John Doe`,
			wantAddr: `jdoe@machine.example`,
		},
		/*
			{
				input: `Mary Smith <mary@example.net>`,
			},
			{
				input: `"Joe Q. Public" <john.q.public@example.com>`,
			},
			{
				input: `Mary Smith <mary@x.test>`,
			},
			{
				input: `jdoe@example.org`,
			},
			{
				input: `Who? <one@y.test>`,
			},
			{
				input: `<boss@nil.test>`,
			},
			{
				input: `"Giant; \"Big\" Box" <sysservices@example.net>`,
			},
			{
				input: `Pete <pete@silly.example>`,
			},
			{
				input: `A Group:Ed Jones <c@a.test>,joe@where.test,John <jdoe@one.test>;`,
			},
			{
				input: `Undisclosed recipients:;`,
			},
			{
				input: `"Mary Smith: Personal Account" <smith@home.example>`,
			},
			{
				input: `Pete(A nice \) chap) <pete(his account)@silly.test(his host)>`,
			},
			{
				input: `(Empty list)(start)Hidden recipients  :(nobody(that I know))  ;`,
			},
			{
				input: `Gogh Fir <gf@example.com>`,
			},
		*/
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			name, addr, err := Parse(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.wantName, name)
			assert.Equal(t, test.wantAddr, addr)
		})
	}
}
