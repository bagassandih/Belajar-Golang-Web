package handler

import (
	//perhatikan package 
	"fmt"
	"path" // untuk suport template parsing files
	"strconv"
	"net/http"
	"html/template" // untuk parsing file html dan eksekusi
	)


func HandleIndex (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" { 
		http.NotFound(w, r)
		return
	}

	// buat var (html) dengan cek error, isi variabel html adalah mengakses template untuk parse file
	// template.ParseFiles(path) error
	// (path) => path.Join("namafolder", "namafiles")

	html, err:= template.ParseFiles(path.Join("views", "index.html")) // untuk parsing file path join nama folder, nama file
	if err != nil { // jika terjadi error
		fmt.Println(err) // ini menampilkan mana yang error untuk developer (spesifik error)
		http.Error(w, "Terjadi Error, harap sabar!", http.StatusInternalServerError) // error untuk user
		return
	}

	// Eksekusi html dengan (variabeltemplate).Execute(wr / w, data tipebebas) error
	// html.Execute(w, nil) wr / w ambil saja parameter w, dan nil jika tidak ada
	err = html.Execute(w, nil) 
	if err != nil {
		fmt.Println(err) // ini menampilkan mana yang error untuk developer
		http.Error(w, "Terjadi Error, harap sabar!", http.StatusInternalServerError) // error untuk user
		return
	}
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
