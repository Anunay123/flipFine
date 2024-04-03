package src

import (
	"errors"
	"sort"
)

type IRankingService interface {
	displayRanking() ([]*Slot, error)
	updateRanking(*Slot) error
}

type RankByStartTime struct {
	orderedList []*Slot
}

func NewRankByStartTime() *RankByStartTime {
	return &RankByStartTime{
		orderedList: make([]*Slot, 0),
	}
}

func (rbst *RankByStartTime) displayRanking() ([]*Slot, error) {
	if rbst != nil {
		return rbst.orderedList, nil
	}

	return nil, errors.New("ranking unavailable")
}

func (rbst *RankByStartTime) updateRanking(slot *Slot) error {
	if rbst != nil {
		newOrderedList := append(rbst.orderedList, slot)

		sort.Slice(newOrderedList, func(i, j int) bool {
			return newOrderedList[i].StartTime < newOrderedList[j].StartTime
		})

		rbst.orderedList = newOrderedList

		return nil
	}

	return errors.New("ranking unavailable")
}
