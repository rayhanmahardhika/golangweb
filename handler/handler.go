package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

// agar menjadi public ini semua nama function harus di awali dengan 
// huruf kapital

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// ini rule solusi untuk root route
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to home"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya sedang be;ajar golang web"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from nintendo"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Product Page : " + id))

	fmt.Fprintf(w, "Product page : %d", idNumb)
}
