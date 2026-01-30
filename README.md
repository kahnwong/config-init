# Config Init

Because I got tired of setting up init scripts and looking up _that_ particular config file every time.

## Usage

```bash
❯ config-init
Utils to init various configs

Usage:
  config-init [command]

Available Commands:
  caddy          Init Caddyfile
  completion     Generate the autocompletion script for the specified shell
  compose        Init compose.yaml
  direnv         Init direnv
  dockerfile     Init Dockerfile
  github         Init GitHub-related configs
  github-actions Init Github Actions configs
  goreleaser     Init goreleaser config
  help           Help about any command
  make           Init Makefile
  pre-commit     Init pre-commit

Flags:
  -h, --help      help for config-init
  -t, --toggle    Help message for toggle
  -v, --version   version for config-init

Use "config-init [command] --help" for more information about a command.
```

## Bonus

You can use [git-xargs](https://github.com/gruntwork-io/git-xargs) to mass update repos

```bash
export GITHUB_TOKEN=
export PR_BRANCH_NAME=
export COMMIT_MESSAGE=
git-xargs \
  --no-skip-ci \
  --branch-name "$PR_BRANCH_NAME" \
  --repo $USER/$REPO \
  --commit-message "$COMMIT_MESSAGE" \
  $COMMAND
```

### Cookbooks

```bash
# pre-commit
yq e -i '(.repos[] | select(.repo== "https://github.com/DavidAnson/markdownlint-cli2").hooks[0].args) = "[--ignores, node_modules, src, --fix]"' .pre-commit-config.yaml
```
