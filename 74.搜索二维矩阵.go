/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0 || len(matrix[0]) <= 0{
		return false
	}

	start, end := 0, len(matrix)-1

	for start < end {
		mid := start + (end - start)/2

		if matrix[mid][0] == target {
			return true
		}else if matrix[mid][0] > target {
			end = mid - 1
		}else if matrix[mid][0] < target {
			if matrix[mid][len(matrix[mid])-1] < target {
				start = mid + 1
			}else if matrix[mid][len(matrix[mid])-1] > target {
				start, end = mid, mid
			}else {
				return true
			}
		}
	}

	row := start
	start, end = 0, len(matrix[row])-1

	for start <= end {
		mid := start + (end - start)/2
		if matrix[row][mid] == target {
			return true
		}else if matrix[row][mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}
// @lc code=end

