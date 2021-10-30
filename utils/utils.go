package disys_mini_project_2_utils

import "math"
import "fmt"

func CalcNextLTime(id int32,ownLTime *[]int32, receivedLTime *[]int32) {
	maxLen := int(math.Max(float64(len(*ownLTime)), float64(len(*receivedLTime))))
	newLTime := []int32{}
	for i := 0; i < int(maxLen); i++ {
		if i < len(*ownLTime)-1 && i < len(*receivedLTime) {
			newLTime = append(newLTime, int32(math.Max(float64((*ownLTime)[i]), float64((*receivedLTime)[i]))))
		} else if i < len(*ownLTime) {
			newLTime = append(newLTime, (*ownLTime)[i])
		} else if i < len(*receivedLTime) {
			newLTime = append(newLTime, (*receivedLTime)[i])
		}
	}
	newLTime[id] = newLTime[id] + 1
	*ownLTime = newLTime
}

func IncrementLTime(id int32, ownLTime *[]int32) {
	(*ownLTime)[id]++
}

func LTimeToString(lTime []int32) string {
	output := " [:"
	for _, i := range lTime {
		output += fmt.Sprintf("%d:", i)
	}
	output += "]"
	return output
}
