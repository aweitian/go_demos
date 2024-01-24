//https://blog.csdn.net/qq_51721904/article/details/128231308
package main    //把test.go文件归属到main

import (
	"fmt"    //引入包fmt
	"taw.me/greetings"
)

func main() {
	//输出内容
	//fmt.Println("hello, My name is Eastmount!")
	fmt.Println(greetings.Hello("gondar"))
	fmt.Println("done!")
}