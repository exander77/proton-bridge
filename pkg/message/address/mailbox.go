package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type mailbox struct{}

func (w *walker) EnterMailbox(ctx *parser.MailboxContext) {
	logrus.Trace("Entering mailbox")
	w.enter(&mailbox{})
}

func (w *walker) ExitMailbox(ctx *parser.MailboxContext) {
	logrus.Trace("Exiting address")

	type withMailbox interface {
		withMailbox(*mailbox)
	}

	w.parent().(withMailbox).withMailbox(w.exit().(*mailbox))
}
