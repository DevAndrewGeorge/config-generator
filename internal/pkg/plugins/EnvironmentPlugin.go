package plugins

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/devandrewgeorge/config-generator/internal/pkg/errors"
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
  var err error = nil

  valid_inputs := []string{"file"}
  for setting_name := range settings {
    for _, valid_input := range valid_inputs {
      if setting_name == valid_input {
        break
      }

      log.WithField("scope", "plugin").WithField("name", e.name).Error(
        fmt.Sprintf("%s is not a valid configuration setting", setting_name),
      )
      err = &errors.PluginError{}
    }
  }

  e.name = name

  if file, ok := settings["file"]; ok {
    e.file = file.(string)
  }

  return err
}
