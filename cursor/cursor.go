package cursor

import "fmt"

func Up(n int) {
	fmt.Printf("\033[%d;A", n)
}
func Down(n int) {
	fmt.Printf("\033[%d;B", n)
}
func Forward(n int) {
	fmt.Printf("\033[%d;C", n)
}
func Back(n int) {
	fmt.Printf("\033[%d;D", n)
}

func Position(n int, m int) {
	fmt.Printf("\033[%d;%d;H", n, m)
}
func EraseDisplay(n int) {
	fmt.Printf("\033[%d;J", n)
}

// EraseLine
// 清除行内的部分区域。如果n是0（或缺失），清除从光标位置到该行末尾的部分。
// 如果n是1，清除从光标位置到该行开头的部分。
// 如果n是2，清除整行。光标位置不变。
func EraseLine(n int) {
	fmt.Printf("\033[%d;K", n)
}

func SavePosition() {
	fmt.Print("\0337")
}

func RestorePosition() {
	fmt.Print("\0338")
}
