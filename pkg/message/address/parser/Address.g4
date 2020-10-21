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

