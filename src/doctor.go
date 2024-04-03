package src

type Doctor struct {
	Name       string
	Speciality string
	Rating     int
}

func NewDoctor(name string, speciality string) *Doctor {
	return &Doctor{
		Name:       name,
		Speciality: speciality,
	}
}

//func (doctor *Doctor) AddSlots(slots [][]int) error {
//
//	if doctor != nil {
//
//		var err error
//		for _, slot := range slots {
//			slotString := generateSlotString(slot)
//
//			if _, exists := doctor.AvailableSlots[slotString]; exists {
//				continue
//			}
//
//			//if slot[1]-slot[0] > 30 {
//			//	err = errors.New(fmt.Sprintf("Sorry %s, %s slot dropped"))
//			//}
//
//			doctor.AvailableSlots[slotString] = &Slot{
//				StartTime: slot[0],
//				EndTime:   slot[1],
//				Display:   slotString,
//			}
//		}
//
//		return nil
//	}
//
//	return errors.New("doctor unavailable")
//}
