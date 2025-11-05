package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var (
	lettersLower = []rune("abcdefghijklmnopqrstuvwxyz")
	lettersUpper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func main() {
	fmt.Println("Starting...")
	start := time.Now()
	const iterations = 100000000
	states := make(map[string]string, iterations)

	for i := 0; i < iterations; i++ {
		key := createRandomString("upper", 3)
		value := createRandomString("lower", 10)
		states[key] = value
	}

	printSortedMap(states)

	fmt.Println("Finished in", time.Since(start))
}

func createRandomString(choice string, numOfLetters int) string {
	var letters []rune
	switch choice {
	case "upper":
		letters = lettersUpper
	default:
		letters = lettersLower
	}

	randomString := make([]rune, numOfLetters)
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))]
	}
	return string(randomString)
}

func printSortedMap(mapToBePrinted map[string]string) {
	keys := make([]string, 0, len(mapToBePrinted))
	for k := range mapToBePrinted {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("The key is %-5s and the value is %s\n", k, mapToBePrinted[k])
	}
}
