package str

import "unicode/utf8"

func Len(s []byte) int {
	return utf8.RuneCount(s)
}
func Slen(s string) int {
	return utf8.RuneCountInString(s)
}

// start: 子串在s中开始的下标
// length: 子串长度
// return string: 子串
// return int: 实际子串长度
func SubString(s string, start int, length int) (string, int) {
	if start < 0 {
		panic("arg start error")
	}
	if length <= 0 {
		panic("arg length error")
	}
	var k int    // 遍历字符串s的字符个数
	var l, r int // l: 子字符串在s的左下标,  r: 子字符串在s的右下标
	var c int
	for i := range s {
		if k < start {
			k++
			continue
		} else if k == start {
			l = i
			r = i
			k++
		}
		_, size := utf8.DecodeRuneInString(s[i:])
		c++
		r += size
		if c >= length {
			break
		}
	}
	return s[l:r], c
}
