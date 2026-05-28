# pi-imgr

`pi-imgr` is a Go HTTP service with its executable entrypoint at `cmd/pi-imgr`.

## Prerequisites

- Go 1.26.2 or newer

## Build

From the repository root, build the binary with standard Go tooling:

```sh
go build -o bin/pi-imgr ./cmd/pi-imgr
```

That creates the executable at `bin/pi-imgr`.

## Run

Run the program directly with Go:

```sh
go run ./cmd/pi-imgr
```

Or run the binary you built:

```sh
./bin/pi-imgr
```

The server listens on `http://localhost:8080`.

## Check The Service

After starting the program, you can verify it is running with:

```sh
curl http://localhost:8080/api/status
```

Expected response:

```json
{"status":"ok"}
```

## Available API Endpoints

- `GET /api/status`
- `GET /api/devices`
- `POST /api/flash`
