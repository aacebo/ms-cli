package cli

type _Text string

func Text() _Text {
	return ""
}

func (self _Text) Bold(text string) _Text {
	self += _Text(ANSI_CODE_BOLD + text + ANSI_CODE_BOLD_RESET)
	return self
}

func (self _Text) Dim(text string) _Text {
	self += _Text(ANSI_CODE_DIM + text + ANSI_CODE_DIM_RESET)
	return self
}

func (self _Text) Italic(text string) _Text {
	self += _Text(ANSI_CODE_ITALIC + text + ANSI_CODE_ITALIC_RESET)
	return self
}

func (self _Text) Underline(text string) _Text {
	self += _Text(ANSI_CODE_UNDERLINE + text + ANSI_CODE_UNDERLINE_RESET)
	return self
}

func (self _Text) Blink(text string) _Text {
	self += _Text(ANSI_CODE_BLINK + text + ANSI_CODE_BLINK_RESET)
	return self
}

func (self _Text) Reverse(text string) _Text {
	self += _Text(ANSI_CODE_REVERSE + text + ANSI_CODE_REVERSE_RESET)
	return self
}

func (self _Text) Hide(text string) _Text {
	self += _Text(ANSI_CODE_HIDE + text + ANSI_CODE_HIDE_RESET)
	return self
}

func (self _Text) Strike(text string) _Text {
	self += _Text(ANSI_CODE_STRIKE + text + ANSI_CODE_STRIKE_RESET)
	return self
}

func (self _Text) BlackForeground(text string) _Text {
	self += _Text(ANSI_CODE_FOREGROUND_BLACK + text + ANSI_CODE_FOREGROUND_RESET)
	return self
}

func (self _Text) BlackBackground(text string) _Text {
	self += _Text(ANSI_CODE_BACKGROUND_BLACK + text + ANSI_CODE_BACKGROUND_RESET)
	return self
}

func (self _Text) RedForeground(text string) _Text {
	self += _Text(ANSI_CODE_FOREGROUND_RED + text + ANSI_CODE_FOREGROUND_RESET)
	return self
}

func (self _Text) RedBackground(text string) _Text {
	self += _Text(ANSI_CODE_BACKGROUND_RED + text + ANSI_CODE_BACKGROUND_RESET)
	return self
}

func (self _Text) GreenForeground(text string) _Text {
	self += _Text(ANSI_CODE_FOREGROUND_GREEN + text + ANSI_CODE_FOREGROUND_RESET)
	return self
}

func (self _Text) GreenBackground(text string) _Text {
	self += _Text(ANSI_CODE_BACKGROUND_GREEN + text + ANSI_CODE_BACKGROUND_RESET)
	return self
}

func (self _Text) YellowForeground(text string) _Text {
	self += _Text(ANSI_CODE_FOREGROUND_YELLOW + text + ANSI_CODE_FOREGROUND_RESET)
	return self
}

func (self _Text) YellowBackground(text string) _Text {
	self += _Text(ANSI_CODE_BACKGROUND_YELLOW + text + ANSI_CODE_BACKGROUND_RESET)
	return self
}

func (self _Text) BlueForeground(text string) _Text {
	self += _Text(ANSI_CODE_FOREGROUND_BLUE + text + ANSI_CODE_FOREGROUND_RESET)
	return self
}

func (self _Text) BlueBackground(text string) _Text {
	self += _Text(ANSI_CODE_BACKGROUND_BLUE + text + ANSI_CODE_BACKGROUND_RESET)
	return self
}

func (self _Text) MagentaForeground(text string) _Text {
	self += _Text(ANSI_CODE_FOREGROUND_MAGENTA + text + ANSI_CODE_FOREGROUND_RESET)
	return self
}

func (self _Text) MagentaBackground(text string) _Text {
	self += _Text(ANSI_CODE_BACKGROUND_MAGENTA + text + ANSI_CODE_BACKGROUND_RESET)
	return self
}

func (self _Text) CyanForeground(text string) _Text {
	self += _Text(ANSI_CODE_FOREGROUND_CYAN + text + ANSI_CODE_FOREGROUND_RESET)
	return self
}

func (self _Text) CyanBackground(text string) _Text {
	self += _Text(ANSI_CODE_BACKGROUND_CYAN + text + ANSI_CODE_BACKGROUND_RESET)
	return self
}

func (self _Text) WhiteForeground(text string) _Text {
	self += _Text(ANSI_CODE_FOREGROUND_WHITE + text + ANSI_CODE_FOREGROUND_RESET)
	return self
}

func (self _Text) WhiteBackground(text string) _Text {
	self += _Text(ANSI_CODE_BACKGROUND_WHITE + text + ANSI_CODE_BACKGROUND_RESET)
	return self
}

func (self _Text) DefaultForeground(text string) _Text {
	self += _Text(ANSI_CODE_FOREGROUND_DEFAULT + text + ANSI_CODE_FOREGROUND_RESET)
	return self
}

func (self _Text) DefaultBackground(text string) _Text {
	self += _Text(ANSI_CODE_BACKGROUND_DEFAULT + text + ANSI_CODE_BACKGROUND_RESET)
	return self
}
