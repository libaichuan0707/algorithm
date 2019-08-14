package algorithm

import (
	"fmt"
	"testing"
)


func TestPermRecursive(t *testing.T) {

	total := []string{"a", "b","c", "d", "e"}

	for i := 1; i <= len(total); i++ {
		PermRecursive(total[:i], 0)
		fmt.Printf("\n\n")
	}

}

func TestPermRecursiveNoRepeat(t *testing.T) {

	total := []string{"a","a","b","b"}

	for i := 1; i <= len(total); i++ {
		PermRecursiveNoRepeat(total[:i], 0)
		fmt.Printf("\n\n")
	}

}