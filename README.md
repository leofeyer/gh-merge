## `gh merge`

`gh merge` is a GitHub CLI extension to squash and merge PRs. The commit
message will be the PR description followed by the list of commits.

## Installation

Make sure you have `gh` and `git` installed. Then run:

```bash
$ gh extension install leofeyer/gh-merge
```

## Usage

```bash
$ gh merge 1234
```

If you are on a branch that you have checked out with `gh pr checkout 1234`,
you can omit the PR number:

```bash
$ gh merge
```

Use the `--auto` flag to enable auto-merging:

```bash
$ gh merge 1234 --auto
```

Use the `--admin` flag to merge the PR with admin privileges:

```bash
$ gh merge 1234 --admin
```
