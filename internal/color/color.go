// Package color provides ANSI color and style codes for terminal output.
// Terminal çıktısı için ANSI renk ve stil kodlarını içerir.
package color

// Reset clears all styles.
// Tüm stilleri temizler.
const Reset = "\033[0m"

// Style codes for text formatting.
// Yazı biçimlendirme için stil kodları.
const (
	Bold      = "\033[1m"
	Faint     = "\033[2m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
	Inverse   = "\033[7m"
	Hidden    = "\033[8m"
	Strike    = "\033[9m"
)

// Standard 8 colors (foreground).
// Standart 8 renk (ön plan).
const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// Bright versions (foreground).
// Parlak renkler (ön plan).
const (
	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
)

// Standard 8 colors (background).
// Standart 8 renk (arka plan).
const (
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

// Bright versions (background).
// Parlak renkler (arka plan).
const (
	BgBrightBlack   = "\033[100m"
	BgBrightRed     = "\033[101m"
	BgBrightGreen   = "\033[102m"
	BgBrightYellow  = "\033[103m"
	BgBrightBlue    = "\033[104m"
	BgBrightMagenta = "\033[105m"
	BgBrightCyan    = "\033[106m"
	BgBrightWhite   = "\033[107m"
)
