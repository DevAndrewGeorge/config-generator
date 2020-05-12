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

func New() (*Generator) {
  g := &Generator{}

  g.Plugins = map[string]plugins.Plugin{}
  for name, create := range plugins.Plugins {
    g.Plugins[name] = create()
  }

  return g
}

func (g Generator) Equal(o Generator) bool {
  for name, _ := range g.Plugins {
    if !g.Plugins[name].Equal(o.Plugins[name]) {
      return false
    }
  }

  for name, _ := range g.Variables {
    if !g.Variables[name].Equal(o.Variables[name]) {
      return false
    }
  }

  for name, _ := range g.Templates {
    if !g.Templates[name].Equal(o.Templates[name]) {
      return false
    }
  }

  for name, _ := range g.Outputs {
    if !g.Outputs[name].Equal(o.Outputs[name]) {
      return false
    }
  }

  return true
}
