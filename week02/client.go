package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golangwork/week02/service"
)

func main() {
	user, err := service.GetPersonById(1)
	if err != nil {
		fmt.Printf("根因：%T\n", errors.Cause(err))
		fmt.Printf("调用栈：%+v\n", err)
	}

	if len(user) > 0 {
		fmt.Println(user[0])
	}

	fmt.Println("over...")
}
