// Code generated from pkg/message/address/parser/Address.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Address

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAddressListener is a complete listener for a parse tree produced by Address.
type BaseAddressListener struct{}

var _ AddressListener = &BaseAddressListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAddressListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAddressListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAddressListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAddressListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuotedPair is called when production quotedPair is entered.
func (s *BaseAddressListener) EnterQuotedPair(ctx *QuotedPairContext) {}

// ExitQuotedPair is called when production quotedPair is exited.
func (s *BaseAddressListener) ExitQuotedPair(ctx *QuotedPairContext) {}

// EnterFws is called when production fws is entered.
func (s *BaseAddressListener) EnterFws(ctx *FwsContext) {}

// ExitFws is called when production fws is exited.
func (s *BaseAddressListener) ExitFws(ctx *FwsContext) {}

// EnterCtext is called when production ctext is entered.
func (s *BaseAddressListener) EnterCtext(ctx *CtextContext) {}

// ExitCtext is called when production ctext is exited.
func (s *BaseAddressListener) ExitCtext(ctx *CtextContext) {}

// EnterCcontent is called when production ccontent is entered.
func (s *BaseAddressListener) EnterCcontent(ctx *CcontentContext) {}

// ExitCcontent is called when production ccontent is exited.
func (s *BaseAddressListener) ExitCcontent(ctx *CcontentContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseAddressListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseAddressListener) ExitComment(ctx *CommentContext) {}

// EnterCfws is called when production cfws is entered.
func (s *BaseAddressListener) EnterCfws(ctx *CfwsContext) {}

// ExitCfws is called when production cfws is exited.
func (s *BaseAddressListener) ExitCfws(ctx *CfwsContext) {}

// EnterAtext is called when production atext is entered.
func (s *BaseAddressListener) EnterAtext(ctx *AtextContext) {}

// ExitAtext is called when production atext is exited.
func (s *BaseAddressListener) ExitAtext(ctx *AtextContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseAddressListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseAddressListener) ExitAtom(ctx *AtomContext) {}

// EnterDotAtomText is called when production dotAtomText is entered.
func (s *BaseAddressListener) EnterDotAtomText(ctx *DotAtomTextContext) {}

// ExitDotAtomText is called when production dotAtomText is exited.
func (s *BaseAddressListener) ExitDotAtomText(ctx *DotAtomTextContext) {}

// EnterDotAtom is called when production dotAtom is entered.
func (s *BaseAddressListener) EnterDotAtom(ctx *DotAtomContext) {}

// ExitDotAtom is called when production dotAtom is exited.
func (s *BaseAddressListener) ExitDotAtom(ctx *DotAtomContext) {}

// EnterQtext is called when production qtext is entered.
func (s *BaseAddressListener) EnterQtext(ctx *QtextContext) {}

// ExitQtext is called when production qtext is exited.
func (s *BaseAddressListener) ExitQtext(ctx *QtextContext) {}

// EnterQcontent is called when production qcontent is entered.
func (s *BaseAddressListener) EnterQcontent(ctx *QcontentContext) {}

// ExitQcontent is called when production qcontent is exited.
func (s *BaseAddressListener) ExitQcontent(ctx *QcontentContext) {}

// EnterQuotedString is called when production quotedString is entered.
func (s *BaseAddressListener) EnterQuotedString(ctx *QuotedStringContext) {}

// ExitQuotedString is called when production quotedString is exited.
func (s *BaseAddressListener) ExitQuotedString(ctx *QuotedStringContext) {}

// EnterWord is called when production word is entered.
func (s *BaseAddressListener) EnterWord(ctx *WordContext) {}

// ExitWord is called when production word is exited.
func (s *BaseAddressListener) ExitWord(ctx *WordContext) {}

// EnterPhrase is called when production phrase is entered.
func (s *BaseAddressListener) EnterPhrase(ctx *PhraseContext) {}

// ExitPhrase is called when production phrase is exited.
func (s *BaseAddressListener) ExitPhrase(ctx *PhraseContext) {}

// EnterUnstructured is called when production unstructured is entered.
func (s *BaseAddressListener) EnterUnstructured(ctx *UnstructuredContext) {}

// ExitUnstructured is called when production unstructured is exited.
func (s *BaseAddressListener) ExitUnstructured(ctx *UnstructuredContext) {}

