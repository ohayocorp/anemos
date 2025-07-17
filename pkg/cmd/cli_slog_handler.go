package cmd

import (
	"bytes"
	"context"
	"encoding"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/ohayocorp/anemos/pkg/core"
)

const (
	timeFormat = "[15:04:05.000]"

	reset = "\033[0m"

	black        = 30
	red          = 31
	green        = 32
	yellow       = 33
	blue         = 34
	magenta      = 35
	cyan         = 36
	lightGray    = 37
	darkGray     = 90
	lightRed     = 91
	lightGreen   = 92
	lightYellow  = 93
	lightBlue    = 94
	lightMagenta = 95
	lightCyan    = 96
	white        = 97
)

type CliSlogHandler struct {
	slogHandler slog.Handler
	buffer      *bytes.Buffer
	bufferLock  *sync.Mutex
	writer      io.Writer
	useColors   bool
	printAttrs  bool
}

func NewCliSlogHandler(writer io.Writer, handlerOptions *slog.HandlerOptions) *CliSlogHandler {
	if handlerOptions == nil {
		handlerOptions = &slog.HandlerOptions{}
	}

	buffer := &bytes.Buffer{}
	handler := &CliSlogHandler{
		buffer:      buffer,
		bufferLock:  &sync.Mutex{},
		slogHandler: slog.NewJSONHandler(buffer, handlerOptions),
		useColors:   true,
		writer:      writer,
	}

	return handler
}

func (handler *CliSlogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return handler.slogHandler.Enabled(ctx, level)
}

func (handler *CliSlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CliSlogHandler{
		slogHandler: handler.slogHandler.WithAttrs(attrs),
		buffer:      handler.buffer,
		bufferLock:  handler.bufferLock,
		writer:      handler.writer,
		useColors:   handler.useColors,
	}
}

func (handler *CliSlogHandler) WithGroup(name string) slog.Handler {
	return &CliSlogHandler{
		slogHandler: handler.slogHandler.WithGroup(name),
		buffer:      handler.buffer,
		bufferLock:  handler.bufferLock,
		writer:      handler.writer,
		useColors:   handler.useColors,
	}
}

func (handler *CliSlogHandler) Handle(context context.Context, record slog.Record) error {
	timestamp := fmt.Sprintf("%s ", record.Time.Format(timeFormat))
	level := fmt.Sprintf("%s: ", record.Level.String())

	if err := handler.write(timestamp); err != nil {
		return err
	}

	if err := handler.writeColored(handler.levelToColor(record.Level), level); err != nil {
		return err
	}

	collector := make(map[string]string)
	record.Attrs(func(attr slog.Attr) bool {
		handler.handleAttr(attr, "", collector)
		return true
	})

	if err := handler.write(handler.renderMessage(record.Message, collector)); err != nil {
		return err
	}

	if err := handler.write("\n"); err != nil {
		return err
	}

	if !handler.printAttrs {
		return nil
	}

	if len(collector) == 0 {
		return nil
	}

	attrKeys := core.SortedKeys(collector)
	maxKeyLength := 0

	for _, key := range attrKeys {
		if len(key) > maxKeyLength {
			maxKeyLength = len(key)
		}
	}

	for _, key := range attrKeys {
		value := collector[key]
		if err := handler.writeColored(lightGray, fmt.Sprintf("%*s%-*s: ", len(timestamp), "", maxKeyLength, key)); err != nil {
			return err
		}

		lines := strings.Split(value, "\n")
		for i, line := range lines {
			prefixLength := 0
			if i > 0 {
				prefixLength = len(timestamp) + maxKeyLength + 2
			}

			if err := handler.writeColored(white, fmt.Sprintf("%*s%s", prefixLength, "", line)); err != nil {
				return err
			}

			if i < len(lines)-1 {
				if err := handler.write("\n"); err != nil {
					return err
				}
			}
		}

		if err := handler.write("\n"); err != nil {
			return err
		}
	}

	return nil
}

