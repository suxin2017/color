package color

import (
	"fmt"
	"github.com/suxin2017/color"
	"testing"
)

func Example() {
	fmt.Println("hello")
	//output: hello
}
func TestRainbow(t *testing.T) {
	color.Rainbow(1)
}

func TestRainbowText(t *testing.T) {
	color.RainbowText("abcdefghijklmnopqrstuvwxyz")
}

func TestFormat(t *testing.T) {
	fmt.Print(color.NewRGB(255, 0, 0).FormatColor("test"))
}

func Test4BitColor(t *testing.T) {
	fmt.Println(color.Black("text"))
	fmt.Println(color.Red("text"))
	fmt.Println(color.Green("text"))
	fmt.Println(color.Yellow("text"))
	fmt.Println(color.Blue("text"))
	fmt.Println(color.Magenta("text"))
	fmt.Println(color.Cyan("text"))
	fmt.Println(color.White("text"))
	fmt.Println(color.BrightBlack("text"))
	fmt.Println(color.BrightRed("text"))
	fmt.Println(color.BrightGreen("text"))
	fmt.Println(color.BrightYellow("text"))
	fmt.Println(color.BrightBlue("text"))
	fmt.Println(color.BrightMagenta("text"))
	fmt.Println(color.BrightCyan("text"))
	fmt.Println(color.BrightWhite("text"))
	fmt.Println(color.BgBlack("text"))
	fmt.Println(color.BgRed("text"))
	fmt.Println(color.BgGreen("text"))
	fmt.Println(color.BgYellow("text"))
	fmt.Println(color.BgBlue("text"))
	fmt.Println(color.BgMagenta("text"))
	fmt.Println(color.BgCyan("text"))
	fmt.Println(color.BgWhite("text"))
	fmt.Println(color.BgBrightBlack("text"))
	fmt.Println(color.BgBrightRed("text"))
	fmt.Println(color.BgBrightGreen("text"))
	fmt.Println(color.BgBrightYellow("text"))
	fmt.Println(color.BgBrightBlue("text"))
	fmt.Println(color.BgBrightMagenta("text"))
	fmt.Println(color.BgBrightCyan("text"))
	fmt.Println(color.BgBrightWhite("text"))
}

func TestCompose(t *testing.T) {
	fmt.Println(color.BgBrightCyan(color.BrightRed("blue bg bright green text")))
}

func Test8BitColor(t *testing.T) {
	fmt.Println(color.Format8BitColor(196, "text"))
	fmt.Println(color.Format8BitBgColor(196, "text"))
}

func TestNeonLight(t *testing.T) {
	color.NeonLight(color.ColorIcon)
}
