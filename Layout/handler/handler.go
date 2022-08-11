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

	html, err:= template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html")) // untuk parsing file path join nama folder, nama file
	if err != nil { // jika terjadi error
		fmt.Println(err) // ini menampilkan mana yang error untuk developer (spesifik error)
		http.Error(w, "Terjadi Error, harap sabar!", http.StatusInternalServerError) // error untuk user
		return
	}

	// data yang ingin di eksekusi atau di tampilkan di index
	data := map[string]interface{}{ // interface agar bisa tipe apa saja (stirng, int, char, dll)
		"title" : "My First Golang Web",
		"content": "Belajar golang Web Otodidak, ini adalah web Dinamis pertama saya",
	}

	// Eksekusi html dengan (variabeltemplate).Execute(wr / w, data tipebebas) error
	// html.Execute(w, nil) wr / w ambil saja parameter w, dan nil jika tidak ada
	err = html.Execute(w, data) // coba cek file index.html
	if err != nil {
		fmt.Println(err) // ini menampilkan mana yang error untuk developer
		http.Error(w, "Terjadi Error, harap sabar!", http.StatusInternalServerError) // error untuk user
		return
	}
}


func HandleProduct(w http.ResponseWriter, r*http.Request){
	pesan := "Product Page with Id"

	id := r.URL.Query().Get("id") // ambil query string id dg get => localhost:8080/product?id= 
	idInt, err := strconv.Atoi(id) // conversi ke var idInt dan cek error

	if err!= nil || idInt < 1 { // kondisi error
		http.NotFound(w, r) // jika error tampilkan notfound
		return
	}

	data := map[string]interface{}{
		"id": idInt,
		"text": pesan,
	}

	html, err:= template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html") ) 
	if err!= nil{
		fmt.Println(err)
		http.Error(w, "Terjadi Error, harap sabar!", http.StatusInternalServerError)
		return
	}

	err = html.Execute(w, data)
	if err!= nil{
		fmt.Println(err)
		http.Error(w, "Terjadi Error, harap sabar!", http.StatusInternalServerError)
		return
	}

}
