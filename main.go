package main

import (
	"fmt"
	cli "timestamp-service/cmd"
)

func main() {
	err := cli.Execute()
	if err != nil {
		return
	}
	fmt.Println("flags: ", cli.CLI)
}
