package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type address struct {
	name, address string
}

func (w *walker) EnterAddress(ctx *parser.AddressContext) {
	logrus.Trace("Entering address")
	w.enter(&address{})
}

func (w *walker) ExitRoot(ctx *parser.AddressContext) {
	logrus.Trace("Exiting address")

	res := w.exit().(*address)

	w.name = res.name
	w.address = res.address
}
