package main

import (
	"fmt"
	"ova-conference-api/internal/utils"
)

func main() {
	fmt.Printf("Hello from %s %s", "ova-conference-api", "\n")
	openConfigInLoop()

}

func openConfigInLoop() {
	result := utils.ReadConfigInLoop("test_config.json", 10)
	fmt.Println(result)
}
