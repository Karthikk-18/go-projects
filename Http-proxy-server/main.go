package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func startBackend() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Backend received request:", r.URL.Path)
		fmt.Fprintf(w, "Hello from backend (9090)")
	})

	fmt.Println("Backend running on :9090")
	http.ListenAndServe(":9090", mux)
}

func startProxy() {
	target, _ := url.Parse("http://localhost:9090")
	proxy := httputil.NewSingleHostReverseProxy(target)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Proxy received:", r.URL.Path)

		proxy.ServeHTTP(w, r)
	})

	fmt.Println(" Proxy running on :8080")
	http.ListenAndServe(":8080", mux)
}

func main() {
	go startBackend()
	startProxy()
}