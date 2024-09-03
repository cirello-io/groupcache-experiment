// Package frontend implements a GRPC interface for the public facing part of
// the system.
package frontend

import (
	"context"
	"errors"

	"github.com/cirello-io/groupcache-experiment/pkg/api"
	"github.com/golang/groupcache"
	"google.golang.org/grpc"
)

// Server implements the public facing part of the system.
type Server struct {
	cacheGroup *groupcache.Group

	api.UnimplementedCacheServer
}

// New creates a new server and registers it to the GRPC server.
func New(cacheGroup *groupcache.Group, srv *grpc.Server) *Server {
	s := &Server{
		cacheGroup: cacheGroup,
	}
	api.RegisterCacheServer(srv, s)
	return s
}

// Get reads a key-pair value from the storage.
func (s *Server) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	var data []byte
	err := s.cacheGroup.Get(
		ctx,
		req.Key,
		groupcache.AllocatingByteSliceSink(&data),
	)
	return &api.GetResponse{
		Kv: &api.KV{
			Key:   req.Key,
			Value: string(data),
		},
	}, err
}

// Store saves a key-pair value in the storage.
func (s *Server) Store(_ context.Context, req *api.StoreRequest) (*api.StoreResponse, error) {
	return nil, errors.New("not supported")
}
