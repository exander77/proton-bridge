package address

import (
	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Parse(input string) (string, string, error) {
	l := parser.NewAddressLexer(antlr.NewInputStream(input))
	p := parser.NewAddressParser(antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel))
	w := &walker{}

	p.RemoveErrorListeners()
	p.AddErrorListener(w)

	antlr.ParseTreeWalkerDefault.Walk(w, p.Address())

	return w.name, w.address, w.err
}
