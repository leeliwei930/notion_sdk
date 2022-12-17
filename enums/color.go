package enums

type Color uint8

const (
	Default Color = iota
	Gray
	Brown
	Orange
	Yellow
	Green
	Blue
	Purple
	Pink
	Red
	GrayBackground
	BrownBackground
	OrangeBackground
	YellowBackground
	GreenBackground
	BlueBackground
	PurpleBackground
	PinkBackground
	RedBackground
)

var ColorEnumsMap = map[Color]string{
	Default:          "default",
	Gray:             "gray",
	Brown:            "brown",
	Orange:           "orange",
	Yellow:           "yellow",
	Green:            "green",
	Blue:             "blue",
	Purple:           "purple",
	Pink:             "pink",
	Red:              "red",
	GrayBackground:   "gray_background",
	BrownBackground:  "brown_background",
	OrangeBackground: "orange_background",
	YellowBackground: "yellow_background",
	GreenBackground:  "green_background",
	BlueBackground:   "blue_background",
	PurpleBackground: "purple_background",
	PinkBackground:   "pink_background",
	RedBackground:    "red_background",
}

func (c Color) String() string {
	return ColorEnumsMap[c]
}
