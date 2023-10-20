package main

import (
	"log"
	"net/http"
	"ticketstore/requesthandlers"
)

func main() {
	startServing()
}

func startServing() {

	registerRouts()

	log.Println("Starting Netsuite Mock Server on port : 8050")
	err := http.ListenAndServe(":8050", nil)

	if err != nil {
		log.Println(err)
		return
	}
}

func registerRouts() {
	http.HandleFunc("/netsuiteRegisterEmployee", requesthandlers.NetsuiteRegisterEmployee)

}
