package handler

import (
	"fmt"
	"strconv"
	"net/http"
	"html/template"
	)


func HandleIndex (w http.ResponseWriter, r *http.Request){
	pesan := "Halo Golang, Saya sedang mempelajari kamu"
	if r.URL.Path != "/" { 
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte(pesan))
	_, err:= template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(pesan))
}


func HandleProduct(w http.ResponseWriter, r*http.Request){
	pesan := "Product Page"

	id := r.URL.Query().Get("id") // ambil query string id dg get => localhost:8080/product?id= 
	idInt, err := strconv.Atoi(id) // conversi ke var idInt dan cek error

	if err!= nil || idInt < 1 { // kondisi error
		http.NotFound(w, r) // jika error tampilkan notfound
		return
	}

	// w.Write([]byte(pesan)) // tidak bisa jika int
	fmt.Fprintf(w, pesan + ": %d", idInt) // pakai Fprint

}
