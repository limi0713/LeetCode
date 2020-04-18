// 给你数字 k ，请你返回和为 k 的斐波那契数字的最少数目，
//其中，每个斐波那契数字都可以被使用多次。
// 斐波那契数字定义为：
// F1 = 1
// F2 = 1
// Fn = Fn-1 + Fn-2 ， 其中 n > 2 。
// 数据保证对于给定的 k ，一定能找到可行解。

func findMinFibonacciNumbers(k int) int {
    arr := handle(k)
    
    return dp(arr, k)
}

func handle(k int) []int{
    arr := make([]int,0)
    arr = append(arr, 1,1)
    
    for i := 2; arr[i-1] < k;i ++ {
        arr = append(arr, arr[i-1]+arr[i-2])
    }
    
    return arr
}

func dp(arr []int, k int) int {
    
    ans := 0
    
    i := len(arr) - 1
    for k > 0 {
        if k >= arr[i] {
            ans ++
            k -= arr[i]
        }
        i --
    }
    
    return ans
}



// 一个 「开心字符串」定义为：
// 仅包含小写字母 ['a', 'b', 'c'].
// 对所有在 1 到 s.length - 1 之间的 i ，满足 s[i] != s[i + 1] （字符串的下标从 1 开始）。
// 比方说，字符串 "abc"，"ac"，"b" 和 "abcbabcbcb" 都是开心字符串，
// 但是 "aa"，"baa" 和 "ababbc" 都不是开心字符串。

// 给你两个整数 n 和 k ，你需要将长度为 n 的所有开心字符串按字典序排序。
// 请你返回排序后的第 k 个开心字符串，如果长度为 n 的开心字符串少于 k 个，那么请你返回 空字符串 。
class Solution {
public:
	int po[20];
	string getHappyString(int n, int k) {
		po[0]=1;
		for(int i=1;i<=n;++i) po[i]=po[i-1]*2;
		if(k>po[n-1]*3) return "";
		int l=0;
		string u;
		for(int i=1;i<=n;++i)
		{
			for(int j='a';j<='c';++j) if(l!=j)
			{
				if(k>po[n-i]) k-=po[n-i];
				else
				{
					u.push_back(j);
					l=j;
					break;
				}
			}
		}
		return u;
	}
};
	




// 某个程序本来应该输出一个整数数组。但是这个程序忘记输出空格了以致输出了一个数字字符串，
// 我们所知道的信息只有：
// 数组中所有整数都在 [1, k] 之间，且数组中的数字都没有前导 0 。

// 给你字符串 s 和整数 k 。可能会有多种不同的数组恢复结果。
// 按照上述程序，请你返回所有可能输出字符串 s 的数组方案数。
// 由于数组方案数可能会很大，请你返回它对 10^9 + 7 取余 后的结果。

func numberOfArrays(s string, k int) int {
    ans := make(map[string]int)
    ans[""] = 1
    return dp(&ans,s,k)
}

func dp(ans *map[string]int, s string, k int) int {
    if v, ok := *ans[s]; ok {
        return v
    }
    
    num := 0
    for i := range s {
        num = 10 *num + int(s[i]-'0')
        
        if i+1< len(s) && s[i+1] == '0' {
            continue
        }
        
        if num <= k {
            *ans[s[i+1:]] = dp(ans,s[i+1:], k)
            *ans[s] = (ans[s] + ans[s[i+1:]])%1000000007
        }else{
            break
        }
    }
    
    return ans[s]
}

class Solution {
public:
	int numberOfArrays(string s, int k) {
		const int MOD=1e9+7;
		int n=s.size();
		int*f=(int*)malloc(sizeof(int)*(n+1));
		memset(f,0,sizeof(int)*(n+1));
		f[0]=1;
		for(int i=0;i<n;++i) if(s[i]!='0')
		{
			long long u=0;
			for(int j=i;j<n;++j)
			{
				u=u*10+s[j]-48;
				if(u>k) break;
				(f[j+1]+=f[i])%=MOD;
			}
		}
		int ans=f[n];
		free(f);
		return ans;
	}
};