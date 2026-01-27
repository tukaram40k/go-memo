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

func TestAdd(t *testing.T) {
	f1 := []string{"add"}
	p1 := []string{"arg1"}
	p2 := []string{"arg1", "arg2", "arg3"}
	p3 := []string{}

	var got string
	var execErr error

	got = captureStdout(func() {
		execErr = internal.Execute(f1, p1)
	})

	if execErr != nil {
		t.Errorf("-add with 1 param returned unexpected error: %s", execErr)
	}

	if got != "saved!\n" {
		t.Errorf("-add missing save confirmation")
	}

	got = captureStdout(func() {
		execErr = internal.Execute(f1, p2)
	})

	if execErr != nil {
		t.Errorf("-add with 2 params returned unexpected error: %s", execErr)
	}

	if got != "saved!\n" {
		t.Errorf("-add missing save confirmation")
	}

	err := internal.Execute(f1, p3)
	if err != nil {
		if err.Error() != "nothing to save!" {
			t.Errorf("-add is supposed to warn on empty args")
		}
	} else {
		t.Errorf("-add is supposed to return error on empty params")
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