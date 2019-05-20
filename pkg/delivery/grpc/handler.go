package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"m15.io/alpha/pkg/delivery/grpc/conf_grpc"
)

// GrpcConfHandler implements ConfHandler declared in confrepository.go
type GrpcConfHandler struct {
	Hostname string
	Port     int
}

// NewGrpcConfHandler returns pointer to newly created instance of GrpcConfHandler
func NewGrpcConfHandler(hostname string, port int) *GrpcConfHandler {
	grpcConfHandler := new(GrpcConfHandler)
	grpcConfHandler.Hostname = hostname
	grpcConfHandler.Port = port

	return grpcConfHandler
}

// GetConf calls remote procedure
func (handler *GrpcConfHandler) GetConf(fetchRequest *conf_grpc.FetchRequest) (*conf_grpc.Conf, error) {
	target := fmt.Sprintf("%v:%d", handler.Hostname, handler.Port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := conf_grpc.NewConfHandlerClient(conn)

	response, err := client.GetConf(context.Background(), fetchRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}
