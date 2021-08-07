package main

import (
	"fmt"
	"ova-conference-api/internal/utils"
)

func main() {
	fmt.Printf("Hello from %s %s", "ova-conference-api", "\n")
	check()
}

func check() {
	checkSplit()
	checkSwap()
	checkDisjoin()
}

func checkDisjoin() {
	fmt.Println("Check disjoin")
	source := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	filter := []string{"1", "2", "3", "9"}
	result := utils.Disjoin(source, filter)
	fmt.Println("source:")
	fmt.Println(source)
	fmt.Println("filter:")
	fmt.Println(filter)
	fmt.Println("Result:")
	fmt.Println(result)
	fmt.Println("")
}

func checkSwap() {
	fmt.Println("Check swap")
	source := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	result := utils.SwapKeyValues(source)
	fmt.Println("source:")
	fmt.Println(source)
	fmt.Println("Result:")
	fmt.Println(result)
	fmt.Println("")
}

func checkSplit() {
	fmt.Println("Check split:")
	source := []int{1, 2, 3, 4, 5, 6}
	result := utils.Split(source, 5)
	fmt.Println("source:")
	fmt.Println(source)
	fmt.Println("Result:")
	fmt.Println(result)
	fmt.Println("")
}
