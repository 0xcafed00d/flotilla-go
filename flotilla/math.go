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

func MaxInt(vals ...int) (max int) {
	if len(vals) > 0 {
		max = vals[0]
		for _, val := range vals {
			if val > max {
				max = val
			}
		}
	}
	return
}

func MinInt(vals ...int) (min int) {
	if len(vals) > 0 {
		min = vals[0]
		for _, val := range vals {
			if val < min {
				min = val
			}
		}
	}
	return
}
