package enums

type UserType uint8

const (
	Person UserType = iota
	Bot
)

var UserTypeMap = map[UserType]string{
	Person: "person",
	Bot:    "bot",
}

func (u UserType) String() string {
	return UserTypeMap[u]
}
