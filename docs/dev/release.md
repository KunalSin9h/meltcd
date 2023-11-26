# Release new version of MeltCD on GitHub

Workflow `.github/workflows/release.yml` is responsible for releasing new version.

It will be triggered by manually running workflow from GitHub Actions UI.

It require new `tag` to be pushed to `main` branch

To create and push new tags

```bash
git tag 1.15.0

git push --force --tag
```

Or just push tag when pushing real prs

```bash
git push origin master --tag
```

This will update the tag in `main` branch.

After this trigger the workflow.
