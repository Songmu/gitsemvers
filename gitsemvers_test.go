package gitsemvers

import (
	"reflect"
	"testing"
)

var input = `dummy
v0.10.1
v0.9.0
v0.9.3
v0.8.4-pre
v0.8.4
v0.8.3+win
v0.8.2-pre.pre+win.win
v0.7.0-pre+win+invalid
`

func TestParseVersions(t *testing.T) {
	expect := []string{
		"v0.10.1",
		"v0.9.3",
		"v0.9.0",
		"v0.8.4",
	}
	sv := &Semvers{}
	if !reflect.DeepEqual(sv.parseVersions(input), expect) {
		t.Errorf("something went wrong %+v", sv.parseVersions(input))
	}
}

func TestParseVersionsWithPreRelease(t *testing.T) {
	expect := []string{
		"v0.10.1",
		"v0.9.3",
		"v0.9.0",
		"v0.8.4",
		"v0.8.4-pre",
	}
	sv := &Semvers{WithPreRelease: true}
	if !reflect.DeepEqual(sv.parseVersions(input), expect) {
		t.Errorf("something went wrong %+v", sv.parseVersions(input))
	}
}

func TestParseVersionsWithBuildMetadata(t *testing.T) {
	expect := []string{
		"v0.10.1",
		"v0.9.3",
		"v0.9.0",
		"v0.8.4",
		"v0.8.3+win",
	}
	sv := &Semvers{WithBuildMetadata: true}
	if !reflect.DeepEqual(sv.parseVersions(input), expect) {
		t.Errorf("something went wrong %+v", sv.parseVersions(input))
	}
}

func TestParseVersionsWithAllExtensions(t *testing.T) {
	expect := []string{
		"v0.10.1",
		"v0.9.3",
		"v0.9.0",
		"v0.8.4",
		"v0.8.4-pre",
		"v0.8.3+win",
		"v0.8.2-pre.pre+win.win",
	}
	sv := &Semvers{WithPreRelease: true, WithBuildMetadata: true}
	if !reflect.DeepEqual(sv.parseVersions(input), expect) {
		t.Errorf("something went wrong %+v", sv.parseVersions(input))
	}
}
