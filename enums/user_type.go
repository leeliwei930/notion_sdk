package enums

import "encoding/json"

type UserType uint8

const (
	Person UserType = iota
	Bot
)

var userTypeMap = map[UserType]string{
	Person: "person",
	Bot:    "bot",
}

var userTypeIndexes = map[string]UserType{
	"person": Person,
	"bot":    Bot,
}

func ParseUserType(s string) UserType {
	return userTypeIndexes[s]
}

func (u UserType) String() string {
	return userTypeMap[u]
}

func (u UserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *UserType) UnmarshalJSON(input []byte) error {
	var j string
	err := json.Unmarshal(input, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*u = ParseUserType(j)
	return nil
}
