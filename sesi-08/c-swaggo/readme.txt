Step swago:
go mod init swago
library:
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/swag/http-swagger
go get -u github.com/alecthomas/template
go get -u github.com/gorilla/mux

global:
go install github.com/swaggo/swag/cmd/swag

add path :
export PATH=$(go env GOPATH)/bin:$PATH