go install github.com/spf13/cobra-cli@latest

cobra-cli init
cobra-cli add args
cobra-cli add flags
cobra-cli add integration

go run main.go
go run main.go add
go run main.go add --port 9090 --debug true

go run main.go args 5 10
go run main.go flags --port 9090 --debug true
go run main.go integration 5 10
go run main.go integration 5 10 --test hello
