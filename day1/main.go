package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calibration_input_online() string {
	res, err := http.Get("https://adventofcode.com/2023/day/1/input")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func calibration_input_local() []string {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(body), "\n")
}

func decode(word string) int {
	var first_number byte
	var last_number byte
	for i := 0; i < len(word); i++ {
		if word[i] < 58 && word[i] > 47 {
			if first_number == 0 {
				first_number = word[i]
			} else {
				last_number = word[i]
			}
		}
	}
	if last_number == 0 {
		last_number = first_number
	}
	res, err := strconv.Atoi(string([]byte{first_number, last_number}))
	if err != nil {
		return 0
	}
	return res
}

func calibrations_decoded(input []string) []int {
	var calibration []int
	for i := 0; i < len(input); i++ {
		calibration = append(calibration, decode(input[i]))
	}
	return calibration
}

func spelled_out_to_int(input string) string {
	output := input
	written_number := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < len(written_number); i++ {
		r, _ := regexp.Compile(written_number[i])
		index := r.FindAllStringIndex(output, -1)
		for j := 0; j < len(index); j++ {
			output = output[:index[j][0]+2] + strconv.Itoa(i+1) + output[index[j][0]+2:]
		}
	}
	return output
}

// func sum_digit(number int) int {
// 	var res int
// 	digits := strconv.Itoa(number)
// 	log.Print(number)
// 	log.Print(digits)
// 	for i := 0; i < len(digits); i++ {
// 		if digits[i] < 58 && digits[i] > 47 {
// 			add, _ := strconv.Atoi(string([]byte{digits[i]}))
// 			res += add
// 		}
// 	}
// 	return res
// }

func main() {
	input := calibration_input_local()
	for i := 0; i < len(input); i++ {
		// log.Printf(input[i])
		input[i] = spelled_out_to_int(input[i])
		// log.Printf(input[i])
	}
	output := calibrations_decoded(input)
	var res int
	for i := 0; i < len(output); i++ {

		res += output[i]
	}
	log.Print(res)
}
