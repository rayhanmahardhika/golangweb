package handler

import (
	"golangweb/entity"
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
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		// dev mode error
		log.Println(err.Error())
		// user friendly error
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// menggunakan slices of struct sebagai parsing data ke views
	data := []entity.Product{
		{
			ID : 1,
			Name : "Mobilio",
			Price : 220000000,
			Stock : 11,
		},
		{
			ID : 2,
			Name : "Xpander",
			Price : 270000000,
			Stock : 8,
		},
		{
			ID : 3,
			Name : "Pajero Sport",
			Price : 400000000,
			Stock : 1,
		},
	}

	err = tmpl.Execute(w, data)
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
	// interface seperti ini memungkinkan bisa mengisi value dengan berbagai type data
	data := map[interface{}]interface{}{
		"content": idNumb,
	}

	// menggunakan template untuk menampilkan HTML
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		// dev mode error
		log.Println(err.Error())
		// user friendly error
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		// dev mode error
		log.Println(err.Error())
		// user friendly error
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}
