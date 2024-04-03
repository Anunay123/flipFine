package src

import (
	"errors"
	"fmt"
)

type DoctorService struct {
	DoctorRepository      map[string]*Doctor
	SpecialityToDoctorMap map[string]IRankingService
	SlotManager           *SlotManager
}

func NewDoctorService() *DoctorService {
	return &DoctorService{
		DoctorRepository:      make(map[string]*Doctor),
		SpecialityToDoctorMap: make(map[string]IRankingService),
		SlotManager:           NewSlotManager(),
	}
}

func (doctorService *DoctorService) RegisterDoctor(name string, speciality string) error {
	if doctorService != nil {
		if _, exists := doctorService.DoctorRepository[name]; exists {
			return errors.New("doctor already exists")
		}

		doctorService.DoctorRepository[name] = NewDoctor(name, speciality)

		if _, exists := doctorService.SpecialityToDoctorMap[speciality]; !exists {
			doctorService.SpecialityToDoctorMap[speciality] = NewRankByStartTime()
		}

		return nil
	}

	return errors.New("doctor service unavailable")
}

func (doctorService *DoctorService) AddSlot(name string, slots [][]int) error {
	if doctorService != nil {
		if doctor, exists := doctorService.DoctorRepository[name]; exists {

			for _, slot := range slots {
				if newSlot, err := doctorService.SlotManager.AddSlot(slot[0], slot[1], doctor); err != nil {
					fmt.Println(err)
				} else {
					if err := doctorService.SpecialityToDoctorMap[doctor.Speciality].updateRanking(newSlot); err != nil {
						fmt.Println(err)
					}
				}
			}

			return nil
		}

		return errors.New("doctor doesn't exists")
	}

	return errors.New("doctor service unavailable")
}

func (doctorService *DoctorService) DisplaySlots(speciality string) ([]*Slot, error) {
	if doctorService != nil {
		if rankingService, exists := doctorService.SpecialityToDoctorMap[speciality]; exists {
			return rankingService.displayRanking()
		}

		return nil, errors.New("speciality doesn't exists")
	}

	return nil, errors.New("doctor service unavailable")
}
