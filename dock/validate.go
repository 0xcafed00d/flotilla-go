package dock

import "fmt"

type param struct {
	minval int
	maxval int
}

type paramsInfo struct {
	paramCounts []int
	paramLimits []param
}

var validationData = map[ModuleType]paramsInfo{
	Motor: {
		[]int{1},
		[]param{{-63, 63}},
	},
	Matrix: paramsInfo{
		[]int{9},
		[]param{
			{0, 255},
			{0, 255},
			{0, 255},
			{0, 255},
			{0, 255},
			{0, 255},
			{0, 255},
			{0, 255},
			{0, 255},
		},
	},
	Rainbow: paramsInfo{
		[]int{3, 15},
		[]param{
			{0, 255}, {0, 255}, {0, 255},
			{0, 255}, {0, 255}, {0, 255},
			{0, 255}, {0, 255}, {0, 255},
			{0, 255}, {0, 255}, {0, 255},
			{0, 255}, {0, 255}, {0, 255},
		},
	},
	Number: paramsInfo{
		[]int{4, 5, 6},
		[]param{
			{0, 255},
			{0, 255},
			{0, 255},
			{0, 255},
			{0, 1},
			{0, 1},
		},
	},
}

func contains(s []int, a int) bool {
	for _, r := range s {
		if a == r {
			return true
		}
	}
	return false
}

func validateParams(mtype ModuleType, params []int) error {
	info, ok := validationData[mtype]
	if !ok {
		return fmt.Errorf("Module: %v not writeable", mtype)
	}

	if !contains(info.paramCounts, len(params)) {
		return fmt.Errorf("Module: %v invalid param count (%v) expecting %v", mtype, len(params), info.paramCounts)
	}

	for i, param := range params {
		lim := info.paramLimits[i]
		if param < lim.minval || param > lim.maxval {
			return fmt.Errorf("Module: %v param %v (%v) out of range of: (%v)",
				mtype, i, params, lim)
		}
	}

	return nil
}
