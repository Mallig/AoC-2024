package day_1

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
	"slices"
	"math"
	"errors"	
)

func Complete() {
	fmt.Println("Solving Day 1")
	
	left_vals, right_vals := ReadInput()
	
	slices.Sort(left_vals)
	slices.Sort(right_vals)
	
	part_1, err := SolvePartOne(left_vals, right_vals)
	
	if err != nil {
		fmt.Errorf("Error ocurred while solving part 1: %e", err)
	}
	
	fmt.Printf("	Part 1: %d\n", part_1)
	
	part_2, err := SolvePartTwo(left_vals, right_vals)
	
	fmt.Printf("	Part 2: %d\n", part_2)

}

func SolvePartOne(left []int, right[]int) (int, error) {
	if !slices.IsSorted(left) || !slices.IsSorted(right) {
		return 0, errors.New("Input slices are not sorted")
	}
	var result float64 = 0
	
	for i:=0; i<len(left); i++ {
		x, y := left[i], right[i]
		
		result += math.Abs(float64(x - y))
	}
	
	return int(result), nil
}

func SolvePartTwo(left []int, right []int) (int, error) {
	if !slices.IsSorted(left) || !slices.IsSorted(right) {
		return 0, errors.New("Input slices are not sorted")
	}
	
	result := 0
	y := 0
	
	for i:=0; i<len(left); i++ {
		x := left[i]
		
		if i>0 && x == left[i-1] {
			result += (x*y)
			y = 0
			continue
		}
		
		y = 0
		
		for j:=y; j<len(right); j++ {
			if x < right[j] {
				break
			} else if x == right[j] {
				y += 1
			} else {
				continue
			}
		}
		
		slices.Delete(right, 0, y)
		
		result += (x*y)
	}
	
	return result, nil
}

func ReadInput() ([]int, []int) {
	f, err := os.Open("../inputs/day_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	
	left_vals := make([]int, 0)
	right_vals := make([]int, 0)
	
	for scanner.Scan() {
		left, right := ParseLine(scanner.Text())
		left_vals = append(left_vals, left)
		right_vals = append(right_vals, right)
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return left_vals, right_vals
}

func ParseLine(line string) (int, int) {
	vals := strings.Split(line, "   ")
	left, err := strconv.Atoi(vals[0])
	
	if err != nil {
		log.Fatal(err)
	}
	
	right, err := strconv.Atoi(vals[1])
	
	if err != nil {
		log.Fatal(err)
	}
	
	return left, right
}
