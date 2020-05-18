package plugins

import (
  "testing"
)

func TestNewMultiPlugin(t *testing.T) {
  if _, ok := NewMultiPlugin().(*MultiPlugin); !ok {
    t.Fail()
  }
}

func TestMultiPluginConfigure(t *testing.T) {
  m := NewMultiPlugin().(*MultiPlugin)
  m.Configure("test", nil)
  if m.name != "test" {
    t.Fail()
  }
}
