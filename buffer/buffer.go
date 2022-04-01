package buffer

import (
	"io"
	"strconv"
	"time"
	"unicode/utf8"
)

// Buffer implements a thin wrapper around a byte slice, that supports a portion
// of the strconv's zero-allocation formatters.
type Buffer struct {
	buf []byte
}

// Bytes returns a mutable reference to the underlying byte slice. The reference
// is only valid until the next buffer modification.
func (b *Buffer) Bytes() []byte {
	return b.buf
}

// String returns a string representation of the underlying byte slice.
func (b *Buffer) String() string {
	return string(b.buf)
}

// Cap returns the capacity of the underlying byte slice.
func (b *Buffer) Cap() int {
	return cap(b.buf)
}

// Len returns the length of the underlying byte slice.
func (b *Buffer) Len() int {
	return len(b.buf)
}

// Reset resets the underlying byte slice to be empty, but retains it capacity.
func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
}

// Write appends given bytes to the underlying byte slice.
func (b *Buffer) Write(p []byte) (int, error) {
	b.buf = append(b.buf, p...)
	return len(p), nil
}

// WriteByte appends the given byte to the underlying byte slice.
func (b *Buffer) WriteByte(c byte) error {
	b.AppendByte(c)
	return nil
}

// WriteString appends the given string to the underlying byte slice.
func (b *Buffer) WriteString(s string) (int, error) {
	b.AppendString(s)
	return len(s), nil
}

// WriteTo writes the contents of the underlying byte slice to the given writer.
func (b *Buffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.buf)
	return int64(n), err
}

// AppendBool appends the string form of the given boolean to the underlying
// byte slice.
func (b *Buffer) AppendBool(v bool) {
	b.buf = strconv.AppendBool(b.buf, v)
}

// AppendByte appends the given byte to the underlying byte slice.
func (b *Buffer) AppendByte(c byte) {
	b.buf = append(b.buf, c)
}

// AppendDuration appends the string form representing the number of seconds of
// the given duration to the underlying byte slice.
func (b *Buffer) AppendDuration(d time.Duration) {
	b.buf = strconv.AppendFloat(b.buf, d.Seconds(), 'f', -1, 64)
}

// AppendFloat appends the string form of the given floating-point number to
// the underlying byte slice.
func (b *Buffer) AppendFloat(f float64, bitSize int) {
	b.buf = strconv.AppendFloat(b.buf, f, 'f', -1, bitSize)
}

// AppendInt appends the string form of the given integer to the underlying
// byte slice.
func (b *Buffer) AppendInt(i int64, base int) {
	b.buf = strconv.AppendInt(b.buf, i, base)
}

// AppendQuote appends a double-quoted string to the underlying byte slice.
func (b *Buffer) AppendQuote(s string) {
	b.buf = strconv.AppendQuote(b.buf, s)
}

// AppendRune appends the utf8 encoding of the given rune to the underlying
// byte slice.
func (b *Buffer) AppendRune(r rune) {
	b.buf = utf8.AppendRune(b.buf, r)
}

// AppendString appends the given string to the underlying byte slice.
func (b *Buffer) AppendString(s string) {
	b.buf = append(b.buf, s...)
}

// AppendTime appends the string form of the given time to the underlying byte
// slice.
func (b *Buffer) AppendTime(t time.Time, layout string) {
	b.buf = t.AppendFormat(b.buf, layout)
}

// AppendUint appends the string form of the given unsigned integer to the
// underlying byte slice.
func (b *Buffer) AppendUint(i uint64, base int) {
	b.buf = strconv.AppendUint(b.buf, i, base)
}
