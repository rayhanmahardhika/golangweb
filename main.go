package main

import (
	"log"
	"net/http"
	"golangweb/handler"
)

func main() {
	// Ada 3 cara membuat Handler function untuk routing
	// 1. Membuat function terpisah [INI PALING BANYAK DIGUNAKAN]
	// 2. Membuat Closure function (function yang di assign ke variabel)
	// 3. membuat anonimus function (function tanpa nama yang langsung didefinisikan di proses routing) [INI JUGA LUMAYAN SERING DIPAKE]

	// routing
	mux := http.NewServeMux()

	// membuat closure function
	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	}

	// jika pattern seperti ini maka semua yang tidak terdaftar akan terhandler ini
	mux.HandleFunc("/", handler.HomeHandler) // root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	// membuat anonimus function
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Page"))
	})

	// loading static files (js, css, or anything)
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Starting web on port 8080")

	// serving web pada port 8080
	err := http.ListenAndServe(":8080", mux)
	// jika terdapat error
	log.Fatal(err)
}
