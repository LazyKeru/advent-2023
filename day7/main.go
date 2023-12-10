package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const FIVE_OF_A_KIND = 6
const FOUR_OF_A_KIND = 5
const FULL_HOUSE = 4
const THREE_OF_A_KIND = 3
const TWO_PAIR = 2
const ONE_PAIR = 1
const HIGH_CARD = 0

type hands struct {
	cards     string
	bid       int
	hand_type int
}

func serializer(path string) []hands {
	var res []hands
	_body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	body := string(_body)
	cards, err := regexp.Compile("[0-9A-Z]*")
	number, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		temp := strings.Split(line, " ")
		hands_bid, _ := strconv.Atoi(number.FindString(temp[1]))
		hands_cards := cards.FindString(temp[0])
		res = append(res, hands{hands_cards, hands_bid, hand_type(hands_cards)})
	}

	return res
}

func hand_type(cards string) int {
	has_three_of_a_kind := false
	has_two_of_a_kind := false
	if strings.Count(cards, string(cards[0])) == 5 {
		return FIVE_OF_A_KIND
	}
	for i := 0; i < len(cards); i++ {
		if strings.Count(cards, string(cards[i])) == 4 {
			return FOUR_OF_A_KIND
		}
		if strings.Count(cards, string(cards[i])) == 3 {
			cards = strings.ReplaceAll(cards, string(cards[i]), "")
			i = 0
			has_three_of_a_kind = true
			continue
		}
		if strings.Count(cards, string(cards[i])) == 2 {
			cards = strings.ReplaceAll(cards, string(cards[i]), "")
			i = 0
			if has_three_of_a_kind {
				return FULL_HOUSE
			}
			if has_two_of_a_kind {
				return TWO_PAIR
			}
			has_two_of_a_kind = true
		}
	}
	if has_three_of_a_kind {
		if has_two_of_a_kind {
			return FULL_HOUSE
		}
		return THREE_OF_A_KIND
	}
	if has_two_of_a_kind {
		return ONE_PAIR
	}
	return HIGH_CARD
}

func to_int(card byte) int {
	res := 0
	switch string(card) {
	case "A":
		res = 14
	case "K":
		res = 13
	case "Q":
		res = 12
	case "J":
		res = 11
	case "T":
		res = 10
	default:
		res, _ = strconv.Atoi(string(card))
	}
	return res
}

func is_stronger(hand1 hands, hand2 hands) bool {
	if hand1.hand_type == hand2.hand_type {
		for i := 0; i < len(hand1.cards); i++ {
			if hand1.cards[i] == hand2.cards[i] {
				continue
			} else {
				return to_int(hand1.cards[i]) > to_int(hand2.cards[i])
			}
		}
	}
	if hand1.hand_type > hand2.hand_type {
		return true
	}
	return false
}

func sort_hands(unsorted_hands []hands) []hands {
	var sorted_hands []hands
	sorted_hands = append(sorted_hands, unsorted_hands[0])
	for i := 1; i < len(unsorted_hands); i++ {
		for j := 0; j < len(sorted_hands); j++ {
			if !is_stronger(unsorted_hands[i], sorted_hands[j]) {
				sorted_hands = append(sorted_hands[:j+1], sorted_hands[j:]...)
				sorted_hands[j] = unsorted_hands[i]
				break
			} else {
				if j == len(sorted_hands)-1 {
					sorted_hands = append(sorted_hands, unsorted_hands[i])
					break
				}
				continue
			}
		}
	}
	return sorted_hands
}

func score_part1(sorted_hands []hands) int {
	var res int
	for i := 0; i < len(sorted_hands); i++ {
		res += sorted_hands[i].bid * (i + 1)
	}
	return res
}

func main() {
	hands := serializer("input.txt")
	sorted_hands := sort_hands(hands)
	log.Print(score_part1(sorted_hands))

}
