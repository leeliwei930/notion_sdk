package enums

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

var BlockTypeEnumMap = map[BlockType]string{
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

func (b BlockType) String() string {
	return BlockTypeEnumMap[b]
}
