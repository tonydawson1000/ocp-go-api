# ocp-go-api
A very simple stateless HTTP API written in Go

---

## Install Go (from linux CLI)

### Set the Go Version
- `export VERSION=1.18.3`

### Download the Go archive.
- `curl -O https://dl.google.com/go/go$VERSION.linux-amd64.tar.gz`

### Verify the SHA256 Checksum (against the downloads page).
- `sha256sum go$VERSION.linux-amd64.tar.gz`

### Extract the tarball into the `/usr/local` directory.
- `tar -C /usr/local -xzf go$VERSION.linux-amd64.tar.gz`

### Add the Go binaries to your $PATH.
- `export PATH=$PATH:/usr/local/go/bin`
- To persist this across bash sessions, add to `~/.bashrc`

### Test Go Install worked
- `go version`

---

## Build, Run and Test

### Build the App
- `go build cmd/hello/hello.go`

### Run the App (from a terminal)
- `./hello`

### Test the App (from a different terminal)
- `curl localhost:8180`
- `curl localhost:8180/test`
- `curl localhost:8180/anyurihere/1/2/3/4`