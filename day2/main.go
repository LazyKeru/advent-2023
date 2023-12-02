package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type games struct {
	list []game
}

type game struct {
	key  int
	sets []set
}

type set struct {
	pulled_cubes []hand
}

type hand struct {
	cube_color string
	number     int
}

func load_games(path string) games {
	var _games games
	identification, err := regexp.Compile("Game ([1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9]):")
	sets, err := regexp.Compile("(([0-9]+) ([0-9 a-z,]*)[^;])")
	pulled, err := regexp.Compile("([0-9]+) ((red)|(blue)|(green))")
	number, err := regexp.Compile("[0-9]+")
	color, err := regexp.Compile("(red)|(blue)|(green)")
	if err != nil {
		log.Fatal(err)
	}
	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(body), "\n")
	for i := 0; i < len(lines); i++ {
		log.Print(lines[i])
		_identification := identification.FindString(lines[i])
		_id, err := strconv.Atoi(number.FindString(_identification))
		if err != nil {
			log.Fatal(err)
		}
		log.Print("For the Game number " + strconv.Itoa(_id) + "this was the following sets:")
		_sets := sets.FindAllString(lines[i], -1)
		var _game game
		_game.key = _id
		for j := 0; j < len(_sets); j++ {
			log.Print("For the set " + strconv.Itoa(j) + " the elf picked up :")
			var _set set
			_pulled := pulled.FindAllString(_sets[j], -1)
			for k := 0; k < len(_pulled); k++ {
				var _hand hand
				_times, err := strconv.Atoi(number.FindString(_pulled[k]))
				if err != nil {
					log.Fatal(err)
				}
				_color := color.FindString(_pulled[k])
				log.Print("Picked " + strconv.Itoa(_times) + " times the " + _color + " cube")
				_hand.cube_color = _color
				_hand.number = _times
				_set.pulled_cubes = append(_set.pulled_cubes, _hand)
			}
			_game.sets = append(_game.sets, _set)
		}
		_games.list = append(_games.list, _game)
	}
	return _games
}

func is_hand_valid(input hand, red_cubes int, green_cubes int, blue_cubes int) bool {
	var max_cubes int
	switch _color := input.cube_color; _color {
	case "red":
		max_cubes = red_cubes
	case "green":
		max_cubes = green_cubes
	case "blue":
		max_cubes = blue_cubes
	}
	if input.number <= max_cubes {
		return true
	}
	return false
}

func is_set_valid(input set, red_cubes int, green_cubes int, blue_cubes int) bool {
	for i := 0; i < len(input.pulled_cubes); i++ {
		if is_hand_valid(input.pulled_cubes[i], red_cubes, green_cubes, blue_cubes) == false {
			return false
		}
	}
	return true
}

func is_game_valid(input game, red_cubes int, green_cubes int, blue_cubes int) bool {
	for i := 0; i < len(input.sets); i++ {
		if is_set_valid(input.sets[i], red_cubes, green_cubes, blue_cubes) == false {
			return false
		}
	}
	return true
}

func addition_id_valid_games(input games, red_cubes int, green_cubes int, blue_cubes int) int {
	res := 0
	for i := 0; i < len(input.list); i++ {
		if is_game_valid(input.list[i], red_cubes, green_cubes, blue_cubes) {
			res += input.list[i].key
		}
	}
	return res
}

func sum_power_game(input game) int {
	var min_red, min_green, min_blue int
	for _, _set := range input.sets {
		for _, _hand := range _set.pulled_cubes {
			switch _hand.cube_color {
			case "red":
				if _hand.number > min_red {
					min_red = _hand.number
				}
			case "green":
				if _hand.number > min_green {
					min_green = _hand.number
				}
			case "blue":
				if _hand.number > min_blue {
					min_blue = _hand.number
				}
			}
		}
	}
	return min_red * min_green * min_blue
}

func sum_power_games(input games) int {
	var res int
	for i := 0; i < len(input.list); i++ {
		res += sum_power_game(input.list[i])
	}
	return res
}

func main() {
	_games := load_games("input.txt")
	first_answer := addition_id_valid_games(_games, 12, 13, 14)
	log.Print(first_answer)
	second_answer := sum_power_games(_games)
	log.Print(second_answer)
}
