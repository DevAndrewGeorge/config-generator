package parser

import (
  "testing"
  "github.com/devandrewgeorge/config-generator/internal/pkg/plugins"
)

func TestParsePlugin (t *testing.T) {
  var empty_config map[string]interface{}
  t.Run("plugin: console", func(t *testing.T) {
    g, e := parsePlugin("console", empty_config)
    if e != nil {
      t.Fail()
    }

    if _, ok := g.(plugins.ConsolePlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: environment", func(t *testing.T) {
    g, e := parsePlugin("environment", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := g.(plugins.EnvironmentPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: file", func(t *testing.T) {
    g, e := parsePlugin("file", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := g.(plugins.FilePlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: http", func(t *testing.T) {
    g, e := parsePlugin("http", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := g.(plugins.HttpPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: kubernetes", func(t *testing.T) {
    g, e := parsePlugin("kubernetes", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := g.(plugins.KubernetesPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: multi", func(t *testing.T) {
    g, e := parsePlugin("multi", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := g.(plugins.MultiPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: script", func(t *testing.T) {
    g, e := parsePlugin("script", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := g.(plugins.ScriptPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: static", func(t *testing.T) {
    g, e := parsePlugin("static", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := g.(plugins.StaticPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: vault", func(t *testing.T) {
    g, e := parsePlugin("vault", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := g.(plugins.VaultPlugin); !ok {
      t.Fail()
    }
  })
}
