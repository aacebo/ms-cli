package cli

type Text string

func (self Text) String() string {
	return string(self)
}

func (self Text) Bold() Text {
	return Text(ANSI_CODE_BOLD + string(self) + ANSI_CODE_BOLD_RESET)
}

func (self Text) Dim() Text {
	return Text(ANSI_CODE_DIM + string(self) + ANSI_CODE_DIM_RESET)
}

func (self Text) Italic() Text {
	return Text(ANSI_CODE_ITALIC + string(self) + ANSI_CODE_ITALIC_RESET)
}

func (self Text) Underline() Text {
	return Text(ANSI_CODE_UNDERLINE + string(self) + ANSI_CODE_UNDERLINE_RESET)
}

func (self Text) Blink() Text {
	return Text(ANSI_CODE_BLINK + string(self) + ANSI_CODE_BLINK_RESET)
}

func (self Text) Reverse() Text {
	return Text(ANSI_CODE_REVERSE + string(self) + ANSI_CODE_REVERSE_RESET)
}

func (self Text) Hide() Text {
	return Text(ANSI_CODE_HIDE + string(self) + ANSI_CODE_HIDE_RESET)
}

func (self Text) Strike() Text {
	return Text(ANSI_CODE_STRIKE + string(self) + ANSI_CODE_STRIKE_RESET)
}

func (self Text) BlackForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_BLACK + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) BlackBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_BLACK + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) RedForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_RED + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) RedBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_RED + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) GreenForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_GREEN + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) GreenBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_GREEN + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) YellowForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_YELLOW + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) YellowBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_YELLOW + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) BlueForeground(text string) Text {
	self += Text(ANSI_CODE_FOREGROUND_BLUE + text + ANSI_CODE_FOREGROUND_RESET)
	return self
}

func (self Text) BlueBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_BLUE + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) MagentaForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_MAGENTA + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) MagentaBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_MAGENTA + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) CyanForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_CYAN + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) CyanBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_CYAN + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) WhiteForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_WHITE + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) WhiteBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_WHITE + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) DefaultForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_DEFAULT + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) DefaultBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_DEFAULT + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) EraseScreenEnd() Text {
	return Text(string(self) + ANSI_CODE_ERASE_SCREEN_END)
}

func (self Text) EraseScreenStart() Text {
	return Text(string(self) + ANSI_CODE_ERASE_SCREEN_END)
}

func (self Text) EraseScreen() Text {
	return Text(string(self) + ANSI_CODE_ERASE_SCREEN)
}

func (self Text) EraseLineEnd() Text {
	return Text(string(self) + ANSI_CODE_ERASE_LINE_END)
}

func (self Text) EraseLineStart() Text {
	return Text(string(self) + ANSI_CODE_ERASE_LINE_END)
}

func (self Text) EraseLine() Text {
	return Text(string(self) + ANSI_CODE_ERASE_LINE)
}
