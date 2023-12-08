## Echo simple example

### Setup:

```
go mod init echo-example
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
```

### Start the server:

```
go run main.go
```

### Start the tests:

```bash
go test
#or
go test -v
```

### Start the test benchmark:

```bash
go test -benchmem -bench=BenchmarkRequestEcho
#or
go test -benchmem -bench RequestEcho -run=^$
```
