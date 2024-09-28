### About
A container with a Go installation, including all dependencies specified in the `go.mod` file

### Run
1. Execute in terminal: `docker compose up`

The following logs should appear, indicating that the container is properly set up:

```bash
go_service-1  | ==> ALL DONE WITH SETTING UP THE GO ENVIRONMENT.
go_service-1  | ==> LOCATION: /go
go_service-1  | ==> VERSION: go version go1.23.1 linux/amd64
```

To inspect the environment setup:

* Enter the container: `docker compose run go_service bash`
* List the installed packages: `go list ...`