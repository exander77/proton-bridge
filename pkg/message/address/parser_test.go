package address

import (
	"net/mail"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		addrs []*mail.Address
	}{
		{
			input: `user@example.com`,
			addrs: []*mail.Address{{
				Address: `user@example.com`,
			}},
		},
		{
			input: `John Doe <jdoe@machine.example>`,
			addrs: []*mail.Address{{
				Name:    `John Doe`,
				Address: `jdoe@machine.example`,
			}},
		},
		{
			input: `Mary Smith <mary@example.net>`,
			addrs: []*mail.Address{{
				Name:    `Mary Smith`,
				Address: `mary@example.net`,
			}},
		},
		{
			input: `"Joe Q. Public" <john.q.public@example.com>`,
			addrs: []*mail.Address{{
				Name:    `Joe Q. Public`,
				Address: `john.q.public@example.com`,
			}},
		},
		{
			input: `Who? <one@y.test>`,
			addrs: []*mail.Address{{
				Name:    `Who?`,
				Address: `one@y.test`,
			}},
		},
		{
			input: `<boss@nil.test>`,
			addrs: []*mail.Address{{
				Address: `boss@nil.test`,
			}},
		},
		{
			input: `"Giant; \"Big\" Box" <sysservices@example.net>`,
			addrs: []*mail.Address{{
				Name:    `Giant; "Big" Box`,
				Address: `sysservices@example.net`,
			}},
		},
		{
			input: `Pete <pete@silly.example>`,
			addrs: []*mail.Address{{
				Name:    `Pete`,
				Address: `pete@silly.example`,
			}},
		},
		{
			input: `"Mary Smith: Personal Account" <smith@home.example>`,
			addrs: []*mail.Address{{
				Name:    `Mary Smith: Personal Account`,
				Address: `smith@home.example`,
			}},
		},
		{
			input: `Gogh Fir <gf@example.com>`,
			addrs: []*mail.Address{{
				Name:    `Gogh Fir`,
				Address: `gf@example.com`,
			}},
		},
		{
			input: `Pete(A nice \) chap) <pete(his account)@silly.test(his host)>`,
			addrs: []*mail.Address{{
				Name:    `Pete`,
				Address: `pete@silly.test`,
			}},
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
			_, err := Parse(test.input)
			assert.NoError(t, err)
		})
	}
}
