// MVCWebServer project main.go
package main

import (
	"Prj/MVCWebServer/common"
	"Prj/MVCWebServer/server"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("ServerStart\n")
	reader := bufio.NewReader(os.Stdin)
	sc := server.GetCurrentServer()
	sc.Start()
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Error:%s", err.Error())
			continue
		}
		command := string(line[:len(line)-2])
		fmt.Printf("Get Command:%s\n", command)
		if command == "stop" {
			break
		}
	}
}
