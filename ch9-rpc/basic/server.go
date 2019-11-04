package main

import (
	"github.com/keets2012/Micro-Go-Pracrise/ch9-rpc/basic/string-service"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	stringService := new(service.StringService)
	rpc.Register(stringService)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "127.0.0.1:1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
