package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	food_list, err := os.Open("./input/food_list.txt")
	check(err)
	defer food_list.Close()

	f, err := os.Create("./output/calorie.txt")
	check(err)
	defer f.Close()

	food_scanner := bufio.NewScanner(food_list)
	food_scanner.Split(bufio.ScanLines)

	current_cal := 0
	elf_list := make([]int, 3)

	for food_scanner.Scan() {
		line := food_scanner.Text()
		if line == "" {
			sort.Ints(elf_list)
			for i := 0; i < 3; i++ {
				if current_cal > elf_list[i] {
					elf_list[i] = current_cal
					break
				}
			}
			current_cal = 0
		} else {
			calorie, err := strconv.Atoi(line)
			check(err)
			current_cal += calorie
		}
	}

	max_cal := elf_list[0] + elf_list[1] + elf_list[2]

	_, err = f.WriteString(fmt.Sprintf("%d\n", max_cal))
	check(err)
}
