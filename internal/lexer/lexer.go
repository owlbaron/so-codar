package lexer

import (
	"fmt"
	"strings"

	"github.com/marcobarao/parser/internal/automata/tokens/grouper"
	"github.com/marcobarao/parser/internal/automata/tokens/identifier"
	"github.com/marcobarao/parser/internal/automata/tokens/literal"
	"github.com/marcobarao/parser/internal/automata/tokens/nl"
	"github.com/marcobarao/parser/internal/automata/tokens/operator"
	"github.com/marcobarao/parser/internal/automata/tokens/ws"
	"github.com/marcobarao/parser/internal/tokens"

	. "github.com/marcobarao/parser/internal/automata"
	. "github.com/marcobarao/parser/internal/tokens"
)

type result struct {
	Accepted bool
	Type     TokenType
	Value    string
}

type testInput struct {
	Automata *Automata
	Type     TokenType
}

type Lexer struct {
	code string
	ptr  uint

	line   uint
	column uint
}

func NewLexer(code string) *Lexer {
	return &Lexer{
		code: code,
		ptr:  0,

		line:   1,
		column: 1,
	}
}

func (l *Lexer) skip() {
  resultChannel := make(chan result)
  found := false
  
  tests := []testInput{
		{Automata: ws.NewWSAutomata(), Type: WS},
		{Automata: nl.NewNLAutomata(), Type: NL},
	}

  for _, test := range tests {
		go l.TestAutomata(resultChannel, test.Automata, test.Type)
	}

  for i := 0; i < len(tests); i++ {
		result := <-resultChannel

		if result.Accepted {
			l.ptr += uint(len(result.Value))
			if result.Type == NL {
				l.line++
				l.column = 1
			} else {
				l.column += uint(len(result.Value))
			}
      found = true
    }
  }

  if found {
    l.skip()
  }
}

func (l *Lexer) BackTo(token *tokens.Token) error {
  if token == nil {
    return fmt.Errorf("token is null\n")
  }
  
  l.ptr = token.Index;
  l.line = token.Line;
  l.column = token.Column;

  return nil
}

var keyword = map[string]bool{
  "true": true,
  "false": true,
}

func (l *Lexer) Peek() *Token {
  resultChannel := make(chan result)

  l.skip()

	tests := []testInput{
		{Automata: grouper.NewGrouperAutomata(), Type: GROUPER},
		{Automata: literal.NewLiteralAutomata(), Type: LITERAL},
		{Automata: operator.NewOperatorAutomata(), Type: OPERATOR},
		{Automata: identifier.NewIdentifierAutomata(), Type: IDENTIFIER},
	}

	for _, test := range tests {
		go l.TestAutomata(resultChannel, test.Automata, test.Type)
	}

	for i := 0; i < len(tests); i++ {
		result := <-resultChannel

    if _, isKeyword := keyword[result.Value]; result.Accepted && result.Type == tokens.IDENTIFIER && isKeyword {
      continue
    }

		if result.Accepted {
           
			return &Token{
				Type:   result.Type,
				Value:  result.Value,
        Index:  l.ptr,
				Line:   l.line,
				Column: l.column,
			}
    }
  }

  if l.ptr < uint(len(l.code)-1) {
		token := &Token{
			Type:   UNKNOWN,
			Value:  string(l.code[l.ptr]),
      Index:  l.ptr,
			Line:   l.line,
			Column: l.column,
		}

		return token
	}

  return nil
}

func (l *Lexer) Next() *Token {
	resultChannel := make(chan result)

  l.skip()

	tests := []testInput{
		{Automata: grouper.NewGrouperAutomata(), Type: GROUPER},
		{Automata: literal.NewLiteralAutomata(), Type: LITERAL},
		{Automata: operator.NewOperatorAutomata(), Type: OPERATOR},
		{Automata: identifier.NewIdentifierAutomata(), Type: IDENTIFIER},
	}

	for _, test := range tests {
		go l.TestAutomata(resultChannel, test.Automata, test.Type)
	}

	for i := 0; i < len(tests); i++ {
		result := <-resultChannel

    if _, isKeyword := keyword[result.Value]; result.Accepted && result.Type == tokens.IDENTIFIER && isKeyword {
      continue
    }

		if result.Accepted {
			token := &Token{
				Type:   result.Type,
				Value:  result.Value,
        Index:  l.ptr,
				Line:   l.line,
				Column: l.column,
			}

			l.ptr += uint(len(result.Value))
			l.column += uint(len(result.Value))
      
			return token
		}
	}

	if l.ptr < uint(len(l.code)-1) {
		token := &Token{
			Type:   UNKNOWN,
			Value:  string(l.code[l.ptr]),
      Index:  l.ptr,
			Line:   l.line,
			Column: l.column,
		}

		l.column++
		l.ptr++

		return token
	}

	return nil
}

func (l *Lexer) TestAutomata(channel chan result, automata *Automata, tokenType TokenType) {
	var sb strings.Builder
	ptr := int(l.ptr)

	for ptr < len(l.code) {
		char := rune(l.code[ptr])

		err := automata.Input(char)

		if err != nil {
			break
		}

		sb.WriteRune(char)
		ptr++
	}

	channel <- result{
		Accepted: automata.Accepted(),
		Type:     tokenType,
		Value:    sb.String(),
	}
}
