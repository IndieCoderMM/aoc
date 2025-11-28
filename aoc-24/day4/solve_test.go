package main

import (
	"testing"

	"github.com/IndieCoderMM/aoc-24/utils"
)

func TestSolve(t *testing.T) {
	file := utils.GetFilePath(4)
	expected := utils.GetAnswer(4)

	t.Run("Day-4 Solution", func(t *testing.T) {
		ans := Solve(file)
		if ans != expected {
			t.Errorf("Expected %d, got %d", expected, ans)
		}
	})
}
