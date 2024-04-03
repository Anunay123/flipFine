package src

import (
	"errors"
	"fmt"
	"strconv"
)

type SlotManager struct {
	Slots map[string]*Slot
}

func NewSlotManager() *SlotManager {
	return &SlotManager{
		Slots: make(map[string]*Slot, 0),
	}
}

func (slotManager *SlotManager) AddSlot(startTime int, endTime int, doctor *Doctor) (*Slot, error) {
	if slotManager != nil {

		slotString := generateSlotString(startTime, endTime)
		if startTime < 0 || startTime >= 2330 || endTime < 0 || endTime > 2400 || endTime-startTime > 30 {
			return nil, errors.New(fmt.Sprintf("%s slot can't be added", slotString))
		}

		newSlot := NewSlot(startTime, endTime, doctor, slotString)

		key := fmt.Sprintf("%s_%s_%s", doctor.Name, strconv.Itoa(startTime), strconv.Itoa(endTime))

		slotManager.Slots[key] = newSlot

		return newSlot, nil
	}

	return nil, errors.New("slot manager unavailable")
}
