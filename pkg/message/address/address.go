package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type address struct {
	name, address string
}

func (a *address) withMailbox(mailbox *mailbox) {
	a.name = mailbox.name
	a.address = mailbox.address
}

func (a *address) withGroup(group *group) {
	// TODO
}

func (w *walker) EnterAddress(ctx *parser.AddressContext) {
	logrus.Trace("Entering address")
	w.enter(&address{})
}

func (w *walker) ExitAddress(ctx *parser.AddressContext) {
	logrus.Trace("Exiting address")

	res := w.exit().(*address)

	w.name = res.name
	w.address = res.address
}
