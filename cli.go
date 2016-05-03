package gitsemvers

import (
	"fmt"
	"io"
	"strings"

	"github.com/jessevdk/go-flags"
)

const (
	exitcodeOK = iota
	exitCodeParseFlagErr
	exitCodeErr
)

// CLI is for command line
type CLI struct {
	OutStream, ErrStream io.Writer
}

// Run the cli
func (cli *CLI) Run(argv []string) int {
	p, sv, err := parseArgs(argv)
	if err != nil {
		if ferr, ok := err.(*flags.Error); !ok || ferr.Type != flags.ErrHelp {
			p.WriteHelp(cli.ErrStream)
		}
		return exitCodeParseFlagErr
	}
	fmt.Fprintln(cli.OutStream, strings.Join(sv.VersionStrings(), "\n"))
	return exitcodeOK
}

func parseArgs(args []string) (*flags.Parser, *Semvers, error) {
	sv := &Semvers{}
	p := flags.NewParser(sv, flags.Default)
	p.Usage = "[OPTIONS]\n\nVersion: " + version
	_, err := p.ParseArgs(args)
	return p, sv, err
}
