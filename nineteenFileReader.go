package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("myfile1.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	//method 2, to read only few bytes
	data2, err2 := os.Open("myfile2.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	reader := bufio.NewReader(data2)
	mydata2, err3 := reader.Peek(10) //take only 10 bytes
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(string(mydata2))
}
