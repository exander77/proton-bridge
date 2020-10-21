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
	logrus.Trace("Exiting group")

	type withGroup interface {
		withGroup(*group)
	}

	res := w.exit().(*group)

	if parent, ok := w.parent().(withGroup); ok {
		parent.withGroup(res)
	}
}
