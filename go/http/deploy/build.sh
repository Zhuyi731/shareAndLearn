cd ../
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o ./bin/testHttpServer
go build -o ./bin/testHttpServer