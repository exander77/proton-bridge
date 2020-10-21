package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input      string
		name, addr string
	}{
		{
			input: `user@example.com`,
			addr:  `user@example.com`,
		},
		{
			input: `John Doe <jdoe@machine.example>`,
			name:  `John Doe`,
			addr:  `jdoe@machine.example`,
		},
		{
			input: `Mary Smith <mary@example.net>`,
			name:  `Mary Smith`,
			addr:  `mary@example.net`,
		},
		{
			input: `"Joe Q. Public" <john.q.public@example.com>`,
			name:  `Joe Q. Public`,
			addr:  `john.q.public@example.com`,
		},
		{
			input: `Who? <one@y.test>`,
			name:  `Who?`,
			addr:  `one@y.test`,
		},
		{
			input: `<boss@nil.test>`,
			addr:  `boss@nil.test`,
		},
		{
			input: `"Giant; \"Big\" Box" <sysservices@example.net>`,
			name:  `Giant; "Big" Box`,
			addr:  `sysservices@example.net`,
		},
		{
			input: `Pete <pete@silly.example>`,
			name:  `Pete`,
			addr:  `pete@silly.example`,
		},
		{
			input: `"Mary Smith: Personal Account" <smith@home.example>`,
			name:  `Mary Smith: Personal Account`,
			addr:  `smith@home.example`,
		},
		{
			input: `Gogh Fir <gf@example.com>`,
			name:  `Gogh Fir`,
			addr:  `gf@example.com`,
		},
		{
			input: `Pete(A nice \) chap) <pete(his account)@silly.test(his host)>`,
			name:  `Pete`,
			addr:  `pete@silly.test`,
		},
		/*
			{
				input: `(Empty list)(start)Hidden recipients  :(nobody(that I know))  ;`,
			},
			{
				input: `A Group:Ed Jones <c@a.test>,joe@where.test,John <jdoe@one.test>;`,
			},
			{
				input: `Undisclosed recipients:;`,
			},
		*/
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			name, addr, err := Parse(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.name, name)
			assert.Equal(t, test.addr, addr)
		})
	}
}
