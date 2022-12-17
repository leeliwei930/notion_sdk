package enums

type BotOwnerType uint8

const (
	UserOwner BotOwnerType = iota
	WorkspaceOwner
)

var BotOwnerTypeMap = map[BotOwnerType]string{
	UserOwner:      "user",
	WorkspaceOwner: "workspace",
}

func (b BotOwnerType) String() string {
	return BotOwnerTypeMap[b]
}
