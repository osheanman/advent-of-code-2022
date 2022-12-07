package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	food_list, err := os.Open("/input/food_list.txt")
	check(err)
	defer food_list.Close()

	f, err := os.Create("./output/calorie.txt")
	check(err)
	defer f.Close()

	food_scanner := bufio.NewScanner(food_list)
	food_scanner.Split(bufio.ScanLines)

	// elf_pack := make([]int, 0)
	// pack_size := 0
	max_cal, elf_cal := 0, 0

	for food_scanner.Scan() {
		line := food_scanner.Text()
		if line == "" {
			if max_cal < elf_cal {
				max_cal = elf_cal
			}
			elf_cal = 0
		} else {
			calorie, err := strconv.Atoi(line)
			check(err)
			elf_cal += calorie
		}
	}

	_, err = f.WriteString(fmt.Sprintf("%d\n", elf_cal))
	check(err)
}
