package src

import (
	"fmt"
	"strconv"
)

func generateSlotString(startTime int, endTime int) string {

	return fmt.Sprintf("%s-%s", strconv.Itoa(startTime), strconv.Itoa(endTime))
}
