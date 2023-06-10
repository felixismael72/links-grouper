package response

type Error struct {
	Msg string `json:"msg"`
}

func NewError(msg string) *Error {
	return &Error{Msg: msg}
}
