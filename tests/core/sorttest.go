package core

import (
	"github.com/revel/revel"
	"github.com/revel/revel/testing"
	"sort"
)

type SortTest struct {
	testing.TestSuite
}

func (t *SortTest) TestSort() {
	var intList [5]int = [5]int{1, 2, 3, 5, 4}
	sort.Ints(intList[:])

	revel.TRACE.Println(intList[:])
}
