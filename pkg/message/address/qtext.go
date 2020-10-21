package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type qtext struct {
	value string
}

func (w *walker) EnterQtext(ctx *parser.QtextContext) {
	logrus.Trace("Entering qtext")

	w.enter(&qtext{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitQtext(ctx *parser.QtextContext) {
	logrus.Trace("Exiting qtext")

	type withQtext interface {
		withQtext(*qtext)
	}

	res := w.exit().(*qtext)

	if parent, ok := w.parent().(withQtext); ok {
		parent.withQtext(res)
	}
}
