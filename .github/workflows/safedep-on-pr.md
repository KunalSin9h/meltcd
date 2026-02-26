---
on:
  pull_request:
    types: [opened, synchronize]
permissions:
      contents: read
engine: copilot
network: defaults
safe-outputs:
  reply-to-pull-request-review-comment:
    max: 10                              # max replies (default: 10)
    target: "triggering"                 # "triggering" (default), "*", or number
    target-repo: "kunalsin9h/meltcd"            # cross-repository
    footer: true                         # add AI-generated footer (default: true)
    github-token: ${{ secrets.TOKEN_COMMENT }} # optional custom token for permissions
---

Say Hello to David, its his birthday!
