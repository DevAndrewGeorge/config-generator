package plugins

import (
  "testing"
)

func TestNewFilePlugin(t *testing.T) {
  if _, ok := NewFilePlugin().(*FilePlugin); !ok {
    t.Fail()
  }
}

func TestFilePluginConfigure(t *testing.T) {
  f := NewFilePlugin().(*FilePlugin)
  f.Configure("test", nil)
  if f.name != "test" { t.Fail() }
}
