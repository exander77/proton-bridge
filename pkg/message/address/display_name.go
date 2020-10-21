package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type displayName struct {
	words []string
}

func (n *displayName) withWord(word *word) {
	n.words = append(n.words, word.value)
}

func (w *walker) EnterDisplayName(ctx *parser.DisplayNameContext) {
	logrus.Trace("Entering displayName")

	w.enter(&displayName{})
}

func (w *walker) ExitDisplayName(ctx *parser.DisplayNameContext) {
	logrus.Trace("Exiting displayName")

	type withDisplayName interface {
		withDisplayName(*displayName)
	}

	res := w.exit().(*displayName)

	w.parent().(withDisplayName).withDisplayName(res)
}
