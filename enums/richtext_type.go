package enums

type RichTextType uint8

const (
	Text RichTextType = iota
	Mention
	Equation
)

var RichTextTypeMap = map[RichTextType]string{
	Text:     "text",
	Mention:  "mention",
	Equation: "equation",
}

func (r RichTextType) String() string {
	return RichTextTypeMap[r]
}
