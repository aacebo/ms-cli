package cli

type _Text string

func Text() _Text {
	return ""
}

func (self _Text) Bold(text string) _Text {
	self += _Text("\033[1m" + text + ANSI_CODE_RESET)
	return self
}
