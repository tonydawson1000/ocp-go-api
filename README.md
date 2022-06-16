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

### Test the App (from a different terminal) - Note the Hostname / Machine Name
- `curl localhost:8180`
- `curl localhost:8180/test`
- `curl localhost:8180/anyurihere/1/2/3/4`

---

## Build and upload a Container Image to public Container Registry

### Build Image
- `nerdctl build -t <image-name> .`
- e.g. - `nerdctl build -t tonydawson1000/ocp-go-api .`

### Verify (List) Image
- `nerdctl images`

### Run a Local Container Instance
- `nerdctl run --name hello-go --rm -p 8180:8180 <image-name>`
- e.g. - `nerdctl run --name hello-go --rm -p 8180:8180 tonydawson1000/ocp-go-api`

### Test the App (from a different terminal) - Note the Hostname / ContainerId
- `curl localhost:8180`
- `curl localhost:8180/test`
- `curl localhost:8180/anyurihere/1/2/3/4`

### Push to public Container Registry (Docker Hub)
- Login (Default Docker Hub)
    - `nerdctl login`
- Push
    - `nerdctl push <image-name>:<tag>`
    - e.g. - `nerdctl push tonydawson1000/ocp-go-api:latest`
