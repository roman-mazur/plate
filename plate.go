// A static HTTP server that will serve files in the working directory.
// By default, will serve current process directory files on port 8080.
// Use --help flag for configuration options.
package main // import "rmazur.io/plate"

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	port    = flag.Int("port", 8080, "TCP port to listen on")
	address = flag.String("address", "", "Network address to listen on (all addresses by default)")
	dir     = flag.String("dir", ".", "Directory to serve files from")
)

func main() {
	flag.Parse()

	h := new(http.ServeMux)
	h.Handle("/", http.FileServer(http.Dir(*dir)))

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", *address, *port),
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("Staring the HTTP server...")
		err := server.ListenAndServe()
		log.Fatalf("HTTP server finished: %s. Finishing the process.", err)
	}()

	intChannel := make(chan os.Signal)
	signal.Notify(intChannel, syscall.SIGINT, syscall.SIGTERM)
	<-intChannel
	log.Println("Shutting down...")
}
