package Func

func Suggest(Error int) string {
	var Message string
	switch Error {
	case 1:
		Message = "Please enter a letter"
	case 2:
		Message = "Letter already used, Please try something else"
	default:
		return ""
	}
	return Message
}
