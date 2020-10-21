package address

import (
	"strings"

	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type nameAddr struct {
	name, address string
}

func (a *nameAddr) withDisplayName(displayName *displayName) {
	a.name = strings.Join(displayName.words, " ")
}

func (a *nameAddr) withAngleAddr(angleAddr *angleAddr) {
	a.address = angleAddr.address
}

func (w *walker) EnterNameAddr(ctx *parser.NameAddrContext) {
	logrus.Trace("Entering nameAddr")
	w.enter(&nameAddr{})
}

func (w *walker) ExitNameAddr(ctx *parser.NameAddrContext) {
	logrus.Trace("Exiting nameAddr")

	type withNameAddr interface {
		withNameAddr(*nameAddr)
	}

	res := w.exit().(*nameAddr)

	w.parent().(withNameAddr).withNameAddr(res)
}
