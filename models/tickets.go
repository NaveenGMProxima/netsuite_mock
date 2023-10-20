package models

import "ticketstore/constants"

type ErrorResponse struct {
	ErrorCode constants.Errorcodes `json:"errorcode"`
	ErrorMsg  string               `json:"errormessage"`
}

type Ticket struct {
	TicketID     int     `json:"ticketid"` // This is the unique key in our system
	TheatreName  string  `json:"theatrename"`
	ShowDateTime string  `json:"showdatetime"`
	TicketCost   float64 `json:"cost"`
}

type PreBookTicketList struct {
	PreBookingTickets []PreBookedTicket `json:"prebookticketlist"`
}

type PreBookedTicket struct {
	PreBookingId     int    `json:"preboookingid"`
	TicketId         int    `json:"ticketid"`
	CustomerMobileNo string `json:"customermobile"`
}

type BookedTicketList struct {
	BookedTickets []BookedTicket `json:"bookedtickets"`
}

type BookedTicket struct {
	BookingId        int    `json:"boookingid"`
	TicketId         int    `json:"ticketid"`
	CustomerMobileNo string `json:"customermobile"`
	PaymentReference string `json:"paymentreference"`
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
