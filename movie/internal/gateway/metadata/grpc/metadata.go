package grpc

import (
	"context"

	"github.com/MSVelan/movieapp/gen"
	"github.com/MSVelan/movieapp/internal/grpcutil"
	"github.com/MSVelan/movieapp/metadata/pkg/model"
	"github.com/MSVelan/movieapp/pkg/discovery"
)

// Gateway defines a movie metadata gRPC gateway.
type Gateway struct {
	registry discovery.Registry
}

// New creates a new gRPC gateway for movie metadata service.
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

// Get returns movie metadata by a movie id.
func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "metadata", g.registry)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := gen.NewMetadataServiceClient(conn)
	resp, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{MovieId: id})
	if err != nil {
		return nil, err
	}
	return model.MetadataFromProto(resp.Metadata), nil
}
