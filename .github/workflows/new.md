---
on: pull_request
permissions:
  contents: read
  pull-requests: read
safe-outputs:
  add-comment:
    target: "triggering" # Ensures the comment is posted on the PR that started the workflow [4, 5]
---

# Pull Request Commenter Agent
You are an assistant that reviews pull requests and provides a helpful summary for 
