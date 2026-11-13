package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, World!"

	// contains
	fmt.Println(strings.Contains(text, "Hello"))

	// prefix
	fmt.Println(strings.HasPrefix("http://example.com", "http://"))

	// suffix
	fmt.Println(strings.HasSuffix("example.com", ".com"))

	// index
	fmt.Println(strings.Index("Hello, World!", "World"))

	// last index
	fmt.Println(strings.LastIndex("Hello, World!", "World"))

	csvData := `name,age,city
John,25,New York
Jane,30,Los Angeles
Jim,35,Chicago`
	lines := strings.Split(csvData, "\n")
	for i, line := range lines {
		if i == 0 {
			fmt.Println("header:", line)
			continue
		}
		fields := strings.Split(line, ",")
		fmt.Println(fields[0], fields[1], fields[2])
	}
}