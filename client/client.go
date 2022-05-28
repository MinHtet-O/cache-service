package main

import (
	"cache-service/z_generated/pb"
	"errors"
	"flag"
	"log"
)

var (
	addr    = flag.String("addr", "localhost:50051", "the address to connect rpc server")
	timeout = flag.Int("timeout", 1, "timeout duration in second")

	set = flag.Bool("set", false, "set raw cache")
	get = flag.Bool("get", false, "get raw cache value")

	setUser = flag.Bool("set-user", false, "set user cache")
	getUser = flag.Bool("get-user", false, "get user cache value")
)

func main() {
	flag.Parse()
	var c pb.CacheClient
	conn := connectRPC(&c)
	if *set {
		if !(len(flag.Args()) >= 2) {
			log.Fatal(errors.New("please provide both key and value"))
			return
		}
		key := flag.Args()[0]
		val := flag.Args()[1]
		setCache(c, key, val)
		return
	}
	if *get {
		if !(len(flag.Args()) >= 1) {
			log.Fatal(errors.New("please provide key"))
			return
		}
		key := flag.Args()[0]
		getCache(c, key)
		return
	}
	if *setUser {
		if !(len(flag.Args()) >= 4) {
			log.Fatal(errors.New("please provide name class roll and metadata for user"))
			return
		}
		name := flag.Args()[0]
		class := flag.Args()[1]
		roll := flag.Args()[2]
		metadata := flag.Args()[3]
		setUserCache(c, name, class, roll, metadata)
		return
	}
	if *getUser {
		if !(len(flag.Args()) >= 3) {
			log.Fatal(errors.New("please provide name class and roll as to find the user"))
			return
		}
		name := flag.Args()[0]
		class := flag.Args()[1]
		roll := flag.Args()[2]
		getUserCache(c, name, class, roll)
		return
	}
	defer conn.Close()
}
