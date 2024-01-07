package cat

import "bytes"

// +=演算子による文字列結合
func Cat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}
	return r
}

// bytes.Buffer による文字列結合
func Buf(ss ...string) string {
	var b bytes.Buffer
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}
