package main

func main() {
	s := "abc"
	for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
		println(s[i])
	}

	n := len(s)
	for n > 0 { // 替代 while (n > 0) {}
		println(s[n-1]) // 替代 for (; n > 0;) {}
		n--
	}

	for i, n := 0, length(s); i < n; i++ { // 避免多次调用 length 函数。
		println(i, s[i])
	}
}
func length(s string) int {
	println("call length.")
	return len(s)
}
