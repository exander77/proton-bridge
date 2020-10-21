package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type angleAddr struct {
	address string
}

func (a *angleAddr) withAddrSpec(addrSpec *addrSpec) {
	a.address = addrSpec.address
}

func (w *walker) EnterAngleAddr(ctx *parser.AngleAddrContext) {
	logrus.Trace("Entering angleAddr")

	w.enter(&angleAddr{
		address: ctx.GetText(),
	})
}

func (w *walker) ExitAngleAddr(ctx *parser.AngleAddrContext) {
	logrus.Trace("Exiting angleAddr")

	type withAngleAddr interface {
		withAngleAddr(*angleAddr)
	}

	res := w.exit().(*angleAddr)

	if parent, ok := w.parent().(withAngleAddr); ok {
		parent.withAngleAddr(res)
	}
}
