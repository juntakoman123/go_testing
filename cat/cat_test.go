package cat_test

import (
	"fmt"
	"testing"

	"github.com/juntakoman123/go_testing/cat"
)

func seed(n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, "a")

	}
	return s
}

func bench(b *testing.B, n int, f func(...string) string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(seed(n)...)
	}
}

func BenchmarkCat3(b *testing.B) {
	bench(b, 3, cat.Cat)
}

func BenchmarkBuf3(b *testing.B) {
	bench(b, 3, cat.Buf)
}

func BenchmarkCat100(b *testing.B) {
	bench(b, 100, cat.Cat)
}

func BenchmarkBuf100(b *testing.B) {
	bench(b, 100, cat.Buf)
}

func BenchmarkCat10000(b *testing.B) {
	bench(b, 10000, cat.Cat)
}

func BenchmarkBuf10000(b *testing.B) {
	bench(b, 10000, cat.Buf)
}

func BenchmarkConcatenete(b *testing.B) {
	benchCases := []struct {
		name string
		n    int
		f    func(...string) string
	}{
		{"Cat", 3, cat.Cat},
		{"Buf", 3, cat.Buf},
		{"Cat", 100, cat.Cat},
		{"Buf", 100, cat.Buf},
		{"Cat", 10000, cat.Cat},
		{"Buf", 10000, cat.Buf},
	}
	for _, c := range benchCases {
		b.Run(fmt.Sprintf("%s%d", c.name, c.n),
			func(b *testing.B) { bench(b, c.n, c.f) })
	}
}

func TestCat(t *testing.T) {
	input := []string{"a", "b", "c"}
	want := "abc"

	if got := cat.Cat(input...); got != want {
		t.Errorf("want %v got %v", want, got)
	}

}

func TestBuf(t *testing.T) {
	input := []string{"a", "b", "c"}
	want := "abc"

	if got := cat.Buf(input...); got != want {
		t.Errorf("want %v got %v", want, got)
	}

}
