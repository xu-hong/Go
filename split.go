package main
import (
    "fmt"
    "strings"
)

func main() {
	//Fields splits the string s around each instance of one or more consecutive white space 
	fmt.Printf("%q\n", strings.Fields("a man a    plan	a canal panama"))
}