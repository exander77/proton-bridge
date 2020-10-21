// Code generated from pkg/message/address/parser/AddressParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AddressParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAddressParserListener is a complete listener for a parse tree produced by AddressParser.
type BaseAddressParserListener struct{}

var _ AddressParserListener = &BaseAddressParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAddressParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAddressParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAddressParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAddressParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuotedPair is called when production quotedPair is entered.
func (s *BaseAddressParserListener) EnterQuotedPair(ctx *QuotedPairContext) {}

// ExitQuotedPair is called when production quotedPair is exited.
func (s *BaseAddressParserListener) ExitQuotedPair(ctx *QuotedPairContext) {}

// EnterFws is called when production fws is entered.
func (s *BaseAddressParserListener) EnterFws(ctx *FwsContext) {}

// ExitFws is called when production fws is exited.
func (s *BaseAddressParserListener) ExitFws(ctx *FwsContext) {}

// EnterCtext is called when production ctext is entered.
func (s *BaseAddressParserListener) EnterCtext(ctx *CtextContext) {}

// ExitCtext is called when production ctext is exited.
func (s *BaseAddressParserListener) ExitCtext(ctx *CtextContext) {}

// EnterCcontent is called when production ccontent is entered.
func (s *BaseAddressParserListener) EnterCcontent(ctx *CcontentContext) {}

// ExitCcontent is called when production ccontent is exited.
func (s *BaseAddressParserListener) ExitCcontent(ctx *CcontentContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseAddressParserListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseAddressParserListener) ExitComment(ctx *CommentContext) {}

// EnterCfws is called when production cfws is entered.
func (s *BaseAddressParserListener) EnterCfws(ctx *CfwsContext) {}

// ExitCfws is called when production cfws is exited.
func (s *BaseAddressParserListener) ExitCfws(ctx *CfwsContext) {}

// EnterAtext is called when production atext is entered.
func (s *BaseAddressParserListener) EnterAtext(ctx *AtextContext) {}

// ExitAtext is called when production atext is exited.
func (s *BaseAddressParserListener) ExitAtext(ctx *AtextContext) {}

// EnterAtextString is called when production atextString is entered.
func (s *BaseAddressParserListener) EnterAtextString(ctx *AtextStringContext) {}

// ExitAtextString is called when production atextString is exited.
func (s *BaseAddressParserListener) ExitAtextString(ctx *AtextStringContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseAddressParserListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseAddressParserListener) ExitAtom(ctx *AtomContext) {}

// EnterDotAtomText is called when production dotAtomText is entered.
func (s *BaseAddressParserListener) EnterDotAtomText(ctx *DotAtomTextContext) {}

// ExitDotAtomText is called when production dotAtomText is exited.
func (s *BaseAddressParserListener) ExitDotAtomText(ctx *DotAtomTextContext) {}

// EnterDotAtom is called when production dotAtom is entered.
func (s *BaseAddressParserListener) EnterDotAtom(ctx *DotAtomContext) {}

// ExitDotAtom is called when production dotAtom is exited.
func (s *BaseAddressParserListener) ExitDotAtom(ctx *DotAtomContext) {}

// EnterQtext is called when production qtext is entered.
func (s *BaseAddressParserListener) EnterQtext(ctx *QtextContext) {}

// ExitQtext is called when production qtext is exited.
func (s *BaseAddressParserListener) ExitQtext(ctx *QtextContext) {}

// EnterQcontent is called when production qcontent is entered.
func (s *BaseAddressParserListener) EnterQcontent(ctx *QcontentContext) {}

// ExitQcontent is called when production qcontent is exited.
func (s *BaseAddressParserListener) ExitQcontent(ctx *QcontentContext) {}

// EnterQuotedString is called when production quotedString is entered.
func (s *BaseAddressParserListener) EnterQuotedString(ctx *QuotedStringContext) {}

// ExitQuotedString is called when production quotedString is exited.
func (s *BaseAddressParserListener) ExitQuotedString(ctx *QuotedStringContext) {}

// EnterWord is called when production word is entered.
func (s *BaseAddressParserListener) EnterWord(ctx *WordContext) {}

// ExitWord is called when production word is exited.
func (s *BaseAddressParserListener) ExitWord(ctx *WordContext) {}

// EnterUnstructured is called when production unstructured is entered.
func (s *BaseAddressParserListener) EnterUnstructured(ctx *UnstructuredContext) {}

// ExitUnstructured is called when production unstructured is exited.
func (s *BaseAddressParserListener) ExitUnstructured(ctx *UnstructuredContext) {}

// EnterAddress is called when production address is entered.
func (s *BaseAddressParserListener) EnterAddress(ctx *AddressContext) {}

