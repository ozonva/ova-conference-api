package main

import (
	"fmt"
	"ova-conference-api/internal/domain"
	"ova-conference-api/internal/utils"
	"time"
)

func main() {
	fmt.Printf("Hello from %s %s", "ova-conference-api", "\n")
	openConfigInLoop()
	check()
}

func openConfigInLoop() {
	result := utils.ReadConfigInLoop("test_config.json", 10)
	fmt.Println(result)
}

func check() {
	checkDomainSplit()
	checkToMap()
}

func checkToMap() {
	count := 10
	fmt.Println("source:")
	conferences := GenerateConferences(count)
	fmt.Println(conferences)
	fmt.Println("Result:")
	mapByUserId, _ := utils.ToMapByUserId(conferences)
	fmt.Println(mapByUserId)
	fmt.Println("")

}

func GenerateConferences(count int) []domain.Conference {
	conferences := make([]domain.Conference, count)
	var counter uint64 = 0
	for counter = 0; counter < 10; counter++ {
		conferences[counter] = *domain.NewConference(counter, "TestConference", &domain.EventTime{Time: time.Now()})
	}
	return conferences
}

func checkDomainSplit() {
	count := 10
	conferences := GenerateConferences(count)

	fmt.Println("source:")
	fmt.Println(conferences)
	result, _ := utils.SplitToBulks(conferences, 3)
	fmt.Println("result:")
	fmt.Println(result)
	fmt.Println("")

}
