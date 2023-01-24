## `gh merge`

`gh merge` is a GitHub CLI extension to squash and merge PRs. The commit
message will be the PR description followed by the list of commits.

## Usage

```bash
$ gh merge --help

Usage:  gh merge {<number>} [options]

Options:
  -h, --help   Display the help information
```

## Installation

Make sure you have `gh` and `git` installed.

Then run:

```bash
$ gh extension install leofeyer/gh-merge
```
