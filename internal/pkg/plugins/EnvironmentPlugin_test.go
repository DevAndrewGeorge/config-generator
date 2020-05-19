package plugins

import (
  "testing"
)

func TestNewEnvironmentPlugin(t *testing.T) {
  if _, ok := NewEnvironmentPlugin().(*EnvironmentPlugin); !ok {
    t.Fail()
  }
}

func TestEnvironmentPluginConfigure(t *testing.T) {
  t.Run("configuration is nil", func(t *testing.T) {
    e := NewEnvironmentPlugin().(*EnvironmentPlugin)
    expected := NewEnvironmentPlugin().(*EnvironmentPlugin)

    err := e.Configure("", nil)
    if err != nil || !e.Equal(expected) { t.Fail() }
  })

  t.Run("configuration is empty", func(t *testing.T) {
    e := NewEnvironmentPlugin().(*EnvironmentPlugin)
    expected := NewEnvironmentPlugin().(*EnvironmentPlugin)

    err := e.Configure("", map[string]interface{}{})
    if err != nil || !e.Equal(expected) { t.Fail() }
  })

  t.Run("configuration is valid", func(t *testing.T) {
    e := NewEnvironmentPlugin().(*EnvironmentPlugin)
    expected := &EnvironmentPlugin{ name: "test", file: "testing" }

    err := e.Configure("test", map[string]interface{}{"File": "testing"})
    if err != nil || !e.Equal(expected) { t.Fail() }
  })



}
