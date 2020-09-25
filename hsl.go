/**
支持 4(255) bit 和 8(rgb) bit颜色显示
 */
package color

import (
	"errors"
	"fmt"
	"math"
)

type HSL struct {
	h float64
	s float64
	l float64
}

func NewHSL(h float64,s float64,l float64) *HSL{

	return &HSL{h,s,l}
}

// HSL2RGB 转换 hsl 到 rgb
func (hsl *HSL)HSL2RGB() (*RGB,error) {
	h := hsl.h
	s := hsl.s
	l := hsl.l
	C := (1 - math.Abs(2*l-1)) * s
	hPrime := h / 60
	X := C * (1 - math.Abs(math.Mod(hPrime,2)-1))
	m := l - C/2
	withLight := func (r, g, b, m float64) *RGB {
		return &RGB{(r + m) * 255, (g + m)* 255, (b + m)*255}
	}
	if hPrime <= 1 {
		return withLight(C, X, 0, m),nil
	} else if hPrime <= 2 {
		return withLight(X, C, 0, m),nil
	} else if hPrime <= 3 {
		return withLight(0, C, X, m),nil
	} else if hPrime <= 4 {
		return withLight(0, X, C, m),nil
	} else if hPrime <= 5 {
		return withLight(X, 0, C, m),nil
	} else if hPrime <= 6 {
		return withLight(C, 0, X, m),nil
	}else{
		return &RGB{0,0,0},errors.New("hsl 解析出错")
	}

}


// BgPrint print background color for text
func (hsl *HSL)BgPrint(text interface{}){
	rgb,_ := hsl.HSL2RGB()
	fmt.Printf(bgColorTemplate24Bit,int(rgb.r),int(rgb.g),int(rgb.b),text)
}
// Print print font color for text
func (hsl *HSL)Print(text interface{}){
	rgb,_ := hsl.HSL2RGB()
	fmt.Printf(colorTemplate24Bit,int(rgb.r),int(rgb.g),int(rgb.b),text)
}


