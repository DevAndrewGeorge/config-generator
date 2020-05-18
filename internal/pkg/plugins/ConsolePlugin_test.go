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
  c := NewConsolePlugin().(*ConsolePlugin)
  c.Configure("test", nil)
  if c.name != "test" { t.Fail() }
}
