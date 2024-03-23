package server

import (
	"fmt"
	"net/http"
)
// membuat route 
func Route() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "ini adalah halaman utama")
	})

	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "ini adalah halaman about")
	})

	mux.HandleFunc("/portofolio", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "ini adalah halaman portofolio")
	})

	// membuat Server

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}