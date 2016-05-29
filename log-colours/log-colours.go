package logColours

type Color struct {
	name   string
	Format string
}

var Red, Green, Yellow, Blue, Gray, White, Default Color

func init() {
	Red = Color{name: "red", Format: "\033[0;31m"}
	Green = Color{name: "green", Format: "\033[0;32m"}
	Yellow = Color{name: "yellow", Format: "\033[0;33m"}
	Blue = Color{name: "blue", Format: "\033[0;34m"}
	Gray = Color{name: "gray", Format: "\033[0;90m"}
	White = Color{name: "white", Format: "\033[0;97m"}
	Default = Color{name: "default", Format: "\033[0m"}
}
