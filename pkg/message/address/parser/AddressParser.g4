parser grammar AddressParser;

options { tokenVocab=AddressLexer; }


// -------------------
// 3.2. Lexical tokens
// -------------------

quotedPair
	: Backslash (vchar | wsp)
//| obsQP
	;


fws 
	: (wsp* crlf)? wsp+
//| obsFWS
	;

ctext
	: Exclamation
	| DQuote
	| Hash
	| Dollar
	| Percent
	| Ampersand
	| SQuote
	| Asterisk
	| Plus
	| Comma
	| Minus
	| Period
	| Slash
	| Digit
	| Colon
	| Semicolon
	| Less
	| Equal
	| Greater
	| Question
	| At
	| AlphaUpper
	| LBracket
	| RBracket
	| Caret
	| Underscore
	| Backtick
	| AlphaLower
	| LCurly
	| Pipe
	| RCurly
	| Tilde
//| obsCtext
	;

ccontent
	: ctext
	| quotedPair
	| comment
	;

comment: LParens (fws? ccontent)* fws? RParens;

cfws
	: (fws? comment)+ fws?
	| fws
	;

atext
	: AlphaUpper 
	| AlphaLower 
	| Digit
	| Exclamation 
	| Hash
	| Dollar 
	| Percent
	| Ampersand 
	| SQuote
	| Asterisk 
	| Plus
	| Minus 
	| Slash
	| Equal 
	| Question
	| Caret 
	| Underscore
	| Backtick 
	| LCurly
	| Pipe 
	| RCurly
	| Tilde
	;

atextString: atext+;

atom: cfws? atextString cfws?;

dotAtomText: atextString (Period atextString)*;

dotAtom: cfws? dotAtomText cfws?;

qtext
	: Exclamation
	| Hash
	| Dollar
	| Percent
	| Ampersand
	| SQuote
	| LParens
	| RParens
	| Asterisk
	| Plus
	| Comma
	| Minus
	| Period
	| Slash
	| Digit
	| Colon
	| Semicolon
	| Less
	| Equal
	| Greater
	| Question
	| At
	| AlphaUpper
	| LBracket
	| RBracket
	| Caret
	| Underscore
	| Backtick
	| AlphaLower
	| LCurly
	| Pipe
	| RCurly
	| Tilde
//| obsQtext
	;

qcontent
	: qtext
	| quotedPair
	;

quotedString: cfws? DQuote (fws? qcontent)* fws? DQuote cfws?;

word
	: atom
	| quotedString
	;

unstructured
	: (fws? vchar)* wsp*
//| obsUnstruct
	;


// --------------------------
// 3.4. Address Specification
// --------------------------

address
	: mailbox
	| group
	;

mailbox
	: nameAddr
	| addrSpec
	;

nameAddr: displayName? angleAddr;

angleAddr
	: cfws? Less addrSpec Greater cfws?
//| obsAngleAddr
	;

group: displayName Colon groupList? Semicolon cfws?;

displayName
	: word+
//| obsPhrase
	;

mailboxList
	: mailbox (Comma mailbox)*
//| obsMboxList
	;

addressList
	: address (Comma address)*
//| obsAddrList
	;

groupList
	: mailboxList
	| cfws
//| obsGroupList
	;

addrSpec: localPart At domain;

localPart
	: dotAtom 
	| quotedString
//| obsLocalPart
	;

domain
	: dotAtom
	| domainLiteral
//| obsDomain
	;

domainLiteral: cfws? LBracket (fws? dtext)* fws? RBracket cfws?;

dtext
	: Exclamation
	| DQuote
	| Hash
	| Dollar
	| Percent
	| Ampersand
	| SQuote
	| LParens
	| RParens
	| Asterisk
	| Plus
	| Comma
	| Minus
	| Period
	| Slash
	| Digit
	| Colon
	| Semicolon
	| Less
	| Equal
	| Greater
	| Question
	| At
	| AlphaUpper
	| Caret
	| Underscore
	| Backtick
	| AlphaLower
	| LCurly
	| Pipe
	| RCurly
	| Tilde
//| obsDtext
	;


// -------------------------
// B.1. Core Rules (RFC5234)
// -------------------------

crlf: CR LF;

wsp: SP | TAB;

vchar
	: Exclamation
	| DQuote
	| Hash
	| Dollar
	| Percent
	| Ampersand
	| SQuote
	| LParens
	| RParens
	| Asterisk
	| Plus
	| Comma
	| Minus
	| Period
	| Slash
	| Digit
	| Colon
	| Semicolon
	| Less
	| Equal
	| Greater
	| Question
	| At
	| AlphaUpper
	| LBracket
	| Backslash
	| RBracket
	| Caret
	| Underscore
	| Backtick
	| AlphaLower
	| LCurly
	| Pipe
	| RCurly
	| Tilde
	;
