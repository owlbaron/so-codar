package parser

import (
  "fmt"
  "strconv"
)

type operation func(*LiteralValue, *LiteralValue) (*LiteralValue, error)

var operations = map[VariableType]map[string]operation{
  INT: numOperations,
  DOUBLE: numOperations,
  BOOL: boolOperations,
  STRING: stringOperations,
}

var stringOperations = map[string]operation {
  "+": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    if right.Type != STRING {
      return nil, fmt.Errorf("invalid string operation")
    }

    return &LiteralValue{
      Type: STRING,
      Value: left.Value + right.Value,
    }, nil
  },
  "==": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    if right.Type != STRING {
      return nil, fmt.Errorf("invalid string operation")
    }

    value := "false"

    if left.Value == right.Value {
      value = "true"
    }

    return &LiteralValue{
      Type: BOOL,
      Value: value,
    }, nil
  },
  "!=": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    if right.Type != STRING {
      return nil, fmt.Errorf("invalid string operation")
    }

    value := "false"

    if left.Value != right.Value {
      value = "true"
    }

    return &LiteralValue{
      Type: BOOL,
      Value: value,
    }, nil
  },
}

var boolOperations = map[string]operation {
  "&&": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    if right.Type != BOOL {
      return nil, fmt.Errorf("invalid bool operation")
    }

    value := "false"

    if left.Value == "true" && right.Value == "true" {
      value = "true"
    }

    return &LiteralValue{
      Type: BOOL,
      Value: value,
    }, nil
  },
  "||": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    if right.Type != BOOL {
      return nil, fmt.Errorf("invalid bool operation")
    }

    value := "false"

    if left.Value == "true" || right.Value == "true" {
      value = "true"
    }

    return &LiteralValue{
      Type: BOOL,
      Value: value,
    }, nil
  },
}

var numOperations = map[string]operation{
  "+": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var v string
    var leftValue, rightValue float64
    var err error
    
    leftValue, err = strconv.ParseFloat(left.Value, 64)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.ParseFloat(right.Value, 64)
    if err != nil {
      return nil, err
    }

    t := INT;
    if left.Type == DOUBLE || right.Type == DOUBLE {
      t = DOUBLE;
    }

    if t == INT {
      v = strconv.Itoa(int(leftValue + rightValue))
    } else {
      v = fmt.Sprintf("%f", leftValue + rightValue)
    }
    
    return &LiteralValue{
      Type: t,
      Value: v,
    }, nil
  },
  "-": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var v string
    var leftValue, rightValue float64
    var err error
    
    leftValue, err = strconv.ParseFloat(left.Value, 64)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.ParseFloat(right.Value, 64)
    if err != nil {
      return nil, err
    }

    t := INT;
    if left.Type == DOUBLE || right.Type == DOUBLE {
      t = DOUBLE;
    }

    if t == INT {
      v = strconv.Itoa(int(leftValue - rightValue))
    } else {
      v = fmt.Sprintf("%f", leftValue - rightValue)
    }
    
    return &LiteralValue{
      Type: t,
      Value: v,
    }, nil
  },
  "*": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var v string
    var leftValue, rightValue float64
    var err error
    
    leftValue, err = strconv.ParseFloat(left.Value, 64)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.ParseFloat(right.Value, 64)
    if err != nil {
      return nil, err
    }

    t := INT;
    if left.Type == DOUBLE || right.Type == DOUBLE {
      t = DOUBLE;
    }

    if t == INT {
      v = strconv.Itoa(int(leftValue * rightValue))
    } else {
      v = fmt.Sprintf("%f", leftValue * rightValue)
    }
    
    return &LiteralValue{
      Type: t,
      Value: v,
    }, nil
  },
  "/": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var v string
    var leftValue, rightValue float64
    var err error
    
    leftValue, err = strconv.ParseFloat(left.Value, 64)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.ParseFloat(right.Value, 64)
    if err != nil {
      return nil, err
    }

    t := INT;
    if left.Type == DOUBLE || right.Type == DOUBLE {
      t = DOUBLE;
    }

    if t == INT {
      v = strconv.Itoa(int(leftValue / rightValue))
    } else {
      v = fmt.Sprintf("%f", leftValue / rightValue)
    }
    
    return &LiteralValue{
      Type: t,
      Value: v,
    }, nil
  },
  "%": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var leftValue, rightValue int
    var err error
    
    leftValue, err = strconv.Atoi(left.Value)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.Atoi(right.Value)
    if err != nil {
      return nil, err
    }
    
    return &LiteralValue{
      Type: INT,
      Value: strconv.Itoa(leftValue % rightValue),
    }, nil
  },
  "==": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var leftValue, rightValue float64
    var err error
    
    leftValue, err = strconv.ParseFloat(left.Value, 64)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.ParseFloat(right.Value, 64)
    if err != nil {
      return nil, err
    }

    value := "false"

    if leftValue == rightValue {
      value = "true"
    }
    
    return &LiteralValue{
      Type: BOOL,
      Value: value,
    }, nil
  },
  "!=": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var leftValue, rightValue float64
    var err error
    
    leftValue, err = strconv.ParseFloat(left.Value, 64)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.ParseFloat(right.Value, 64)
    if err != nil {
      return nil, err
    }

    value := "false"

    if leftValue != rightValue {
      value = "true"
    }
    
    return &LiteralValue{
      Type: BOOL,
      Value: value,
    }, nil
  },
  ">=": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var leftValue, rightValue float64
    var err error
    
    leftValue, err = strconv.ParseFloat(left.Value, 64)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.ParseFloat(right.Value, 64)
    if err != nil {
      return nil, err
    }

    value := "false"

    if leftValue >= rightValue {
      value = "true"
    }
    
    return &LiteralValue{
      Type: BOOL,
      Value: value,
    }, nil
  },
  "<=": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var leftValue, rightValue float64
    var err error
    
    leftValue, err = strconv.ParseFloat(left.Value, 64)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.ParseFloat(right.Value, 64)
    if err != nil {
      return nil, err
    }

    value := "false"

    if leftValue <= rightValue {
      value = "true"
    }
    
    return &LiteralValue{
      Type: BOOL,
      Value: value,
    }, nil
  },
  "<": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var leftValue, rightValue float64
    var err error
    
    leftValue, err = strconv.ParseFloat(left.Value, 64)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.ParseFloat(right.Value, 64)
    if err != nil {
      return nil, err
    }

    value := "false"

    if leftValue < rightValue {
      value = "true"
    }
    
    return &LiteralValue{
      Type: BOOL,
      Value: value,
    }, nil
  },
  ">": func (left *LiteralValue, right *LiteralValue) (*LiteralValue, error) {
    var leftValue, rightValue float64
    var err error
    
    leftValue, err = strconv.ParseFloat(left.Value, 64)
    if err != nil {
      return nil, err
    }
    
    rightValue, err = strconv.ParseFloat(right.Value, 64)
    if err != nil {
      return nil, err
    }

    value := "false"

    if leftValue > rightValue {
      value = "true"
    }
    
    return &LiteralValue{
      Type: BOOL,
      Value: value,
    }, nil
  },
}


