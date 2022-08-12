package main
/*
file html ada di folder views
*/
import (
	"fmt"
	"net/http"
	"golangweb/handler"
	)

func main (){
	server := http.NewServeMux()
	address:= "localhost:8080"

	server.HandleFunc("/", handler.HandleIndex) 
	server.HandleFunc("/index", handler.HandleIndex) 
	server.HandleFunc("/product", handler.HandleProduct)

	
	fmt.Println("Server starting at", address)

	err := http.ListenAndServe(address, server)
	if err!=nil{ // error handler
		fmt.Print("Gagal starting server, Cek kembali\n", err.Error() )
	}

}