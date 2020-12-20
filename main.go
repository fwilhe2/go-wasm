package main

import (
	"fmt"
)

func main() {
	fmt.Println(title("An experiment"))
	fmt.Printf("This is an experiment with %s and %s", hyperlink("Go", "https://golang.org/"), hyperlink("WebAssembly", "https://webassembly.org/"))
}

func title(title string) string {
	return fmt.Sprintf("# %s", title)
}

func hyperlink(title string, url string) string {
	return fmt.Sprintf("[%s](%s)", title, url)
}