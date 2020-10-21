package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type localPart struct {
	value string
}

func (p *localPart) withDotAtom(dotAtom *dotAtom) {
	p.value = dotAtom.value
}

func (p *localPart) withQuotedString(quotedString *quotedString) {
	p.value = quotedString.value
}

func (w *walker) EnterLocalPart(ctx *parser.LocalPartContext) {
	logrus.Trace("Entering localPart")

	w.enter(&localPart{})
}

func (w *walker) ExitLocalPart(ctx *parser.LocalPartContext) {
	logrus.Trace("Exiting localPart")

	type withLocalPart interface {
		withLocalPart(*localPart)
	}

	res := w.exit().(*localPart)

	if parent, ok := w.parent().(withLocalPart); ok {
		parent.withLocalPart(res)
	}
}
