package models

import "ticketstore/constants"

type ErrorResponse struct {
	ErrorCode constants.Errorcodes `json:"errorcode"`
	ErrorMsg  string               `json:"errormessage"`
}

/*
{
   "accountHead": "travel_fund",
   "eventJobCode": "Cassandra Summit 202 2341205",
   "employeeEmail": "dqualls@contractor.linuxfoundation.org",
   "ticketID": 1566074986,
   "approvedAmount": 1234.56
 }
*/
type NetsuiteEmployeeRegister struct {
	AccountHead    string  `json:"accountHead"`
	EventJobCode   string  `json:"eventJobCode"`
	EmployeeEmail  string  `json:"employeeEmail"`
	TicketID       float64 `json:"ticketID"`
	ApprovedAmount float64 `json:"approvedAmount"`
}
