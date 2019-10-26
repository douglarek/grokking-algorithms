package chap4

// Sum 4.1 请编写前述 sum 函数的代码
func Sum(a []int) int {
	n := len(a)

	if n == 0 {
		return 0
	}
	if n == 1 {
		return a[0]
	}

	return a[0] + Sum(a[1:]) // a[1:n]
}

// Len 4.2 编写一个递归函数来计算列表包含的元素数
func Len(a []int) int {
	n := len(a)

	if n == 0 {
		return 0
	}

	return 1 + len(a[1:])
}

//// 找出列表中最大的数字 ////

// Max 普通循环版
func Max(a []int) int {
	m := 0

	for _, v := range a {
		if v > m {
			m = v
		}
	}

	return m
}

// RecursiveMax 递归版
func RecursiveMax(a []int) int {
	return recursiveMax(a, 0)
}

func recursiveMax(a []int, m int) int {
	n := len(a)

	if n == 0 {
		return m
	}

	if a[0] > m {
		m = a[0]
	}

	return recursiveMax(a[1:], m)
}
