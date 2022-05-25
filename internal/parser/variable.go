package parser

type VariableType string

const (
  INT VariableType = "INT"
  DOUBLE VariableType = "DOUBLE"
  STRING VariableType = "STRING"
  BOOL VariableType = "BOOL"
)

type LiteralValue struct {
	Type  VariableType
	Value string
}

type Variable struct {
  Level int
  Value *LiteralValue
}
