<div align="center">
    <img alt="MeltCD Logo" height="200px" src="https://github.com/meltred/meltcd/assets/82411321/52b0c441-0d63-4afb-b5a6-fec145e3ba26">
</div>

# meltcd

#### Docs: https://cd.meltred.tech/docs
![Discord](https://img.shields.io/discord/1086894797622624257)


> [!IMPORTANT]
> `meltcd` is very far from production use, unless we achieve **1.0.0**

Argo-cd like GitDevOps Continuous Development platform for docker swarm.

## Install

#### Linux, MacOS and WSL.

```bash
curl -s https://install.meltred.tech/meltcd | bash
```

#### Windows

Download From [latest release](https://github.com/meltred/meltcd/releases/latest)

#### Go Install

```bash
go install github.com/meltred/meltcd@latest
```

## Architecture

![architecture](https://github.com/meltred/meltcd/assets/82411321/9af15c33-627d-4e10-9952-0bd9e6422bbd)

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

3. install frontend packages

```bash
pnpm --prefix=./ui install
```

4. build the frontend

```bash
pnpm --prefix=./ui build
```

This will update the latest frontend to `server/static`

5. run the app

```bash
go run main.go serve --verbose
```

This will start the server on port `11771`

Go to **Developer Docs** for more info. [Developer Docs](https://github.com/meltred/meltcd/tree/main/docs/dev)

---

# About Meltred

This project is sponsored and maintained by [Meltred](https://meltred.com). Meltred builds tools to manage software.

<a href="https://meltred.com"><img src="https://i.imgur.com/Lq1q7vO.png" alt="Meltred Logo" loading="lazy" height="50px" /></a>
