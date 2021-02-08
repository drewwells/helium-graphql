package mock

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"
)

func Handle(pattern string, handle http.HandlerFunc) {

	handleFuncs[pattern] = handle
}

var handleFuncs = make(map[string]http.HandlerFunc)

func NewServer(t *testing.T) (*http.ServeMux, net.Listener) {

	mux := http.NewServeMux()
	for p, h := range handleFuncs {
		mux.HandleFunc(p, h)
	}

	lis, err := net.Listen("tcp4", ":0")
	if err != nil {
		log.Fatal(err)
	}
	root := "http://" + lis.Addr().String()
	fmt.Println("starting mock server:", root)

	go http.Serve(lis, mux)
	return mux, lis
}
