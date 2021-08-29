package main

import (
	"fmt"
	"math/rand"

	"time"
)

func main(){
	intlist:=InitList(100000,[]int{})
	WrapTimer("bubble",intlist,BubbleSort)
	intlist=InitList(100000,[]int{})
	WrapTimer("select",intlist,SelectionSort)
	intlist=InitList(100000,[]int{})
	WrapTimer("insert",intlist,InsertionSort)
	intlist=InitList(100000,[]int{})
	WrapTimer("shell",intlist,ShellSort)
	intlist=InitList(100000,[]int{})
	WrapTimer("merge",intlist,MergeSort)
	intlist=InitList(100000,[]int{})
	//intlist=[]int{5,4,3,2,1,6,7,8,9}
	WrapTimer("quick",intlist,QuickSort)
}
func WrapTimer(funcName string,in []int,f func([]int)){

	now:=time.Now().UnixNano()
	defer func(){
		fmt.Println(funcName,":",(time.Now().UnixNano()-now)/1000000)
	}()
	f(in)


}
func InitList(num int,list []int)[]int{
	rand.Seed(time.Now().Unix())
	for i:=0;i<num;i++{
		tmp:=rand.Int31n(1000)
		list=append(list, int(tmp))

	}
	return list
}

func BubbleSort(list []int){
	//冒泡排序
	l:=len(list)
	if l==0{
		return
	}
	var tempInt int
	for i:=0;i<l;i++{
		//每次迭代相当于确定最后一个人元素位置

		for j := 0; j < l-1-i; j++ {
			//每一大轮都将迭代长度截短1
			if list[j]>list[j+1]{
				tempInt=list[j]
				list[j]=list[j+1]
				list[j+1]=tempInt
			}
		}

	}
	return
}
func SelectionSort(list []int){
	//选择
	l:=len(list)
	var minIndex int
	//最小元素的地址
	var tempInt int
	for i:=0;i<l-1;i++{
		minIndex=i
		for j:=i+1;j<l;j++{
			//找到区间内的最小元素
			if list[minIndex]>list[j]{
				minIndex=j
			}
		}
		if minIndex!=i{
			tempInt=list[i]
			list[i]=list[minIndex]
			list[minIndex]=tempInt
		}
	}
}
func InsertionSort(list []int){
	//插入
	l:=len(list)
	for i:=0;i<l;i++{
		tmp:=list[i]
		for j:=i-1;j>=0;j--{
			if list[j]>tmp{
				list[j+1]=list[j]
				if j==0{
					list[j]=tmp
				}
			}else{
				list[j+1]=tmp
				break
			}
		}
	}
	//fmt.Println(list)
}

func ShellSort(list []int){
	var tmp int
	l:=len(list)
	for i:=l/2;i>=1;i/=2{
		//i:间隔
		for j:=0;j<i;j++{
			//j:起始点
			for k:=j+i;k<l;k+=i{
				//k:游标
				tmp=list[k]
				for m:=k-i;m>0;m-=i{
					//m:比较游标
					if list[m]>tmp{
						list[m+i]=list[m]
						if m==j{
							list[m]=tmp
						}
					}else{
						list[m+i]=tmp
						break
					}
				}
			}
		}
	}
	InsertionSort(list)
	//fmt.Println(list)
}

func sort(list []int)[]int{
	result:=make([]int,0)
	l:=len(list)
	if l==1{
		return list
	}
	left:=list[:l/2]
	right:=list[l/2:]
	leftResult:=sort(left)
	rightResult:=sort(right)
	result=merge(leftResult,rightResult,l)
	return result

}
func merge(left []int,right []int,l int)[]int{
	result:=make([]int,l)

	//fmt.Println(l,left,right)
	leftCurfor:=0
	rightCursor:=0
	ll:=len(left)
	rl:=len(right)
	for i:=0;i<l;i++{
		//fmt.Println(i,leftCurfor,rightCursor,left,right)
		if leftCurfor==ll{
			copy(result[i:],right[rightCursor:])
			//fmt.Println("lcopy",result[i:],right[rightCursor:])
			break
		}
		if rightCursor==rl{
			copy(result[i:],left[leftCurfor:])
			//fmt.Println("rcopy",result[i:],left[leftCurfor:])
			break
		}
		//fmt.Println(i,leftCurfor,rightCursor,left,right)

		if left[leftCurfor]<right[rightCursor]{
			result[i]=left[leftCurfor]
			leftCurfor+=1
		}else{
			result[i]=right[rightCursor]
			rightCursor+=1
		}
	}

	return result
}

func MergeSort(list []int){
	//归并
	list=sort(list)
	//fmt.Println(list)
}

func quickSort(list []int,left int,right int) {
	//标准值
	key:=list[left]
	//左坐标
	i:=left
	//右坐标
	j:=right
	//空位所在的坐标，基准值已经转载到缓存key中
	p:=left
	//左右坐标不能越界
	for i<j{
		for i<=j{
			if list[j]<key{
				list[i]=list[j]
				p=j
				break

			}
			j--
		}

		for i<j{
			if list[i]>key{
				list[p]=list[i]
				p=i
				break

			}
			i++

		}
	}
	list[p]=key
	if p-left>1{
		quickSort(list,left,p-1)
	}
	if right-p>1{
		quickSort(list,p+1,right)
	}

}
func QuickSort(list []int){
	//快速排序
	quickSort(list,0,len(list)-1)
	//fmt.Println(list)
}