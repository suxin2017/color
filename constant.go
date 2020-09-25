package color

const (
	black = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
	brightBlack = iota + 82
	brightRed
	brightGreen
	brightYellow
	brightBlue
	brightMagenta
	brightCyan
	brightWhite
)
const (
	bgBlack = iota + 40
	bgRed
	bgGreen
	bgYellow
	bgBlue
	bgMagenta
	bgCyan
	bgWhite
	bgBrightBlack = iota + 92
	bgBrightRed
	bgBrightGreen
	bgBrightYellow
	bgBrightBlue
	bgBrightMagenta
	bgBrightCyan
	bgBrightWhite
)

const colorTemplate4Bit = "\033[1;%dm%v\033[0m"
const colorTemplate8Bit = "\033[38;5;%dm%v\033[0m"
const bgColorTemplate8Bit = "\033[48;5;%dm%v\033[0m"
const colorTemplate24Bit = "\033[38;2;%d;%d;%dm%v\033[0m"
const bgColorTemplate24Bit = "\033[48;2;%d;%d;%dm%v\033[0m"

