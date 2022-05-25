package parser

import (
	"fmt"

	"github.com/marcobarao/parser/internal/lexer"
	"github.com/marcobarao/parser/internal/tokens"
)

type Parser struct {
	lexer *lexer.Lexer
  memory map[string]*Variable
}

type Variables map[string]*LiteralValue

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer: lexer,
    memory: make(map[string]*Variable),
	}
}

func (p *Parser) skip() error {
  token := p.lexer.Next()

  if token == nil {
    return fmt.Errorf("expected skip a token but found <EOF>\n")
  }

  return nil
}

func (p *Parser) match(tokenType tokens.TokenType, values ...string) (*tokens.Token, error) {
  found := false
  token := p.lexer.Next()

  if token == nil {
    return nil, fmt.Errorf("expected %v but found <EOF>\n", values)
  }

  if len(values) == 0 {
    if token.Type != tokenType {
      fmt.Println(token, tokenType)
      return nil, fmt.Errorf("expected value of type \"%s\" at line %d and column %d\n", tokenType, token.Line, token.Column)
    }

    return token, nil
  }

  for _, value := range values {
    if token.Type == tokenType && token.Value == value {
      found = true
      break;
    }
  }
  if !found {
    return nil, fmt.Errorf("expected %v of type %s and found \"%s\" of type %s at line %d and column %d\n", values, tokenType, token.Value, token.Type, token.Line, token.Column)
  }

  return token, nil
}

func (p *Parser) peek(tokenType tokens.TokenType, values ...string) bool {
  found := false
  token := p.lexer.Peek()

  if token == nil {
    return false
  }

  if len(values) == 0 {
    return token.Type == tokenType
  }

  for _, value := range values {
    if token.Type == tokenType && token.Value == value {
      found = true
      break;
    }
  }
  
  return found
}

func (p *Parser) backTo(token *tokens.Token) error {
  return p.lexer.BackTo(token)
}

func (p *Parser) getLiteralValue() (*LiteralValue, error) {
  token := p.lexer.Peek()
  if token == nil {
    return nil, fmt.Errorf("expected a literal but found <EOF>\n")
  }

  if token.Type != tokens.LITERAL {
    return nil, fmt.Errorf("expected a literal but found %s at line %d and column %d\n", token.Type, token.Line, token.Column)
  }

  if isBool(token.Value) {
    return &LiteralValue{Type: BOOL, Value: token.Value}, nil
  } else if isString(token.Value) {
    return &LiteralValue{Type: STRING, Value: token.Value[1:len(token.Value)-1]}, nil
  } else if isDouble(token.Value) {
    return &LiteralValue{Type: DOUBLE, Value: token.Value}, nil
  } else if isInt(token.Value) {
    return &LiteralValue{Type: INT, Value: token.Value}, nil
  } else {
    return nil, fmt.Errorf("type not found at line %d and column %d\n", token.Line, token.Column)
  }
}

func (p *Parser) Program() error {
  fmt.Println("Starting program")
  var err error
  
  _, err = p.match(tokens.IDENTIFIER, "program")
  if err != nil {
    return err
  }

  err = p.block(0)
  if err != nil {
    return err
  }

  fmt.Println("Ending program")

	return nil
}

func (p *Parser) garbageCollector(level int) {
  for key, value := range p.memory {
    if value.Level >= level {
      delete(p.memory, key)
    }
  }
}

func (p *Parser) block(level int) error {
  var err error

  _, err = p.match(tokens.GROUPER, "{")
  if err != nil {
    return err
  }
  err = p.statementList(level)
  if err != nil {
    return err
  }

  _, err = p.match(tokens.GROUPER, "}")
  if err != nil {
    return err
  }

  p.garbageCollector(level);

	return nil
}

func (p *Parser) statementList(level int) error {
  for !p.peek(tokens.GROUPER, "}") {
    err := p.statement(level)
    if err != nil {
      return err
    }
  }

  return nil
}

func (p *Parser) statement(level int) error {
  var err error
  
  if p.peek(tokens.IDENTIFIER, "if", "while") {
    err = p.structuredStatement(level)
    if err != nil {
      return err
    }
  } else if p.peek(tokens.IDENTIFIER, "var") {
    err = p.variableDeclaration(level)
    if err != nil {
      return err
    }
  } else if p.peek(tokens.IDENTIFIER) {
    err = p.assignmentStatement(level)
    if err != nil {
      return err
    }
  } else {
    token := p.lexer.Next()
    return fmt.Errorf("invalid \"%s\" at line %d and column %d", token.Value, token.Line, token.Column)
  }

  return nil
}

