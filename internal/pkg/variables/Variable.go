package variables

import (
  "github.com/devandrewgeorge/config-generator/internal/pkg/variables/validators"
)

type Variable struct {
	required    bool
	sensitive   bool
	validators []validators.Validator `yaml:"validate"`
}

func (v Variable) Equal(o Variable) bool {
  if v.required != o.required {
    return false
  }

  if v.sensitive != o.sensitive {
    return false
  }

  if len(v.validators) != len(o.validators) {
    return false
  }

  for i, validator := range v.validators {
    if !validator.Equal(o.validators[i]) {
      return false
    }
  }

  return true
}
