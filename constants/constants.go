package constants

const TicketsCount int = 100

type Errorcodes = int

const ClientApikey string = "abcd@12345"

// Error codes
const (
	ErrorNoErrors              = 999
	ErrorCodeBadRequest        = 1001
	ErrorCodeTicktNotAvailable = 1002
	ErrorCodeAuthError         = 1003
)

// Error string
const (
	ErrorStringNoErrors               = "Success"
	ErrorStringCodeBadRequest         = "Bad Request"
	ErrorStringNotAvailableForBooking = "not available for booking"
	ErrorStringAuthError              = "Not authorized"
)
