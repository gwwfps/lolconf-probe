package main

import (
  "fmt"
  "github.com/gwwfps/lolconf-probe/display"
)

func main() {
  resolutions := display.ListAvailableResolutions()
  for _, res := range resolutions {
    fmt.Printf("%dx%d\n", res.Width, res.Height)
  }
}
