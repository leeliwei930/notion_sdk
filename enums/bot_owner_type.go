package enums

type BotOwnerType int

const (
	UserOwner BotOwnerType = iota
	WorkspaceOwner
)

var BotOwnerTypeMap = map[BotOwnerType]string{
	UserOwner:      "user",
	WorkspaceOwner: "workspace",
}
