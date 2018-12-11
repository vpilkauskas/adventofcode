package main

import (
	"flag"
	"log"
	"time"

	"github.com/vpilkauskas/adventofcode/day9"
	"github.com/vpilkauskas/adventofcode/solver"
)

func main() {
	day := flag.Int("day", 9, "day")
	input := flag.String("input", "446 players; last marble is worth 7152200 points", "task input")
	flag.Parse()

	client := getClient(*day, *input)
	if client == nil {
		log.Fatal("Cannot create client with these parameters", *day, *input)
	}

	startTime := time.Now()

	result := client.Solve()

	elapsedTime := time.Since(startTime)

	log.Println(result, elapsedTime)
}

func getClient(day int, input string) solver.Client {
	switch day {
	case 9:
		return day9.New(input)
	}
	return nil
}
