package requesthandlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"ticketstore/constants"
	"ticketstore/models"
)

var errorrespponse models.ErrorResponse

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

	auth := r.Header.Get("X-API-KEY")
	if auth != constants.ClientApikey {
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
		strconv.FormatInt(int64(netsuiteEmpRegister.TicketID), 10), netsuiteEmpRegister.ApprovedAmount)

	log.Print("Registration Sucessful")

	jsondata, _ := json.Marshal(netsuiteEmpRegister)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)

}
