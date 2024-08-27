package cliformat

var (
	Reset = "\033[0m"

	Bold      = "\033[1m"
	Underline = "\033[4m"

	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"

	BlackBackground  = "\033[40m"
	RedBackground    = "\033[41m"
	GreenBackground  = "\033[42m"
	YellowBackground = "\033[43m"
	BlueBackground   = "\033[44m"
	PurpleBackground = "\033[45m"
	CyanBackground   = "\033[46m"
	GrayBackground   = "\033[47m"
	WhiteBackground  = "\033[107m"
)

var (
	enabled = true
)

func Colorize(color string, text string) string {
	if !enabled {
		return text
	}
	return color + text + Reset
}

func With(color string, text string) string {
	return Colorize(color, text)
}

func InBold(text string) string {
	return Colorize(Bold, text)
}

func InUnderline(text string) string {
	return Colorize(Underline, text)
}

func InBlue(text string) string {
	return Colorize(Blue, text)
}

func InPurple(text string) string {
	return Colorize(Purple, text)
}

func InWhite(text string) string {
	return Colorize(White, text)
}

func BackgroundGreen(text string) string {
	return Colorize(GreenBackground, text)
}

func BackgroundRed(text string) string {
	return Colorize(RedBackground, text)
}

func BackgroundYellow(text string) string {
	return Colorize(YellowBackground, text)
}

func BackgroundWhite(text string) string {
	return Colorize(WhiteBackground, text)
}

func Error(text string) string {
	return BackgroundRed(InWhite(InBold(text)))
}

func Success(text string) string {
	return BackgroundGreen(InWhite(InBold(text)))
}

func Info(text string) string {
	return BackgroundYellow(InWhite(InBold(text)))
}
