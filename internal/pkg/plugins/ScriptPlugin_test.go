package plugins

import (
  "testing"
  "os"
)

func TestNewScriptPlugin(t *testing.T) {
  s, ok := NewScriptPlugin().(*ScriptPlugin)
  if !ok {
    t.Fail()
  }

  if s.user != string(os.Getuid()) || s.group != string(os.Getgid()) {
    t.Fail()
  }
}

func TestScriptPluginConfigure(t *testing.T) {
  s := NewScriptPlugin().(*ScriptPlugin)
  with_names := map[string]interface{}{
    "shell": "/bin/sh",
    "user": "root",
    "group": "root",
  }

  with_ids := map[string]interface{}{
    "shell": "/bin/sh",
    "user": 0,
    "group": 0,
  }

  t.Run("[plugins] [ScriptPlugin.go] Configure() with names", func(t *testing.T) {
    s.Configure("test", with_names)
    if s.name != "test" || s.shell != "/bin/sh" || s.user != "0" || s.group != "0" {
      t.Fail()
    }
  })

  t.Run("[plugins] [ScriptPlugin.go] Configure() with IDs", func(t *testing.T) {
    s.Configure("test", with_ids)
    if s.name != "test" || s.shell != "/bin/sh" || s.user != "0" || s.group != "0" {
      t.Fail()
    }
  })
}
