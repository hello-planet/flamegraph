package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"flamegraph/handler"
)

//port
const hostPort = ":8080"

func main() {
	flag.Parse()

	// Listen
	http.HandleFunc("/stats", handler.DealHandler(handler.Msg))
	http.HandleFunc("/", index)

	fmt.Println("Server started on", hostPort)
	if err := http.ListenAndServe(hostPort, nil); err != nil {
		log.Fatalf("HTTP Server Failed: %v", err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, `<a href="%v">%v</a>`, "stats", "check")
}
