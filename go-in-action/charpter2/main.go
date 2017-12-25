
package main



import (
    "fmt"
	"log"
	"os"
	_ "./search"
)

func init(){
	log.SetOutput(os.Stdout)
}

func main(){
    fmt.Println("Hello World!")
}

