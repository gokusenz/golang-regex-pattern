package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := `100.100.100.100 - - [23/Feb/2015:03:03:56 +0100] "GET /folder/file.mp3 HTTP/1.1" 206 5637064 "-" "AppleCoreMedia/1.0.0.12B466 (iPhone; U; CPU OS 8_1_3 like Mac OS X; da_dk)"`

	r, _ := regexp.Compile(`([.0-9]+) .*?\[([0-9a-zA-Z:\/+ ]+)\].*?"[A-Z]+ \/([^\/ ]+)\/([a-zA-Z0-9\-.]+).*" ([0-9]{3}) .*"(.*?)"$`)

	for index, match := range r.FindStringSubmatch(str) {
		fmt.Printf("[%d] %s\n", index, match)
	}
	// [0] 100.100.100.100 - - [23/Feb/2015:03:03:56 +0100] "GET /folder/file.mp3 HTTP/1.1" 206 5637064 "-" "AppleCoreMedia/1.0.0.12B466 (iPhone; U; CPU OS 8_1_3 like Mac OS X; da_dk)"
	// [1] 100.100.100.100
	// [2] 23/Feb/2015:03:03:56 +0100
	// [3] folder
	// [4] file.mp3
	// [5] 206
	// [6] AppleCoreMedia/1.0.0.12B466 (iPhone; U; CPU OS 8_1_3 like Mac OS X; da_dk)
}
