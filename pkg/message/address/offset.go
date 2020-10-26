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
	"fmt"
	"time"

	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type offset struct {
	rep   string
	value int
}

func (w *walker) EnterOffset(ctx *parser.OffsetContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering offset")

	sgn := ctx.GetText()[0:1]
	hrs := ctx.GetText()[1:3]
	min := ctx.GetText()[3:5]

	dur, err := time.ParseDuration(fmt.Sprintf("%v%vh%vm", sgn, hrs, min))
	if err != nil {
		w.err = err
	}

	w.enter(&offset{
		rep:   ctx.GetText(),
		value: int(dur.Seconds()),
	})
}

func (w *walker) ExitOffset(ctx *parser.OffsetContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting offset")

	type withOffset interface {
		withOffset(*offset)
	}

	res := w.exit().(*offset)

	if parent, ok := w.parent().(withOffset); ok {
		parent.withOffset(res)
	}
}
