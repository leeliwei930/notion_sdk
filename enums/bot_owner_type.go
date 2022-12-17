package enums

import "encoding/json"

type BotOwnerType uint8

const (
	UserOwner BotOwnerType = iota
	WorkspaceOwner
)

var botOwnerTypeMap = map[BotOwnerType]string{
	UserOwner:      "user",
	WorkspaceOwner: "workspace",
}

var botOwnerTypeIndexes = map[string]BotOwnerType{
	"user":      UserOwner,
	"workspace": WorkspaceOwner,
}

func ParseBotOwnerType(s string) BotOwnerType {
	return botOwnerTypeIndexes[s]
}
func (b BotOwnerType) String() string {
	return botOwnerTypeMap[b]
}

func (b BotOwnerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *BotOwnerType) UnmarshalJSON(d []byte) error {
	var j string
	err := json.Unmarshal(d, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*b = ParseBotOwnerType(j)
	return nil
}
