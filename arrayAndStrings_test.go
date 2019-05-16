package main

import   (
	"github.com/stretchr/testify/assert"
	"testing"
)


//{2}
func TestFirstduplicate_SmallestArray(t *testing.T) {
	expectedResult := 2
	assert.Equal(t, expectedResult, firstDuplicate(smallestArray))
}

//{2, 2}
func TestFirstduplicate_SmallerArray(t *testing.T) {
	expectedResult := 2
	assert.Equal(t, expectedResult, firstDuplicate(smallerArray))
}
// [2, 1, 3, 5, 3, 2]
func TestFirstduplicate_SmallArray(t *testing.T) {
	expectedResult := 3
	assert.Equal(t, expectedResult, firstDuplicate(smallArray))
}
//[[11 4 1] [13 7 2] [15 9 3]]
func TestRotateArray_ThreeByThree(t *testing.T) {
	var expectedResult = [][]int {
		{11, 4, 1},
		{13, 7, 2},
		{15, 9, 3},
	}
	assert.Equal(t, expectedResult, rotateArray(a))
}

func TestRotateArray_One(t *testing.T) {
	var expectedResult = [][]int {
		{1},
	}
	assert.Equal(t, expectedResult, rotateArray(b))
}

func TestRotateArray_FiveByFive(t *testing.T) {
	var expectedResult = [][]int {
		{6,8,7,6,10},
		{8,9,6,10,9},
		{6,7,3,2,6},
		{8,9,8,9,3},
		{2,9,2,7,7},
	}
	assert.Equal(t, expectedResult, rotateArray(c))
}

