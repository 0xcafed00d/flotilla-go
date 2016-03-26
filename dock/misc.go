package dock

const (
	decpnt = 1
	hmid   = 2
	vltop  = 4
	vlbot  = 8
	hbot   = 16
	vrbot  = 32
	vrtop  = 64
	htop   = 128
)

var digits = []int{
	htop | hbot | vlbot | vltop | vrbot | vrtop,
	vrbot | vrtop,
	htop | hbot | hmid | vlbot | vrtop,
	htop | hbot | hmid | vrbot | vrtop,
	hmid | vrtop | vrtop | vltop | vrbot,
	htop | hbot | hmid | vrbot | vltop,
	htop | hbot | hmid | vrbot | vltop | vlbot,
	htop | vrbot | vrtop | vrbot,
	htop | hmid | hbot | vlbot | vltop | vrbot | vrtop,
	htop | hmid | hbot | vltop | vrbot | vrtop,
}

func GetDigitPattern(n int, dpoint bool) int {
	if n < 0 || n > 9 {
		return 0
	}

	if dpoint {
		return digits[n] | decpnt
	}
	return digits[n]

}
