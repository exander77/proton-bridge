package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type quotedString struct {
	value string
}

func (w *walker) EnterQuotedString(ctx *parser.QuotedStringContext) {
	logrus.Trace("Entering quotedString")

	w.enter(&quotedString{
		value: ctx.QuotedStringValue().GetText(),
	})
}

func (w *walker) ExitQuotedString(ctx *parser.QuotedStringContext) {
	logrus.Trace("Exiting quotedString")

	type withQuotedString interface {
		withQuotedString(*quotedString)
	}

	res := w.exit().(*quotedString)

	w.parent().(withQuotedString).withQuotedString(res)
}
