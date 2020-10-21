package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type domainLiteral struct {
	value string
}

func (w *walker) EnterDomainLiteral(ctx *parser.DomainLiteralContext) {
	logrus.Trace("Entering domainLiteral")

	w.enter(&domainLiteral{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitDomainLiteral(ctx *parser.DomainLiteralContext) {
	logrus.Trace("Exiting domainLiteral")

	type withDomainLiteral interface {
		withDomainLiteral(*domainLiteral)
	}

	res := w.exit().(*domainLiteral)

	if parent, ok := w.parent().(withDomainLiteral); ok {
		parent.withDomainLiteral(res)
	}
}
