package main

import (
	"fmt"
	"github.com/msbu-tech/go-sf/framework"
)

func main() {
	fmt.Print("new app: \n")
	err := framework.NewApp("demo", "template")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success!\n")
}
