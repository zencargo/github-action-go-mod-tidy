# github-action-go-mod-tidy

This action checks the tidiness of a Go module.

## Example usage

Until GitHub Actions [supports `uses` in composite actions](https://github.com/actions/runner/issues/646), you'll need to add `actions/cache` before this action.

```yaml
- uses: actions/cache@v2
  with:
    path: ~/go/pkg/mod
    key: go-mod-${{ hashFiles('**/go.sum') }}
    restore-keys: |
      go-mod-

- uses: zencargo/github-action-go-mod-tidy@v1
  with:
    path: my/go/module
    go-version: 1.15
```

## Inputs

### path

**Required**, default: `.`

Path of the module to check, relative to the GitHub workspace.

### go-version

**Required**, default: `1`

Version of the [golang](https://hub.docker.com/_/golang) `-alpine` Docker image to use.
