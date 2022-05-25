package automataerrors

import (
  "fmt"
)

type StateNotFoundError struct {
  State uint
}

func (e StateNotFoundError) Error() string {
  return fmt.Sprintf("state %d not found", e.State)
}