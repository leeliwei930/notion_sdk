package models

import (
	"time"

	"github.com/google/uuid"
)

type NotionBlock struct {
	Object           string                       `json:"object,omitempty"`
	ID               string                       `json:"id,omitempty"`
	Type             NotionBlockType              `json:"type,omitempty"`
	CreatedTime      *time.Time                   `json:"created_time,omitempty"`
	LastEditedTime   *time.Time                   `json:"last_edited_time,omitempty"`
	Archived         bool                         `json:"archived,omitempty"`
	HasChildren      bool                         `json:"has_children,omitempty"`
	Paragraph        *NotionParagraphBlock        `json:"paragraph,omitempty"`
	Heading1         *NotionHeading1Block         `json:"heading_1,omitempty"`
	Heading2         *NotionHeading2Block         `json:"heading_2,omitempty"`
	Heading3         *NotionHeading3Block         `json:"heading_3,omitempty"`
	Callout          *NotionCalloutBlock          `json:"callout,omitempty"`
	Quote            *NotionQuoteBlock            `json:"quote,omitempty"`
	BulletedListItem *NotionBulletedListItemBlock `json:"bulleted_list_item,omitempty"`
	NumberedListItem *NotionNumberedListItemBlock `json:"numbered_list_item,omitempty"`
	ToDo             *NotionToDoBlock             `json:"to_do,omitempty"`
	Toggle           *NotionToggleBlock           `json:"toggle,omitempty"`
	Code             *NotionCodeBlock             `json:"code,omitempty"`
	ChildPage        *NotionChildPageBlock        `json:"child_page,omitempty"`
	ChildDatabase    *NotionChildDatabaseBlock    `json:"child_database,omitempty"`
	Embed            *NotionEmbedBlock            `json:"embed,omitempty"`
	Image            *NotionFile                  `json:"image,omitempty"`
	Video            *NotionFile                  `json:"video,omitempty"`
	File             *NotionFile                  `json:"file,omitempty"`
	PDF              *NotionFile                  `json:"pdf,omitempty"`
	Bookmark         *NotionBookmarkBlock         `json:"bookmark,omitempty"`
	Equation         *NotionEquation              `json:"equation,omitempty"`
	Divider          *map[string]string           `json:"divider,omitempty"`
	TableOfContents  *map[string]string           `json:"table_of_contents,omitempty"`
	Breadcrumb       *map[string]string           `json:"breadcrumb,omitempty"`
	ColumnList       *map[string]string           `json:"column_list,omitempty"`
	Column           *map[string]string           `json:"column,omitempty"`
	LinkPreview      *NotionLinkPreviewBlock      `json:"link_preview,omitempty"`
	Template         *NotionTemplateBlock         `json:"template,omitempty"`
	LinkToPage       *NotionLinkToPageBlock       `json:"link_to_page,omitempty"`
	SyncedBlock      *NotionSyncedBlockBlocks     `json:"synced_block,omitempty"`
}

type NotionParagraphBlock struct {
	Text     *[]NotionRichText `json:"text,omitempty"`
	Children *[]NotionBlock    `json:"children,omitempty"`
}

type NotionHeading1Block struct {
	Text *[]NotionRichText `json:"text"`
}

type NotionHeading2Block NotionHeading1Block

type NotionHeading3Block NotionHeading1Block

type NotionCalloutBlock struct {
	Text     *[]NotionRichText `json:"text"`
	Icon     *[]NotionIcon     `json:"icon"`
	Children *[]NotionBlock    `json:"children"`
}

type NotionQuoteBlock struct {
	Text     *[]NotionRichText `json:"text"`
	Children *[]NotionBlock    `json:"children"`
}

type NotionToDoBlock struct {
	Text     *[]NotionRichText `json:"text"`
	Children *[]NotionBlock    `json:"children"`
	Checked  bool              `json:"checked,omitempty"`
}
type NotionBulletedListItemBlock NotionQuoteBlock
type NotionNumberedListItemBlock NotionQuoteBlock
type NotionToggleBlock NotionQuoteBlock
type NotionTemplateBlock NotionQuoteBlock

type NotionCodeBlock struct {
	Text     *[]NotionRichText `json:"text"`
	Language string            `json:"language"`
}

type NotionChildPageBlock struct {
	Title string `json:"title"`
}

type NotionChildDatabaseBlock NotionChildPageBlock

type NotionEmbedBlock struct {
	Url string `json:"url"`
}

type NotionLinkPreviewBlock NotionEmbedBlock
type NotionBookmarkBlock struct {
	Url     string            `json:"url"`
	Caption *[]NotionRichText `json:"caption"`
}

type NotionLinkToPageBlock struct {
	Type       *NotionPageParentType `json:"type"`
	PageID     *uuid.UUID            `json:"page_id,omitempty"`
	DatabaseID *uuid.UUID            `json:"database_id,omitempty"`
}

type NotionSyncedBlockBlocks struct {
	SyncedFrom *NotionSyncedFromBlock `json:"synced_from,omitempty"`
}

type NotionSyncedFromBlock struct {
	BlockID *uuid.UUID `json:"block_id,omitempty"`
}
