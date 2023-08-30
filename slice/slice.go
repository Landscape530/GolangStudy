package slice

import "fmt"

func y() {
	slice := make([]int, 3, 3)
	slice[0] = 1
	fmt.Println(slice)
	fmt.Println(&slice[0])
	fmt.Println("初始长度:", len(slice))
	fmt.Println("初始容量:", cap(slice))

	slice = append(slice, 1)
	fmt.Println(&slice[0])
	fmt.Println("初始长度:", len(slice))
	fmt.Println("初始容量:", cap(slice))

	news := slice
	fmt.Println(news)
	fmt.Println(&news[0])

	slice[0] = 5
	fmt.Println(news)
}
