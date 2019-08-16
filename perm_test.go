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

func TestNextPerm(t *testing.T) {
	arr := []int{2,3,5,1}
	fmt.Println(NextPerm(arr))
}

func TestPermNoRecursive(t *testing.T) {
	arr := []int{2, 1, 1}
	PermNoRecursive(arr)
}


func TestPermRecursiveOther(t *testing.T) {
	total := []string{"a", "b","c"}
	tempTotal:= make([]string, len(total), len(total))
	used := map[string]bool{}

	PermRecursiveOther(total, tempTotal, used, 0)
}