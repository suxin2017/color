package color

import "fmt"

type RGB struct {
	r float64
	g float64
	b float64
}

// BgPrint print background color for text
func (rgb *RGB)BgPrint(text interface{}){

	fmt.Printf(rgb.Format8BitBgColor(text))
}
// Print print font color for text
func (rgb *RGB)Print(text interface{}){
	fmt.Printf(rgb.Format8BitColor(text))
}
// BgPrint print background color for text
func (rgb *RGB)BgPrint24Bit(text interface{}){

	fmt.Printf(rgb.FormatBgColor(text))
}
// Print print font color for text
func (rgb *RGB)Print124Bit(text interface{}){
	fmt.Printf(rgb.FormatColor(text))
}

func NewRGB(r,g,b float64) *RGB{
	return &RGB{r,g,b}
}

func (rgb *RGB) FormatColor(text interface{}) string{
	return Format(colorTemplate24Bit,int(rgb.r),int(rgb.g),int(rgb.b),text)
}
func (rgb *RGB) FormatBgColor(text interface{}) string{
	return Format(bgColorTemplate24Bit,int(rgb.r),int(rgb.g),int(rgb.b),text)
}

// ColorTo8Bit convert 24 bit color to 8 bit color support all terminal
func (rgb *RGB) ColorTo8Bit() int {
	r,g,b:=rgb.r,rgb.g,rgb.b;
	return 36 * (int(r)/51) + 6*(int(g)/51) + (int(b)/51) + 16
}
func (rgb *RGB) Format8BitColor(text interface{}) string{
	return Format(colorTemplate8Bit,rgb.ColorTo8Bit(),text)
}
func (rgb *RGB) Format8BitBgColor(text interface{}) string{
	return Format(bgColorTemplate8Bit,rgb.ColorTo8Bit(),text)
}

