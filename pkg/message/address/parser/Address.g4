parser grammar Address;

options { tokenVocab=AddressLexer; }


// 3.2.1. Quoted characters

// quoted-pair     =   ("\" (VCHAR / WSP)) / obs-qp
quotedPair
	: Backslash (vchar | wsp)
//| obsQP
	;


// 3.2.2. Folding White Space and Comments

// FWS             =   ([*WSP CRLF] 1*WSP) /  obs-FWS
fws 
	: (wsp* CRLF) wsp+
//| obsFWS
	;

// ctext           =   ascii 33-39 / ascii 42-91 / ascii 93-126 / obs-ctext
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


// 3.2.3. Atom

// atext         =   ALPHA / DIGIT /    ; Printable US-ASCII
//                   "!" / "#" /        ;  characters not including
//                   "$" / "%" /        ;  specials.  Used for atoms.
//                   "&" / "'" /
//                   "*" / "+" /
//                   "-" / "/" /
//                   "=" / "?" /
//                   "^" / "_" /
//                   "`" / "{" /
//                   "|" / "}" /
//                   "~"
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

// atom            =   [CFWS] 1*atext [CFWS]
atom: cfws? atext+ cfws?;

// dot-atom-text   =   1*atext *("." 1*atext)
dotAtomText: atext+ (Period atext+)*;

// dot-atom        =   [CFWS] dot-atom-text [CFWS]
dotAtom: cfws? dotAtomText cfws?;


// 3.2.4. Quoted Strings

// qtext           =   %d33 / d35-91 / d93-126 / obs-qtext
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

// qcontent        =   qtext / quoted-pair
qcontent
	: qtext
	| quotedPair
	;

// quoted-string   =   [CFWS] DQUOTE *([FWS] qcontent) [FWS] DQUOTE [CFWS]
quotedString: cfws? DQuote (fws? qcontent)* fws? DQuote cfws?;


// 3.2.5. Miscellaneous Tokens

// word            =   atom / quoted-string
word
	: atom
	| quotedString
	;

// phrase          =   1*word / obs-phrase
phrase
	: word+
//| obsPhrase
	;

// unstructured    =   (*([FWS] VCHAR) *WSP) / obs-unstruct
unstructured
	: (fws? vchar)* wsp*
//| obsUnstruct
	;


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
	: cfws? Less addrSpec Greater cfws?
//| obsAngleAddr
	;

// group           =   display-name ":" [group-list] ";" [CFWS]
group: displayName Colon groupList? Semicolon cfws?;

// display-name    =   phrase
displayName: phrase;

// mailbox-list    =   (mailbox *("," mailbox)) / obs-mbox-list
mailboxList
	: mailbox (Comma mailbox)*
//| obsMboxList
	;

// address-list    =   (address *("," address)) / obs-addr-list
addressList
	: address (Comma address)*
//| obsAddrList
	;

// group-list      =   mailbox-list / CFWS / obs-group-list
groupList
	: mailboxList
	| cfws
//| obsGroupList
	;


// 3.4.1. Addr-Spec Specification

// addr-spec       =   local-part "@" domain
addrSpec: localPart At domain;

// local-part      =   dot-atom / quoted-string / obs-local-part
localPart
	: dotAtom 
	| quotedString
//| obsLocalPart
	;

// domain          =   dot-atom / domain-literal / obs-domain
domain
	: dotAtom
	| domainLiteral
//| obsDomain
	;

// domain-literal  =   [CFWS] "[" *([FWS] dtext) [FWS] "]" [CFWS]
domainLiteral: cfws? LBracket (fws? dtext)* RBracket cfws?;

// dtext           =   ascii 33-90 / ascii 94-126 / obs-dtext
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


// B.1. Core Rules (RFC5234)

// WSP            =  SP / HTAB
wsp: SP | HTAB;

// VCHAR          =  %x21-7E
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
