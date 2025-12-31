package main

import (
  "fmt"
  "os"

  "go-memo/internal"
)

func main() {
  args := os.Args[1:]
  flags, params := internal.GetParams(args)
  mode := internal.GetMode(flags)

  fmt.Printf("flags: %q\n", flags)
  fmt.Printf("mode: %q\n", mode)
  fmt.Printf("params: %q\n", params)
}