func (p *Parser) variableDeclaration(level int) error {
  var id *tokens.Token
  var literalValue *LiteralValue
	var err error

  _, err = p.match(tokens.IDENTIFIER, "var")
  if err != nil {
    return err
  }

  id, err = p.match(tokens.IDENTIFIER)
  if err != nil {
    return err
  }  

  _, err = p.match(tokens.OPERATOR, "=")
  if err != nil {
    return err
  }

  literalValue, err = p.expression(level)
  if err != nil {
    return err
  }

  _, err = p.match(tokens.GROUPER, ";")
  if err != nil {
    return err
  }

  if val, exists := p.memory[id.Value]; !exists || val.Level >= level {
    p.memory[id.Value] = &Variable{
      Level: level,
      Value: literalValue,
    }

    fmt.Printf("Initializing %s variable with value %s\n", id.Value, literalValue.Value)
  } else {
    return fmt.Errorf("declaration of a variable that already exists at line %d and column %d\n", id.Line, id.Column)
  }
  
  return nil
}

func (p *Parser) assignmentStatement(level int) error {
  var id *tokens.Token
  var literalValue *LiteralValue
	var err error

  id, err = p.match(tokens.IDENTIFIER)
  if err != nil {
    return err
  }  

  _, err = p.match(tokens.OPERATOR, "=")
  if err != nil {
    return err
  }

  literalValue, err = p.expression(level)
  if err != nil {
    return err
  }

  _, err = p.match(tokens.GROUPER, ";")
  if err != nil {
    return err
  }

  if val, exists := p.memory[id.Value]; exists && val.Level <= level {
    p.memory[id.Value] = &Variable{
      Level: val.Level,
      Value: literalValue,
    }

    fmt.Printf("Overriding %s variable with value %s\n", id.Value, literalValue.Value)
  } else {
    return fmt.Errorf("assignment of a variable that does not exists at line %d and column %d\n", id.Line, id.Column)
  }
  
  return nil
}

func (p *Parser) structuredStatement(level int) error {
  var err error
  
  if p.peek(tokens.IDENTIFIER, "if") {
    err = p.ifStatement(level)
    if err != nil {
      return err
    }
  } else if p.peek(tokens.IDENTIFIER, "while") {
    err = p.whileStatement(level)
    if err != nil {
      return err
    }
  }

  return nil
} 

func (p *Parser) whileStatement(level int) error {
  var begin *tokens.Token
  var condition *LiteralValue
  var err error

  begin, err = p.match(tokens.IDENTIFIER, "while")
  if err != nil {
    return err
  }

  condition, err = p.expression(level)
  if err != nil {
    return err
  }

  for isTruthy(condition) {
    fmt.Println("Executing loop block")
    err = p.block(level + 1)
    if err != nil {
      return err
    }

    p.backTo(begin)
    err = p.skip()
    if err != nil {
      return err
    }
    
    condition, err = p.expression(level)
    if err != nil {
      return err
    }
  }

  p.ignoreBlock()

  return nil
}

func (p *Parser) ignoreBlock() error {
  level := 0
  if p.peek(tokens.GROUPER, "{") {
    level = 1
  }
  
  for level > 0 {
    err := p.skip()
    if err != nil {
      return err
    }
    
    if p.peek(tokens.GROUPER, "{") {
      level += 1
    }

    if p.peek(tokens.GROUPER, "}") {
      level -= 1
    }
  }

  err := p.skip()
  if err != nil {
    return err
  }

  return nil
}

func (p *Parser) ifStatement(level int) error {
  var condition *LiteralValue
	var err error

  _, err = p.match(tokens.IDENTIFIER, "if")
  if err != nil {
    return err
  }

  condition, err = p.expression(level)
  if err != nil {
    return err
  }

  if isTruthy(condition) {
    fmt.Println("Executing truthy block")
    
    err = p.block(level + 1)
    if err != nil {
      return err
    } 

    if p.peek(tokens.IDENTIFIER, "else") {
      fmt.Println("Ignoring falsy block")
      
      _, err = p.match(tokens.IDENTIFIER, "else")
      if err != nil {
        return err
      }
      
      err = p.ignoreBlock();
      if err != nil {
        return err
      }
    }
  } else {
    fmt.Println("Ignoring truthy block")
    err = p.ignoreBlock();
    if err != nil {
      return err
    }
    
    if p.peek(tokens.IDENTIFIER, "else") {
      fmt.Println("Executing falsy block")
      _, err = p.match(tokens.IDENTIFIER, "else")
      if err != nil {
        return err
      }
  
      err = p.block(level + 1)
      if err != nil {
        return err
      }
    }
  }

  return nil
}

