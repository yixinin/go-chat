cd protocol
protoc -I. --csharp_out=./code/cs --go_out=plugins=grpc:. *.proto