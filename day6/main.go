package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func string_list_to_int_list(list []string) []int {
	var res []int
	for i := 0; i < len(list); i++ {
		tmp, err := strconv.Atoi(list[i])
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, tmp)
	}
	return res
}

func serializer(path string) ([]int, []int) {
	_body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	body := string(_body)
	number, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(body, "\n")
	input_time := string_list_to_int_list(number.FindAllString(lines[0], -1))
	input_distance := string_list_to_int_list(number.FindAllString(lines[1], -1))
	return input_time, input_distance
}

func boat_distance(time_race int, time_holding_button int) int {
	return (time_race - time_holding_button) * time_holding_button
}

func number_of_ways_to_win(time int, record int) int {
	var res int
	for i := 1; i < time; i++ {
		if boat_distance(time, i) > record {
			res++
		}
	}
	return res
}

func times_number_of_ways_to_win(times []int, records []int) int {
	res := 1
	for i := 0; i < len(times); i++ {
		res *= number_of_ways_to_win(times[i], records[i])
	}
	return res
}

func list_to_big(list []int) int {
	tmp := ""
	for i := 0; i < len(list); i++ {
		tmp += strconv.Itoa(list[i])
	}
	res, err := strconv.Atoi(tmp)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func main() {
	times, distances := serializer("input.txt")
	log.Print(times_number_of_ways_to_win(times, distances))
	log.Print(number_of_ways_to_win(list_to_big(times), list_to_big(distances)))
}
