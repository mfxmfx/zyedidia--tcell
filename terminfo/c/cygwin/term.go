// Generated automatically.  DO NOT HAND-EDIT.

package cygwin

import "github.com/zyedidia/tcell/terminfo"

func init() {

	// ANSI emulation for Cygwin
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:         "cygwin",
		Colors:       8,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[J",
		EnterCA:      "\x1b7\x1b[?47h",
		ExitCA:       "\x1b[2J\x1b[?47l\x1b8",
		AttrOff:      "\x1b[0;10m",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Reverse:      "\x1b[7m",
		SetFg:        "\x1b[3%p1%dm",
		SetBg:        "\x1b[4%p1%dm",
		SetFgBg:      "\x1b[3%p1%d;4%p2%dm",
		PadChar:      "\x00",
		AltChars:     "+\x10,\x11-\x18.\x190\xdb`\x04a\xb1f\xf8g\xf1h\xb0j\xd9k\xbfl\xdam\xc0n\xc5o~p\xc4q\xc4r\xc4s_t\xc3u\xb4v\xc1w\xc2x\xb3y\xf3z\xf2{\xe3|\xd8}\x9c~\xfe",
		EnterAcs:     "\x1b[11m",
		ExitAcs:      "\x1b[10m",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b",
		CursorUp1:    "\x1b[A",
		KeyUp:        "\x1b[A",
		KeyDown:      "\x1b[B",
		KeyRight:     "\x1b[C",
		KeyLeft:      "\x1b[D",
		KeyInsert:    "\x1b[2~",
		KeyDelete:    "\x1b[3~",
		KeyBackspace: "\b",
		KeyHome:      "\x1b[1~",
		KeyEnd:       "\x1b[4~",
		KeyPgUp:      "\x1b[5~",
		KeyPgDn:      "\x1b[6~",
		KeyF1:        "\x1b[[A",
		KeyF2:        "\x1b[[B",
		KeyF3:        "\x1b[[C",
		KeyF4:        "\x1b[[D",
		KeyF5:        "\x1b[[E",
		KeyF6:        "\x1b[17~",
		KeyF7:        "\x1b[18~",
		KeyF8:        "\x1b[19~",
		KeyF9:        "\x1b[20~",
		KeyF10:       "\x1b[21~",
		KeyF11:       "\x1b[23~",
		KeyF12:       "\x1b[24~",
		KeyF13:       "\x1b[25~",
		KeyF14:       "\x1b[26~",
		KeyF15:       "\x1b[28~",
		KeyF16:       "\x1b[29~",
		KeyF17:       "\x1b[31~",
		KeyF18:       "\x1b[32~",
		KeyF19:       "\x1b[33~",
		KeyF20:       "\x1b[34~",
	})
}
