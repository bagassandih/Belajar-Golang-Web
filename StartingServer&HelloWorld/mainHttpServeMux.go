package main
import (
	"fmt"
	"net/http"
	)


func helloHandler(w http.ResponseWriter, r *http.Request){
	w.Write( []byte("hello World, Saya sedang belajar golang web") )
}

func bagasHandler(w http.ResponseWriter, r *http.Request){
	w.Write( []byte("Pusing juga ya ges ya") )

}	

func main (){
	// surfing dengan mux
	Mux := http.NewServeMux()

	Mux.HandleFunc("/hello", helloHandler)
	Mux.HandleFunc("/bagas", bagasHandler)


	fmt.Println("Starting web on 8080")

	err:= http.ListenAndServe(":8080", Mux) // :8080 sama saja localhost:8080
	log.Fatal(err)
}

