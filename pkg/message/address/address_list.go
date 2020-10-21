package address

import (
	"net/mail"

	"github.com/ProtonMail/proton-bridge/pkg/message/address/parser"
	"github.com/sirupsen/logrus"
)

type addressList struct {
	addresses []*mail.Address
}

func (a *addressList) withAddress(address *address) {
	a.addresses = append(a.addresses, &mail.Address{
		Name:    address.name,
		Address: address.address,
	})
}

func (w *walker) EnterAddressList(ctx *parser.AddressListContext) {
	logrus.Trace("Entering addressList")
	w.enter(&addressList{})
}

func (w *walker) ExitAddressList(ctx *parser.AddressListContext) {
	logrus.Trace("Exiting addressList")
	w.addresses = w.exit().(*addressList).addresses
}
