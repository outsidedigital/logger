package logger

import (
	"io"

	"github.com/outsidedigital/logger/buffer"
	"github.com/outsidedigital/logger/encoding/json"
	"github.com/outsidedigital/logger/encoding/text"
)

// Writer is a generic interface of the logging writer.
type Writer interface {
	// Write encodes given fields and writes them to the destination.
	Write(...Field)
}

// WriterFunc is an adapter to allow the use of ordinary functions as writers.
type WriterFunc func(...Field)

// Write encodes given fields and writes them to the destination.
func (w WriterFunc) Write(ff ...Field) {
	w(ff...)
}

// HookWriter creates a wrapper around the given hook that outputs to
// the specified writer.
func HookWriter(w Writer, hook Hook) WriterFunc {
	return func(ff ...Field) {
		hook.Hook(w, ff...)
	}
}

var jsonPool = &buffer.Pool{}

// JSONWriter creates a new logging writer that encode entries into json format.
func JSONWriter(out io.Writer) WriterFunc {
	return func(ff ...Field) {
		buf := jsonPool.Get()
		defer jsonPool.Put(buf)

		enc := json.NewEncoder(buf)
		buf.AppendByte('{')
		for _, f := range ff {
			f.Encode(enc)
		}
		buf.AppendString("}\n")
		buf.WriteTo(out)
	}
}

var textPool = &buffer.Pool{}

// TextWriter creates a new logging writer that encode entries into text format.
func TextWriter(out io.Writer) WriterFunc {
	return func(ff ...Field) {
		buf := textPool.Get()
		defer textPool.Put(buf)

		enc := text.NewEncoder(buf)
		for _, f := range ff {
			f.Encode(enc)
		}

		if buf.Len() > 0 {
			buf.AppendByte('\n')
			buf.WriteTo(out)
		}
	}
}
