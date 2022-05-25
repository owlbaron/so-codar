package automataerrors

import (
  "fmt"
)

type TransitionNotFoundError struct {
  Input rune
  State uint
}

func (e TransitionNotFoundError) Error() string {
  return fmt.Sprintf("transition with rune %c for state %d not found", 
e.Input, e.State)
}