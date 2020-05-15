package plugins

import (
  "testing"
)

func TestNewConsolePlugin(t *testing.T) {
  if _, ok := NewConsolePlugin().(*ConsolePlugin); !ok {
    t.Fail()
  }
}

func TestConsolePluginConfigure(t *testing.T) {}
