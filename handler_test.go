package belajar_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//tempat logic web
		fmt.Fprint(writer, "hello world")
	}

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Menggunakan banyak endpoint menggunakan func servemux
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Selamat datang di web akooh")
	})
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hi broo wasuuup ma nugga")
	})

	// Menggunakan url pattern dengan cara menambahkan garis miring dibelakang
	mux.HandleFunc("/haloo/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "menggunakan url pattern")
	})

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Test Request
func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Body)
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.URL)
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}
