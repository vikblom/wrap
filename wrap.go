package wrap

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type locError struct {
	inner error
	pc    uintptr
}

// Errorf behaves like fmt.Errorf but prepends callsite info.
func Errorf(format string, args ...any) error {
	pcs := make([]uintptr, 1)
	runtime.Callers(2, pcs)
	return &locError{
		inner: fmt.Errorf(format, args...),
		pc:    pcs[0],
	}
}

func (e *locError) Error() string {
	pcs := []uintptr{e.pc}
	frame, _ := runtime.CallersFrames(pcs).Next()
	filename := filepath.Base(frame.File)
	return fmt.Sprintf("(%s:%d): ", filename, frame.Line) + e.inner.Error()
}

func (e *locError) Unwrap() error {
	return e.inner
}
