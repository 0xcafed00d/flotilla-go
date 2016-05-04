package flotilla

func Limit(val, min, max int) int {

	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

func Map(val, fromMin, fromMax, toMin, toMax int) int {
	fromMax++
	toMax++
	res := (val-fromMin)*(toMax-toMin)/(fromMax-fromMin) + toMin
	toMax--
	return Limit(res, toMin, toMax)
}
