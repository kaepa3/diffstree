package sourcetree

func CheckArg(args []string) bool {
	count := len(args)
	if count < 2 || count > 3 {
		return false
	}
	return true
}
