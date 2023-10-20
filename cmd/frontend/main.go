// Command frontend implements the publicly facing part of the system. Users
// looking for the content will hit here first, and groupcache will take of
// actually load data from the backend.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/cirello-io/groupcache-experiment/pkg/client"
	"github.com/cirello-io/groupcache-experiment/pkg/frontend"
	"github.com/golang/groupcache"
	"google.golang.org/grpc"
)

var (
	listenAddr   = flag.String("listen", "http://localhost:8001", "groupcache listen address")
	frontendAddr = flag.String("frontend", "localhost:9001", "frontend listen address")
	backend      = flag.String("backend", "localhost:8080", "backend listen address")
	gcpeers      = flag.String("peers", "http://localhost:8001,http://localhost:8002,http://localhost:8003", "groupcache peers")
)

func init() {
	flag.Parse()
}

func main() {
	peers := groupcache.NewHTTPPool(*listenAddr)
	backendClient := client.New(*backend)

	var stringCache = groupcache.NewGroup(
		"BackendCache",
		64<<20,
		groupcache.GetterFunc(func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			result, err := backendClient.Get(key)
			if err != nil {
				return err
			}
			dest.SetBytes([]byte(result))
			return nil
		}),
	)

	peers.Set(strings.Split(*gcpeers, ",")...)

	l, err := net.Listen("tcp", *frontendAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	frontend.New(stringCache, grpcServer)
	grpcServer.Serve(l)

	fmt.Println("cachegroup slave listening address:", *listenAddr)
	fmt.Println("frontend  listening address:", *frontendAddr)
	fmt.Println("peers pool:", strings.Split(*gcpeers, ","))
	log.Fatalln(http.ListenAndServe(strings.Replace(*listenAddr, "http://", "", 1), http.HandlerFunc(peers.ServeHTTP)))
}
