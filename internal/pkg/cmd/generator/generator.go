package generator

import (
  "github.com/devandrewgeorge/config-generator/internal/pkg/parser"
  log "github.com/sirupsen/logrus"
)

func Run() {
	args := parse_cli()

  // TODO: have this value set by argument
  setup_logging(log.DebugLevel)
  parser.ParseFile(args.Path)
}
