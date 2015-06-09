package main
import (
    "fmt"
    "strings"
    "./evaluation"
)

func main() {
	//Fields splits the string s around each instance of one or more consecutive white space 
	fmt.Printf("%q\n", strings.Fields("a man a    plan	a canal panama"))

	al := []string{"a", "aa"}
	bl := []string{"b", "bb"}

	testlist := [][]string{al, bl}

	testlist = Shuffle(testlist)
	
	fmt.Println("%q", testlist)
}