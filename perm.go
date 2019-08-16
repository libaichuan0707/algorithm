package algorithm

import (
	"fmt"
	"sort"
)

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

//递归的另外一种实现
//temp存当前这次从根节点到子节点，used存当前使用的字符,beginindex是树的层数
func PermRecursiveOther(arr []string, temp []string, used map[string]bool, beginIndex int) {
	if beginIndex == len(arr) {
		fmt.Println(temp)
		return
	}

	for i := 0; i < len(arr); i++ {
		curr := arr[i]
		if isUsed, ok := used[curr]; ok && isUsed {
			continue
		}

		temp[beginIndex] = curr
		used[curr] = true

		PermRecursiveOther(arr, temp, used, beginIndex + 1)

		used[curr] = false
	}
}

//非递归算法


//从小到大排序->获得比自己大的一个数

//怎么获得相同数组比当前这个数字大的数?

//找到最大的指数k，使得nums[k] < nums[k + 1]。如果不存在这样的索引，只需反转nums并完成。
//找到最大的指数l > k，使得nums[k] < nums[l]。
//交换nums[k]和nums[l]。
//反转子阵列nums[k + 1:]。

func ReverseSlice(nums []int, beginIndex int) {
	endIndex := len(nums) - 1

	for {
		if beginIndex >= endIndex {
			break
		}

		nums[beginIndex], nums[endIndex] = nums[endIndex], nums[beginIndex]
		beginIndex++
		endIndex--
	}
}

func NextPerm(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	findIndex := -1
	endIndex := len(nums) - 1

	for tempIndex := endIndex; tempIndex >= 1; tempIndex-- {
		if nums[tempIndex - 1] < nums[tempIndex] {
			findIndex = tempIndex - 1
			break
		}
	}

	if findIndex == -1 {
		ReverseSlice(nums, 0)
	}else{
		for i := endIndex; i > findIndex; i-- {
			if nums[i] > nums[findIndex] {
				nums[i], nums[findIndex] = nums[findIndex], nums[i]
				break
			}
		}

		ReverseSlice(nums, findIndex + 1)
	}

	if findIndex == -1 {
		return false
	}else{
		return true
	}
}

func PermNoRecursive(nums []int) {
	sort.Ints(nums)

	fmt.Println(nums)
	needBreak := true

	for {
		needBreak = NextPerm(nums)
		fmt.Println(nums)

		if !needBreak {
			break
		}
	}
}