// ExitAddress is called when production address is exited.
func (s *BaseAddressParserListener) ExitAddress(ctx *AddressContext) {}

// EnterMailbox is called when production mailbox is entered.
func (s *BaseAddressParserListener) EnterMailbox(ctx *MailboxContext) {}

// ExitMailbox is called when production mailbox is exited.
func (s *BaseAddressParserListener) ExitMailbox(ctx *MailboxContext) {}

// EnterNameAddr is called when production nameAddr is entered.
func (s *BaseAddressParserListener) EnterNameAddr(ctx *NameAddrContext) {}

// ExitNameAddr is called when production nameAddr is exited.
func (s *BaseAddressParserListener) ExitNameAddr(ctx *NameAddrContext) {}

// EnterAngleAddr is called when production angleAddr is entered.
func (s *BaseAddressParserListener) EnterAngleAddr(ctx *AngleAddrContext) {}

// ExitAngleAddr is called when production angleAddr is exited.
func (s *BaseAddressParserListener) ExitAngleAddr(ctx *AngleAddrContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseAddressParserListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseAddressParserListener) ExitGroup(ctx *GroupContext) {}

// EnterDisplayName is called when production displayName is entered.
func (s *BaseAddressParserListener) EnterDisplayName(ctx *DisplayNameContext) {}

// ExitDisplayName is called when production displayName is exited.
func (s *BaseAddressParserListener) ExitDisplayName(ctx *DisplayNameContext) {}

// EnterMailboxList is called when production mailboxList is entered.
func (s *BaseAddressParserListener) EnterMailboxList(ctx *MailboxListContext) {}

// ExitMailboxList is called when production mailboxList is exited.
func (s *BaseAddressParserListener) ExitMailboxList(ctx *MailboxListContext) {}

// EnterAddressList is called when production addressList is entered.
func (s *BaseAddressParserListener) EnterAddressList(ctx *AddressListContext) {}

// ExitAddressList is called when production addressList is exited.
func (s *BaseAddressParserListener) ExitAddressList(ctx *AddressListContext) {}

// EnterGroupList is called when production groupList is entered.
func (s *BaseAddressParserListener) EnterGroupList(ctx *GroupListContext) {}

// ExitGroupList is called when production groupList is exited.
func (s *BaseAddressParserListener) ExitGroupList(ctx *GroupListContext) {}

// EnterAddrSpec is called when production addrSpec is entered.
func (s *BaseAddressParserListener) EnterAddrSpec(ctx *AddrSpecContext) {}

// ExitAddrSpec is called when production addrSpec is exited.
func (s *BaseAddressParserListener) ExitAddrSpec(ctx *AddrSpecContext) {}

// EnterLocalPart is called when production localPart is entered.
func (s *BaseAddressParserListener) EnterLocalPart(ctx *LocalPartContext) {}

// ExitLocalPart is called when production localPart is exited.
func (s *BaseAddressParserListener) ExitLocalPart(ctx *LocalPartContext) {}

// EnterDomain is called when production domain is entered.
func (s *BaseAddressParserListener) EnterDomain(ctx *DomainContext) {}

// ExitDomain is called when production domain is exited.
func (s *BaseAddressParserListener) ExitDomain(ctx *DomainContext) {}

// EnterDomainLiteral is called when production domainLiteral is entered.
func (s *BaseAddressParserListener) EnterDomainLiteral(ctx *DomainLiteralContext) {}

// ExitDomainLiteral is called when production domainLiteral is exited.
func (s *BaseAddressParserListener) ExitDomainLiteral(ctx *DomainLiteralContext) {}

// EnterDtext is called when production dtext is entered.
func (s *BaseAddressParserListener) EnterDtext(ctx *DtextContext) {}

// ExitDtext is called when production dtext is exited.
func (s *BaseAddressParserListener) ExitDtext(ctx *DtextContext) {}

// EnterCrlf is called when production crlf is entered.
func (s *BaseAddressParserListener) EnterCrlf(ctx *CrlfContext) {}

// ExitCrlf is called when production crlf is exited.
func (s *BaseAddressParserListener) ExitCrlf(ctx *CrlfContext) {}

// EnterWsp is called when production wsp is entered.
func (s *BaseAddressParserListener) EnterWsp(ctx *WspContext) {}

// ExitWsp is called when production wsp is exited.
func (s *BaseAddressParserListener) ExitWsp(ctx *WspContext) {}

// EnterVchar is called when production vchar is entered.
func (s *BaseAddressParserListener) EnterVchar(ctx *VcharContext) {}

// ExitVchar is called when production vchar is exited.
func (s *BaseAddressParserListener) ExitVchar(ctx *VcharContext) {}
