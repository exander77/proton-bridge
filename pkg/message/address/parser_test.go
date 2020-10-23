// Copyright (c) 2020 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

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
		test := test

		t.Run(test.input, func(t *testing.T) {
			addrs, err := Parse(test.input)
			assert.NoError(t, err)
			assert.ElementsMatch(t, test.addrs, addrs)
		})
	}
}

func TestParseSingleAddressEncodedWord(t *testing.T) {
	tests := []struct {
		input string
		addrs []*mail.Address
	}{
		{
			input: `=?US-ASCII?Q?Keith_Moore?= <moore@cs.utk.edu>`,
			addrs: []*mail.Address{{
				Name:    `Keith Moore`,
				Address: `moore@cs.utk.edu`,
			}},
		},
		{
			input: `=?ISO-8859-1?Q?Keld_J=F8rn_Simonsen?= <keld@dkuug.dk>`,
			addrs: []*mail.Address{{
				Name:    `Keld Jørn Simonsen`,
				Address: `keld@dkuug.dk`,
			}},
		},
		{
			input: `=?ISO-8859-1?Q?Andr=E9?= Pirard <PIRARD@vm1.ulg.ac.be>`,
			addrs: []*mail.Address{{
				Name:    `André Pirard`,
				Address: `PIRARD@vm1.ulg.ac.be`,
			}},
		},
		{
			input: `=?ISO-8859-1?Q?Olle_J=E4rnefors?= <ojarnef@admin.kth.se>`,
			addrs: []*mail.Address{{
				Name:    `Olle Järnefors`,
				Address: `ojarnef@admin.kth.se`,
			}},
		},
		{
			input: `=?ISO-8859-1?Q?Patrik_F=E4ltstr=F6m?= <paf@nada.kth.se>`,
			addrs: []*mail.Address{{
				Name:    `Patrik Fältström`,
				Address: `paf@nada.kth.se`,
			}},
		},
		{
			input: `Nathaniel Borenstein <nsb@thumper.bellcore.com> (=?iso-8859-8?b?7eXs+SDv4SDp7Oj08A==?=)`,
			addrs: []*mail.Address{{
				Name:    `Nathaniel Borenstein`,
				Address: `nsb@thumper.bellcore.com`,
			}},
		},
	}
	for _, test := range tests {
		test := test

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
	}
	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			addrs, err := Parse(test.input)
			assert.NoError(t, err)
			assert.ElementsMatch(t, test.addrs, addrs)
		})
	}
}

func TestParseGroup(t *testing.T) {
	tests := []struct {
		input string
		addrs []*mail.Address
	}{
		{
			input: `A Group:Ed Jones <c@a.test>,joe@where.test,John <jdoe@one.test>;`,
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
			input: `Undisclosed recipients:;`,
			addrs: []*mail.Address{},
		},
	}
	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			addrs, err := Parse(test.input)
			assert.NoError(t, err)
			assert.ElementsMatch(t, test.addrs, addrs)
		})
	}
}

func TestParseStrangeAddresses(t *testing.T) {
	tests := []struct {
		input string
		addrs []*mail.Address
	}{
		{input: `<somebody@somebody.com >`},
		{input: `somebody@somebody.com,`},
		// {input: `somebody`},
		// {input: `<some random ascii text with spaces>`},
		{input: `=?UTF-8?B?PEJlemUgam3DqW5hPg==?= <somebody@somebody.com>`},
		{input: `somebody@somebody.com:81`},
		// {input: `undisclosed-recipients:`},
		// {input: `Name somebody@somewhere.com`},
		{input: `somebody@somewhere.com,`},
		{input: `<postmaster@[10.10.10.10]>`},
		// {input: `"GES-ELECTRONICS <GES-ELECTRONICS"@nos.ges.cz, a.s. <ges@ges.cz>>`},
		// {input: `=?ISO-8859-2?Q?Somebody_Somewhere?= <somebody@somewhere.com>, <somebody@somewhere.com,Somebody/AAA/BBB/CCC,>`},
		// {input: `somebody%somewhere...com`},
		// {input: `"Mailer Daemon" <>`},
		// {input: `=?windows-1250?Q?Spr=E1vce_syst=E9mu?=`},
		// {input: `"'somebody@somewhere.com.'"`},
		// {input: `Somebody Somewhere <somebody@somewhere. com> <somebody@somewhere.com>`},
		// {input: `"somebody@somewhere.com." <somebody@somewhere.com.>`},
		// {input: ` "comma, name"  <username@server.com>, another, name <address@server.com>`},
		// {input: ` normal name  <username@server.com>, (comment)All.(around)address@(the)server.com`},
		// {input: ` normal name  <username@server.com>, All.("comma, in comment")address@(the)server.com`},
	}
	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			_, err := Parse(test.input) // Ignore errors for now, these messages are maybe not valid; just checking the parser doesn't panic.
			assert.NoError(t, err)
		})
	}
}
