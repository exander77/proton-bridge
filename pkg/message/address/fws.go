package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type fws struct {
	value string
}

func (w *walker) EnterFws(ctx *parser.FwsContext) {
	logrus.Trace("Entering fws")

	w.enter(&fws{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitFws(ctx *parser.FwsContext) {
	logrus.Trace("Exiting fws")

	type withFws interface {
		withFws(*fws)
	}

	res := w.exit().(*fws)

	if parent, ok := w.parent().(withFws); ok {
		parent.withFws(res)
	}
}
