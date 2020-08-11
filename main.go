package main

import (
	"./app"
	"log"
	"net/http"
)

func main() {
	log.Println("http://localhost:8080")
	http.HandleFunc("/provinsi/", app.Provinsi)
	http.HandleFunc("/kabupaten/", app.Kabupaten)
	http.HandleFunc("/kecamatan/", app.Kecamatan)
	http.HandleFunc("/kelurahan/", app.Kelurahan)
	http.HandleFunc("/data-wilayah/", app.Wilayah)
	log.Fatal(http.ListenAndServe(":8080", nil))
}