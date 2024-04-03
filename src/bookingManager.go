package src

import "errors"

type BookingManager struct {
	BookingsMap     map[int]*Booking
	PatientsBooking map[string]*Booking
	DoctorsBooking  map[string]*Booking
	Counter         int
}

func (bookingManager *BookingManager) MakeBooking(patient *Patient, slot *Slot) (int, error) {
	if bookingManager != nil {

		if bookingManager.isSlotAlreadyBooked(patient.Name, slot) {
			return -1, errors.New("booking can't be made, select a different time slot")
		}

		bookingManager.Counter++
		booking := NewBooking(bookingManager.Counter, patient, slot)

		bookingManager.BookingsMap[bookingManager.Counter] = booking

		return bookingManager.Counter, nil

	}

	return -1, errors.New("booking manager unavailable")
}

func (bookingManager *BookingManager) isSlotAlreadyBooked(patientName string, slot *Slot) bool {
	if bookingManager != nil {
		for _, bookings := range bookingManager.BookingsMap {
			if bookings.BookedBy.Name == patientName && bookings.Slot.StartTime == slot.StartTime && bookings.Slot.EndTime == bookings.Slot.EndTime {
				return true
			}
		}

		return false
	}

	return false
}

func (bookingManager *BookingManager) CancelBooking(bookingId int) error {
	if bookingManager != nil {

		if booking, exists := bookingManager.BookingsMap[bookingId]; exists {

			delete(bookingManager.BookingsMap, bookingId)

			if len(booking.Slot.WaitList) > 0 && booking.Slot.Front < len(booking.Slot.WaitList) {
				waitListedPatient := booking.Slot.WaitList[booking.Slot.Front]
				booking.Slot.Front += 1

				if _, err := bookingManager.MakeBooking(waitListedPatient, booking.Slot); err != nil {
					return err
				}

				return nil

			}

			return nil
		}

		return errors.New("bookingId doesn't exist")

	}

	return errors.New("booking manager unavailable")
}
