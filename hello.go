package main

import "fmt"

const (
	messageConst = "Hello, Go!"
)

var (
	messageVar   = "Hello, from var"
	messageEmpty string
)

func main() {
	fmt.Println(messageVar)
	fmt.Println(messageEmpty)
	fmt.Println(messageConst)
}
func init() {
	messageEmpty = "init string"

}
