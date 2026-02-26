---
on:
  pull_request:
    types: [opened, synchronize]
permissions:
      contents: read
engine: copilot
network: defaults
safe-outputs:
  submit-pull-request-review:
    max: 1            # max reviews to submit (default: 1)
    target: "triggering"  # or "*", or e.g. ${{ github.event.inputs.pr_number }} when not in pull_request trigger
    footer: false     # omit AI-generated footer from review body (default: true)
---

Say Hello to David, its his birthday!
