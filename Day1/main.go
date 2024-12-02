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

	fmt.Printf("list 1 index 0: %d list 1 last index: %d\n", list1[0], list1[len(list1)-1])
	fmt.Printf("list 2 index 0: %d list 2 last index: %d\n", list2[0], list2[len(list2)-1])
	//bytes, err := io.ReadAll(file)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(bytes))

	//r := csv.NewReader(file)
	//r.Comma = '\t'
	//records, err := r.ReadAll()
	//if err != nil {
	//	panic(err)
	//}
	//for i, record := range records {
	//	fmt.Println("row", i)
	//	fmt.Println("columns", record)
	//}
}

//There's just one problem: by holding the two lists up side by side (your puzzle input), it quickly becomes clear that the lists aren't very similar. Maybe you can help The Historians reconcile their lists?
//
//For example:
//
//3   4
//4   3
//2   5
//1   3
//3   9
//3   3
//Maybe the lists are only off by a small amount! To find out, pair up the numbers and measure how far apart they are. Pair up the smallest number in the left list with the smallest number in the right list, then the second-smallest left number with the second-smallest right number, and so on.
//
//Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those distances. For example, if you pair up a 3 from the left list with a 7 from the right list, the distance apart is 4; if you pair up a 9 with a 3, the distance apart is 6.
//
//In the example list above, the pairs and distances would be as follows:
//
//The smallest number in the left list is 1, and the smallest number in the right list is 3. The distance between them is 2.
//The second-smallest number in the left list is 2, and the second-smallest number in the right list is another 3. The distance between them is 1.
//The third-smallest number in both lists is 3, so the distance between them is 0.
//The next numbers to pair up are 3 and 4, a distance of 1.
//The fifth-smallest numbers in each list are 3 and 5, a distance of 2.
//Finally, the largest number in the left list is 4, while the largest number in the right list is 9; these are a distance 5 apart.
//To find the total distance between the left list and the right list, add up the distances between all of the pairs you found. In the example above, this is 2 + 1 + 0 + 1 + 2 + 5, a total distance of 11!
//
//Your actual left and right lists contain many location IDs. What is the total distance between your lists?

// this link is locked behind login
//https://adventofcode.com/2024/day/1/input
