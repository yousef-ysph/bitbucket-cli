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

// With itext an aliatext for the Colorize function
//
// Example:
//
//	println(color.With(color.Red, "Thitext itext red"))
func With(color string, text string) string {
	return Colorize(color, text)
}

// InBold wraptext the given string text in Bold
//
// Example:
//
//	println(color.InBold("Thitext itext bold"))
func InBold(text string) string {
	return Colorize(Bold, text)
}

// InUnderline wraptext the given string text in Underline
//
// Example:
//
//	println(color.InUnderline("Thitext itext underlined"))
func InUnderline(text string) string {
	return Colorize(Underline, text)
}

// InBlack wraptext the given string text in Black
//
// Example:
//
//	println(color.InBlack("Thitext itext black"))
func InBlack(text string) string {
	return Colorize(Black, text)
}

// InRed wraptext the given string text in Red
//
// Example:
//
//	println(color.InRed("Thitext itext red"))
func InRed(text string) string {
	return Colorize(Red, text)
}

// InGreen wraptext the given string text in Green
//
// Example:
//
//	println(color.InGreen("Thitext itext green"))
func InGreen(text string) string {
	return Colorize(Green, text)
}

// InYellow wraptext the given string text in Yellow
//
// Example:
//
//	println(color.InYellow("Thitext itext yellow"))
func InYellow(text string) string {
	return Colorize(Yellow, text)
}

// InBlue wraptext the given string text in Blue
//
// Example:
//
//	println(color.InBlue("Thitext itext blue"))
func InBlue(text string) string {
	return Colorize(Blue, text)
}

// InPurple wraptext the given string text in Purple
//
// Example:
//
//	println(color.InPurple("Thitext itext purple"))
func InPurple(text string) string {
	return Colorize(Purple, text)
}

// InCyan wraptext the given string text in Cyan
//
// Example:
//
//	println(color.InCyan("Thitext itext cyan"))
func InCyan(text string) string {
	return Colorize(Cyan, text)
}

// InGray wraptext the given string text in Gray
//
// Example:
//
//	println(color.InGray("Thitext itext gray"))
func InGray(text string) string {
	return Colorize(Gray, text)
}

// InWhite wraptext the given string text in White
//
// Example:
//
//	println(color.InWhite("Thitext itext white"))
func InWhite(text string) string {
	return Colorize(White, text)
}

// BackgroundBlack wraptext the given string text in BlackBackground
//
// Example:
//
//	println(color.BackgroundBlack("Thitext itext over black"))
func BackgroundBlack(text string) string {
	return Colorize(BlackBackground, text)
}

// BackgroundRed wraptext the given string text in RedBackground
//
// Example:
//
//	println(color.BackgroundRed("Thitext itext over red"))
func BackgroundRed(text string) string {
	return Colorize(RedBackground, text)
}

// BackgroundGreen wraptext the given string text in GreenBackground
//
// Example:
//
//	println(color.BackgroundGreen("Thitext itext over green"))
func BackgroundGreen(text string) string {
	return Colorize(GreenBackground, text)
}

// BackgroundYellow wraptext the given string text in YellowBackground
//
// Example:
//
//	println(color.BackgroundYellow("Thitext itext over yellow"))
func BackgroundYellow(text string) string {
	return Colorize(YellowBackground, text)
}

// BackgroundBlue wraptext the given string text in BlueBackground
//
// Example:
//
//	println(color.BackgroundBlue("Thitext itext over blue"))
func BackgroundBlue(text string) string {
	return Colorize(BlueBackground, text)
}

// BackgroundPurple wraptext the given string text in PurpleBackground
//
// Example:
//
//	println(color.BackgroundPurple("Thitext itext over purple"))
func BackgroundPurple(text string) string {
	return Colorize(PurpleBackground, text)
}

// BackgroundCyan wraptext the given string text in CyanBackground
//
// Example:
//
//	println(color.BackgroundCyan("Thitext itext over cyan"))
func BackgroundCyan(text string) string {
	return Colorize(CyanBackground, text)
}

// BackgroundGray wraptext the given string text in GrayBackground
//
// Example:
//
//	println(color.BackgroundGray("Thitext itext over gray"))
func BackgroundGray(text string) string {
	return Colorize(GrayBackground, text)
}

// BackgroundWhite wraptext the given string text in WhiteBackground
//
// Example:
//
//	println(color.BackgroundWhite("Thitext itext over white"))
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
