---
on: pull_request
permissions:
  contents: read
  pull-requests: read
safe-outputs:
  add-comment:
    target: "triggering" # Ensures the comment is posted on the PR that started the workflow
mcp-servers:
  safedep:
    url: "https://mcp.safedep.io/model-context-protocol/threats/v1/mcp"
    headers:
      Authorization: "${{ secrets.API_TOKEN }}"
      X-Tenant-ID: "${{ secrets.TENANT }}"
    allowed: ["*"]
---

# SafeDep Security Check

Your task is to, using SafeDep security checks, find if the OSS pacakges introduced / updated in the Pull Reqeust is safe to merge. 
