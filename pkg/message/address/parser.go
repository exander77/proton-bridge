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

	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Parse parses one or more valid RFC5322 (with RFC2047) addresses.
func Parse(input string) ([]*mail.Address, error) {
	if len(input) == 0 {
		return []*mail.Address{}, nil
	}

	l := parser.NewAddressLexer(antlr.NewInputStream(input))
	p := parser.NewAddressParser(antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel))
	w := &walker{}

	p.RemoveErrorListeners()
	p.AddErrorListener(w)

	antlr.ParseTreeWalkerDefault.Walk(w, p.AddressList())

	return w.addresses, w.err
}
