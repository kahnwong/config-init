# Config Init

Because I got tired of setting up init scripts and looking up _that_ particular config file every time.

## Bonus

You can use [git-xargs](https://github.com/gruntwork-io/git-xargs) to mass update repo

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
