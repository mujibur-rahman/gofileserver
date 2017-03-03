//fileserver will be running on port 4000 default
//it will provide the html/css/javascript/images
//basically all static resources.
//Nginx is the reverse proxy which will be using
// this fileserver.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

var (
	port *int
)

func init() {
	port = flag.Int("port", 4000, "Http file server running on port")
}

func panicRecover() {
	defer func() {
		if rec := recover(); rec != nil {
			log.Println(rec)
			log.Printf("%s\n", debug.Stack())
		}
	}()
}

func main() {
	panicRecover()
	flag.Parse()
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	listenAddr := fmt.Sprintf(":%d", *port)
	log.Printf("File Server Listening at port %s...", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalf("Error: %v", err)
	}

}
