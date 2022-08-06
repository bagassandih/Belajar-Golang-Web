package main

import "fmt"
import "net/http"
/*
sumber : dasarpemrogramangolang.novalagung.com

Route handler atau handler atau parameter kedua fungsi http.HandleFunc(), 
adalah sebuah fungsi dengan ber-skema func (ResponseWriter, *Request)

Function Handler Wajib!
Parameter ke-1 bertipe http.ResponseWrite
Parameter ke-2 bertipe *http.Request
Fungsi dengan struktur itu diperlukan oleh http.HandleFunc untuk keperluan penanganan request ke rute yang ditentukan

*/

func handlerIndex(w http.ResponseWriter, r *http.Request) { // handler untuk route Index
	var message = "Welcome"
	w.Write([]byte(message))

	/*Method Write() milik parameter pertama (yang bertipe http.ResponseWrite), 
	digunakan untuk meng-output-kan nilai balik data. 
	Argumen method adalah data yang ingin dijadikan output, 
	ditulis dalam bentuk []byte.

	Pada contoh ini, data yang akan kita tampilkan bertipe string, 
	maka perlu dilakukan casting dari string ke []byte. 
	Contohnya bisa dilihat seperta pada kode di atas, di bagian w.Write([]byte(message)*/
}

func handlerHello(w http.ResponseWriter, r *http.Request) { // Handler utk route Hello
	var message = "Hello world!"
	w.Write([]byte(message))
}

func main() {
	// starting server 
	http.HandleFunc("/", handlerIndex) // route index(halaman depan)
	http.HandleFunc("/index", handlerIndex) // route index
	http.HandleFunc("/hello", handlerHello) // route hello
	// http.HandleFunc() digunakan untuk routing. Parameter pertama adalah rute dan parameter ke-2 adalah handler-nya

	var address = "localhost:9000" // alamat localhost dan portnya (9000, 8080, dll)
	fmt.Printf("server started at %s\n", address) // menampilkan status di console
	err := http.ListenAndServe(address, nil) // start
	if err != nil { // cek jika ada error
		fmt.Println(err.Error())

	/*
	Fungsi http.ListenAndServe() bersifat blocking, 
	menjadikan semua statement setelahnya tidak akan dieksekusi, sebelum di-stop.

	Fungsi ini mengembalikan nilai balik ber-tipe error. 
	Jika proses pembuatan web server baru gagal, maka kita bisa mengetahui root-cause nya apa
	*/

	}
}
