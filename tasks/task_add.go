package tasks

import (
	"fmt"
)

func init() {
	TaskMap["add"] = Add
	TaskMap["do"] = Do
}

func Add(data string) {
	fmt.Println("add--->", data)
}

func Do(data string) {
	fmt.Println("Do-->", data)
}
