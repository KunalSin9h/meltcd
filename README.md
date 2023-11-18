# meltcd

Argo-cd like GitDevOps Continuous Development platform for docker swarm.

# Local Setup

### Requirements
1. GoLang
2. pnpm

1. Clone the
2. Download go packages
```bash
go mod download
```
3. install fontend packages
```bash
pnpm --prefix=./ui install
```

4. build the fontend
```bash
pnpm --prefix=./ui build
```
This will update the latest fontend to `server/static`

5. run the app

```bash
go run main.go
```

---

# About Meltred

This project is sponsored and maintained by [Meltred](https://meltred.com). Meltred builds tools to manage software.

<a href="https://meltred.com"><img src="https://i.imgur.com/Lq1q7vO.png" alt="Meltred Logo" loading="lazy" height="50px" /></a>
