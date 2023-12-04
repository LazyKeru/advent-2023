package main

import (
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	number            int
	winning_numbers   []int
	scratched_numbers []int
}

func import_card_data(path string) []string {
	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(body), "\n")
}

func struct_cards(raw_cards []string) []card {
	var res []card
	for i := 0; i < len(raw_cards); i++ {
		tmp, err := string_to_card(raw_cards[i])
		if err != nil {
			log.Print(err)
		}
		res = append(res, tmp)
	}
	return res
}

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

func string_to_card(raw_card string) (card, error) {
	var res card
	start_card_regex, err := regexp.Compile("Card[ ]+[0-9]*:")
	if err != nil {
		log.Fatal(err)
	}
	number_regex, err := regexp.Compile("[1-9]+[0-9]*")
	if err != nil {
		log.Fatal(err)
	}
	winning_numbers, err := regexp.Compile("[0-9 ]*[|]")
	if err != nil {
		log.Fatal(err)
	}
	scratched_numbers, err := regexp.Compile("[|][0-9 ]*")
	if err != nil {
		log.Fatal(err)
	}
	res.number, err = strconv.Atoi(number_regex.FindString(start_card_regex.FindString(raw_card)))
	if err != nil {
		log.Fatal(err)
	}
	res.winning_numbers = string_list_to_int_list(number_regex.FindAllString(winning_numbers.FindString(raw_card), -1))
	res.scratched_numbers = string_list_to_int_list(number_regex.FindAllString(scratched_numbers.FindString(raw_card), -1))
	return res, err
}

func total_winning_pair(formated_card card) int {
	var winning_pair int
	for _, scratched_number := range formated_card.scratched_numbers {
		if slices.Contains(formated_card.winning_numbers, scratched_number) {
			winning_pair += 1
		}
	}
	return winning_pair
}

func total_points_cards(formated_cards []card) int {
	var res int
	for i := 0; i < len(formated_cards); i++ {
		winning_pair := total_winning_pair(formated_cards[i])
		if winning_pair == 0 {
			continue
		}
		tmp := 1
		for i := 0; i < winning_pair-1; i++ {
			tmp *= 2
		}
		res += tmp
	}
	return res
}

func scratch_card(n int, winning_pairs_cards []int) int {
	res := 1
	for i := 1; i <= winning_pairs_cards[n]; i++ {
		res += scratch_card(n+i, winning_pairs_cards)
	}
	return res
}

func total_scratch_cards(formated_cards []card) int {
	var scratched_cards int
	var winning_pairs_cards []int
	for _, formated_card := range formated_cards {
		winning_pairs_cards = append(winning_pairs_cards, total_winning_pair(formated_card))
	}
	for n := 0; n < len(formated_cards); n++ {
		scratched_cards += scratch_card(n, winning_pairs_cards)
	}
	return scratched_cards
}

func main() {
	raw_cards := import_card_data("input.txt")
	log.Print(raw_cards)
	formated_cards := struct_cards(raw_cards)
	log.Print(formated_cards)
	part1_result := total_points_cards(formated_cards)
	log.Print(part1_result)
	part2_result := total_scratch_cards(formated_cards)
	log.Print(part2_result)
}
