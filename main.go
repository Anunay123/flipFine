package main

import (
	"flipfine/src"
	"fmt"
)

func main() {
	ff := src.NewFlipFine()

	if err := ff.RegisterDoctor("Anunay", "Ortho"); err != nil {
		fmt.Println(err)
	}

	slots := make([][]int, 0)

	slots = append(slots, []int{900, 930})
	slots = append(slots, []int{1000, 1030})
	slots = append(slots, []int{1200, 1230})

	if err := ff.AddSlots("Anunay", slots); err != nil {
		fmt.Println(err)
	}

	if err := ff.RegisterPatient("Akhil"); err != nil {
		fmt.Println(err)
	}

	if err := ff.Display("Ortho"); err != nil {
		fmt.Println(err)
	}

	if err := ff.RegisterDoctor("Naresh", "Ortho"); err != nil {
		fmt.Println(err)
	}

	newSlots := make([][]int, 0)

	newSlots = append(newSlots, []int{900, 930})

	if err := ff.AddSlots("Naresh", newSlots); err != nil {
		fmt.Println(err)
	}

	if err := ff.Display("Ortho"); err != nil {
		fmt.Println(err)
	}

	if bookingId, err := ff.MakeBooking("Anunay", "Akhil", 900, 930); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bookingId)
	}

	if err := ff.RegisterPatient("Aman"); err != nil {
		fmt.Println(err)
	}

	if bookingId, err := ff.MakeBooking("Anunay", "Aman", 900, 930); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bookingId)
	}

	if err := ff.CancelBooking(1); err != nil {
		fmt.Println(err)
	}

	if err := ff.ViewBookingsPatient("Aman"); err != nil {
		fmt.Println(err)
	}

	if err := ff.ViewBookingsDoctor("Anunay"); err != nil {
		fmt.Println(err)
	}

	if err := ff.Display("Ortho"); err != nil {
		fmt.Println(err)
	}

}
