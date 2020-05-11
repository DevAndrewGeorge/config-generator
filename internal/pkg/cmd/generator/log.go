package generator

import (
  "fmt"
  log "github.com/sirupsen/logrus"
)

type GeneratorFormatter struct {}
func (f *GeneratorFormatter) Format(entry *log.Entry) ([]byte, error) {
  log_line := fmt.Sprintf("[%s] [%s] [%s] %s\n", entry.Level, entry.Data["scope"], entry.Data["name"], entry.Message)
  return []byte(log_line), nil
}

func setup_logging(log_level log.Level) {
  log.SetFormatter(new(GeneratorFormatter))
}
