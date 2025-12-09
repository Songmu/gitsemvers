gitsemvers
=======

[![Test Status](https://github.com/Songmu/gitsemvers/workflows/test/badge.svg?branch=main)][actions]
[![codecov.io](https://codecov.io/github/Songmu/gitsemvers/coverage.svg?branch=main)][codecov]
[![MIT License](https://img.shields.io/github/license/Songmu/gitsemvers)][license]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/Songmu/gitsemvers)][PkgGoDev]

[actions]: https://github.com/Songmu/gitsemvers/actions?workflow=test
[codecov]: https://codecov.io/github/Songmu/gitsemvers?branch=main
[license]: https://github.com/Songmu/gitsemvers/blob/main/LICENSE
[PkgGoDev]: https://pkg.go.dev/github.com/Songmu/gitsemvers

## Description

Retrieve semvers from git tags

## Synopsis

```go
sv := &gitsemvers.Semvers{RepoPath: "path/to/repo"}
semvers := sv.VersionStrings()
```

### Monorepo Support

For monorepo projects with prefixed tags (e.g., `tools/v1.0.0`):

```go
sv := &gitsemvers.Semvers{
    RepoPath:  "path/to/repo",
    TagPrefix: "tools",
}
semvers := sv.VersionStrings() // returns ["tools/v1.1.0", "tools/v1.0.0", ...]
```

## Command Line Tool

    % go get github.com/Songmu/gitsemvers/cmd/git-semvers
    % git-semvers
    v0.9.0
    ...

## Author

[Songmu](https://github.com/Songmu)
