package src

type Slot struct {
	StartTime int
	EndTime   int
	Display   string
	Doctor    *Doctor
	Status    string
	BookingId int
	WaitList  []*Patient
	Front     int
}

func NewSlot(startTime int, endTime int, doctor *Doctor, displayString string) *Slot {
	return &Slot{
		StartTime: startTime,
		EndTime:   endTime,
		Doctor:    doctor,
		Display:   displayString,
		Status:    "Available",
		WaitList:  make([]*Patient, 0),
	}
}
