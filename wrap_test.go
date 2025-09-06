package wrap_test

import (
	"errors"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/vikblom/wrap"
)

func TestWrapErrorf(t *testing.T) {
	err := wrap.Errorf("kaboom: %d", 123)
	got := err.Error()
	want := "(wrap_test.go:14) kaboom: 123"
	if d := cmp.Diff(want, got); d != "" {
		t.Fatalf("error mismatch (-want, +got):\n%s", d)
	}
}

func TestWrapErrorsIs(t *testing.T) {
	err := wrap.Errorf("foo: %w", io.EOF)

	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected errors.Is(err, io.EOF) to be true")
	}
}
