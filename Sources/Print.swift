let Red = "\u{001B}[0;31m"
let Green = "\u{001B}[0;32m"
let Yellow = "\u{001B}[0;33m"
let Blue = "\u{001B}[0;34m"
let Gray = "\u{001B}[0;90m"
let White = "\u{001B}[0;97m"
let Default = "\u{001B}[0;0m"

let LollipopEmoji = "🍭 "
let SmugEmoji = "😎 "
let ShitEmoji = "💩 "
let GhostEmoji = "👻 "
let NamasteEmoji = "🙏 "
let MiddleFingerEmoji = "🖕 "
let VictoryEmoji = "✌️ "
let HeartBreakEmoji = "💔 "
let BombEmoji = "💣 "
let BananaEmoji = "🍌 "
let BeersEmoji = "🍻 "
let BeerEmoji = "🍺 "
let RainEmoji = "⛈ "
let CorrectEmoji = "✅ "
let WrongEmoji = "❌ "
let CryEmoji = "😭 "

struct Log {
	func Error(_ message: String?) {
		if let data = message {
			print("\(Red)\(BombEmoji) Error\(Default)")
			print(data)
		} else {
			print("\(Red)◀◀ ERROR\(Default)")
		}
	}

	func Highlight(_ message: String?) {
		if let data = message {
			print("\n\(Blue)\(LollipopEmoji) \(data)\(Default)")
		}
	}
}

var log = Log()