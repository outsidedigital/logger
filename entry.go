package logger

import (
	"fmt"
	"sync"
	"time"
)

// Entry represents a structured logging entry.
type Entry struct {
	w  Writer
	ff []Field
}

var entryPool = &sync.Pool{
	New: func() any {
		return Entry{}
	},
}

// NewEntry creates a new entry that outputs to the given writer.
func NewEntry(w Writer) Entry {
	entry, _ := entryPool.Get().(Entry)
	entry.w = w
	return entry
}

// Message appends a given message to the entry and sends it to the underlying
// writer. Once this method is called, the entry should be disposed.
func (e Entry) Message(msg string) {
	e.ff = append(e.ff, Message(msg))
	e.w.Write(e.ff...)
	e.Discard()
}

// Messagef appends a given formatted message to the entry and sends it to the
// underlying writer. Once this method is called, the entry should be disposed.
func (e Entry) Messagef(format string, a ...any) {
	e.Message(fmt.Sprintf(format, a...))
}

// Discard discards the entry, so it won't be logged. Once this method is
// called, the entry should be disposed.
func (e Entry) Discard() {
	entryPool.Put(e.Reset())
}

// Reset resets previously stored fields.
func (e Entry) Reset() Entry {
	e.ff = e.ff[:0]
	return e
}

// Bool appends a new field with the given key and boolean value.
func (e Entry) Bool(key string, b bool) Entry {
	e.ff = append(e.ff, Bool(key, b))
	return e
}

// Bytes appends a new field with the given key and bytes value.
func (e Entry) Bytes(key string, p []byte) Entry {
	e.ff = append(e.ff, Bytes(key, p))
	return e
}

// Caller appends a new field with current file and line number.
func (e Entry) Caller(skip int) Entry {
	e.ff = append(e.ff, Caller(skip+1))
	return e
}

// Duration appends a new field with the given key and duration value.
func (e Entry) Duration(key string, d time.Duration) Entry {
	e.ff = append(e.ff, Duration(key, d))
	return e
}

// Error appends given error to the entry.
func (e Entry) Error(err error) Entry {
	e.ff = append(e.ff, Error{err})
	return e
}

// Errorf appends a new formatted error to the entry.
func (e Entry) Errorf(format string, a ...any) Entry {
	//nolint:goerr113 // Errorf is a wrapper for errorf.
	e.ff = append(e.ff, Error{fmt.Errorf(format, a...)})
	return e
}

// Float32 appends a new field with the given key and float32 value.
func (e Entry) Float32(key string, f float32) Entry {
	e.ff = append(e.ff, Float32(key, f))
	return e
}

// Float64 appends a new field with the given key and float64 value.
func (e Entry) Float64(key string, f float64) Entry {
	e.ff = append(e.ff, Float64(key, f))
	return e
}

// Int appends a new field with the given key and integer value.
func (e Entry) Int(key string, i int) Entry {
	e.ff = append(e.ff, Int(key, i))
	return e
}

// Int32 appends a new field with the given key and int32 value.
func (e Entry) Int32(key string, i int32) Entry {
	e.ff = append(e.ff, Int32(key, i))
	return e
}

// Int64 appends a new field with the given key and int64 value.
func (e Entry) Int64(key string, i int64) Entry {
	e.ff = append(e.ff, Int64(key, i))
	return e
}

// Name appends a new field with the given logger name.
func (e Entry) Name(name string) Entry {
	e.ff = append(e.ff, Name(name))
	return e
}

// Span appends a new time span field that begins at the current time.
func (e Entry) Span() Entry {
	e.ff = append(e.ff, Span(time.Now()))
	return e
}

// String appends a new field with the given key and string value.
func (e Entry) String(key, s string) Entry {
	e.ff = append(e.ff, String(key, s))
	return e
}

// Stringer appends a new field with the given key and value that implements
// stringer interface.
func (e Entry) Stringer(key string, v fmt.Stringer) Entry {
	e.ff = append(e.ff, String(key, v.String()))
	return e
}

// Stringf appends a new field with the given key and formatted string value.
func (e Entry) Stringf(key, format string, a ...any) Entry {
	e.ff = append(e.ff, String(key, fmt.Sprintf(format, a...)))
	return e
}

// Time appends a new field with the given key and time value.
func (e Entry) Time(key string, t time.Time) Entry {
	e.ff = append(e.ff, Time(key, t))
	return e
}

// Timestamp appends a new field with the current time.
func (e Entry) Timestamp() Entry {
	e.ff = append(e.ff, Timestamp{})
	return e
}

// Uint appends a new field with the given key and unsigned integer value.
func (e Entry) Uint(key string, i uint) Entry {
	e.ff = append(e.ff, Uint(key, i))
	return e
}

// Uint32 appends a new field with the given key and uint32 value.
func (e Entry) Uint32(key string, i uint32) Entry {
	e.ff = append(e.ff, Uint32(key, i))
	return e
}

// Uint64 appends a new field with the given key and uint64 value.
func (e Entry) Uint64(key string, i uint64) Entry {
	e.ff = append(e.ff, Uint64(key, i))
	return e
}

// With appends given fields to the entry.
func (e Entry) With(ff ...Field) Entry {
	e.ff = append(e.ff, ff...)
	return e
}
