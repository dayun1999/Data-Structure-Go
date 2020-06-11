package Sorting

//CountingSort的基本流程如下:
//目标数组: [2,4,5,2,4,6,2,1] 中最大元素为6,所以新建一个大小为6+1的数组count
//计数: 有1个1,3个2,1个4,1个5,1个6
//count数组如下: [0,1,3,0,1,1,1,1]
//局限性：目标数组元素只能为非负整数
func CountingSort(array []int) []int {
	out := make([]int, 0)
	//计数用的数组，大小为需要排序数组中最大元素值+1
	count := make([]int, max(array)+1)
	//遍历原始数组,将array里面的元素对应count数组里面的索引,array里面有重复的元素
	//那么count对应的索引的值就为元素的个数
	for _, e := range array {
		count[e]++
	}
	for index, item := range count {
		for i := 0; i < item; i++ {
			out = append(out, index)
		}
	}
	return out
}

//max return the max element of target array
func max(arr []int) int {
	max := arr[0]
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	return max
}
