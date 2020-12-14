package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string //Here we use Name. Where it starts  with Upper case.
	Age      int    `json:"age"` //To use json variable starts with lower cast. use this
	Dept     string
	Subjects []string
}

func main() {
	fmt.Println("start")
	pojoTojson()
	jsonTopojo()
	fmt.Println("end")
}

func pojoTojson() {
	p1 := &Person{Name: "ram", Age: 23, Dept: "EEE", Subjects: []string{"Maths", "English", "Tamil"}}
	data, _ := json.Marshal(p1) //This will convert pojo to json
	fmt.Println(string(data))
}

func jsonTopojo() {
	data2 := `{"Name":"Kumar","age":15,"Dept":"ECE","Subjects":["Social","English","Tamil"]}`
	p2 := &Person{}
	json.Unmarshal([]byte(data2), p2) //This will convert json to pojo
	fmt.Println(p2.Name)
	fmt.Println(p2.Age)
	fmt.Println(p2.Subjects[1])
}
