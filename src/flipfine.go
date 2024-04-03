package src

import (
	"errors"
	"fmt"
	"strconv"
)

type FlipFine struct {
	DoctorService  *DoctorService
	PatientService *PatientService
	BookingManager *BookingManager
}

func NewFlipFine() *FlipFine {
	return &FlipFine{
		DoctorService: NewDoctorService(),
		PatientService: &PatientService{
			PatientMap: make(map[string]*Patient),
		},
		BookingManager: &BookingManager{
			BookingsMap:     make(map[int]*Booking),
			PatientsBooking: make(map[string]*Booking),
			DoctorsBooking:  make(map[string]*Booking),
		},
	}
}

func (ff *FlipFine) RegisterDoctor(name string, speciality string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if ff != nil {
		return ff.DoctorService.RegisterDoctor(name, speciality)
	}

	return errors.New("service unavailable")
}

func (ff *FlipFine) RegisterPatient(name string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if ff != nil {
		return ff.PatientService.AddPatient(name)
	}

	return errors.New("service unavailable")
}

func (ff *FlipFine) AddSlots(name string, slots [][]int) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if ff != nil {
		return ff.DoctorService.AddSlot(name, slots)
	}

	return errors.New("service unavailable")
}

func (ff *FlipFine) Display(speciality string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if ff != nil {
		slotList, err := ff.DoctorService.DisplaySlots(speciality)
		if err != nil {
			return err
		}

		for _, slot := range slotList {
			displayString := fmt.Sprintf("%s:  %s:  %s", slot.Doctor.Name, slot.Display, slot.Status)
			fmt.Println(displayString)
		}

		return nil

	}

	return errors.New("service unavailable")
}

func (ff *FlipFine) MakeBooking(doctorName string, patientName string, startTime int, endTime int) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if ff != nil {

		if doctor, exists := ff.DoctorService.DoctorRepository[doctorName]; exists && doctor != nil {

			if patient, pExists := ff.PatientService.PatientMap[patientName]; pExists && patient != nil {

				key := fmt.Sprintf("%s_%s_%s", doctor.Name, strconv.Itoa(startTime), strconv.Itoa(endTime))

				if slot, sExists := ff.DoctorService.SlotManager.Slots[key]; sExists && slot != nil {
					if slot.Status == "Available" {
						slot.Status = "Booked"
						return ff.BookingManager.MakeBooking(patient, slot)
					} else if slot.Status == "Booked" {
						slot.WaitList = append(slot.WaitList, patient)
						return slot.BookingId, errors.New("slot already booked added to waitlist")
					}

				}
			}
		}

		return -1, errors.New("booking failure")

	}

	return -1, errors.New("service unavailable")
}

func (ff *FlipFine) CancelBooking(bookingID int) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if ff != nil {
		return ff.BookingManager.CancelBooking(bookingID)
	}

	return errors.New("service unavailable")
}

func (ff *FlipFine) ViewBookingsDoctor(doctorName string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if ff != nil {
		for _, booking := range ff.BookingManager.BookingsMap {
			if booking.Slot.Doctor.Name == doctorName {
				fmt.Println(booking.Slot.Display)
			}
		}

		return nil
	}

	return errors.New("service unavailable")
}

func (ff *FlipFine) ViewBookingsPatient(patientName string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if ff != nil {
		for _, booking := range ff.BookingManager.BookingsMap {
			if booking.BookedBy.Name == patientName {
				fmt.Println(booking.Slot.Display)
			}
		}

		return nil
	}

	return errors.New("service unavailable")
}
