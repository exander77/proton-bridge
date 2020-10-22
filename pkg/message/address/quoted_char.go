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
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type quotedChar struct {
	value string
}

func (w *walker) EnterQuotedChar(ctx *parser.QuotedCharContext) {
	logrus.Trace("Entering quotedChar")

	w.enter(&quotedChar{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitQuotedChar(ctx *parser.QuotedCharContext) {
	logrus.Trace("Exiting quotedChar")

	type withQuotedChar interface {
		withQuotedChar(*quotedChar)
	}

	res := w.exit().(*quotedChar)

	if parent, ok := w.parent().(withQuotedChar); ok {
		parent.withQuotedChar(res)
	}
}