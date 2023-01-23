## `gh merge`

`gh merge` is a GitHub CLI extension to squash and merge PRs. The commit
message will be the PR description followed by the list of commits.

## Usage

```bash
$ gh merge --help

  usage:
    gh merge                  Squash and merge the current PR

  [options]
    -n, --number              Set a specific PR number
    -h, --help                Display the help information
```

## Installation

Make sure you have `gh` and `git` installed.

Then run:

```bash
$ gh extension install leofeyer/gh-merge
```
