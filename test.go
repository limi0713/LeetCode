/*

ethanlong
void Split(const std::string &strTarget, char ch, std::vector<std::string> &vctStr);
ethanlong
用ch将strTarget分割成数组，保存到vctStr
*/

package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	// strTarget := "aabba" // aabaa baa aab
	// var ch byte = 'b'    // b
	// resLs := list.New()
	// Split(strTarget, byte(ch), resLs)

	// Print(resLs)

	strs := []string{"0", "123", "-123", "+234", "2147483648", "-2147483649", strconv.Itoa(1<<31 - 1), strconv.Itoa(-(1 << 31))}

	for i := range strs {
		res, err := atoi(strs[i])
		if err != nil {
			fmt.Printf("str:%s, err:%v\n", strs[i], err)
		} else {
			fmt.Printf("str:%s, res:%d\n", strs[i], res)
		}
	}
}

/*
实现atoi函数，如果超过了整形表示的范围，则返回错误
	int atoi(char *pData);
*/

func atoi(str string) (int, error) {

	if len(str) <= 0 {
		return 0, fmt.Errorf("str is nil")
	}

	var flag int64 = 1
	if str[0] == '-' {
		flag = -1
		str = str[1:]
	} else if str[0] == '+' {
		flag = 1
		str = str[1:]
	}

	var max int64 = 1<<31 - 1
	var min int64 = -(1 << 31)

	var res int64 = 0

	for i := range str {
		if str[i] < '0' || str[i] > '9' {
			return 0, fmt.Errorf("char error")
		} else {
			res = res*10 + int64(str[i]-'0')
			// if res > max {
			// 	return 0, fmt.Errorf("num is over max")
			// }
		}
	}

	res *= flag
	if res < min || res > max {
		return 0, fmt.Errorf("num is over max")
	}

	return int(res), nil
}

func Print(ls *list.List) {
	for ls.Len() > 0 {
		node := ls.Front()
		if str, ok := node.Value.(string); ok {
			fmt.Println(str)
		}
		ls.Remove(ls.Front())
	}
}

func Split(strTarget string, ch byte, ls *list.List) {

	indexBegin := 0
	for i := range strTarget {
		if strTarget[i] == ch {
			if i > indexBegin {
				ls.PushBack(strTarget[indexBegin:i])

			}
			indexBegin = i + 1
		}
	}

	if indexBegin < len(strTarget) {
		ls.PushBack(strTarget[indexBegin:])
	}
}
