lexer grammar AddressLexer;

// Control chars (unprintable)
CRLF: '\r\n';
DEL:  '\u007F';

HTAB:							'\t';     // \u0009
SP:               ' ';      // \u0020

// Printable chars.
Exclamation:      '!';      // \u0021
DQuote:           '"';      // \u0022
Hash:             '#';      // \u0023
Dollar:           '$';      // \u0024
Percent:          '%';      // \u0025
Ampersand:        '&';      // \u0026
SQuote:           '\'';     // \u0027
LParens:          '(';      // \u0028
RParens:          ')';      // \u0029
Asterisk:         '*';      // \u002A
Plus:             '+';      // \u002B
Comma:            ',';      // \u002C
Minus:            '-';      // \u002D
Period:           '.';      // \u002E
Slash:            '/';      // \u002F
Digit:            [0-9];    // \u0030 -- \u0039
Colon:            ':';      // \u003A
Semicolon:        ';';      // \u003B
Less:             '<';      // \u003C
Equal:            '=';      // \u003D
Greater:          '>';      // \u003E
Question:         '?';      // \u003F
At:               '@';      // \u0040
AlphaUpper:       [A-Z];    // \u0041 -- \u005A
LBracket:         '[';      // \u005B
Backslash:        '\\';     // \u005C
RBracket:         ']';      // \u005D
Caret:            '^';      // \u005E
Underscore:       '_';      // \u005F
Backtick:         '`';      // \u0060
AlphaLower:       [a-z];    // \u0061 -- \u007A
LCurly:           '{';      // \u007B
Pipe:             '|';      // \u007C
RCurly:           '}';      // \u007D
Tilde:            '~';      // \u007E


fragment A: 'A'|'a';
fragment B: 'B'|'b';
fragment C: 'C'|'c';
fragment D: 'D'|'d';
fragment E: 'E'|'e';
fragment F: 'F'|'f';
fragment G: 'G'|'g';
fragment H: 'H'|'h';
fragment I: 'I'|'i';
fragment J: 'J'|'j';
fragment K: 'K'|'k';
fragment L: 'L'|'l';
fragment M: 'M'|'m';
fragment N: 'N'|'n';
fragment O: 'O'|'o';
fragment P: 'P'|'p';
fragment Q: 'Q'|'q';
fragment R: 'R'|'r';
fragment S: 'S'|'s';
fragment T: 'T'|'t';
fragment U: 'U'|'u';
fragment V: 'V'|'v';
fragment W: 'W'|'w';
fragment X: 'X'|'x';
fragment Y: 'Y'|'y';
fragment Z: 'Z'|'z';
