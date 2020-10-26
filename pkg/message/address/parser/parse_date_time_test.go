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

package parser

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func TestParseDateTimeValid(t *testing.T) {
	tests := []string{
		`Fri, 21 Nov 1997 09:55:06 -0600`,

		`2 Jan 2006 15:04:05 -0700`,
		`2 Jan 2006 15:04:05 MST`,
		`2 Jan 2006 15:04 -0700`,
		`2 Jan 2006 15:04 MST`,
		`2 Jan 06 15:04:05 -0700`,
		`2 Jan 06 15:04:05 MST`,
		`2 Jan 06 15:04 -0700`,
		`2 Jan 06 15:04 MST`,
		`02 Jan 2006 15:04:05 -0700`,
		`02 Jan 2006 15:04:05 MST`,
		`02 Jan 2006 15:04 -0700`,
		`02 Jan 2006 15:04 MST`,
		`02 Jan 06 15:04:05 -0700`,
		`02 Jan 06 15:04:05 MST`,
		`02 Jan 06 15:04 -0700`,
		`02 Jan 06 15:04 MST`,
		`Mon, 2 Jan 2006 15:04:05 -0700`,
		`Mon, 2 Jan 2006 15:04:05 MST`,
		`Mon, 2 Jan 2006 15:04 -0700`,
		`Mon, 2 Jan 2006 15:04 MST`,
		`Mon, 2 Jan 06 15:04:05 -0700`,
		`Mon, 2 Jan 06 15:04:05 MST`,
		`Mon, 2 Jan 06 15:04 -0700`,
		`Mon, 2 Jan 06 15:04 MST`,
		`Mon, 02 Jan 2006 15:04:05 -0700`,
		`Mon, 02 Jan 2006 15:04:05 MST`,
		`Mon, 02 Jan 2006 15:04 -0700`,
		`Mon, 02 Jan 2006 15:04 MST`,
		`Mon, 02 Jan 06 15:04:05 -0700`,
		`Mon, 02 Jan 06 15:04:05 MST`,
		`Mon, 02 Jan 06 15:04 -0700`,
		`Mon, 02 Jan 06 15:04 MST`,

		// Missing time zone is allowed; we assume GMT+0000.
		`Mon, 2 Jan 2006 15:04:05`,
	}
	for _, input := range tests {
		input := input

		t.Run(input, func(t *testing.T) {
			assert.Empty(t, parseDateTime(input))
		})
	}
}

func parseDateTime(dateTime string) string {
	stream := antlr.NewInputStream(dateTime)
	lexer := NewRFC5322Lexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	l := &errorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}

	p := NewRFC5322Parser(tokens)
	p.AddErrorListener(l)

	antlr.ParseTreeWalkerDefault.Walk(&BaseRFC5322ParserListener{}, p.DateTime())

	return l.msg
}
