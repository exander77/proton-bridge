package address

import (
	"fmt"
	"net/mail"

	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type walker struct {
	parser.BaseAddressParserListener
	antlr.DefaultErrorListener

	nodes []interface{}

	addresses []*mail.Address
	err       error
}

func (w *walker) enter(b interface{}) {
	w.nodes = append(w.nodes, b)
}

func (w *walker) exit() interface{} {
	b := w.nodes[len(w.nodes)-1]
	w.nodes = w.nodes[:len(w.nodes)-1]
	return b
}

func (w *walker) parent() (b interface{}) {
	return w.nodes[len(w.nodes)-1]
}

func (w *walker) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, msg string, _ antlr.RecognitionException) {
	w.err = fmt.Errorf("error parsing address: %v", msg)
}
