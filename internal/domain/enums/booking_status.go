package enums

type BookingStatus string

const (
    BookingPending   BookingStatus = "pending"
    BookingAccepted  BookingStatus = "accepted"
    BookingArriving  BookingStatus = "arriving"
    BookingStarted   BookingStatus = "started"
    BookingCompleted BookingStatus = "completed"
    BookingCancelled BookingStatus = "cancelled"
)