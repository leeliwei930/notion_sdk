package models

type NotionRichTextType string

const (
	Text     NotionRichTextType = "text"
	Mention  NotionRichTextType = "mention"
	Equation NotionRichTextType = "equation"
)

type NotionColor string

const (
	Default          NotionColor = "default"
	Gray             NotionColor = "gray"
	Brown            NotionColor = "brown"
	Orange           NotionColor = "orange"
	Yellow           NotionColor = "yellow"
	Green            NotionColor = "green"
	Blue             NotionColor = "blue"
	Purple           NotionColor = "purple"
	Pink             NotionColor = "pink"
	Red              NotionColor = "red"
	GrayBackground   NotionColor = "gray_background"
	BrownBackground  NotionColor = "brown_background"
	OrangeBackground NotionColor = "orange_background"
	YellowBackground NotionColor = "yellow_background"
	GreenBackground  NotionColor = "green_background"
	BlueBackground   NotionColor = "blue_background"
	PurpleBackground NotionColor = "purple_background"
	PinkBackground   NotionColor = "pink_background"
	RedBackground    NotionColor = "red_background"
)

type NotionMentionType string

const (
	User        NotionMentionType = "user"
	Page        NotionMentionType = "page"
	Database    NotionMentionType = "database"
	Data        NotionMentionType = "data"
	LinkPreview NotionMentionType = "link_preview"
)

type NotionUserType string

const (
	Person NotionUserType = "person"
	Bot    NotionUserType = "bot"
)

type NotionBotOwnerType string

const (
	UserOwner      NotionBotOwnerType = "user"
	WorkspaceOwner NotionBotOwnerType = "workspace"
)

type NotionPropertyType string

const (
	NotionRichTextPropertyType       NotionPropertyType = "rich_text"
	NotionNumberPropertyType         NotionPropertyType = "number"
	NotionSelectPropertyType         NotionPropertyType = "select"
	NotionMultiSelectPropertyType    NotionPropertyType = "multi_select"
	NotionDatePropertyType           NotionPropertyType = "date"
	NotionFormulaPropertyType        NotionPropertyType = "formula"
	NotionRelationPropertyType       NotionPropertyType = "relation"
	NotionRollUpPropertyType         NotionPropertyType = "rollup"
	NotionTitlePropertyType          NotionPropertyType = "title"
	NotionPeoplePropertyType         NotionPropertyType = "people"
	NotionFilesPropertyType          NotionPropertyType = "files"
	NotionCheckboxPropertyType       NotionPropertyType = "checkbox"
	NotionUrlPropertyType            NotionPropertyType = "url"
	NotionEmailPropertyType          NotionPropertyType = "email"
	NotionPhoneNumberPropertyType    NotionPropertyType = "phone_number"
	NotionCreatedTimePropertyType    NotionPropertyType = "created_time"
	NotionCreatedByPropertyType      NotionPropertyType = "created_by"
	NotionLastEditedTimePropertyType NotionPropertyType = "last_edited_time"
	NotionLastEditedByPropertyType   NotionPropertyType = "last_edited_by"
)

type NotionFormulaPropertyValueType string

const (
	NotionFormulaPropertyStringValue  NotionFormulaPropertyValueType = "string"
	NotionFormulaPropertyNumberValue  NotionFormulaPropertyValueType = "number"
	NotionFormulaPropertyBooleanValue NotionFormulaPropertyValueType = "boolean"
	NotionFormulaPropertyDateValue    NotionFormulaPropertyValueType = "date"
)

type NotionRollUpPropertyValueType string

const (
	NotionRollUpNumberPropertyValue NotionRollUpPropertyValueType = "number"
	NotionRollUpDatePropertyValue   NotionRollUpPropertyValueType = "date"
	NotionRollUpArrayPropertyValue  NotionRollUpPropertyValueType = "array"
)

type NotionPageParentType string

const (
	NotionPageParentDatabase  NotionPageParentType = "database_id"
	NotionPageParentPage      NotionPageParentType = "page_id"
	NotionPageParentWorkspace NotionPageParentType = "workspace"
)

type NotionBlockType string

const (
	NotionParagraphBlockType        NotionBlockType = "paragraph"
	NotionH1BlockType               NotionBlockType = "heading_1"
	NotionH2BlockType               NotionBlockType = "heading_2"
	NotionH3BlockType               NotionBlockType = "heading_3"
	NotionBulletListItemBlockType   NotionBlockType = "bulleted_list_item"
	NotionNumberedListItemBlockType NotionBlockType = "numbered_list_item"
	NotionToDoBlockType             NotionBlockType = "to_do"
	NotionToggleBlockType           NotionBlockType = "toggle"
	NotionChildPageBlockType        NotionBlockType = "child_page"
	NotionChildDatabaseBlockType    NotionBlockType = "child_database"
	NotionEmbedBlockType            NotionBlockType = "embed"
	NotionImageBlockType            NotionBlockType = "image"
	NotionVideoBlockType            NotionBlockType = "video"
	NotionFileBlockType             NotionBlockType = "file"
	NotionPdfBlockType              NotionBlockType = "pdf"
	NotionBookmarkBlockType         NotionBlockType = "bookmark"
	NotionCalloutBlockType          NotionBlockType = "callout"
	NotionCodeBlockType             NotionBlockType = "code"
	NotionQuoteBlockType            NotionBlockType = "quote"
	NotionEquationBlockType         NotionBlockType = "equation"
	NotionDividerBlockType          NotionBlockType = "divider"
	NotionTableOfContentsBlockType  NotionBlockType = "table_of_contents"
	NotionColumnBlockType           NotionBlockType = "column"
	NotionColumnListBlockType       NotionBlockType = "column_list"
	NotionLinkPreviewBlockType      NotionBlockType = "link_preview"
	NotionSyncedBlockBlockType      NotionBlockType = "synced_block"
	NotionTemplateBlockType         NotionBlockType = "template"
	NotionLinkToPageBlockType       NotionBlockType = "link_to_page"
	NotionUnsupportedBlockType      NotionBlockType = "unsupported"
)
