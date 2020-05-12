package generator

import (
	"github.com/docopt/docopt-go"
)

type Args struct {
  Path string `docopt:"<path>"`
}

func parse_cli() (Args) {
	opts, _ := docopt.ParseDoc(`
generator

Usage:
	generator [options] <path>
`)

  args := Args{}
  opts.Bind(&args)
  return args
}

// func coerce(args *Args, opts *docopt.Opt) {
//
// }
