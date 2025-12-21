package cmd

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/Songmu/gitsemvers"
)

const cmdName = "git-semvers"

func Run(ctx context.Context, argv []string, outStream, errStream io.Writer) error {
	log.SetOutput(errStream)
	log.SetPrefix(fmt.Sprintf("[%s] ", cmdName))
	fs := flag.NewFlagSet(
		fmt.Sprintf("%s (v%s ref:%s)", cmdName, gitsemvers.Version, gitsemvers.Revision),
		flag.ContinueOnError,
	)
	fs.SetOutput(errStream)

	sv := gitsemvers.Semvers{}
	fs.StringVar(&sv.RepoPath, "r", ".", "Path to the git repository")
	fs.StringVar(&sv.RepoPath, "repo", ".", "Path to the git repository")
	fs.StringVar(&sv.GitPath, "g", "git", "Path to the git executable")
	fs.StringVar(&sv.GitPath, "git", "git", "Path to the git executable")
	fs.BoolVar(&sv.WithPreRelease, "P", false, "Include pre-release versions")
	fs.BoolVar(&sv.WithPreRelease, "with-pre-release", false, "Include pre-release versions")
	fs.BoolVar(&sv.WithBuildMetadata, "B", false, "Include build metadata versions")
	fs.BoolVar(&sv.WithBuildMetadata, "with-build-metadata", false, "Include build metadata versions")
	fs.StringVar(&sv.TagPrefix, "p", "", "Tag prefix for monorepo (e.g., 'tools' for 'tools/v1.0.0')")
	fs.StringVar(&sv.TagPrefix, "prefix", "", "Tag prefix for monorepo (e.g., 'tools' for 'tools/v1.0.0')")

	if err := fs.Parse(argv); err != nil {
		return err
	}
	fmt.Fprintln(outStream, strings.Join(sv.VersionStrings(), "\n"))
	return nil
}
