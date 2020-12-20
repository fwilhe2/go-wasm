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

func TestFloatingPointMath(t *testing.T) {
	if floatingPointMath(0.3, 0.122) != 2.459016393442623 {
		t.Fail()
	}
}
