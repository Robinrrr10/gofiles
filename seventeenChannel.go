package main

import "fmt"

func main() {
	mychannel := make(chan string)
	go func() {
		mychannel <- "hello"
		mychannel <- "world"
	}()
	fmt.Println(<-mychannel)
	fmt.Println(<-mychannel)
}
