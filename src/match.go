package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, err := regexp.Compile(`Nattawut`)

	if err != nil {
		fmt.Printf("Error : Don't have Nattawut")
		return
	}

	if r.MatchString("Hello Regular Expression. from Nattawut") == true {
		fmt.Printf("Match ")
	} else {
		fmt.Printf("No match ")
	}
}
