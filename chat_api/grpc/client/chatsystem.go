package client

import (
	"chat/proto"
	"google.golang.org/grpc"
)

type RailsGrpcClient struct {
	Client proto.ChatSystemClient
}

func NewRailsGrpcClient(grpcConnection *grpc.ClientConn) *RailsGrpcClient {
	return &RailsGrpcClient{
		Client: proto.NewChatSystemClient(grpcConnection),
	}
}
