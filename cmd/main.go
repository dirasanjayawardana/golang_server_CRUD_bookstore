package main

import (
	"golang_server_bookstore/internals/routes"
	"golang_server_bookstore/pkg"
	"log"
)

// Depedency Injection (DI) --> lebih pasti dalam validasi error

func main() {

	// Inisialisasi Database (install sqlx dan driver go sql driverr)
	_, error := pkg.InitMySql()
	if error != nil {
		log.Fatal(error) // log.Fatal --> ketika terjadi error akan langsung memberhentikan program
		// return
	}

	// Inisialisasi Router
	router := routes.InitRouter()

	// Inisialisasi Server
	server := pkg.InitServer(router)

	// jalankan server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
