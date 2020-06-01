//package SqList实现了顺序储存结构(顺序表)
package SqList

import (
	. "dataStructureGo/TextBook/01_PreDefinition" //为了使用status.go里面预定义的变量
)

//宏定义
const (
	ListInitSize  = 100 //顺序表储存空间的初始大小
	ListIncrement = 10  //顺序表储存空间的分配增量
)

//顺序表元素的类型定义
type ElemType int //类型定义

//顺序表结构
type SqList struct {
	element  *[]ElemType //顺序表储存空间的首地址
	length   int         //当前顺序表的长度,默认为0
	capacity int         //当前顺序表的储存容量,默认为0
}

//下面是基于顺序表的一系列操作

//InitList是初始化顺序表,return OK表示初始化成功，Error表示失败
func InitList() (*SqList, Status) {
	//分配指定的容量
	sqlist := new(SqList)
	list := make([]ElemType, ListInitSize, ListInitSize)
	if sqlist.element = &list; sqlist.element == nil {
		return nil, Error
	}
	sqlist.capacity = ListInitSize
	return sqlist, OK
}

//DestroyList是删除顺序表(释放内存)
func DestroyList(L *SqList) Status {
	//确保顺序表存在
	if L == nil || L.element == nil {
		return Error
	}
	//释放顺序表的内存
	L = nil
	return OK
}

//ClearList只是清空顺序表里面的数据(不释放内存)
func ClearList(L *SqList) Status {
	//确保顺序表的存在
	if L == nil || L.element == nil {
		return Error
	}
	L.length = 0
	return OK
}

//IsEmpty 判断顺序表里面有无数据
func IsEmpty(L *SqList) Status {
	if L.length == 0 {
		return OK
	}
	return Error
}

//ListLength 返回顺序表中有效元素的数量
func ListLength(L *SqList) int {
	return L.length
}

//GetElem 根据索引获取顺序表中指定元素
func GetElem(L *SqList, index int) (ElemType, Status) {
	//索引的有效范围是[1,L.length]
	if index < 1 || index > L.length {
		return 0, Error
	}
	return (*L.element)[index-1], OK
}

//LocateElem 返回顺序表中首个与e满足Compare关系的元素的位置索引

type compare func(a, b ElemType) bool

//var compareGreater = func(a, b ElemType) bool {
//	return a > b
//}

func LocateElem(L SqList, e ElemType, compareFunc compare) (ElemType, Status) {
	//确保顺序表的存在
	if L.element == nil {
		return 0, Error
	}
	p, i := L.element, 1
	for i < L.length && compareFunc(e, (*p)[i]) {
		i++
	}
	if i <= L.length {
		return ElemType(i), OK
	}
	return 0, Error
}

//PriorElem 返回给定参数的前驱(就是前一个元素)
func PriorElem(L SqList, e ElemType) (ElemType, Status) {
	//确保顺序表存在而且至少包含两个元素
	if L.element == nil || L.length <= 2 {
		return 0, Error
	}
	i := 0
	for i < L.length-1 && (*L.element)[i] != e {
		i++
	}
	if i == 0 || i >= L.length {
		return 0, Error
	}
	return (*L.element)[i-1], OK
}
