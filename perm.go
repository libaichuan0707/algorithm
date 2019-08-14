package algorithm

import "fmt"

//全排列算法

//例子:
//给定a,b,c  生成abc,acb,bac,bca,cab,cba

//递归算法
func PermRecursive(arr []string, beginIndex int) {
	if beginIndex == len(arr) {
		fmt.Println(arr)
		return
	}

	for i := beginIndex; i < len(arr); i++ {
		arr[beginIndex], arr[i] = arr[i], arr[beginIndex]
		PermRecursive(arr, i)
		arr[beginIndex], arr[i] = arr[i], arr[beginIndex]
	}
}


//非递归算法