package enums

import "encoding/json"

type BlockType uint8

const (
	ParagraphBlockType = iota
	H1BlockType
	H2BlockType
	H3BlockType
	BulletListItemBlockType
	NumberedListItemBlockType
	ToDoBlockType
	ToggleBlockType
	ChildPageBlockType
	ChildDatabaseBlockType
	EmbedBlockType
	ImageBlockType
	VideoBlockType
	FileBlockType
	PdfBlockType
	BookmarkBlockType
	CalloutBlockType
	CodeBlockType
	QuoteBlockType
	EquationBlockType
	DividerBlockType
	TableOfContentsBlockType
	ColumnBlockType
	ColumnListBlockType
	LinkPreviewBlockType
	SyncedBlockBlockType
	TemplateBlockType
	LinkToPageBlockType
	UnsupportedBlockType
)

var blockTypeEnumMap = map[BlockType]string{
	ParagraphBlockType:        "paragraph",
	H1BlockType:               "heading_1",
	H2BlockType:               "heading_2",
	H3BlockType:               "heading_3",
	BulletListItemBlockType:   "bulleted_list_item",
	NumberedListItemBlockType: "numbered_list_item",
	ToDoBlockType:             "to_do",
	ToggleBlockType:           "toggle",
	ChildPageBlockType:        "child_page",
	ChildDatabaseBlockType:    "child_database",
	EmbedBlockType:            "embed",
	ImageBlockType:            "image",
	VideoBlockType:            "video",
	FileBlockType:             "file",
	PdfBlockType:              "pdf",
	BookmarkBlockType:         "bookmark",
	CalloutBlockType:          "callout",
	CodeBlockType:             "code",
	QuoteBlockType:            "quote",
	EquationBlockType:         "equation",
	DividerBlockType:          "divider",
	TableOfContentsBlockType:  "table_of_contents",
	ColumnBlockType:           "column",
	ColumnListBlockType:       "column_list",
	LinkPreviewBlockType:      "link_preview",
	SyncedBlockBlockType:      "synced_block",
	TemplateBlockType:         "template",
	LinkToPageBlockType:       "link_to_page",
	UnsupportedBlockType:      "unsupported",
}

var blockTypeEnumIndexes = map[string]BlockType{
	"paragraph":          ParagraphBlockType,
	"heading_1":          H1BlockType,
	"heading_2":          H2BlockType,
	"heading_3":          H3BlockType,
	"bulleted_list_item": BulletListItemBlockType,
	"numbered_list_item": NumberedListItemBlockType,
	"to_do":              ToDoBlockType,
	"toggle":             ToggleBlockType,
	"child_page":         ChildPageBlockType,
	"child_database":     ChildDatabaseBlockType,
	"embed":              EmbedBlockType,
	"image":              ImageBlockType,
	"video":              VideoBlockType,
	"file":               FileBlockType,
	"pdf":                PdfBlockType,
	"bookmark":           BookmarkBlockType,
	"callout":            CalloutBlockType,
	"code":               CodeBlockType,
	"quote":              QuoteBlockType,
	"equation":           EquationBlockType,
	"divider":            DividerBlockType,
	"table_of_contents":  TableOfContentsBlockType,
	"column":             ColumnBlockType,
	"column_list":        ColumnListBlockType,
	"link_preview":       LinkPreviewBlockType,
	"synced_block":       SyncedBlockBlockType,
	"template":           TemplateBlockType,
	"link_to_page":       LinkToPageBlockType,
	"unsupported":        UnsupportedBlockType,
}

func ParseBlockType(s string) BlockType {
	return blockTypeEnumIndexes[s]
}
func (b BlockType) String() string {
	return blockTypeEnumMap[b]
}

func (b BlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *BlockType) UnmarshalJSON(d []byte) error {
	var j string
	err := json.Unmarshal(d, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*b = ParseBlockType(j)
	return nil
}
