package variables

import (
  "github.com/devandrewgeorge/config-generator/internal/pkg/variables/validators"
)

type Variable struct {
	required    bool
	sensitive   bool
	validations []validators.Validator `yaml:"validate"`
}
