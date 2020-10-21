package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type word struct {
	value string
}

func (w *word) withAtom(atom *atom) {
	w.value = atom.value
}

func (w *word) withQuotedString(quotedString *quotedString) {
	w.value = quotedString.value
}

func (w *walker) EnterWord(ctx *parser.WordContext) {
	logrus.Trace("Entering word")

	w.enter(&word{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitWord(ctx *parser.WordContext) {
	logrus.Trace("Exiting word")

	type withWord interface {
		withWord(*word)
	}

	res := w.exit().(*word)

	if parent, ok := w.parent().(withWord); ok {
		parent.withWord(res)
	}
}
