package colours

var Red, Green, Yellow, Blue, Gray, White, Default, Bold, Light, Underline string

func init() {
	Red = "\033[0;31m"
	Green = "\033[0;32m"
	Yellow = "\033[0;33m"
	Blue = "\033[0;34m"
	Gray = "\033[0;90m"
	White = "\033[0;97m"
	Default = "\033[0m"

	Bold = "\033[1m"
	Light = "\033[2m"
	Underline = "\033[4m"
}
