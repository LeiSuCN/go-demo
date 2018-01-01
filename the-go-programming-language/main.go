package main

import "fmt"
import "runtime"

type Msg struct{
	code int "int"
	message string "message"
}


func main() {
	fmt.Println("Start...")
	var msg Msg
	fmt.Println(msg)
	fmt.Println(runtime.NumCPU())
	fmt.Println("End!")
}
