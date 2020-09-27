package color

import (
	"fmt"
	"math"
)

type RGB struct {
	r float64
	g float64
	b float64
}

//
func NewRGB(r, g, b float64) *RGB {
	return &RGB{r, g, b}
}

// BgPrint print background color for text
func (rgb *RGB) BgPrint(text interface{}) {
	fmt.Printf(rgb.Format8BitBgColor(text))
}

// Print print font color for text
func (rgb *RGB) Print(text interface{}) {
	fmt.Printf(rgb.Format8BitColor(text))
}

// BgPrint24Bit print background color for text
func (rgb *RGB) BgPrint24Bit(text interface{}) {

	fmt.Printf(rgb.FormatBgColor(text))
}

// Print24Bit print font color for text
func (rgb *RGB) Print24Bit(text interface{}) {
	fmt.Printf(rgb.FormatColor(text))
}

func (rgb *RGB) FormatColor(text interface{}) string {
	return Format(colorTemplate24Bit, int(rgb.r), int(rgb.g), int(rgb.b), text)
}
func (rgb *RGB) FormatBgColor(text interface{}) string {
	return Format(bgColorTemplate24Bit, int(rgb.r), int(rgb.g), int(rgb.b), text)
}

// ColorTo8Bit convert 24 bit color to 8 bit color support all terminal
func (rgb *RGB) ColorTo8Bit() int {
	r, g, b := rgb.r, rgb.g, rgb.b
	return 36*(int(r)/51) + 6*(int(g)/51) + (int(b) / 51) + 16
}
func (rgb *RGB) Format8BitColor(text interface{}) string {
	return Format(colorTemplate8Bit, rgb.ColorTo8Bit(), text)
}
func (rgb *RGB) Format8BitBgColor(text interface{}) string {
	//println(rgb.ColorTo8Bit())
	return Format(bgColorTemplate8Bit, rgb.ColorTo8Bit(), text)

}
func (rgb *RGB) getRgb() (float64, float64, float64) {
	return rgb.r, rgb.g, rgb.b
}
func (rgb *RGB) rgb2Hue() float64 {
	r, g, b := rgb.getRgb()
	hue := math.Round(math.Atan2(math.Sqrt(3)*(g-b), 2*r-g-b)) * 100 / math.Pi
	for hue < 0 {
		hue = hue + 360
	}
	return hue
}
func (rgb *RGB) rgb2Saturation() float64 {
	r, g, b := rgb.getRgb()
	 max := math.Max(r,math.Max(g,b))
	 min := math.Min(r,math.Min(g,b))
	 l := rgb.rgb2Lightness()

	 if l == 0 || l == 1{
	 	return 0
	 }else {
	 	return (max - min)/(1 - math.Abs(2 * l - 1))
	 }
}
func (rgb *RGB) rgb2Lightness() float64 {
	r, g, b := rgb.getRgb()
	max := math.Max(r,math.Max(g,b))
	min := math.Min(r,math.Min(g,b))
	return 1/2 * (max + min)
}
func (rgb *RGB) RGB2HSL() *HSL {
	h := rgb.rgb2Hue()
	s := rgb.rgb2Saturation()
	l := rgb.rgb2Lightness()
	return NewHSL(h, s, l)

}
