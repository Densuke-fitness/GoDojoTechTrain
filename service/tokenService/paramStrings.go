package tokenService

const (
	USER_ID string = "user_id"
)

const (
	TYPE_INT Type = 1 << iota

	// TODO:
	// If　you want to declare token filed,
	// You should write TYPE~ for convert jwt's filed interface obj .
	//ex) TYPE_DATETIME
)

//Type express bit's type
type Type int
