package main

import (
	"fmt"
	"net/http"
	)


func HandleIndex (w http.ResponseWriter, r *http.Request){
	pesan := "Halo Golang, Saya sedang mempelajari kamu"
		if r.URL.Path != "/" && r.URL.Path != "/index" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(pesan))
}

func main (){
	server := http.NewServeMux()
	address:= "localhost:8080"

	// Membuat Handler didalam HandleFunc dengan Closure(membuat func didalam variabel, harus di func main)
	aboutHandler := func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Ini adalah halaman about"))
	}


	// daftarkan rutenya dari Closure
	server.HandleFunc("/", HandleIndex) 
	server.HandleFunc("/index", HandleIndex) 
	server.HandleFunc("/about", aboutHandler) 
	
	server.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Ini adalah halaman profile"))
	} ) // func tanpa nama, atau anonim func


	
	fmt.Println("Server starting at", address)

	err := http.ListenAndServe(address, server)
	if err!=nil{ // error handler
		fmt.Print("Gagal starting server, Cek kembali\n", err.Error() )
	}

}