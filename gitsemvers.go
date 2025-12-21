package gitsemvers

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/mod/semver"
)

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
	RepoPath          string
	GitPath           string
	WithPreRelease    bool
	WithBuildMetadata bool
	TagPrefix         string
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

type ver struct {
	orig, version string
}

func (sv *Semvers) parseVersions(out string) []string {
	rawTags := strings.Split(out, "\n")
	var vers []*ver

	// Compute prefix once outside the loop
	var prefix string
	if sv.TagPrefix != "" {
		prefix = strings.TrimSuffix(sv.TagPrefix, "/") + "/"
	}

	for _, tag := range rawTags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}

		// Handle tag prefix for monorepo support
		semvPart := tag
		if prefix != "" {
			if !strings.HasPrefix(tag, prefix) {
				continue
			}
			semvPart = strings.TrimPrefix(tag, prefix)
			if semvPart == "" {
				continue
			}
		}

		semv := semvPart
		if semv != "" && semv[0] != 'v' {
			semv = "v" + semv
		}
		v := &ver{tag, semv}
		if semver.IsValid(semv) {
			hasBuild := semver.Build(semv) != ""
			isPrerelease := semver.Prerelease(semv) != ""
			if hasBuild && !sv.WithBuildMetadata {
				continue
			}
			if isPrerelease && !sv.WithPreRelease {
				continue
			}
			vers = append(vers, v)
		}
	}
	sort.Slice(vers, func(i, j int) bool {
		return semver.Compare(vers[i].version, vers[j].version) > 0
	})
	ret := make([]string, len(vers))
	for i, v := range vers {
		ret[i] = v.orig
	}
	return ret
}
