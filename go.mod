module go-chat

go 1.13

replace go-lib => ../go-lib

require (
	github.com/davyxu/cellnet v4.1.0+incompatible // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/emirpasic/gods v1.12.0
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/protobuf v1.3.2
	go-lib v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.3.1
	google.golang.org/grpc v1.26.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.2.8
)
