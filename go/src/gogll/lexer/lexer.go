// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"gogll/token"
)

const (
	NoState    = -1
	NumStates  = 70
	NumSymbols = 71
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '''
1: '''
2: '"'
3: '"'
4: 'p'
5: 'a'
6: 'c'
7: 'k'
8: 'a'
9: 'g'
10: 'e'
11: ':'
12: ';'
13: '|'
14: 'e'
15: 'm'
16: 'p'
17: 't'
18: 'y'
19: 'A'
20: 'l'
21: 't'
22: 'a'
23: 'n'
24: 'y'
25: 'l'
26: 'e'
27: 't'
28: 't'
29: 'e'
30: 'r'
31: 'n'
32: 'u'
33: 'm'
34: 'b'
35: 'e'
36: 'r'
37: 'u'
38: 'p'
39: 'c'
40: 'a'
41: 's'
42: 'e'
43: 'l'
44: 'o'
45: 'w'
46: 'c'
47: 'a'
48: 's'
49: 'e'
50: 'n'
51: 'o'
52: 't'
53: 's'
54: 'p'
55: 'a'
56: 'c'
57: 'e'
58: '\'
59: '"'
60: 'n'
61: '''
62: '_'
63: ' '
64: '\t'
65: '\n'
66: '\r'
67: 'A'-'Z'
68: 'a'-'z'
69: '0'-'9'
70: .
*/
