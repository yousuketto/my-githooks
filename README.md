# README
This repository is my githooks.

## insert-branch-prefix-into-comment
`insert-branch-prefix-into-comment` insert branch prefix(the part before '/' the branch name) into commit comment.

1. Build
```
go build github.com/yousuketto/my-githooks/insert-branch-prefix-into-comment.go
```
2. Copy generated file to `.git/hooks/`.
3. Create the following `prepare-commit-msg`.
```
#!/bin/sh
$(dirname $0)/insert-branch-prefix-into-comment $@
```
