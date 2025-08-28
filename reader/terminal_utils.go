package reader


const TerminalEscapeBase = "\033["

var TerminalEscapeCodes = struct {
	ClearScreen string
	HideCursor string
	ShowCursor string
	MoveCursorTopLeft string
}{
	ClearScreen: TerminalEscapeBase + "2J",
	HideCursor: TerminalEscapeBase + "?25l",
	ShowCursor: TerminalEscapeBase + "?25h",
	MoveCursorTopLeft: TerminalEscapeBase + "H",
}
