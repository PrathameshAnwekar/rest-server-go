# rest-server-go

A simple REST server written in Go for learning best practices.

## Table of Contents

- [rest-server-go](#rest-server-go)
  - [Table of Contents](#table-of-contents)
    - [Prerequisites](#prerequisites)
    - [Building the Project](#building-the-project)

### Prerequisites

build_and_run.sh requires the following dependencies to be installed and present in PATH.

* [go 1.21](https://go.dev/dl/)
* [golangci-lint](https://golangci-lint.run/usage/quick-start/)
  ```
  brew install golangci-lint
  ```
* [gofumpt](https://pkg.go.dev/github.com/vearutop/gofumpt)
  ```
  brew install gofumpt
  ```
* [opencv](https://gocv.io/getting-started/macos/)
  ```
  brew install opencv
  brew install pkgconfig
  ```

### Building the Project

```bash
sh build_and_run.sh
```
