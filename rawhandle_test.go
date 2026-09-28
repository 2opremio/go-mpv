package mpv_test

import (
	"testing"

	"github.com/gen2brain/go-mpv"
)

func TestRawHandle(t *testing.T) {
	a := mpv.New()
	defer a.TerminateDestroy()
	b := mpv.New()
	defer b.TerminateDestroy()

	if a.RawHandle() == nil {
		t.Fatal("RawHandle is nil")
	}
	if a.RawHandle() != a.RawHandle() {
		t.Error("RawHandle changed between calls")
	}
	if a.RawHandle() == b.RawHandle() {
		t.Error("two clients share a RawHandle")
	}
}
