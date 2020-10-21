package address

import (
	"net/mail"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSingleAddress(t *testing.T) {
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
			input: ` normal name  <username@server.com>`,
			addrs: []*mail.Address{{
				Name:    `normal name`,
				Address: `username@server.com`,
			}},
		},
		{
			input: ` "comma, name"  <username@server.com>`,
			addrs: []*mail.Address{{
				Name:    `comma, name`,
				Address: `username@server.com`,
			}},
		},
		{
			input: ` name  <username@server.com> (ignore comment)`,
			addrs: []*mail.Address{{
				Name:    `name`,
				Address: `username@server.com`,
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
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			addrs, err := Parse(test.input)
			assert.NoError(t, err)
			assert.ElementsMatch(t, test.addrs, addrs)
		})
	}
}

func TestParseAddressList(t *testing.T) {
	tests := []struct {
		input string
		addrs []*mail.Address
	}{
		{
			input: `Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>`,
			addrs: []*mail.Address{
				{
					Name:    `Alice`,
					Address: `alice@example.com`,
				},
				{
					Name:    `Bob`,
					Address: `bob@example.com`,
				},
				{
					Name:    `Eve`,
					Address: `eve@example.com`,
				},
			},
		},
		{
			input: `Ed Jones <c@a.test>,joe@where.test,John <jdoe@one.test>`,
			addrs: []*mail.Address{
				{
					Name:    `Ed Jones`,
					Address: `c@a.test`,
				},
				{
					Address: `joe@where.test`,
				},
				{
					Name:    `John`,
					Address: `jdoe@one.test`,
				},
			},
		},
		{
			input: ` name (ignore comment)  <username@server.com>,  (Comment as name) username2@server.com`,
			addrs: []*mail.Address{
				{
					Name:    `name`,
					Address: `username@server.com`,
				},
				{
					Address: `username2@server.com`,
				},
			},
		},
		{
			input: ` "normal name"  <username@server.com>, "comma, name" <address@server.com>`,
			addrs: []*mail.Address{
				{
					Name:    `normal name`,
					Address: `username@server.com`,
				},
				{
					Name:    `comma, name`,
					Address: `address@server.com`,
				},
			},
		},
		{
			input: ` "comma, one"  <username@server.com>, "comma, two" <address@server.com>`,
			addrs: []*mail.Address{
				{
					Name:    `comma, one`,
					Address: `username@server.com`,
				},
				{
					Name:    `comma, two`,
					Address: `address@server.com`,
				},
			},
		},
		/*
			{
				input: ` "comma, name"  <username@server.com>, another, name <address@server.com>`,
			},
			{
				input: ` normal name  <username@server.com>, (comment)All.(around)address@(the)server.com`,
			},
			{
				input: ` normal name  <username@server.com>, All.("comma, in comment")address@(the)server.com`,
			},
		*/
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			addrs, err := Parse(test.input)
			assert.NoError(t, err)
			assert.ElementsMatch(t, test.addrs, addrs)
		})
	}
}
