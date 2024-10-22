package loggers

import (
	"log/slog"
	"strings"
)

func ReplaceOption(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.SourceKey {
		source := a.Value.Any().(*slog.Source)
		userName := "username"

		if strings.Contains(source.File, userName) {
			startIndex := strings.Index(source.File, userName)
			if startIndex != -1 {
				newPath := source.File[startIndex+len(userName):]
				source.File = newPath
				return slog.Attr{Key: a.Key, Value: a.Value}
			}
		}
	}
	return a
}
