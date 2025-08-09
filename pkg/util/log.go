package util

import "log/slog"

const (
	NoLineBreakAttrKey = "no_line_break"
)

func SlogNoLineBreakAttr() slog.Attr {
	return slog.Attr{
		Key:   NoLineBreakAttrKey,
		Value: slog.BoolValue(true),
	}
}
