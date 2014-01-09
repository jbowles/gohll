package gohll

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSparse(t *testing.T) {
	h, err := NewHLL(4, 1000)
	assert.Equal(t, err, nil, "Could not create HLL")

	for i := 0; i <= 500; i += 1 {
		h.Add(fmt.Sprintf("%d", i))
	}

	assert.Equal(t, h.format, SPARSE, "Not using sparse mode")

	c := h.Cardinality()

	fmt.Println("m2: ", h.m2)
	fmt.Println("Sizeof sparseList: ", h.sparseList.Len())
	fmt.Println("Sizeof tempSet: ", h.tempSet.Len())
    //for _, v := range h.sparseList.Data {
    //    fmt.Printf("%0.32b\n", v)
    //}
	assert.Equal(t, c, 0, "blah")
}