This is a playground Go project with a simple web request using standard net/http Go library.

_09-03-2025_

`go run main.go` - To run the server on [localhost:](http://localhost:8090/hello)

`go mod init <file-name>` - To initializes a new module in the current directory

`go build` - To compile the code into an executable

Run `docker-compose up --build` to enable to auto-recompiling and auto-restarting the Go server.

`go mod download` - Downloads modules without modifying go.mod or go.sum

`go mod tidy` - Cleans and updates go.mod and go.sum by ensuring they reflect the actual dependencies required by the code.
