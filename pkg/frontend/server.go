// Package frontend implements a GRPC interface for the public facing part of
// the system.
package frontend

import (
	"context"

	"cirello.io/errors"
	"github.com/golang/groupcache"
	"github.com/ucirello/groupcache-experiment/pkg/api"
	"google.golang.org/grpc"
)

// Server implements the public facing part of the system.
type Server struct {
	cacheGroup *groupcache.Group
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
func (s *Server) Get(_ context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	var data []byte
	err := s.cacheGroup.Get(
		nil,
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
	return nil, errors.E(errors.Invalid, "not supported")
}
