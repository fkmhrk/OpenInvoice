package main

import (
	"./v1"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	err := v1.InitRouter(r)
	if err != nil {
		fmt.Printf("Failed to init server : %s", err)
		return
	}
	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Printf("Failed to call net.Listen : %s", err)
		return
	}
	//err = fcgi.Serve(l, r)
	err = http.Serve(l, r)
	if err != nil {
		fmt.Printf("failed to stop : %s", err)
	}
}
