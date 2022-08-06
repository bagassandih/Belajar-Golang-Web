package main

import (
	"fmt"
	"net/http"
	"time"
	)


func HandlerIndex (w http.ResponseWriter, r *http.Request) {
	pesan := "Hallo"
	w.Write([]byte(pesan))
}

func BagasIndex (w http.ResponseWriter, r *http.Request) {
	pesan := "Hallo, Bagas disini. Pusing juga nie"
	w.Write([]byte(pesan))
}

func main(){
	http.HandleFunc("/", HandlerIndex)
	http.HandleFunc("/index", HandlerIndex)
	http.HandleFunc("/bagas", BagasIndex)

	address:= "localhost:8080"
	fmt.Println("Starting server at", address)	

	server:= new(http.Server)	
	server.Addr = address
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10

	err:= server.ListenAndServe()
	if err!= nil{
		fmt.Println(err.Error())
	}
}


/*
Kelebihan menggunakan http.Server salah satunya adalah kemampuan untuk mengubah beberapa konfigurasi default web server Go.
Contoh, pada kode timeout untuk read request dan write request di ubah menjadi 10 detik.
*/