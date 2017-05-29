package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  Is golang good for real world web apps?  "
	// <Is golang good for real world web apps?>
	fmt.Printf("<%v>", strings.TrimSpace(s))
}
