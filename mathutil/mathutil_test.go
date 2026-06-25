package mathutil

import "testing"

func TestIsPrime(t *testing.T) {
	cases := map[int]bool{7: true, 8: false, 1: false, 4: false}
	for input, want := range cases {
		if got := IsPrime(input); got != want {
			t.Errorf("isPrime(%d) = %v; want %v", input, got, want)
		}
	}
}
