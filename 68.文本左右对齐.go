/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] 文本左右对齐
 *
 * https://leetcode-cn.com/problems/text-justification/description/
 *
 * algorithms
 * Hard (42.93%)
 * Likes:    59
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 17.6K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text", "justification."]\n16'
 *
 * 给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
 * 
 * 你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
 * 
 * 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
 * 
 * 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
 * 
 * 说明:
 * 
 * 
 * 单词是指由非空格字符组成的字符序列。
 * 每个单词的长度大于 0，小于等于 maxWidth。
 * 输入单词数组 words 至少包含一个单词。
 * 
 * 
 * 示例:
 * 
 * 输入:
 * words = ["This", "is", "an", "example", "of", "text", "justification."]
 * maxWidth = 16
 * 输出:
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * words = ["What","must","be","acknowledgment","shall","be"]
 * maxWidth = 16
 * 输出:
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * 解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
 * 因为最后一行应为左对齐，而不是左右两端对齐。       
 * ⁠    第二行同样为左对齐，这是因为这行只包含一个单词。
 * 
 * 
 * 示例 3:
 * 
 * 输入:
 * words =
 * ["Science","is","what","we","understand","well","enough","to","explain",
 * "to","a","computer.","Art","is","everything","else","we","do"]
 * maxWidth = 20
 * 输出:
 * [
 * "Science  is  what we",
 * ⁠ "understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 * 
 * 
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	if len(words) <= 0 {
		return []string{getLengSpace(maxWidth)}
	}

	ans := make([]string, 0)
	
	index := 0
	leng := len(words[index])
	for i := 1;i < len(words); i ++ {
		if leng + 1 + len(words[i]) < maxWidth {
			leng = leng + 1 + len(words[i])
		}else if leng + 1 + len(words[i]) == maxWidth {
			ans = append(ans, getStr(&words,index,i,maxWidth,maxWidth),)
			index = i + 1
			leng = -1
		} else {
			ans = append(ans,getStr(&words,index,i-1,maxWidth,leng))
			index = i
			leng = -1
			i --
		}
	}

	if index < len(words) {
		ans = append(ans,getStr(&words,index,len(words) - 1,maxWidth,leng))
	}

	return ans
}

func getStr(words *[]string, start, end, maxWidth, curLen int) string {
	if curLen == maxWidth {
		return joinStrWithSpace(words,start,end,1)
	}

	if start == end {
		return (*words)[start] + getLengSpace(maxWidth-len((*words)[start]))
	}

	if end == len(*words) - 1{
		words := joinStrWithSpace(words,start,end,1)
		return words + getLengSpace(maxWidth -len(words))
	}

	wordsCount := end - start + 1
	spaceCount := maxWidth - curLen + wordsCount - 1
	minSpaceLen := spaceCount / (wordsCount - 1)
	maxSpaceCount := spaceCount % (wordsCount - 1)

	return joinStrWithSpace(words,start,start+maxSpaceCount,minSpaceLen+1) + getLengSpace(minSpaceLen) + joinStrWithSpace(words,start+maxSpaceCount+1,end,minSpaceLen)
}

func joinStrWithSpace(words *[]string, start, end, count int) string {
	res := (*words)[start]

	for i:=start+1;i<=end;i++{
		res += getLengSpace(count) + (*words)[i]
	}

	return res
}

func getLengSpace(leng int) string {
	res := ""

	for i:=0;i<leng;i++{
		res += " "
	}

	return res
}
// @lc code=end

