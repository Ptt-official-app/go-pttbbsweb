package types

type Color struct {
	Foreground ColorMap `json:"foreground" bson:"f"`
	Background ColorMap `json:"background" bson:"b"`
	Blink      bool     `json:"blink" bson:"k"`
	Highlight  bool     `json:"highlight" bson:"h"`
	IsReset    bool     `json:"-" bson:"-"`
}

type ColorMap int8

var (
	DefaultColor = Color{Foreground: COLOR_FOREGROUND_WHITE, Background: COLOR_BACKGROUND_BLACK}
	InvalidColor = Color{Foreground: COLOR_INVALID, Background: COLOR_INVALID}
	ResetColor   = Color{Foreground: COLOR_FOREGROUND_WHITE, Background: COLOR_BACKGROUND_BLACK, IsReset: true}
)

const (
	COLOR_FOREGROUND_BLACK   ColorMap = 30
	COLOR_FOREGROUND_RED     ColorMap = 31
	COLOR_FOREGROUND_GREEN   ColorMap = 32
	COLOR_FOREGROUND_YELLOW  ColorMap = 33
	COLOR_FOREGROUND_BLUE    ColorMap = 34
	COLOR_FOREGROUND_MAGENTA ColorMap = 35
	COLOR_FOREGROUND_CYAN    ColorMap = 36
	COLOR_FOREGROUND_WHITE   ColorMap = 37
	COLOR_BACKGROUND_BLACK   ColorMap = 40
	COLOR_BACKGROUND_RED     ColorMap = 41
	COLOR_BACKGROUND_GREEN   ColorMap = 42
	COLOR_BACKGROUND_YELLOW  ColorMap = 43
	COLOR_BACKGROUND_BLUE    ColorMap = 44
	COLOR_BACKGROUND_MAGENTA ColorMap = 45
	COLOR_BACKGROUND_CYAN    ColorMap = 46
	COLOR_BACKGROUND_WHITE   ColorMap = 47
	COLOR_INVALID            ColorMap = -1
)
