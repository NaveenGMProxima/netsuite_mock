package ticketsrepo

import (
	"fmt"
	"sync"
	"ticketstore/constants"
	"ticketstore/models"
)

var availabletickets []models.Ticket
var reservedTickets []models.Ticket
var soldoutTickets []models.Ticket
var prebookedtickets []models.PreBookedTicket
var bookedtickets []models.BookedTicket
var prebookingIdGenerator = 1
var bookingIdGenerator = 1

var lock = sync.RWMutex{}

// Getters
func GetAvailableTickets() []models.Ticket {
	var tickets []models.Ticket

	lock.RLock()
	tickets = availabletickets
	lock.RUnlock()

	return tickets
}

func isTicketBooked(pTicketID int) bool {
	for i := 0; i < len(bookedtickets); i++ {
		if bookedtickets[i].TicketId == pTicketID {
			return true
		}
	}
	return false
}

func isAnyOfTicketBooked(pTicketsForPrebooking models.PreBookTicketList) bool {

	for i := 0; i < len(pTicketsForPrebooking.PreBookingTickets); i++ {
		if isTicketBooked(pTicketsForPrebooking.PreBookingTickets[i].TicketId) {
			return true
		}
	}
	return false
}

func isTicketPreBooked(pTicketID int) bool {
	for i := 0; i < len(prebookedtickets); i++ {
		if prebookedtickets[i].TicketId == pTicketID {
			return true
		}
	}
	return false
}

func isTicketPreBookedBy(pTicketID int, pCustomerMobileNo string) bool {
	for i := 0; i < len(prebookedtickets); i++ {
		if (prebookedtickets[i].TicketId == pTicketID) && (prebookedtickets[i].CustomerMobileNo == pCustomerMobileNo) {
			return true
		}
	}
	return false
}

func isAllTicketsPreBooked(pBookingConfirmedTickets models.BookedTicketList) bool {
	for i := 0; i < len(pBookingConfirmedTickets.BookedTickets); i++ {
		if !isTicketPreBookedBy(pBookingConfirmedTickets.BookedTickets[i].TicketId, pBookingConfirmedTickets.BookedTickets[i].CustomerMobileNo) {
			return false
		}
	}
	return true
}

func isAnyOfTicketPreBooked(pTicketsForPrebooking models.PreBookTicketList) bool {

	for i := 0; i < len(pTicketsForPrebooking.PreBookingTickets); i++ {
		if isTicketPreBooked(pTicketsForPrebooking.PreBookingTickets[i].TicketId) {
			return true
		}
	}
	return false
}

func markTicketAsReserved(pTicketID int) {
	for i := 0; i < len(availabletickets); i++ {
		if availabletickets[i].TicketID == pTicketID {
			reservedTickets = append(reservedTickets, availabletickets[i])
			availabletickets = append(availabletickets[:i], availabletickets[i+1:]...)
			return
		}
	}
}

func markTicketAsAvaialable(pTicketID int) {
	for i := 0; i < len(reservedTickets); i++ {
		if reservedTickets[i].TicketID == pTicketID {
			availabletickets = append(availabletickets, reservedTickets[i])
			reservedTickets = append(reservedTickets[:i], reservedTickets[i+1:]...)
			break
		}
	}

	for i := 0; i < len(prebookedtickets); i++ {
		if prebookedtickets[i].TicketId == pTicketID {
			prebookedtickets = append(prebookedtickets[:i], prebookedtickets[i+1:]...)
			break
		}
	}
}

func markTicketAsBookingConfirmed(pTicketID int) {
	for i := 0; i < len(reservedTickets); i++ {
		if reservedTickets[i].TicketID == pTicketID {
			soldoutTickets = append(soldoutTickets, reservedTickets[i])
			reservedTickets = append(reservedTickets[:i], reservedTickets[i+1:]...)
			break
		}
	}

	for i := 0; i < len(prebookedtickets); i++ {
		if prebookedtickets[i].TicketId == pTicketID {
			prebookedtickets = append(prebookedtickets[:i], prebookedtickets[i+1:]...)
			break
		}
	}
}

func PreBookTickets(pTicketsForPrebooking models.PreBookTicketList) (int, constants.Errorcodes, error) {
	var prebookid int
	lock.RLock()
	// Check if tickets are already booked
	if isAnyOfTicketBooked(pTicketsForPrebooking) {
		lock.RUnlock()
		return -1, constants.ErrorCodeTicktNotAvailable, fmt.Errorf(constants.ErrorStringNotAvailableForBooking)
	}

	// Check if tickets are in prebooking
	if isAnyOfTicketPreBooked(pTicketsForPrebooking) {
		lock.RUnlock()
		return -1, constants.ErrorCodeTicktNotAvailable, fmt.Errorf(constants.ErrorStringNotAvailableForBooking)
	}
	lock.RUnlock()

	// Tickets are available for booking
	lock.Lock()
	prebookid = prebookingIdGenerator
	for i := 0; i < len(pTicketsForPrebooking.PreBookingTickets); i++ {
		pTicketsForPrebooking.PreBookingTickets[i].PreBookingId = prebookid
		markTicketAsReserved(pTicketsForPrebooking.PreBookingTickets[i].TicketId)
		prebookedtickets = append(prebookedtickets, pTicketsForPrebooking.PreBookingTickets[i])
	}
	lock.Unlock()

	prebookingIdGenerator++

	return prebookid, constants.ErrorNoErrors, nil
}

func ConfirmBookedTickets(pBookingConfirmedTickets models.BookedTicketList) (int, constants.Errorcodes, error) {
	var bookingid int

	lock.RLock()
	// Check if tickets are in prebooking
	if !isAllTicketsPreBooked(pBookingConfirmedTickets) {
		lock.RUnlock()
		return -1, constants.ErrorCodeTicktNotAvailable, fmt.Errorf(constants.ErrorStringNotAvailableForBooking)
	}
	lock.RUnlock()

	lock.Lock()
	bookingid = bookingIdGenerator
	for i := 0; i < len(pBookingConfirmedTickets.BookedTickets); i++ {
		pBookingConfirmedTickets.BookedTickets[i].BookingId = bookingid
		markTicketAsBookingConfirmed(pBookingConfirmedTickets.BookedTickets[i].TicketId)
		bookedtickets = append(bookedtickets, pBookingConfirmedTickets.BookedTickets[i])
	}
	lock.Unlock()

	bookingIdGenerator++

	return bookingid, constants.ErrorNoErrors, nil
}

func RevertPreBookedTickets(pTicketsPrebookedForRevert models.PreBookTicketList) (constants.Errorcodes, error) {

	lock.Lock()

	for i := 0; i < len(pTicketsPrebookedForRevert.PreBookingTickets); i++ {
		markTicketAsAvaialable(pTicketsPrebookedForRevert.PreBookingTickets[i].TicketId)
	}

	lock.Unlock()

	return constants.ErrorNoErrors, nil
}

// Helper functions
func GenerateTickets(pTheatreName string, pShowDateTime string, pTicketCost float64, pNumberofTickets int) {
	var ticket models.Ticket

	for i := 1; i <= pNumberofTickets; i++ {
		ticket = models.Ticket{
			TicketID:     i,
			TheatreName:  pTheatreName,
			ShowDateTime: pShowDateTime,
			TicketCost:   pTicketCost,
		}

		availabletickets = append(availabletickets, ticket)
	}

}
