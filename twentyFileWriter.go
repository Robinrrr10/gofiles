package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	message := []byte("This is my fileeeee")
	err := ioutil.WriteFile("myfile1.txt", message, 0644) //To write in existing file
	if err != nil {
		fmt.Println(err)
	}

	//method2
	f, err := os.Create("myfile3.txt") //To create a file
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString("My content in new file.")
	f.Sync()

	w := bufio.NewWriter(f) //adding existing file which has content
	w.WriteString("appending my content in existing content")
	w.Flush()

}
