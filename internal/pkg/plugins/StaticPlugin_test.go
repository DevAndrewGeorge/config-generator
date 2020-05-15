package plugins

import (
  "testing"
)

func TestNewStaticPlugin(t *testing.T) {
  if _, ok := NewStaticPlugin().(*StaticPlugin); !ok {
    t.Fail()
  }
}

func TestStaticPluginConfigure(t *testing.T) {}
