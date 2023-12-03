package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func string_int(input string) int {
	res, _ := strconv.Atoi(input)
	return res
}

func extract_engine_schematic(path string) []string {
	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(body), "\n")
}

func numbers_adjacent_symbol(input []string) []int {
	var res []int
	numbers_regex, err := regexp.Compile("[1-9]+[0-9]*")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(input); i++ {
		numbers_value := numbers_regex.FindAllString(input[i], -1)
		numbers_index := numbers_regex.FindAllStringIndex(input[i], -1)
		for k, number_index := range numbers_index {
			_left_limit := number_index[0]
			if _left_limit != 0 {
				_left_limit -= 1
				if string(input[i][_left_limit]) != "." {
					res = append(res, string_int(numbers_value[k]))
					continue
				}
			}
			_right_limit := number_index[1]
			if _right_limit == len(input[i]) {
				_right_limit -= 1
			} else {
				if string(input[i][_right_limit]) != "." {
					res = append(res, string_int(numbers_value[k]))
					continue
				}
			}
			for j := _left_limit; j <= _right_limit; j++ {
				switch {
				case i == 0:
					if string(input[i+1][j]) != "." {
						res = append(res, string_int(numbers_value[k]))
						break
					}
				case i == len(input)-1:
					if string(input[i-1][j]) != "." {
						res = append(res, string_int(numbers_value[k]))
						break
					}
				default:
					if string(input[i+1][j]) != "." {
						res = append(res, string_int(numbers_value[k]))
						break
					}
					if string(input[i-1][j]) != "." {
						res = append(res, string_int(numbers_value[k]))
						break
					}
				}
			}
		}
	}
	return res
}

func pairs_adjacent_star(input []string) []int {
	var res []int
	star, err := regexp.Compile("[*]")
	number_regex, err := regexp.Compile("[0-9]")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(input); i++ {
		found_stars := star.FindAllStringIndex(input[i], -1)
		for _, found_star := range found_stars {
			left_border := found_star[0] - 1
			right_border := found_star[0] + 1
			top_border := i - 1
			bottom_border := i + 1
			if found_star[0] == 0 {
				left_border = found_star[0]
			}
			if found_star[0] == len(input[i])-1 {
				right_border = found_star[0]
			}
			if i == 0 {
				top_border = i
			}
			if i == len(input)-1 {
				bottom_border = i
			}
			var _tmp_adjacent_numbers []int
			for j := top_border; j <= bottom_border; j++ {
				for k := left_border; k <= right_border; k++ {
					if number_regex.FindString(string(input[j][k])) != "" {
						_tmp := numbers_touching_border(input[j], left_border, right_border)
						_tmp_adjacent_numbers = append(_tmp_adjacent_numbers, _tmp...)
						break
					}
				}
			}
			if len(_tmp_adjacent_numbers) == 2 {
				res = append(res, _tmp_adjacent_numbers[0]*_tmp_adjacent_numbers[1])
			}
		}
	}
	return res
}

func numbers_touching_border(line string, left_border int, right_border int) []int {
	var res []int
	numbers_regex, err := regexp.Compile("[1-9]+[0-9]*")
	if err != nil {
		log.Fatal(err)
	}
	values := numbers_regex.FindAllString(line, -1)
	indexes := numbers_regex.FindAllStringIndex(line, -1)
	for i := 0; i < len(values); i++ {
		if left_border <= indexes[i][0] && indexes[i][0] <= right_border {
			res = append(res, string_int(values[i]))
			continue
		}
		if left_border <= indexes[i][1]-1 && indexes[i][1]-1 <= right_border {
			res = append(res, string_int(values[i]))
			continue
		}
		if indexes[i][0] < left_border && right_border < indexes[i][1]-1 {
			res = append(res, string_int(values[i]))
			continue
		}
	}
	return res
}

func sum(input []int) int {
	var res int
	for _, value := range input {
		res += value
	}
	return res
}

func main() {
	engine_schematic := extract_engine_schematic("input.txt")
	log.Print(engine_schematic)
	part_numbers := numbers_adjacent_symbol(engine_schematic)
	sum_parts := sum(part_numbers)
	log.Print(sum_parts)
	gear_ratios := pairs_adjacent_star(engine_schematic)
	sum_gear_ratios := sum(gear_ratios)
	log.Print(gear_ratios)
	log.Print(len(gear_ratios))
	log.Print(sum_gear_ratios)
}
