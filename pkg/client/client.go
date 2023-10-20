// Package client implements the client-side interface of the protobuf defined
// API.
package client

import (
	"context"
	"fmt"
	"log"

	"github.com/cirello-io/groupcache-experiment/pkg/api"
	"google.golang.org/grpc"
)

// Client talks to GRPC enabled servers.
type Client struct {
	api api.CacheClient
}

// New instantiates a new client by dialing the GRPC server first.
func New(addr string) *Client {
	cc, err := grpc.Dial(addr, grpc.WithInsecure()) // for testing purpose, no security.
	if err != nil {
		log.Fatalln("error:", err)
	}
	return &Client{
		api: api.NewCacheClient(cc),
	}
}

// Get loads key from the server.
func (c *Client) Get(key string) (string, error) {
	req := &api.GetRequest{Key: key}
	resp, err := c.api.Get(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("cannot get key-value pair: %w", err)
	}
	return resp.Kv.Value, nil
}

// Set saves a key-value in the server.
func (c *Client) Set(key string, value string) error {
	req := &api.StoreRequest{Kv: &api.KV{
		Key:   key,
		Value: value,
	}}
	_, err := c.api.Store(context.Background(), req)
	if err != nil {
		return fmt.Errorf("cannot store key-value pair: %w", err)
	}
	return nil
}
