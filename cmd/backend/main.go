// Command backend implements the storage. It uses a random sleep to introduce
// performance variance.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ucirello/groupcache-experiment/pkg/backend"
	"github.com/ucirello/groupcache-experiment/pkg/storage"
	"google.golang.org/grpc"
)

var (
	listenAddr = flag.String("listen", "localhost:8080", "backend listening address")
)

func init() {
	flag.Parse()
}

func main() {
	grpcServer := grpc.NewServer()
	backend.New(storage.New(), grpcServer)
	fmt.Println("backend listening address:", *listenAddr)
	l, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(l)
}
