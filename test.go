//Go 程序的执行（程序启动）顺序如下：
//
//按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
//如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
//然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
//在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。

package main

import "fmt"

const c = "C"

var v = 5

type T struct{}

func init() { // initialization of package
}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	//...
}

func Func1() { // exported function Func1
	//...
}
