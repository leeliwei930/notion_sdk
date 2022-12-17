package enums

import "encoding/json"

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

var colorEnumsMap = map[Color]string{
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

var colorEnumsIndexes = map[string]Color{
	"default":           Default,
	"gray":              Gray,
	"brown":             Brown,
	"orange":            Orange,
	"yellow":            Yellow,
	"green":             Green,
	"blue":              Blue,
	"purple":            Purple,
	"pink":              Pink,
	"red":               Red,
	"gray_background":   GrayBackground,
	"brown_background":  BrownBackground,
	"orange_background": OrangeBackground,
	"yellow_background": YellowBackground,
	"green_background":  GreenBackground,
	"blue_background":   BlueBackground,
	"purple_background": PurpleBackground,
	"pink_background":   PinkBackground,
	"red_background":    RedBackground,
}

func ParseColor(s string) Color {
	return colorEnumsIndexes[s]
}

func (c Color) String() string {
	return colorEnumsMap[c]
}

func (c Color) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Color) UnmarshalJSON(d []byte) error {
	var j string
	err := json.Unmarshal(d, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*c = ParseColor(j)
	return nil
}
