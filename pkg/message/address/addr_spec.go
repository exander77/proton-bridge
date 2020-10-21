package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type addrSpec struct {
	address string
}

func (w *walker) EnterAddrSpec(ctx *parser.AddrSpecContext) {
	logrus.Trace("Entering addrSpec")

	w.enter(&addrSpec{
		address: ctx.GetText(),
	})
}

func (w *walker) ExitAddrSpec(ctx *parser.AddrSpecContext) {
	logrus.Trace("Exiting addrSpec")

	type withAddrSpec interface {
		withAddrSpec(*addrSpec)
	}

	res := w.exit().(*addrSpec)

	if parent, ok := w.parent().(withAddrSpec); ok {
		parent.withAddrSpec(res)
	}
}
