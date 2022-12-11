package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_sdk/enums"
)

type Block struct {
	Object           string                 `json:"object,omitempty"`
	ID               string                 `json:"id,omitempty"`
	Type             enums.BlockType        `json:"type,omitempty"`
	CreatedTime      *time.Time             `json:"created_time,omitempty"`
	LastEditedTime   *time.Time             `json:"last_edited_time,omitempty"`
	Archived         bool                   `json:"archived,omitempty"`
	HasChildren      bool                   `json:"has_children,omitempty"`
	Paragraph        *ParagraphBlock        `json:"paragraph,omitempty"`
	Heading1         *Heading1Block         `json:"heading_1,omitempty"`
	Heading2         *Heading2Block         `json:"heading_2,omitempty"`
	Heading3         *Heading3Block         `json:"heading_3,omitempty"`
	Callout          *CalloutBlock          `json:"callout,omitempty"`
	Quote            *QuoteBlock            `json:"quote,omitempty"`
	BulletedListItem *BulletedListItemBlock `json:"bulleted_list_item,omitempty"`
	NumberedListItem *NumberedListItemBlock `json:"numbered_list_item,omitempty"`
	ToDo             *ToDoBlock             `json:"to_do,omitempty"`
	Toggle           *ToggleBlock           `json:"toggle,omitempty"`
	Code             *CodeBlock             `json:"code,omitempty"`
	ChildPage        *ChildPageBlock        `json:"child_page,omitempty"`
	ChildDatabase    *ChildDatabaseBlock    `json:"child_database,omitempty"`
	Embed            *EmbedBlock            `json:"embed,omitempty"`
	Image            *File                  `json:"image,omitempty"`
	Video            *File                  `json:"video,omitempty"`
	File             *File                  `json:"file,omitempty"`
	PDF              *File                  `json:"pdf,omitempty"`
	Bookmark         *BookmarkBlock         `json:"bookmark,omitempty"`
	Equation         *Equation              `json:"equation,omitempty"`
	Divider          *map[string]string     `json:"divider,omitempty"`
	TableOfContents  *map[string]string     `json:"table_of_contents,omitempty"`
	Breadcrumb       *map[string]string     `json:"breadcrumb,omitempty"`
	ColumnList       *map[string]string     `json:"column_list,omitempty"`
	Column           *map[string]string     `json:"column,omitempty"`
	LinkPreview      *LinkPreviewBlock      `json:"link_preview,omitempty"`
	Template         *TemplateBlock         `json:"template,omitempty"`
	LinkToPage       *LinkToPageBlock       `json:"link_to_page,omitempty"`
	SyncedBlock      *SyncedBlockBlocks     `json:"synced_block,omitempty"`
}

type ParagraphBlock struct {
	RichText *[]RichText `json:"rich_text,omitempty"`
	Children *[]Block    `json:"children,omitempty"`
}

type Heading1Block struct {
	RichText *[]RichText `json:"rich_text"`
}

type Heading2Block Heading1Block

type Heading3Block Heading1Block

type CalloutBlock struct {
	RichText *[]RichText `json:"rich_text"`
	Icon     *[]Icon     `json:"icon"`
	Children *[]Block    `json:"children"`
}

type QuoteBlock struct {
	RichText *[]RichText `json:"rich_text"`
	Children *[]Block    `json:"children"`
}

type ToDoBlock struct {
	RichText *[]RichText `json:"rich_text"`
	Children *[]Block    `json:"children"`
	Checked  bool        `json:"checked,omitempty"`
}
type BulletedListItemBlock QuoteBlock
type NumberedListItemBlock QuoteBlock
type ToggleBlock QuoteBlock
type TemplateBlock QuoteBlock

type CodeBlock struct {
	RichText *[]RichText `json:"rich_text"`
	Language string      `json:"language"`
}

type ChildPageBlock struct {
	Title string `json:"title"`
}

type ChildDatabaseBlock ChildPageBlock

type EmbedBlock struct {
	Url string `json:"url"`
}

type LinkPreviewBlock EmbedBlock
type BookmarkBlock struct {
	Url     string      `json:"url"`
	Caption *[]RichText `json:"caption"`
}

type LinkToPageBlock struct {
	Type       *enums.PageParentType `json:"type"`
	PageID     *uuid.UUID            `json:"page_id,omitempty"`
	DatabaseID *uuid.UUID            `json:"database_id,omitempty"`
}

type SyncedBlockBlocks struct {
	SyncedFrom *SyncedFromBlock `json:"synced_from,omitempty"`
}

type SyncedFromBlock struct {
	BlockID *uuid.UUID `json:"block_id,omitempty"`
}