func (handler *CliSlogHandler) handleAttrs(attrs []slog.Attr, groupIdentifier string, collector map[string]string) {
	for _, attr := range attrs {
		handler.handleAttr(attr, groupIdentifier, collector)
	}
}

func (handler *CliSlogHandler) handleAttr(attr slog.Attr, groupIdentifier string, collector map[string]string) {
	attr.Value = attr.Value.Resolve()

	if attr.Equal(slog.Attr{}) {
		return
	}

	if attr.Value.Kind() == slog.KindGroup {
		attrs := attr.Value.Group()
		if len(attrs) == 0 {
			return
		}

		if groupIdentifier != "" {
			groupIdentifier = fmt.Sprintf("%s.%s", groupIdentifier, attr.Key)
		} else {
			groupIdentifier = attr.Key
		}

		handler.handleAttrs(attrs, groupIdentifier, collector)

		return
	}

	key := attr.Key
	if groupIdentifier != "" {
		key = fmt.Sprintf("%s.%s", groupIdentifier, key)
	}

	defer func() {
		if r := recover(); r != nil {
			// If it panics with a nil pointer, the most likely cases are
			// an encoding.TextMarshaler or error fails to guard against nil,
			// in which case "<nil>" seems to be the feasible choice.
			//
			// Adapted from the code in fmt/print.go.
			if v := reflect.ValueOf(attr.Value.Any()); v.Kind() == reflect.Pointer && v.IsNil() {
				collector[key] = "<nil>"
				return
			}

			// Otherwise just print the original panic message.
			collector[key] = fmt.Sprintf("!PANIC: %v", r)
		}
	}()

	collector[key] = handler.getValueString(attr.Value)
}

func (handler *CliSlogHandler) getValueString(value slog.Value) string {
	switch value.Kind() {
	case slog.KindString:
		return value.String()
	case slog.KindInt64:
		return strconv.FormatInt(value.Int64(), 10)
	case slog.KindUint64:
		return strconv.FormatUint(value.Uint64(), 10)
	case slog.KindFloat64:
		return strconv.FormatFloat(value.Float64(), 'g', -1, 64)
	case slog.KindBool:
		return strconv.FormatBool(value.Bool())
	case slog.KindDuration:
		return value.Duration().String()
	case slog.KindTime:
		return value.Time().String()
	case slog.KindAny:
		switch anyValue := value.Any().(type) {
		case encoding.TextMarshaler:
			data, err := anyValue.MarshalText()
			if err != nil {
				break
			}
			return string(data)
		case *slog.Source:
			return handler.getSourceString(anyValue)
		}
	}

	return fmt.Sprintf("%+v", value.Any())
}

func (handler *CliSlogHandler) getSourceString(source *slog.Source) string {
	if source == nil {
		return ""
	}

	var out strings.Builder
	out.WriteString(source.File)
	out.WriteString(":")
	out.WriteString(strconv.Itoa(source.Line))
	out.WriteString(":")

	return out.String()
}

func (handler *CliSlogHandler) levelToColor(level slog.Level) int {
	switch level {
	case slog.LevelInfo:
		return cyan
	case slog.LevelWarn:
		return lightYellow
	case slog.LevelError:
		return lightRed
	}

	if level > slog.LevelError {
		return magenta
	}

	return lightGray
}

func (handler *CliSlogHandler) writeColored(colorCode int, v string) error {
	if handler.useColors {
		v = fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
	}

	return handler.write(v)
}

func (handler *CliSlogHandler) renderMessage(message string, collector map[string]string) string {
	for key, value := range collector {
		if handler.useColors {
			value = fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(lightGreen), value, reset)
		}

		message = strings.ReplaceAll(message, fmt.Sprintf("${%s}", key), value)
	}

	return message
}

func (handler *CliSlogHandler) write(v string) error {
	_, err := handler.writer.Write([]byte(v))
	return err
}
