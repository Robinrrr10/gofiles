package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting")
	file, err := os.OpenFile("myfile4.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.WriteString("this is my line\n")
	fmt.Println("End")
}
