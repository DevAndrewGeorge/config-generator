package parser

import (
  "testing"
  "github.com/devandrewgeorge/config-generator/internal/pkg/plugins"
)

func TestParsePlugin (t *testing.T) {
  var empty_config map[string]interface{}

  t.Run("plugin, invalid", func (t *testing.T) {
    _, e := parsePlugin("blah", "blah", empty_config)
    if e == nil {
      t.Fail()
    }
  })

  t.Run("plugin: console", func(t *testing.T) {
    p, e := parsePlugin("console", "console", empty_config)
    if e != nil {
      t.Fail()
    }

    if _, ok := p.(*plugins.ConsolePlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: environment", func(t *testing.T) {
    p, e := parsePlugin("environment", "environment", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := p.(*plugins.EnvironmentPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: file", func(t *testing.T) {
    p, e := parsePlugin("file", "file", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := p.(*plugins.FilePlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: http", func(t *testing.T) {
    p, e := parsePlugin("http", "http", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := p.(*plugins.HttpPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: kubernetes", func(t *testing.T) {
    p, e := parsePlugin("kubernetes", "kubernetes", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := p.(*plugins.KubernetesPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: multi", func(t *testing.T) {
    p, e := parsePlugin("multi", "multi", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := p.(*plugins.MultiPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: script", func(t *testing.T) {
    p, e := parsePlugin("script", "script", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := p.(*plugins.ScriptPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: static", func(t *testing.T) {
    p, e := parsePlugin("static", "static", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := p.(*plugins.StaticPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: vault", func(t *testing.T) {
    p, e := parsePlugin("vault", "vault", empty_config)

    if e != nil {
      t.Fail()
    }

    if _, ok := p.(*plugins.VaultPlugin); !ok {
      t.Fail()
    }
  })

  t.Run("plugin: configures plugin", func(t *testing.T) {
    config := map[string]interface{}{"file": "hello"}
    expected := plugins.NewEnvironmentPlugin()
    expected.Configure(config)
    p, e := parsePlugin("environment", "environment", config)

    if e != nil {
      t.Fail()
    }

    if p == nil || !p.Equal(expected) {
      t.Fail()
    }
  })
}
