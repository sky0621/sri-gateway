package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/sky0621/sri-gateway/server"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := server.Run(); err != nil {
		glog.Fatal(err)
	}
}
