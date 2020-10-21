package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type group struct{}

func (w *walker) EnterGroup(ctx *parser.GroupContext) {
	logrus.Trace("Entering group")
	w.enter(&group{})
}

func (w *walker) ExitGroup(ctx *parser.GroupContext) {
	logrus.Trace("Exiting address")

	type withGroup interface {
		withGroup(*group)
	}

	w.parent().(withGroup).withGroup(w.exit().(*group))
}
