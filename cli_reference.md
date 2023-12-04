# CLI Commands

1. Create a new `Application` [DONE]

```bash
meltcd app create <app-name> --repo <repo-url> --path <path-to-spec>
```

2. Create a new `Application` with file [DONE]

```bash
meltcd app create --file <path-to-file>
```

3. Update existing `Application` [DONE]

```bash
meltcd app update <app-name> --repo <repo-url> --path <path-to-spec>

# Or using file

meltcd app update --file <path-to-file>
```

4. Get details about `Application` [DONE]

```bash
meltcd app get <app-name>

# or
meltcd app inspect <app-name>
```

5. List all the running applications

```bash
meltcd app list

# or

meltcd app ls
```

6. Force refresh (synchronize) the application [DONE]

```bash
meltcd app refresh <app-name>

# or

meltcd app sync <app-name>
```

7. Remove an application [DONE]

```bash
meltcd app remove <app-name>

# or

meltcd app rm <app-name>
```

# Private Repository

1. Add a private repository auth credentials [DONE]

```bash
meltcd repo add <repo-url> --username <username> --password <password>
```

2. List all added repositories [DONE]

```bash
meltcd repo ls

#or

meltcd repo list
```

3. Remove a repository [DONE]

```bash
meltcd repo rm <repo-url>

# or

meltcd repo remove <repo-url>
```

4. Update a repository [DONE]

```bash
meltcd repo update <repo-url> --username <username> --password <password>
```

# Private Image

1. Add a image auth credentials [TODO]

```bash
meltcd image add <image-name> --username <username> --password <password>
```
