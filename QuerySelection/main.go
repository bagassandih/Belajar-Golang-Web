package main

import (
	"fmt"
	"strconv"
	"net/http"
	)


func HandleIndex (w http.ResponseWriter, r *http.Request){
	pesan := "Halo Golang, Saya sedang mempelajari kamu"
	if r.URL.Path != "/" { 
		http.NotFound(w, r)
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

func main (){
	server := http.NewServeMux()
	address:= "localhost:8080"

	server.HandleFunc("/", HandleIndex) 
	server.HandleFunc("/index", HandleIndex) 
	server.HandleFunc("/product", HandleProduct)

	
	fmt.Println("Server starting at", address)

	err := http.ListenAndServe(address, server)
	if err!=nil{ // error handler
		fmt.Print("Gagal starting server, Cek kembali\n", err.Error() )
	}

}