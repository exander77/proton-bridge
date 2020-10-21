parser grammar Address;

options { tokenVocab=AddressLexer; }

// 3.4. Address Specification

// address         =   mailbox / group
address
	: mailbox
	| group
	;

// mailbox         =   name-addr / addr-spec
mailbox
	: nameAddr
	| addrSpec
	;

// name-addr       =   [display-name] angle-addr
nameAddr: displayName? angleAddr;

// angle-addr      =   [CFWS] "<" addr-spec ">" [CFWS] / obs-angle-addr
angleAddr
	: cfws? LAngle addrSpec RAngle cfws?
	| obsAngleAddr
	;

// group           =   display-name ":" [group-list] ";" [CFWS]
group: displayName Colon groupList? Semicolon cfws?;

// display-name    =   phrase
displayName: phrase;

// mailbox-list    =   (mailbox *("," mailbox)) / obs-mbox-list
mailboxList
	: mailbox (Comma mailbox)*
	| obsMboxList
	;

// address-list    =   (address *("," address)) / obs-addr-list
addressList
	: address (Comma address)*
	| obsAddrList
	;

// group-list      =   mailbox-list / CFWS / obs-group-list
groupList
	: mailboxList
	| cfws
	| obsGroupList
	;


// 3.4.1. Addr-Spec Specification

// addr-spec       =   local-part "@" domain
addrSpec: localPart At domain;

// local-part      =   dot-atom / quoted-string / obs-local-part
localPart
	: dotAtom 
	| quotedString
	| obsLocalPart
	;

// domain          =   dot-atom / domain-literal / obs-domain
domain
	: dotAtom
	| domainLiteral
	| obsDomain
	;

// domain-literal  =   [CFWS] "[" *([FWS] dtext) [FWS] "]" [CFWS]
domainLiteral: cfws? LBracket (fws? dtext)* RBracket cfws?;

// dtext           =   ascii 33-90 / ascii 94-126 / obs-dtext
/*
dtext
	: ascii 33-90  // TODO
	| ascii 94-126 // TODO
	| obsDtext
	;
*/


// 3.2.1. Quoted characters

// quoted-pair     =   ("\" (VCHAR / WSP)) / obs-qp
quotedPair
	: Backslash (vchar | wsp)
	| obsQP
	;


// 3.2.2. Folding White Space and Comments

// FWS             =   ([*WSP CRLF] 1*WSP) /  obs-FWS
fws 
	: (WSP* CRLF) WSP+
	| obsFWS
	;

// ctext           =   ascii 33-39 / ascii 42-91 / ascii 93-126 / obs-ctext
/*
ctext
	: ascii 33-39  // TODO
	| aasci 42-91  // TODO
	| ascii 93-126 // TODO
	| obsCtext
	;
*/

// ccontent        =   ctext / quoted-pair / comment
ccontent
	: ctext
	| quotedPair
	| comment
	;

// comment         =   "(" *([FWS] ccontent) [FWS] ")"
comment: LParens (fws? ccontent)* fws? RParens;

// CFWS            =   (1*([FWS] comment) [FWS]) / FWS
cfws
	: (fws? comment)+ fws?
	| fws
	;

// vchar: hex 21-7E; // TODO
