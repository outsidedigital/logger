package logger

import (
	"bytes"
	"errors"
	"fmt"
)

// Level represents a logging priority level.
type Level uint

// Well-known logging priority levels.
const (
	LevelNone Level = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	levelCount
)

// Encode encodes the logging priority level with the given encoder.
func (lvl Level) Encode(enc Encoder) {
	enc.EncodeString(FieldLevel, lvl.String())
}

// Hook intercepts the logging entry and ensure that the it has a correct
// priority level.
func (lvl Level) Hook(w Writer, ff ...Field) {
	var level Level
	for _, f := range ff {
		if field, ok := f.(Level); ok {
			level = field
		}
	}
	if level.Equal(LevelNone) || lvl.Less(level) {
		return
	}
	w.Write(ff...)
}

// Equal checks whether the logging priority level is equal to the given one.
func (lvl Level) Equal(other Level) bool {
	return lvl == other
}

// Less checks whether the logging priority level is less than the given one.
func (lvl Level) Less(other Level) bool {
	return lvl < other
}

var levelOutput = []string{
	"none",
	"error",
	"warn",
	"info",
	"debug",
	"invalid",
}

// String returns the string form of the logging priority level.
func (lvl Level) String() string {
	if lvl >= levelCount {
		return fmt.Sprintf("%s (%d)", levelOutput[levelCount], lvl)
	}
	return levelOutput[lvl]
}

// ErrLevelInvalid is returned when the logging priority level is invalid.
var ErrLevelInvalid = errors.New("invalid level")

// MarshalText marshals the logging priority level into text form.
func (lvl Level) MarshalText() ([]byte, error) {
	if lvl >= levelCount {
		return nil, ErrLevelInvalid
	}
	return []byte(levelOutput[lvl]), nil
}

var levelInput = map[string]Level{
	"none":     LevelNone,
	"disabled": LevelNone,
	"error":    LevelError,
	"warn":     LevelWarn,
	"warning":  LevelWarn,
	"info":     LevelInfo,
	"debug":    LevelDebug,
	"trace":    LevelDebug,
}

// UnmarshalText unmarshals a text form of the logging priority level.
func (lvl *Level) UnmarshalText(text []byte) error {
	v, ok := levelInput[string(bytes.ToLower(text))]
	if !ok {
		return fmt.Errorf("%s: %w", text, ErrLevelInvalid)
	}
	*lvl = v
	return nil
}
