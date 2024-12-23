<div align="center">
    <img alt="MeltCD Logo" height="200px" src="https://github.com/kunalsin9h/meltcd/assets/82411321/9065c92d-79a5-44ff-aa53-3e0bd40f0080">
</div>

# meltcd

#### Docs: https://cd.kunalsin9h.com/docs

![Discord](https://img.shields.io/discord/1086894797622624257)

> [!Caution]
> `meltcd` is not ready for production use, unless we achieve **1.0.0**

Argo-cd like GitDevOps Continuous Delivery platform for docker swarm.

<div align="center">
    <img alt="MeltCD Demo Page" src="https://i.imgur.com/LxaD7qM.png">
</div>

## Install

#### Linux, MacOS and WSL.

```bash
curl -s https://install.kunalsin9h.com/meltcd | bash
```

#### Windows

Download From [latest release](https://github.com/kunalsin9h/meltcd/releases/latest)

#### Go Install

```bash
go install github.com/kunalsin9h/meltcd@latest
```

## Architecture

![architecture](https://github.com/kunalsin9h/meltcd/assets/82411321/f73f80a5-a533-420d-aee9-6a06e2b13976)

## Local Setup

#### Requirements

1. GoLang
2. pnpm

#### Run

1. Clone the
2. Download go packages

```bash
go mod download
```

3. Install `husky`

```bash
pnpm install
```

4. Install `swag` from [here](https://github.com/swaggo/swag)

5. install frontend packages

```bash
pnpm --prefix=./ui install
```

6. build the frontend

```bash
pnpm --prefix=./ui build
```

This will update the latest frontend to `server/static`

7. run the app

```bash
go run main.go serve --verbose

# Using `gnu make`
make run
```
This will start the server on port `1771`

> [!TIP]
> If you get error saying **"Error response from daemon: This node is not a swarm manager. Use \"docker swarm init\" or \"docker swarm join\" to connect this node to swarm and try again."**
> This means you have docker working but the node is not a `Docker Swarm` Node, to make it run `docker swarm init`.

> [!TIP]
> If applications are unable to run, there might be a case of `root` privilege. To allow docker run without `sudo` do..
> ```bash
> sudo groupadd docker
> sudo usermod -aG docker $USER
> newgrp docker
> ```

Go to **Developer Docs** for more info. [Developer Docs](https://github.com/kunalsin9h/meltcd/tree/main/docs/dev)

## Contributing

We welcome contributions to `meltcd` in many forms. There's always plenty to do!

See the [Contribution Guide](https://github.com/kunalsin9h/meltcd/blob/main/CONTRIBUTING.md) for more information.
