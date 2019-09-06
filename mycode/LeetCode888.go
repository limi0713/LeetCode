package main

import "fmt"

func fairCandySwap(A []int, B []int) []int {
	countA := 0
	for _,i:=range A{
		countA+=i
	}
	countB := 0
	for _,i:=range B{
		countB+=i
	}
	res := make([]int,2,2)
	arg := (countA+countB)/2
	for _,i:=range A{
		for _,j:=range B{
			if arg == countA+j-i{
				res[0]=i
				res[1]=j
				return res
			}
		}
	}
	return res
}

func main(){
	A := []int{1,2,5}
	B := []int{2,4}
	res := fairCandySwap(A, B)
	fmt.Println(res[0],res[1])
}