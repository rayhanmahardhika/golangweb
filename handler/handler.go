package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
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
	// w.Write([]byte("Welcome to home"))
	// menggunakan template untuk menampilkan HTML
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		// dev mode error
		log.Println(err.Error())
		// user friendly error
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return 
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		// dev mode error
		log.Println(err.Error())
		// user friendly error
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return 
	}

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