var relationalOperators = []string{"==", "!=", "<=", "<", ">", ">="}
func (p *Parser) expression(level int) (*LiteralValue, error) {
  var op *tokens.Token
  var left, right *LiteralValue
	var err error

  left, err = p.simpleExpression(level)
  if err != nil {
    return nil, err
  }
  
  for p.peek(tokens.OPERATOR, relationalOperators...) {
    op, err = p.match(tokens.OPERATOR, relationalOperators...)
    if err != nil {
      return nil, err
    }

    right, err = p.simpleExpression(level)
    if err != nil {
      return nil, err
    }

    fmt.Printf("Doing operation %s %s %s\n", left.Value, op.Value, right.Value)
    
    left, err = operations[left.Type][op.Value](left, right);
    if err != nil {
      return nil, fmt.Errorf("%w at line %d and column %d\n", err, op.Line, op.Column)
    }
  }
  
  return left, nil
} 

// simple-expression ::= (sign)? term (lower-precedence-operator term)*
var signOperators = []string{"+", "-"}
var lowerPrecedenceOperators = []string{"+", "-", "||"}
func (p *Parser) simpleExpression(level int) (*LiteralValue, error) {
  var op *tokens.Token
  var left, right *LiteralValue
  var err error
  
	if p.peek(tokens.OPERATOR, signOperators...) {
    op, err = p.match(tokens.OPERATOR, signOperators...)
    if err != nil {
      return nil, err
    }
  }

  left, err = p.term(level)
  if err != nil {
    return nil, err
  }

  if op != nil {
    fmt.Printf("Doing operation %s %s %s\n", "0", op.Value, left.Value)
    left, err = operations[left.Type][op.Value](&LiteralValue{ Type: left.Type, Value: "0" }, left)
    if err != nil {
      return nil, fmt.Errorf("%w at line %d and column %d\n", err, op.Line, op.Column)
    }
  }

  for p.peek(tokens.OPERATOR, lowerPrecedenceOperators...) {
    op, err = p.match(tokens.OPERATOR, lowerPrecedenceOperators...)
    if err != nil {
      return nil, err
    }

    right, err = p.term(level)
    if err != nil {
      return nil, err
    }

    fmt.Printf("Doing operation %s %s %s\n", left.Value, op.Value, right.Value)
    
    left, err = operations[left.Type][op.Value](left, right);
    if err != nil {
      return nil, fmt.Errorf("%w at line %d and column %d\n", err, op.Line, op.Column)
    }
  }

  return left, nil
} 


// term ::= factor (higher-precendence-operator factor)*
var higherPrecedenceOperators = []string{"*", "/", "%", "&&"}
func (p *Parser) term(level int) (*LiteralValue, error) {
  var op *tokens.Token
  var left, right *LiteralValue
	var err error

  left, err = p.factor(level)
  if err != nil {
    return nil, err
  }

  for p.peek(tokens.OPERATOR, higherPrecedenceOperators...) {
    op, err = p.match(tokens.OPERATOR, higherPrecedenceOperators...)
    if err != nil {
      return nil, err
    }

    right, err = p.factor(level)
    if err != nil {
      return nil, err
    }

    fmt.Printf("Doing operation %s %s %s\n", left.Value, op.Value, right.Value)

    left, err = operations[left.Type][op.Value](left, right);
    if err != nil {
      return nil, fmt.Errorf("%w at line %d and column %d\n", err, op.Line, op.Column)
    }
  }
  
  return left, nil
} 


func (p *Parser) factor(level int) (*LiteralValue, error) {
  var literalValue *LiteralValue
  var err error
  
	if p.peek(tokens.GROUPER, "(") {
    _, err = p.match(tokens.GROUPER, "(")
    if err != nil {
      return nil, err
    }

    literalValue, err = p.expression(level)
    if err != nil {
      return nil, err
    }

    _, err = p.match(tokens.GROUPER, ")")
    if err != nil {
      return nil, err
    }
  } else if p.peek(tokens.LITERAL) {
    literalValue, err = p.literal()
    if err != nil {
      return nil, err
    }
  } else if p.peek(tokens.IDENTIFIER) {
    var id *tokens.Token
    id, err = p.match(tokens.IDENTIFIER)
    if err != nil {
      return nil, err
    }

    if variable, exists := p.memory[id.Value]; exists && variable.Level <= level {
      literalValue = variable.Value
    } else {
      return nil, fmt.Errorf("use of a variable that does not exists at line %d and column %d\n", id.Line, id.Column)
    }
  }

  return literalValue, nil
}

func (p *Parser) literal() (*LiteralValue, error) {
  var literalValue *LiteralValue
  var err error

  if p.peek(tokens.LITERAL) {
    literalValue, err = p.getLiteralValue();
    if err != nil {
      return nil, err
    }

    
    _, err = p.match(tokens.LITERAL)
    if err != nil {
      return nil, err
    }
  }
  
  return literalValue, nil
}
