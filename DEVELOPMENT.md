# Development

Download the codebase from GitHub to your local machine.

```bash
cd /home/user/workspace/
git clone git@github.com:initialcapacity/freshcloud.git
cd freshcloud
```

## Tests

Run the test suite.

```bash
go test  ./.../
```

Clean up your test cache.

```bash
go clean -testcache
```

## Building locally

Build and install the Fresh Cloud command line interface.

```bash
go build cmd/freshctl.go
go install cmd/freshctl.go
```

Check the `freshctl` installation path as needed.

```bash
go list -f '{{.Target}}' cmd/freshctl.go
```
