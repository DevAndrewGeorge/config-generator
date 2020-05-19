package plugins

import (
  "testing"
)

func TestNewConsolePlugin(t *testing.T) {
  if _, ok := NewConsolePlugin().(*ConsolePlugin); !ok {
    t.Fail()
  }
}

func TestConsolePluginConfigure(t *testing.T) {
  t.Run("configuration is nil", func(t *testing.T) {
    c := NewConsolePlugin().(*ConsolePlugin)
    expected := NewConsolePlugin().(*ConsolePlugin)

    err := c.Configure("", nil)
    if err != nil || !c.Equal(expected) { t.Fail() }
  })

  t.Run("configuration is empty", func(t *testing.T) {
    c := NewConsolePlugin().(*ConsolePlugin)
    expected := NewConsolePlugin().(*ConsolePlugin)

    err := c.Configure("", map[string]interface{}{})
    if err != nil || !c.Equal(expected) { t.Fail() }
  })

  t.Run("configuration is valid", func(t *testing.T) {
    c := NewConsolePlugin().(*ConsolePlugin)
    expected := &ConsolePlugin{name: "test"}

    err := c.Configure("test", map[string]interface{}{})
    if err != nil || !c.Equal(expected) { t.Fail() }
  })
}
