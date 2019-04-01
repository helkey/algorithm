// CGO_ENABLED=0 go test

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	// Empty list
	lst := list{}
	assert.Equal(t, flatten(lst), list{}, "[] => []")

	// Nominal test case
	lst = list{list{1, 2, list{3}}, 4}
	f := flatten(lst)
	assert.Equal(t, f, list{1, 2, 3, 4}, "in: [[1,2,[3]],4]")

	// Interior empty list
	lst = list{list{1, 2, list{}}, 3}
	f = flatten(lst)
	assert.Equal(t, f, list{1, 2, 3}, "in: [[1,2,[]],3]")

	// Mix of int, str, float, bool
	lst = list{list{1, 2.1, list{"3"}}, true}
	f = flatten(lst)
	assert.Equal(t, f, list{1, 2.1, "3", true}, "in: [[1,2,['3']],4]")
}
