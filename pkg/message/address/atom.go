package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type atom struct {
	value string
}

func (w *walker) EnterAtom(ctx *parser.AtomContext) {
	logrus.Trace("Entering atom")

	w.enter(&atom{
		value: ctx.AtextString().GetText(),
	})
}

func (w *walker) ExitAtom(ctx *parser.AtomContext) {
	logrus.Trace("Exiting atom")

	type withAtom interface {
		withAtom(*atom)
	}

	res := w.exit().(*atom)

	w.parent().(withAtom).withAtom(res)
}
