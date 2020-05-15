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
  e := NewEnvironmentPlugin().(*EnvironmentPlugin)
  e.Configure(map[string]interface{}{"file": "testing"})

  if e.file != "testing" {
    t.Fail()
  }
}
