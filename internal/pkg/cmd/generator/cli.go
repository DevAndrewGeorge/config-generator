package generator

import (
	"github.com/docopt/docopt-go"
)

func parse_cli() (docopt.Opts, error) {
	return docopt.ParseDoc(`
generator

Usage:
	generator <path>
	`)
}
