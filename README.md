# Gin demo
For `Gin` learning, a Go web framework.
## learn_1, go hello world
1. touch main.go, run by `go run main.go`

2. config the package management tool, `mod`, using `go mod init 0_gin_learn`.

## learn_2, gin hello world

1. `go get -u github.com/gin-gonic/gin`
2. New engine instance.
3. Config specific Route, like "/".
4. Run && listen http.
5. Run main.go using `go run ./`.

## learn_3, Config hot reload by `fresh` package

1. `go get github.com/pilu/fresh`
2. init run project using `fresh`.(generate a tmp/)
3. It'll do hot reload automatically every time you save your code.

