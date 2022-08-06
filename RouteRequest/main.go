package main

import (
	"fmt"
	"net/http"
	)


func HandleIndex (w http.ResponseWriter, r *http.Request){
	pesan := "Halo Golang, Saya sedang mempelajari kamu"
	if r.URL.Path != "/" { // handling jika memasukan route yang tdk didaftar
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(pesan))
}

func main (){
	server := http.NewServeMux()
	address:= "localhost:8080"

	server.HandleFunc("/", HandleIndex) 
	//  "/" halaman depan/index juga tetapi juga menampilkan handle yang sama untuk semua rute yg tidak terdaftar
	server.HandleFunc("/index", HandleIndex) // rute hanya "/index" saja


	
	fmt.Println("Server starting at", address)

	err := http.ListenAndServe(address, server) // start dg adress, nama http multiplexer
	if err!=nil{ // error handler
		fmt.Print("Gagal starting server, Cek kembali\n", err.Error() )
	}

}