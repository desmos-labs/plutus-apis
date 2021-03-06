package types

var (
	ErrInvalidCommand = NewError("Invalid command. Please use !tip <amount> <user> to tip a user")
	ErrInvalidAmount  = NewError("Invalid amount provided")
)

type Error struct {
	Text string
}

func (e Error) Error() string {
	return e.Text
}

func NewError(text string) Error {
	return Error{
		Text: text,
	}
}
