//查找重复的行
/*这段 Go 语言代码的主要功能是从标准输入读取文本行，
统计每行文本出现的次数，然后输出那些出现次数超过一次的行及其出现的次数。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
