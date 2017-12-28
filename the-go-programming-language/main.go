package main
import "fmt" 


func main(){
	fmt.Println("Hello") 
	defer fmt.Println(1)
    fmt.Println("World")
	defer fmt.Println(2)
}
