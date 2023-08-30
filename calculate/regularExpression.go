package calculate

import (
	"regexp"
)

func Regular(s string) (result string) {
	// 创建正则表达式模式
	pattern := `\s*=\s*`

	// 编译正则表达式
	regex := regexp.MustCompile(pattern)

	// 使用正则表达式替换字符串
	result = regex.ReplaceAllString(s, "=")

	//b := strings.Fields(result)
	return
}
