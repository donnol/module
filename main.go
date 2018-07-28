package main

import (
	"fmt"

	"github.com/donnol/module/a"
)

// Main 包的导出变量
var Main = "this is main."

func main() {
	fmt.Println("hello,", Main, a.A)
}
