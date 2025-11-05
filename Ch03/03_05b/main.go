package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("Starting ...")
	nowTime := time.Now()
	states := make(map[string]string)
	counter := 0
	for {
		states[createRandomStrings("upper", 3)] = createRandomStrings("lower", 10)


		counter++
		if counter == 100 {
			break
		}
	}
	finishedTime := time.Now()
	// getKeyAndValue(states)
	fmt.Println("Finished in", finishedTime.Sub(nowTime))
	// delete(states,"CA")
	sortedStatesMap := sortMap(states)
	// fmt.Println("sortedStatesMap", sortedStatesMap)
	getKeyAndValue(sortedStatesMap)
}

func createRandomStrings (choice string, lenOfString int) string{
	lettersLower := []rune("abcdefghijklmnopqrstuvwxyz")
	lettersUpper := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomString := make([]rune, lenOfString)
	var letters []rune
	switch choice {
	case "upper":
		letters = lettersUpper
	case "lower":
		letters = lettersLower
	default:
		letters = lettersLower
	}
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))]
	}
	// fmt.Println(randomString, string(randomString))
	return string(randomString)
}



func getKeyAndValue(states map[string]string)  {
	for key, value := range states {
		fmt.Println("the key is", key, "and the value is", value) 
	}
}



func sortMap(mapToBeSorted map[string]string)map[string]string  {
	keys := make([]string, len(mapToBeSorted))
	i := 0
	for k := range mapToBeSorted {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	newSortedMap := make(map[string]string)
	for i:= range keys {
		fmt.Println(keys[i])
		newSortedMap[keys[i]] = mapToBeSorted[keys[i]]
	}
	fmt.Println("sortedStatesMap", newSortedMap)
	return newSortedMap
}