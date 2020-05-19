package parser

import (
  "io/ioutil"
  log "github.com/sirupsen/logrus"
  "gopkg.in/yaml.v2"
  "github.com/devandrewgeorge/config-generator/internal/pkg/generator"
  "github.com/devandrewgeorge/config-generator/internal/pkg/errors"
)

func ParseFile(config_path string) (*generator.Generator, error) {
  data, err := ioutil.ReadFile(config_path)
  if err != nil {
    log.WithField("scope", "config").WithField("name", config_path).Error(err.Error())
    return nil, &errors.ConfigError{}
  }
  return Parse(data)
}

func Parse(config_raw []byte) (*generator.Generator, error) {
  log := log.WithField("scope", "config")

  var config GeneratorConfig
  if err := yaml.UnmarshalStrict(config_raw, &config); err != nil {
    log.Error(err.Error())
    return nil, &errors.ConfigError{}
  }

  g := generator.New()
  if plugins, err := parsePlugins(config.Plugins); err != nil {
    return nil, err
  } else {
    for plugin_name, plugin := range plugins {
      g.Plugins[plugin_name] = plugin
    }
  }

  return g, nil
}
