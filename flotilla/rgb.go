package flotilla

type RGB struct {
	R, G, B byte
}

func Blend(rgb1, rgb2 RGB) (rgb RGB) {
	rgb.R = byte((uint(rgb1.R) + uint(rgb2.R)) / 2)
	rgb.G = byte((uint(rgb1.G) + uint(rgb2.G)) / 2)
	rgb.B = byte((uint(rgb1.B) + uint(rgb2.B)) / 2)
	return
}
