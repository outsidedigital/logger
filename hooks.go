package logger

// Hook is a generic interface of the logging hook.
type Hook interface {
	// Hook intercepts the logging entry.
	Hook(Writer, ...Field)
}

// HookFunc is an adapter to allow the use of ordinary functions as hooks.
type HookFunc func(Writer, ...Field)

// Hook intercepts the logging entry.
func (h HookFunc) Hook(w Writer, ff ...Field) {
	h(w, ff...)
}
