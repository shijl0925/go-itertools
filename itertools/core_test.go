package itertools

import (
	"fmt"
	"testing"
)

func Test_MapSliceArray(t *testing.T) {
	rChan := Map(func(x int) int { return x * 3 }, []int{1, 2, 3})
	//fmt.Printf("got: %v\n", rChan)
	for _, i := range *rChan {
		fmt.Printf("got: %v\n", i)
	}
	//fmt.Printf("got: %v\n", len(*rChan))
}
