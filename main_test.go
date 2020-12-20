package main

import "testing"

func TestTitle(t *testing.T) {
	if title("abc") != "# abc" {
		t.Fail()
	}
}

func TestHyperlink(t *testing.T) {
	if hyperlink("Some ting", "https://example.com/abc") != "[Some ting](https://example.com/abc)" {
		t.Fail()
	}
}