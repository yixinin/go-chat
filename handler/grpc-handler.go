package handler

type Grpc struct {
}

func NewGrpcHandler() *Grpc {
	return &Grpc{}
}

func (*Grpc) String() string {
	return "grpc"
}
