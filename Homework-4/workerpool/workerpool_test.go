package workerpool_test

import (
	"github.com/alcor67/repo-Go-level-2-home-work/Homework-4/workerpool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_workerpool(t *testing.T) {

	var n uint = 1000
	var expected int = 1000
	received := workerpool.WorkerPool(n)
	assert.Equal(t, expected, received, "they should be equal")
}
