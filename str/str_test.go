package str

import (
	"testing"
)

func TestSlen(t *testing.T) {
	var useCases = []struct {
		s string
		i int
	}{
		{"123", 3}, {"1234", 4}, {"", 0}, {"dadqw", 5}, {"Go语言社区", 6},
	}
	for _, uc := range useCases {
		l := Slen(uc.s)
		if l != uc.i {
			t.Fail()
		}
	}
}

func TestLen(t *testing.T) {
	var useCases = []struct {
		s []byte
		i int
	}{
		{[]byte("123"), 3}, {[]byte("1234"), 4}, {[]byte(""), 0}, {[]byte("dadqw"), 5}, {[]byte("Go语言社区"), 6},
	}
	for _, uc := range useCases {
		l := Len(uc.s)
		if l != uc.i {
			t.Fail()
		}
	}
}

func TestSubString(t *testing.T) {
	var useCases = []struct {
		s         string
		start     int
		length    int
		result    string
		resultLen int
	}{
		{"", 0, 1, "", 0},
		{"123", 0, 2, "12", 2},
		{"1234", 1, 3, "234", 3},
		{"dadqw", 4, 1, "w", 1},
		{"dadqw", 4, 2, "w", 1},
		{"dadqw", 4, 4, "w", 1},
		{"dadqw", 5, 1, "", 0},
		{"dadqw", 5, 5, "", 0},
		{"Go语言社区", 6, 1, "", 0},
		{"Go语言社区", 0, 3, "Go语", 3},
		{"Go语言社区", 0, 8, "Go语言社区", 6},
		{"Go语言社区", 2, 2, "语言", 2},
		{"Go语言社区", 2, 8, "语言社区", 4},
		{"Go语言社区", 5, 1, "区", 1},
		{"Go语言社区", 6, 1, "", 0},
		{"Go语言社区", 6, 2, "", 0},
		{"Go语言社区", 6, 6, "", 0},
		{"Go语言社区，，", 2, 8, "语言社区，，", 6},
	}
	for _, uc := range useCases {
		l, c := SubString(uc.s, uc.start, uc.length)
		if l != uc.result || c != uc.resultLen {
			t.Errorf("%#v: %v, %v", uc, l, c)
		}
	}
}
