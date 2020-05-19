package plugins

import (
  "reflect"
)

type EnvironmentPlugin struct {
  name string
  file string
}

func NewEnvironmentPlugin() Plugin {
  return &EnvironmentPlugin{}
}

func (e *EnvironmentPlugin) Equal(p Plugin) bool {
  a, ok := p.(*EnvironmentPlugin)
  if ok {
    return *e == *a
  }

  return false
}

func (e *EnvironmentPlugin) Configure(name string, settings map[string]interface{}) error {
  e.name = name
  if file, ok := settings["File"]; ok && !reflect.ValueOf(file).IsNil() {
    e.file = *file.(*string)
  }

  return nil
}
