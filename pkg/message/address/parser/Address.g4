parser grammar Address;

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
		: CFWS? LAngle addrSpec RAngle CFWS?
		| obsAngleAddr
		;

// group           =   display-name ":" [group-list] ";" [CFWS]
group: displayName Colon groupList? Semicolon CFWS?;

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
		| CFWS
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
domainLiteral: CFWS? LBracket (FWS? dtext)* RBracket CFWS?;

// dtext           =   ascii 33-90 / ascii 94-126 / obs-dtext
dtext
		: ascii 33-90  // TODO
		| ascii 94-126 // TODO
		| obsDtext
		;
