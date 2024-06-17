package main

import (
	"fmt"
	"strings"
)

func main() {
	tolower := strings.ToLower("To Hello")
	toupper := strings.ToUpper("To Hello")
	totitle := strings.ToTitle("To Hello")
	fmt.Println("ToLower:", tolower, "\nToUpper:", toupper, "\nToTitle:", totitle)
}
