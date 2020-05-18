package plugins

import (
  "testing"
)

func TestNewHttpPlugin(t *testing.T) {
  if _, ok := NewHttpPlugin().(*HttpPlugin); !ok {
    t.Fail()
  }
}

func TestHttpPluginConfigure(t *testing.T) {
  h := NewHttpPlugin().(*HttpPlugin)
  h.Configure("test", nil)
  if h.name != "test" { t.Fail() }
}
