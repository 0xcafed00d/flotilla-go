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

func int2byte(a int) byte {
	if a > 255 {
		return 255
	}
	return byte(a)
}

func Add(rgb1, rgb2 RGB) (rgb RGB) {
	rgb.R = int2byte((int(rgb1.R) + int(rgb2.R)))
	rgb.G = int2byte((int(rgb1.G) + int(rgb2.G)))
	rgb.B = int2byte((int(rgb1.B) + int(rgb2.B)))
	return
}

func LerpRGB(rgb1, rgb2 RGB, t float64) (rgb RGB) {
	rgb.R = byte(LerpFloat(float64(rgb1.R), float64(rgb2.R), t))
	rgb.G = byte(LerpFloat(float64(rgb1.G), float64(rgb2.G), t))
	rgb.B = byte(LerpFloat(float64(rgb1.B), float64(rgb2.B), t))
	return
}
