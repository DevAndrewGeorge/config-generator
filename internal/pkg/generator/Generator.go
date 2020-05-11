package generator

import (
  "github.com/devandrewgeorge/config-generator/internal/pkg/plugins"
  "github.com/devandrewgeorge/config-generator/internal/pkg/variables"
  "github.com/devandrewgeorge/config-generator/internal/pkg/templates"
  "github.com/devandrewgeorge/config-generator/internal/pkg/outputs"
)

type Generator struct {
  Plugins map[string]plugins.Plugin
  Variables map[string]variables.Variable
  Templates map[string]templates.Template
  Outputs map[string]outputs.Output
}
