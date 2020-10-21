package address

import (
	"net/mail"

	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Parse(input string) ([]*mail.Address, error) {
	l := parser.NewAddressLexer(antlr.NewInputStream(input))
	p := parser.NewAddressParser(antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel))
	w := &walker{}

	p.RemoveErrorListeners()
	p.AddErrorListener(w)

	antlr.ParseTreeWalkerDefault.Walk(w, p.AddressList())

	return w.addresses, w.err
}
