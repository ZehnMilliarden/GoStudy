package example

import "fmt"

func init() {
	fmt.Println("example module init function")
}

func Dumpname(name *string) bool {
	/* 这是我的第一个简单的程序 */
	fmt.Println(*name)
	return true
}
