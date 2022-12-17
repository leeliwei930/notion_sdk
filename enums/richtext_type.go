package enums

type RichTextType int

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
