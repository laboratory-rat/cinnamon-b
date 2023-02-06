package impl

import (
	proto2 "api/main/src/bclient/proto"
	"context"
)

type Client struct {
	name    string
	version string
	token   string
	proto2.UnsafeConnectorServer
}

// NewClient creates new Client
func NewClient(token string) *Client {
	return &Client{
		token:   token,
		name:    "privat_bank",
		version: "0.0.1",
	}
}

func NewConnectorServer(token string) proto2.ConnectorServer {
	return Client{
		token:   token,
		name:    "privat_bank",
		version: "0.0.1",
	}
}

func (c Client) MyType(ctx context.Context, request *proto2.MyTypeRequest) (*proto2.MyTypeResponse, error) {
	return &proto2.MyTypeResponse{
		Name:    c.name,
		Version: c.version,
	}, nil
}

func (c Client) CanPing(ctx context.Context, request *proto2.CanPingRequest) (*proto2.CanPingResponse, error) {
	return &proto2.CanPingResponse{
		Success: true,
	}, nil
}

func (c Client) Ping(ctx context.Context, request *proto2.PingRequest) (*proto2.PingResponse, error) {
	return &proto2.PingResponse{
		Success: true,
	}, nil
}
