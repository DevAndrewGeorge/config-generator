package parser

import (
  "testing"
  "github.com/devandrewgeorge/config-generator/internal/pkg/plugins"
)

func TestParsePlugin (t *testing.T) {
  test_string := "test"
  t.Run("named plugins map is nil", func(t *testing.T) {
    plugins, err := parsePlugins(nil);
    if err != nil || len(plugins) != 0 {
      t.Fail()
    }
  })

  t.Run("named plugins map is empty", func(t *testing.T) {
    plugins, error := parsePlugins(map[string]*namedPluginConfig{})
    if error != nil || len(plugins) != 0 {
      t.Fail()
    }
  })

  t.Run("named plugins map contains a properly configured named plugin", func(t *testing.T) {
    input := map[string]*namedPluginConfig{
      "test": &namedPluginConfig{
        Environment: &environmentPluginConfig{ File: &test_string },
      },
    }
    expected := &plugins.EnvironmentPlugin{}
    expected.Configure("test", map[string]interface{}{"File": &test_string})

    plugins, error := parsePlugins(input)
    if error != nil { t.Fail() }
    if !plugins["test"].Equal(expected) { t.Fail() }
  })

  t.Run("named plugins map contains multiple properly configured named plugins", func(t *testing.T) {
    input := map[string]*namedPluginConfig{
      "test1": &namedPluginConfig{
        Environment: &environmentPluginConfig{ File: &test_string },
      },
      "test2": &namedPluginConfig{
        Console: &consolePluginConfig{},
      },
    }

    plugins, error := parsePlugins(input)
    if error != nil { t.Fail() }
    if _, ok := plugins["test1"]; !ok { t.Fail() }
    if _, ok := plugins["test2"]; !ok { t.Fail() }
  })

  t.Run("named plugin contains no actual plugin configuration", func(t *testing.T) {
    input := map[string]*namedPluginConfig{
      "test": &namedPluginConfig{},
    }

    if _, error := parsePlugins(input); error == nil { t.Fail() }
  })

  t.Run("actual plugin configuration has multiple configurations", func (t *testing.T) {
    input := map[string]*namedPluginConfig{
      "test": &namedPluginConfig{
        Environment: &environmentPluginConfig{},
        Console: &consolePluginConfig{},
      },
    }

    if _, error := parsePlugins(input); error == nil {
      t.Fail()
    }
  })

  t.Run("actual plugin configuration map is empty", func (t *testing.T) {
    input := map[string]*namedPluginConfig{
      "test": &namedPluginConfig{
        Environment: &environmentPluginConfig{},
      },
    }
    expected := &plugins.EnvironmentPlugin{}
    expected.Configure("test", nil)

    plugins, error := parsePlugins(input)
    if error != nil { t.Error("error returned") }
    if !plugins["test"].Equal(expected) { t.Error(plugins["test"], expected) }
  })
}
