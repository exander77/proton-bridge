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

type quotedContent struct {
	value string
}

func (c *quotedContent) withQtext(qtext *qtext) {
	c.value = qtext.value
}

func (c *quotedContent) withQuotedPair(quotedPair *quotedPair) {
	c.value = quotedPair.value
}

func (w *walker) EnterQuotedContent(ctx *parser.QuotedContentContext) {
	logrus.Trace("Entering quotedContent")
	w.enter(&quotedContent{})
}

func (w *walker) ExitQuotedContent(ctx *parser.QuotedContentContext) {
	logrus.Trace("Exiting quotedContent")

	type withQuotedContent interface {
		withQuotedContent(*quotedContent)
	}

	res := w.exit().(*quotedContent)

	if parent, ok := w.parent().(withQuotedContent); ok {
		parent.withQuotedContent(res)
	}
}