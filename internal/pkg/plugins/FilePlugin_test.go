package plugins

import (
  "testing"
)

func TestNewFilePlugin(t *testing.T) {
  if _, ok := NewFilePlugin().(*FilePlugin); !ok {
    t.Fail()
  }
}

func TestFilePluginConfigure(t *testing.T) {}
