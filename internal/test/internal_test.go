package internal_test

import (
	"go-memo/internal"
	"io"
	"os"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	flags := []string{"help"}
	params := []string{""}

	var got string
	var execErr error
	
	got = captureStdout(func() {
		execErr = internal.Execute(flags, params)
	})

	if execErr != nil {
		t.Errorf("-h crashed: %v", execErr)
	}

	if !strings.Contains(got, "Usage: go-memo [options] [args]") {
		t.Errorf("-h returned unexpected message: %s", got)
	}
}

func TestError(t *testing.T) {
	flags := []string{"error"}
	params := []string{"bad-arg"}

	var got string
	var execErr error
	
	got = captureStdout(func() {
		execErr = internal.Execute(flags, params)
	})

	if execErr != nil {
		t.Errorf("error flag crashed: %v", execErr)
	}

	if !strings.Contains(got, "Usage: go-memo [options] [args]") {
		t.Errorf("-h returned unexpected message: %s", got)
	}
}

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()
	
	w.Close()
	os.Stdout = old

	out, _ := io.ReadAll(r)
	return string(out)
}