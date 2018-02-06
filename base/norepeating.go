/*
从字符串中寻找最长未重复的字符串长度

例：寻找最长不含有重复字符的子串
对于每一个字母a：
- lastOccurred[a]不存在，或者 < start -> 无操作
- lastOccurred[a] >= start -> 更新start
- 更新lastOccurred[x], 更新maxLength

 */
package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	// 不支持中文
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s){
		// 因为lastOccurred[ch] 如果ch不存在会返回0，要处理下
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			// 更新start
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i- start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLength
}

func lengthOfNonRepeatingSubStr02(s string) int {
	// 支持中文
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s){
		// 因为lastOccurred[ch] 如果ch不存在会返回0，要处理下
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			// 更新start
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i- start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	s1 := "abcelalsgjlad;gadfgjll"
	fmt.Println(s1, "最长未重复的是",
		lengthOfNonRepeatingSubStr(s1))

	s2 := "abcdefghi"
	fmt.Println(s2, "最长未重复的是",
		lengthOfNonRepeatingSubStr(s2))

	s3 := "bbbbbbb"
	fmt.Println(s3, "最长未重复的是",
		lengthOfNonRepeatingSubStr(s3))

	s4 := "中文字符哈哈哈哈"
	fmt.Println(s4, "01版本：最长未重复的是：",
		lengthOfNonRepeatingSubStr(s4))
	fmt.Println(s4, "02版本：最长未重复的是：",
		lengthOfNonRepeatingSubStr02(s4))
}
