// Command cli operates the setup.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ucirello/groupcache-experiment/pkg/client"
)

var (
	listenAddr = flag.String("listen", "localhost:9001", "frontend listening address")
	backend    = flag.String("backend", "localhost:8080", "backend listening address")

	get = flag.Bool("get", false, "get a key-value pair")
	set = flag.Bool("set", false, "set a new key-value pair")

	key   = flag.String("k", "", "key to get/set")
	value = flag.String("v", "", "value to set")
)

func init() {
	flag.Parse()
}

func main() {
	switch true {
	case *get:
		getKV(*listenAddr, *key)
	case *set:
		setKV(*backend, *key, *value)
	default:
		flag.PrintDefaults()
	}
}

func getKV(addr, key string) {
	c := client.New(addr)
	v, err := c.Get(key)
	checkErr(err)
	fmt.Println(v)
}

func setKV(addr, k, v string) {
	c := client.New(addr)
	err := c.Set(k, v)
	checkErr(err)
}

func checkErr(err interface{}) {
	if err != nil {
		log.Fatalln("error:", err)
	}
}
