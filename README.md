# git-rm-merged

git-rm-merged is git subcommand that removes the merged local branches.

## Description

You can use the `git-rm-merged` as a git subcommand.
This removes the merged local branches.

**This confirms if you want to delete the target branches.**

## Install

```
go get -u github.com/yasukotelin/git-rm-merged
```

## Usage

```
$ git rm-merged
feature/#1 [y/n]:y
feature/#2 [y/n]:n
feature/#3 [y/n]:y
```

## Licence

MIT

## Author

yasukotelin
