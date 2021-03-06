package _03_二维数组中的查找


// 从右上或者左下开始查找
func Search(array [][]int, target int) bool{
	row := len(array)
	col := len(array[0])

	i := 0
	j := col -1
	for i<row && j>=0{
		tmp := array[i][j]
		if tmp == target{
			return true
		}
		if tmp < target{
			i++
		}
		if tmp > target{
			j--
		}

	}
	return false
}
