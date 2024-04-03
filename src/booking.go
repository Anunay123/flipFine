package src

type Booking struct {
	BookingId int
	BookedBy  *Patient
	Slot      *Slot
}

func NewBooking(bookingId int, patient *Patient, slot *Slot) *Booking {
	slot.BookingId = bookingId
	return &Booking{
		BookingId: bookingId,
		BookedBy:  patient,
		Slot:      slot,
	}
}
