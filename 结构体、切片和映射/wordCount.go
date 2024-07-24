// 练习：映射
// 实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。 函数 wc.Test 会为此函数执行一系列测试用例，并输出成功还是失败。

// 你会发现 strings.Fields 很有用。

package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	wordMap := make(map[string]int)

	for _, word := range arr {
		if count, ok := wordMap[word]; ok {
			wordMap[word] = count + 1
		} else {
			wordMap[word] = 1
		}
	}

	fmt.Println(wordMap)

	return wordMap
}

func main_wordCount() {
	wc.Test(WordCount)
}

// map[Go!:1 I:1 am:1 learning:1]
// PASS
//  f("I am learning Go!") =
//   map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
// map[The:1 brown:1 dog.:1 fox:1 jumped:1 lazy:1 over:1 quick:1 the:1]
// PASS
//  f("The quick brown fox jumped over the lazy dog.") =
//   map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
// map[I:2 Then:1 a:1 another:1 ate:2 donut.:2]
// PASS
//  f("I ate a donut. Then I ate another donut.") =
//   map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
// map[A:1 a:2 canal:1 man:1 panama.:1 plan:1]
// PASS
//  f("A man a plan a canal panama.") =
//   map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
