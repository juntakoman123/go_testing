package calc_test

import (
	"testing"

	"github.com/juntakoman123/go_testing/calc"
)

func TestSum(t *testing.T) {
	if calc.Sum(1, 2) != 3 {
		t.Fatal("sum(1,2) should be 3, but doesn't match")
	}
}

// tdt sample

func TestSumTdt(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{1, 2}, want: 3},
		{args: args{5, 5}, want: 10},
		{args: args{6, 6}, want: 12},
		{args: args{0, 2}, want: 2},
	}

	for _, tt := range tests {
		if calc.Sum(tt.args.a, tt.args.b) != tt.want {
			t.Errorf("sum(%d,%d) should be %d, but doesn't match", tt.args.a, tt.args.b, tt.want)
		}
	}

}
