# Backend Game Simulator
-
## Installation & Run

### Install and run with Docker

```bash
$ docker build -t backendsimulator .
$ docker run simulator
```

### Run with local
Before run, you may want to configure local environment (backend api url), you can update settings in under .env file.
```bash
$ go run main.go
```

### Run Unittest

```bash
$ go test -v ./... -count=1
```
