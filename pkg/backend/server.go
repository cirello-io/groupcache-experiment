// Package backend provides the GRPC interface for the backend.
package backend

import (
	"context"

	"github.com/ucirello/groupcache-experiment/pkg/api"
	"github.com/ucirello/groupcache-experiment/pkg/storage"
	"google.golang.org/grpc"
)

// Server operates the storage.
type Server struct {
	storage *storage.Storage
}

// New creates a new server and registers it to the GRPC server.
func New(storage *storage.Storage, srv *grpc.Server) *Server {
	s := &Server{
		storage: storage,
	}
	api.RegisterCacheServer(srv, s)
	return s
}

// Get reads a key-pair value from the storage.
func (s *Server) Get(_ context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	v := s.storage.Get(req.Key)
	resp := &api.GetResponse{
		Kv: &api.KV{
			Key:   req.Key,
			Value: v,
		},
	}
	return resp, nil
}

// Store saves a key-pair value in the storage.
func (s *Server) Store(_ context.Context, req *api.StoreRequest) (*api.StoreResponse, error) {
	s.storage.Set(req.Kv.Key, req.Kv.Value)
	return &api.StoreResponse{}, nil
}
