package requesthandlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"ticketstore/constants"
	"ticketstore/models"
	"ticketstore/ticketsrepo"
)

var errorrespponse models.ErrorResponse

func GetAvailableTickets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	auth := r.Header.Get("Authorization")
	if auth != constants.ClientApikey {
		w.WriteHeader(http.StatusUnauthorized)
		errorrespponse.ErrorCode = constants.ErrorCodeAuthError
		errorrespponse.ErrorMsg = constants.ErrorStringAuthError
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s GetAvailableTickets\n", r.Method)

	tickets := ticketsrepo.GetAvailableTickets()
	jsondata, _ := json.Marshal(tickets)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

	log.Println("Available Tickets returned.")
}

func PreBookTickets(w http.ResponseWriter, r *http.Request) {
	var ticketsForPrebooking models.PreBookTicketList
	var preBookingId int
	var errorcode constants.Errorcodes

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	auth := r.Header.Get("Authorization")
	if auth != constants.ClientApikey {
		w.WriteHeader(http.StatusUnauthorized)
		errorrespponse.ErrorCode = constants.ErrorCodeAuthError
		errorrespponse.ErrorMsg = constants.ErrorStringAuthError
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s PreBookTickets\n", r.Method)

	payload, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(payload, &ticketsForPrebooking)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	preBookingId, errorcode, err = ticketsrepo.PreBookTickets(ticketsForPrebooking)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = errorcode
		errorrespponse.ErrorMsg = err.Error()
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	// On success, store prebooking id and return
	for i := 0; i < len(ticketsForPrebooking.PreBookingTickets); i++ {
		ticketsForPrebooking.PreBookingTickets[i].PreBookingId = preBookingId
	}

	jsondata, _ := json.Marshal(ticketsForPrebooking)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)
}

func ConfirmPreBookTickets(w http.ResponseWriter, r *http.Request) {
	var bookingConfirmedTickets models.BookedTicketList
	var bookingId int
	var errorcode constants.Errorcodes

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	auth := r.Header.Get("Authorization")
	if auth != constants.ClientApikey {
		w.WriteHeader(http.StatusUnauthorized)
		errorrespponse.ErrorCode = constants.ErrorCodeAuthError
		errorrespponse.ErrorMsg = constants.ErrorStringAuthError
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s ConfirmPreBookTickets\n", r.Method)

	payload, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(payload, &bookingConfirmedTickets)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	bookingId, errorcode, err = ticketsrepo.ConfirmBookedTickets(bookingConfirmedTickets)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = errorcode
		errorrespponse.ErrorMsg = err.Error()
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	// On success, store Booking id and return
	for i := 0; i < len(bookingConfirmedTickets.BookedTickets); i++ {
		bookingConfirmedTickets.BookedTickets[i].BookingId = bookingId
	}

	jsondata, _ := json.Marshal(bookingConfirmedTickets)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

}

func RestorePreBookedTickets(w http.ResponseWriter, r *http.Request) {
	var ticketsPrebookedForRevert models.PreBookTicketList
	var errorcode constants.Errorcodes

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	auth := r.Header.Get("Authorization")
	if auth != constants.ClientApikey {
		w.WriteHeader(http.StatusUnauthorized)
		errorrespponse.ErrorCode = constants.ErrorCodeAuthError
		errorrespponse.ErrorMsg = constants.ErrorStringAuthError
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s RestorePreBookedTickets\n", r.Method)

	payload, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(payload, &ticketsPrebookedForRevert)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	errorcode, err = ticketsrepo.RevertPreBookedTickets(ticketsPrebookedForRevert)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = errorcode
		errorrespponse.ErrorMsg = err.Error()
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tickets Reverted"))
}

func NetsuiteRegisterEmployee(w http.ResponseWriter, r *http.Request) {
	//var ticketsForPrebooking models.PreBookTicketList
	//var preBookingId int
	//var errorcode constants.Errorcodes
	var netsuiteEmpRegister models.NetsuiteEmployeeRegister

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	auth := r.Header.Get("Authorization")
	if auth != "Bearer "+constants.ClientApikey {
		w.WriteHeader(http.StatusUnauthorized)
		errorrespponse.ErrorCode = constants.ErrorCodeAuthError
		errorrespponse.ErrorMsg = constants.ErrorStringAuthError
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Request Received : %s NetsuiteRegisterEmployee\n", r.Method)

	payload, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(payload, &netsuiteEmpRegister)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorrespponse.ErrorCode = constants.ErrorCodeBadRequest
		errorrespponse.ErrorMsg = constants.ErrorStringCodeBadRequest
		jsondata, _ := json.Marshal(errorrespponse)
		w.Write(jsondata)
		return
	}

	log.Printf("Data Received for Netsuite Registration: %v \n", netsuiteEmpRegister)
	log.Printf("\n\t\"accountHead\": \t\t\"%v\",\n\t\"eventJobCode\": \t\"%v\",\n\t\"employeeEmail\": \t\"%v\",\n\t\"ticketID\": \t\t%v,\n\t\"approvedAmount\": \t%v",
		netsuiteEmpRegister.AccountHead, netsuiteEmpRegister.EventJobCode, netsuiteEmpRegister.EmployeeEmail,
		netsuiteEmpRegister.TicketID, netsuiteEmpRegister.ApprovedAmount)

	log.Print("Registration Cucessful")

	jsondata, _ := json.Marshal(netsuiteEmpRegister)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

}
