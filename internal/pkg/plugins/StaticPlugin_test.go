package plugins

import (
  "testing"
)

func TestNewStaticPlugin(t *testing.T) {
  if _, ok := NewStaticPlugin().(*StaticPlugin); !ok {
    t.Fail()
  }
}

func TestStaticPluginConfigure(t *testing.T) {
  s := NewStaticPlugin().(*StaticPlugin)
  s.Configure("test", nil)
  if s.name != "test" { t.Fail() }
}
