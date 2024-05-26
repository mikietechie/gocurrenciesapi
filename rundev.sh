export PATH=$(go env GOPATH)/bin:$PATH
swag init --parseDependency --parseInternal
go run main.go