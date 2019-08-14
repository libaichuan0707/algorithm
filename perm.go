package algorithm

import "fmt"

//全排列算法

//例子:
//给定a,b,c  生成abc,acb,bac,bca,cab,cba

//递归算法
//固定任意一个元素，求剩下的全排列

func PermRecursive(arr []string, beginIndex int) {
	if beginIndex == len(arr) - 1 {
		fmt.Println(arr)
		return
	}

	for i := beginIndex; i < len(arr); i++ {
		arr[beginIndex], arr[i] = arr[i], arr[beginIndex]
		PermRecursive(arr, beginIndex + 1)
		arr[beginIndex], arr[i] = arr[i], arr[beginIndex]
	}
}


//去重递归全排列
func IsPermRecursiveRepeat(arr []string, beginIndex int, endIndex int) bool{
	for temp := beginIndex; temp < endIndex; temp++ {
		if arr[temp] == arr[endIndex] {
			return true
		}
	}

	return false
}

func PermRecursiveNoRepeat(arr []string, beginIndex int) {
	if beginIndex == len(arr) - 1 {
		fmt.Println(arr)
		return
	}

	for i := beginIndex; i < len(arr); i++ {
		isRepeat := IsPermRecursiveRepeat(arr, beginIndex, i)
		if !isRepeat {
			arr[beginIndex], arr[i] = arr[i], arr[beginIndex]
			PermRecursiveNoRepeat(arr, beginIndex + 1)
			arr[beginIndex], arr[i] = arr[i], arr[beginIndex]
		}
	}
}


//非递归算法