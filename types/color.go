package types

type Color struct {
	Foreground ColorMap `json:"foreground" bson:"f"`
	Background ColorMap `json:"background" bson:"b"`
	Blink      bool     `json:"blink" bson:"k"`
	Highlight  bool     `json:"highlight" bson:"h"`
	IsReset    bool     `json:"reset" bson:"r"`
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

func (c ColorMap) String() string {
	switch c {
	case COLOR_FOREGROUND_BLACK:
		return "f-black"
	case COLOR_FOREGROUND_RED:
		return "f-red"
	case COLOR_FOREGROUND_GREEN:
		return "f-green"
	case COLOR_FOREGROUND_YELLOW:
		return "f-yellow"
	case COLOR_FOREGROUND_BLUE:
		return "f-blue"
	case COLOR_FOREGROUND_MAGENTA:
		return "f-magenta"
	case COLOR_FOREGROUND_CYAN:
		return "f-cyan"
	case COLOR_FOREGROUND_WHITE:
		return "f-white"
	case COLOR_BACKGROUND_BLACK:
		return "b-black"
	case COLOR_BACKGROUND_RED:
		return "b-red"
	case COLOR_BACKGROUND_GREEN:
		return "b-green"
	case COLOR_BACKGROUND_YELLOW:
		return "b-yellow"
	case COLOR_BACKGROUND_BLUE:
		return "b-blue"
	case COLOR_BACKGROUND_MAGENTA:
		return "b-magenta"
	case COLOR_BACKGROUND_CYAN:
		return "b-cyan"
	case COLOR_BACKGROUND_WHITE:
		return "b-white"
	case COLOR_INVALID:
		return "invalid"
	default:
		return "unknown"
	}
}

var colorBytesMap = map[ColorMap][]byte{
	COLOR_FOREGROUND_BLACK:   {'3', '0'},
	COLOR_FOREGROUND_RED:     {'3', '1'},
	COLOR_FOREGROUND_GREEN:   {'3', '2'},
	COLOR_FOREGROUND_YELLOW:  {'3', '3'},
	COLOR_FOREGROUND_BLUE:    {'3', '4'},
	COLOR_FOREGROUND_MAGENTA: {'3', '5'},
	COLOR_FOREGROUND_CYAN:    {'3', '6'},
	COLOR_FOREGROUND_WHITE:   {'3', '7'},
	COLOR_BACKGROUND_BLACK:   {'4', '0'},
	COLOR_BACKGROUND_RED:     {'4', '1'},
	COLOR_BACKGROUND_GREEN:   {'4', '2'},
	COLOR_BACKGROUND_YELLOW:  {'4', '3'},
	COLOR_BACKGROUND_BLUE:    {'4', '4'},
	COLOR_BACKGROUND_MAGENTA: {'4', '5'},
	COLOR_BACKGROUND_CYAN:    {'4', '6'},
	COLOR_BACKGROUND_WHITE:   {'4', '7'},
}

const (
	DEFAULT_LEN_COLOR_BYTES = 20 //\x1b[0;1;5;37;40m
)

func (c *Color) BytesWithPreColor(color *Color, isForceReset bool) (theBytes []byte) {
	if c.IsReset {
		if color.IsReset && !isForceReset {
			return nil
		}
		return []byte("\x1b[m")
	}

	if color.IsReset {
		color = &DefaultColor
	}

	if *c == *color {
		return nil
	}

	theBytes = make([]byte, 0, DEFAULT_LEN_COLOR_BYTES)
	theBytes = append(theBytes, []byte{'\x1b', '['}...)

	isFirst := true
	isReset := false
	if c.Highlight != color.Highlight && !c.Highlight ||
		c.Blink != color.Blink && !c.Blink {
		theBytes = append(theBytes, '0')
		isReset = true
		isFirst = false
	}

	if c.Highlight != color.Highlight && c.Highlight {
		if !isFirst {
			theBytes = append(theBytes, ';')
		} else {
			isFirst = false
		}
		theBytes = append(theBytes, '1')
	}

	if c.Blink != color.Blink && c.Blink {
		if !isFirst {
			theBytes = append(theBytes, ';')
		} else {
			isFirst = false
		}
		theBytes = append(theBytes, '5')
	}

	if c.Foreground != color.Foreground || isReset && c.Foreground != DefaultColor.Foreground {
		if !isFirst {
			theBytes = append(theBytes, ';')
		} else {
			isFirst = false
		}
		theBytes = append(theBytes, colorBytesMap[c.Foreground]...)
	}

	if c.Background != color.Background || isReset && c.Background != DefaultColor.Background {
		if !isFirst {
			theBytes = append(theBytes, ';')
		}
		theBytes = append(theBytes, colorBytesMap[c.Background]...)
	}

	theBytes = append(theBytes, 'm')

	return theBytes
}
