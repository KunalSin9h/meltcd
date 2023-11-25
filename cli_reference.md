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
```

5. List all the running applications

```bash
meltcd app list
```

6. Force refresh (synchronize) the application [DONE]

```bash
meltcd app refresh <app-name>

# or

meltcd app sync <app-name>
```
