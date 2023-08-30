package memory

import (
	"GolandProjects/addSum"
	"fmt"
)

func xxx(i *int64) int64 {
	*i++
	addSum.AddFunction(1, 2)
	var x int64 = 1
	p := &x // 地址
	fmt.Println(p)
	fmt.Println(*p)
	*p = 3 // 根据地址取值
	fmt.Println(xxx(p))
	fmt.Println(x)
	return *i
}
