package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type domain struct {
	value string
}

func (d *domain) withDotAtom(dotAtom *dotAtom) {
	d.value = dotAtom.value
}

func (d *domain) withDomainLiteral(domainLiteral *domainLiteral) {
	d.value = domainLiteral.value
}

func (w *walker) EnterDomain(ctx *parser.DomainContext) {
	logrus.Trace("Entering domain")

	w.enter(&domain{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitDomain(ctx *parser.DomainContext) {
	logrus.Trace("Exiting domain")

	type withDomain interface {
		withDomain(*domain)
	}

	res := w.exit().(*domain)

	if parent, ok := w.parent().(withDomain); ok {
		parent.withDomain(res)
	}
}
