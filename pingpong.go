package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Ping struct{}

func (p Ping) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf("pinged from %s", r.RemoteAddr))
	fmt.Fprintln(w, fmt.Sprintf("pong back to %s", r.RemoteAddr))
}

func main() {
	addrFlag := flag.String("addr", "0.0.0.0", "web server host/ip")
	portFlag := flag.Uint("port", 8000, "web server port")
	flag.Parse()

	endpoint := "ping"
	http.Handle("/"+endpoint, Ping{})

	listenTo := fmt.Sprintf("%s:%d", *addrFlag, *portFlag)
	fmt.Fprintf(os.Stderr, "listening on %s/%s\n", listenTo, endpoint)
	http.ListenAndServe(listenTo, nil)
}
