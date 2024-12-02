package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day1/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	list1 := []int{}
	list2 := []int{}
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "   ")
		item1, err := strconv.Atoi(items[0])
		if err != nil {
			panic(errors.New("failed to convert item 1 to int"))
		}
		item2, err := strconv.Atoi(items[1])
		if err != nil {
			panic(errors.New("failed to convert item 2 to int"))
		}

		list1 = append(list1, item1)
		list2 = append(list2, item2)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input: ", err)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	totalDistance := 0
	for i, entry := range list1 {
		list2Item := list2[i]
		diff := entry - list2Item
		if diff < 0 {
			diff = diff * -1
		}
		totalDistance += diff
	}
	fmt.Println(totalDistance)
}
