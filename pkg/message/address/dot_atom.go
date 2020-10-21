package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type dotAtom struct {
	value string
}

func (w *walker) EnterDotAtom(ctx *parser.DotAtomContext) {
	logrus.Trace("Entering dotAtom")

	w.enter(&dotAtom{
		value: ctx.DotAtomText().GetText(),
	})
}

func (w *walker) ExitDotAtom(ctx *parser.DotAtomContext) {
	logrus.Trace("Exiting dotAtom")

	type withDotAtom interface {
		withDotAtom(*dotAtom)
	}

	res := w.exit().(*dotAtom)

	if parent, ok := w.parent().(withDotAtom); ok {
		parent.withDotAtom(res)
	}
}
