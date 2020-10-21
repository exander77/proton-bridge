package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type displayName struct {
	name string
}

func (w *walker) EnterDisplayName(ctx *parser.DisplayNameContext) {
	logrus.Trace("Entering displayName")

	w.enter(&displayName{
		name: ctx.GetText(),
	})
}

func (w *walker) ExitDisplayName(ctx *parser.DisplayNameContext) {
	logrus.Trace("Exiting displayName")

	type withDisplayName interface {
		withDisplayName(*displayName)
	}

	res := w.exit().(*displayName)

	w.parent().(withDisplayName).withDisplayName(res)
}
