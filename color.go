package color

import (
	"fmt"
)

// Rainbow print rainbow color to terminal
func Rainbow(colorOffsetStep int) {
	for h := 0; h < 360; h += colorOffsetStep {
		rgb, err := NewHSL(float64(h), 100/100.0, 50/100.0).HSL2RGB()
		if err == nil {
			rgb.BgPrint(" ")
		}
	}
}

// RainbowText print rainbow text
func RainbowText(text string) {
	for i, character := range text {
		currentColor := i * 360 / len(text)
		hsl := NewHSL(float64(currentColor), 100/100.0, 60/100.0)
		rgb, err := hsl.HSL2RGB()
		if err == nil {
			rgb.Print(string(character))
		}
	}
}

func Format(template string, a ...interface{}) string {
	return fmt.Sprintf(template, a...)
}

func Format8BitColor(color int,text string)string{
	return Format(colorTemplate8Bit,color,text)
}
func Format8BitBgColor(color int,text string)string{
	return Format(bgColorTemplate8Bit,color,text)
}


func Black(text string) string {
	return Format(colorTemplate4Bit, black, text)
}
func Red(text string) string {
	return Format(colorTemplate4Bit, red, text)
}
func Green(text string) string {
	return Format(colorTemplate4Bit, green, text)
}
func Yellow(text string) string {
	return Format(colorTemplate4Bit, yellow, text)
}
func Blue(text string) string {
	return Format(colorTemplate4Bit, blue, text)
}
func Magenta(text string) string {
	return Format(colorTemplate4Bit, magenta, text)
}
func Cyan(text string) string {
	return Format(colorTemplate4Bit, cyan, text)
}
func White(text string) string {
	return Format(colorTemplate4Bit, white, text)
}
func BrightBlack(text string) string {
	return Format(colorTemplate4Bit, brightBlack, text)
}
func BrightRed(text string) string {
	return Format(colorTemplate4Bit, brightRed, text)
}
func BrightGreen(text string) string {
	return Format(colorTemplate4Bit, brightGreen, text)
}
func BrightYellow(text string) string {
	return Format(colorTemplate4Bit, brightYellow, text)
}
func BrightBlue(text string) string {
	return Format(colorTemplate4Bit, brightBlue, text)
}
func BrightMagenta(text string) string {
	return Format(colorTemplate4Bit, brightMagenta, text)
}
func BrightCyan(text string) string {
	return Format(colorTemplate4Bit, brightCyan, text)
}
func BrightWhite(text string) string {
	return Format(colorTemplate4Bit, brightWhite, text)
}

func BgBlack(text string) string {
	return Format(colorTemplate4Bit, bgBlack, text)
}
func BgRed(text string) string {
	return Format(colorTemplate4Bit, bgRed, text)
}
func BgGreen(text string) string {
	return Format(colorTemplate4Bit, bgGreen, text)
}
func BgYellow(text string) string {
	return Format(colorTemplate4Bit, bgYellow, text)
}
func BgBlue(text string) string {
	return Format(colorTemplate4Bit, bgBlue, text)
}
func BgMagenta(text string) string {
	return Format(colorTemplate4Bit, bgMagenta, text)
}
func BgCyan(text string) string {
	return Format(colorTemplate4Bit, bgCyan, text)
}
func BgWhite(text string) string {
	return Format(colorTemplate4Bit, bgWhite, text)
}
func BgBrightBlack(text string) string {
	return Format(colorTemplate4Bit, bgBrightBlack, text)
}
func BgBrightRed(text string) string {
	return Format(colorTemplate4Bit, bgBrightRed, text)
}
func BgBrightGreen(text string) string {
	return Format(colorTemplate4Bit, bgBrightGreen, text)
}
func BgBrightYellow(text string) string {
	return Format(colorTemplate4Bit, bgBrightYellow, text)
}
func BgBrightBlue(text string) string {
	return Format(colorTemplate4Bit, bgBrightBlue, text)
}
func BgBrightMagenta(text string) string {
	return Format(colorTemplate4Bit, bgBrightMagenta, text)
}
func BgBrightCyan(text string) string {
	return Format(colorTemplate4Bit, bgBrightCyan, text)
}
func BgBrightWhite(text string) string {
	return Format(colorTemplate4Bit, bgBrightWhite, text)
}
