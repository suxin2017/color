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

const ColorIcon = "   ____   U  ___ u   _       U  ___ u   ____     \nU /\"___|   \\/\"_ \\/  |\"|       \\/\"_ \\" +
	"/U |  _\"\\ u  \n\\| | u     | | | |U | | u     | | | | \\| |_) |/  \n | |/__.-,_| |_| | \\| |/__.-,_| |_| |  |" +
	"  _ <    \n  \\____|\\_)-\\___/   |_____|\\_)-\\___/   |_| \\_\\   \n _// \\\\      \\\\     //  \\\\      \\\\" +
	"     //   \\\\_  \n(__)(__)    (__)   (_\")(\"_)    (__)   (__)  (__) "

var RainbowColor = map[string]*RGB{
	"Red":    {255, 0, 0},
	"Origin": {255, 152, 0},
	"Yellow": {255, 255, 0},
	"Green":  {0, 255, 0},
	"Cyan":   {0, 127, 255},
	"Blue":   {0, 0, 255},
	"Purple": {139, 0, 255},
}

var RedRGB = NewRGB(255, 0, 0)
var OriginRGB = NewRGB(255, 152, 0)
