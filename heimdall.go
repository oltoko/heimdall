package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	log.SetFlags(0)

	htmlDir := flag.String("html", ".", "Directory of the static HTML files")
	port := flag.Uint("port", 4242, "Port on which the Server runs")
	flag.Parse()

	absHtmlDir, err := filepath.Abs(*htmlDir)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf(":%d", *port)
	
	fileServer := http.FileServer(http.Dir(absHtmlDir))
	http.Handle("/", fileServer)
	log.Println("Serving Files from", absHtmlDir, "and listening on", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
