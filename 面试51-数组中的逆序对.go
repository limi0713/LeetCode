// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

 

// 示例 1:

// 输入: [7,5,6,4]
// 输出: 5

func reversePairs(nums []int) int {
    ans := 0
    
    merge(&nums,0,len(nums)-1,&ans)
    
    return ans
    
}

func merge(nums *[]int,start,end int,ans *int){
    if start >= end {
        return
    }
    
    if start + 1 == end {
        if (*nums)[start] > (*nums)[end] {
            *ans += 1
            swap(nums, start,end)
        }
        return
    }
    
    mid := start + (end - start)/2
    merge(nums,start,mid,ans)
    merge(nums,mid+1,end,ans)
    
    i,j := mid, end
    
    temp := make([]int,end-start+1)
    indexOfTemp := len(temp) -1
    
    for i >= start && j > mid {
        if (*nums)[i] > (*nums)[j] {
            *ans += j - mid
            temp[indexOfTemp] = (*nums)[i]
            i --
            
        }else {
            temp[indexOfTemp] = (*nums)[j]
            j --
        }
        
        indexOfTemp --
    }
    
    for i >= start{
        temp[indexOfTemp] = (*nums)[i]
        i --
        indexOfTemp --
    }
    
    for j > mid {
        temp[indexOfTemp] = (*nums)[j]
        j --
        indexOfTemp -- 
        // *ans += mid - start + 1
    }
    
    for i := range temp {
        (*nums)[start+i] = temp[i]
    }
}

func swap(nums *[]int,i,j int){
    (*nums)[i],(*nums)[j] = (*nums)[j],(*nums)[i]
}
