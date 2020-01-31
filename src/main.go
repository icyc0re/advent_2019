package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
)

var (
	inputFile = flag.String("input", "input/day1.txt", "File with input of day1")
)

func loadInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var masses []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		masses = append(masses, mass)
	}

	return masses
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func fuelCount(mass int) int {
	return max(0, (mass / 3) - 2)
}

func countFuel(modulesMass []int) int {
	res := 0
	for _, mass := range modulesMass {
		res += fuelCount(mass)
	}
	return res
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	modulesMass := loadInput(*inputFile)

	log.Println(countFuel(modulesMass))
}
