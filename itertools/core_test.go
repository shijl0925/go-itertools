package itertools

import (
	"fmt"
	"testing"
)

func Test_MapSliceArray(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(slice)
	rChan := Map(func(x int) int { return x * 3 }, slice)

	for _, i := range rChan {
		fmt.Printf("got: %v\n", i)
	}

	if len(rChan) != n {
		t.Errorf("expected %v, got %v", n, len(rChan))
	}
}