// EnterAddress is called when production address is entered.
func (s *BaseAddressListener) EnterAddress(ctx *AddressContext) {}

// ExitAddress is called when production address is exited.
func (s *BaseAddressListener) ExitAddress(ctx *AddressContext) {}

// EnterMailbox is called when production mailbox is entered.
func (s *BaseAddressListener) EnterMailbox(ctx *MailboxContext) {}

// ExitMailbox is called when production mailbox is exited.
func (s *BaseAddressListener) ExitMailbox(ctx *MailboxContext) {}

// EnterNameAddr is called when production nameAddr is entered.
func (s *BaseAddressListener) EnterNameAddr(ctx *NameAddrContext) {}

// ExitNameAddr is called when production nameAddr is exited.
func (s *BaseAddressListener) ExitNameAddr(ctx *NameAddrContext) {}

// EnterAngleAddr is called when production angleAddr is entered.
func (s *BaseAddressListener) EnterAngleAddr(ctx *AngleAddrContext) {}

// ExitAngleAddr is called when production angleAddr is exited.
func (s *BaseAddressListener) ExitAngleAddr(ctx *AngleAddrContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseAddressListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseAddressListener) ExitGroup(ctx *GroupContext) {}

// EnterDisplayName is called when production displayName is entered.
func (s *BaseAddressListener) EnterDisplayName(ctx *DisplayNameContext) {}

// ExitDisplayName is called when production displayName is exited.
func (s *BaseAddressListener) ExitDisplayName(ctx *DisplayNameContext) {}

// EnterMailboxList is called when production mailboxList is entered.
func (s *BaseAddressListener) EnterMailboxList(ctx *MailboxListContext) {}

// ExitMailboxList is called when production mailboxList is exited.
func (s *BaseAddressListener) ExitMailboxList(ctx *MailboxListContext) {}

// EnterAddressList is called when production addressList is entered.
func (s *BaseAddressListener) EnterAddressList(ctx *AddressListContext) {}

// ExitAddressList is called when production addressList is exited.
func (s *BaseAddressListener) ExitAddressList(ctx *AddressListContext) {}

// EnterGroupList is called when production groupList is entered.
func (s *BaseAddressListener) EnterGroupList(ctx *GroupListContext) {}

// ExitGroupList is called when production groupList is exited.
func (s *BaseAddressListener) ExitGroupList(ctx *GroupListContext) {}

// EnterAddrSpec is called when production addrSpec is entered.
func (s *BaseAddressListener) EnterAddrSpec(ctx *AddrSpecContext) {}

// ExitAddrSpec is called when production addrSpec is exited.
func (s *BaseAddressListener) ExitAddrSpec(ctx *AddrSpecContext) {}

// EnterLocalPart is called when production localPart is entered.
func (s *BaseAddressListener) EnterLocalPart(ctx *LocalPartContext) {}

// ExitLocalPart is called when production localPart is exited.
func (s *BaseAddressListener) ExitLocalPart(ctx *LocalPartContext) {}

// EnterDomain is called when production domain is entered.
func (s *BaseAddressListener) EnterDomain(ctx *DomainContext) {}

// ExitDomain is called when production domain is exited.
func (s *BaseAddressListener) ExitDomain(ctx *DomainContext) {}

// EnterDomainLiteral is called when production domainLiteral is entered.
func (s *BaseAddressListener) EnterDomainLiteral(ctx *DomainLiteralContext) {}

// ExitDomainLiteral is called when production domainLiteral is exited.
func (s *BaseAddressListener) ExitDomainLiteral(ctx *DomainLiteralContext) {}

// EnterDtext is called when production dtext is entered.
func (s *BaseAddressListener) EnterDtext(ctx *DtextContext) {}

// ExitDtext is called when production dtext is exited.
func (s *BaseAddressListener) ExitDtext(ctx *DtextContext) {}

// EnterWsp is called when production wsp is entered.
func (s *BaseAddressListener) EnterWsp(ctx *WspContext) {}

// ExitWsp is called when production wsp is exited.
func (s *BaseAddressListener) ExitWsp(ctx *WspContext) {}

// EnterVchar is called when production vchar is entered.
func (s *BaseAddressListener) EnterVchar(ctx *VcharContext) {}

// ExitVchar is called when production vchar is exited.
func (s *BaseAddressListener) ExitVchar(ctx *VcharContext) {}
