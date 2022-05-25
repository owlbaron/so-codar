package parser

import (
	"strconv"
	"strings"
	"unicode"
)

func isBool(value string) bool {
  return value == "true" || value == "false";
}

func isString(value string) bool {
  if len(value) < 2 {
    return false
  }
  
  first, last := value[0], value[len(value)-1]
  
  return first == '"' && last == '"';
}

func isDouble(value string) bool {  
  is := true

  for _, r := range value {
    if !unicode.IsNumber(r) && r != '.' && r != 'e' && r != '+' && r != '-' {
      is = false
    }
  }
  
  return is && strings.Contains(value, ".");
}

func isInt(value string) bool {
  is := true

  for _, r := range value {
    if !unicode.IsNumber(r) {
      is = false
    }
  }
  
  return is;
}

func isTruthy(literalValue *LiteralValue) bool {
  if literalValue.Type == BOOL {
    return literalValue.Value == "true"
  }

  if literalValue.Type == STRING {
    return literalValue.Value != ""
  }

  if literalValue.Type == INT || literalValue.Type == DOUBLE {
    value, err := strconv.ParseFloat(literalValue.Value, 64)
    if err != nil {
      return false
    }
    
    return value > 0
  }
  
  return false
}