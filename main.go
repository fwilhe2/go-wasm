package main

import (
	"fmt"
)

func main() {
	fmt.Println(title("An experiment"))
	fmt.Printf("This is an experiment with %s and %s\n", hyperlink("Go", "https://golang.org/"), hyperlink("WebAssembly", "https://webassembly.org/"))
	fmt.Println(floatingPointMath(0.3, 0.122))
}

func title(title string) string {
	return fmt.Sprintf("# %s", title)
}

func hyperlink(title string, url string) string {
	return fmt.Sprintf("[%s](%s)", title, url)
}

func floatingPointMath(a float64, b float64) float64 {
	return a / b
}
