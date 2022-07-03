package main

import (
	"fmt"
	cli "timestamp-service/cmd"
	"timestamp-service/server"
)

func main() {
	err := cli.Execute()
	if err != nil {
		return
	}
	fmt.Println("flags: ", cli.CLI)
	router := server.InitServer()
	router.Run(fmt.Sprintf("%v:%v", cli.CLI.Address, cli.CLI.Port))
}
