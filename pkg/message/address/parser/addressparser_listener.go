// Code generated from pkg/message/address/parser/AddressParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AddressParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AddressParserListener is a complete listener for a parse tree produced by AddressParser.
type AddressParserListener interface {
	antlr.ParseTreeListener

	// EnterQuotedPair is called when entering the quotedPair production.
	EnterQuotedPair(c *QuotedPairContext)

	// EnterFws is called when entering the fws production.
	EnterFws(c *FwsContext)

	// EnterCtext is called when entering the ctext production.
	EnterCtext(c *CtextContext)

	// EnterCcontent is called when entering the ccontent production.
	EnterCcontent(c *CcontentContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterCfws is called when entering the cfws production.
	EnterCfws(c *CfwsContext)

	// EnterAtext is called when entering the atext production.
	EnterAtext(c *AtextContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterDotAtomText is called when entering the dotAtomText production.
	EnterDotAtomText(c *DotAtomTextContext)

	// EnterDotAtom is called when entering the dotAtom production.
	EnterDotAtom(c *DotAtomContext)

	// EnterQtext is called when entering the qtext production.
	EnterQtext(c *QtextContext)

	// EnterQcontent is called when entering the qcontent production.
	EnterQcontent(c *QcontentContext)

	// EnterQuotedString is called when entering the quotedString production.
	EnterQuotedString(c *QuotedStringContext)

	// EnterWord is called when entering the word production.
	EnterWord(c *WordContext)

	// EnterPhrase is called when entering the phrase production.
	EnterPhrase(c *PhraseContext)

	// EnterUnstructured is called when entering the unstructured production.
	EnterUnstructured(c *UnstructuredContext)

	// EnterAddress is called when entering the address production.
	EnterAddress(c *AddressContext)

	// EnterMailbox is called when entering the mailbox production.
	EnterMailbox(c *MailboxContext)

	// EnterNameAddr is called when entering the nameAddr production.
	EnterNameAddr(c *NameAddrContext)

	// EnterAngleAddr is called when entering the angleAddr production.
	EnterAngleAddr(c *AngleAddrContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterDisplayName is called when entering the displayName production.
	EnterDisplayName(c *DisplayNameContext)

	// EnterMailboxList is called when entering the mailboxList production.
	EnterMailboxList(c *MailboxListContext)

	// EnterAddressList is called when entering the addressList production.
	EnterAddressList(c *AddressListContext)

	// EnterGroupList is called when entering the groupList production.
	EnterGroupList(c *GroupListContext)

	// EnterAddrSpec is called when entering the addrSpec production.
	EnterAddrSpec(c *AddrSpecContext)

	// EnterLocalPart is called when entering the localPart production.
	EnterLocalPart(c *LocalPartContext)

	// EnterDomain is called when entering the domain production.
	EnterDomain(c *DomainContext)

	// EnterDomainLiteral is called when entering the domainLiteral production.
	EnterDomainLiteral(c *DomainLiteralContext)

	// EnterDtext is called when entering the dtext production.
	EnterDtext(c *DtextContext)

	// EnterCrlf is called when entering the crlf production.
	EnterCrlf(c *CrlfContext)

	// EnterWsp is called when entering the wsp production.
	EnterWsp(c *WspContext)

	// EnterVchar is called when entering the vchar production.
	EnterVchar(c *VcharContext)

	// ExitQuotedPair is called when exiting the quotedPair production.
	ExitQuotedPair(c *QuotedPairContext)

	// ExitFws is called when exiting the fws production.
	ExitFws(c *FwsContext)

	// ExitCtext is called when exiting the ctext production.
	ExitCtext(c *CtextContext)

	// ExitCcontent is called when exiting the ccontent production.
	ExitCcontent(c *CcontentContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitCfws is called when exiting the cfws production.
	ExitCfws(c *CfwsContext)

	// ExitAtext is called when exiting the atext production.
	ExitAtext(c *AtextContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitDotAtomText is called when exiting the dotAtomText production.
	ExitDotAtomText(c *DotAtomTextContext)

	// ExitDotAtom is called when exiting the dotAtom production.
	ExitDotAtom(c *DotAtomContext)

	// ExitQtext is called when exiting the qtext production.
	ExitQtext(c *QtextContext)

	// ExitQcontent is called when exiting the qcontent production.
	ExitQcontent(c *QcontentContext)

	// ExitQuotedString is called when exiting the quotedString production.
	ExitQuotedString(c *QuotedStringContext)

	// ExitWord is called when exiting the word production.
	ExitWord(c *WordContext)

	// ExitPhrase is called when exiting the phrase production.
	ExitPhrase(c *PhraseContext)

	// ExitUnstructured is called when exiting the unstructured production.
	ExitUnstructured(c *UnstructuredContext)

	// ExitAddress is called when exiting the address production.
	ExitAddress(c *AddressContext)

	// ExitMailbox is called when exiting the mailbox production.
	ExitMailbox(c *MailboxContext)

	// ExitNameAddr is called when exiting the nameAddr production.
	ExitNameAddr(c *NameAddrContext)

	// ExitAngleAddr is called when exiting the angleAddr production.
	ExitAngleAddr(c *AngleAddrContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitDisplayName is called when exiting the displayName production.
	ExitDisplayName(c *DisplayNameContext)

	// ExitMailboxList is called when exiting the mailboxList production.
	ExitMailboxList(c *MailboxListContext)

	// ExitAddressList is called when exiting the addressList production.
	ExitAddressList(c *AddressListContext)

	// ExitGroupList is called when exiting the groupList production.
	ExitGroupList(c *GroupListContext)

	// ExitAddrSpec is called when exiting the addrSpec production.
	ExitAddrSpec(c *AddrSpecContext)

	// ExitLocalPart is called when exiting the localPart production.
	ExitLocalPart(c *LocalPartContext)

	// ExitDomain is called when exiting the domain production.
	ExitDomain(c *DomainContext)

	// ExitDomainLiteral is called when exiting the domainLiteral production.
	ExitDomainLiteral(c *DomainLiteralContext)

	// ExitDtext is called when exiting the dtext production.
	ExitDtext(c *DtextContext)

	// ExitCrlf is called when exiting the crlf production.
	ExitCrlf(c *CrlfContext)

	// ExitWsp is called when exiting the wsp production.
	ExitWsp(c *WspContext)

	// ExitVchar is called when exiting the vchar production.
	ExitVchar(c *VcharContext)
}
