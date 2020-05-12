package parser

import (
  "io/ioutil"
  "github.com/devandrewgeorge/config-generator/internal/pkg/generator"
  "github.com/sirupsen/logrus"
)

var log *logrus.Entry
func init() {
  log = logrus.WithField("scope", "praser")
}

type GeneratorConfig struct {
	Plugins   map[string]interface{}
	Variables map[string]interface{}
	Templates map[string]interface{}
	Outputs   map[string]interface{}
}

func ParseFile(config_path string) (*generator.Generator, error) {
  data, err := ioutil.ReadFile(config_path)
  if err != nil {
    log.WithField("name", config_path).Fatal("file not found")
  }
  return Parse(data)
}

func Parse(config []byte) (*generator.Generator, error) {
  return nil, nil
}
