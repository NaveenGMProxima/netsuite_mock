package main

import (
	"log"
	"net/http"
	"ticketstore/constants"
	"ticketstore/requesthandlers"
	"ticketstore/ticketsrepo"
)

func main() {
	generateTickets(constants.TicketsCount)
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
	http.HandleFunc("/GetAvailableTickets", requesthandlers.GetAvailableTickets)
	http.HandleFunc("/PreBookTickets", requesthandlers.PreBookTickets)
	http.HandleFunc("/ConfirmPreBookTickets", requesthandlers.ConfirmPreBookTickets)
	http.HandleFunc("/RestorePreBookedTickets", requesthandlers.RestorePreBookedTickets)

	http.HandleFunc("/netsuiteRegisterEmployee", requesthandlers.NetsuiteRegisterEmployee)

}

// Genetare Tickets for a show
func generateTickets(numberOfTickets int) {

	theatreName := "Abhinav Theatre - Screen 1"
	showDateTime := "22-Apr-2021 16:00"
	ticketCost := 176.00

	ticketsrepo.GenerateTickets(theatreName, showDateTime, ticketCost, numberOfTickets)
}
