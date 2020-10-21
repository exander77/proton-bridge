package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type quotedString struct {
	value string
}

func (s *quotedString) withQuotedValue(quotedValue *quotedValue) {
	s.value = quotedValue.value
}

func (w *walker) EnterQuotedString(ctx *parser.QuotedStringContext) {
	logrus.Trace("Entering quotedString")
	w.enter(&quotedString{})
}

func (w *walker) ExitQuotedString(ctx *parser.QuotedStringContext) {
	logrus.Trace("Exiting quotedString")

	type withQuotedString interface {
		withQuotedString(*quotedString)
	}

	res := w.exit().(*quotedString)

	if parent, ok := w.parent().(withQuotedString); ok {
		parent.withQuotedString(res)
	}
}
