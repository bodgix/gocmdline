package cmds

import docopt "github.com/docopt/docopt-go"

// ParseArgs cmd line arguments parsing
func ParseArgs() (map[string]interface{}, error) {
	usage := `gocmdline
Usage:
  gocmdline start server <name> [options]
  gocmdline start client <name> [options]

Options:
  -h --help     Show this screen.
  --version     Show version
  -v --verbose  Verbose output`

	return docopt.Parse(usage, nil, true, "gocmdline", false)
}
