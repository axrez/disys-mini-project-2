package disys_mini_project_2_utils

import "math"

func CalcNextLTime(ownLTime *int32, recivedLTime *int32) {
	max := math.Max(float64(*ownLTime), float64(*recivedLTime))
	*ownLTime = int32(max) + 1
}
