# Contributing to Meltcd ✨

We really thank you for your interest in contributing to Meltcd!

This is a guide for you to get started with contributing to Meltcd. Please read it carefully before you start.

## Table of Contents

- [Getting Started](#getting-started)

  - [Project Structure](#project-structure)

    - [Using GNU Make](#using-gnu-make)

  - [Setting up locally](#setting-up-locally)

  - [Contributing to different parts of the project](#contributing-to-different-parts-of-the-project)

    - [Server](#server)

    - [Web Client](#web-client)

    - [CLI](#cli)

  - [Testing](#testing)

- [How to contribute?](#how-to-contribute)

Before we start, here are the important links you should know:

- [docs](https://cd.kunalsin9h.tech/docs)
- [to file an **issues**](https://github.com/kunalsin9h/meltcd/issues)
- [chat](https://discord.gg/Y2C6mEhhf3)

## Getting Started

Things you need on your local machine:

- [Docker](https://docs.docker.com/get-docker/)

> You should have `docker swarm` up and running
> To start `docker swarm` just do
>
> ```bash
> docker swarm init
> ```
>
> After this we are all done

### Project Structure

Meltcd is a monorepo, which have three things in it:

- Meltcd **API Server** in `server` directory, build with GoLang using Fiber framework.

- A **CLI** Client in `cmd` directory, build with GoLang using Cobra framework.

- A **Web** Client in `ui` directory, build with ReactJS.

The application is build as a CLI, where Server can be started using `meltcd server` command.

The server will server the API on `http://localhost:11771/api` and the web client on `http://localhost:11771`.

#### Using GNU Make

We highly recommend to use GNU Make to run the commands from the `Makefile`.

### Setting up locally

1. Fork the repository
2. Clone the repository

```bash
git clone https://github.com/your_username/meltcd
```

3. Change the directory

```bash
cd meltcd
```

4. Download frontend dependencies

```bash
pnpm --prefix ./ui install
pnpm --prefix ./ui build
# see the commands under `frontend` in `Makefile`, to use it without GNU Make
```

5. Download server / cli dependencies

```bash
go mod download

pnpm install # some npm packages are used like husky for pre-commit
```

6. Start the server

```bash
go run main.go server --verbose
```

7. Using CLI

```bash
go run main.go app ls
```

8. Using WEB Client

Open https://localhost:11771 in your browser.

## Contributing to different parts of the project

### Server

To contribute to the server, work in the `server` directory. Save the changes and run the server using `go run main.go server --verbose`.

Server also embeds `swagger spec` and `swagger ui` which can be accessed at `http://localhost:11771/swagger/index.html`.

To update the swagger spec, run

```bash
swag init --output ./docs/swagger
```

So every-time you need to update swagger spec and run the swerver you need to do

```bash
swag init --output ./docs/swagger
go run main.go server --verbose
```

The server will be running at post `11771` unless changed using environment variable, see the [docs](https://cd.kunalsin9h.tech/docs) for more information.

The api will be served at `http://localhost:11771/api`.

### Web Client

To contribute to the web client, work in the `ui` directory. The `vite` app in the `ui` directory will be build to the `server/static` directory so that the API Server can serve it.

Go to `ui` directory and run:

```bash
pnpm run dev
```

To start the development server. This will only start the frontend, so it will not be able to communicate with the server.

To start frontend with server, run:

```bash
make run

# or

pnpm --prefix build # you need to build the frontend first
go run main.go server --verbose
```

So to work on frontend, after changing the code, you need to build the frontend and start the server.

**Deploying application**: Head over to [docs](https://cd.kunalsin9h.tech/docs/) to see how to deploy the application.

### CLI

To contribute to the CLI, work in the `cmd` directory. CLI is just cobra app which just do http request to the server.

You can run the CLI using

```bash
go run main.go app ls
```

For other CLI command see [docs](https://cd.kunalsin9h.tech/docs/).

While the server is running, you can update the CLI code and test it.

## Testing

To run tests, run:

```bash
make test
```

# How to contribute?

Make sure you create a new `branch` before working on new feature or fixing a bug.

After you are done with the changes, create a `pull request`

Contributions are always welcome, no matter how large or small. If stuck, feel free to ask for help in the [chat](https://discord.gg/Y2C6mEhhf3).

🚀
