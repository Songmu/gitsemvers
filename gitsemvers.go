package gitsemvers

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
)

const version = "0.0.1"

var verRegStr = `^v?[0-9]+(?:\.[0-9]+){0,2}`
var extension = `[-0-9A-Za-z]+(?:\.[-0-9A-Za-z]+)*`
var withPreReleaseRegStr = "(?:-" + extension + ")?"
var withBuildMetadataRegStr = `(?:\+` + extension + ")?"

type regBuilder uint

const (
	naked regBuilder = 0

	withPreRelease = 1 << (iota - 1)
	withBuildMetadata

	withPreReleaseAndBuildMetadata = withPreRelease | withBuildMetadata
)

var cache = make(map[regBuilder]*regexp.Regexp)

func (rb regBuilder) build() string {
	b := bytes.NewBufferString(verRegStr)
	if rb&withPreRelease != 0 {
		b.WriteString(withPreReleaseRegStr)
	}
	if rb&withBuildMetadata != 0 {
		b.WriteString(withBuildMetadataRegStr)
	}
	b.WriteString("$")
	return b.String()
}

func (rb regBuilder) reg() *regexp.Regexp {
	return cache[rb]
}

func init() {
	regs := []regBuilder{naked, withPreRelease, withBuildMetadata, withPreReleaseAndBuildMetadata}
	for _, v := range regs {
		cache[v] = regexp.MustCompile(v.build())
	}
}

// Semvers retrieve semvers from git tags
type Semvers struct {
	RepoPath          string `short:"r" long:"repo" default:"." description:"git repository path"`
	GitPath           string `short:"g" long:"git" default:"git" description:"git path"`
	WithPreRelease    bool   `short:"P" long:"with-pre-release" description:"display pre-release versions"`
	WithBuildMetadata bool   `short:"B" long:"with-build-metadata" description:"display build-metadata versions"`
}

// VersionStrings returns version strings
func (sv *Semvers) VersionStrings() []string {
	tags, err := sv.gitTags()
	if err != nil {
		return nil
	}
	return sv.parseVersions(tags)
}

func (sv *Semvers) reg() *regexp.Regexp {
	regB := regBuilder(0)
	if sv.WithPreRelease {
		regB |= withPreRelease
	}
	if sv.WithBuildMetadata {
		regB |= withBuildMetadata
	}
	return regB.reg()
}

func (sv *Semvers) gitProg() string {
	if sv.GitPath != "" {
		return sv.GitPath
	}
	return "git"
}

func (sv *Semvers) gitTags() (string, error) {
	cmd := exec.Command(sv.gitProg(), "-C", sv.RepoPath, "tag")
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return b.String(), err
}

func (sv *Semvers) parseVersions(out string) []string {
	rawTags := strings.Split(out, "\n")
	var versions []*semver.Version
	for _, tag := range rawTags {
		t := strings.TrimSpace(tag)
		if sv.reg().MatchString(t) {
			v, _ := semver.NewVersion(t)
			versions = append(versions, v)
		}
	}
	sort.Sort(sort.Reverse(semver.Collection(versions)))
	var vers = make([]string, len(versions))
	for i, v := range versions {
		vers[i] = v.Original()
	}
	return vers
}
