package parser

import (
  "io/ioutil"
  "github.com/sirupsen/logrus"
  "github.com/devandrewgeorge/config-generator/internal/pkg/generator"
)

var log *logrus.Entry
func init() {
  log = logrus.WithField("scope", "praser")
}

func ParseFile(config_path string) (*generator.Generator, error) {
  data, err := ioutil.ReadFile(config_path)
  if err != nil {
    log.WithField("name", config_path).Fatal("file not found")
  }
  return Parse(data)
}

func Parse(config []byte) (*generator.Generator, error) {
  return generator.New(), nil
}
