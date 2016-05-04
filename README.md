gitsemvers
=======

[![Build Status](https://travis-ci.org/Songmu/gitsemvers.png?branch=master)][travis]
[![Coverage Status](https://coveralls.io/repos/Songmu/gitsemvers/badge.png?branch=master)][coveralls]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/gitsemvers?status.svg)](godoc)

[travis]: https://travis-ci.org/Songmu/gitsemvers
[coveralls]: https://coveralls.io/r/Songmu/gitsemvers?branch=master
[license]: https://github.com/Songmu/gitsemvers/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/Songmu/gitsemvers

## Description

Retrieve semvers from git tags

## Synopsis

```go
sv := &gitsemvers.Semvers{RepoPath: "path/to/repo"}
semvers := sv.VersionStrings()
```

## Command Line Tool

    % go get github.com/Songmu/gitsemvers/cmd/git-semvers
    % git-semvers
    v0.9.0
    ...

## Author

[Songmu](https://github.com/Songmu)
