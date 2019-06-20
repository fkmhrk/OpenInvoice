package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/fkmhrk/OpenInvoice/v1"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	err := v1.InitRouter(r)
	if err != nil {
		fmt.Printf("Failed to init server : %s", err)
		return
	}
	var standalone *bool = flag.Bool("standalone", false, "if true, runs standalone mode")
	flag.Parse()

	if *standalone {
		initStandalone(r)
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

func initStandalone(r *mux.Router) {
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("../web/"))))
}
