package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type mailbox struct {
	name, address string
}

func (m *mailbox) withNameAddr(nameAddr *nameAddr) {
	m.name = nameAddr.name
	m.address = nameAddr.address
}

func (m *mailbox) withAddrSpec(addrSpec *addrSpec) {
	m.address = addrSpec.address
}

func (w *walker) EnterMailbox(ctx *parser.MailboxContext) {
	logrus.Trace("Entering mailbox")
	w.enter(&mailbox{})
}

func (w *walker) ExitMailbox(ctx *parser.MailboxContext) {
	logrus.Trace("Exiting mailbox")

	type withMailbox interface {
		withMailbox(*mailbox)
	}

	res := w.exit().(*mailbox)

	w.parent().(withMailbox).withMailbox(res)
}
