package enums

import "encoding/json"

type RichTextType uint8

const (
	Text RichTextType = iota
	Mention
	Equation
)

var richTextTypeMap = map[RichTextType]string{
	Text:     "text",
	Mention:  "mention",
	Equation: "equation",
}

var richTextTypeIndexes = map[string]RichTextType{
	"text":     Text,
	"mention":  Mention,
	"equation": Equation,
}

func ParseRichTextType(s string) RichTextType {
	return richTextTypeIndexes[s]
}

func (r RichTextType) String() string {
	return richTextTypeMap[r]
}

func (r RichTextType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *RichTextType) UnmarshalJSON(input []byte) error {
	var j string
	err := json.Unmarshal(input, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*r = ParseRichTextType(j)
	return nil
}
