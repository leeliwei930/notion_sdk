package enums

type UserType int

const (
	Person UserType = iota
	Bot
)

var UserTypeMap = map[UserType]string{
	Person: "person",
	Bot:    "bot",
}
