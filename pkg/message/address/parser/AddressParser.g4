// Copyright (c) 2020 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

parser grammar AddressParser;

options { tokenVocab=AddressLexer; }


// -------------------
// 3.2. Lexical tokens
// -------------------

quotedChar: vchar | wsp;

quotedPair
	: Backslash quotedChar
	| obsQP
	;

fws 
	: (wsp* crlf)? wsp+
	| obsFWS
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
	| obsCtext
	| UTF8NonAscii
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
	| UTF8NonAscii
	;

atom: atext+;

dotAtom: atext+ (Period atext+)*;

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
	| obsQtext
	| UTF8NonAscii
	;

quotedContent
	: qtext
	| quotedPair
	;

quotedValue: (fws? quotedContent)*;

quotedString: DQuote quotedValue fws? DQuote;

word
	: cfws? encodedWord cfws?
	| cfws? atom cfws?
	| cfws? quotedString cfws?
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
	: cfws? Less addrSpec? Greater cfws?
	| obsAngleAddr
	;

group: displayName Colon groupList? Semicolon cfws?;

displayName
	: word+
	| obsPhrase
	;

mailboxList
	: mailbox (Comma mailbox)*
	| obsMboxList
	;

addressList
	: address (Comma address)*
	| obsAddrList
	;

groupList
	: mailboxList
	| cfws
	| obsGroupList
	;

addrSpec: localPart At domain;

localPart
	: cfws? dotAtom cfws?
	| cfws? quotedString cfws?
	| obsLocalPart
	;

domain
	: cfws? dotAtom cfws? 
	| cfws? domainLiteral cfws?
	| cfws? obsDomain cfws?
	;

domainLiteral: LBracket (fws? dtext)* fws? RBracket;

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
	| UTF8NonAscii
	;


// ----------------------------------
// 4.1. Miscellaneous Obsolete Tokens
// ----------------------------------

obsNoWSCTL
	: U_01_08
	| U_0B
	| U_0C
	| U_0E_1F
	| Delete
	;

obsPhrase: word (word | Period | cfws)*;

obsCtext: obsNoWSCTL;

obsQtext: obsNoWSCTL;

obsQP: Backslash (U_00 | obsNoWSCTL | LF | CR);


// ------------------------
// 4.2. Obsolete Addressing
// ------------------------

obsFWS: wsp+ (crlf wsp+);


// ------------------------
// 4.4. Obsolete Addressing
// ------------------------

obsAngleAddr: cfws? Less obsRoute addrSpec Greater cfws?;

obsRoute: obsDomainList Colon;

obsDomainList: (cfws | Comma)* At domain (Comma cfws? (At domain)?)*;

obsMboxList: (cfws? Comma)* mailbox (Comma (mailbox | cfws)?)*;

obsAddrList: (cfws? Comma)* address (Comma (address | cfws)?)*;

obsGroupList: (cfws? Comma)+ cfws?;

obsLocalPart: word (Period word)*;

obsDomain: atom (Period atom)*;


// ------------------------------------
// 2. Syntax of encoded-words (RFC2047)
// ------------------------------------

encodedWord: Equal Question charset Question encoding Question encodedText Question Equal;

charset: token;

encoding: token;

token: tokenChar+;

tokenChar
	: Exclamation
	| Hash
	| Dollar
	| Percent
	| Ampersand
	| SQuote
	| Asterisk
	| Plus
	| Minus
	| Digit
	| AlphaUpper
	| Backslash
	| Caret
	| Underscore
	| Backtick
	| AlphaLower
	| LCurly
	| Pipe
	| RCurly
	| Tilde
	;

encodedText: encodedChar+;

encodedChar
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
	| UTF8NonAscii
	;
